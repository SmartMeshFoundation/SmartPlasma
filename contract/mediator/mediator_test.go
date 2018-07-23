package mediator

import (
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartmeshfoundation/smartplasma/blockchan/account"
	"github.com/smartmeshfoundation/smartplasma/blockchan/backend"
	"github.com/smartmeshfoundation/smartplasma/contract/erc20token"
	"github.com/smartmeshfoundation/smartplasma/contract/rootchain"
)

const (
	fakeString = "fake"
	one        = 1
)

var (
	server        backend.Backend
	accounts      []*bind.TransactOpts
	mediatorAddr  common.Address
	tokenAddr     common.Address
	rootChainAddr common.Address

	owner *bind.TransactOpts
	user1 *bind.TransactOpts
	user2 *bind.TransactOpts

	tokenOwnerSession   *erc20token.ExampleTokenSession
	tokenUserSession    *erc20token.ExampleTokenSession
	mediatorUserSession *MediatorSession
)

func newInstance(t *testing.T) {
	address, mediator, err := deployMediator(owner)
	if err != nil {
		t.Fatal(err)
	}

	mediatorAddr = address

	rootChainAddr, err = mediator.RootChain(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}

	tokenAddr, _ = deployToken(t, owner)

	tokenOwnerSession = tokenSession(t, owner, tokenAddr)
	tokenUserSession = tokenSession(t, user1, tokenAddr)
	mediatorUserSession = mediatorSession(t, user1, mediatorAddr)
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

func TestMediatorWithdraw(t *testing.T) {
	// TODO: to refactor after RootChain implementation
	newInstance(t)

	mint(t, tokenOwnerSession, user1.From, big.NewInt(one))
	increaseApproval(t, tokenUserSession, mediatorAddr, big.NewInt(one))

	deposit(t, mediatorUserSession, tokenAddr, big.NewInt(one), true)
	withdraw(t, mediatorUserSession, []byte(fakeString),
		[]byte(fakeString), big.NewInt(one), []byte(fakeString),
		[]byte(fakeString), big.NewInt(one), true)

	// withdraw test 1
	logs, err := erc20token.LogsTransfer(tokenOwnerSession.Contract)
	if err != nil {
		t.Fatal("failed to parse transfer logs")
	}

	if len(logs) != 3 {
		t.Fatal("invalid number of transfer transactions")
	}
	if logs[2].From.String() != mediatorAddr.String() ||
		logs[2].To.String() != user1.From.String() ||
		logs[2].Value.String() != big.NewInt(one).String() {
		t.Fatal("invalid withdraw")
	}

	// withdraw test 2
	withdraw(t, mediatorUserSession, []byte(fakeString),
		[]byte(fakeString), big.NewInt(one), []byte(fakeString),
		[]byte(fakeString), big.NewInt(one), false)
}

func TestMediatorDeposit(t *testing.T) {
	// TODO: to refactor after RootChain implementation
	newInstance(t)

	// deposit test1
	deposit(t, mediatorUserSession, tokenAddr, big.NewInt(one), false)
	mint(t, tokenOwnerSession, user1.From, big.NewInt(one))
	increaseApproval(t, tokenUserSession, mediatorAddr, big.NewInt(one))
	deposit(t, mediatorUserSession, tokenAddr, big.NewInt(one), true)

	rootSession := rootChainSession(t, user1, rootChainAddr)

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
	newInstance(t)

	// checkToken test1
	checkToken(t, mediatorUserSession, tokenAddr, false)
	mint(t, tokenOwnerSession, user2.From, big.NewInt(1))
	checkToken(t, mediatorUserSession, tokenAddr, false)

	// checkToken test2
	mint(t, tokenOwnerSession, user1.From, big.NewInt(1))
	checkToken(t, mediatorUserSession, tokenAddr, true)

	// checkToken test3
	// checkToken test4
	changeApproveState(t, tokenOwnerSession)
	checkToken(t, mediatorUserSession, tokenAddr, false)
	changeApproveState(t, tokenOwnerSession)

	// checkToken test5
	changeTransferFromState(t, tokenOwnerSession)
	checkToken(t, mediatorUserSession, tokenAddr, false)

	changeTransferFromState(t, tokenOwnerSession)

	// normal flow
	checkToken(t, mediatorUserSession, tokenAddr, true)
}

func TestMain(m *testing.M) {
	accounts = account.GenAccounts(3)
	owner = accounts[0]
	user1 = accounts[1]
	user2 = accounts[2]

	server = backend.NewSimulatedBackend(account.Addresses(accounts))

	os.Exit(m.Run())
}
