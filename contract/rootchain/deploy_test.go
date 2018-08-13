package rootchain

import (
	"testing"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
)

func TestDeploy(t *testing.T) {
	accounts := account.GenAccounts(3)
	owner := accounts[0]
	server := backend.NewSimulatedBackend(account.Addresses(accounts))

	_, _, err := Deploy(owner.TransactOpts, server)
	if err != nil {
		t.Fatal(err)
	}
}
