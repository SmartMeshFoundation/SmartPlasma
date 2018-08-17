package transactions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/merkle"
)

// TxBlock defines the methods for standard Transactions block.
type TxBlock interface {
	block.Block
	AddTx(tx *transaction.Transaction) error
}

type txBlock struct {
	mtx  sync.Mutex
	uIDs []string
	txs  map[string]*transaction.Transaction
	tree *merkle.Tree

	built bool
}

// NewTxBlock creates new Transactions block in memory.
func NewTxBlock() TxBlock {
	return &txBlock{
		mtx: sync.Mutex{},
		txs: make(map[string]*transaction.Transaction),
	}
}

// Hash returns block hash.
func (bl *txBlock) Hash() common.Hash {
	if !bl.built {
		return common.Hash{}
	}
	return bl.tree.Root()
}

// AddTx adds a transaction to the block.
func (bl *txBlock) AddTx(tx *transaction.Transaction) error {
	bl.mtx.Lock()
	defer bl.mtx.Unlock()

	if _, ok := bl.txs[tx.UID().String()]; ok {
		return errors.Errorf("transaction for uid %s already"+
			" exist in the block", tx.UID().String())
	}

	bl.uIDs = append(bl.uIDs, tx.UID().String())
	sort.Strings(bl.uIDs)
	bl.txs[tx.UID().String()] = tx
	return nil
}

// NumberOfTX returns number of transactions in the block.
func (bl *txBlock) NumberOfTX() int64 {
	return int64(len(bl.txs))
}

// Build finalizes the block.
func (bl *txBlock) Build() (common.Hash, error) {
	if bl.built {
		return common.Hash{}, errors.New("block is already built")
	}

	bl.mtx.Lock()
	defer bl.mtx.Unlock()

	if !sort.StringsAreSorted(bl.uIDs) {
		bl.mtx.Lock()
		sort.Strings(bl.uIDs)
		bl.mtx.Unlock()
	}

	leaves := make(map[string]common.Hash)

	for _, uid := range bl.uIDs {
		leaves[uid] = bl.txs[uid].Hash()
	}

	tree, err := merkle.NewTree(leaves, merkle.Depth257)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to build block")
	}

	bl.tree = tree
	bl.built = true
	return bl.tree.Root(), nil
}

// CreateProof creates merkle proof for particular uid.
func (bl *txBlock) CreateProof(uid *big.Int) []byte {
	if !bl.built {
		return nil
	}
	return merkle.CreateProof(uid, merkle.Depth257, bl.tree.GetStructure(),
		bl.tree.DefaultNodes)
}

// Marshal encodes block object to raw json data.
func (bl *txBlock) Marshal() ([]byte, error) {
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
func (bl *txBlock) Unmarshal(raw []byte) error {
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

		if err := bl.AddTx(tx); err != nil {
			return errors.Wrap(err, "failed to add transaction in the block")
		}
	}
	return nil
}
