package mediator

import (
	"bytes"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/erc20token"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
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

	mediatorUser1Session *MediatorSession
	mediatorUser2Session *MediatorSession

	rootOwnerSession *rootchain.RootChainSession
	rootUser1Session *rootchain.RootChainSession
	rootUser2Session *rootchain.RootChainSession

	mediatorAddr  common.Address
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

	address, mediator, err := deployMediator(owner.TransactOpts)
	if err != nil {
		t.Fatal(err)
	}

	i.mediatorAddr = address
	i.rootChainAddr, err = mediator.RootChain(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}

	i.tokenAddr, _ = deployToken(t, owner.TransactOpts)

	i.tokenOwnerSession = tokenSession(t, owner.TransactOpts, i.tokenAddr)
	i.tokenUser1Session = tokenSession(t, user1.TransactOpts, i.tokenAddr)

	i.mediatorUser1Session = mediatorSession(t, user1.TransactOpts,
		i.mediatorAddr)
	i.mediatorUser2Session = mediatorSession(t, user2.TransactOpts,
		i.mediatorAddr)

	i.rootOwnerSession = rootChainSession(t, owner.TransactOpts,
		i.rootChainAddr)
	i.rootUser1Session = rootChainSession(t, user1.TransactOpts,
		i.rootChainAddr)
	i.rootUser2Session = rootChainSession(t, user2.TransactOpts,
		i.rootChainAddr)
	return i
}

func timeMachine(t *testing.T, adjustment time.Duration) {
	if sim, ok := server.(backend.Simulator); ok {
		if err := sim.AdjustTime(adjustment); err != nil {
			t.Fatal(err)
		}
	}
}

func deposit(t *testing.T, session *MediatorSession,
	token common.Address, amount *big.Int, executable bool) {
	tx, err := session.Deposit(token, amount)
	if !executable {
		if err == nil {
			t.Fatal("It should not be executed")
		}
		return
	}
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(tx) {
		t.Fatal("failed to deposit tokens")
	}
}

func withdraw(t *testing.T, session *MediatorSession,
	prevTx []byte, prevTxProof []byte, prevTxBlkNum *big.Int, txRaw []byte,
	txProof []byte, txBlkNum *big.Int, executable bool) {
	tx, err := session.Withdraw(prevTx, prevTxProof, prevTxBlkNum, txRaw,
		txProof, txBlkNum)
	if !executable {
		if err == nil {
			t.Fatal("It should not be executed")
		}
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to withdraw tokens")
	}
}

func mint(t *testing.T, session *erc20token.ExampleTokenSession,
	acc common.Address, val *big.Int) {
	tx, err := session.Mint(acc, val)
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to mint tokens")
	}
}

func increaseApproval(t *testing.T, session *erc20token.ExampleTokenSession,
	spender common.Address, addedValue *big.Int) {
	tx, err := session.IncreaseApproval(spender, addedValue)
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to approval tokens")
	}
}

func checkToken(t *testing.T, mediatorSession *MediatorSession,
	token common.Address, expected bool) {
	valid, _ := mediatorSession.CheckToken(token)
	if expected && !valid {
		t.Fatal("function must return - true")
	}

	if !expected && valid {
		t.Fatal("function must return - false")
	}
}

func tokenSession(t *testing.T, account *bind.TransactOpts,
	contact common.Address) (session *erc20token.ExampleTokenSession) {
	session, err := erc20token.NewExampleTokenSession(*account,
		contact, server)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func rootChainSession(t *testing.T, account *bind.TransactOpts,
	contact common.Address) (session *rootchain.RootChainSession) {
	session, err := rootchain.NewRootChainSession(*account,
		contact, server)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func mediatorSession(t *testing.T, account *bind.TransactOpts,
	contact common.Address) (session *MediatorSession) {
	session, err := NewMediatorSession(*account,
		contact, server)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func deployToken(t *testing.T,
	account *bind.TransactOpts) (address common.Address,
	contract *erc20token.ExampleToken) {
	address, contract, err := erc20token.Deploy(account, server)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func deployMediator(account *bind.TransactOpts) (address common.Address,
	contract *Mediator, err error) {
	address, contract, err = Deploy(account, server)
	return
}

func changeApproveState(t *testing.T,
	session *erc20token.ExampleTokenSession) {
	tx, err := session.ChangeApproveState()
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to change Approve state")
	}
}

func changeTransferFromState(t *testing.T,
	session *erc20token.ExampleTokenSession) {
	tx, err := session.ChangeTransferFromState()
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to change TransferFrom state")
	}
}

func testDeposit(t *testing.T, i *instance) (uid *big.Int) {
	mint(t, i.tokenOwnerSession, user1.From, one)
	increaseApproval(t, i.tokenUser1Session, i.mediatorAddr, one)

	// deposit
	deposit(t, i.mediatorUser1Session, i.tokenAddr, one, true)

	// receive logs with deposit. Get uid
	logs, err := rootchain.LogsDeposit(i.rootOwnerSession.Contract)
	if err != nil {
		t.Fatalf("failed to parse deposit logs %s", err)
	}

	// TODO: single deposit. Not applicable for multiple deposits.
	if len(logs) != 1 {
		t.Fatal("wrong number of logs")
	}

	return logs[0].Uid
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

func TestMediatorNormalFlow(t *testing.T) {
	i := newInstance(t)

	uid := testDeposit(t, i)

	tx1 := newPlasmaTestTx(t, i, zero, uid, one, zero, user1.From, user1)
	tx2 := newPlasmaTestTx(t, i, one, uid, one, one, user2.From, user1)

	ethTX3, err := i.rootUser2Session.StartExit(tx1.rawTx, tx1.proof, one,
		tx2.rawTx, tx2.proof, two)
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(ethTX3) {
		t.Fatal("failed to start exit")
	}

	exits1, err := i.rootUser2Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits1.State.Int64() != 2 {
		t.Fatal("The exit operation did not start")
	}

	// + 3 week
	timeMachine(t, time.Duration(504*time.Hour))

	withdraw(t, i.mediatorUser2Session, tx1.rawTx, tx1.proof,
		one, tx2.rawTx, tx2.proof, two, true)

	// withdraw test 1
	logsTransfer, err := erc20token.LogsTransfer(i.tokenOwnerSession.Contract)
	if err != nil {
		t.Fatal("failed to parse transfer logs")
	}

	if len(logsTransfer) != 3 {
		t.Fatal("invalid number of transfer transactions")
	}

	if logsTransfer[2].From.String() != i.mediatorAddr.String() ||
		logsTransfer[2].To.String() != user2.From.String() ||
		logsTransfer[2].Value.Int64() != one.Int64() {
		t.Fatal("invalid withdraw")
	}

	amount, err := i.rootUser1Session.Wallet(common.BigToHash(uid))
	if err != nil {
		t.Fatal(err)
	}

	if amount.Int64() != 0 {
		t.Fatal("the coin is not spent")
	}

	exits2, err := i.rootUser2Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits2.State.Int64() != 3 {
		t.Fatal("The exit operation did not finish")
	}

	// withdraw test 2
	withdraw(t, i.mediatorUser2Session, tx1.rawTx, tx1.proof,
		one, tx2.rawTx, tx2.proof, two, false)
}

func TestMediatorDeposit(t *testing.T) {
	i := newInstance(t)

	// deposit test1
	deposit(t, i.mediatorUser1Session, i.tokenAddr, one, false)
	mint(t, i.tokenOwnerSession, user1.From, one)
	increaseApproval(t, i.tokenUser1Session, i.mediatorAddr, one)
	deposit(t, i.mediatorUser1Session, i.tokenAddr, one, true)

	rootSession := rootChainSession(t, user1.TransactOpts, i.rootChainAddr)

	// deposit test2
	logs, err := rootchain.LogsDeposit(rootSession.Contract)
	if err != nil {
		t.Fatal("failed to parse deposit logs")
	}
	if len(logs) != 1 {
		t.Fatal("invalid number of deposit transactions")
	}
}

func TestMediatorCheckToken(t *testing.T) {
	i := newInstance(t)

	// checkToken test1
	checkToken(t, i.mediatorUser1Session, i.tokenAddr, false)
	mint(t, i.tokenOwnerSession, user2.From, one)
	checkToken(t, i.mediatorUser1Session, i.tokenAddr, false)

	// checkToken test2
	mint(t, i.tokenOwnerSession, user1.From, one)
	checkToken(t, i.mediatorUser1Session, i.tokenAddr, true)

	// checkToken test3
	// checkToken test4
	changeApproveState(t, i.tokenOwnerSession)
	checkToken(t, i.mediatorUser1Session, i.tokenAddr, false)
	changeApproveState(t, i.tokenOwnerSession)

	// checkToken test5
	changeTransferFromState(t, i.tokenOwnerSession)
	checkToken(t, i.mediatorUser1Session, i.tokenAddr, false)

	changeTransferFromState(t, i.tokenOwnerSession)

	// normal flow
	checkToken(t, i.mediatorUser1Session, i.tokenAddr, true)
}

func TestMain(m *testing.M) {
	accounts := account.GenAccounts(3)
	owner = accounts[0]
	user1 = accounts[1]
	user2 = accounts[2]

	server = backend.NewSimulatedBackend(account.Addresses(accounts))

	os.Exit(m.Run())
}
