package merkle

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	dummyVal = common.BytesToHash(genHash("\x01"))
	emptyVal = common.Hash{}
)

func genHash(pattern string) []byte {
	var hash []byte

	for i := 0; i < 32; i++ {
		hash = append(hash, []byte(pattern)...)
	}
	return hash
}

func TestNewTreeAllLeaves(t *testing.T) {
	leaves := map[int]common.Hash{0: dummyVal, 1: dummyVal, 2: dummyVal,
		3: dummyVal}

	tree, err := NewTree(leaves, 3)
	if err != nil {
		t.Fatal(err)
	}

	midLevelVal := crypto.Keccak256Hash(dummyVal.Bytes(), dummyVal.Bytes())

	result := crypto.Keccak256Hash(midLevelVal.Bytes(), midLevelVal.Bytes())

	if !bytes.Equal(tree.root.Bytes(), result.Bytes()) {
		t.Fatal("hashes not equal")
	}
}

func TestNewTreeEmptyLeaves(t *testing.T) {
	tree, err := NewTree(nil, 3)
	if err != nil {
		t.Fatal(err)
	}

	midLevelVal := crypto.Keccak256Hash(emptyVal.Bytes(), emptyVal.Bytes())

	result := crypto.Keccak256Hash(midLevelVal.Bytes(), midLevelVal.Bytes())

	if !bytes.Equal(tree.root.Bytes(), result.Bytes()) {
		t.Fatal("hashes not equal")
	}
}

func TestNewTreeEmptyLeftLeave(t *testing.T) {
	leaves := map[int]common.Hash{1: dummyVal, 2: dummyVal,
		3: dummyVal}

	tree, err := NewTree(leaves, 3)
	if err != nil {
		t.Fatal(err)
	}

	midLeftVal := crypto.Keccak256Hash(emptyVal.Bytes(), dummyVal.Bytes())
	midRightVal := crypto.Keccak256Hash(dummyVal.Bytes(), dummyVal.Bytes())

	result := crypto.Keccak256Hash(midLeftVal.Bytes(), midRightVal.Bytes())

	if !bytes.Equal(tree.root.Bytes(), result.Bytes()) {
		t.Fatal("hashes not equal")
	}
}

func TestNewTreeEmptyRightLeave(t *testing.T) {
	leaves := map[int]common.Hash{0: dummyVal, 2: dummyVal,
		3: dummyVal}

	tree, err := NewTree(leaves, 3)
	if err != nil {
		t.Fatal(err)
	}

	midLeftVal := crypto.Keccak256Hash(dummyVal.Bytes(), emptyVal.Bytes())
	midRightVal := crypto.Keccak256Hash(dummyVal.Bytes(), dummyVal.Bytes())

	result := crypto.Keccak256Hash(midLeftVal.Bytes(), midRightVal.Bytes())

	if !bytes.Equal(tree.root.Bytes(), result.Bytes()) {
		t.Fatal("hashes not equal")
	}
}

func TestExceedTreeSize(t *testing.T) {
	leaves := map[int]common.Hash{0: {}, 1: {}}
	_, err := NewTree(leaves, 1)
	if err == nil {
		t.Fatal("expect not null error")
	}
}

func TestCreateProof(t *testing.T) {
	leaves := map[int]common.Hash{0: dummyVal, 2: dummyVal,
		3: dummyVal}
	tree, err := NewTree(leaves, 3)
	if err != nil {
		t.Fatal(err)
	}
	midLeftVal := crypto.Keccak256Hash(dummyVal.Bytes(), emptyVal.Bytes())
	midRightVal := crypto.Keccak256Hash(dummyVal.Bytes(), dummyVal.Bytes())
	proof1 := CreateProof(0, 3, tree.tree, tree.defaultNodes)
	proof2 := CreateProof(1, 3, tree.tree, tree.defaultNodes)
	proof3 := CreateProof(2, 3, tree.tree, tree.defaultNodes)
	proof4 := CreateProof(3, 3, tree.tree, tree.defaultNodes)

	if !bytes.Equal(proof1, append(emptyVal.Bytes(), midRightVal.Bytes()...)) {
		t.Fatal("hashes not equal")
	}
	if !bytes.Equal(proof2, append(dummyVal.Bytes(), midRightVal.Bytes()...)) {
		t.Fatal("hashes not equal")
	}

	if !bytes.Equal(proof3, append(dummyVal.Bytes(), midLeftVal.Bytes()...)) {
		t.Fatal("hashes not equal")
	}

	if !bytes.Equal(proof4, append(dummyVal.Bytes(), midLeftVal.Bytes()...)) {
		t.Fatal("hashes not equal")
	}
}
