package block

import (
	"errors"
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/common"
)

// Errors.
var (
	ErrAlreadyBuilt = errors.New("block is already built")
)

// Block defines the methods for abstract block.
type Block interface {
	Hash() common.Hash
	Build() (common.Hash, error)
	IsBuilt() bool
	CreateProof(uid *big.Int) []byte
	Marshal() ([]byte, error)
	Unmarshal(raw []byte) error
}
