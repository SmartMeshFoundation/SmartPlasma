package block

import (
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/common"
)

// Block defines the methods for abstract block.
type Block interface {
	Hash() common.Hash
	Build() (common.Hash, error)
	CreateProof(uid *big.Int) []byte
	Marshal() ([]byte, error)
	Unmarshal(raw []byte) error
}
