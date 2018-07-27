package merkle

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	dummyVal = common.BytesToHash(genHash("\x01"))
	emptyVal = common.Hash{}

	uid0 = big.NewInt(0)
	uid1 = big.NewInt(1)
	uid2 = big.NewInt(2)
	uid3 = big.NewInt(3)

	uidMax = new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2),
		big.NewInt(256), nil), big.NewInt(1))

	depth3   = big.NewInt(3)
	depth257 = big.NewInt(257)
)

func testTree(t *testing.T, leaves map[string]common.Hash,
	depth *big.Int) *Tree {
	tree, err := NewTree(leaves, depth)
	if err != nil {
		t.Fatal(err)
	}
	return tree
}

func genHash(pattern string) []byte {
	var hash []byte

	for i := 0; i < 32; i++ {
		hash = append(hash, []byte(pattern)...)
	}
	return hash
}

func TestNewTreeAllLeaves(t *testing.T) {
	leaves := map[string]common.Hash{
		uid0.String(): dummyVal,
		uid1.String(): dummyVal,
		uid2.String(): dummyVal,
		uid3.String(): dummyVal,
	}

	tree := testTree(t, leaves, depth3)

	midLevelVal := crypto.Keccak256Hash(dummyVal.Bytes(), dummyVal.Bytes())

	if !bytes.Equal(tree.root.Bytes(), crypto.Keccak256Hash(
		midLevelVal.Bytes(), midLevelVal.Bytes()).Bytes()) {
		t.Fatal("hashes not equal")
	}
}

func TestNewTreeEmptyLeaves(t *testing.T) {
	tree := testTree(t, nil, depth3)

	midLevelVal := crypto.Keccak256Hash(emptyVal.Bytes(), emptyVal.Bytes())

	if !bytes.Equal(tree.root.Bytes(), crypto.Keccak256Hash(
		midLevelVal.Bytes(), midLevelVal.Bytes()).Bytes()) {
		t.Fatal("hashes not equal")
	}
}

func TestNewTreeEmptyLeftLeave(t *testing.T) {
	leaves := map[string]common.Hash{
		uid1.String(): dummyVal,
		uid2.String(): dummyVal,
		uid3.String(): dummyVal,
	}

	tree := testTree(t, leaves, depth3)

	midLeftVal := crypto.Keccak256Hash(emptyVal.Bytes(), dummyVal.Bytes())
	midRightVal := crypto.Keccak256Hash(dummyVal.Bytes(), dummyVal.Bytes())

	result := crypto.Keccak256Hash(midLeftVal.Bytes(), midRightVal.Bytes())

	if !bytes.Equal(tree.root.Bytes(), result.Bytes()) {
		t.Fatal("hashes not equal")
	}
}

func TestNewTreeEmptyRightLeave(t *testing.T) {
	leaves := map[string]common.Hash{
		uid0.String(): dummyVal,
		uid2.String(): dummyVal,
		uid3.String(): dummyVal,
	}

	tree := testTree(t, leaves, depth3)

	midLeftVal := crypto.Keccak256Hash(dummyVal.Bytes(), emptyVal.Bytes())
	midRightVal := crypto.Keccak256Hash(dummyVal.Bytes(), dummyVal.Bytes())

	if !bytes.Equal(tree.root.Bytes(), crypto.Keccak256Hash(midLeftVal.Bytes(),
		midRightVal.Bytes()).Bytes()) {
		t.Fatal("hashes not equal")
	}
}

func TestExceedTreeSize(t *testing.T) {
	leaves := map[string]common.Hash{
		uid0.String(): dummyVal,
		uid1.String(): dummyVal,
	}
	_, err := NewTree(leaves, big.NewInt(1))
	if err == nil {
		t.Fatal("expect not null error")
	}
}

func TestCreateProof(t *testing.T) {
	leaves := map[string]common.Hash{
		uid0.String(): dummyVal,
		uid2.String(): dummyVal,
		uid3.String(): dummyVal,
	}

	tree := testTree(t, leaves, depth3)

	midLeftVal := crypto.Keccak256Hash(dummyVal.Bytes(), emptyVal.Bytes())
	midRightVal := crypto.Keccak256Hash(dummyVal.Bytes(), dummyVal.Bytes())

	proof1 := CreateProof(uid0, depth3, tree.Tree, tree.DefaultNodes)
	proof2 := CreateProof(uid1, depth3, tree.Tree, tree.DefaultNodes)
	proof3 := CreateProof(uid2, depth3, tree.Tree, tree.DefaultNodes)
	proof4 := CreateProof(uid3, depth3, tree.Tree, tree.DefaultNodes)

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

func TestLimit(t *testing.T) {
	leaves := map[string]common.Hash{
		uid0.String():   dummyVal,
		uidMax.String(): dummyVal,
	}

	tree := testTree(t, leaves, depth257)

	proof1 := CreateProof(uidMax, depth257, tree.Tree, tree.DefaultNodes)

	if !CheckMembership(uidMax, dummyVal.Bytes(), tree.root.Bytes(), proof1) {
		t.Fatal("membership is not confirmed")
	}

}
