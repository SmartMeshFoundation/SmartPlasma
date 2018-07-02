package mediator

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartmeshfoundation/smartplasma/blockchan/backend"
)

// MewMediatorSession function
func MewMediatorSession(account bind.TransactOpts, contact common.Address,
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

// Deploy function
func Deploy(account *bind.TransactOpts, server backend.Backend) (common.Address,
	*types.Receipt, *Mediator, error) {
	addr, tx, contract, err := DeployMediator(account, server.Connect())
	if err != nil {
		return [20]byte{}, nil, nil, err
	}

	tr, err := server.Mine(tx)
	if err != nil {
		return [20]byte{}, nil, nil, err
	}

	return addr, tr, contract, nil
}
