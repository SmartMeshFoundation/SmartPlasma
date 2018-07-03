package mediator

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/smartmeshfoundation/smartplasma/blockchan/account"
	"github.com/smartmeshfoundation/smartplasma/blockchan/backend"
	"github.com/smartmeshfoundation/smartplasma/contract/erc20token"
	"log"
	"math/big"
	"os"
	"testing"
	"github.com/ethereum/go-ethereum/common"
)

var (
	server   backend.Backend
	accounts []*bind.TransactOpts
	session  *MediatorSession
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

func checkToken(t *testing.T, token common.Address, expected bool) {
	valid, _ := session.CheckToken(token)
	if expected && !valid {
		t.Fatal("function must return - true")
	}

	if !expected && valid {
		t.Fatal("function must return - false")
	}
}

func TestMediatorSession_CheckToken(t *testing.T) {
	address, _, _, err := erc20token.Deploy(accounts[0], server)
	if err != nil {
		log.Fatal(err)
	}

	// checkToken test1
	checkToken(t, address, false)

	tokenSess, err := erc20token.NewExampleTokenSession(*accounts[0],
		address, server)
	if err != nil {
		log.Fatal(err)
	}

	mint(t, tokenSess, accounts[1].From, big.NewInt(1))

	// checkToken test2
	checkToken(t, address, false)
	mint(t, tokenSess, accounts[0].From, big.NewInt(1))

	// checkToken test3
	// checkToken test4
	tx, err := tokenSess.ChangeApproveState()
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to mint tokens")
	}

	checkToken(t, address, false)

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

	checkToken(t, address, false)

	tx, err = tokenSess.ChangeTransferFromState()
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to mint tokens")
	}

	// normal flow
	checkToken(t, address, true)
}

func TestMain(m *testing.M) {
	accounts = account.GenAccounts(2)

	server = backend.NewSimulatedBackend(account.Addresses(accounts))

	address, tr, _, err := Deploy(accounts[0], server)
	if err != nil {
		log.Fatal(err)
	}

	if tr.Status != 1 {
		log.Fatal("mediator contract not deployed")
	}

	session, err = NewMediatorSession(*accounts[0], address, server)

	os.Exit(m.Run())
}
