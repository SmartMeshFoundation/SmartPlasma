package backend

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
)

func newSimulator(accounts []common.Address) *backends.SimulatedBackend {
	alloc := make(core.GenesisAlloc)

	balance := new(big.Int)
	fmt.Sscan("1000000000000000000000000000000000000000000000000000", balance)

	for _, acc := range accounts {
		alloc[acc] = core.GenesisAccount{Balance: balance}
	}

	return backends.NewSimulatedBackend(alloc)
}
