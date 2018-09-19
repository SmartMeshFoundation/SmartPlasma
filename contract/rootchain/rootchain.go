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
const RootChainABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"previousTx\",\"type\":\"bytes\"},{\"name\":\"previousTxProof\",\"type\":\"bytes\"},{\"name\":\"previousTxBlockNum\",\"type\":\"uint256\"},{\"name\":\"lastTx\",\"type\":\"bytes\"},{\"name\":\"lastTxProof\",\"type\":\"bytes\"},{\"name\":\"lastTxBlockNum\",\"type\":\"uint256\"}],\"name\":\"startExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wallet2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChallenge\",\"outputs\":[{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"challengeBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exits\",\"outputs\":[{\"name\":\"state\",\"type\":\"uint256\"},{\"name\":\"exitTime\",\"type\":\"uint256\"},{\"name\":\"exitTxBlkNum\",\"type\":\"uint256\"},{\"name\":\"exitTx\",\"type\":\"bytes\"},{\"name\":\"txBeforeExitTxBlkNum\",\"type\":\"uint256\"},{\"name\":\"txBeforeExitTx\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpoint\",\"type\":\"bytes32\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"}],\"name\":\"checkpointIsChallenge\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"challengeBlockNum\",\"type\":\"uint256\"}],\"name\":\"challengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpointRoot\",\"type\":\"bytes32\"},{\"name\":\"checkpointProof\",\"type\":\"bytes\"},{\"name\":\"wrongNonce\",\"type\":\"uint256\"},{\"name\":\"lastTx\",\"type\":\"bytes\"},{\"name\":\"lastTxProof\",\"type\":\"bytes\"},{\"name\":\"lastTxBlockNum\",\"type\":\"uint256\"}],\"name\":\"challengeCheckpoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpointRoot\",\"type\":\"bytes32\"},{\"name\":\"checkpointProof\",\"type\":\"bytes\"},{\"name\":\"historicalCheckpointRoot\",\"type\":\"bytes32\"},{\"name\":\"historicalCheckpointProof\",\"type\":\"bytes\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"moreNonce\",\"type\":\"uint256\"}],\"name\":\"respondWithHistoricalCheckpoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"}],\"name\":\"challengeExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"currency\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"respondTx\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"respondChallengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"newBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpoint\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getCheckpointChallenge\",\"outputs\":[{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"challengeBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpoint\",\"type\":\"bytes32\"}],\"name\":\"checkpointChallengesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"newCheckpoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpointRoot\",\"type\":\"bytes32\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"respondTx\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"respondCheckpointChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"challengesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"previousTx\",\"type\":\"bytes\"},{\"name\":\"previousTxProof\",\"type\":\"bytes\"},{\"name\":\"previousTxBlockNum\",\"type\":\"uint256\"},{\"name\":\"lastTx\",\"type\":\"bytes\"},{\"name\":\"lastTxProof\",\"type\":\"bytes\"},{\"name\":\"lastTxBlockNum\",\"type\":\"uint256\"}],\"name\":\"finishExit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"NewBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"NewCheckpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastBlock\",\"type\":\"uint256\"}],\"name\":\"StartExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"FinishExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"ChallengeExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"checkpoint\",\"type\":\"bytes32\"}],\"name\":\"ChallengeCheckpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"RespondChallengeExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"checkpoint\",\"type\":\"bytes32\"}],\"name\":\"RespondCheckpointChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"checkpoint\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"historicalCheckpoint\",\"type\":\"bytes32\"}],\"name\":\"RespondWithHistoricalCheckpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `0x6060604052341561000f57600080fd5b60405160208061391c8339810160405280805160008054600160a060020a03338116600160a060020a03199283161783556002839055621275006003556001929092556004805492909316911617905550506138ac806100706000396000f3006060604052600436106101695763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166304675c65811461016e57806326f3d093146102925780632dfdf0b5146102ba5780632e8f396d146102cd578063342de17914610365578063370beab91461049d578063449b2f441461050d5780634a2ee19814610523578063570ca735146105bd57806357e871e7146105ec5780635e8cdf44146105ff5780636d2975c7146106ea5780636e606251146107d1578063715018a6146108275780638340f5491461083a5780638b593183146108625780638da5cb5b1461093e578063981adca5146109515780639af35551146109675780639ced689014610983578063a5e450f41461099c578063ad131e13146109b2578063b15e42f814610a94578063eb5e91ff14610aaa578063f2fde38b14610ac0578063f310f2b214610adf578063f3f480d914610c0f578063f95643b114610c22575b600080fd5b341561017957600080fd5b61029060046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505093359350610c3892505050565b005b341561029d57600080fd5b6102a8600435610ef3565b60405190815260200160405180910390f35b34156102c557600080fd5b6102a8610f05565b34156102d857600080fd5b6102e6600435602435610f0b565b6040518080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610329578082015183820152602001610311565b50505050905090810190601f1680156103565780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b341561037057600080fd5b61037b600435610ffd565b60405186815260208101869052604081018590526080810183905260c0606082018181528554600260001961010060018416150201909116049183018290529060a083019060e0840190879080156104145780601f106103e957610100808354040283529160200191610414565b820191906000526020600020905b8154815290600101906020018083116103f757829003601f168201915b50508381038252845460026000196101006001841615020190911604808252602090910190859080156104885780601f1061045d57610100808354040283529160200191610488565b820191906000526020600020905b81548152906001019060200180831161046b57829003601f168201915b50509850505050505050505060405180910390f35b34156104a857600080fd5b6104f9600480359060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061102e95505050505050565b604051901515815260200160405180910390f35b341561051857600080fd5b6102a86004356110fa565b341561052e57600080fd5b610290600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061110c92505050565b34156105c857600080fd5b6105d0611475565b604051600160a060020a03909116815260200160405180910390f35b34156105f757600080fd5b6102a8611484565b341561060a57600080fd5b610290600480359060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061148a92505050565b34156106f557600080fd5b610290600480359060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094968635969095506040808201955060209182013587018083019550359350839250601f830182900482029091019051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965050933593506115b492505050565b34156107dc57600080fd5b6104f9600480359060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506116ba95505050505050565b341561083257600080fd5b610290611775565b341561084557600080fd5b6102a8600160a060020a03600435811690602435166044356117e7565b341561086d57600080fd5b610290600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965050933593506118ad92505050565b341561094957600080fd5b6105d0611a36565b341561095c57600080fd5b610290600435611a45565b341561097257600080fd5b6102e6600435602435604435611ac2565b341561098e57600080fd5b6102a8600435602435611bc0565b34156109a757600080fd5b610290600435611c03565b34156109bd57600080fd5b610290600480359060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505093359350611c8092505050565b3415610a9f57600080fd5b6102a8600435611db9565b3415610ab557600080fd5b6102a8600435611df2565b3415610acb57600080fd5b610290600160a060020a0360043516611e04565b3415610aea57600080fd5b6102a860048035600160a060020a03169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505093359350611e2b92505050565b3415610c1a57600080fd5b6102a8612070565b3415610c2d57600080fd5b6102a8600435612076565b610c4061368b565b610c4861368b565b600080600080610c578c612088565b9550610c6289612088565b945084518a14610c7157600080fd5b8460200151866020015114610c8557600080fd5b8460400151866040015114610c9957600080fd5b8460a00151600160a060020a03168660600151600160a060020a031614610cbf57600080fd5b610cd5600187608001519063ffffffff61217916565b856080015114610ce457600080fd5b8460600151600160a060020a031633600160a060020a0316141515610d0857600080fd5b60076000866020015181526020810191909152604001600020541515610d2d57600080fd5b8560c0015160008b815260056020526040902054909450925060c0850151915060056000888152602001908152602001600020549050610d7a86602001518590858e63ffffffff61218c16565b1515610d8557600080fd5b610d9c85602001518390838b63ffffffff61218c16565b1515610da757600080fd5b600660008660200151815260208101919091526040016000205415610dcb57600080fd5b610dd88560200151611db9565b15610de257600080fd5b60c06040519081016040528060028152602001610e0a6003544261217990919063ffffffff16565b81526020018881526020018a81526020018b81526020018d8152506006600087602001518152602001908152602001600020600082015181556020820151816001015560408201518160020155606082015181600301908051610e719291602001906136c7565b506080820151816004015560a082015181600501908051610e969291602001906136c7565b509050507f4d3db44958203fdb34cb08c6f6ecda2c9e182cd426b0871efec5ade7ac94580d86602001518b8960405180848152602001838152602001828152602001935050505060405180910390a1505050505050505050505050565b60086020526000908152604090205481565b60015481565b610f13613745565b600083815260096020526040812081906001908101908290610f3c90879063ffffffff61217916565b81526020019081526020016000209050806001018160020154818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fea5780601f10610fbf57610100808354040283529160200191610fea565b820191906000526020600020905b815481529060010190602001808311610fcd57829003601f168201915b5050505050915092509250509250929050565b6006602052600090815260409020805460018201546002830154600484015492939192909160038101919060050186565b6000838152600b6020908152604080832085845290915280822082916002909101908490518082805190602001908083835b6020831061107f5780518252601f199092019160209182019101611060565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054905080600014156110c757600091506110f2565b6000858152600b60209081526040808320878452825280832084845260010190915290205460ff1691505b509392505050565b60076020526000908152604090205481565b61111461368b565b61111c61368b565b61112461368b565b600087815260066020526040812054819060021461114157600080fd5b6111f8600660008b81526020019081526020016000206003018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111ee5780601f106111c3576101008083540402835291602001916111ee565b820191906000526020600020905b8154815290600101906020018083116111d157829003601f168201915b5050505050612088565b945061127c600660008b81526020019081526020016000206005018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111ee5780601f106111c3576101008083540402835291602001916111ee565b935061128788612088565b9250826020015185602001511461129d57600080fd5b82604001518560400151146112b157600080fd5b8260c0015160008781526005602052604090205490925090506112dc828a838a63ffffffff61218c16565b15156112e757600080fd5b8260a00151600160a060020a03168560600151600160a060020a0316148015611317575082608001518560800151105b156113685760008981526006602052604081208181556001810182905560028101829055906113496003830182613757565b60048201600090556005820160006113619190613757565b505061146a565b600089815260066020526040902060020154861080156113b357508260a00151600160a060020a03168460600151600160a060020a03161480156113b3575083608001518360800151115b156113e55760008981526006602052604081208181556001810182905560028101829055906113496003830182613757565b60008981526006602052604090206004015486101561141b5760008981526006602052604090206001905561141b898988612203565b60008981526006602052604090205460011461143657600080fd5b7fd3103af5a62a3d64dc3b30e901e2aa37bd30bde8b6f61a70e79b23e20d3bbd468960405190815260200160405180910390a15b505050505050505050565b600454600160a060020a031681565b60025481565b61149261368b565b6000878152600a602052604081205481908190158015906114d4575060035460008b8152600a602052604090205442916114d2919063ffffffff61217916565b115b15156114df57600080fd5b6114ea8b8b8961102e565b156114f457600080fd5b6114fd87612088565b93508360c00151600086815260056020526040902054909350915087905061152d838c848963ffffffff61218c16565b151561153857600080fd5b61154a818c8c8c63ffffffff61218c16565b151561155557600080fd5b836080015188111561156d5761156d8b8b89886123e1565b7f5785fb6481777a31c76595b8b701cf63e247203b0e63c12dd7da5d731d3769618b8b60405191825260208201526040908101905180910390a15050505050505050505050565b6115bc61368b565b60006115c989898661102e565b15156115d457600080fd5b6115dd84612088565b9150829050608082015183116115f257600080fd5b611604818a888863ffffffff61218c16565b151561160f57600080fd5b6003546000878152600a60205260409020544291611633919063ffffffff61217916565b1061163d57600080fd5b6000888152600a6020526040808220548883529120541061165d57600080fd5b6116688989866125f1565b7f8e7f7ccff34f7d0789427738f22785ddd8a9fee340613ffc9b8b24ad0e2a20f189898860405192835260208301919091526040808301919091526060909101905180910390a1505050505050505050565b60008060096000858152602001908152602001600020600201836040518082805190602001908083835b602083106117035780518252601f1990920191602091820191016116e4565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050806000141561174b576000915061176e565b600084815260096020908152604080832084845260010190915290205460ff1691505b5092915050565b60005433600160a060020a0390811691161461179057600080fd5b600054600160a060020a03167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b60008054819033600160a060020a0390811691161461180557600080fd5b611812858560015461291f565b600081815260076020908152604080832087905560025460089092529091205560018054919250611849919063ffffffff61217916565b6001557f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a158584836040518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390a1949350505050565b6118b561368b565b6118bd61368b565b6000806118ca89896116ba565b15156118d557600080fd5b6000898152600660205260409020546001146118f057600080fd5b6118f988612088565b935061190487612088565b9250826020015184602001511461191a57600080fd5b826040015184604001511461192e57600080fd5b8260a00151600160a060020a03168460600151600160a060020a03161461195457600080fd5b826080015161196f600186608001519063ffffffff61217916565b1461197957600080fd5b600089815260066020526040902060040154851061199657600080fd5b8260c0015160008681526005602052604090205490925090506119c1828a838963ffffffff61218c16565b15156119cc57600080fd5b6119d68989612969565b6119df89611db9565b15156119f8576000898152600660205260409020600290555b7f755b2676ab5f5b54bffac288782b3b18ae132ffd1950e416c1cfeb98eeb3c5c28960405190815260200160405180910390a1505050505050505050565b600054600160a060020a031681565b60045433600160a060020a03908116911614611a6057600080fd5b600254611a7490600163ffffffff61217916565b6002819055600090815260056020526040908190208290557f5f11b60a71ba7b4124fe41971a682a44d1af8fff92e0c4852a2701e56323218a9082905190815260200160405180910390a150565b611aca613745565b6000848152600b60209081526040808320868452909152812081906001908101908290611afe90879063ffffffff61217916565b81526020019081526020016000209050806001018160020154818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611bac5780601f10611b8157610100808354040283529160200191611bac565b820191906000526020600020905b815481529060010190602001808311611b8f57829003601f168201915b505050505091509250925050935093915050565b6000828152600b60209081526040808320848452909152812054801515611bea576000915061176e565b611bfb81600163ffffffff612c3f16565b949350505050565b60045433600160a060020a03908116911614611c1e57600080fd5b6000818152600a602052604090205415611c3757600080fd5b6000818152600a6020526040908190204290557f3dfae83a0b2f3013f409fd97c7e72574fcb10cd81987893771d8a2707d533d219082905190815260200160405180910390a150565b611c8861368b565b611c9061368b565b600080611c9e8a8a8a61102e565b1515611ca957600080fd5b611cb288612088565b9350611cbd87612088565b92508260200151846020015114611cd357600080fd5b8260400151846040015114611ce757600080fd5b8260a00151600160a060020a03168460600151600160a060020a031614611d0d57600080fd5b8260800151611d28600186608001519063ffffffff61217916565b14611d3257600080fd5b8260c001516000868152600560205260409020549092509050611d5d828b838963ffffffff61218c16565b1515611d6857600080fd5b611d738a8a8a6125f1565b7f12e9fe9a5ae32610f4f992917e433e03162fb143bc4614d245ff7f30482866b08a8a60405191825260208201526040908101905180910390a150505050505050505050565b600081815260096020526040812054801515611dd85760009150611dec565b611de981600163ffffffff612c3f16565b91505b50919050565b600a6020526000908152604090205481565b60005433600160a060020a03908116911614611e1f57600080fd5b611e2881612c51565b50565b6000611e3561368b565b611e3d61368b565b6000805481908190819033600160a060020a03908116911614611e5f57600080fd5b611e688d612088565b9550611e738a612088565b945084518b14611e8257600080fd5b8460200151866020015114611e9657600080fd5b8460400151866040015114611eaa57600080fd5b8460a00151600160a060020a03168660600151600160a060020a031614611ed057600080fd5b8460600151600160a060020a038f8116911614611eec57600080fd5b8560c0015160008c815260056020526040902054909450925060c0850151915060056000898152602001908152602001600020549050611f3986602001518590858f63ffffffff61218c16565b1515611f4457600080fd5b611f5b85602001518390838c63ffffffff61218c16565b1515611f6657600080fd5b42600660008760200151815260200190815260200160002060010154101515611f8e57600080fd5b6006600086602001518152602081019190915260400160002054600214611fb457600080fd5b611fc18560200151611db9565b15611fcb57600080fd5b6003600660008760200151815260200190815260200160002060000181905550600760008660200151815260208082019290925260400160009081208190556008918701518152602001908152602001600020600090557f7c59798283502bd302b18828d4f858808c79d023b3784676d6984b3c23ae45b5856020015160405190815260200160405180910390a184602001519e9d5050505050505050505050505050565b60035481565b60056020526000908152604090205481565b61209061368b565b612098613745565b6120b260066120a685612cd1565b9063ffffffff612d0316565b905060e0604051908101604052806120df836000815181106120d057fe5b90602001906020020151612d9f565b81526020016120f4836001815181106120d057fe5b8152602001612109836002815181106120d057fe5b815260200161212d8360038151811061211e57fe5b90602001906020020151612dc6565b600160a060020a0316815260200161214b836004815181106120d057fe5b815260200161215983612dff565b600160a060020a0316815260200161217083612f4b565b90529392505050565b8181018281101561218657fe5b92915050565b6000808560205b845181116121f5578085015192506002870615156121cb578183604051918252602082015260409081019051809103902091506121e7565b8282604051918252602082015260409081019051809103902091505b600287049650602001612193565b509390931495945050505050565b600061220d61379b565b60008581526009602052604090819020600201908590518082805190602001908083835b602083106122505780518252601f199092019160209182019101612231565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020549150811561229257600080fd5b6060604051908101604090815260018252602080830187905281830186905260008881526009909152205490915015156122d9576000858152600960205260409020600190555b600085815260096020908152604080832080548452600101909152902081908151815460ff19169015151781556020820151816001019080516123209291602001906136c7565b5060408201516002918201556000878152600960205260409081902080549350909101908690518082805190602001908083835b602083106123735780518252601f199092019160209182019101612354565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020556000858152600960205260409020546123c890600163ffffffff61217916565b6000958652600960205260409095209490945550505050565b60006123eb61379b565b6000868152600b60209081526040808320888452909152808220600201908690518082805190602001908083835b602083106124385780518252601f199092019160209182019101612419565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020549250821561247a57600080fd5b6060604051908101604090815260018252602080830188905281830187905260008a8152600b82528281208a8252909152205490925015156124d4576000878152600b602090815260408083208984529091529020600190555b506000868152600b6020908152604080832088845282528083208054808552600190910190925290912082908151815460ff19169015151781556020820151816001019080516125289291602001906136c7565b5060408201516002918201556000898152600b602090815260408083208b845290915290819020849350909101908790518082805190602001908083835b602083106125855780518252601f199092019160209182019101612566565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020556125cb81600163ffffffff61217916565b6000978852600b60209081526040808a20988a5297905295909620949094555050505050565b6000838152600b6020908152604080832085845290915280822082918291600201908590518082805190602001908083835b602083106126425780518252601f199092019160209182019101612623565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902054925082151561268557600080fd5b6000868152600b60209081526040808320888452825280832086845260019081019092528220805460ff1916815591906126c190830182613757565b5060006002918201819055878152600b6020908152604080832089845290915290819020909101908590518082805190602001908083835b602083106127185780518252601f1990920191602091820191016126f9565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020600090819055868152600b6020908152604080832088845290915290205461277b90600163ffffffff612c3f16565b91508282146128d257506000858152600b6020908152604080832087845282528083208484526001908101808452828520878652935292208154815460ff909116151560ff199091161781558183018054929384936127ef92848301929091600291811615610100026000190116046137c3565b50600291820154908201556000878152600b60209081526040808320898452909152908190208592019060018401905180828054600181600116156101000203166002900480156128775780601f10612855576101008083540402835291820191612877565b820191906000526020600020905b815481529060010190602001808311612863575b50509283525050602001604051908190039020556000868152600b60209081526040808320888452825280832085845260019081019092528220805460ff1916815591906128c790830182613757565b600282016000905550505b81600114156128fa576000868152600b60209081526040808320888452909152812055612917565b6000868152600b6020908152604080832088845290915290208290555b505050505050565b60008284836040516c01000000000000000000000000600160a060020a03948516810282529290931690910260148301526028820152604801604051809103902090509392505050565b60008281526009602052604080822082918291600201908590518082805190602001908083835b602083106129af5780518252601f199092019160209182019101612990565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205492508215156129f257600080fd5b600085815260096020908152604080832086845260019081019092528220805460ff191681559190612a2690830182613757565b6002820160009055505060096000868152602001908152602001600020600201846040518082805190602001908083835b60208310612a765780518252601f199092019160209182019101612a57565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902060009081905585815260096020526040902054612ace90600163ffffffff612c3f16565b9150828214612c0957506000848152600960209081526040808320848452600190810190925280832085845292208254815460ff191660ff909116151517815582820180548493612b39928482019290916002610100918316159190910260001901909116046137c3565b506002918201549082015560008681526009602052604090819020859201906001840190518082805460018160011615610100020316600290048015612bb65780601f10612b94576101008083540402835291820191612bb6565b820191906000526020600020905b815481529060010190602001808311612ba2575b5050928352505060200160405190819003902055600085815260096020908152604080832085845260019081019092528220805460ff191681559190612bfe90830182613757565b600282016000905550505b8160011415612c2657600085815260096020526040812055612c38565b60008581526009602052604090208290555b5050505050565b600082821115612c4b57fe5b50900390565b600160a060020a0381161515612c6657600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b612cd9613838565b600080835191505060208301604080519081016040528181526020810183905292505b5050919050565b612d0b613745565b612d1361384f565b600083604051805910612d235750595b908082528060200260200182016040528015612d5957816020015b612d46613838565b815260200190600190039081612d3e5790505b509250612d6585613057565b91505b83811015612d9757612d798261307c565b838281518110612d8557fe5b60209081029091010152600101612d68565b505092915050565b6000806000612dad846130b1565b909250905060208190036101000a825104949350505050565b6000806000612dd4846130b1565b909250905060148114612de657600080fd5b6c01000000000000000000000000825104949350505050565b6000612e09613745565b6000612e13613745565b612e1b613745565b6005604051805910612e2a5750595b908082528060200260200182016040528015612e6057816020015b612e4d613745565b815260200190600190039081612e455790505b509350600092505b6005831015612eb057612e8f868481518110612e8057fe5b90602001906020020151613118565b848481518110612e9b57fe5b60209081029091010152600190920191612e68565b612eb98461316c565b9150612eda86600581518110612ecb57fe5b90602001906020020151613198565b9050612f41826040518082805190602001908083835b60208310612f0f5780518252601f199092019160209182019101612ef0565b6001836020036101000a03801982511681845116179092525050509190910192506040915050518091039020826131e8565b9695505050505050565b6000612f55613745565b6000612f5f613745565b6005604051805910612f6e5750595b908082528060200260200182016040528015612fa457816020015b612f91613745565b815260200190600190039081612f895790505b509250600091505b6005821015612fe557612fc4858381518110612e8057fe5b838381518110612fd057fe5b60209081029091010152600190910190612fac565b612fee8361316c565b9050806040518082805190602001908083835b602083106130205780518252601f199092019160209182019101613001565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051809103902093505b505050919050565b61305f61384f565b600061306a836132c8565b83519383529092016020820152919050565b613084613838565b6000808360200151915061309782613347565b828452602080850182905292019390910192909252919050565b60008080808085519150815160001a925060808310156130d75781945060019350613110565b60b88310156130f55760018660200151039350816001019450613110565b5060b619820180600160208801510303935080820160010194505b505050915091565b613120613745565b6000808360200151915081151561313657612cfc565b816040518059106131445750595b818152601f19601f8301168101602001604052905092505060208201612cfc818551846133d9565b613174613745565b61317c613745565b613184613745565b61318d8461341e565b9150611bfb82613502565b6131a0613745565b60008060006131ae856130b1565b9093509150816040518059106131c15750595b818152601f19601f830116810160200160405290509350506020830161304f8184846133d9565b600080600080845160411461320057600093506132bf565b6020850151925060408501519150606085015160001a9050601b8160ff16101561322857601b015b8060ff16601b1415801561324057508060ff16601c14155b1561324e57600093506132bf565b6001868285856040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f115156132b357600080fd5b50506020604051035193505b50505092915050565b6000806000836020015115156132e15760009250612cfc565b83519050805160001a915060808210156132fe5760009250612cfc565b60b8821080613319575060c08210158015613319575060f882105b156133275760019250612cfc565b60c082101561333c5760b51982019250612cfc565b5060f5190192915050565b600080825160001a905060808110156133635760019150611dec565b60b881101561337857607e1981019150611dec565b60c08110156133a25760b78103806020036101000a60018501510480820160010193505050611dec565b60f88110156133b75760be1981019150611dec565b60f78103806020036101000a6001850151048082016001019350505050919050565b60005b602082106133ff57825184526020840193506020830192506020820391506133dc565b6001826020036101000a03905080198351168185511617909352505050565b613426613745565b600080613431613745565b600061343b613745565b60008094505b87518510156134705787858151811061345657fe5b906020019060200201515190950194600190940193613441565b8560405180591061347e5750595b8181526020601f909201601f191681018201604052600096509450840192505b87518510156134f6578785815181106134b357fe5b90602001906020020151915050602081016134d0838284516133d9565b8785815181106134dc57fe5b90602001906020020151516001909501949092019161349e565b50919695505050505050565b61350a613745565b6000613514613745565b60008060008060208801955087519250600090505b82816101000210156135415760019182019101613529565b603783116135b6578260010160405180591061355a5750595b818152601f19601f8301168101602001604052905094508260c00160f860020a028560008151811061358857fe5b906020010190600160f860020a031916908160001a9053506021850193506135b18487856133d9565b61367f565b8282600101016040518059106135c95750595b818152601f19601f83011681016020016040529050945060f860020a60f7830102856000815181106135f757fe5b906020010190600160f860020a031916908160001a905350600190505b81811161366c576101008183036101000a8481151561362f57fe5b0481151561363957fe5b0660f860020a0285828151811061364c57fe5b906020010190600160f860020a031916908160001a905350600101613614565b816021860101935061367f8487856133d9565b50929695505050505050565b60e06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a0820181905260c082015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061370857805160ff1916838001178555613735565b82800160010185558215613735579182015b8281111561373557825182559160200191906001019061371a565b50613741929150613863565b5090565b60206040519081016040526000815290565b50805460018160011615610100020316600290046000825580601f1061377d5750611e28565b601f016020900490600052602060002090810190611e289190613863565b606060405190810160405260008152602081016137b6613745565b8152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106137fc5780548555613735565b8280016001018555821561373557600052602060002091601f016020900482015b8281111561373557825482559160010191906001019061381d565b604080519081016040526000808252602082015290565b6060604051908101604052806137b6613838565b61387d91905b808211156137415760008155600101613869565b905600a165627a7a72305820a35edf839bcec9a70390cd4b460f5568d87ada479069ecddef7bd52779d8947d0029`

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
