package merkle

import (
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// TODO: this is a naive implementation. Need refactor with bigInt.

// Tree is object to storage merkle tree.
type Tree struct {
	root         common.Hash
	tree         []map[int]common.Hash
	depth        int
	defaultNodes map[int]common.Hash
}

// NewTree creates new merkle tree.
func NewTree(leaves map[int]common.Hash, depth int) (*Tree, error) {
	length := new(big.Int).SetInt64(int64(len(leaves)))
	capacity := new(big.Int).Exp(big.NewInt(2),
		new(big.Int).Sub(big.NewInt(int64(depth)), big.NewInt(1)), nil)

	if length.Cmp(capacity) == 1 {
		return nil, errors.Errorf("tree with depth %d could not have %d"+
			" leaves", depth, len(leaves))
	}

	tree := &Tree{depth: depth}

	defaultNodes := createDefaultNodes(depth)
	tree.defaultNodes = defaultNodes

	if leaves != nil {
		tree.tree = create(leaves, depth, defaultNodes)
		tree.root = tree.tree[len(tree.tree)-1][0]
	} else {
		tree.tree = []map[int]common.Hash{}
		tree.root = defaultNodes[depth-1]
	}
	return tree, nil
}

func create(leaves map[int]common.Hash, depth int,
	defaultNodes map[int]common.Hash) []map[int]common.Hash {
	tree := []map[int]common.Hash{leaves}
	treeLevel := leaves

	for level := 0; level < depth-1; level++ {
		nextLevel := make(map[int]common.Hash)
		prevIndex := -1

		keys := sortKeys(treeLevel)

		for _, uid := range keys {
			index := uid
			value := treeLevel[uid]

			if index%2 == 0 {
				nextLevel[index/2] = crypto.Keccak256Hash(value.Bytes(),
					defaultNodes[level].Bytes())

			} else {
				if index == prevIndex+1 {
					nextLevel[index/2] = crypto.Keccak256Hash(
						treeLevel[prevIndex].Bytes(), value.Bytes())
				} else {
					nextLevel[index/2] = crypto.Keccak256Hash(
						defaultNodes[level].Bytes(), value.Bytes())
				}
			}
			prevIndex = index
		}
		treeLevel = nextLevel
		tree = append(tree, treeLevel)
	}
	return tree
}

func createDefaultNodes(depth int) map[int]common.Hash {
	defaultNodes := map[int]common.Hash{0: {}}

	for level := 1; level < depth; level++ {
		prevDefault := defaultNodes[level-1]
		defaultNodes[level] = crypto.Keccak256Hash(prevDefault.Bytes(),
			prevDefault.Bytes())
	}
	return defaultNodes
}

// CreateProof creates merkle proof for particular uid.
func CreateProof(uid, depth int, tree []map[int]common.Hash,
	defaultNodes map[int]common.Hash) []byte {
	index := uid
	var proof []byte

	for level := 0; level < depth-1; level++ {
		var siblingIndex int

		if index%2 == 0 {
			siblingIndex = index + 1
		} else {
			siblingIndex = index - 1
		}
		index = index / 2

		l := tree[level]
		if _, ok := l[siblingIndex]; ok {
			proof = append(proof, l[siblingIndex].Bytes()...)
		} else {
			proof = append(proof, defaultNodes[level].Bytes()...)
		}
	}
	return proof
}

func sortKeys(dict map[int]common.Hash) []int {
	var keys []int

	for k := range dict {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	if !sort.IntsAreSorted(keys) {
		panic("slice are not sorted")
	}
	return keys
}
