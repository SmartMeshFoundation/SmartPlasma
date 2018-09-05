package service

import (
	"bytes"
	"context"
	"testing"

	"github.com/SmartMeshFoundation/Spectrum/common"
)

func TestAcceptTransaction(t *testing.T) {
	i := newInstance(t)
	tx := testTx(t, zero, one, two, three, owner.From, owner)

	if err := i.service.AcceptTransaction(tx); err != nil {
		t.Fatal(err)
	}
}

func TestCreateProof(t *testing.T) {
	i := newInstance(t)
	tx := testTx(t, zero, one, two, three, owner.From, owner)

	if err := i.service.AcceptTransaction(tx); err != nil {
		t.Fatal(err)
	}

	_, err := i.service.BuildBlock()
	if err != nil {
		t.Fatal(err)
	}

	block1 := i.service.CurrentBlock()

	sendBlockTx, err := i.service.SendBlockHash(context.Background(), block1.Hash())
	if err != nil {
		t.Fatal(err)
	}

	err = i.service.mineTx(context.Background(), sendBlockTx)
	if err != nil {
		t.Fatal(err)
	}

	err = i.service.SaveBlockToDB(one.Uint64(), block1)
	if err != nil {
		t.Fatal(err)
	}

	blockNum, err := i.service.LastBlockNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	proof, err := i.service.CreateProof(one, blockNum.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if len(proof) == 0 {
		t.Fatal("empty proof")
	}

	exist, err := i.service.VerifyTxProof(
		one, tx.Hash(), blockNum.Uint64(), proof)
	if err != nil {
		t.Fatal(err)
	}

	if !exist {
		t.Fatal("the transaction in the block")
	}

	proof2, err := i.service.CreateProof(two, blockNum.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if len(proof2) == 0 {
		t.Fatal("empty proof")
	}

	exist, err = i.service.VerifyTxProof(
		two, tx.Hash(), blockNum.Uint64(), proof2)
	if err != nil {
		t.Fatal(err)
	}

	if exist {
		t.Fatal("the transaction not in the block")
	}
}

func TestInitBlock(t *testing.T) {
	i := newInstance(t)
	tx := testTx(t, zero, one, two, three, owner.From, owner)

	if err := i.service.AcceptTransaction(tx); err != nil {
		t.Fatal(err)
	}

	_, err := i.service.BuildBlock()
	if err != nil {
		t.Fatal(err)
	}

	hash := i.service.CurrentBlock().Hash()

	i.service.InitBlock()

	hash2 := i.service.CurrentBlock().Hash()

	if bytes.Equal(hash.Bytes(), hash2.Bytes()) || (hash2 != common.Hash{}) {
		t.Fatal("the current block was not initialized")
	}
}

func TestRawBlockFromDB(t *testing.T) {
	i := newInstance(t)
	tx := testTx(t, zero, one, two, three, owner.From, owner)

	if err := i.service.AcceptTransaction(tx); err != nil {
		t.Fatal(err)
	}

	_, err := i.service.BuildBlock()
	if err != nil {
		t.Fatal(err)
	}

	rawBlock, err := i.service.RawBlockFromDB(one.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if len(rawBlock) != 0 {
		t.Fatal("block must not be in the database")
	}

	block := i.service.CurrentBlock()

	err = i.service.SaveBlockToDB(one.Uint64(), block)
	if err != nil {
		t.Fatal(err)
	}

	rawBlock, err = i.service.RawBlockFromDB(one.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if len(rawBlock) == 0 {
		t.Fatal("block must be in the database")
	}
}
