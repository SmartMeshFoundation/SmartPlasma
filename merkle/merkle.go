package merkle

import (
	"bytes"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

var (
	one = big.NewInt(1)
	two = big.NewInt(2)
)

// Tree is structure to storage Merkle tree.
type Tree struct {
	root         common.Hash
	Tree         []map[string]common.Hash
	depth        *big.Int
	DefaultNodes map[string]common.Hash
}

// NewTree creates new Merkle tree.
func NewTree(leaves map[string]common.Hash, depth *big.Int) (*Tree, error) {
	length := new(big.Int).SetInt64(int64(len(leaves)))
	capacity := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2),
		depth, nil), big.NewInt(1))

	if length.Cmp(capacity) == 1 {
		return nil, errors.Errorf("Tree with depth %d could not have %d"+
			" leaves", depth, len(leaves))
	}

	tree := &Tree{depth: depth}

	defaultNodes := createDefaultNodes(depth)
	tree.DefaultNodes = defaultNodes

	if leaves != nil {
		tree.Tree = create(leaves, depth, defaultNodes)
		tree.root = tree.Tree[len(tree.Tree)-1]["0"]
	} else {
		tree.Tree = []map[string]common.Hash{}
		tree.root = defaultNodes[new(big.Int).Sub(depth, big.NewInt(1)).String()]
	}
	return tree, nil
}

// Root returns Merkle root.
func (tr *Tree) Root() common.Hash {
	return tr.root
}

func create(leaves map[string]common.Hash, depth *big.Int,
	defaultNodes map[string]common.Hash) []map[string]common.Hash {
	tree := []map[string]common.Hash{leaves}
	treeLevel := leaves

	for level := big.NewInt(0); level.Cmp(new(big.Int).Sub(depth, one)) < 0; level.Add(level, one) {
		nextLevel := make(map[string]common.Hash)
		prevIndex := big.NewInt(-1)

		keys := sortKeys(treeLevel)

		for _, strUID := range keys {
			index, _ := new(big.Int).SetString(strUID, 10)
			value := treeLevel[strUID]

			if new(big.Int).Rem(index, two).Cmp(big.NewInt(0)) == 0 {
				div := new(big.Int).Div(index, two)
				nextLevel[div.String()] = crypto.Keccak256Hash(value.Bytes(),
					defaultNodes[level.String()].Bytes())
			} else {
				div := new(big.Int).Div(index, two)

				if index.Cmp(new(big.Int).Add(prevIndex, one)) == 0 {
					nextLevel[div.String()] = crypto.Keccak256Hash(
						treeLevel[prevIndex.String()].Bytes(), value.Bytes())
				} else {
					nextLevel[div.String()] = crypto.Keccak256Hash(
						defaultNodes[level.String()].Bytes(), value.Bytes())
				}
			}
			prevIndex = index
		}
		treeLevel = nextLevel
		tree = append(tree, treeLevel)
	}

	return tree
}

func createDefaultNodes(depth *big.Int) map[string]common.Hash {
	defaultNodes := map[string]common.Hash{"0": {}}

	for level := new(big.Int).Set(one); level.Cmp(depth) < 0; level.Add(level, one) {
		nextLevel := new(big.Int).Sub(level, one)
		prevDefault := defaultNodes[nextLevel.String()]
		defaultNodes[level.String()] = crypto.Keccak256Hash(
			prevDefault.Bytes(), prevDefault.Bytes())
	}
	return defaultNodes
}

// CreateProof creates merkle proof for particular uid.
func CreateProof(uid, depth *big.Int, tree []map[string]common.Hash,
	defaultNodes map[string]common.Hash) []byte {
	index := new(big.Int).Set(uid)
	var proof []byte

	limit := new(big.Int).Sub(depth, one)

	for level := big.NewInt(0); level.Cmp(limit) < 0; level.Add(level, one) {
		var siblingIndex *big.Int

		if new(big.Int).Rem(index, two).Cmp(big.NewInt(0)) == 0 {
			siblingIndex = new(big.Int).Add(index, one)
		} else {
			siblingIndex = new(big.Int).Sub(index, one)
		}
		index = index.Div(index, two)

		l := tree[level.Uint64()]
		if _, ok := l[siblingIndex.String()]; ok {
			proof = append(proof, l[siblingIndex.String()].Bytes()...)
		} else {
			proof = append(proof, defaultNodes[level.String()].Bytes()...)
		}
	}
	return proof
}

// CheckMembership checks membership.
func CheckMembership(uid *big.Int, leaf, rootHash []byte,
	proof []byte) bool {
	if len(proof) == 0 || len(proof)%32 != 0 {
		return false
	}

	index := new(big.Int).Set(uid)

	computedHash := leaf

	for i := 0; i < len(proof); i += 32 {
		proofElement := proof[i : i+32]

		if new(big.Int).Rem(index, big.NewInt(2)).Uint64() == 0 {
			computedHash = crypto.Keccak256(computedHash, proofElement)
		} else {
			computedHash = crypto.Keccak256(proofElement, computedHash)
		}
		index = index.Div(index, big.NewInt(2))
	}
	return bytes.Equal(computedHash, rootHash)
}

func sortKeys(dict map[string]common.Hash) []string {
	var keys []string

	for k := range dict {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	if !sort.StringsAreSorted(keys) {
		panic("slice are not sorted")
	}
	return keys
}
