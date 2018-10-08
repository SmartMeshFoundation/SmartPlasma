// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rootchain

import (
	"math/big"
	"strings"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi"
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
)

// ECRecoveryABI is the input ABI used to generate the binding from.
const ECRecoveryABI = "[]"

// ECRecoveryBin is the compiled bytecode used for deploying new contracts.
const ECRecoveryBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820bff236bc8fdbcea7198032adbd682222b2e06baa27ce9331c353643cfff62cd90029`

// DeployECRecovery deploys a new Ethereum contract, binding an instance of ECRecovery to it.
func DeployECRecovery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECRecovery, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECRecoveryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}}, nil
}

// ECRecovery is an auto generated Go binding around an Ethereum contract.
type ECRecovery struct {
	ECRecoveryCaller     // Read-only binding to the contract
	ECRecoveryTransactor // Write-only binding to the contract
}

// ECRecoveryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECRecoveryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECRecoveryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoverySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECRecoverySession struct {
	Contract     *ECRecovery             // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ECRecoveryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECRecoveryCallerSession struct {
	Contract *ECRecoveryCaller       // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// ECRecoveryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECRecoveryTransactorSession struct {
	Contract     *ECRecoveryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ECRecoveryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECRecoveryRaw struct {
	Contract *ECRecovery // Generic contract binding to access the raw methods on
}

// ECRecoveryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECRecoveryCallerRaw struct {
	Contract *ECRecoveryCaller // Generic read-only contract binding to access the raw methods on
}

// ECRecoveryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECRecoveryTransactorRaw struct {
	Contract *ECRecoveryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECRecovery creates a new instance of ECRecovery, bound to a specific deployed contract.
func NewECRecovery(address common.Address, backend bind.ContractBackend) (*ECRecovery, error) {
	contract, err := bindECRecovery(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}}, nil
}

// NewECRecoveryCaller creates a new read-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryCaller(address common.Address, caller bind.ContractCaller) (*ECRecoveryCaller, error) {
	contract, err := bindECRecovery(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryCaller{contract: contract}, nil
}

// NewECRecoveryTransactor creates a new write-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryTransactor(address common.Address, transactor bind.ContractTransactor) (*ECRecoveryTransactor, error) {
	contract, err := bindECRecovery(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryTransactor{contract: contract}, nil
}

// bindECRecovery binds a generic wrapper to an already deployed contract.
func bindECRecovery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.ECRecoveryCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transact(opts, method, params...)
}

// MerkleABI is the input ABI used to generate the binding from.
const MerkleABI = "[]"

// MerkleBin is the compiled bytecode used for deploying new contracts.
const MerkleBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a7230582071fa3ce60d21f638c169fbbff448b06deeecbc899902b809d8422c2dea1723fb0029`

// DeployMerkle deploys a new Ethereum contract, binding an instance of Merkle to it.
func DeployMerkle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Merkle, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Merkle{MerkleCaller: MerkleCaller{contract: contract}, MerkleTransactor: MerkleTransactor{contract: contract}}, nil
}

// Merkle is an auto generated Go binding around an Ethereum contract.
type Merkle struct {
	MerkleCaller     // Read-only binding to the contract
	MerkleTransactor // Write-only binding to the contract
}

// MerkleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleSession struct {
	Contract     *Merkle                 // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MerkleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleCallerSession struct {
	Contract *MerkleCaller           // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// MerkleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTransactorSession struct {
	Contract     *MerkleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleRaw struct {
	Contract *Merkle // Generic contract binding to access the raw methods on
}

// MerkleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleCallerRaw struct {
	Contract *MerkleCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTransactorRaw struct {
	Contract *MerkleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkle creates a new instance of Merkle, bound to a specific deployed contract.
func NewMerkle(address common.Address, backend bind.ContractBackend) (*Merkle, error) {
	contract, err := bindMerkle(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Merkle{MerkleCaller: MerkleCaller{contract: contract}, MerkleTransactor: MerkleTransactor{contract: contract}}, nil
}

// NewMerkleCaller creates a new read-only instance of Merkle, bound to a specific deployed contract.
func NewMerkleCaller(address common.Address, caller bind.ContractCaller) (*MerkleCaller, error) {
	contract, err := bindMerkle(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleCaller{contract: contract}, nil
}

// NewMerkleTransactor creates a new write-only instance of Merkle, bound to a specific deployed contract.
func NewMerkleTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTransactor, error) {
	contract, err := bindMerkle(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MerkleTransactor{contract: contract}, nil
}

// bindMerkle binds a generic wrapper to an already deployed contract.
func bindMerkle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merkle *MerkleRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Merkle.Contract.MerkleCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merkle *MerkleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkle.Contract.MerkleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merkle *MerkleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merkle.Contract.MerkleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merkle *MerkleCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Merkle.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merkle *MerkleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merkle *MerkleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merkle.Contract.contract.Transact(opts, method, params...)
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

// PlasmaLibABI is the input ABI used to generate the binding from.
const PlasmaLibABI = "[]"

// PlasmaLibBin is the compiled bytecode used for deploying new contracts.
const PlasmaLibBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820b87b35813ecf436cf08813b22af77119bbdbed52c6a30b46ebf581a5892f2fbc0029`

// DeployPlasmaLib deploys a new Ethereum contract, binding an instance of PlasmaLib to it.
func DeployPlasmaLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PlasmaLib, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PlasmaLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PlasmaLib{PlasmaLibCaller: PlasmaLibCaller{contract: contract}, PlasmaLibTransactor: PlasmaLibTransactor{contract: contract}}, nil
}

// PlasmaLib is an auto generated Go binding around an Ethereum contract.
type PlasmaLib struct {
	PlasmaLibCaller     // Read-only binding to the contract
	PlasmaLibTransactor // Write-only binding to the contract
}

// PlasmaLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlasmaLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlasmaLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlasmaLibSession struct {
	Contract     *PlasmaLib              // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PlasmaLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlasmaLibCallerSession struct {
	Contract *PlasmaLibCaller        // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// PlasmaLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlasmaLibTransactorSession struct {
	Contract     *PlasmaLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PlasmaLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlasmaLibRaw struct {
	Contract *PlasmaLib // Generic contract binding to access the raw methods on
}

// PlasmaLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlasmaLibCallerRaw struct {
	Contract *PlasmaLibCaller // Generic read-only contract binding to access the raw methods on
}

// PlasmaLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlasmaLibTransactorRaw struct {
	Contract *PlasmaLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlasmaLib creates a new instance of PlasmaLib, bound to a specific deployed contract.
func NewPlasmaLib(address common.Address, backend bind.ContractBackend) (*PlasmaLib, error) {
	contract, err := bindPlasmaLib(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlasmaLib{PlasmaLibCaller: PlasmaLibCaller{contract: contract}, PlasmaLibTransactor: PlasmaLibTransactor{contract: contract}}, nil
}

// NewPlasmaLibCaller creates a new read-only instance of PlasmaLib, bound to a specific deployed contract.
func NewPlasmaLibCaller(address common.Address, caller bind.ContractCaller) (*PlasmaLibCaller, error) {
	contract, err := bindPlasmaLib(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaLibCaller{contract: contract}, nil
}

// NewPlasmaLibTransactor creates a new write-only instance of PlasmaLib, bound to a specific deployed contract.
func NewPlasmaLibTransactor(address common.Address, transactor bind.ContractTransactor) (*PlasmaLibTransactor, error) {
	contract, err := bindPlasmaLib(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PlasmaLibTransactor{contract: contract}, nil
}

// bindPlasmaLib binds a generic wrapper to an already deployed contract.
func bindPlasmaLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaLib *PlasmaLibRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _PlasmaLib.Contract.PlasmaLibCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaLib *PlasmaLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaLib.Contract.PlasmaLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaLib *PlasmaLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaLib.Contract.PlasmaLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaLib *PlasmaLibCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _PlasmaLib.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaLib *PlasmaLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaLib *PlasmaLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaLib.Contract.contract.Transact(opts, method, params...)
}

// RLPABI is the input ABI used to generate the binding from.
const RLPABI = "[]"

// RLPBin is the compiled bytecode used for deploying new contracts.
const RLPBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a7230582052e729c6f562ff7bb33400c60a408484f6b7b551850bb75051f420f918471d990029`

// DeployRLP deploys a new Ethereum contract, binding an instance of RLP to it.
func DeployRLP(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLP, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLP{RLPCaller: RLPCaller{contract: contract}, RLPTransactor: RLPTransactor{contract: contract}}, nil
}

// RLP is an auto generated Go binding around an Ethereum contract.
type RLP struct {
	RLPCaller     // Read-only binding to the contract
	RLPTransactor // Write-only binding to the contract
}

// RLPCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPSession struct {
	Contract     *RLP                    // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RLPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPCallerSession struct {
	Contract *RLPCaller              // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// RLPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPTransactorSession struct {
	Contract     *RLPTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPRaw struct {
	Contract *RLP // Generic contract binding to access the raw methods on
}

// RLPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPCallerRaw struct {
	Contract *RLPCaller // Generic read-only contract binding to access the raw methods on
}

// RLPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPTransactorRaw struct {
	Contract *RLPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLP creates a new instance of RLP, bound to a specific deployed contract.
func NewRLP(address common.Address, backend bind.ContractBackend) (*RLP, error) {
	contract, err := bindRLP(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLP{RLPCaller: RLPCaller{contract: contract}, RLPTransactor: RLPTransactor{contract: contract}}, nil
}

// NewRLPCaller creates a new read-only instance of RLP, bound to a specific deployed contract.
func NewRLPCaller(address common.Address, caller bind.ContractCaller) (*RLPCaller, error) {
	contract, err := bindRLP(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RLPCaller{contract: contract}, nil
}

// NewRLPTransactor creates a new write-only instance of RLP, bound to a specific deployed contract.
func NewRLPTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPTransactor, error) {
	contract, err := bindRLP(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RLPTransactor{contract: contract}, nil
}

// bindRLP binds a generic wrapper to an already deployed contract.
func bindRLP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLP *RLPRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _RLP.Contract.RLPCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLP *RLPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLP.Contract.RLPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLP *RLPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLP.Contract.RLPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLP *RLPCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _RLP.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLP *RLPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLP *RLPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLP.Contract.contract.Transact(opts, method, params...)
}

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"previousTx\",\"type\":\"bytes\"},{\"name\":\"previousTxProof\",\"type\":\"bytes\"},{\"name\":\"previousTxBlockNum\",\"type\":\"uint256\"},{\"name\":\"lastTx\",\"type\":\"bytes\"},{\"name\":\"lastTxProof\",\"type\":\"bytes\"},{\"name\":\"lastTxBlockNum\",\"type\":\"uint256\"}],\"name\":\"startExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wallet2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChallenge\",\"outputs\":[{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"challengeBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exits\",\"outputs\":[{\"name\":\"state\",\"type\":\"uint256\"},{\"name\":\"exitTime\",\"type\":\"uint256\"},{\"name\":\"exitTxBlkNum\",\"type\":\"uint256\"},{\"name\":\"exitTx\",\"type\":\"bytes\"},{\"name\":\"txBeforeExitTxBlkNum\",\"type\":\"uint256\"},{\"name\":\"txBeforeExitTx\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpoint\",\"type\":\"bytes32\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"}],\"name\":\"checkpointIsChallenge\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"challengeBlockNum\",\"type\":\"uint256\"}],\"name\":\"challengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpointRoot\",\"type\":\"bytes32\"},{\"name\":\"checkpointProof\",\"type\":\"bytes\"},{\"name\":\"wrongNonce\",\"type\":\"uint256\"},{\"name\":\"lastTx\",\"type\":\"bytes\"},{\"name\":\"lastTxProof\",\"type\":\"bytes\"},{\"name\":\"lastTxBlockNum\",\"type\":\"uint256\"}],\"name\":\"challengeCheckpoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpointRoot\",\"type\":\"bytes32\"},{\"name\":\"checkpointProof\",\"type\":\"bytes\"},{\"name\":\"historicalCheckpointRoot\",\"type\":\"bytes32\"},{\"name\":\"historicalCheckpointProof\",\"type\":\"bytes\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"moreNonce\",\"type\":\"uint256\"}],\"name\":\"respondWithHistoricalCheckpoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"}],\"name\":\"challengeExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"currency\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"respondTx\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"respondChallengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"newBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpoint\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getCheckpointChallenge\",\"outputs\":[{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"challengeBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpoint\",\"type\":\"bytes32\"}],\"name\":\"checkpointChallengesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"checkpointRoot\",\"type\":\"bytes32\"},{\"name\":\"checkpointProof\",\"type\":\"bytes\"},{\"name\":\"checkpointNonce\",\"type\":\"bytes32\"}],\"name\":\"respondChallengeExitWithCheckpoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"newCheckpoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpointRoot\",\"type\":\"bytes32\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"respondTx\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"respondCheckpointChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"challengesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"previousTx\",\"type\":\"bytes\"},{\"name\":\"previousTxProof\",\"type\":\"bytes\"},{\"name\":\"previousTxBlockNum\",\"type\":\"uint256\"},{\"name\":\"lastTx\",\"type\":\"bytes\"},{\"name\":\"lastTxProof\",\"type\":\"bytes\"},{\"name\":\"lastTxBlockNum\",\"type\":\"uint256\"}],\"name\":\"finishExit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"NewBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"NewCheckpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastBlock\",\"type\":\"uint256\"}],\"name\":\"StartExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"FinishExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"ChallengeExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"checkpoint\",\"type\":\"bytes32\"}],\"name\":\"ChallengeCheckpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"RespondChallengeExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"checkpoint\",\"type\":\"bytes32\"}],\"name\":\"RespondCheckpointChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"checkpoint\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"historicalCheckpoint\",\"type\":\"bytes32\"}],\"name\":\"RespondWithHistoricalCheckpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `0x6060604052341561000f57600080fd5b604051602080613acc8339810160405280805160008054600160a060020a03338116600160a060020a0319928316178355600283905562127500600355600192909255600480549290931691161790555050613a5c806100706000396000f3006060604052600436106101745763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166304675c65811461017957806326f3d0931461029d5780632dfdf0b5146102c55780632e8f396d146102d8578063342de17914610370578063370beab9146104a8578063449b2f44146105185780634a2ee1981461052e578063570ca735146105c857806357e871e7146105f75780635e8cdf441461060a5780636d2975c7146106f55780636e606251146107dc578063715018a6146108325780638340f549146108455780638b5931831461086d5780638da5cb5b14610949578063981adca51461095c5780639af35551146109725780639ced68901461098e5780639f26026e146109a7578063a5e450f414610a46578063ad131e1314610a5c578063b15e42f814610b3e578063eb5e91ff14610b54578063f2fde38b14610b6a578063f310f2b214610b89578063f3f480d914610cb9578063f95643b114610ccc575b600080fd5b341561018457600080fd5b61029b60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505093359350610ce292505050565b005b34156102a857600080fd5b6102b3600435610f9d565b60405190815260200160405180910390f35b34156102d057600080fd5b6102b3610faf565b34156102e357600080fd5b6102f1600435602435610fb5565b6040518080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561033457808201518382015260200161031c565b50505050905090810190601f1680156103615780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b341561037b57600080fd5b6103866004356110a7565b60405186815260208101869052604081018590526080810183905260c0606082018181528554600260001961010060018416150201909116049183018290529060a083019060e08401908790801561041f5780601f106103f45761010080835404028352916020019161041f565b820191906000526020600020905b81548152906001019060200180831161040257829003601f168201915b50508381038252845460026000196101006001841615020190911604808252602090910190859080156104935780601f1061046857610100808354040283529160200191610493565b820191906000526020600020905b81548152906001019060200180831161047657829003601f168201915b50509850505050505050505060405180910390f35b34156104b357600080fd5b610504600480359060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506110d895505050505050565b604051901515815260200160405180910390f35b341561052357600080fd5b6102b36004356111a4565b341561053957600080fd5b61029b600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965050933593506111b692505050565b34156105d357600080fd5b6105db61151f565b604051600160a060020a03909116815260200160405180910390f35b341561060257600080fd5b6102b361152e565b341561061557600080fd5b61029b600480359060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061153492505050565b341561070057600080fd5b61029b600480359060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094968635969095506040808201955060209182013587018083019550359350839250601f830182900482029091019051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061165e92505050565b34156107e757600080fd5b610504600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061176495505050505050565b341561083d57600080fd5b61029b61181f565b341561085057600080fd5b6102b3600160a060020a0360043581169060243516604435611891565b341561087857600080fd5b61029b600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061195792505050565b341561095457600080fd5b6105db611ae1565b341561096757600080fd5b61029b600435611af0565b341561097d57600080fd5b6102f1600435602435604435611b6d565b341561099957600080fd5b6102b3600435602435611c6b565b34156109b257600080fd5b61029b600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094968635969095506040808201955060209182013587018083019550359350839250601f83018290048202909101905190810160405281815292919060208401838380828437509496505093359350611cae92505050565b3415610a5157600080fd5b61029b600435611db3565b3415610a6757600080fd5b61029b600480359060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505093359350611e3092505050565b3415610b4957600080fd5b6102b3600435611f69565b3415610b5f57600080fd5b6102b3600435611fa2565b3415610b7557600080fd5b61029b600160a060020a0360043516611fb4565b3415610b9457600080fd5b6102b360048035600160a060020a03169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505093359350611fdb92505050565b3415610cc457600080fd5b6102b3612220565b3415610cd757600080fd5b6102b3600435612226565b610cea61383b565b610cf261383b565b600080600080610d018c612238565b9550610d0c89612238565b945084518a14610d1b57600080fd5b8460200151866020015114610d2f57600080fd5b8460400151866040015114610d4357600080fd5b8460a00151600160a060020a03168660600151600160a060020a031614610d6957600080fd5b610d7f600187608001519063ffffffff61232916565b856080015114610d8e57600080fd5b8460600151600160a060020a031633600160a060020a0316141515610db257600080fd5b60076000866020015181526020810191909152604001600020541515610dd757600080fd5b8560c0015160008b815260056020526040902054909450925060c0850151915060056000888152602001908152602001600020549050610e2486602001518590858e63ffffffff61233c16565b1515610e2f57600080fd5b610e4685602001518390838b63ffffffff61233c16565b1515610e5157600080fd5b600660008660200151815260208101919091526040016000205415610e7557600080fd5b610e828560200151611f69565b15610e8c57600080fd5b60c06040519081016040528060028152602001610eb46003544261232990919063ffffffff16565b81526020018881526020018a81526020018b81526020018d8152506006600087602001518152602001908152602001600020600082015181556020820151816001015560408201518160020155606082015181600301908051610f1b929160200190613877565b506080820151816004015560a082015181600501908051610f40929160200190613877565b509050507f4d3db44958203fdb34cb08c6f6ecda2c9e182cd426b0871efec5ade7ac94580d86602001518b8960405180848152602001838152602001828152602001935050505060405180910390a1505050505050505050505050565b60086020526000908152604090205481565b60015481565b610fbd6138f5565b600083815260096020526040812081906001908101908290610fe690879063ffffffff61232916565b81526020019081526020016000209050806001018160020154818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110945780601f1061106957610100808354040283529160200191611094565b820191906000526020600020905b81548152906001019060200180831161107757829003601f168201915b5050505050915092509250509250929050565b6006602052600090815260409020805460018201546002830154600484015492939192909160038101919060050186565b6000838152600b6020908152604080832085845290915280822082916002909101908490518082805190602001908083835b602083106111295780518252601f19909201916020918201910161110a565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390205490508060001415611171576000915061119c565b6000858152600b60209081526040808320878452825280832084845260010190915290205460ff1691505b509392505050565b60076020526000908152604090205481565b6111be61383b565b6111c661383b565b6111ce61383b565b60008781526006602052604081205481906002146111eb57600080fd5b6112a2600660008b81526020019081526020016000206003018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112985780601f1061126d57610100808354040283529160200191611298565b820191906000526020600020905b81548152906001019060200180831161127b57829003601f168201915b5050505050612238565b9450611326600660008b81526020019081526020016000206005018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112985780601f1061126d57610100808354040283529160200191611298565b935061133188612238565b9250826020015185602001511461134757600080fd5b826040015185604001511461135b57600080fd5b8260c001516000878152600560205260409020549092509050611386828a838a63ffffffff61233c16565b151561139157600080fd5b8260a00151600160a060020a03168560600151600160a060020a03161480156113c1575082608001518560800151105b156114125760008981526006602052604081208181556001810182905560028101829055906113f36003830182613907565b600482016000905560058201600061140b9190613907565b5050611514565b6000898152600660205260409020600201548610801561145d57508260a00151600160a060020a03168460600151600160a060020a031614801561145d575083608001518360800151115b1561148f5760008981526006602052604081208181556001810182905560028101829055906113f36003830182613907565b6000898152600660205260409020600401548610156114c5576000898152600660205260409020600190556114c58989886123b3565b6000898152600660205260409020546001146114e057600080fd5b7fd3103af5a62a3d64dc3b30e901e2aa37bd30bde8b6f61a70e79b23e20d3bbd468960405190815260200160405180910390a15b505050505050505050565b600454600160a060020a031681565b60025481565b61153c61383b565b6000878152600a6020526040812054819081901580159061157e575060035460008b8152600a6020526040902054429161157c919063ffffffff61232916565b115b151561158957600080fd5b6115948b8b896110d8565b1561159e57600080fd5b6115a787612238565b93508360c0015160008681526005602052604090205490935091508790506115d7838c848963ffffffff61233c16565b15156115e257600080fd5b6115f4818c8c8c63ffffffff61233c16565b15156115ff57600080fd5b8360800151881115611617576116178b8b8988612591565b7f5785fb6481777a31c76595b8b701cf63e247203b0e63c12dd7da5d731d3769618b8b60405191825260208201526040908101905180910390a15050505050505050505050565b61166661383b565b60006116738989866110d8565b151561167e57600080fd5b61168784612238565b91508290506080820151831161169c57600080fd5b6116ae818a888863ffffffff61233c16565b15156116b957600080fd5b6003546000878152600a602052604090205442916116dd919063ffffffff61232916565b106116e757600080fd5b6000888152600a6020526040808220548883529120541061170757600080fd5b6117128989866127a1565b7f8e7f7ccff34f7d0789427738f22785ddd8a9fee340613ffc9b8b24ad0e2a20f189898860405192835260208301919091526040808301919091526060909101905180910390a1505050505050505050565b60008060096000858152602001908152602001600020600201836040518082805190602001908083835b602083106117ad5780518252601f19909201916020918201910161178e565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054905080600014156117f55760009150611818565b600084815260096020908152604080832084845260010190915290205460ff1691505b5092915050565b60005433600160a060020a0390811691161461183a57600080fd5b600054600160a060020a03167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b60008054819033600160a060020a039081169116146118af57600080fd5b6118bc8585600154612acf565b6000818152600760209081526040808320879055600254600890925290912055600180549192506118f3919063ffffffff61232916565b6001557f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a158584836040518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390a1949350505050565b61195f61383b565b61196761383b565b6000806119748989611764565b151561197f57600080fd5b60008981526006602052604090205460011461199a57600080fd5b6119a388612238565b93506119ae87612238565b925082602001518460200151146119c457600080fd5b82604001518460400151146119d857600080fd5b8260a00151600160a060020a03168460600151600160a060020a0316146119fe57600080fd5b8260800151611a19600186608001519063ffffffff61232916565b14611a2357600080fd5b600089815260066020526040902060040154851115611a4157600080fd5b8260c001516000868152600560205260409020549092509050611a6c828a838963ffffffff61233c16565b1515611a7757600080fd5b611a818989612b19565b611a8a89611f69565b1515611aa3576000898152600660205260409020600290555b7f755b2676ab5f5b54bffac288782b3b18ae132ffd1950e416c1cfeb98eeb3c5c28960405190815260200160405180910390a1505050505050505050565b600054600160a060020a031681565b60045433600160a060020a03908116911614611b0b57600080fd5b600254611b1f90600163ffffffff61232916565b6002819055600090815260056020526040908190208290557f5f11b60a71ba7b4124fe41971a682a44d1af8fff92e0c4852a2701e56323218a9082905190815260200160405180910390a150565b611b756138f5565b6000848152600b60209081526040808320868452909152812081906001908101908290611ba990879063ffffffff61232916565b81526020019081526020016000209050806001018160020154818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611c575780601f10611c2c57610100808354040283529160200191611c57565b820191906000526020600020905b815481529060010190602001808311611c3a57829003601f168201915b505050505091509250925050935093915050565b6000828152600b60209081526040808320848452909152812054801515611c955760009150611818565b611ca681600163ffffffff612def16565b949350505050565b611cb661383b565b611cbf85612238565b9050611ccb8686611764565b1515611cd657600080fd5b600086815260066020526040902054600114611cf157600080fd5b6003546000858152600a60205260409020544291611d15919063ffffffff61232916565b10611d1f57600080fd5b80608001518211611d2f57600080fd5b611d418287868663ffffffff61233c16565b1515611d4c57600080fd5b611d568686612b19565b611d5f86611f69565b1515611d78576000868152600660205260409020600290555b7f755b2676ab5f5b54bffac288782b3b18ae132ffd1950e416c1cfeb98eeb3c5c28660405190815260200160405180910390a1505050505050565b60045433600160a060020a03908116911614611dce57600080fd5b6000818152600a602052604090205415611de757600080fd5b6000818152600a6020526040908190204290557f3dfae83a0b2f3013f409fd97c7e72574fcb10cd81987893771d8a2707d533d219082905190815260200160405180910390a150565b611e3861383b565b611e4061383b565b600080611e4e8a8a8a6110d8565b1515611e5957600080fd5b611e6288612238565b9350611e6d87612238565b92508260200151846020015114611e8357600080fd5b8260400151846040015114611e9757600080fd5b8260a00151600160a060020a03168460600151600160a060020a031614611ebd57600080fd5b8260800151611ed8600186608001519063ffffffff61232916565b14611ee257600080fd5b8260c001516000868152600560205260409020549092509050611f0d828b838963ffffffff61233c16565b1515611f1857600080fd5b611f238a8a8a6127a1565b7f12e9fe9a5ae32610f4f992917e433e03162fb143bc4614d245ff7f30482866b08a8a60405191825260208201526040908101905180910390a150505050505050505050565b600081815260096020526040812054801515611f885760009150611f9c565b611f9981600163ffffffff612def16565b91505b50919050565b600a6020526000908152604090205481565b60005433600160a060020a03908116911614611fcf57600080fd5b611fd881612e01565b50565b6000611fe561383b565b611fed61383b565b6000805481908190819033600160a060020a0390811691161461200f57600080fd5b6120188d612238565b95506120238a612238565b945084518b1461203257600080fd5b846020015186602001511461204657600080fd5b846040015186604001511461205a57600080fd5b8460a00151600160a060020a03168660600151600160a060020a03161461208057600080fd5b8460600151600160a060020a038f811691161461209c57600080fd5b8560c0015160008c815260056020526040902054909450925060c08501519150600560008981526020019081526020016000205490506120e986602001518590858f63ffffffff61233c16565b15156120f457600080fd5b61210b85602001518390838c63ffffffff61233c16565b151561211657600080fd5b4260066000876020015181526020019081526020016000206001015410151561213e57600080fd5b600660008660200151815260208101919091526040016000205460021461216457600080fd5b6121718560200151611f69565b1561217b57600080fd5b6003600660008760200151815260200190815260200160002060000181905550600760008660200151815260208082019290925260400160009081208190556008918701518152602001908152602001600020600090557f7c59798283502bd302b18828d4f858808c79d023b3784676d6984b3c23ae45b5856020015160405190815260200160405180910390a184602001519e9d5050505050505050505050505050565b60035481565b60056020526000908152604090205481565b61224061383b565b6122486138f5565b612262600661225685612e81565b9063ffffffff612eb316565b905060e06040519081016040528061228f8360008151811061228057fe5b90602001906020020151612f4f565b81526020016122a48360018151811061228057fe5b81526020016122b98360028151811061228057fe5b81526020016122dd836003815181106122ce57fe5b90602001906020020151612f76565b600160a060020a031681526020016122fb8360048151811061228057fe5b815260200161230983612faf565b600160a060020a03168152602001612320836130fb565b90529392505050565b8181018281101561233657fe5b92915050565b6000808560205b845181116123a55780850151925060028706151561237b57818360405191825260208201526040908101905180910390209150612397565b8282604051918252602082015260409081019051809103902091505b600287049650602001612343565b509390931495945050505050565b60006123bd61394b565b60008581526009602052604090819020600201908590518082805190602001908083835b602083106124005780518252601f1990920191602091820191016123e1565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020549150811561244257600080fd5b606060405190810160409081526001825260208083018790528183018690526000888152600990915220549091501515612489576000858152600960205260409020600190555b600085815260096020908152604080832080548452600101909152902081908151815460ff19169015151781556020820151816001019080516124d0929160200190613877565b5060408201516002918201556000878152600960205260409081902080549350909101908690518082805190602001908083835b602083106125235780518252601f199092019160209182019101612504565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205560008581526009602052604090205461257890600163ffffffff61232916565b6000958652600960205260409095209490945550505050565b600061259b61394b565b6000868152600b60209081526040808320888452909152808220600201908690518082805190602001908083835b602083106125e85780518252601f1990920191602091820191016125c9565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020549250821561262a57600080fd5b6060604051908101604090815260018252602080830188905281830187905260008a8152600b82528281208a825290915220549092501515612684576000878152600b602090815260408083208984529091529020600190555b506000868152600b6020908152604080832088845282528083208054808552600190910190925290912082908151815460ff19169015151781556020820151816001019080516126d8929160200190613877565b5060408201516002918201556000898152600b602090815260408083208b845290915290819020849350909101908790518082805190602001908083835b602083106127355780518252601f199092019160209182019101612716565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205561277b81600163ffffffff61232916565b6000978852600b60209081526040808a20988a5297905295909620949094555050505050565b6000838152600b6020908152604080832085845290915280822082918291600201908590518082805190602001908083835b602083106127f25780518252601f1990920191602091820191016127d3565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902054925082151561283557600080fd5b6000868152600b60209081526040808320888452825280832086845260019081019092528220805460ff19168155919061287190830182613907565b5060006002918201819055878152600b6020908152604080832089845290915290819020909101908590518082805190602001908083835b602083106128c85780518252601f1990920191602091820191016128a9565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020600090819055868152600b6020908152604080832088845290915290205461292b90600163ffffffff612def16565b9150828214612a8257506000858152600b6020908152604080832087845282528083208484526001908101808452828520878652935292208154815460ff909116151560ff1990911617815581830180549293849361299f9284830192909160029181161561010002600019011604613973565b50600291820154908201556000878152600b6020908152604080832089845290915290819020859201906001840190518082805460018160011615610100020316600290048015612a275780601f10612a05576101008083540402835291820191612a27565b820191906000526020600020905b815481529060010190602001808311612a13575b50509283525050602001604051908190039020556000868152600b60209081526040808320888452825280832085845260019081019092528220805460ff191681559190612a7790830182613907565b600282016000905550505b8160011415612aaa576000868152600b60209081526040808320888452909152812055612ac7565b6000868152600b6020908152604080832088845290915290208290555b505050505050565b60008284836040516c01000000000000000000000000600160a060020a03948516810282529290931690910260148301526028820152604801604051809103902090509392505050565b60008281526009602052604080822082918291600201908590518082805190602001908083835b60208310612b5f5780518252601f199092019160209182019101612b40565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020549250821515612ba257600080fd5b600085815260096020908152604080832086845260019081019092528220805460ff191681559190612bd690830182613907565b6002820160009055505060096000868152602001908152602001600020600201846040518082805190602001908083835b60208310612c265780518252601f199092019160209182019101612c07565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902060009081905585815260096020526040902054612c7e90600163ffffffff612def16565b9150828214612db957506000848152600960209081526040808320848452600190810190925280832085845292208254815460ff191660ff909116151517815582820180548493612ce992848201929091600261010091831615919091026000190190911604613973565b506002918201549082015560008681526009602052604090819020859201906001840190518082805460018160011615610100020316600290048015612d665780601f10612d44576101008083540402835291820191612d66565b820191906000526020600020905b815481529060010190602001808311612d52575b5050928352505060200160405190819003902055600085815260096020908152604080832085845260019081019092528220805460ff191681559190612dae90830182613907565b600282016000905550505b8160011415612dd657600085815260096020526040812055612de8565b60008581526009602052604090208290555b5050505050565b600082821115612dfb57fe5b50900390565b600160a060020a0381161515612e1657600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b612e896139e8565b600080835191505060208301604080519081016040528181526020810183905292505b5050919050565b612ebb6138f5565b612ec36139ff565b600083604051805910612ed35750595b908082528060200260200182016040528015612f0957816020015b612ef66139e8565b815260200190600190039081612eee5790505b509250612f1585613207565b91505b83811015612f4757612f298261322c565b838281518110612f3557fe5b60209081029091010152600101612f18565b505092915050565b6000806000612f5d84613261565b909250905060208190036101000a825104949350505050565b6000806000612f8484613261565b909250905060148114612f9657600080fd5b6c01000000000000000000000000825104949350505050565b6000612fb96138f5565b6000612fc36138f5565b612fcb6138f5565b6005604051805910612fda5750595b90808252806020026020018201604052801561301057816020015b612ffd6138f5565b815260200190600190039081612ff55790505b509350600092505b60058310156130605761303f86848151811061303057fe5b906020019060200201516132c8565b84848151811061304b57fe5b60209081029091010152600190920191613018565b6130698461331c565b915061308a8660058151811061307b57fe5b90602001906020020151613348565b90506130f1826040518082805190602001908083835b602083106130bf5780518252601f1990920191602091820191016130a0565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051809103902082613398565b9695505050505050565b60006131056138f5565b600061310f6138f5565b600560405180591061311e5750595b90808252806020026020018201604052801561315457816020015b6131416138f5565b8152602001906001900390816131395790505b509250600091505b60058210156131955761317485838151811061303057fe5b83838151811061318057fe5b6020908102909101015260019091019061315c565b61319e8361331c565b9050806040518082805190602001908083835b602083106131d05780518252601f1990920191602091820191016131b1565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051809103902093505b505050919050565b61320f6139ff565b600061321a83613478565b83519383529092016020820152919050565b6132346139e8565b60008083602001519150613247826134f7565b828452602080850182905292019390910192909252919050565b60008080808085519150815160001a9250608083101561328757819450600193506132c0565b60b88310156132a557600186602001510393508160010194506132c0565b5060b619820180600160208801510303935080820160010194505b505050915091565b6132d06138f5565b600080836020015191508115156132e657612eac565b816040518059106132f45750595b818152601f19601f8301168101602001604052905092505060208201612eac81855184613589565b6133246138f5565b61332c6138f5565b6133346138f5565b61333d846135ce565b9150611ca6826136b2565b6133506138f5565b600080600061335e85613261565b9093509150816040518059106133715750595b818152601f19601f83011681016020016040529050935050602083016131ff818484613589565b60008060008084516041146133b0576000935061346f565b6020850151925060408501519150606085015160001a9050601b8160ff1610156133d857601b015b8060ff16601b141580156133f057508060ff16601c14155b156133fe576000935061346f565b6001868285856040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f1151561346357600080fd5b50506020604051035193505b50505092915050565b6000806000836020015115156134915760009250612eac565b83519050805160001a915060808210156134ae5760009250612eac565b60b88210806134c9575060c082101580156134c9575060f882105b156134d75760019250612eac565b60c08210156134ec5760b51982019250612eac565b5060f5190192915050565b600080825160001a905060808110156135135760019150611f9c565b60b881101561352857607e1981019150611f9c565b60c08110156135525760b78103806020036101000a60018501510480820160010193505050611f9c565b60f88110156135675760be1981019150611f9c565b60f78103806020036101000a6001850151048082016001019350505050919050565b60005b602082106135af578251845260208401935060208301925060208203915061358c565b6001826020036101000a03905080198351168185511617909352505050565b6135d66138f5565b6000806135e16138f5565b60006135eb6138f5565b60008094505b87518510156136205787858151811061360657fe5b9060200190602002015151909501946001909401936135f1565b8560405180591061362e5750595b8181526020601f909201601f191681018201604052600096509450840192505b87518510156136a65787858151811061366357fe5b906020019060200201519150506020810161368083828451613589565b87858151811061368c57fe5b90602001906020020151516001909501949092019161364e565b50919695505050505050565b6136ba6138f5565b60006136c46138f5565b60008060008060208801955087519250600090505b82816101000210156136f157600191820191016136d9565b60378311613766578260010160405180591061370a5750595b818152601f19601f8301168101602001604052905094508260c00160f860020a028560008151811061373857fe5b906020010190600160f860020a031916908160001a905350602185019350613761848785613589565b61382f565b8282600101016040518059106137795750595b818152601f19601f83011681016020016040529050945060f860020a60f7830102856000815181106137a757fe5b906020010190600160f860020a031916908160001a905350600190505b81811161381c576101008183036101000a848115156137df57fe5b048115156137e957fe5b0660f860020a028582815181106137fc57fe5b906020010190600160f860020a031916908160001a9053506001016137c4565b816021860101935061382f848785613589565b50929695505050505050565b60e06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a0820181905260c082015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106138b857805160ff19168380011785556138e5565b828001600101855582156138e5579182015b828111156138e55782518255916020019190600101906138ca565b506138f1929150613a13565b5090565b60206040519081016040526000815290565b50805460018160011615610100020316600290046000825580601f1061392d5750611fd8565b601f016020900490600052602060002090810190611fd89190613a13565b606060405190810160405260008152602081016139666138f5565b8152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106139ac57805485556138e5565b828001600101855582156138e557600052602060002091601f016020900482015b828111156138e55782548255916001019190600101906139cd565b604080519081016040526000808252602082015290565b6060604051908101604052806139666139e8565b613a2d91905b808211156138f15760008155600101613a19565b905600a165627a7a72305820b5b290a620dc9967c72ec1a943b466a34fed11bf5ad998900fcc9aa787eaec9b0029`

// DeployRootChain deploys a new Ethereum contract, binding an instance of RootChain to it.
func DeployRootChain(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address) (common.Address, *types.Transaction, *RootChain, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainBin), backend, _operator)
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

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() constant returns(uint256)
func (_RootChain *RootChainCaller) BlockNumber(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "blockNumber")
	return *ret0, err
}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() constant returns(uint256)
func (_RootChain *RootChainSession) BlockNumber() (*big.Int, error) {
	return _RootChain.Contract.BlockNumber(&_RootChain.CallOpts)
}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() constant returns(uint256)
func (_RootChain *RootChainCallerSession) BlockNumber() (*big.Int, error) {
	return _RootChain.Contract.BlockNumber(&_RootChain.CallOpts)
}

// ChallengeExists is a free data retrieval call binding the contract method 0x6e606251.
//
// Solidity: function challengeExists(uid uint256, challengeTx bytes) constant returns(bool)
func (_RootChain *RootChainCaller) ChallengeExists(opts *bind.CallOptsWithNumber, uid *big.Int, challengeTx []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "challengeExists", uid, challengeTx)
	return *ret0, err
}

// ChallengeExists is a free data retrieval call binding the contract method 0x6e606251.
//
// Solidity: function challengeExists(uid uint256, challengeTx bytes) constant returns(bool)
func (_RootChain *RootChainSession) ChallengeExists(uid *big.Int, challengeTx []byte) (bool, error) {
	return _RootChain.Contract.ChallengeExists(&_RootChain.CallOpts, uid, challengeTx)
}

// ChallengeExists is a free data retrieval call binding the contract method 0x6e606251.
//
// Solidity: function challengeExists(uid uint256, challengeTx bytes) constant returns(bool)
func (_RootChain *RootChainCallerSession) ChallengeExists(uid *big.Int, challengeTx []byte) (bool, error) {
	return _RootChain.Contract.ChallengeExists(&_RootChain.CallOpts, uid, challengeTx)
}

// ChallengePeriod is a free data retrieval call binding the contract method 0xf3f480d9.
//
// Solidity: function challengePeriod() constant returns(uint256)
func (_RootChain *RootChainCaller) ChallengePeriod(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "challengePeriod")
	return *ret0, err
}

// ChallengePeriod is a free data retrieval call binding the contract method 0xf3f480d9.
//
// Solidity: function challengePeriod() constant returns(uint256)
func (_RootChain *RootChainSession) ChallengePeriod() (*big.Int, error) {
	return _RootChain.Contract.ChallengePeriod(&_RootChain.CallOpts)
}

// ChallengePeriod is a free data retrieval call binding the contract method 0xf3f480d9.
//
// Solidity: function challengePeriod() constant returns(uint256)
func (_RootChain *RootChainCallerSession) ChallengePeriod() (*big.Int, error) {
	return _RootChain.Contract.ChallengePeriod(&_RootChain.CallOpts)
}

// ChallengesLength is a free data retrieval call binding the contract method 0xb15e42f8.
//
// Solidity: function challengesLength(uid uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) ChallengesLength(opts *bind.CallOptsWithNumber, uid *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "challengesLength", uid)
	return *ret0, err
}

// ChallengesLength is a free data retrieval call binding the contract method 0xb15e42f8.
//
// Solidity: function challengesLength(uid uint256) constant returns(uint256)
func (_RootChain *RootChainSession) ChallengesLength(uid *big.Int) (*big.Int, error) {
	return _RootChain.Contract.ChallengesLength(&_RootChain.CallOpts, uid)
}

// ChallengesLength is a free data retrieval call binding the contract method 0xb15e42f8.
//
// Solidity: function challengesLength(uid uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) ChallengesLength(uid *big.Int) (*big.Int, error) {
	return _RootChain.Contract.ChallengesLength(&_RootChain.CallOpts, uid)
}

// CheckpointChallengesLength is a free data retrieval call binding the contract method 0x9ced6890.
//
// Solidity: function checkpointChallengesLength(uid uint256, checkpoint bytes32) constant returns(uint256)
func (_RootChain *RootChainCaller) CheckpointChallengesLength(opts *bind.CallOptsWithNumber, uid *big.Int, checkpoint [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "checkpointChallengesLength", uid, checkpoint)
	return *ret0, err
}

// CheckpointChallengesLength is a free data retrieval call binding the contract method 0x9ced6890.
//
// Solidity: function checkpointChallengesLength(uid uint256, checkpoint bytes32) constant returns(uint256)
func (_RootChain *RootChainSession) CheckpointChallengesLength(uid *big.Int, checkpoint [32]byte) (*big.Int, error) {
	return _RootChain.Contract.CheckpointChallengesLength(&_RootChain.CallOpts, uid, checkpoint)
}

// CheckpointChallengesLength is a free data retrieval call binding the contract method 0x9ced6890.
//
// Solidity: function checkpointChallengesLength(uid uint256, checkpoint bytes32) constant returns(uint256)
func (_RootChain *RootChainCallerSession) CheckpointChallengesLength(uid *big.Int, checkpoint [32]byte) (*big.Int, error) {
	return _RootChain.Contract.CheckpointChallengesLength(&_RootChain.CallOpts, uid, checkpoint)
}

// CheckpointIsChallenge is a free data retrieval call binding the contract method 0x370beab9.
//
// Solidity: function checkpointIsChallenge(uid uint256, checkpoint bytes32, challengeTx bytes) constant returns(bool)
func (_RootChain *RootChainCaller) CheckpointIsChallenge(opts *bind.CallOptsWithNumber, uid *big.Int, checkpoint [32]byte, challengeTx []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "checkpointIsChallenge", uid, checkpoint, challengeTx)
	return *ret0, err
}

// CheckpointIsChallenge is a free data retrieval call binding the contract method 0x370beab9.
//
// Solidity: function checkpointIsChallenge(uid uint256, checkpoint bytes32, challengeTx bytes) constant returns(bool)
func (_RootChain *RootChainSession) CheckpointIsChallenge(uid *big.Int, checkpoint [32]byte, challengeTx []byte) (bool, error) {
	return _RootChain.Contract.CheckpointIsChallenge(&_RootChain.CallOpts, uid, checkpoint, challengeTx)
}

// CheckpointIsChallenge is a free data retrieval call binding the contract method 0x370beab9.
//
// Solidity: function checkpointIsChallenge(uid uint256, checkpoint bytes32, challengeTx bytes) constant returns(bool)
func (_RootChain *RootChainCallerSession) CheckpointIsChallenge(uid *big.Int, checkpoint [32]byte, challengeTx []byte) (bool, error) {
	return _RootChain.Contract.CheckpointIsChallenge(&_RootChain.CallOpts, uid, checkpoint, challengeTx)
}

// Checkpoints is a free data retrieval call binding the contract method 0xeb5e91ff.
//
// Solidity: function checkpoints( bytes32) constant returns(uint256)
func (_RootChain *RootChainCaller) Checkpoints(opts *bind.CallOptsWithNumber, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "checkpoints", arg0)
	return *ret0, err
}

// Checkpoints is a free data retrieval call binding the contract method 0xeb5e91ff.
//
// Solidity: function checkpoints( bytes32) constant returns(uint256)
func (_RootChain *RootChainSession) Checkpoints(arg0 [32]byte) (*big.Int, error) {
	return _RootChain.Contract.Checkpoints(&_RootChain.CallOpts, arg0)
}

// Checkpoints is a free data retrieval call binding the contract method 0xeb5e91ff.
//
// Solidity: function checkpoints( bytes32) constant returns(uint256)
func (_RootChain *RootChainCallerSession) Checkpoints(arg0 [32]byte) (*big.Int, error) {
	return _RootChain.Contract.Checkpoints(&_RootChain.CallOpts, arg0)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(bytes32)
func (_RootChain *RootChainCaller) ChildChain(opts *bind.CallOptsWithNumber, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "childChain", arg0)
	return *ret0, err
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(bytes32)
func (_RootChain *RootChainSession) ChildChain(arg0 *big.Int) ([32]byte, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(bytes32)
func (_RootChain *RootChainCallerSession) ChildChain(arg0 *big.Int) ([32]byte, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() constant returns(uint256)
func (_RootChain *RootChainCaller) DepositCount(opts *bind.CallOptsWithNumber) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "depositCount")
	return *ret0, err
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() constant returns(uint256)
func (_RootChain *RootChainSession) DepositCount() (*big.Int, error) {
	return _RootChain.Contract.DepositCount(&_RootChain.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() constant returns(uint256)
func (_RootChain *RootChainCallerSession) DepositCount() (*big.Int, error) {
	return _RootChain.Contract.DepositCount(&_RootChain.CallOpts)
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(state uint256, exitTime uint256, exitTxBlkNum uint256, exitTx bytes, txBeforeExitTxBlkNum uint256, txBeforeExitTx bytes)
func (_RootChain *RootChainCaller) Exits(opts *bind.CallOptsWithNumber, arg0 *big.Int) (struct {
	State                *big.Int
	ExitTime             *big.Int
	ExitTxBlkNum         *big.Int
	ExitTx               []byte
	TxBeforeExitTxBlkNum *big.Int
	TxBeforeExitTx       []byte
}, error) {
	ret := new(struct {
		State                *big.Int
		ExitTime             *big.Int
		ExitTxBlkNum         *big.Int
		ExitTx               []byte
		TxBeforeExitTxBlkNum *big.Int
		TxBeforeExitTx       []byte
	})
	out := ret
	err := _RootChain.contract.CallWithNumber(opts, out, "exits", arg0)
	return *ret, err
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(state uint256, exitTime uint256, exitTxBlkNum uint256, exitTx bytes, txBeforeExitTxBlkNum uint256, txBeforeExitTx bytes)
func (_RootChain *RootChainSession) Exits(arg0 *big.Int) (struct {
	State                *big.Int
	ExitTime             *big.Int
	ExitTxBlkNum         *big.Int
	ExitTx               []byte
	TxBeforeExitTxBlkNum *big.Int
	TxBeforeExitTx       []byte
}, error) {
	return _RootChain.Contract.Exits(&_RootChain.CallOpts, arg0)
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(state uint256, exitTime uint256, exitTxBlkNum uint256, exitTx bytes, txBeforeExitTxBlkNum uint256, txBeforeExitTx bytes)
func (_RootChain *RootChainCallerSession) Exits(arg0 *big.Int) (struct {
	State                *big.Int
	ExitTime             *big.Int
	ExitTxBlkNum         *big.Int
	ExitTx               []byte
	TxBeforeExitTxBlkNum *big.Int
	TxBeforeExitTx       []byte
}, error) {
	return _RootChain.Contract.Exits(&_RootChain.CallOpts, arg0)
}

// GetChallenge is a free data retrieval call binding the contract method 0x2e8f396d.
//
// Solidity: function getChallenge(uid uint256, index uint256) constant returns(challengeTx bytes, challengeBlock uint256)
func (_RootChain *RootChainCaller) GetChallenge(opts *bind.CallOptsWithNumber, uid *big.Int, index *big.Int) (struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
}, error) {
	ret := new(struct {
		ChallengeTx    []byte
		ChallengeBlock *big.Int
	})
	out := ret
	err := _RootChain.contract.CallWithNumber(opts, out, "getChallenge", uid, index)
	return *ret, err
}

// GetChallenge is a free data retrieval call binding the contract method 0x2e8f396d.
//
// Solidity: function getChallenge(uid uint256, index uint256) constant returns(challengeTx bytes, challengeBlock uint256)
func (_RootChain *RootChainSession) GetChallenge(uid *big.Int, index *big.Int) (struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
}, error) {
	return _RootChain.Contract.GetChallenge(&_RootChain.CallOpts, uid, index)
}

// GetChallenge is a free data retrieval call binding the contract method 0x2e8f396d.
//
// Solidity: function getChallenge(uid uint256, index uint256) constant returns(challengeTx bytes, challengeBlock uint256)
func (_RootChain *RootChainCallerSession) GetChallenge(uid *big.Int, index *big.Int) (struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
}, error) {
	return _RootChain.Contract.GetChallenge(&_RootChain.CallOpts, uid, index)
}

// GetCheckpointChallenge is a free data retrieval call binding the contract method 0x9af35551.
//
// Solidity: function getCheckpointChallenge(uid uint256, checkpoint bytes32, index uint256) constant returns(challengeTx bytes, challengeBlock uint256)
func (_RootChain *RootChainCaller) GetCheckpointChallenge(opts *bind.CallOptsWithNumber, uid *big.Int, checkpoint [32]byte, index *big.Int) (struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
}, error) {
	ret := new(struct {
		ChallengeTx    []byte
		ChallengeBlock *big.Int
	})
	out := ret
	err := _RootChain.contract.CallWithNumber(opts, out, "getCheckpointChallenge", uid, checkpoint, index)
	return *ret, err
}

// GetCheckpointChallenge is a free data retrieval call binding the contract method 0x9af35551.
//
// Solidity: function getCheckpointChallenge(uid uint256, checkpoint bytes32, index uint256) constant returns(challengeTx bytes, challengeBlock uint256)
func (_RootChain *RootChainSession) GetCheckpointChallenge(uid *big.Int, checkpoint [32]byte, index *big.Int) (struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
}, error) {
	return _RootChain.Contract.GetCheckpointChallenge(&_RootChain.CallOpts, uid, checkpoint, index)
}

// GetCheckpointChallenge is a free data retrieval call binding the contract method 0x9af35551.
//
// Solidity: function getCheckpointChallenge(uid uint256, checkpoint bytes32, index uint256) constant returns(challengeTx bytes, challengeBlock uint256)
func (_RootChain *RootChainCallerSession) GetCheckpointChallenge(uid *big.Int, checkpoint [32]byte, index *big.Int) (struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
}, error) {
	return _RootChain.Contract.GetCheckpointChallenge(&_RootChain.CallOpts, uid, checkpoint, index)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChain *RootChainCaller) Operator(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChain *RootChainSession) Operator() (common.Address, error) {
	return _RootChain.Contract.Operator(&_RootChain.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChain *RootChainCallerSession) Operator() (common.Address, error) {
	return _RootChain.Contract.Operator(&_RootChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RootChain *RootChainCaller) Owner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RootChain *RootChainSession) Owner() (common.Address, error) {
	return _RootChain.Contract.Owner(&_RootChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RootChain *RootChainCallerSession) Owner() (common.Address, error) {
	return _RootChain.Contract.Owner(&_RootChain.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x449b2f44.
//
// Solidity: function wallet( bytes32) constant returns(uint256)
func (_RootChain *RootChainCaller) Wallet(opts *bind.CallOptsWithNumber, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "wallet", arg0)
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x449b2f44.
//
// Solidity: function wallet( bytes32) constant returns(uint256)
func (_RootChain *RootChainSession) Wallet(arg0 [32]byte) (*big.Int, error) {
	return _RootChain.Contract.Wallet(&_RootChain.CallOpts, arg0)
}

// Wallet is a free data retrieval call binding the contract method 0x449b2f44.
//
// Solidity: function wallet( bytes32) constant returns(uint256)
func (_RootChain *RootChainCallerSession) Wallet(arg0 [32]byte) (*big.Int, error) {
	return _RootChain.Contract.Wallet(&_RootChain.CallOpts, arg0)
}

// Wallet2 is a free data retrieval call binding the contract method 0x26f3d093.
//
// Solidity: function wallet2( uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) Wallet2(opts *bind.CallOptsWithNumber, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.CallWithNumber(opts, out, "wallet2", arg0)
	return *ret0, err
}

// Wallet2 is a free data retrieval call binding the contract method 0x26f3d093.
//
// Solidity: function wallet2( uint256) constant returns(uint256)
func (_RootChain *RootChainSession) Wallet2(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.Wallet2(&_RootChain.CallOpts, arg0)
}

// Wallet2 is a free data retrieval call binding the contract method 0x26f3d093.
//
// Solidity: function wallet2( uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) Wallet2(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.Wallet2(&_RootChain.CallOpts, arg0)
}

// ChallengeCheckpoint is a paid mutator transaction binding the contract method 0x5e8cdf44.
//
// Solidity: function challengeCheckpoint(uid uint256, checkpointRoot bytes32, checkpointProof bytes, wrongNonce uint256, lastTx bytes, lastTxProof bytes, lastTxBlockNum uint256) returns()
func (_RootChain *RootChainTransactor) ChallengeCheckpoint(opts *bind.TransactOpts, uid *big.Int, checkpointRoot [32]byte, checkpointProof []byte, wrongNonce *big.Int, lastTx []byte, lastTxProof []byte, lastTxBlockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeCheckpoint", uid, checkpointRoot, checkpointProof, wrongNonce, lastTx, lastTxProof, lastTxBlockNum)
}

// ChallengeCheckpoint is a paid mutator transaction binding the contract method 0x5e8cdf44.
//
// Solidity: function challengeCheckpoint(uid uint256, checkpointRoot bytes32, checkpointProof bytes, wrongNonce uint256, lastTx bytes, lastTxProof bytes, lastTxBlockNum uint256) returns()
func (_RootChain *RootChainSession) ChallengeCheckpoint(uid *big.Int, checkpointRoot [32]byte, checkpointProof []byte, wrongNonce *big.Int, lastTx []byte, lastTxProof []byte, lastTxBlockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeCheckpoint(&_RootChain.TransactOpts, uid, checkpointRoot, checkpointProof, wrongNonce, lastTx, lastTxProof, lastTxBlockNum)
}

// ChallengeCheckpoint is a paid mutator transaction binding the contract method 0x5e8cdf44.
//
// Solidity: function challengeCheckpoint(uid uint256, checkpointRoot bytes32, checkpointProof bytes, wrongNonce uint256, lastTx bytes, lastTxProof bytes, lastTxBlockNum uint256) returns()
func (_RootChain *RootChainTransactorSession) ChallengeCheckpoint(uid *big.Int, checkpointRoot [32]byte, checkpointProof []byte, wrongNonce *big.Int, lastTx []byte, lastTxProof []byte, lastTxBlockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeCheckpoint(&_RootChain.TransactOpts, uid, checkpointRoot, checkpointProof, wrongNonce, lastTx, lastTxProof, lastTxBlockNum)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x4a2ee198.
//
// Solidity: function challengeExit(uid uint256, challengeTx bytes, proof bytes, challengeBlockNum uint256) returns()
func (_RootChain *RootChainTransactor) ChallengeExit(opts *bind.TransactOpts, uid *big.Int, challengeTx []byte, proof []byte, challengeBlockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeExit", uid, challengeTx, proof, challengeBlockNum)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x4a2ee198.
//
// Solidity: function challengeExit(uid uint256, challengeTx bytes, proof bytes, challengeBlockNum uint256) returns()
func (_RootChain *RootChainSession) ChallengeExit(uid *big.Int, challengeTx []byte, proof []byte, challengeBlockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, uid, challengeTx, proof, challengeBlockNum)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x4a2ee198.
//
// Solidity: function challengeExit(uid uint256, challengeTx bytes, proof bytes, challengeBlockNum uint256) returns()
func (_RootChain *RootChainTransactorSession) ChallengeExit(uid *big.Int, challengeTx []byte, proof []byte, challengeBlockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, uid, challengeTx, proof, challengeBlockNum)
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

// NewBlock is a paid mutator transaction binding the contract method 0x981adca5.
//
// Solidity: function newBlock(hash bytes32) returns()
func (_RootChain *RootChainTransactor) NewBlock(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "newBlock", hash)
}

// NewBlock is a paid mutator transaction binding the contract method 0x981adca5.
//
// Solidity: function newBlock(hash bytes32) returns()
func (_RootChain *RootChainSession) NewBlock(hash [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.NewBlock(&_RootChain.TransactOpts, hash)
}

// NewBlock is a paid mutator transaction binding the contract method 0x981adca5.
//
// Solidity: function newBlock(hash bytes32) returns()
func (_RootChain *RootChainTransactorSession) NewBlock(hash [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.NewBlock(&_RootChain.TransactOpts, hash)
}

// NewCheckpoint is a paid mutator transaction binding the contract method 0xa5e450f4.
//
// Solidity: function newCheckpoint(hash bytes32) returns()
func (_RootChain *RootChainTransactor) NewCheckpoint(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "newCheckpoint", hash)
}

// NewCheckpoint is a paid mutator transaction binding the contract method 0xa5e450f4.
//
// Solidity: function newCheckpoint(hash bytes32) returns()
func (_RootChain *RootChainSession) NewCheckpoint(hash [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.NewCheckpoint(&_RootChain.TransactOpts, hash)
}

// NewCheckpoint is a paid mutator transaction binding the contract method 0xa5e450f4.
//
// Solidity: function newCheckpoint(hash bytes32) returns()
func (_RootChain *RootChainTransactorSession) NewCheckpoint(hash [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.NewCheckpoint(&_RootChain.TransactOpts, hash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootChain *RootChainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootChain *RootChainSession) RenounceOwnership() (*types.Transaction, error) {
	return _RootChain.Contract.RenounceOwnership(&_RootChain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootChain *RootChainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RootChain.Contract.RenounceOwnership(&_RootChain.TransactOpts)
}

// RespondChallengeExit is a paid mutator transaction binding the contract method 0x8b593183.
//
// Solidity: function respondChallengeExit(uid uint256, challengeTx bytes, respondTx bytes, proof bytes, blockNum uint256) returns()
func (_RootChain *RootChainTransactor) RespondChallengeExit(opts *bind.TransactOpts, uid *big.Int, challengeTx []byte, respondTx []byte, proof []byte, blockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "respondChallengeExit", uid, challengeTx, respondTx, proof, blockNum)
}

// RespondChallengeExit is a paid mutator transaction binding the contract method 0x8b593183.
//
// Solidity: function respondChallengeExit(uid uint256, challengeTx bytes, respondTx bytes, proof bytes, blockNum uint256) returns()
func (_RootChain *RootChainSession) RespondChallengeExit(uid *big.Int, challengeTx []byte, respondTx []byte, proof []byte, blockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.RespondChallengeExit(&_RootChain.TransactOpts, uid, challengeTx, respondTx, proof, blockNum)
}

// RespondChallengeExit is a paid mutator transaction binding the contract method 0x8b593183.
//
// Solidity: function respondChallengeExit(uid uint256, challengeTx bytes, respondTx bytes, proof bytes, blockNum uint256) returns()
func (_RootChain *RootChainTransactorSession) RespondChallengeExit(uid *big.Int, challengeTx []byte, respondTx []byte, proof []byte, blockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.RespondChallengeExit(&_RootChain.TransactOpts, uid, challengeTx, respondTx, proof, blockNum)
}

// RespondChallengeExitWithCheckpoint is a paid mutator transaction binding the contract method 0x9f26026e.
//
// Solidity: function respondChallengeExitWithCheckpoint(uid uint256, challengeTx bytes, checkpointRoot bytes32, checkpointProof bytes, checkpointNonce bytes32) returns()
func (_RootChain *RootChainTransactor) RespondChallengeExitWithCheckpoint(opts *bind.TransactOpts, uid *big.Int, challengeTx []byte, checkpointRoot [32]byte, checkpointProof []byte, checkpointNonce [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "respondChallengeExitWithCheckpoint", uid, challengeTx, checkpointRoot, checkpointProof, checkpointNonce)
}

// RespondChallengeExitWithCheckpoint is a paid mutator transaction binding the contract method 0x9f26026e.
//
// Solidity: function respondChallengeExitWithCheckpoint(uid uint256, challengeTx bytes, checkpointRoot bytes32, checkpointProof bytes, checkpointNonce bytes32) returns()
func (_RootChain *RootChainSession) RespondChallengeExitWithCheckpoint(uid *big.Int, challengeTx []byte, checkpointRoot [32]byte, checkpointProof []byte, checkpointNonce [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.RespondChallengeExitWithCheckpoint(&_RootChain.TransactOpts, uid, challengeTx, checkpointRoot, checkpointProof, checkpointNonce)
}

// RespondChallengeExitWithCheckpoint is a paid mutator transaction binding the contract method 0x9f26026e.
//
// Solidity: function respondChallengeExitWithCheckpoint(uid uint256, challengeTx bytes, checkpointRoot bytes32, checkpointProof bytes, checkpointNonce bytes32) returns()
func (_RootChain *RootChainTransactorSession) RespondChallengeExitWithCheckpoint(uid *big.Int, challengeTx []byte, checkpointRoot [32]byte, checkpointProof []byte, checkpointNonce [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.RespondChallengeExitWithCheckpoint(&_RootChain.TransactOpts, uid, challengeTx, checkpointRoot, checkpointProof, checkpointNonce)
}

// RespondCheckpointChallenge is a paid mutator transaction binding the contract method 0xad131e13.
//
// Solidity: function respondCheckpointChallenge(uid uint256, checkpointRoot bytes32, challengeTx bytes, respondTx bytes, proof bytes, blockNum uint256) returns()
func (_RootChain *RootChainTransactor) RespondCheckpointChallenge(opts *bind.TransactOpts, uid *big.Int, checkpointRoot [32]byte, challengeTx []byte, respondTx []byte, proof []byte, blockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "respondCheckpointChallenge", uid, checkpointRoot, challengeTx, respondTx, proof, blockNum)
}

// RespondCheckpointChallenge is a paid mutator transaction binding the contract method 0xad131e13.
//
// Solidity: function respondCheckpointChallenge(uid uint256, checkpointRoot bytes32, challengeTx bytes, respondTx bytes, proof bytes, blockNum uint256) returns()
func (_RootChain *RootChainSession) RespondCheckpointChallenge(uid *big.Int, checkpointRoot [32]byte, challengeTx []byte, respondTx []byte, proof []byte, blockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.RespondCheckpointChallenge(&_RootChain.TransactOpts, uid, checkpointRoot, challengeTx, respondTx, proof, blockNum)
}

// RespondCheckpointChallenge is a paid mutator transaction binding the contract method 0xad131e13.
//
// Solidity: function respondCheckpointChallenge(uid uint256, checkpointRoot bytes32, challengeTx bytes, respondTx bytes, proof bytes, blockNum uint256) returns()
func (_RootChain *RootChainTransactorSession) RespondCheckpointChallenge(uid *big.Int, checkpointRoot [32]byte, challengeTx []byte, respondTx []byte, proof []byte, blockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.RespondCheckpointChallenge(&_RootChain.TransactOpts, uid, checkpointRoot, challengeTx, respondTx, proof, blockNum)
}

// RespondWithHistoricalCheckpoint is a paid mutator transaction binding the contract method 0x6d2975c7.
//
// Solidity: function respondWithHistoricalCheckpoint(uid uint256, checkpointRoot bytes32, checkpointProof bytes, historicalCheckpointRoot bytes32, historicalCheckpointProof bytes, challengeTx bytes, moreNonce uint256) returns()
func (_RootChain *RootChainTransactor) RespondWithHistoricalCheckpoint(opts *bind.TransactOpts, uid *big.Int, checkpointRoot [32]byte, checkpointProof []byte, historicalCheckpointRoot [32]byte, historicalCheckpointProof []byte, challengeTx []byte, moreNonce *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "respondWithHistoricalCheckpoint", uid, checkpointRoot, checkpointProof, historicalCheckpointRoot, historicalCheckpointProof, challengeTx, moreNonce)
}

// RespondWithHistoricalCheckpoint is a paid mutator transaction binding the contract method 0x6d2975c7.
//
// Solidity: function respondWithHistoricalCheckpoint(uid uint256, checkpointRoot bytes32, checkpointProof bytes, historicalCheckpointRoot bytes32, historicalCheckpointProof bytes, challengeTx bytes, moreNonce uint256) returns()
func (_RootChain *RootChainSession) RespondWithHistoricalCheckpoint(uid *big.Int, checkpointRoot [32]byte, checkpointProof []byte, historicalCheckpointRoot [32]byte, historicalCheckpointProof []byte, challengeTx []byte, moreNonce *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.RespondWithHistoricalCheckpoint(&_RootChain.TransactOpts, uid, checkpointRoot, checkpointProof, historicalCheckpointRoot, historicalCheckpointProof, challengeTx, moreNonce)
}

// RespondWithHistoricalCheckpoint is a paid mutator transaction binding the contract method 0x6d2975c7.
//
// Solidity: function respondWithHistoricalCheckpoint(uid uint256, checkpointRoot bytes32, checkpointProof bytes, historicalCheckpointRoot bytes32, historicalCheckpointProof bytes, challengeTx bytes, moreNonce uint256) returns()
func (_RootChain *RootChainTransactorSession) RespondWithHistoricalCheckpoint(uid *big.Int, checkpointRoot [32]byte, checkpointProof []byte, historicalCheckpointRoot [32]byte, historicalCheckpointProof []byte, challengeTx []byte, moreNonce *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.RespondWithHistoricalCheckpoint(&_RootChain.TransactOpts, uid, checkpointRoot, checkpointProof, historicalCheckpointRoot, historicalCheckpointProof, challengeTx, moreNonce)
}

// StartExit is a paid mutator transaction binding the contract method 0x04675c65.
//
// Solidity: function startExit(previousTx bytes, previousTxProof bytes, previousTxBlockNum uint256, lastTx bytes, lastTxProof bytes, lastTxBlockNum uint256) returns()
func (_RootChain *RootChainTransactor) StartExit(opts *bind.TransactOpts, previousTx []byte, previousTxProof []byte, previousTxBlockNum *big.Int, lastTx []byte, lastTxProof []byte, lastTxBlockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startExit", previousTx, previousTxProof, previousTxBlockNum, lastTx, lastTxProof, lastTxBlockNum)
}

// StartExit is a paid mutator transaction binding the contract method 0x04675c65.
//
// Solidity: function startExit(previousTx bytes, previousTxProof bytes, previousTxBlockNum uint256, lastTx bytes, lastTxProof bytes, lastTxBlockNum uint256) returns()
func (_RootChain *RootChainSession) StartExit(previousTx []byte, previousTxProof []byte, previousTxBlockNum *big.Int, lastTx []byte, lastTxProof []byte, lastTxBlockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, previousTx, previousTxProof, previousTxBlockNum, lastTx, lastTxProof, lastTxBlockNum)
}

// StartExit is a paid mutator transaction binding the contract method 0x04675c65.
//
// Solidity: function startExit(previousTx bytes, previousTxProof bytes, previousTxBlockNum uint256, lastTx bytes, lastTxProof bytes, lastTxBlockNum uint256) returns()
func (_RootChain *RootChainTransactorSession) StartExit(previousTx []byte, previousTxProof []byte, previousTxBlockNum *big.Int, lastTx []byte, lastTxProof []byte, lastTxBlockNum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, previousTx, previousTxProof, previousTxBlockNum, lastTx, lastTxProof, lastTxBlockNum)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RootChain *RootChainTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RootChain *RootChainSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.TransferOwnership(&_RootChain.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RootChain *RootChainTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.TransferOwnership(&_RootChain.TransactOpts, _newOwner)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a723058202a93768167fc29533133884e902f0af9a3dc3941e8c6ecd3eab60b65541f8dce0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath               // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller         // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TransactionABI is the input ABI used to generate the binding from.
const TransactionABI = "[]"

// TransactionBin is the compiled bytecode used for deploying new contracts.
const TransactionBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820c7f0fbef8aeb7bfd61b58f92382d5521a9738c8486b1488d98edff0ca1f51c090029`

// DeployTransaction deploys a new Ethereum contract, binding an instance of Transaction to it.
func DeployTransaction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TransactionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Transaction{TransactionCaller: TransactionCaller{contract: contract}, TransactionTransactor: TransactionTransactor{contract: contract}}, nil
}

// Transaction is an auto generated Go binding around an Ethereum contract.
type Transaction struct {
	TransactionCaller     // Read-only binding to the contract
	TransactionTransactor // Write-only binding to the contract
}

// TransactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransactionSession struct {
	Contract     *Transaction            // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TransactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransactionCallerSession struct {
	Contract *TransactionCaller      // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// TransactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransactionTransactorSession struct {
	Contract     *TransactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TransactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransactionRaw struct {
	Contract *Transaction // Generic contract binding to access the raw methods on
}

// TransactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransactionCallerRaw struct {
	Contract *TransactionCaller // Generic read-only contract binding to access the raw methods on
}

// TransactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactionTransactorRaw struct {
	Contract *TransactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransaction creates a new instance of Transaction, bound to a specific deployed contract.
func NewTransaction(address common.Address, backend bind.ContractBackend) (*Transaction, error) {
	contract, err := bindTransaction(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transaction{TransactionCaller: TransactionCaller{contract: contract}, TransactionTransactor: TransactionTransactor{contract: contract}}, nil
}

// NewTransactionCaller creates a new read-only instance of Transaction, bound to a specific deployed contract.
func NewTransactionCaller(address common.Address, caller bind.ContractCaller) (*TransactionCaller, error) {
	contract, err := bindTransaction(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionCaller{contract: contract}, nil
}

// NewTransactionTransactor creates a new write-only instance of Transaction, bound to a specific deployed contract.
func NewTransactionTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactionTransactor, error) {
	contract, err := bindTransaction(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TransactionTransactor{contract: contract}, nil
}

// bindTransaction binds a generic wrapper to an already deployed contract.
func bindTransaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transaction *TransactionRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Transaction.Contract.TransactionCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transaction *TransactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.Contract.TransactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transaction *TransactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transaction.Contract.TransactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transaction *TransactionCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Transaction.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transaction *TransactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transaction *TransactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transaction.Contract.contract.Transact(opts, method, params...)
}
