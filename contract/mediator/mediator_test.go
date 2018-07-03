package mediator

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartmeshfoundation/smartplasma/blockchan/account"
	"github.com/smartmeshfoundation/smartplasma/blockchan/backend"
	"github.com/smartmeshfoundation/smartplasma/contract/erc20token"
	"log"
	"math/big"
	"os"
	"testing"
)

var (
	server       backend.Backend
	accounts     []*bind.TransactOpts
	mediatorAddr common.Address

	owner *bind.TransactOpts
	user1 *bind.TransactOpts
	user2 *bind.TransactOpts
)

func mint(t *testing.T, session *erc20token.ExampleTokenSession,
	acc common.Address, val *big.Int) {
	tx, err := session.Mint(acc, big.NewInt(1))
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to mint tokens")
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

func TestDeposit(t *testing.T) {
	tokenAddr, _, token, err := erc20token.Deploy(owner, server)
	if err != nil {
		log.Fatal(err)
	}

	tokenOwnerSess, err := erc20token.NewExampleTokenSession(*owner,
		tokenAddr, server)
	if err != nil {
		log.Fatal(err)
	}

	mint(t, tokenOwnerSess, user1.From, big.NewInt(1))

	tokenUserSess, err := erc20token.NewExampleTokenSession(*user1,
		tokenAddr, server)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := tokenUserSess.IncreaseApproval(mediatorAddr, big.NewInt(1))
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to approval tokens")
	}

	sessU1, err := NewMediatorSession(*user1, mediatorAddr, server)
	if err != nil {
		log.Fatal(err)
	}

	tx, err = sessU1.Deposit(tokenAddr, big.NewInt(1))
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to transfer tokens")
	}

	erc20token.LogsTransfer(token)
}

func TestMediatorSession_CheckToken(t *testing.T) {
	tokenAddr, _, _, err := erc20token.Deploy(owner, server)
	if err != nil {
		log.Fatal(err)
	}

	sessU1, err := NewMediatorSession(*user1, mediatorAddr, server)
	if err != nil {
		log.Fatal(err)
	}

	tokenSess, err := erc20token.NewExampleTokenSession(*owner,
		tokenAddr, server)
	if err != nil {
		log.Fatal(err)
	}

	// checkToken test1
	checkToken(t, sessU1, tokenAddr, false)

	mint(t, tokenSess, user2.From, big.NewInt(1))

	checkToken(t, sessU1, tokenAddr, false)

	// checkToken test2
	mint(t, tokenSess, user1.From, big.NewInt(1))
	checkToken(t, sessU1, tokenAddr, true)

	// checkToken test3
	// checkToken test4
	tx, err := tokenSess.ChangeApproveState()
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to mint tokens")
	}

	checkToken(t, sessU1, tokenAddr, false)

	tx, err = tokenSess.ChangeApproveState()
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to mint tokens")
	}

	// checkToken test5
	tx, err = tokenSess.ChangeTransferFromState()
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to mint tokens")
	}

	checkToken(t, sessU1, tokenAddr, false)

	tx, err = tokenSess.ChangeTransferFromState()
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to mint tokens")
	}

	// normal flow
	checkToken(t, sessU1, tokenAddr, true)
}

func TestMain(m *testing.M) {
	accounts = account.GenAccounts(3)
	owner = accounts[0]
	user1 = accounts[1]
	user2 = accounts[2]

	server = backend.NewSimulatedBackend(account.Addresses(accounts))

	address, tr, _, err := Deploy(owner, server)
	if err != nil {
		log.Fatal(err)
	}

	if tr.Status != 1 {
		log.Fatal("mediator contract not deployed")
	}

	mediatorAddr = address

	os.Exit(m.Run())
}
