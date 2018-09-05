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
const RootChainABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"previousTx\",\"type\":\"bytes\"},{\"name\":\"previousTxProof\",\"type\":\"bytes\"},{\"name\":\"previousTxBlockNum\",\"type\":\"uint256\"},{\"name\":\"lastTx\",\"type\":\"bytes\"},{\"name\":\"lastTxProof\",\"type\":\"bytes\"},{\"name\":\"lastTxBlockNum\",\"type\":\"uint256\"}],\"name\":\"startExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChallenge\",\"outputs\":[{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"challengeBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exits\",\"outputs\":[{\"name\":\"state\",\"type\":\"uint256\"},{\"name\":\"exitTime\",\"type\":\"uint256\"},{\"name\":\"exitTxBlkNum\",\"type\":\"uint256\"},{\"name\":\"exitTx\",\"type\":\"bytes\"},{\"name\":\"txBeforeExitTxBlkNum\",\"type\":\"uint256\"},{\"name\":\"txBeforeExitTx\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpoint\",\"type\":\"bytes32\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"}],\"name\":\"checkpointIsChallenge\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"challengeBlockNum\",\"type\":\"uint256\"}],\"name\":\"challengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpointRoot\",\"type\":\"bytes32\"},{\"name\":\"checkpointProof\",\"type\":\"bytes\"},{\"name\":\"wrongNonce\",\"type\":\"uint256\"},{\"name\":\"lastTx\",\"type\":\"bytes\"},{\"name\":\"lastTxProof\",\"type\":\"bytes\"},{\"name\":\"lastTxBlockNum\",\"type\":\"uint256\"}],\"name\":\"challengeCheckpoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpointRoot\",\"type\":\"bytes32\"},{\"name\":\"checkpointProof\",\"type\":\"bytes\"},{\"name\":\"historicalCheckpointRoot\",\"type\":\"bytes32\"},{\"name\":\"historicalCheckpointProof\",\"type\":\"bytes\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"moreNonce\",\"type\":\"uint256\"}],\"name\":\"respondWithHistoricalCheckpoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"}],\"name\":\"challengeExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"currency\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"respondTx\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"respondChallengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"newBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpoint\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getCheckpointChallenge\",\"outputs\":[{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"challengeBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpoint\",\"type\":\"bytes32\"}],\"name\":\"checkpointChallengesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"newCheckpoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"checkpointRoot\",\"type\":\"bytes32\"},{\"name\":\"challengeTx\",\"type\":\"bytes\"},{\"name\":\"respondTx\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"respondCheckpointChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"challengesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"previousTx\",\"type\":\"bytes\"},{\"name\":\"previousTxProof\",\"type\":\"bytes\"},{\"name\":\"previousTxBlockNum\",\"type\":\"uint256\"},{\"name\":\"lastTx\",\"type\":\"bytes\"},{\"name\":\"lastTxProof\",\"type\":\"bytes\"},{\"name\":\"lastTxBlockNum\",\"type\":\"uint256\"}],\"name\":\"finishExit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"NewBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"NewCheckpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastBlock\",\"type\":\"uint256\"}],\"name\":\"StartExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"FinishExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"ChallengeExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"checkpoint\",\"type\":\"bytes32\"}],\"name\":\"ChallengeCheckpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"}],\"name\":\"RespondChallengeExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"checkpoint\",\"type\":\"bytes32\"}],\"name\":\"RespondCheckpointChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"checkpoint\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"historicalCheckpoint\",\"type\":\"bytes32\"}],\"name\":\"RespondWithHistoricalCheckpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `0x6060604052341561000f57600080fd5b6040516020806138bf8339810160405280805160008054600160a060020a03338116600160a060020a031992831617835560028390556212750060035560019290925560048054929093169116179055505061384f806100706000396000f30060606040526004361061015e5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166304675c6581146101635780632dfdf0b5146102875780632e8f396d146102ac578063342de17914610344578063370beab91461047c578063449b2f44146104ec5780634a2ee19814610502578063570ca7351461059c57806357e871e7146105cb5780635e8cdf44146105de5780636d2975c7146106c95780636e606251146107b0578063715018a6146108065780638340f549146108195780638b593183146108415780638da5cb5b1461091d578063981adca5146109305780639af35551146109465780639ced689014610962578063a5e450f41461097b578063ad131e1314610991578063b15e42f814610a73578063eb5e91ff14610a89578063f2fde38b14610a9f578063f310f2b214610abe578063f3f480d914610bee578063f95643b114610c01575b600080fd5b341561016e57600080fd5b61028560046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505093359350610c1792505050565b005b341561029257600080fd5b61029a610ed2565b60405190815260200160405180910390f35b34156102b757600080fd5b6102c5600435602435610ed8565b6040518080602001838152602001828103825284818151815260200191508051906020019080838360005b838110156103085780820151838201526020016102f0565b50505050905090810190601f1680156103355780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b341561034f57600080fd5b61035a600435610fca565b60405186815260208101869052604081018590526080810183905260c0606082018181528554600260001961010060018416150201909116049183018290529060a083019060e0840190879080156103f35780601f106103c8576101008083540402835291602001916103f3565b820191906000526020600020905b8154815290600101906020018083116103d657829003601f168201915b50508381038252845460026000196101006001841615020190911604808252602090910190859080156104675780601f1061043c57610100808354040283529160200191610467565b820191906000526020600020905b81548152906001019060200180831161044a57829003601f168201915b50509850505050505050505060405180910390f35b341561048757600080fd5b6104d8600480359060248035919060649060443590810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610ffb95505050505050565b604051901515815260200160405180910390f35b34156104f757600080fd5b61029a6004356110c7565b341561050d57600080fd5b610285600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965050933593506110d992505050565b34156105a757600080fd5b6105af611442565b604051600160a060020a03909116815260200160405180910390f35b34156105d657600080fd5b61029a611451565b34156105e957600080fd5b610285600480359060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061145792505050565b34156106d457600080fd5b610285600480359060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094968635969095506040808201955060209182013587018083019550359350839250601f830182900482029091019051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061158192505050565b34156107bb57600080fd5b6104d8600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061168795505050505050565b341561081157600080fd5b610285611742565b341561082457600080fd5b61029a600160a060020a03600435811690602435166044356117b4565b341561084c57600080fd5b610285600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061186b92505050565b341561092857600080fd5b6105af6119f4565b341561093b57600080fd5b610285600435611a03565b341561095157600080fd5b6102c5600435602435604435611a80565b341561096d57600080fd5b61029a600435602435611b7e565b341561098657600080fd5b610285600435611bc1565b341561099c57600080fd5b610285600480359060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505093359350611c3e92505050565b3415610a7e57600080fd5b61029a600435611d77565b3415610a9457600080fd5b61029a600435611db0565b3415610aaa57600080fd5b610285600160a060020a0360043516611dc2565b3415610ac957600080fd5b61029a60048035600160a060020a03169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505093359350611de992505050565b3415610bf957600080fd5b61029a612013565b3415610c0c57600080fd5b61029a600435612019565b610c1f61362e565b610c2761362e565b600080600080610c368c61202b565b9550610c418961202b565b945084518a14610c5057600080fd5b8460200151866020015114610c6457600080fd5b8460400151866040015114610c7857600080fd5b8460a00151600160a060020a03168660600151600160a060020a031614610c9e57600080fd5b610cb4600187608001519063ffffffff61211c16565b856080015114610cc357600080fd5b8460600151600160a060020a031633600160a060020a0316141515610ce757600080fd5b60076000866020015181526020810191909152604001600020541515610d0c57600080fd5b8560c0015160008b815260056020526040902054909450925060c0850151915060056000888152602001908152602001600020549050610d5986602001518590858e63ffffffff61212f16565b1515610d6457600080fd5b610d7b85602001518390838b63ffffffff61212f16565b1515610d8657600080fd5b600660008660200151815260208101919091526040016000205415610daa57600080fd5b610db78560200151611d77565b15610dc157600080fd5b60c06040519081016040528060028152602001610de96003544261211c90919063ffffffff16565b81526020018881526020018a81526020018b81526020018d8152506006600087602001518152602001908152602001600020600082015181556020820151816001015560408201518160020155606082015181600301908051610e5092916020019061366a565b506080820151816004015560a082015181600501908051610e7592916020019061366a565b509050507f4d3db44958203fdb34cb08c6f6ecda2c9e182cd426b0871efec5ade7ac94580d86602001518b8960405180848152602001838152602001828152602001935050505060405180910390a1505050505050505050505050565b60015481565b610ee06136e8565b600083815260086020526040812081906001908101908290610f0990879063ffffffff61211c16565b81526020019081526020016000209050806001018160020154818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fb75780601f10610f8c57610100808354040283529160200191610fb7565b820191906000526020600020905b815481529060010190602001808311610f9a57829003601f168201915b5050505050915092509250509250929050565b6006602052600090815260409020805460018201546002830154600484015492939192909160038101919060050186565b6000838152600a6020908152604080832085845290915280822082916002909101908490518082805190602001908083835b6020831061104c5780518252601f19909201916020918201910161102d565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050806000141561109457600091506110bf565b6000858152600a60209081526040808320878452825280832084845260010190915290205460ff1691505b509392505050565b60076020526000908152604090205481565b6110e161362e565b6110e961362e565b6110f161362e565b600087815260066020526040812054819060021461110e57600080fd5b6111c5600660008b81526020019081526020016000206003018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111bb5780601f10611190576101008083540402835291602001916111bb565b820191906000526020600020905b81548152906001019060200180831161119e57829003601f168201915b505050505061202b565b9450611249600660008b81526020019081526020016000206005018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111bb5780601f10611190576101008083540402835291602001916111bb565b93506112548861202b565b9250826020015185602001511461126a57600080fd5b826040015185604001511461127e57600080fd5b8260c0015160008781526005602052604090205490925090506112a9828a838a63ffffffff61212f16565b15156112b457600080fd5b8260a00151600160a060020a03168560600151600160a060020a03161480156112e4575082608001518560800151105b1561133557600089815260066020526040812081815560018101829055600281018290559061131660038301826136fa565b600482016000905560058201600061132e91906136fa565b5050611437565b6000898152600660205260409020600201548610801561138057508260a00151600160a060020a03168460600151600160a060020a0316148015611380575083608001518360800151115b156113b257600089815260066020526040812081815560018101829055600281018290559061131660038301826136fa565b6000898152600660205260409020600401548610156113e8576000898152600660205260409020600190556113e88989886121a6565b60008981526006602052604090205460011461140357600080fd5b7fd3103af5a62a3d64dc3b30e901e2aa37bd30bde8b6f61a70e79b23e20d3bbd468960405190815260200160405180910390a15b505050505050505050565b600454600160a060020a031681565b60025481565b61145f61362e565b60008781526009602052604081205481908190158015906114a1575060035460008b815260096020526040902054429161149f919063ffffffff61211c16565b115b15156114ac57600080fd5b6114b78b8b89610ffb565b156114c157600080fd5b6114ca8761202b565b93508360c0015160008681526005602052604090205490935091508790506114fa838c848963ffffffff61212f16565b151561150557600080fd5b611517818c8c8c63ffffffff61212f16565b151561152257600080fd5b836080015188111561153a5761153a8b8b8988612384565b7f5785fb6481777a31c76595b8b701cf63e247203b0e63c12dd7da5d731d3769618b8b60405191825260208201526040908101905180910390a15050505050505050505050565b61158961362e565b6000611596898986610ffb565b15156115a157600080fd5b6115aa8461202b565b9150829050608082015183116115bf57600080fd5b6115d1818a888863ffffffff61212f16565b15156115dc57600080fd5b6003546000878152600960205260409020544291611600919063ffffffff61211c16565b1061160a57600080fd5b600088815260096020526040808220548883529120541061162a57600080fd5b611635898986612594565b7f8e7f7ccff34f7d0789427738f22785ddd8a9fee340613ffc9b8b24ad0e2a20f189898860405192835260208301919091526040808301919091526060909101905180910390a1505050505050505050565b60008060086000858152602001908152602001600020600201836040518082805190602001908083835b602083106116d05780518252601f1990920191602091820191016116b1565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390205490508060001415611718576000915061173b565b600084815260086020908152604080832084845260010190915290205460ff1691505b5092915050565b60005433600160a060020a0390811691161461175d57600080fd5b600054600160a060020a03167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b60008054819033600160a060020a039081169116146117d257600080fd5b6117df85856001546128c2565b600081815260076020526040902084905560018054919250611807919063ffffffff61211c16565b6001557f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a158584836040518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390a1949350505050565b61187361362e565b61187b61362e565b6000806118888989611687565b151561189357600080fd5b6000898152600660205260409020546001146118ae57600080fd5b6118b78861202b565b93506118c28761202b565b925082602001518460200151146118d857600080fd5b82604001518460400151146118ec57600080fd5b8260a00151600160a060020a03168460600151600160a060020a03161461191257600080fd5b826080015161192d600186608001519063ffffffff61211c16565b1461193757600080fd5b600089815260066020526040902060040154851061195457600080fd5b8260c00151600086815260056020526040902054909250905061197f828a838963ffffffff61212f16565b151561198a57600080fd5b611994898961290c565b61199d89611d77565b15156119b6576000898152600660205260409020600290555b7f755b2676ab5f5b54bffac288782b3b18ae132ffd1950e416c1cfeb98eeb3c5c28960405190815260200160405180910390a1505050505050505050565b600054600160a060020a031681565b60045433600160a060020a03908116911614611a1e57600080fd5b600254611a3290600163ffffffff61211c16565b6002819055600090815260056020526040908190208290557f5f11b60a71ba7b4124fe41971a682a44d1af8fff92e0c4852a2701e56323218a9082905190815260200160405180910390a150565b611a886136e8565b6000848152600a60209081526040808320868452909152812081906001908101908290611abc90879063ffffffff61211c16565b81526020019081526020016000209050806001018160020154818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611b6a5780601f10611b3f57610100808354040283529160200191611b6a565b820191906000526020600020905b815481529060010190602001808311611b4d57829003601f168201915b505050505091509250925050935093915050565b6000828152600a60209081526040808320848452909152812054801515611ba8576000915061173b565b611bb981600163ffffffff612be216565b949350505050565b60045433600160a060020a03908116911614611bdc57600080fd5b60008181526009602052604090205415611bf557600080fd5b600081815260096020526040908190204290557f3dfae83a0b2f3013f409fd97c7e72574fcb10cd81987893771d8a2707d533d219082905190815260200160405180910390a150565b611c4661362e565b611c4e61362e565b600080611c5c8a8a8a610ffb565b1515611c6757600080fd5b611c708861202b565b9350611c7b8761202b565b92508260200151846020015114611c9157600080fd5b8260400151846040015114611ca557600080fd5b8260a00151600160a060020a03168460600151600160a060020a031614611ccb57600080fd5b8260800151611ce6600186608001519063ffffffff61211c16565b14611cf057600080fd5b8260c001516000868152600560205260409020549092509050611d1b828b838963ffffffff61212f16565b1515611d2657600080fd5b611d318a8a8a612594565b7f12e9fe9a5ae32610f4f992917e433e03162fb143bc4614d245ff7f30482866b08a8a60405191825260208201526040908101905180910390a150505050505050505050565b600081815260086020526040812054801515611d965760009150611daa565b611da781600163ffffffff612be216565b91505b50919050565b60096020526000908152604090205481565b60005433600160a060020a03908116911614611ddd57600080fd5b611de681612bf4565b50565b6000611df361362e565b611dfb61362e565b6000805481908190819033600160a060020a03908116911614611e1d57600080fd5b611e268d61202b565b9550611e318a61202b565b945084518b14611e4057600080fd5b8460200151866020015114611e5457600080fd5b8460400151866040015114611e6857600080fd5b8460a00151600160a060020a03168660600151600160a060020a031614611e8e57600080fd5b8460600151600160a060020a038f8116911614611eaa57600080fd5b8560c0015160008c815260056020526040902054909450925060c0850151915060056000898152602001908152602001600020549050611ef786602001518590858f63ffffffff61212f16565b1515611f0257600080fd5b611f1985602001518390838c63ffffffff61212f16565b1515611f2457600080fd5b42600660008760200151815260200190815260200160002060010154101515611f4c57600080fd5b6006600086602001518152602081019190915260400160002054600214611f7257600080fd5b611f7f8560200151611d77565b15611f8957600080fd5b600360066000876020015181526020019081526020016000206000018190555060076000866020015181526020808201929092526040016000908120557f7c59798283502bd302b18828d4f858808c79d023b3784676d6984b3c23ae45b59086015160405190815260200160405180910390a184602001519e9d5050505050505050505050505050565b60035481565b60056020526000908152604090205481565b61203361362e565b61203b6136e8565b612055600661204985612c74565b9063ffffffff612ca616565b905060e0604051908101604052806120828360008151811061207357fe5b90602001906020020151612d42565b81526020016120978360018151811061207357fe5b81526020016120ac8360028151811061207357fe5b81526020016120d0836003815181106120c157fe5b90602001906020020151612d69565b600160a060020a031681526020016120ee8360048151811061207357fe5b81526020016120fc83612da2565b600160a060020a0316815260200161211383612eee565b90529392505050565b8181018281101561212957fe5b92915050565b6000808560205b845181116121985780850151925060028706151561216e5781836040519182526020820152604090810190518091039020915061218a565b8282604051918252602082015260409081019051809103902091505b600287049650602001612136565b509390931495945050505050565b60006121b061373e565b60008581526008602052604090819020600201908590518082805190602001908083835b602083106121f35780518252601f1990920191602091820191016121d4565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020549150811561223557600080fd5b60606040519081016040908152600182526020808301879052818301869052600088815260089091522054909150151561227c576000858152600860205260409020600190555b600085815260086020908152604080832080548452600101909152902081908151815460ff19169015151781556020820151816001019080516122c392916020019061366a565b5060408201516002918201556000878152600860205260409081902080549350909101908690518082805190602001908083835b602083106123165780518252601f1990920191602091820191016122f7565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205560008581526008602052604090205461236b90600163ffffffff61211c16565b6000958652600860205260409095209490945550505050565b600061238e61373e565b6000868152600a60209081526040808320888452909152808220600201908690518082805190602001908083835b602083106123db5780518252601f1990920191602091820191016123bc565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020549250821561241d57600080fd5b6060604051908101604090815260018252602080830188905281830187905260008a8152600a82528281208a825290915220549092501515612477576000878152600a602090815260408083208984529091529020600190555b506000868152600a6020908152604080832088845282528083208054808552600190910190925290912082908151815460ff19169015151781556020820151816001019080516124cb92916020019061366a565b5060408201516002918201556000898152600a602090815260408083208b845290915290819020849350909101908790518082805190602001908083835b602083106125285780518252601f199092019160209182019101612509565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205561256e81600163ffffffff61211c16565b6000978852600a60209081526040808a20988a5297905295909620949094555050505050565b6000838152600a6020908152604080832085845290915280822082918291600201908590518082805190602001908083835b602083106125e55780518252601f1990920191602091820191016125c6565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902054925082151561262857600080fd5b6000868152600a60209081526040808320888452825280832086845260019081019092528220805460ff191681559190612664908301826136fa565b5060006002918201819055878152600a6020908152604080832089845290915290819020909101908590518082805190602001908083835b602083106126bb5780518252601f19909201916020918201910161269c565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020600090819055868152600a6020908152604080832088845290915290205461271e90600163ffffffff612be216565b915082821461287557506000858152600a6020908152604080832087845282528083208484526001908101808452828520878652935292208154815460ff909116151560ff199091161781558183018054929384936127929284830192909160029181161561010002600019011604613766565b50600291820154908201556000878152600a602090815260408083208984529091529081902085920190600184019051808280546001816001161561010002031660029004801561281a5780601f106127f857610100808354040283529182019161281a565b820191906000526020600020905b815481529060010190602001808311612806575b50509283525050602001604051908190039020556000868152600a60209081526040808320888452825280832085845260019081019092528220805460ff19168155919061286a908301826136fa565b600282016000905550505b816001141561289d576000868152600a602090815260408083208884529091528120556128ba565b6000868152600a6020908152604080832088845290915290208290555b505050505050565b60008284836040516c01000000000000000000000000600160a060020a03948516810282529290931690910260148301526028820152604801604051809103902090509392505050565b60008281526008602052604080822082918291600201908590518082805190602001908083835b602083106129525780518252601f199092019160209182019101612933565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902054925082151561299557600080fd5b600085815260086020908152604080832086845260019081019092528220805460ff1916815591906129c9908301826136fa565b6002820160009055505060086000868152602001908152602001600020600201846040518082805190602001908083835b60208310612a195780518252601f1990920191602091820191016129fa565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902060009081905585815260086020526040902054612a7190600163ffffffff612be216565b9150828214612bac57506000848152600860209081526040808320848452600190810190925280832085845292208254815460ff191660ff909116151517815582820180548493612adc92848201929091600261010091831615919091026000190190911604613766565b506002918201549082015560008681526008602052604090819020859201906001840190518082805460018160011615610100020316600290048015612b595780601f10612b37576101008083540402835291820191612b59565b820191906000526020600020905b815481529060010190602001808311612b45575b5050928352505060200160405190819003902055600085815260086020908152604080832085845260019081019092528220805460ff191681559190612ba1908301826136fa565b600282016000905550505b8160011415612bc957600085815260086020526040812055612bdb565b60008581526008602052604090208290555b5050505050565b600082821115612bee57fe5b50900390565b600160a060020a0381161515612c0957600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b612c7c6137db565b600080835191505060208301604080519081016040528181526020810183905292505b5050919050565b612cae6136e8565b612cb66137f2565b600083604051805910612cc65750595b908082528060200260200182016040528015612cfc57816020015b612ce96137db565b815260200190600190039081612ce15790505b509250612d0885612ffa565b91505b83811015612d3a57612d1c8261301f565b838281518110612d2857fe5b60209081029091010152600101612d0b565b505092915050565b6000806000612d5084613054565b909250905060208190036101000a825104949350505050565b6000806000612d7784613054565b909250905060148114612d8957600080fd5b6c01000000000000000000000000825104949350505050565b6000612dac6136e8565b6000612db66136e8565b612dbe6136e8565b6005604051805910612dcd5750595b908082528060200260200182016040528015612e0357816020015b612df06136e8565b815260200190600190039081612de85790505b509350600092505b6005831015612e5357612e32868481518110612e2357fe5b906020019060200201516130bb565b848481518110612e3e57fe5b60209081029091010152600190920191612e0b565b612e5c8461310f565b9150612e7d86600581518110612e6e57fe5b9060200190602002015161313b565b9050612ee4826040518082805190602001908083835b60208310612eb25780518252601f199092019160209182019101612e93565b6001836020036101000a038019825116818451161790925250505091909101925060409150505180910390208261318b565b9695505050505050565b6000612ef86136e8565b6000612f026136e8565b6005604051805910612f115750595b908082528060200260200182016040528015612f4757816020015b612f346136e8565b815260200190600190039081612f2c5790505b509250600091505b6005821015612f8857612f67858381518110612e2357fe5b838381518110612f7357fe5b60209081029091010152600190910190612f4f565b612f918361310f565b9050806040518082805190602001908083835b60208310612fc35780518252601f199092019160209182019101612fa4565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051809103902093505b505050919050565b6130026137f2565b600061300d8361326b565b83519383529092016020820152919050565b6130276137db565b6000808360200151915061303a826132ea565b828452602080850182905292019390910192909252919050565b60008080808085519150815160001a9250608083101561307a57819450600193506130b3565b60b883101561309857600186602001510393508160010194506130b3565b5060b619820180600160208801510303935080820160010194505b505050915091565b6130c36136e8565b600080836020015191508115156130d957612c9f565b816040518059106130e75750595b818152601f19601f8301168101602001604052905092505060208201612c9f8185518461337c565b6131176136e8565b61311f6136e8565b6131276136e8565b613130846133c1565b9150611bb9826134a5565b6131436136e8565b600080600061315185613054565b9093509150816040518059106131645750595b818152601f19601f8301168101602001604052905093505060208301612ff281848461337c565b60008060008084516041146131a35760009350613262565b6020850151925060408501519150606085015160001a9050601b8160ff1610156131cb57601b015b8060ff16601b141580156131e357508060ff16601c14155b156131f15760009350613262565b6001868285856040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f1151561325657600080fd5b50506020604051035193505b50505092915050565b6000806000836020015115156132845760009250612c9f565b83519050805160001a915060808210156132a15760009250612c9f565b60b88210806132bc575060c082101580156132bc575060f882105b156132ca5760019250612c9f565b60c08210156132df5760b51982019250612c9f565b5060f5190192915050565b600080825160001a905060808110156133065760019150611daa565b60b881101561331b57607e1981019150611daa565b60c08110156133455760b78103806020036101000a60018501510480820160010193505050611daa565b60f881101561335a5760be1981019150611daa565b60f78103806020036101000a6001850151048082016001019350505050919050565b60005b602082106133a2578251845260208401935060208301925060208203915061337f565b6001826020036101000a03905080198351168185511617909352505050565b6133c96136e8565b6000806133d46136e8565b60006133de6136e8565b60008094505b8751851015613413578785815181106133f957fe5b9060200190602002015151909501946001909401936133e4565b856040518059106134215750595b8181526020601f909201601f191681018201604052600096509450840192505b87518510156134995787858151811061345657fe5b90602001906020020151915050602081016134738382845161337c565b87858151811061347f57fe5b906020019060200201515160019095019490920191613441565b50919695505050505050565b6134ad6136e8565b60006134b76136e8565b60008060008060208801955087519250600090505b82816101000210156134e457600191820191016134cc565b6037831161355957826001016040518059106134fd5750595b818152601f19601f8301168101602001604052905094508260c00160f860020a028560008151811061352b57fe5b906020010190600160f860020a031916908160001a90535060218501935061355484878561337c565b613622565b82826001010160405180591061356c5750595b818152601f19601f83011681016020016040529050945060f860020a60f78301028560008151811061359a57fe5b906020010190600160f860020a031916908160001a905350600190505b81811161360f576101008183036101000a848115156135d257fe5b048115156135dc57fe5b0660f860020a028582815181106135ef57fe5b906020010190600160f860020a031916908160001a9053506001016135b7565b816021860101935061362284878561337c565b50929695505050505050565b60e06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a0820181905260c082015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106136ab57805160ff19168380011785556136d8565b828001600101855582156136d8579182015b828111156136d85782518255916020019190600101906136bd565b506136e4929150613806565b5090565b60206040519081016040526000815290565b50805460018160011615610100020316600290046000825580601f106137205750611de6565b601f016020900490600052602060002090810190611de69190613806565b606060405190810160405260008152602081016137596136e8565b8152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061379f57805485556136d8565b828001600101855582156136d857600052602060002091601f016020900482015b828111156136d85782548255916001019190600101906137c0565b604080519081016040526000808252602082015290565b6060604051908101604052806137596137db565b61382091905b808211156136e4576000815560010161380c565b905600a165627a7a72305820f4ac83951efd24637a721fb1e2c0a7db05b03099ffe01b33676b04078322fd250029`

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
