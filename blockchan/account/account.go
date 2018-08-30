package account

import (
	"crypto/ecdsa"
	"log"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
)

// PlasmaSignerFn is a signer function callback when requires a method
// to sign the Plasma cash transaction before submission.
type PlasmaSignerFn func(address common.Address,
	tx *transaction.Transaction) (*transaction.Transaction, error)

// PlasmaTransactOpts is the collection of authorization data required
// to create a valid Plasma Cash transaction.
type PlasmaTransactOpts struct {
	PlasmaSigner PlasmaSignerFn
	*bind.TransactOpts
}

// GenKey makes new private key.
func GenKey() *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	return key
}

// Account makes new account.
func Account(key *ecdsa.PrivateKey) *PlasmaTransactOpts {
	return NewPlasmaKeyedTransactor(key)
}

// GenAccounts makes new accounts.
func GenAccounts(number int) (accounts []*PlasmaTransactOpts) {
	for i := 0; i < number; i++ {
		accounts = append(accounts, Account(GenKey()))
	}
	return
}

// Address makes new public key.
func Address(acc *PlasmaTransactOpts) common.Address {
	return acc.From
}

// Addresses extracts public keys from accounts.
func Addresses(accounts []*PlasmaTransactOpts) (addresses []common.Address) {
	for _, a := range accounts {
		addresses = append(addresses, Address(a))
	}
	return
}

// NewPlasmaKeyedTransactor is a utility method to easily create
// a transaction signer from a single private key.
func NewPlasmaKeyedTransactor(key *ecdsa.PrivateKey) *PlasmaTransactOpts {
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)

	return &PlasmaTransactOpts{
		TransactOpts: bind.NewKeyedTransactor(key),
		PlasmaSigner: func(address common.Address,
			tx *transaction.Transaction) (*transaction.Transaction, error) {
			if address != keyAddr {
				return nil, errors.New("not authorized to sign this account")
			}
			return tx.SignTx(key)
		},
	}
}
