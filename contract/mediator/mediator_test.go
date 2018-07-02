package mediator

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartmeshfoundation/smartplasma/blockchan/backend"
	"github.com/smartmeshfoundation/smartplasma/contract/erc20token"
	"log"
	"math/big"
	"os"
	"testing"
)

var (
	server   backend.Backend
	accounts []*bind.TransactOpts
	session  *MediatorSession
)

func TestMediatorSession_CheckToken(t *testing.T) {
	address, _, _, err := erc20token.Deploy(accounts[0], server)
	if err != nil {
		log.Fatal(err)
	}

	valid, _ := session.CheckToken(address)
	if valid {
		t.Fatal("wrong result")
	}

	tokenSess, err := erc20token.MewExampleTokenSession(*accounts[0],
		address, server)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := tokenSess.Mint(accounts[0].From, big.NewInt(1))
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to mint tokens")
	}

	valid, _ = session.CheckToken(address)
	if !valid {
		t.Fatal("wrong result")
	}
}

func TestMain(m *testing.M) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	accounts = []*bind.TransactOpts{bind.NewKeyedTransactor(key)}

	var adders []common.Address
	for _, a := range accounts {
		adders = append(adders, a.From)
	}

	server = backend.NewSimulatedBackend(adders)

	address, tr, _, err := Deploy(accounts[0], server)
	if err != nil {
		log.Fatal(err)
	}

	if tr.Status != 1 {
		log.Fatal("mediator contract not deployed")
	}

	session, err = MewMediatorSession(*accounts[0], address, server)

	os.Exit(m.Run())
}
