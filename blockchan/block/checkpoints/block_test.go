package checkpoints

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartmeshfoundation/smartplasma/blockchan/account"
	"github.com/smartmeshfoundation/smartplasma/merkle"
)

const (
	numberCheckpoints = 4
)

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
	if !merkle.CheckMembership(uid, common.BigToHash(number),
		common.BytesToHash(root), proof) {
		t.Fatal("the checkpoint was incorrectly" +
			" included in the block")
	}
}

func TestAddCheckpoint(t *testing.T) {
	chs := generateCheckpoints(numberCheckpoints)
	bl := NewBlock()

	for _, ch := range chs {
		if err := bl.AddCheckpoint(ch.uid, ch.nonce); err != nil {
			t.Fatal(err)
		}
	}

	if len(bl.numbers) != numberCheckpoints {
		t.Fatalf("tx number must be %d, got %d",
			numberCheckpoints, len(bl.numbers))
	}
}

func TestBlockBuild(t *testing.T) {
	chs := generateCheckpoints(numberCheckpoints)
	bl := NewBlock()

	for _, ch := range chs {
		if err := bl.AddCheckpoint(ch.uid, ch.nonce); err != nil {
			t.Fatal(err)
		}
	}

	root, err := bl.Build()
	if err != nil {
		t.Fatal(err)
	}

	for _, ch := range chs {
		proof := merkle.CreateProof(ch.uid, merkle.Depth257,
			bl.tree.GetStructure(), bl.tree.DefaultNodes)
		validateCheckpoint(t, ch.uid, ch.nonce, root.Bytes(), proof)
	}
}

func TestBlockEncodeDecode(t *testing.T) {
	chs := generateCheckpoints(numberCheckpoints)
	bl := NewBlock()

	for _, ch := range chs {
		if err := bl.AddCheckpoint(ch.uid, ch.nonce); err != nil {
			t.Fatal(err)
		}
	}

	root1, err := bl.Build()
	if err != nil {
		t.Fatal(err)
	}

	proof := merkle.CreateProof(chs[0].uid, merkle.Depth257,
		bl.tree.GetStructure(), bl.tree.DefaultNodes)

	raw, err := bl.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	reconstructed := NewBlock()

	if err := reconstructed.Unmarshal(raw); err != nil {
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
	bl := NewBlock()

	for _, ch := range chs {
		if err := bl.AddCheckpoint(ch.uid, ch.nonce); err != nil {
			t.Fatal(err)
		}
	}

	if err := bl.AddCheckpoint(chs[0].uid, chs[0].nonce); err == nil {
		t.Fatal("the transaction already exists in the block")
	}
}
