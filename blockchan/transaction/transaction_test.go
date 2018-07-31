package transaction

import (
	"bytes"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func testKey(t *testing.T) *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	return key
}

func testAccount(key *ecdsa.PrivateKey) *bind.TransactOpts {
	return bind.NewKeyedTransactor(key)
}

func TestTransaction(t *testing.T) {
	key1 := testKey(t)
	key2 := testKey(t)

	oldOwner := testAccount(key1)
	newOwner := testAccount(key2)

	unsignedTx, err := NewTransaction(
		big.NewInt(-1), big.NewInt(43), big.NewInt(1), newOwner.From)
	if err != nil {
		t.Fatal(err)
	}

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
	key1 := testKey(t)
	key2 := testKey(t)

	oldOwner := testAccount(key1)
	newOwner := testAccount(key2)

	unsignedTx, err := NewTransaction(
		big.NewInt(0), big.NewInt(43), big.NewInt(1), newOwner.From)
	if err != nil {
		t.Fatal(err)
	}

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
