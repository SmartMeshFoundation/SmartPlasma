package account

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// GenKey makes new private key
func GenKey() *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	return key
}

// Account makes new account
func Account(key *ecdsa.PrivateKey) *bind.TransactOpts {
	return bind.NewKeyedTransactor(key)
}

// GenAccounts makes new accounts
func GenAccounts(number int) (accounts []*bind.TransactOpts) {
	for i := 0; i < number; i++ {
		accounts = append(accounts, Account(GenKey()))
	}
	return
}

// Address makes new public key
func Address(acc *bind.TransactOpts) common.Address {
	return acc.From
}

// Addresses extracts public keys from accounts
func Addresses(accounts []*bind.TransactOpts) (addresses []common.Address) {
	for _, a := range accounts {
		addresses = append(addresses, Address(a))
	}
	return
}
