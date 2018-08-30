package rootchain

import (
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"github.com/pkg/errors"
	"math/big"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
)

// NewRootChainSession returns RootChain session.
func NewRootChainSession(account bind.TransactOpts, contact common.Address,
	server backend.Backend) (*RootChainSession, error) {
	contract, err := NewRootChain(contact, server.Connect())
	if err != nil {
		return nil, err
	}

	return &RootChainSession{
		Contract: contract,
		CallOpts: bind.CallOptsWithNumber{CallOpts: bind.CallOpts{
			Pending: true,
			From:    account.From,
		}},
		TransactOpts: account,
	}, err
}

// Deploy deploys RootChain contract.
func Deploy(account *bind.TransactOpts,
	server backend.Backend) (common.Address, *RootChain, error) {
	addr, tx, contract, err := DeployRootChain(account,
		server.Connect(), account.From)
	if err != nil {
		return common.Address{}, nil, err
	}

	if !server.GoodTransaction(tx) {
		return common.Address{}, nil,
			errors.New("failed to deploy RootChain contract")
	}

	return addr, contract, nil
}

// GenerateNextUID generates the following UID. Usages for tests only.
func GenerateNextUID(session *RootChainSession,
	address common.Address, token common.Address) (*big.Int, error) {
	depositCount, err := session.DepositCount()
	if err != nil {
		return nil, err
	}

	uidData := crypto.Keccak256(token.Bytes(), address.Bytes(),
		common.BigToHash(depositCount).Bytes())

	return new(big.Int).SetBytes(uidData), nil
}
