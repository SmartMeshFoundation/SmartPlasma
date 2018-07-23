package erc20token

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartmeshfoundation/smartplasma/blockchan/backend"
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
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    account.From,
		},
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

	_, err = server.Mine(tx)
	if err != nil {
		return common.Address{}, nil, err
	}

	if !server.GoodTransaction(tx) {
		return common.Address{}, nil,
			errors.New("failed to deploy Token contract")
	}

	return addr, contract, nil
}

// LogsApproval returns approval logs.
func LogsApproval(token *ExampleToken) (logs []*ExampleTokenApproval,
	err error) {
	iterator, err2 := token.FilterApproval(&bind.FilterOpts{},
		[]common.Address{}, []common.Address{})
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

// LogsTransfer returns transfer logs.
func LogsTransfer(token *ExampleToken) (logs []*ExampleTokenTransfer,
	err error) {
	iterator, err2 := token.FilterTransfer(&bind.FilterOpts{},
		[]common.Address{}, []common.Address{})
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
