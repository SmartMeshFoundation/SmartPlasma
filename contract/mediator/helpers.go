package mediator

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartmeshfoundation/smartplasma/blockchan/backend"
)

// NewMediatorSession returns Mediator session
func NewMediatorSession(account bind.TransactOpts, contact common.Address,
	server backend.Backend) (*MediatorSession, error) {
	contract, err := NewMediator(contact, server.Connect())
	if err != nil {
		return nil, err
	}

	return &MediatorSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    account.From,
		},
		TransactOpts: account,
	}, err
}

// Deploy deploys Mediator contract
func Deploy(account *bind.TransactOpts,
	server backend.Backend) (common.Address, *Mediator, error) {
	addr, tx, contract, err := DeployMediator(account, server.Connect())
	if err != nil {
		return common.Address{}, nil, err
	}

	_, err = server.Mine(tx)
	if err != nil {
		return common.Address{}, nil, err
	}

	if !server.GoodTransaction(tx) {
		return common.Address{}, nil,
			errors.New("failed to deploy Mediator contract")
	}

	return addr, contract, nil
}
