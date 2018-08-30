package backend

import (
	"fmt"
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind/backends"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core"
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
