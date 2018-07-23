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
