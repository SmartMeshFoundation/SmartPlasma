package mediator

import (
	"bytes"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartmeshfoundation/smartplasma/blockchan/account"
	"github.com/smartmeshfoundation/smartplasma/blockchan/backend"
	"github.com/smartmeshfoundation/smartplasma/blockchan/block"
	"github.com/smartmeshfoundation/smartplasma/blockchan/transaction"
	"github.com/smartmeshfoundation/smartplasma/contract/erc20token"
	"github.com/smartmeshfoundation/smartplasma/contract/rootchain"
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

func testBlock(t *testing.T, txs ...*transaction.Transaction) *block.Block {
	plasmaBlock := block.NewBlock()

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

func TestMediatorNormalFlow(t *testing.T) {
	i := newInstance(t)

	uid := testDeposit(t, i)

	tx := testTx(t, zero, uid, one, zero, user1.From, user1)

	plasmaBlock1 := testBlock(t, tx)

	ethTX, err := i.rootOwnerSession.NewBlock(plasmaBlock1.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTX) {
		t.Fatal("failed to create new block")
	}

	// create proof to transaction #1
	proof1 := plasmaBlock1.CreateProof(uid)

	// validate transaction #1
	validateTX1 := block.CheckMembership(uid, tx.Hash(),
		plasmaBlock1.Hash(), proof1)
	if !validateTX1 {
		t.Fatal("tx1 invalid")
	}

	tx2 := testTx(t, one, uid, one, one, user2.From, user1)

	plasmaBlock2 := testBlock(t, tx2)

	ethTX2, err := i.rootOwnerSession.NewBlock(plasmaBlock2.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTX2) {
		t.Fatal("failed to create new block")
	}

	// create proof to transaction #2
	proof2 := plasmaBlock2.CreateProof(uid)

	// validate transaction #2
	validateTX2 := block.CheckMembership(uid, tx2.Hash(),
		plasmaBlock2.Hash(), proof2)
	if !validateTX2 {
		t.Fatal("tx2 invalid")
	}

	rawTX1 := txToBytes(t, tx)
	rawTX2 := txToBytes(t, tx2)

	ethTX3, err := i.rootUser2Session.StartExit(rawTX1, proof1, one,
		rawTX2, proof2, two)
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

	withdraw(t, i.mediatorUser2Session, rawTX1, proof1,
		one, rawTX2, proof2, two, true)

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

	exits2, err := i.rootUser2Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits2.State.Int64() != 3 {
		t.Fatal("The exit operation did not finish")
	}

	// withdraw test 2
	withdraw(t, i.mediatorUser2Session, rawTX1, proof1,
		one, rawTX2, proof2, two, false)
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
