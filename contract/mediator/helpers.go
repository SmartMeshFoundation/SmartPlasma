package mediator

import (
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
)

// NewMediatorSession returns Mediator session.
func NewMediatorSession(account bind.TransactOpts, contact common.Address,
	server backend.Backend) (*MediatorSession, error) {
	contract, err := NewMediator(contact, server.Connect())
	if err != nil {
		return nil, err
	}

	return &MediatorSession{
		Contract: contract,
		CallOpts: bind.CallOptsWithNumber{CallOpts: bind.CallOpts{
			Pending: true,
			From:    account.From,
		}},
		TransactOpts: account,
	}, err
}

// Deploy deploys Mediator contract.
func Deploy(account *bind.TransactOpts,
	server backend.Backend) (common.Address, *Mediator, error) {
	addr, tx, contract, err := DeployMediator(account, server.Connect())
	if err != nil {
		return common.Address{}, nil, err
	}

	if !server.GoodTransaction(tx) {
		return common.Address{}, nil,
			errors.New("failed to deploy Mediator contract")
	}

	plasmaAddr, plasmaContract, err := rootchain.Deploy(account, server)
	if err != nil {
		return common.Address{}, nil, err
	}

	tx2, err := plasmaContract.TransferOwnership(account, addr)
	if err != nil {
		return common.Address{}, nil, err
	}

	if !server.GoodTransaction(tx2) {
		return common.Address{}, nil,
			errors.New("failed to transfer ownership")
	}

	joinTx, err := contract.JoinPlasma(account, plasmaAddr)
	if err != nil {
		return common.Address{}, nil, err
	}
	if !server.GoodTransaction(joinTx) {
		return common.Address{}, nil,
			errors.New("failed to join Plasma")
	}

	return addr, contract, nil
}
