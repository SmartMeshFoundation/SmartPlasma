package rootchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartmeshfoundation/smartplasma/blockchan/backend"
)

// NewRootChainSession returns RootChain session
func NewRootChainSession(account bind.TransactOpts, contact common.Address,
	server backend.Backend) (*RootChainSession, error) {
	contract, err := NewRootChain(contact, server.Connect())
	if err != nil {
		return nil, err
	}

	return &RootChainSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    account.From,
		},
		TransactOpts: account,
	}, err
}

// Deploy deploys RootChain contract
func Deploy(account *bind.TransactOpts,
	server backend.Backend) (common.Address, *RootChain, error) {
	addr, tx, contract, err := DeployRootChain(account,
		server.Connect())
	if err != nil {
		return common.Address{}, nil, err
	}

	_, err = server.Mine(tx)
	if err != nil {
		return common.Address{}, nil, err
	}

	if !server.GoodTransaction(tx) {
		return common.Address{}, nil,
			errors.New("failed to deploy RootChain contract")
	}

	return addr, contract, nil
}

// LogsDeposit returns deposit logs
func LogsDeposit(contract *RootChain) (logs []*RootChainDeposit, err error) {
	iterator, err2 := contract.FilterDeposit(&bind.FilterOpts{})
	if err2 != nil {
		err = err2
		return
	}

	defer iterator.Close()

	for iterator.Next() {
		logs = append(logs, iterator.Event)
	}

	if err2 := iterator.Error(); err2 != nil {
		err = err2
		return
	}

	return
}
