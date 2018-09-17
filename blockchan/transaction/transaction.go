package transaction

import (
	"crypto/ecdsa"
	"io"
	"math/big"

	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"github.com/SmartMeshFoundation/Spectrum/crypto/sha3"
	"github.com/SmartMeshFoundation/Spectrum/rlp"
)

// Errors transaction-related errors.
var (
	ErrInvalidSig           = errors.New("invalid transaction v, r, s values")
	ErrInvalidArguments     = errors.New("invalid arguments")
	ErrInvalidPreviousBlock = errors.New("invalid previous block number")
	ErrInvalidNewOwner      = errors.New("invalid new owner address")
	ErrInvalidUID           = errors.New("invalid UID")
	ErrInvalidPrivateKey    = errors.New("invalid private key")
	ErrInvalidPublicKey     = errors.New("invalid public key")
	ErrInvalidTx            = errors.New("invalid transaction")
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
	Nonce     *big.Int       `json:"nonce"`
	Sig       []byte         `json:"sig"`
}

// NewTransaction creates new unsigned transaction.
func NewTransaction(prevBlock, uid, amount, nonce *big.Int,
	newOwner common.Address) (*Transaction, error) {
	if prevBlock == nil || uid == nil || amount == nil || nonce == nil {
		return nil, ErrInvalidArguments
	}

	if prevBlock == nil || prevBlock.Cmp(big.NewInt(-1)) < 0 {
		return nil, ErrInvalidPreviousBlock
	}
	if (newOwner.Hash() == common.Hash{}) {
		return nil, ErrInvalidNewOwner
	}

	if uid == nil || uid.Int64() == 0 {
		return nil, ErrInvalidUID
	}

	return &Transaction{
		data: txData{
			PrevBlock: prevBlock,
			UID:       uid,
			Amount:    amount,
			NewOwner:  newOwner,
			Nonce:     nonce,
		},
	}, nil
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
		tx.data.Nonce,
	})
}

// UID returns uid from the transaction.
func (tx *Transaction) UID() *big.Int {
	return tx.data.UID
}

// SignatureValues returns signature values.
func SignatureValues(sig []byte) (r, s, v *big.Int, err error) {
	if len(sig) != 65 {
		return nil, nil, nil, ErrInvalidSig
	}
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	return r, s, v, nil
}

// SignTx signs the transaction using and a private key.
func (tx *Transaction) SignTx(key *ecdsa.PrivateKey) (*Transaction, error) {
	if key == nil {
		return nil, ErrInvalidPrivateKey
	}

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
	if tx == nil {
		return ErrInvalidTx
	}
	return rlp.Decode(r, &tx.data)
}

// Sender returns the address derived from the signature (V, R, S) using
// secp256k1 elliptic curve and an error if it failed deriving
// or upon an incorrect signature.
func Sender(tx *Transaction) (common.Address, error) {
	if tx == nil {
		return common.Address{}, ErrInvalidTx
	}

	r, s, v, err := SignatureValues(tx.data.Sig)
	if err != nil {
		return common.Address{}, err
	}
	return recoverPlain(tx.Hash(), r, s, v)
}

func recoverPlain(sighash common.Hash, R, S,
	Vb *big.Int) (common.Address, error) {
	if R == nil || S == nil || Vb == nil {
		return common.Address{}, ErrInvalidArguments
	}

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
		return common.Address{}, ErrInvalidPublicKey
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(pub[1:])[12:])
	return addr, nil
}
