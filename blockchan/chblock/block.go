package chblock

import (
	"encoding/json"
	"math/big"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartmeshfoundation/smartplasma/merkle"
)

// TODO: union with tx block, create interface - block
// TODO: need refactor

// Default value to depth.
const (
	DefaultDepth = 257
)

var (
	depth257 = big.NewInt(DefaultDepth)
)

// Block object.
type Block struct {
	mtx     sync.Mutex
	uids    []string
	numbers map[string]common.Hash
	tree    *merkle.Tree

	built bool
}

func NewBlock() *Block {
	return &Block{
		mtx:     sync.Mutex{},
		numbers: make(map[string]common.Hash),
	}
}

func (bl *Block) Hash() common.Hash {
	if !bl.built {
		return common.Hash{}
	}
	return bl.tree.Root()
}

func (bl *Block) AddCheckpoint(uid, number *big.Int) error {
	bl.mtx.Lock()
	defer bl.mtx.Unlock()

	if _, ok := bl.numbers[uid.String()]; ok {
		return errors.Errorf("checkpoint for uid %s already"+
			" exist in the block", uid.String())
	}

	bl.uids = append(bl.uids, uid.String())
	sort.Strings(bl.uids)
	bl.numbers[uid.String()] = common.BigToHash(number)
	return nil
}

func (bl *Block) NumberOfCheckpoints() int64 {
	return int64(len(bl.numbers))
}

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

	tree, err := merkle.NewTree(bl.numbers, depth257)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to build block")
	}

	bl.tree = tree
	bl.built = true
	return bl.tree.Root(), nil
}

func (bl *Block) Marshal() ([]byte, error) {
	raw, err := json.Marshal(bl.numbers)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode"+
			" checkpoints")
	}

	return raw, nil
}

func Unmarshal(raw []byte, block *Block) error {
	var checkpoints map[string]common.Hash

	if err := json.Unmarshal(raw, &checkpoints); err != nil {
		return errors.Wrap(err, "failed to decode"+
			" checkpoints")
	}

	for uidStr, checkpoint := range checkpoints {
		id, ok := new(big.Int).SetString(uidStr, 10)
		if !ok {
			continue
		}

		if err := block.AddCheckpoint(id, checkpoint.Big()); err != nil {
			return errors.Wrap(err, "failed to add checkpoint in the block")
		}
	}
	return nil
}

func (bl *Block) CreateProof(uid *big.Int) []byte {
	if !bl.built {
		return nil
	}
	return merkle.CreateProof(uid, depth257, bl.tree.GetStructure(),
		bl.tree.DefaultNodes)
}

// CheckMembership checks membership.
func CheckMembership(uid *big.Int, txHash, blockHash common.Hash,
	proof []byte) bool {
	return merkle.CheckMembership(uid, txHash.Bytes(),
		blockHash.Bytes(), proof)
}
