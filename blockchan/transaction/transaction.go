package transaction

import (
	"crypto/ecdsa"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
)

// Errors transaction-related errors.
var (
	ErrInvalidSig = errors.New("invalid transaction v, r, s values")
)

// Transaction structure.
type Transaction struct {
	data txData
}

type txData struct {
	PrevBlock *big.Int       `json:"prevBlock"`
	UID       *big.Int       `json:"uid"`
	Amount    *big.Int       `json:"amount"`
	NewOwner  common.Address `json:"newOwner"`
	Sig       []byte         `json:"sig"`
}

// NewTransaction creates new unsigned transaction.
func NewTransaction(prevBlock, uid, amount *big.Int,
	newOwner common.Address) *Transaction {
	// TODO: check args
	return &Transaction{
		data: txData{
			PrevBlock: prevBlock,
			UID:       uid,
			Amount:    amount,
			NewOwner:  newOwner,
		},
	}
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

// Hash returns hash of transaction.
func (tx *Transaction) Hash() common.Hash {
	return rlpHash([]interface{}{
		tx.data.PrevBlock,
		tx.data.UID,
		tx.data.Amount,
		tx.data.NewOwner,
	})
}

// SignatureValues returns signature values.
func SignatureValues(sig []byte) (r, s, v *big.Int, err error) {
	if len(sig) != 65 {
		return nil, nil, nil, errors.New(
			fmt.Sprintf("wrong size for signature:"+
				" got %d, want 65", len(sig)))
	}
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	return r, s, v, nil
}

// SignTx signs the transaction using and a private key.
func (tx *Transaction) SignTx(key *ecdsa.PrivateKey) (*Transaction, error) {
	sig, err := crypto.Sign(tx.Hash().Bytes(), key)
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(sig)
}

// WithSignature returns a new transaction with the given signature.
func (tx *Transaction) WithSignature(sig []byte) (*Transaction, error) {
	cpy := &Transaction{data: tx.data}
	cpy.data.Sig = sig
	return cpy, nil
}

// EncodeRLP implements rlp.Encoder
func (tx *Transaction) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &tx.data)
}

// DecodeRLP implements rlp.Decoder
func DecodeRLP(r io.Reader, tx *Transaction) error {
	return rlp.Decode(r, &tx.data)
}

// Sender returns the address derived from the signature (V, R, S) using
// secp256k1 elliptic curve and an error if it failed deriving
// or upon an incorrect signature.
func Sender(tx *Transaction) (common.Address, error) {
	r, s, v, err := SignatureValues(tx.data.Sig)
	if err != nil {
		return common.Address{}, err
	}
	return recoverPlain(tx.Hash(), r, s, v)
}

func recoverPlain(sighash common.Hash, R, S,
	Vb *big.Int) (common.Address, error) {
	if Vb.BitLen() > 8 {
		return common.Address{}, ErrInvalidSig
	}
	V := byte(Vb.Uint64() - 27)

	// encode the snature in uncompressed format
	r, s := R.Bytes(), S.Bytes()
	sig := make([]byte, 65)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = V
	// recover the public key from the snature
	pub, err := crypto.Ecrecover(sighash[:], sig)
	if err != nil {
		return common.Address{}, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return common.Address{}, errors.New("invalid public key")
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(pub[1:])[12:])
	return addr, nil
}
