package build

import (
	"context"
	"fmt"
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Transactor defines the methods needed to allow operating with contract
// on a write only basis.
type Transactor interface {
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error)
}

// Errors.
var (
	ErrNoCode = errors.New("no contract code at given address")
)

// Contract is the wrapper object that reflects a contract on the
// Ethereum network.
type Contract struct {
	abi        abi.ABI
	address    common.Address
	transactor Transactor
}

// NewContract creates a low level contract interface through which calls
// and transactions may be made through.
func NewContract(address common.Address, abi abi.ABI,
	backend Transactor) (*Contract, error) {
	if (address == common.Address{}) {
		return nil, errors.New("address is null")
	}
	return &Contract{
		abi:        abi,
		address:    address,
		transactor: backend, // TODO: change to plasma backend
	}, nil
}

// UnmarshalTransaction decodes raw transaction.
func (c *Contract) UnmarshalTransaction(raw []byte) (*types.Transaction, error) {
	tx := &types.Transaction{}
	err := tx.UnmarshalJSON(raw)
	if err != nil {
		return nil, err
	}
	if tx.To() == nil || *tx.To() != c.address {
		return nil, errors.New("wrong recipient address")
	}
	return tx, err
}

// Transaction creates signed transaction.
func (c *Contract) Transaction(opts *bind.TransactOpts, method string,
	params ...interface{}) (*types.Transaction, error) {
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return nil, err
	}

	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}

	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = c.transactor.PendingNonceAt(
			ensureContext(opts.Context), opts.From)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}

	gasPrice := opts.GasPrice
	if gasPrice == nil {
		gasPrice, err = c.transactor.SuggestGasPrice(
			ensureContext(opts.Context))
		if err != nil {
			return nil, fmt.Errorf(
				"failed to suggest gas price: %v", err)
		}
	}
	gasLimit := opts.GasLimit
	if gasLimit == 0 {
		if (c.address != common.Address{}) {
			if code, err := c.transactor.PendingCodeAt(
				ensureContext(opts.Context), c.address); err != nil {
				return nil, err
			} else if len(code) == 0 {
				return nil, ErrNoCode
			}
		}
		msg := ethereum.CallMsg{From: opts.From,
			To: &c.address, Value: value, Data: input}
		gasLimit, err = c.transactor.EstimateGas(
			ensureContext(opts.Context), msg)
		if err != nil {
			return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
		}
	}

	rawTx := types.NewTransaction(nonce, c.address, value,
		gasLimit, gasPrice, input)

	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}
	return opts.Signer(types.HomesteadSigner{}, opts.From, rawTx)
}

func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}
