package account

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func GenKey() *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	return key
}

func Account(key *ecdsa.PrivateKey) *bind.TransactOpts {
	return bind.NewKeyedTransactor(key)
}

func GenAccounts(number int) (accounts []*bind.TransactOpts) {
	for i := 0; i < number; i++ {
		accounts = append(accounts, Account(GenKey()))
	}
	return
}

func Address(acc *bind.TransactOpts) common.Address {
	return acc.From
}

func Addresses(accounts []*bind.TransactOpts) (addresses []common.Address) {
	for _, a := range accounts {
		addresses = append(addresses, Address(a))
	}
	return
}
