package chblock

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"bytes"
	"github.com/smartmeshfoundation/smartplasma/blockchan/account"
	"github.com/smartmeshfoundation/smartplasma/merkle"
)

const (
	numberCheckpoints = 4
	testPrevBlock     = 0
)

type acc struct {
	account *account.PlasmaTransactOpts
	key     *ecdsa.PrivateKey
}

type checkpoint struct {
	uid   *big.Int
	nonce *big.Int
}

func generate() *checkpoint {
	key := account.GenKey()

	return &checkpoint{
		uid:   key.D,
		nonce: key.X,
	}
}

func generateCheckpoints(number int) (result []*checkpoint) {
	for i := 0; i < number; i++ {
		result = append(result, generate())
	}
	return result
}

func validateCheckpoint(t *testing.T, uid, number *big.Int,
	root, proof []byte) {
	if !merkle.CheckMembership(uid, number.Bytes(), root, proof) {
		t.Fatal("the checkpoint was incorrectly" +
			" included in the block")
	}
}

func TestAddCheckpoint(t *testing.T) {
	number := 4
	chs := generateCheckpoints(number)
	block := NewBlock()

	for _, ch := range chs {
		if err := block.AddCheckpoint(ch.uid, ch.nonce); err != nil {
			t.Fatal(err)
		}
	}

	if len(block.numbers) != number {
		t.Fatalf("tx number must be %d, got %d", number, len(block.numbers))
	}
}

func TestBlockBuild(t *testing.T) {
	chs := generateCheckpoints(numberCheckpoints)
	block := NewBlock()

	for _, ch := range chs {
		if err := block.AddCheckpoint(ch.uid, ch.nonce); err != nil {
			t.Fatal(err)
		}
	}

	root, err := block.Build()
	if err != nil {
		t.Fatal(err)
	}

	for _, ch := range chs {
		proof := merkle.CreateProof(ch.uid, depth257,
			block.tree.GetStructure(), block.tree.DefaultNodes)
		validateCheckpoint(t, ch.uid, ch.nonce, root.Bytes(), proof)
	}
}

func TestBlockEncodeDecode(t *testing.T) {
	chs := generateCheckpoints(numberCheckpoints)
	block := NewBlock()

	for _, ch := range chs {
		if err := block.AddCheckpoint(ch.uid, ch.nonce); err != nil {
			t.Fatal(err)
		}
	}

	root1, err := block.Build()
	if err != nil {
		t.Fatal(err)
	}

	proof := merkle.CreateProof(chs[0].uid, depth257,
		block.tree.GetStructure(), block.tree.DefaultNodes)

	raw, err := block.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	reconstructed := NewBlock()

	if err := Unmarshal(raw, reconstructed); err != nil {
		t.Fatal(err)
	}

	root2, err := reconstructed.Build()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(root1.Bytes(), root2.Bytes()) {
		t.Fatal("blocks not equal")
	}
	validateCheckpoint(t, chs[0].uid, chs[0].nonce, root2.Bytes(), proof)
}

func TestBlockAddExistsTx(t *testing.T) {
	chs := generateCheckpoints(numberCheckpoints)
	block := NewBlock()

	for _, ch := range chs {
		if err := block.AddCheckpoint(ch.uid, ch.nonce); err != nil {
			t.Fatal(err)
		}
	}

	if err := block.AddCheckpoint(chs[0].uid, chs[0].nonce); err == nil {
		t.Fatal("the transaction already exists in the block")
	}
}
