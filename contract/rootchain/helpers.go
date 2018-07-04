package rootchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartmeshfoundation/smartplasma/blockchan/backend"
	"log"
)

// NewRootChainSession func
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

// Deploy func
func Deploy(account *bind.TransactOpts,
	server backend.Backend) (common.Address,
	*types.Receipt, *RootChain, error) {
	addr, tx, contract, err := DeployRootChain(account,
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

// LogsDeposit func
func LogsDeposit(rootchain *RootChain) {
	iterator, err := rootchain.FilterDeposit(&bind.FilterOpts{})
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	for iterator.Next() {
		log.Printf("Depositor: %s, Amount: %s, Uid: %s",
			iterator.Event.Depositor.String(), iterator.Event.Amount.String(),
			iterator.Event.Uid.String())
	}

	if err := iterator.Error(); err != nil {
		log.Fatal(err)
	}
}
