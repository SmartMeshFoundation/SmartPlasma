// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mediator

import (
	"math/big"
	"strings"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi"
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
)

// MediatorABI is the input ABI used to generate the binding from.
const MediatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"currency\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootChain\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"prevTx\",\"type\":\"bytes\"},{\"name\":\"prevTxProof\",\"type\":\"bytes\"},{\"name\":\"prevTxBlkNum\",\"type\":\"uint256\"},{\"name\":\"txRaw\",\"type\":\"bytes\"},{\"name\":\"txProof\",\"type\":\"bytes\"},{\"name\":\"txBlkNum\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"checkToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"plasma\",\"type\":\"address\"}],\"name\":\"joinPlasma\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// MediatorBin is the compiled bytecode used for deploying new contracts.
const MediatorBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556001805460a060020a60ff0219169055610baa8061004b6000396000f3006060604052600436106100745763ffffffff60e060020a60003504166347e7ef248114610079578063715018a61461009d5780638da5cb5b146100b0578063987ab9db146100df578063e15d9788146100f2578063f1880b2414610214578063f2fde38b14610247578063fb1ff1d714610266575b600080fd5b341561008457600080fd5b61009b600160a060020a0360043516602435610285565b005b34156100a857600080fd5b61009b61040c565b34156100bb57600080fd5b6100c361047e565b604051600160a060020a03909116815260200160405180910390f35b34156100ea57600080fd5b6100c361048d565b34156100fd57600080fd5b61009b60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061049c92505050565b341561021f57600080fd5b610233600160a060020a03600435166107a6565b604051901515815260200160405180910390f35b341561025257600080fd5b61009b600160a060020a0360043516610a35565b341561027157600080fd5b61009b600160a060020a0360043516610a5c565b60008080831161029457600080fd5b83915081600160a060020a03166323b872dd33308660006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561030157600080fd5b6102c65a03f1151561031257600080fd5b50505060405180515050600154600160a060020a0316638340f54933868660006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561038857600080fd5b6102c65a03f1151561039957600080fd5b505050604051805190509050604080519081016040908152600160a060020a0386168252602080830186905260008481526002909152208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617815560208201516001909101555050505050565b60005433600160a060020a0390811691161461042757600080fd5b600054600160a060020a03167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600154600160a060020a031681565b60015460009081908190600160a060020a031663f310f2b2338b8b8b8b8b8b89604051602001526040518863ffffffff1660e060020a0281526004018088600160a060020a0316600160a060020a031681526020018060200180602001878152602001806020018060200186815260200185810385528b818151815260200191508051906020019080838360005b8381101561054257808201518382015260200161052a565b50505050905090810190601f16801561056f5780820380516001836020036101000a031916815260200191505b5085810384528a818151815260200191508051906020019080838360005b838110156105a557808201518382015260200161058d565b50505050905090810190601f1680156105d25780820380516001836020036101000a031916815260200191505b50858103835288818151815260200191508051906020019080838360005b838110156106085780820151838201526020016105f0565b50505050905090810190601f1680156106355780820380516001836020036101000a031916815260200191505b50858103825287818151815260200191508051906020019080838360005b8381101561066b578082015183820152602001610653565b50505050905090810190601f1680156106985780820380516001836020036101000a031916815260200191505b509b505050505050505050505050602060405180830381600087803b15156106bf57600080fd5b6102c65a03f115156106d057600080fd5b505050604051805160008181526002602052604080822080546001820154949850909650600160a060020a0316945084935063a9059cbb9233929091516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561075457600080fd5b6102c65a03f1151561076557600080fd5b5050506040518051505050600091825250600260205260408120805473ffffffffffffffffffffffffffffffffffffffff1916815560010155505050505050565b60008181600160a060020a0382166318160ddd82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156107ef57600080fd5b6102c65a03f1151561080057600080fd5b5050506040518051905011151561081657600080fd5b600081600160a060020a03166370a082313360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561086f57600080fd5b6102c65a03f1151561088057600080fd5b5050506040518051905011151561089657600080fd5b80600160a060020a031663095ea7b3306000806040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156108f357600080fd5b6102c65a03f1151561090457600080fd5b50505060405180519050151561091957600080fd5b80600160a060020a031663d73dd623306000806040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561097657600080fd5b6102c65a03f1151561098757600080fd5b50505060405180519050151561099c57600080fd5b80600160a060020a03166323b872dd33306000806040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610a0657600080fd5b6102c65a03f11515610a1757600080fd5b505050604051805190501515610a2c57600080fd5b50600192915050565b60005433600160a060020a03908116911614610a5057600080fd5b610a5981610afe565b50565b60005433600160a060020a03908116911614610a7757600080fd5b60015474010000000000000000000000000000000000000000900460ff1615610a9f57600080fd5b6001805474ff000000000000000000000000000000000000000019600160a060020a0390931673ffffffffffffffffffffffffffffffffffffffff19909116179190911674010000000000000000000000000000000000000000179055565b600160a060020a0381161515610b1357600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820bdf473e78bd98faedc04f7b25bf6c5a95d67d1c653570665db87155ad13a04be0029`

// DeployMediator deploys a new Ethereum contract, binding an instance of Mediator to it.
func DeployMediator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mediator, error) {
	parsed, err := abi.JSON(strings.NewReader(MediatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MediatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mediator{MediatorCaller: MediatorCaller{contract: contract}, MediatorTransactor: MediatorTransactor{contract: contract}}, nil
}

// Mediator is an auto generated Go binding around an Ethereum contract.
type Mediator struct {
	MediatorCaller     // Read-only binding to the contract
	MediatorTransactor // Write-only binding to the contract
}

// MediatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MediatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MediatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MediatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MediatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MediatorSession struct {
	Contract     *Mediator               // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MediatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MediatorCallerSession struct {
	Contract *MediatorCaller         // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// MediatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MediatorTransactorSession struct {
	Contract     *MediatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MediatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MediatorRaw struct {
	Contract *Mediator // Generic contract binding to access the raw methods on
}

// MediatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MediatorCallerRaw struct {
	Contract *MediatorCaller // Generic read-only contract binding to access the raw methods on
}

// MediatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MediatorTransactorRaw struct {
	Contract *MediatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMediator creates a new instance of Mediator, bound to a specific deployed contract.
func NewMediator(address common.Address, backend bind.ContractBackend) (*Mediator, error) {
	contract, err := bindMediator(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mediator{MediatorCaller: MediatorCaller{contract: contract}, MediatorTransactor: MediatorTransactor{contract: contract}}, nil
}

// NewMediatorCaller creates a new read-only instance of Mediator, bound to a specific deployed contract.
func NewMediatorCaller(address common.Address, caller bind.ContractCaller) (*MediatorCaller, error) {
	contract, err := bindMediator(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MediatorCaller{contract: contract}, nil
}

// NewMediatorTransactor creates a new write-only instance of Mediator, bound to a specific deployed contract.
func NewMediatorTransactor(address common.Address, transactor bind.ContractTransactor) (*MediatorTransactor, error) {
	contract, err := bindMediator(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MediatorTransactor{contract: contract}, nil
}

// bindMediator binds a generic wrapper to an already deployed contract.
func bindMediator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MediatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mediator *MediatorRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Mediator.Contract.MediatorCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mediator *MediatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediator.Contract.MediatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mediator *MediatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mediator.Contract.MediatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mediator *MediatorCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Mediator.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mediator *MediatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mediator *MediatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mediator.Contract.contract.Transact(opts, method, params...)
}

// CheckToken is a free data retrieval call binding the contract method 0xf1880b24.
//
// Solidity: function checkToken(addr address) constant returns(bool)
func (_Mediator *MediatorCaller) CheckToken(opts *bind.CallOptsWithNumber, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mediator.contract.CallWithNumber(opts, out, "checkToken", addr)
	return *ret0, err
}

// CheckToken is a free data retrieval call binding the contract method 0xf1880b24.
//
// Solidity: function checkToken(addr address) constant returns(bool)
func (_Mediator *MediatorSession) CheckToken(addr common.Address) (bool, error) {
	return _Mediator.Contract.CheckToken(&_Mediator.CallOpts, addr)
}

// CheckToken is a free data retrieval call binding the contract method 0xf1880b24.
//
// Solidity: function checkToken(addr address) constant returns(bool)
func (_Mediator *MediatorCallerSession) CheckToken(addr common.Address) (bool, error) {
	return _Mediator.Contract.CheckToken(&_Mediator.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Mediator *MediatorCaller) Owner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mediator.contract.CallWithNumber(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Mediator *MediatorSession) Owner() (common.Address, error) {
	return _Mediator.Contract.Owner(&_Mediator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Mediator *MediatorCallerSession) Owner() (common.Address, error) {
	return _Mediator.Contract.Owner(&_Mediator.CallOpts)
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() constant returns(address)
func (_Mediator *MediatorCaller) RootChain(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mediator.contract.CallWithNumber(opts, out, "rootChain")
	return *ret0, err
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() constant returns(address)
func (_Mediator *MediatorSession) RootChain() (common.Address, error) {
	return _Mediator.Contract.RootChain(&_Mediator.CallOpts)
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() constant returns(address)
func (_Mediator *MediatorCallerSession) RootChain() (common.Address, error) {
	return _Mediator.Contract.RootChain(&_Mediator.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(currency address, amount uint256) returns()
func (_Mediator *MediatorTransactor) Deposit(opts *bind.TransactOpts, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mediator.contract.Transact(opts, "deposit", currency, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(currency address, amount uint256) returns()
func (_Mediator *MediatorSession) Deposit(currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mediator.Contract.Deposit(&_Mediator.TransactOpts, currency, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(currency address, amount uint256) returns()
func (_Mediator *MediatorTransactorSession) Deposit(currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mediator.Contract.Deposit(&_Mediator.TransactOpts, currency, amount)
}

// JoinPlasma is a paid mutator transaction binding the contract method 0xfb1ff1d7.
//
// Solidity: function joinPlasma(plasma address) returns()
func (_Mediator *MediatorTransactor) JoinPlasma(opts *bind.TransactOpts, plasma common.Address) (*types.Transaction, error) {
	return _Mediator.contract.Transact(opts, "joinPlasma", plasma)
}

// JoinPlasma is a paid mutator transaction binding the contract method 0xfb1ff1d7.
//
// Solidity: function joinPlasma(plasma address) returns()
func (_Mediator *MediatorSession) JoinPlasma(plasma common.Address) (*types.Transaction, error) {
	return _Mediator.Contract.JoinPlasma(&_Mediator.TransactOpts, plasma)
}

// JoinPlasma is a paid mutator transaction binding the contract method 0xfb1ff1d7.
//
// Solidity: function joinPlasma(plasma address) returns()
func (_Mediator *MediatorTransactorSession) JoinPlasma(plasma common.Address) (*types.Transaction, error) {
	return _Mediator.Contract.JoinPlasma(&_Mediator.TransactOpts, plasma)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mediator *MediatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mediator *MediatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mediator.Contract.RenounceOwnership(&_Mediator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mediator *MediatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mediator.Contract.RenounceOwnership(&_Mediator.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Mediator *MediatorTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Mediator.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Mediator *MediatorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Mediator.Contract.TransferOwnership(&_Mediator.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Mediator *MediatorTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Mediator.Contract.TransferOwnership(&_Mediator.TransactOpts, _newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xe15d9788.
//
// Solidity: function withdraw(prevTx bytes, prevTxProof bytes, prevTxBlkNum uint256, txRaw bytes, txProof bytes, txBlkNum uint256) returns()
func (_Mediator *MediatorTransactor) Withdraw(opts *bind.TransactOpts, prevTx []byte, prevTxProof []byte, prevTxBlkNum *big.Int, txRaw []byte, txProof []byte, txBlkNum *big.Int) (*types.Transaction, error) {
	return _Mediator.contract.Transact(opts, "withdraw", prevTx, prevTxProof, prevTxBlkNum, txRaw, txProof, txBlkNum)
}

// Withdraw is a paid mutator transaction binding the contract method 0xe15d9788.
//
// Solidity: function withdraw(prevTx bytes, prevTxProof bytes, prevTxBlkNum uint256, txRaw bytes, txProof bytes, txBlkNum uint256) returns()
func (_Mediator *MediatorSession) Withdraw(prevTx []byte, prevTxProof []byte, prevTxBlkNum *big.Int, txRaw []byte, txProof []byte, txBlkNum *big.Int) (*types.Transaction, error) {
	return _Mediator.Contract.Withdraw(&_Mediator.TransactOpts, prevTx, prevTxProof, prevTxBlkNum, txRaw, txProof, txBlkNum)
}

// Withdraw is a paid mutator transaction binding the contract method 0xe15d9788.
//
// Solidity: function withdraw(prevTx bytes, prevTxProof bytes, prevTxBlkNum uint256, txRaw bytes, txProof bytes, txBlkNum uint256) returns()
func (_Mediator *MediatorTransactorSession) Withdraw(prevTx []byte, prevTxProof []byte, prevTxBlkNum *big.Int, txRaw []byte, txProof []byte, txBlkNum *big.Int) (*types.Transaction, error) {
	return _Mediator.Contract.Withdraw(&_Mediator.TransactOpts, prevTx, prevTxProof, prevTxBlkNum, txRaw, txProof, txBlkNum)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556102128061003b6000396000f3006060604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a6811461005b5780638da5cb5b14610070578063f2fde38b1461009f575b600080fd5b341561006657600080fd5b61006e6100be565b005b341561007b57600080fd5b610083610130565b604051600160a060020a03909116815260200160405180910390f35b34156100aa57600080fd5b61006e600160a060020a036004351661013f565b60005433600160a060020a039081169116146100d957600080fd5b600054600160a060020a03167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b60005433600160a060020a0390811691161461015a57600080fd5b61016381610166565b50565b600160a060020a038116151561017b57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820068557bc6585ac354383392f378fcbe2984e3cde70b7e1863d4a8eb15c414f000029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable                // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller          // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.CallWithNumber(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"currency\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"previousTx\",\"type\":\"bytes\"},{\"name\":\"previousTxProof\",\"type\":\"bytes\"},{\"name\":\"previousTxBlockNum\",\"type\":\"uint256\"},{\"name\":\"lastTx\",\"type\":\"bytes\"},{\"name\":\"lastTxProof\",\"type\":\"bytes\"},{\"name\":\"lastTxBlockNum\",\"type\":\"uint256\"}],\"name\":\"finishExit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `0x`

// DeployRootChain deploys a new Ethereum contract, binding an instance of RootChain to it.
func DeployRootChain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RootChain, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}}, nil
}

// RootChain is an auto generated Go binding around an Ethereum contract.
type RootChain struct {
	RootChainCaller     // Read-only binding to the contract
	RootChainTransactor // Write-only binding to the contract
}

// RootChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainSession struct {
	Contract     *RootChain              // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RootChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainCallerSession struct {
	Contract *RootChainCaller        // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// RootChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainTransactorSession struct {
	Contract     *RootChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RootChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainRaw struct {
	Contract *RootChain // Generic contract binding to access the raw methods on
}

// RootChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainCallerRaw struct {
	Contract *RootChainCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainTransactorRaw struct {
	Contract *RootChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChain creates a new instance of RootChain, bound to a specific deployed contract.
func NewRootChain(address common.Address, backend bind.ContractBackend) (*RootChain, error) {
	contract, err := bindRootChain(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}}, nil
}

// NewRootChainCaller creates a new read-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainCaller(address common.Address, caller bind.ContractCaller) (*RootChainCaller, error) {
	contract, err := bindRootChain(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainCaller{contract: contract}, nil
}

// NewRootChainTransactor creates a new write-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainTransactor, error) {
	contract, err := bindRootChain(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RootChainTransactor{contract: contract}, nil
}

// bindRootChain binds a generic wrapper to an already deployed contract.
func bindRootChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.RootChainCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(account address, currency address, amount uint256) returns(bytes32)
func (_RootChain *RootChainTransactor) Deposit(opts *bind.TransactOpts, account common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "deposit", account, currency, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(account address, currency address, amount uint256) returns(bytes32)
func (_RootChain *RootChainSession) Deposit(account common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.Deposit(&_RootChain.TransactOpts, account, currency, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(account address, currency address, amount uint256) returns(bytes32)
func (_RootChain *RootChainTransactorSession) Deposit(account common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.Deposit(&_RootChain.TransactOpts, account, currency, amount)
}

// FinishExit is a paid mutator transaction binding the contract method 0xf310f2b2.
//
// Solidity: function finishExit(account address, previousTx bytes, previousTxProof bytes, previousTxBlockNum uint256, lastTx bytes, lastTxProof bytes, lastTxBlockNum uint256) returns(bytes32)
func (_RootChain *RootChainTransactor) FinishExit(opts *bind.TransactOpts, account common.Address, previousTx []byte, previousTxProof []byte, previousTxBlockNum *big.Int, lastTx []byte, lastTxProof []byte, lastTxBlockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finishExit", account, previousTx, previousTxProof, previousTxBlockNum, lastTx, lastTxProof, lastTxBlockNum)
}

// FinishExit is a paid mutator transaction binding the contract method 0xf310f2b2.
//
// Solidity: function finishExit(account address, previousTx bytes, previousTxProof bytes, previousTxBlockNum uint256, lastTx bytes, lastTxProof bytes, lastTxBlockNum uint256) returns(bytes32)
func (_RootChain *RootChainSession) FinishExit(account common.Address, previousTx []byte, previousTxProof []byte, previousTxBlockNum *big.Int, lastTx []byte, lastTxProof []byte, lastTxBlockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.FinishExit(&_RootChain.TransactOpts, account, previousTx, previousTxProof, previousTxBlockNum, lastTx, lastTxProof, lastTxBlockNum)
}

// FinishExit is a paid mutator transaction binding the contract method 0xf310f2b2.
//
// Solidity: function finishExit(account address, previousTx bytes, previousTxProof bytes, previousTxBlockNum uint256, lastTx bytes, lastTxProof bytes, lastTxBlockNum uint256) returns(bytes32)
func (_RootChain *RootChainTransactorSession) FinishExit(account common.Address, previousTx []byte, previousTxProof []byte, previousTxBlockNum *big.Int, lastTx []byte, lastTxProof []byte, lastTxBlockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.FinishExit(&_RootChain.TransactOpts, account, previousTx, previousTxProof, previousTxBlockNum, lastTx, lastTxProof, lastTxBlockNum)
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token                  // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller            // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOptsWithNumber, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.CallWithNumber(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_Token *TokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.CallWithNumber(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_Token *TokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_Token *TokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseApproval(&_Token.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_Token *TokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseApproval(&_Token.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_Token *TokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_Token *TokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_Token *TokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, value)
}
