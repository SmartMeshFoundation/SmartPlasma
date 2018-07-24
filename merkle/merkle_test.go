package merkle

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestCheckMembership(t *testing.T) {
	a := "1"
	b := "2"

	ha := crypto.Keccak256([]byte(a))
	hb := crypto.Keccak256([]byte(b))

	root := crypto.Keccak256(ha, hb)

	var proof []byte
	proof = append(proof, hb...)

	if !CheckMembership(big.NewInt(0), ha, root, proof) {
		t.Fatal("")
	}
}
