package account

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

// GenKey func
func GenKey() *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	return key
}

// Account func
func Account(key *ecdsa.PrivateKey) *bind.TransactOpts {
	return bind.NewKeyedTransactor(key)
}

// GenAccounts func
func GenAccounts(number int) (accounts []*bind.TransactOpts) {
	for i := 0; i < number; i++ {
		accounts = append(accounts, Account(GenKey()))
	}
	return
}

// Address func
func Address(acc *bind.TransactOpts) common.Address {
	return acc.From
}

// Addresses func
func Addresses(accounts []*bind.TransactOpts) (addresses []common.Address) {
	for _, a := range accounts {
		addresses = append(addresses, Address(a))
	}
	return
}
