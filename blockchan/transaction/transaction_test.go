package transaction

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/smartmeshfoundation/smartplasma/blockchan/account"
)

func TestTransaction(t *testing.T) {
	key1 := account.GenKey()
	key2 := account.GenKey()

	oldOwner := account.Account(key1)
	newOwner := account.Account(key2)

	unsignedTx := NewTransaction(
		big.NewInt(0), big.NewInt(43), big.NewInt(1), newOwner.From)

	tx, err := unsignedTx.SignTx(key1)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(tx.Hash().Bytes(), unsignedTx.Hash().Bytes()) {
		t.Fatal("hashes not equal")
	}

	addr, err := Sender(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(oldOwner.From.Bytes(), addr.Bytes()) {
		t.Fatal("addresses not equal")
	}
}

func TestEncodeDecodeRLP(t *testing.T) {
	key1 := account.GenKey()
	key2 := account.GenKey()

	oldOwner := account.Account(key1)
	newOwner := account.Account(key2)

	unsignedTx := NewTransaction(
		big.NewInt(0), big.NewInt(43), big.NewInt(1), newOwner.From)

	tx, err := unsignedTx.SignTx(key1)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(tx.Hash().Bytes(), unsignedTx.Hash().Bytes()) {
		t.Fatal("hashes not equal")
	}

	addr, err := Sender(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(oldOwner.From.Bytes(), addr.Bytes()) {
		t.Fatal("addresses not equal")
	}

	var b []byte

	buf := bytes.NewBuffer(b)

	if err := tx.EncodeRLP(buf); err != nil {
		t.Fatal(err)
	}

	tx2 := &Transaction{}

	if err := DecodeRLP(buf, tx2); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(tx.Hash().Bytes(), tx2.Hash().Bytes()) {
		t.Fatal("hashes not equal")
	}

	addr2, err := Sender(tx2)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(oldOwner.From.Bytes(), addr2.Bytes()) {
		t.Fatal("addresses not equal")
	}
}
