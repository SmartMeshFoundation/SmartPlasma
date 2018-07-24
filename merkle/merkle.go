package merkle

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

func CheckMembership(index *big.Int, leaf, rootHash []byte,
	proof []byte) bool {
	if len(proof) == 0 || len(proof)%32 != 0 {
		return false
	}

	computedHash := leaf
	// TODO: bigInt
	for i := 0; i < len(proof); i += 32 {
		proofElement := proof[i : i+32]

		if new(big.Int).Rem(index, big.NewInt(2)).Uint64() == 0 {
			computedHash = crypto.Keccak256(computedHash, proofElement)
		} else {
			computedHash = crypto.Keccak256(proofElement, computedHash)
		}
		index = index.Div(index, big.NewInt(2))
	}
	return bytes.Equal(computedHash, rootHash)
}
