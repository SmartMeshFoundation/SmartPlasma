package block

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartmeshfoundation/smartplasma/blockchan/transaction"
	"github.com/smartmeshfoundation/smartplasma/merkle"
)

// Default value to depth.
const (
	DefaultDepth = 257
)

var (
	depth257 = big.NewInt(DefaultDepth)
)

// TODO: check memory synchronization

// Block object.
type Block struct {
	mtx  sync.Mutex
	uids []string
	txs  map[string]*transaction.Transaction
	tree *merkle.Tree

	built bool
}

// NewBlock creates new block in memory.
func NewBlock() *Block {
	return &Block{
		mtx: sync.Mutex{},
		txs: make(map[string]*transaction.Transaction),
	}
}

// Hash returns block hash.
func (bl *Block) Hash() common.Hash {
	if !bl.built {
		return common.Hash{}
	}
	return bl.tree.Root()
}

// AddTx adds a transaction to the block.
func (bl *Block) AddTx(tx *transaction.Transaction) error {
	bl.mtx.Lock()
	bl.uids = append(bl.uids, tx.UID().String())
	sort.Strings(bl.uids)
	bl.mtx.Unlock()

	bl.txs[tx.UID().String()] = tx
	return nil
}

// Build finalizes the block.
func (bl *Block) Build() (common.Hash, error) {
	if bl.built {
		return common.Hash{}, errors.New("block is already built")
	}

	bl.mtx.Lock()
	defer bl.mtx.Unlock()

	if !sort.StringsAreSorted(bl.uids) {
		bl.mtx.Lock()
		sort.Strings(bl.uids)
		bl.mtx.Unlock()
	}

	leaves := make(map[string]common.Hash)

	for _, uid := range bl.uids {
		leaves[uid] = bl.txs[uid].Hash()
	}

	tree, err := merkle.NewTree(leaves, depth257)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to build block")
	}

	bl.tree = tree
	bl.built = true
	return bl.tree.Root(), nil
}

// Marshal encodes block object to raw json data.
func (bl *Block) Marshal() ([]byte, error) {
	txs := make(map[string][]byte)

	for uid, tx := range bl.txs {
		var data []byte
		buf := bytes.NewBuffer(data)

		if err := tx.EncodeRLP(buf); err != nil {
			msg := fmt.Sprintf("failed to encode transaction %s",
				tx.Hash().String())
			return nil, errors.Wrap(err, msg)
		}
		txs[uid] = buf.Bytes()
	}
	raw, err := json.Marshal(txs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode"+
			" transactions")
	}

	return raw, nil
}

// Unmarshal decodes raw json data to block object.
func Unmarshal(raw []byte, block *Block) error {
	var txs map[string][]byte

	if err := json.Unmarshal(raw, &txs); err != nil {
		return errors.Wrap(err, "failed to decode"+
			" transactions")
	}

	for _, rawTX := range txs {
		tx := &transaction.Transaction{}

		buf := bytes.NewBuffer(rawTX)

		if err := transaction.DecodeRLP(buf, tx); err != nil {
			return errors.Wrap(err, "failed to decode"+
				" transaction")
		}
		block.AddTx(tx)
	}
	return nil
}
