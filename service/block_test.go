package service

import (
	"bytes"
	"context"
	"testing"
)

func TestServiceAcceptTransaction(t *testing.T) {
	instance := newInstance(t)
	tx := testTx(t, zero, one, two, three, owner.From, owner)

	if err := instance.service.AcceptTransaction(tx); err != nil {
		t.Fatal(err)
	}
}

func TestServiceCreateTxProof(t *testing.T) {
	instance := newInstance(t)
	tx := testTx(t, zero, one, two, three, owner.From, owner)

	if err := instance.service.AcceptTransaction(tx); err != nil {
		t.Fatal(err)
	}

	_, err := instance.service.BuildTxBlock()
	if err != nil {
		t.Fatal(err)
	}

	err = instance.service.SaveCurrentTxBlock()
	if err != nil {
		t.Fatal(err)
	}

	blockNum, err := instance.service.LastTxBlockNumber()
	if err != nil {
		t.Fatal(err)
	}

	err = instance.service.SendCurrentTxBlock(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	proof, err := instance.service.CreateTxProof(one, blockNum.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if len(proof) == 0 {
		t.Fatal("empty proof")
	}
}

func TestInitCurrentTxBlock(t *testing.T) {
	instance := newInstance(t)
	tx := testTx(t, zero, one, two, three, owner.From, owner)

	if err := instance.service.AcceptTransaction(tx); err != nil {
		t.Fatal(err)
	}

	_, err := instance.service.BuildTxBlock()
	if err != nil {
		t.Fatal(err)
	}

	hash := instance.service.CurrentTxBlockHash()

	instance.service.InitNewTxBlock()

	hash2 := instance.service.CurrentTxBlockHash()

	if bytes.Equal(hash.Bytes(), hash2.Bytes()) {
		t.Fatal("the current block was not initialized")
	}
}

func TestRawTxBlock(t *testing.T) {
	instance := newInstance(t)
	tx := testTx(t, zero, one, two, three, owner.From, owner)

	if err := instance.service.AcceptTransaction(tx); err != nil {
		t.Fatal(err)
	}

	_, err := instance.service.BuildTxBlock()
	if err != nil {
		t.Fatal(err)
	}

	rawBlock, err := instance.service.RawTxBlock(one.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if len(rawBlock) != 0 {
		t.Fatal("block must not be in the database")
	}

	err = instance.service.SaveCurrentTxBlock()
	if err != nil {
		t.Fatal(err)
	}

	rawBlock, err = instance.service.RawTxBlock(zero.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if len(rawBlock) == 0 {
		t.Fatal("block must be in the database")
	}
}
