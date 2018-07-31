package block

import (
	"bytes"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartmeshfoundation/smartplasma/blockchan/account"
	"github.com/smartmeshfoundation/smartplasma/blockchan/transaction"
	"github.com/smartmeshfoundation/smartplasma/merkle"
)

const (
	numberTx      = 4
	testPrevBlock = 0
)

type acc struct {
	account *account.PlasmaTransactOpts
	key     *ecdsa.PrivateKey
}

func testAcc() *acc {
	key := account.GenKey()

	return &acc{
		key:     key,
		account: account.Account(key),
	}
}

func testTX(t *testing.T, prevBlock *big.Int, newOwner common.Address,
	key *ecdsa.PrivateKey) *transaction.Transaction {
	randKey := account.GenKey()

	unsignedTx, err := transaction.NewTransaction(
		prevBlock, randKey.X, randKey.Y, newOwner)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := unsignedTx.SignTx(key)
	if err != nil {
		t.Fatal(err)
	}
	return tx
}

func generateTXs(t *testing.T, number int,
	prevBlock int) (txs []*transaction.Transaction) {
	for i := 0; i < number; i++ {
		owner := testAcc()
		newOwner := testAcc()

		tx := testTX(t, big.NewInt(int64(prevBlock)),
			newOwner.account.From, owner.key)

		txs = append(txs, tx)
	}
	return txs
}

func validateTx(t *testing.T, tx *transaction.Transaction,
	root, proof []byte) {
	if !merkle.CheckMembership(tx.UID(), tx.Hash().Bytes(), root, proof) {
		t.Fatal("the transaction was incorrectly" +
			" included in the block")
	}
}

func TestBlockAddTx(t *testing.T) {
	number := 4
	txs := generateTXs(t, number, 0)
	block := NewBlock()

	for _, tx := range txs {
		if err := block.AddTx(tx); err != nil {
			t.Fatal(err)
		}
	}

	if len(block.txs) != number {
		t.Fatalf("tx number must be %d, got %d", number, len(block.txs))
	}
}

func TestBlockBuild(t *testing.T) {
	txs := generateTXs(t, numberTx, testPrevBlock)
	block := NewBlock()

	for _, tx := range txs {
		if err := block.AddTx(tx); err != nil {
			t.Fatal(err)
		}
	}

	root, err := block.Build()
	if err != nil {
		t.Fatal(err)
	}

	for _, tx := range txs {
		proof := merkle.CreateProof(tx.UID(), depth257, block.tree.Tree,
			block.tree.DefaultNodes)
		validateTx(t, tx, root.Bytes(), proof)
	}
}

func TestBlockEncodeDecode(t *testing.T) {
	txs := generateTXs(t, numberTx, testPrevBlock)
	block := NewBlock()

	for _, tx := range txs {
		if err := block.AddTx(tx); err != nil {
			t.Fatal(err)
		}
	}

	root1, err := block.Build()
	if err != nil {
		t.Fatal(err)
	}

	proof := merkle.CreateProof(txs[0].UID(), depth257, block.tree.Tree,
		block.tree.DefaultNodes)

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

	validateTx(t, txs[0], root2.Bytes(), proof)
}

func TestBlockAddExistsTx(t *testing.T) {
	txs := generateTXs(t, numberTx, testPrevBlock)

	block := NewBlock()

	for _, tx := range txs {
		if err := block.AddTx(tx); err != nil {
			t.Fatal(err)
		}
	}

	if err := block.AddTx(txs[0]); err == nil {
		t.Fatal("the transaction already exists in the block")
	}
}
