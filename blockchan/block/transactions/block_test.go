package transactions

import (
	"bytes"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/merkle"
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
	key *ecdsa.PrivateKey, nonce *big.Int) *transaction.Transaction {
	randKey := account.GenKey()

	unsignedTx, err := transaction.NewTransaction(
		prevBlock, randKey.X, randKey.Y, nonce, newOwner)
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
			newOwner.account.From, owner.key, big.NewInt(0))

		txs = append(txs, tx)
	}
	return txs
}

func validateTx(t *testing.T, tx *transaction.Transaction,
	root, proof []byte) {
	if !merkle.CheckMembership(tx.UID(), tx.Hash(),
		common.BytesToHash(root), proof) {
		t.Fatal("the transaction was incorrectly" +
			" included in the block")
	}
}

func TestBlockAddTx(t *testing.T) {
	number := 4
	txs := generateTXs(t, number, 0)
	bl := NewTxBlock().(*TrBlock)

	for _, tx := range txs {
		if err := bl.AddTx(tx); err != nil {
			t.Fatal(err)
		}
	}

	if len(bl.txs) != number {
		t.Fatalf("tx number must be %d, got %d", number, len(bl.txs))
	}
}

func TestBlockBuild(t *testing.T) {
	txs := generateTXs(t, numberTx, testPrevBlock)
	bl := NewTxBlock().(*TrBlock)

	for _, tx := range txs {
		if err := bl.AddTx(tx); err != nil {
			t.Fatal(err)
		}
	}

	root, err := bl.Build()
	if err != nil {
		t.Fatal(err)
	}

	for _, tx := range txs {
		proof := merkle.CreateProof(tx.UID(), merkle.Depth257,
			bl.tree.GetStructure(), bl.tree.DefaultNodes)
		validateTx(t, tx, root.Bytes(), proof)
	}
}

func TestBlockEncodeDecode(t *testing.T) {
	txs := generateTXs(t, numberTx, testPrevBlock)
	bl := NewTxBlock().(*TrBlock)

	for _, tx := range txs {
		if err := bl.AddTx(tx); err != nil {
			t.Fatal(err)
		}
	}

	root1, err := bl.Build()
	if err != nil {
		t.Fatal(err)
	}

	proof := merkle.CreateProof(txs[0].UID(), merkle.Depth257,
		bl.tree.GetStructure(), bl.tree.DefaultNodes)

	raw, err := bl.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	reconstructed := NewTxBlock()

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

	validateTx(t, txs[0], root2.Bytes(), proof)
}

func TestBlockAddExistsTx(t *testing.T) {
	txs := generateTXs(t, numberTx, testPrevBlock)

	bl := NewTxBlock()

	for _, tx := range txs {
		if err := bl.AddTx(tx); err != nil {
			t.Fatal(err)
		}
	}

	if err := bl.AddTx(txs[0]); err == nil {
		t.Fatal("the transaction already exists in the block")
	}
}
