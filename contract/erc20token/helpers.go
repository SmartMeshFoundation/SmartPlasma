package erc20token

import (
	"context"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
)

// NewExampleTokenSession returns Token session.
func NewExampleTokenSession(account bind.TransactOpts, contact common.Address,
	server backend.Backend) (*ExampleTokenSession, error) {
	contract, err := NewExampleToken(contact, server.Connect())
	if err != nil {
		return nil, err
	}

	return &ExampleTokenSession{
		Contract: contract,
		CallOpts: bind.CallOptsWithNumber{CallOpts: bind.CallOpts{
			Pending: true,
			From:    account.From,
		}},
		TransactOpts: account,
	}, err
}

// Deploy deploys Token contract.
func Deploy(account *bind.TransactOpts,
	server backend.Backend) (common.Address, *ExampleToken, error) {
	addr, tx, contract, err := DeployExampleToken(account,
		server.Connect())
	if err != nil {
		return common.Address{}, nil, err
	}

	_, err = server.Mine(context.Background(), tx)
	if err != nil {
		return common.Address{}, nil, err
	}

	if !server.GoodTransaction(tx) {
		return common.Address{}, nil,
			errors.New("failed to deploy Token contract")
	}

	return addr, contract, nil
}
