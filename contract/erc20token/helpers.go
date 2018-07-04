package erc20token

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartmeshfoundation/smartplasma/blockchan/backend"
	"log"
)

// NewExampleTokenSession func
func NewExampleTokenSession(account bind.TransactOpts, contact common.Address,
	server backend.Backend) (*ExampleTokenSession, error) {
	contract, err := NewExampleToken(contact, server.Connect())
	if err != nil {
		return nil, err
	}

	return &ExampleTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    account.From,
		},
		TransactOpts: account,
	}, err
}

// Deploy func
func Deploy(account *bind.TransactOpts,
	server backend.Backend) (common.Address,
	*types.Receipt, *ExampleToken, error) {
	addr, tx, contract, err := DeployExampleToken(account,
		server.Connect())
	if err != nil {
		return [20]byte{}, nil, nil, err
	}

	tr, err := server.Mine(tx)
	if err != nil {
		return [20]byte{}, nil, nil, err
	}

	return addr, tr, contract, nil
}

// LogsApproval func
func LogsApproval(token *ExampleToken) {
	iterator, err := token.FilterApproval(&bind.FilterOpts{},
		[]common.Address{}, []common.Address{})
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	for iterator.Next() {
		log.Printf("Owner: %s, Spender: %s, Value: %s",
			iterator.Event.Owner.String(), iterator.Event.Spender.String(),
			iterator.Event.Value.String())
	}

	if err := iterator.Error(); err != nil {
		log.Fatal(err)
	}
}

// LogsTransfer func
func LogsTransfer(token *ExampleToken) {
	iterator, err := token.FilterTransfer(&bind.FilterOpts{},
		[]common.Address{}, []common.Address{})
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	for iterator.Next() {
		log.Printf("From: %s, To: %s, Value: %s", iterator.Event.From.String(),
			iterator.Event.To.String(), iterator.Event.Value.String())
	}

	if err := iterator.Error(); err != nil {
		log.Fatal(err)
	}
}
