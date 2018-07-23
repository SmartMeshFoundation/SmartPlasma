package backend

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

const (
	rinkeby = "https://rinkeby.infura.io/nanodzeck"
)

// Backend interface.
type Backend interface {
	Connect() bind.ContractBackend
	Mine(tx *types.Transaction) (*types.Receipt, error)
	GoodTransaction(tx *types.Transaction) bool
}

type backend struct {
	connect bind.ContractBackend
}

// NewBackend makes new Backend.
func NewBackend() Backend {
	cli, err := ethclient.Dial(rinkeby)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &backend{connect: cli}
}

// NewSimulatedBackend makes new backend simulator.
func NewSimulatedBackend(accounts []common.Address) Backend {
	return &backend{connect: newSimulator(accounts)}
}

// Connect gets connect to Ethereum backend.
func (back *backend) Connect() bind.ContractBackend {
	return back.connect
}

// Mine to wait mining.
func (back *backend) Mine(tx *types.Transaction) (*types.Receipt, error) {
	switch conn := back.connect.(type) {
	case *ethclient.Client:
		return bind.WaitMined(context.Background(), conn, tx)
	case *backends.SimulatedBackend:
		conn.Commit()
		return bind.WaitMined(context.Background(), conn, tx)
	}
	return nil, errors.New("backend is wrong")
}

// GoodTransaction returns true if transaction status = 1.
func (back *backend) GoodTransaction(tx *types.Transaction) bool {
	tr, err := back.Mine(tx)
	if err != nil {
		return false
	}
	if tr.Status != 1 {
		return false
	}
	return true
}
