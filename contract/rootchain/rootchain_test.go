package rootchain

import (
	"bytes"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/erc20token"
)

var (
	server backend.Backend

	owner *account.PlasmaTransactOpts
	user1 *account.PlasmaTransactOpts
	user2 *account.PlasmaTransactOpts

	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
	four  = big.NewInt(4)
	five  = big.NewInt(5)
	zero  = big.NewInt(0)
)

type instance struct {
	tokenOwnerSession *erc20token.ExampleTokenSession
	tokenUser1Session *erc20token.ExampleTokenSession

	rootOwnerSession *RootChainSession
	rootUser1Session *RootChainSession
	rootUser2Session *RootChainSession

	tokenAddr     common.Address
	rootChainAddr common.Address
}

type PlasmaTestTx struct {
	tx    *transaction.Transaction
	block block.Block
	proof []byte
	rawTx []byte
}

func newInstance(t *testing.T) *instance {
	i := &instance{}

	rootChainAddr, _, err := Deploy(owner.TransactOpts, server)
	if err != nil {
		t.Fatal(err)
	}

	i.rootChainAddr = rootChainAddr

	i.rootOwnerSession = rootChainSession(t, owner.TransactOpts,
		i.rootChainAddr)
	i.rootUser1Session = rootChainSession(t, user1.TransactOpts,
		i.rootChainAddr)
	i.rootUser2Session = rootChainSession(t, user2.TransactOpts,
		i.rootChainAddr)
	return i
}

func rootChainSession(t *testing.T, account *bind.TransactOpts,
	contact common.Address) (session *RootChainSession) {
	session, err := NewRootChainSession(*account,
		contact, server)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func testDeposit(t *testing.T, i *instance) (uid *big.Int) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	deposit(t, i, common.BigToAddress(key.X), common.BigToAddress(key.Y), one)

	// receive logs with deposit. Get uid
	logs, err := LogsDeposit(i.rootOwnerSession.Contract)
	if err != nil {
		t.Fatalf("failed to parse deposit logs %s", err)
	}

	// TODO: single deposit. Not applicable for multiple deposits.
	if len(logs) != 1 {
		t.Fatal("wrong number of logs")
	}

	return logs[0].Uid
}

func deposit(t *testing.T, i *instance,
	account, token common.Address, amount *big.Int) {
	tx, err := i.rootOwnerSession.Deposit(account, token, amount)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(tx) {
		t.Fatal("failed to deposit")
	}
}

func testTx(t *testing.T, prevBlock, uid,
	amount *big.Int, nonce *big.Int, newOwner common.Address,
	signer *account.PlasmaTransactOpts) *transaction.Transaction {
	unsignedTx, err := transaction.NewTransaction(
		prevBlock, uid, amount, nonce, newOwner)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := signer.PlasmaSigner(signer.From, unsignedTx)
	if err != nil {
		t.Fatalf("failed to sign transaction %s", err)
	}

	addr, err := transaction.Sender(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(signer.From.Bytes(), addr.Bytes()) {
		t.Fatal("addresses not equal")
	}
	return tx
}

func testBlock(t *testing.T, txs ...*transaction.Transaction) block.Block {
	plasmaBlock := transactions.NewTxBlock()

	for _, tx := range txs {
		if err := plasmaBlock.AddTx(tx); err != nil {
			t.Fatal(err)
		}
	}

	_, err := plasmaBlock.Build()
	if err != nil {
		t.Fatal(err)
	}
	return plasmaBlock
}

func txToBytes(t *testing.T, tx *transaction.Transaction) []byte {
	buf := bytes.NewBuffer([]byte{})
	if err := tx.EncodeRLP(buf); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func newPlasmaTestTx(t *testing.T, i *instance, prevBlock, uid,
	amount *big.Int, nonce *big.Int, newOwner common.Address,
	signer *account.PlasmaTransactOpts) *PlasmaTestTx {

	tx := testTx(t, prevBlock, uid, amount, nonce, newOwner, signer)

	plasmaBlock := testBlock(t, tx)

	ethTX, err := i.rootOwnerSession.NewBlock(plasmaBlock.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTX) {
		t.Fatal("failed to create new block")
	}

	proof := plasmaBlock.CreateProof(uid)
	rawTx := txToBytes(t, tx)

	return &PlasmaTestTx{
		tx:    tx,
		block: plasmaBlock,
		proof: proof,
		rawTx: rawTx,
	}
}

func timeMachine(t *testing.T, adjustment time.Duration) {
	if sim, ok := server.(backend.Simulator); ok {
		if err := sim.AdjustTime(adjustment); err != nil {
			t.Fatal(err)
		}
	}
}

func TestMain(m *testing.M) {
	accounts := account.GenAccounts(3)
	owner = accounts[0]
	user1 = accounts[1]
	user2 = accounts[2]

	server = backend.NewSimulatedBackend(account.Addresses(accounts))

	os.Exit(m.Run())
}
