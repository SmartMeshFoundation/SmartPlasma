//go:generate abigen --sol Mediator.sol --pkg mediator  --out ../contract/mediator/mediator.go
//go:generate abigen --sol ExampleToken.sol --pkg erc20token  --out ../contract/erc20token/erc20token.go
//go:generate abigen --sol RootChain.sol --pkg rootchain  --out ../contract/rootchain/rootchain.go

package solidity
