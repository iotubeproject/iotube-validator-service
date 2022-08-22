// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AssetRegistryABI is the input ABI used to generate the binding from.
const AssetRegistryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AssetActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AssetDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tubeID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"AssetOnTubeActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tubeID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"AssetOnTubeDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tubeID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"AssetRemovedOnTube\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tubeID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"AssetSetOnTube\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tubeID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"NewAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"TubeActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"TubeDeactivated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"activateAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"}],\"name\":\"activateAssetOnTube\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"activateTube\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeAssetIDs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeTubeIDs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"}],\"name\":\"assetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_assetAddr\",\"type\":\"address\"}],\"name\":\"assetID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deactivateAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"}],\"name\":\"deactivateAssetOnTube\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deactivateTube\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"grant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isActiveAssetID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isActiveTubeID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_assetAddr\",\"type\":\"address\"}],\"name\":\"newAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"}],\"name\":\"removeAssetOnTube\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_assetAddr\",\"type\":\"address\"}],\"name\":\"setAssetOnTube\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AssetRegistry is an auto generated Go binding around an Ethereum contract.
type AssetRegistry struct {
	AssetRegistryCaller     // Read-only binding to the contract
	AssetRegistryTransactor // Write-only binding to the contract
	AssetRegistryFilterer   // Log filterer for contract events
}

// AssetRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetRegistrySession struct {
	Contract     *AssetRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetRegistryCallerSession struct {
	Contract *AssetRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AssetRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetRegistryTransactorSession struct {
	Contract     *AssetRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AssetRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetRegistryRaw struct {
	Contract *AssetRegistry // Generic contract binding to access the raw methods on
}

// AssetRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetRegistryCallerRaw struct {
	Contract *AssetRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AssetRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetRegistryTransactorRaw struct {
	Contract *AssetRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetRegistry creates a new instance of AssetRegistry, bound to a specific deployed contract.
func NewAssetRegistry(address common.Address, backend bind.ContractBackend) (*AssetRegistry, error) {
	contract, err := bindAssetRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetRegistry{AssetRegistryCaller: AssetRegistryCaller{contract: contract}, AssetRegistryTransactor: AssetRegistryTransactor{contract: contract}, AssetRegistryFilterer: AssetRegistryFilterer{contract: contract}}, nil
}

// NewAssetRegistryCaller creates a new read-only instance of AssetRegistry, bound to a specific deployed contract.
func NewAssetRegistryCaller(address common.Address, caller bind.ContractCaller) (*AssetRegistryCaller, error) {
	contract, err := bindAssetRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryCaller{contract: contract}, nil
}

// NewAssetRegistryTransactor creates a new write-only instance of AssetRegistry, bound to a specific deployed contract.
func NewAssetRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetRegistryTransactor, error) {
	contract, err := bindAssetRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryTransactor{contract: contract}, nil
}

// NewAssetRegistryFilterer creates a new log filterer instance of AssetRegistry, bound to a specific deployed contract.
func NewAssetRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetRegistryFilterer, error) {
	contract, err := bindAssetRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryFilterer{contract: contract}, nil
}

// bindAssetRegistry binds a generic wrapper to an already deployed contract.
func bindAssetRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetRegistry *AssetRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetRegistry.Contract.AssetRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetRegistry *AssetRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetRegistry.Contract.AssetRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetRegistry *AssetRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetRegistry.Contract.AssetRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetRegistry *AssetRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetRegistry *AssetRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetRegistry *AssetRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetRegistry.Contract.contract.Transact(opts, method, params...)
}

// ActiveAssetIDs is a free data retrieval call binding the contract method 0x28f6ce9a.
//
// Solidity: function activeAssetIDs(uint256 ) view returns(bool)
func (_AssetRegistry *AssetRegistryCaller) ActiveAssetIDs(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "activeAssetIDs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActiveAssetIDs is a free data retrieval call binding the contract method 0x28f6ce9a.
//
// Solidity: function activeAssetIDs(uint256 ) view returns(bool)
func (_AssetRegistry *AssetRegistrySession) ActiveAssetIDs(arg0 *big.Int) (bool, error) {
	return _AssetRegistry.Contract.ActiveAssetIDs(&_AssetRegistry.CallOpts, arg0)
}

// ActiveAssetIDs is a free data retrieval call binding the contract method 0x28f6ce9a.
//
// Solidity: function activeAssetIDs(uint256 ) view returns(bool)
func (_AssetRegistry *AssetRegistryCallerSession) ActiveAssetIDs(arg0 *big.Int) (bool, error) {
	return _AssetRegistry.Contract.ActiveAssetIDs(&_AssetRegistry.CallOpts, arg0)
}

// ActiveTubeIDs is a free data retrieval call binding the contract method 0x09498f64.
//
// Solidity: function activeTubeIDs(uint256 ) view returns(bool)
func (_AssetRegistry *AssetRegistryCaller) ActiveTubeIDs(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "activeTubeIDs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActiveTubeIDs is a free data retrieval call binding the contract method 0x09498f64.
//
// Solidity: function activeTubeIDs(uint256 ) view returns(bool)
func (_AssetRegistry *AssetRegistrySession) ActiveTubeIDs(arg0 *big.Int) (bool, error) {
	return _AssetRegistry.Contract.ActiveTubeIDs(&_AssetRegistry.CallOpts, arg0)
}

// ActiveTubeIDs is a free data retrieval call binding the contract method 0x09498f64.
//
// Solidity: function activeTubeIDs(uint256 ) view returns(bool)
func (_AssetRegistry *AssetRegistryCallerSession) ActiveTubeIDs(arg0 *big.Int) (bool, error) {
	return _AssetRegistry.Contract.ActiveTubeIDs(&_AssetRegistry.CallOpts, arg0)
}

// AssetAddress is a free data retrieval call binding the contract method 0x6fbd2b3c.
//
// Solidity: function assetAddress(uint256 _assetID, uint256 _tubeID) view returns(address)
func (_AssetRegistry *AssetRegistryCaller) AssetAddress(opts *bind.CallOpts, _assetID *big.Int, _tubeID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "assetAddress", _assetID, _tubeID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetAddress is a free data retrieval call binding the contract method 0x6fbd2b3c.
//
// Solidity: function assetAddress(uint256 _assetID, uint256 _tubeID) view returns(address)
func (_AssetRegistry *AssetRegistrySession) AssetAddress(_assetID *big.Int, _tubeID *big.Int) (common.Address, error) {
	return _AssetRegistry.Contract.AssetAddress(&_AssetRegistry.CallOpts, _assetID, _tubeID)
}

// AssetAddress is a free data retrieval call binding the contract method 0x6fbd2b3c.
//
// Solidity: function assetAddress(uint256 _assetID, uint256 _tubeID) view returns(address)
func (_AssetRegistry *AssetRegistryCallerSession) AssetAddress(_assetID *big.Int, _tubeID *big.Int) (common.Address, error) {
	return _AssetRegistry.Contract.AssetAddress(&_AssetRegistry.CallOpts, _assetID, _tubeID)
}

// AssetID is a free data retrieval call binding the contract method 0x4b7c88e3.
//
// Solidity: function assetID(uint256 _tubeID, address _assetAddr) view returns(uint256)
func (_AssetRegistry *AssetRegistryCaller) AssetID(opts *bind.CallOpts, _tubeID *big.Int, _assetAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "assetID", _tubeID, _assetAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetID is a free data retrieval call binding the contract method 0x4b7c88e3.
//
// Solidity: function assetID(uint256 _tubeID, address _assetAddr) view returns(uint256)
func (_AssetRegistry *AssetRegistrySession) AssetID(_tubeID *big.Int, _assetAddr common.Address) (*big.Int, error) {
	return _AssetRegistry.Contract.AssetID(&_AssetRegistry.CallOpts, _tubeID, _assetAddr)
}

// AssetID is a free data retrieval call binding the contract method 0x4b7c88e3.
//
// Solidity: function assetID(uint256 _tubeID, address _assetAddr) view returns(uint256)
func (_AssetRegistry *AssetRegistryCallerSession) AssetID(_tubeID *big.Int, _assetAddr common.Address) (*big.Int, error) {
	return _AssetRegistry.Contract.AssetID(&_AssetRegistry.CallOpts, _tubeID, _assetAddr)
}

// IsActive is a free data retrieval call binding the contract method 0x7c1e6091.
//
// Solidity: function isActive(uint256 _assetID, uint256 _tubeID) view returns(bool)
func (_AssetRegistry *AssetRegistryCaller) IsActive(opts *bind.CallOpts, _assetID *big.Int, _tubeID *big.Int) (bool, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "isActive", _assetID, _tubeID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x7c1e6091.
//
// Solidity: function isActive(uint256 _assetID, uint256 _tubeID) view returns(bool)
func (_AssetRegistry *AssetRegistrySession) IsActive(_assetID *big.Int, _tubeID *big.Int) (bool, error) {
	return _AssetRegistry.Contract.IsActive(&_AssetRegistry.CallOpts, _assetID, _tubeID)
}

// IsActive is a free data retrieval call binding the contract method 0x7c1e6091.
//
// Solidity: function isActive(uint256 _assetID, uint256 _tubeID) view returns(bool)
func (_AssetRegistry *AssetRegistryCallerSession) IsActive(_assetID *big.Int, _tubeID *big.Int) (bool, error) {
	return _AssetRegistry.Contract.IsActive(&_AssetRegistry.CallOpts, _assetID, _tubeID)
}

// IsActiveAssetID is a free data retrieval call binding the contract method 0x522dffc5.
//
// Solidity: function isActiveAssetID(uint256 _id) view returns(bool)
func (_AssetRegistry *AssetRegistryCaller) IsActiveAssetID(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "isActiveAssetID", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveAssetID is a free data retrieval call binding the contract method 0x522dffc5.
//
// Solidity: function isActiveAssetID(uint256 _id) view returns(bool)
func (_AssetRegistry *AssetRegistrySession) IsActiveAssetID(_id *big.Int) (bool, error) {
	return _AssetRegistry.Contract.IsActiveAssetID(&_AssetRegistry.CallOpts, _id)
}

// IsActiveAssetID is a free data retrieval call binding the contract method 0x522dffc5.
//
// Solidity: function isActiveAssetID(uint256 _id) view returns(bool)
func (_AssetRegistry *AssetRegistryCallerSession) IsActiveAssetID(_id *big.Int) (bool, error) {
	return _AssetRegistry.Contract.IsActiveAssetID(&_AssetRegistry.CallOpts, _id)
}

// IsActiveTubeID is a free data retrieval call binding the contract method 0x9ac0c145.
//
// Solidity: function isActiveTubeID(uint256 _id) view returns(bool)
func (_AssetRegistry *AssetRegistryCaller) IsActiveTubeID(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "isActiveTubeID", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveTubeID is a free data retrieval call binding the contract method 0x9ac0c145.
//
// Solidity: function isActiveTubeID(uint256 _id) view returns(bool)
func (_AssetRegistry *AssetRegistrySession) IsActiveTubeID(_id *big.Int) (bool, error) {
	return _AssetRegistry.Contract.IsActiveTubeID(&_AssetRegistry.CallOpts, _id)
}

// IsActiveTubeID is a free data retrieval call binding the contract method 0x9ac0c145.
//
// Solidity: function isActiveTubeID(uint256 _id) view returns(bool)
func (_AssetRegistry *AssetRegistryCallerSession) IsActiveTubeID(_id *big.Int) (bool, error) {
	return _AssetRegistry.Contract.IsActiveTubeID(&_AssetRegistry.CallOpts, _id)
}

// NumOfAssets is a free data retrieval call binding the contract method 0xf2cf057f.
//
// Solidity: function numOfAssets() view returns(uint256)
func (_AssetRegistry *AssetRegistryCaller) NumOfAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "numOfAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfAssets is a free data retrieval call binding the contract method 0xf2cf057f.
//
// Solidity: function numOfAssets() view returns(uint256)
func (_AssetRegistry *AssetRegistrySession) NumOfAssets() (*big.Int, error) {
	return _AssetRegistry.Contract.NumOfAssets(&_AssetRegistry.CallOpts)
}

// NumOfAssets is a free data retrieval call binding the contract method 0xf2cf057f.
//
// Solidity: function numOfAssets() view returns(uint256)
func (_AssetRegistry *AssetRegistryCallerSession) NumOfAssets() (*big.Int, error) {
	return _AssetRegistry.Contract.NumOfAssets(&_AssetRegistry.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_AssetRegistry *AssetRegistryCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_AssetRegistry *AssetRegistrySession) Operators(arg0 common.Address) (bool, error) {
	return _AssetRegistry.Contract.Operators(&_AssetRegistry.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_AssetRegistry *AssetRegistryCallerSession) Operators(arg0 common.Address) (bool, error) {
	return _AssetRegistry.Contract.Operators(&_AssetRegistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetRegistry *AssetRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetRegistry *AssetRegistrySession) Owner() (common.Address, error) {
	return _AssetRegistry.Contract.Owner(&_AssetRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetRegistry *AssetRegistryCallerSession) Owner() (common.Address, error) {
	return _AssetRegistry.Contract.Owner(&_AssetRegistry.CallOpts)
}

// ActivateAsset is a paid mutator transaction binding the contract method 0xffba205f.
//
// Solidity: function activateAsset(uint256 _id) returns()
func (_AssetRegistry *AssetRegistryTransactor) ActivateAsset(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "activateAsset", _id)
}

// ActivateAsset is a paid mutator transaction binding the contract method 0xffba205f.
//
// Solidity: function activateAsset(uint256 _id) returns()
func (_AssetRegistry *AssetRegistrySession) ActivateAsset(_id *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.ActivateAsset(&_AssetRegistry.TransactOpts, _id)
}

// ActivateAsset is a paid mutator transaction binding the contract method 0xffba205f.
//
// Solidity: function activateAsset(uint256 _id) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) ActivateAsset(_id *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.ActivateAsset(&_AssetRegistry.TransactOpts, _id)
}

// ActivateAssetOnTube is a paid mutator transaction binding the contract method 0x6c52c3aa.
//
// Solidity: function activateAssetOnTube(uint256 _assetID, uint256 _tubeID) returns()
func (_AssetRegistry *AssetRegistryTransactor) ActivateAssetOnTube(opts *bind.TransactOpts, _assetID *big.Int, _tubeID *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "activateAssetOnTube", _assetID, _tubeID)
}

// ActivateAssetOnTube is a paid mutator transaction binding the contract method 0x6c52c3aa.
//
// Solidity: function activateAssetOnTube(uint256 _assetID, uint256 _tubeID) returns()
func (_AssetRegistry *AssetRegistrySession) ActivateAssetOnTube(_assetID *big.Int, _tubeID *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.ActivateAssetOnTube(&_AssetRegistry.TransactOpts, _assetID, _tubeID)
}

// ActivateAssetOnTube is a paid mutator transaction binding the contract method 0x6c52c3aa.
//
// Solidity: function activateAssetOnTube(uint256 _assetID, uint256 _tubeID) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) ActivateAssetOnTube(_assetID *big.Int, _tubeID *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.ActivateAssetOnTube(&_AssetRegistry.TransactOpts, _assetID, _tubeID)
}

// ActivateTube is a paid mutator transaction binding the contract method 0xcbc532a0.
//
// Solidity: function activateTube(uint256 _id) returns()
func (_AssetRegistry *AssetRegistryTransactor) ActivateTube(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "activateTube", _id)
}

// ActivateTube is a paid mutator transaction binding the contract method 0xcbc532a0.
//
// Solidity: function activateTube(uint256 _id) returns()
func (_AssetRegistry *AssetRegistrySession) ActivateTube(_id *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.ActivateTube(&_AssetRegistry.TransactOpts, _id)
}

// ActivateTube is a paid mutator transaction binding the contract method 0xcbc532a0.
//
// Solidity: function activateTube(uint256 _id) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) ActivateTube(_id *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.ActivateTube(&_AssetRegistry.TransactOpts, _id)
}

// DeactivateAsset is a paid mutator transaction binding the contract method 0x319af5e0.
//
// Solidity: function deactivateAsset(uint256 _id) returns()
func (_AssetRegistry *AssetRegistryTransactor) DeactivateAsset(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "deactivateAsset", _id)
}

// DeactivateAsset is a paid mutator transaction binding the contract method 0x319af5e0.
//
// Solidity: function deactivateAsset(uint256 _id) returns()
func (_AssetRegistry *AssetRegistrySession) DeactivateAsset(_id *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.DeactivateAsset(&_AssetRegistry.TransactOpts, _id)
}

// DeactivateAsset is a paid mutator transaction binding the contract method 0x319af5e0.
//
// Solidity: function deactivateAsset(uint256 _id) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) DeactivateAsset(_id *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.DeactivateAsset(&_AssetRegistry.TransactOpts, _id)
}

// DeactivateAssetOnTube is a paid mutator transaction binding the contract method 0xe14a48fc.
//
// Solidity: function deactivateAssetOnTube(uint256 _assetID, uint256 _tubeID) returns()
func (_AssetRegistry *AssetRegistryTransactor) DeactivateAssetOnTube(opts *bind.TransactOpts, _assetID *big.Int, _tubeID *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "deactivateAssetOnTube", _assetID, _tubeID)
}

// DeactivateAssetOnTube is a paid mutator transaction binding the contract method 0xe14a48fc.
//
// Solidity: function deactivateAssetOnTube(uint256 _assetID, uint256 _tubeID) returns()
func (_AssetRegistry *AssetRegistrySession) DeactivateAssetOnTube(_assetID *big.Int, _tubeID *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.DeactivateAssetOnTube(&_AssetRegistry.TransactOpts, _assetID, _tubeID)
}

// DeactivateAssetOnTube is a paid mutator transaction binding the contract method 0xe14a48fc.
//
// Solidity: function deactivateAssetOnTube(uint256 _assetID, uint256 _tubeID) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) DeactivateAssetOnTube(_assetID *big.Int, _tubeID *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.DeactivateAssetOnTube(&_AssetRegistry.TransactOpts, _assetID, _tubeID)
}

// DeactivateTube is a paid mutator transaction binding the contract method 0xf7596d92.
//
// Solidity: function deactivateTube(uint256 _id) returns()
func (_AssetRegistry *AssetRegistryTransactor) DeactivateTube(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "deactivateTube", _id)
}

// DeactivateTube is a paid mutator transaction binding the contract method 0xf7596d92.
//
// Solidity: function deactivateTube(uint256 _id) returns()
func (_AssetRegistry *AssetRegistrySession) DeactivateTube(_id *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.DeactivateTube(&_AssetRegistry.TransactOpts, _id)
}

// DeactivateTube is a paid mutator transaction binding the contract method 0xf7596d92.
//
// Solidity: function deactivateTube(uint256 _id) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) DeactivateTube(_id *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.DeactivateTube(&_AssetRegistry.TransactOpts, _id)
}

// Grant is a paid mutator transaction binding the contract method 0x70284d19.
//
// Solidity: function grant(address _account) returns()
func (_AssetRegistry *AssetRegistryTransactor) Grant(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "grant", _account)
}

// Grant is a paid mutator transaction binding the contract method 0x70284d19.
//
// Solidity: function grant(address _account) returns()
func (_AssetRegistry *AssetRegistrySession) Grant(_account common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.Grant(&_AssetRegistry.TransactOpts, _account)
}

// Grant is a paid mutator transaction binding the contract method 0x70284d19.
//
// Solidity: function grant(address _account) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) Grant(_account common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.Grant(&_AssetRegistry.TransactOpts, _account)
}

// NewAsset is a paid mutator transaction binding the contract method 0x56a24c3d.
//
// Solidity: function newAsset(uint256 _tubeID, address _assetAddr) returns()
func (_AssetRegistry *AssetRegistryTransactor) NewAsset(opts *bind.TransactOpts, _tubeID *big.Int, _assetAddr common.Address) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "newAsset", _tubeID, _assetAddr)
}

// NewAsset is a paid mutator transaction binding the contract method 0x56a24c3d.
//
// Solidity: function newAsset(uint256 _tubeID, address _assetAddr) returns()
func (_AssetRegistry *AssetRegistrySession) NewAsset(_tubeID *big.Int, _assetAddr common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.NewAsset(&_AssetRegistry.TransactOpts, _tubeID, _assetAddr)
}

// NewAsset is a paid mutator transaction binding the contract method 0x56a24c3d.
//
// Solidity: function newAsset(uint256 _tubeID, address _assetAddr) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) NewAsset(_tubeID *big.Int, _assetAddr common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.NewAsset(&_AssetRegistry.TransactOpts, _tubeID, _assetAddr)
}

// RemoveAssetOnTube is a paid mutator transaction binding the contract method 0x41fb810a.
//
// Solidity: function removeAssetOnTube(uint256 _assetID, uint256 _tubeID) returns()
func (_AssetRegistry *AssetRegistryTransactor) RemoveAssetOnTube(opts *bind.TransactOpts, _assetID *big.Int, _tubeID *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "removeAssetOnTube", _assetID, _tubeID)
}

// RemoveAssetOnTube is a paid mutator transaction binding the contract method 0x41fb810a.
//
// Solidity: function removeAssetOnTube(uint256 _assetID, uint256 _tubeID) returns()
func (_AssetRegistry *AssetRegistrySession) RemoveAssetOnTube(_assetID *big.Int, _tubeID *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.RemoveAssetOnTube(&_AssetRegistry.TransactOpts, _assetID, _tubeID)
}

// RemoveAssetOnTube is a paid mutator transaction binding the contract method 0x41fb810a.
//
// Solidity: function removeAssetOnTube(uint256 _assetID, uint256 _tubeID) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) RemoveAssetOnTube(_assetID *big.Int, _tubeID *big.Int) (*types.Transaction, error) {
	return _AssetRegistry.Contract.RemoveAssetOnTube(&_AssetRegistry.TransactOpts, _assetID, _tubeID)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AssetRegistry *AssetRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AssetRegistry *AssetRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _AssetRegistry.Contract.RenounceOwnership(&_AssetRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AssetRegistry *AssetRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AssetRegistry.Contract.RenounceOwnership(&_AssetRegistry.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _account) returns()
func (_AssetRegistry *AssetRegistryTransactor) Revoke(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "revoke", _account)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _account) returns()
func (_AssetRegistry *AssetRegistrySession) Revoke(_account common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.Revoke(&_AssetRegistry.TransactOpts, _account)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _account) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) Revoke(_account common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.Revoke(&_AssetRegistry.TransactOpts, _account)
}

// SetAssetOnTube is a paid mutator transaction binding the contract method 0x05771c5a.
//
// Solidity: function setAssetOnTube(uint256 _assetID, uint256 _tubeID, address _assetAddr) returns()
func (_AssetRegistry *AssetRegistryTransactor) SetAssetOnTube(opts *bind.TransactOpts, _assetID *big.Int, _tubeID *big.Int, _assetAddr common.Address) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "setAssetOnTube", _assetID, _tubeID, _assetAddr)
}

// SetAssetOnTube is a paid mutator transaction binding the contract method 0x05771c5a.
//
// Solidity: function setAssetOnTube(uint256 _assetID, uint256 _tubeID, address _assetAddr) returns()
func (_AssetRegistry *AssetRegistrySession) SetAssetOnTube(_assetID *big.Int, _tubeID *big.Int, _assetAddr common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.SetAssetOnTube(&_AssetRegistry.TransactOpts, _assetID, _tubeID, _assetAddr)
}

// SetAssetOnTube is a paid mutator transaction binding the contract method 0x05771c5a.
//
// Solidity: function setAssetOnTube(uint256 _assetID, uint256 _tubeID, address _assetAddr) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) SetAssetOnTube(_assetID *big.Int, _tubeID *big.Int, _assetAddr common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.SetAssetOnTube(&_AssetRegistry.TransactOpts, _assetID, _tubeID, _assetAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AssetRegistry *AssetRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AssetRegistry *AssetRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.TransferOwnership(&_AssetRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.TransferOwnership(&_AssetRegistry.TransactOpts, newOwner)
}

// AssetRegistryAssetActivatedIterator is returned from FilterAssetActivated and is used to iterate over the raw logs and unpacked data for AssetActivated events raised by the AssetRegistry contract.
type AssetRegistryAssetActivatedIterator struct {
	Event *AssetRegistryAssetActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetRegistryAssetActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryAssetActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetRegistryAssetActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetRegistryAssetActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryAssetActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryAssetActivated represents a AssetActivated event raised by the AssetRegistry contract.
type AssetRegistryAssetActivated struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAssetActivated is a free log retrieval operation binding the contract event 0x23a39165e7ea63387856bbb51381519f6966967b3e854e9b16035b778c1cb97d.
//
// Solidity: event AssetActivated(uint256 indexed id)
func (_AssetRegistry *AssetRegistryFilterer) FilterAssetActivated(opts *bind.FilterOpts, id []*big.Int) (*AssetRegistryAssetActivatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "AssetActivated", idRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryAssetActivatedIterator{contract: _AssetRegistry.contract, event: "AssetActivated", logs: logs, sub: sub}, nil
}

// WatchAssetActivated is a free log subscription operation binding the contract event 0x23a39165e7ea63387856bbb51381519f6966967b3e854e9b16035b778c1cb97d.
//
// Solidity: event AssetActivated(uint256 indexed id)
func (_AssetRegistry *AssetRegistryFilterer) WatchAssetActivated(opts *bind.WatchOpts, sink chan<- *AssetRegistryAssetActivated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "AssetActivated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryAssetActivated)
				if err := _AssetRegistry.contract.UnpackLog(event, "AssetActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetActivated is a log parse operation binding the contract event 0x23a39165e7ea63387856bbb51381519f6966967b3e854e9b16035b778c1cb97d.
//
// Solidity: event AssetActivated(uint256 indexed id)
func (_AssetRegistry *AssetRegistryFilterer) ParseAssetActivated(log types.Log) (*AssetRegistryAssetActivated, error) {
	event := new(AssetRegistryAssetActivated)
	if err := _AssetRegistry.contract.UnpackLog(event, "AssetActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryAssetDeactivatedIterator is returned from FilterAssetDeactivated and is used to iterate over the raw logs and unpacked data for AssetDeactivated events raised by the AssetRegistry contract.
type AssetRegistryAssetDeactivatedIterator struct {
	Event *AssetRegistryAssetDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetRegistryAssetDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryAssetDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetRegistryAssetDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetRegistryAssetDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryAssetDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryAssetDeactivated represents a AssetDeactivated event raised by the AssetRegistry contract.
type AssetRegistryAssetDeactivated struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAssetDeactivated is a free log retrieval operation binding the contract event 0x24eeae42f647ed9085bd63c948c8201b093844d2ad153d02e9c42950d0306ccc.
//
// Solidity: event AssetDeactivated(uint256 indexed id)
func (_AssetRegistry *AssetRegistryFilterer) FilterAssetDeactivated(opts *bind.FilterOpts, id []*big.Int) (*AssetRegistryAssetDeactivatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "AssetDeactivated", idRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryAssetDeactivatedIterator{contract: _AssetRegistry.contract, event: "AssetDeactivated", logs: logs, sub: sub}, nil
}

// WatchAssetDeactivated is a free log subscription operation binding the contract event 0x24eeae42f647ed9085bd63c948c8201b093844d2ad153d02e9c42950d0306ccc.
//
// Solidity: event AssetDeactivated(uint256 indexed id)
func (_AssetRegistry *AssetRegistryFilterer) WatchAssetDeactivated(opts *bind.WatchOpts, sink chan<- *AssetRegistryAssetDeactivated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "AssetDeactivated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryAssetDeactivated)
				if err := _AssetRegistry.contract.UnpackLog(event, "AssetDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetDeactivated is a log parse operation binding the contract event 0x24eeae42f647ed9085bd63c948c8201b093844d2ad153d02e9c42950d0306ccc.
//
// Solidity: event AssetDeactivated(uint256 indexed id)
func (_AssetRegistry *AssetRegistryFilterer) ParseAssetDeactivated(log types.Log) (*AssetRegistryAssetDeactivated, error) {
	event := new(AssetRegistryAssetDeactivated)
	if err := _AssetRegistry.contract.UnpackLog(event, "AssetDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryAssetOnTubeActivatedIterator is returned from FilterAssetOnTubeActivated and is used to iterate over the raw logs and unpacked data for AssetOnTubeActivated events raised by the AssetRegistry contract.
type AssetRegistryAssetOnTubeActivatedIterator struct {
	Event *AssetRegistryAssetOnTubeActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetRegistryAssetOnTubeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryAssetOnTubeActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetRegistryAssetOnTubeActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetRegistryAssetOnTubeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryAssetOnTubeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryAssetOnTubeActivated represents a AssetOnTubeActivated event raised by the AssetRegistry contract.
type AssetRegistryAssetOnTubeActivated struct {
	AssetID      *big.Int
	TubeID       *big.Int
	AssetAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAssetOnTubeActivated is a free log retrieval operation binding the contract event 0x6c9c1d109893f0bc1950a33215d2f2c5ea1d681db5be8e7bed77e6e9d1c836c6.
//
// Solidity: event AssetOnTubeActivated(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) FilterAssetOnTubeActivated(opts *bind.FilterOpts, assetID []*big.Int, tubeID []*big.Int, assetAddress []common.Address) (*AssetRegistryAssetOnTubeActivatedIterator, error) {

	var assetIDRule []interface{}
	for _, assetIDItem := range assetID {
		assetIDRule = append(assetIDRule, assetIDItem)
	}
	var tubeIDRule []interface{}
	for _, tubeIDItem := range tubeID {
		tubeIDRule = append(tubeIDRule, tubeIDItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "AssetOnTubeActivated", assetIDRule, tubeIDRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryAssetOnTubeActivatedIterator{contract: _AssetRegistry.contract, event: "AssetOnTubeActivated", logs: logs, sub: sub}, nil
}

// WatchAssetOnTubeActivated is a free log subscription operation binding the contract event 0x6c9c1d109893f0bc1950a33215d2f2c5ea1d681db5be8e7bed77e6e9d1c836c6.
//
// Solidity: event AssetOnTubeActivated(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) WatchAssetOnTubeActivated(opts *bind.WatchOpts, sink chan<- *AssetRegistryAssetOnTubeActivated, assetID []*big.Int, tubeID []*big.Int, assetAddress []common.Address) (event.Subscription, error) {

	var assetIDRule []interface{}
	for _, assetIDItem := range assetID {
		assetIDRule = append(assetIDRule, assetIDItem)
	}
	var tubeIDRule []interface{}
	for _, tubeIDItem := range tubeID {
		tubeIDRule = append(tubeIDRule, tubeIDItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "AssetOnTubeActivated", assetIDRule, tubeIDRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryAssetOnTubeActivated)
				if err := _AssetRegistry.contract.UnpackLog(event, "AssetOnTubeActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetOnTubeActivated is a log parse operation binding the contract event 0x6c9c1d109893f0bc1950a33215d2f2c5ea1d681db5be8e7bed77e6e9d1c836c6.
//
// Solidity: event AssetOnTubeActivated(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) ParseAssetOnTubeActivated(log types.Log) (*AssetRegistryAssetOnTubeActivated, error) {
	event := new(AssetRegistryAssetOnTubeActivated)
	if err := _AssetRegistry.contract.UnpackLog(event, "AssetOnTubeActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryAssetOnTubeDeactivatedIterator is returned from FilterAssetOnTubeDeactivated and is used to iterate over the raw logs and unpacked data for AssetOnTubeDeactivated events raised by the AssetRegistry contract.
type AssetRegistryAssetOnTubeDeactivatedIterator struct {
	Event *AssetRegistryAssetOnTubeDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetRegistryAssetOnTubeDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryAssetOnTubeDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetRegistryAssetOnTubeDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetRegistryAssetOnTubeDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryAssetOnTubeDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryAssetOnTubeDeactivated represents a AssetOnTubeDeactivated event raised by the AssetRegistry contract.
type AssetRegistryAssetOnTubeDeactivated struct {
	AssetID      *big.Int
	TubeID       *big.Int
	AssetAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAssetOnTubeDeactivated is a free log retrieval operation binding the contract event 0x7d03d325ea9bc55df47604a2470923b475a892e3e7597869be979a2f08241f31.
//
// Solidity: event AssetOnTubeDeactivated(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) FilterAssetOnTubeDeactivated(opts *bind.FilterOpts, assetID []*big.Int, tubeID []*big.Int, assetAddress []common.Address) (*AssetRegistryAssetOnTubeDeactivatedIterator, error) {

	var assetIDRule []interface{}
	for _, assetIDItem := range assetID {
		assetIDRule = append(assetIDRule, assetIDItem)
	}
	var tubeIDRule []interface{}
	for _, tubeIDItem := range tubeID {
		tubeIDRule = append(tubeIDRule, tubeIDItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "AssetOnTubeDeactivated", assetIDRule, tubeIDRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryAssetOnTubeDeactivatedIterator{contract: _AssetRegistry.contract, event: "AssetOnTubeDeactivated", logs: logs, sub: sub}, nil
}

// WatchAssetOnTubeDeactivated is a free log subscription operation binding the contract event 0x7d03d325ea9bc55df47604a2470923b475a892e3e7597869be979a2f08241f31.
//
// Solidity: event AssetOnTubeDeactivated(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) WatchAssetOnTubeDeactivated(opts *bind.WatchOpts, sink chan<- *AssetRegistryAssetOnTubeDeactivated, assetID []*big.Int, tubeID []*big.Int, assetAddress []common.Address) (event.Subscription, error) {

	var assetIDRule []interface{}
	for _, assetIDItem := range assetID {
		assetIDRule = append(assetIDRule, assetIDItem)
	}
	var tubeIDRule []interface{}
	for _, tubeIDItem := range tubeID {
		tubeIDRule = append(tubeIDRule, tubeIDItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "AssetOnTubeDeactivated", assetIDRule, tubeIDRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryAssetOnTubeDeactivated)
				if err := _AssetRegistry.contract.UnpackLog(event, "AssetOnTubeDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetOnTubeDeactivated is a log parse operation binding the contract event 0x7d03d325ea9bc55df47604a2470923b475a892e3e7597869be979a2f08241f31.
//
// Solidity: event AssetOnTubeDeactivated(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) ParseAssetOnTubeDeactivated(log types.Log) (*AssetRegistryAssetOnTubeDeactivated, error) {
	event := new(AssetRegistryAssetOnTubeDeactivated)
	if err := _AssetRegistry.contract.UnpackLog(event, "AssetOnTubeDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryAssetRemovedOnTubeIterator is returned from FilterAssetRemovedOnTube and is used to iterate over the raw logs and unpacked data for AssetRemovedOnTube events raised by the AssetRegistry contract.
type AssetRegistryAssetRemovedOnTubeIterator struct {
	Event *AssetRegistryAssetRemovedOnTube // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetRegistryAssetRemovedOnTubeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryAssetRemovedOnTube)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetRegistryAssetRemovedOnTube)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetRegistryAssetRemovedOnTubeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryAssetRemovedOnTubeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryAssetRemovedOnTube represents a AssetRemovedOnTube event raised by the AssetRegistry contract.
type AssetRegistryAssetRemovedOnTube struct {
	AssetID      *big.Int
	TubeID       *big.Int
	AssetAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAssetRemovedOnTube is a free log retrieval operation binding the contract event 0x39a9692b88c509cbe139bbcb3a7b1255588ea31cccafdc2243300374356dd65c.
//
// Solidity: event AssetRemovedOnTube(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) FilterAssetRemovedOnTube(opts *bind.FilterOpts, assetID []*big.Int, tubeID []*big.Int, assetAddress []common.Address) (*AssetRegistryAssetRemovedOnTubeIterator, error) {

	var assetIDRule []interface{}
	for _, assetIDItem := range assetID {
		assetIDRule = append(assetIDRule, assetIDItem)
	}
	var tubeIDRule []interface{}
	for _, tubeIDItem := range tubeID {
		tubeIDRule = append(tubeIDRule, tubeIDItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "AssetRemovedOnTube", assetIDRule, tubeIDRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryAssetRemovedOnTubeIterator{contract: _AssetRegistry.contract, event: "AssetRemovedOnTube", logs: logs, sub: sub}, nil
}

// WatchAssetRemovedOnTube is a free log subscription operation binding the contract event 0x39a9692b88c509cbe139bbcb3a7b1255588ea31cccafdc2243300374356dd65c.
//
// Solidity: event AssetRemovedOnTube(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) WatchAssetRemovedOnTube(opts *bind.WatchOpts, sink chan<- *AssetRegistryAssetRemovedOnTube, assetID []*big.Int, tubeID []*big.Int, assetAddress []common.Address) (event.Subscription, error) {

	var assetIDRule []interface{}
	for _, assetIDItem := range assetID {
		assetIDRule = append(assetIDRule, assetIDItem)
	}
	var tubeIDRule []interface{}
	for _, tubeIDItem := range tubeID {
		tubeIDRule = append(tubeIDRule, tubeIDItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "AssetRemovedOnTube", assetIDRule, tubeIDRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryAssetRemovedOnTube)
				if err := _AssetRegistry.contract.UnpackLog(event, "AssetRemovedOnTube", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetRemovedOnTube is a log parse operation binding the contract event 0x39a9692b88c509cbe139bbcb3a7b1255588ea31cccafdc2243300374356dd65c.
//
// Solidity: event AssetRemovedOnTube(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) ParseAssetRemovedOnTube(log types.Log) (*AssetRegistryAssetRemovedOnTube, error) {
	event := new(AssetRegistryAssetRemovedOnTube)
	if err := _AssetRegistry.contract.UnpackLog(event, "AssetRemovedOnTube", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryAssetSetOnTubeIterator is returned from FilterAssetSetOnTube and is used to iterate over the raw logs and unpacked data for AssetSetOnTube events raised by the AssetRegistry contract.
type AssetRegistryAssetSetOnTubeIterator struct {
	Event *AssetRegistryAssetSetOnTube // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetRegistryAssetSetOnTubeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryAssetSetOnTube)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetRegistryAssetSetOnTube)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetRegistryAssetSetOnTubeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryAssetSetOnTubeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryAssetSetOnTube represents a AssetSetOnTube event raised by the AssetRegistry contract.
type AssetRegistryAssetSetOnTube struct {
	AssetID      *big.Int
	TubeID       *big.Int
	AssetAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAssetSetOnTube is a free log retrieval operation binding the contract event 0x167597a772037aedd0d77fa52a5511a2c9a95ed66da54dc8ae85eac64d95f1cd.
//
// Solidity: event AssetSetOnTube(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) FilterAssetSetOnTube(opts *bind.FilterOpts, assetID []*big.Int, tubeID []*big.Int, assetAddress []common.Address) (*AssetRegistryAssetSetOnTubeIterator, error) {

	var assetIDRule []interface{}
	for _, assetIDItem := range assetID {
		assetIDRule = append(assetIDRule, assetIDItem)
	}
	var tubeIDRule []interface{}
	for _, tubeIDItem := range tubeID {
		tubeIDRule = append(tubeIDRule, tubeIDItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "AssetSetOnTube", assetIDRule, tubeIDRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryAssetSetOnTubeIterator{contract: _AssetRegistry.contract, event: "AssetSetOnTube", logs: logs, sub: sub}, nil
}

// WatchAssetSetOnTube is a free log subscription operation binding the contract event 0x167597a772037aedd0d77fa52a5511a2c9a95ed66da54dc8ae85eac64d95f1cd.
//
// Solidity: event AssetSetOnTube(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) WatchAssetSetOnTube(opts *bind.WatchOpts, sink chan<- *AssetRegistryAssetSetOnTube, assetID []*big.Int, tubeID []*big.Int, assetAddress []common.Address) (event.Subscription, error) {

	var assetIDRule []interface{}
	for _, assetIDItem := range assetID {
		assetIDRule = append(assetIDRule, assetIDItem)
	}
	var tubeIDRule []interface{}
	for _, tubeIDItem := range tubeID {
		tubeIDRule = append(tubeIDRule, tubeIDItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "AssetSetOnTube", assetIDRule, tubeIDRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryAssetSetOnTube)
				if err := _AssetRegistry.contract.UnpackLog(event, "AssetSetOnTube", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetSetOnTube is a log parse operation binding the contract event 0x167597a772037aedd0d77fa52a5511a2c9a95ed66da54dc8ae85eac64d95f1cd.
//
// Solidity: event AssetSetOnTube(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) ParseAssetSetOnTube(log types.Log) (*AssetRegistryAssetSetOnTube, error) {
	event := new(AssetRegistryAssetSetOnTube)
	if err := _AssetRegistry.contract.UnpackLog(event, "AssetSetOnTube", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryNewAssetIterator is returned from FilterNewAsset and is used to iterate over the raw logs and unpacked data for NewAsset events raised by the AssetRegistry contract.
type AssetRegistryNewAssetIterator struct {
	Event *AssetRegistryNewAsset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetRegistryNewAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryNewAsset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetRegistryNewAsset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetRegistryNewAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryNewAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryNewAsset represents a NewAsset event raised by the AssetRegistry contract.
type AssetRegistryNewAsset struct {
	AssetID      *big.Int
	TubeID       *big.Int
	AssetAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewAsset is a free log retrieval operation binding the contract event 0x8fad59b9ca87cdb9db03e7b23dd89dc8017b72b82abc02ebfe824e79f59e8108.
//
// Solidity: event NewAsset(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) FilterNewAsset(opts *bind.FilterOpts, assetID []*big.Int, tubeID []*big.Int, assetAddress []common.Address) (*AssetRegistryNewAssetIterator, error) {

	var assetIDRule []interface{}
	for _, assetIDItem := range assetID {
		assetIDRule = append(assetIDRule, assetIDItem)
	}
	var tubeIDRule []interface{}
	for _, tubeIDItem := range tubeID {
		tubeIDRule = append(tubeIDRule, tubeIDItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "NewAsset", assetIDRule, tubeIDRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryNewAssetIterator{contract: _AssetRegistry.contract, event: "NewAsset", logs: logs, sub: sub}, nil
}

// WatchNewAsset is a free log subscription operation binding the contract event 0x8fad59b9ca87cdb9db03e7b23dd89dc8017b72b82abc02ebfe824e79f59e8108.
//
// Solidity: event NewAsset(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) WatchNewAsset(opts *bind.WatchOpts, sink chan<- *AssetRegistryNewAsset, assetID []*big.Int, tubeID []*big.Int, assetAddress []common.Address) (event.Subscription, error) {

	var assetIDRule []interface{}
	for _, assetIDItem := range assetID {
		assetIDRule = append(assetIDRule, assetIDItem)
	}
	var tubeIDRule []interface{}
	for _, tubeIDItem := range tubeID {
		tubeIDRule = append(tubeIDRule, tubeIDItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "NewAsset", assetIDRule, tubeIDRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryNewAsset)
				if err := _AssetRegistry.contract.UnpackLog(event, "NewAsset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAsset is a log parse operation binding the contract event 0x8fad59b9ca87cdb9db03e7b23dd89dc8017b72b82abc02ebfe824e79f59e8108.
//
// Solidity: event NewAsset(uint256 indexed assetID, uint256 indexed tubeID, address indexed assetAddress)
func (_AssetRegistry *AssetRegistryFilterer) ParseNewAsset(log types.Log) (*AssetRegistryNewAsset, error) {
	event := new(AssetRegistryNewAsset)
	if err := _AssetRegistry.contract.UnpackLog(event, "NewAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryOperatorGrantedIterator is returned from FilterOperatorGranted and is used to iterate over the raw logs and unpacked data for OperatorGranted events raised by the AssetRegistry contract.
type AssetRegistryOperatorGrantedIterator struct {
	Event *AssetRegistryOperatorGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetRegistryOperatorGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryOperatorGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetRegistryOperatorGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetRegistryOperatorGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryOperatorGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryOperatorGranted represents a OperatorGranted event raised by the AssetRegistry contract.
type AssetRegistryOperatorGranted struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorGranted is a free log retrieval operation binding the contract event 0x5b6a420b639feabfb011a07742a6ae14ef72f35d1a330ed65f42216c95bc6c2e.
//
// Solidity: event OperatorGranted(address indexed operator)
func (_AssetRegistry *AssetRegistryFilterer) FilterOperatorGranted(opts *bind.FilterOpts, operator []common.Address) (*AssetRegistryOperatorGrantedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "OperatorGranted", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryOperatorGrantedIterator{contract: _AssetRegistry.contract, event: "OperatorGranted", logs: logs, sub: sub}, nil
}

// WatchOperatorGranted is a free log subscription operation binding the contract event 0x5b6a420b639feabfb011a07742a6ae14ef72f35d1a330ed65f42216c95bc6c2e.
//
// Solidity: event OperatorGranted(address indexed operator)
func (_AssetRegistry *AssetRegistryFilterer) WatchOperatorGranted(opts *bind.WatchOpts, sink chan<- *AssetRegistryOperatorGranted, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "OperatorGranted", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryOperatorGranted)
				if err := _AssetRegistry.contract.UnpackLog(event, "OperatorGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorGranted is a log parse operation binding the contract event 0x5b6a420b639feabfb011a07742a6ae14ef72f35d1a330ed65f42216c95bc6c2e.
//
// Solidity: event OperatorGranted(address indexed operator)
func (_AssetRegistry *AssetRegistryFilterer) ParseOperatorGranted(log types.Log) (*AssetRegistryOperatorGranted, error) {
	event := new(AssetRegistryOperatorGranted)
	if err := _AssetRegistry.contract.UnpackLog(event, "OperatorGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryOperatorRevokedIterator is returned from FilterOperatorRevoked and is used to iterate over the raw logs and unpacked data for OperatorRevoked events raised by the AssetRegistry contract.
type AssetRegistryOperatorRevokedIterator struct {
	Event *AssetRegistryOperatorRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetRegistryOperatorRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryOperatorRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetRegistryOperatorRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetRegistryOperatorRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryOperatorRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryOperatorRevoked represents a OperatorRevoked event raised by the AssetRegistry contract.
type AssetRegistryOperatorRevoked struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRevoked is a free log retrieval operation binding the contract event 0xa5f3b7626fd86ff989f1d22cf3d41d74591ea6eb99241079400b0c332a9a8f11.
//
// Solidity: event OperatorRevoked(address indexed operator)
func (_AssetRegistry *AssetRegistryFilterer) FilterOperatorRevoked(opts *bind.FilterOpts, operator []common.Address) (*AssetRegistryOperatorRevokedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "OperatorRevoked", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryOperatorRevokedIterator{contract: _AssetRegistry.contract, event: "OperatorRevoked", logs: logs, sub: sub}, nil
}

// WatchOperatorRevoked is a free log subscription operation binding the contract event 0xa5f3b7626fd86ff989f1d22cf3d41d74591ea6eb99241079400b0c332a9a8f11.
//
// Solidity: event OperatorRevoked(address indexed operator)
func (_AssetRegistry *AssetRegistryFilterer) WatchOperatorRevoked(opts *bind.WatchOpts, sink chan<- *AssetRegistryOperatorRevoked, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "OperatorRevoked", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryOperatorRevoked)
				if err := _AssetRegistry.contract.UnpackLog(event, "OperatorRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRevoked is a log parse operation binding the contract event 0xa5f3b7626fd86ff989f1d22cf3d41d74591ea6eb99241079400b0c332a9a8f11.
//
// Solidity: event OperatorRevoked(address indexed operator)
func (_AssetRegistry *AssetRegistryFilterer) ParseOperatorRevoked(log types.Log) (*AssetRegistryOperatorRevoked, error) {
	event := new(AssetRegistryOperatorRevoked)
	if err := _AssetRegistry.contract.UnpackLog(event, "OperatorRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AssetRegistry contract.
type AssetRegistryOwnershipTransferredIterator struct {
	Event *AssetRegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetRegistryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the AssetRegistry contract.
type AssetRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AssetRegistry *AssetRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AssetRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryOwnershipTransferredIterator{contract: _AssetRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AssetRegistry *AssetRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AssetRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryOwnershipTransferred)
				if err := _AssetRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AssetRegistry *AssetRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*AssetRegistryOwnershipTransferred, error) {
	event := new(AssetRegistryOwnershipTransferred)
	if err := _AssetRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryTubeActivatedIterator is returned from FilterTubeActivated and is used to iterate over the raw logs and unpacked data for TubeActivated events raised by the AssetRegistry contract.
type AssetRegistryTubeActivatedIterator struct {
	Event *AssetRegistryTubeActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetRegistryTubeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryTubeActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetRegistryTubeActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetRegistryTubeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryTubeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryTubeActivated represents a TubeActivated event raised by the AssetRegistry contract.
type AssetRegistryTubeActivated struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTubeActivated is a free log retrieval operation binding the contract event 0x7568aede893c13f81344bda19f47ebdc185794a3e607e9e7551633f3a2b6018f.
//
// Solidity: event TubeActivated(uint256 indexed id)
func (_AssetRegistry *AssetRegistryFilterer) FilterTubeActivated(opts *bind.FilterOpts, id []*big.Int) (*AssetRegistryTubeActivatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "TubeActivated", idRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryTubeActivatedIterator{contract: _AssetRegistry.contract, event: "TubeActivated", logs: logs, sub: sub}, nil
}

// WatchTubeActivated is a free log subscription operation binding the contract event 0x7568aede893c13f81344bda19f47ebdc185794a3e607e9e7551633f3a2b6018f.
//
// Solidity: event TubeActivated(uint256 indexed id)
func (_AssetRegistry *AssetRegistryFilterer) WatchTubeActivated(opts *bind.WatchOpts, sink chan<- *AssetRegistryTubeActivated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "TubeActivated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryTubeActivated)
				if err := _AssetRegistry.contract.UnpackLog(event, "TubeActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTubeActivated is a log parse operation binding the contract event 0x7568aede893c13f81344bda19f47ebdc185794a3e607e9e7551633f3a2b6018f.
//
// Solidity: event TubeActivated(uint256 indexed id)
func (_AssetRegistry *AssetRegistryFilterer) ParseTubeActivated(log types.Log) (*AssetRegistryTubeActivated, error) {
	event := new(AssetRegistryTubeActivated)
	if err := _AssetRegistry.contract.UnpackLog(event, "TubeActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryTubeDeactivatedIterator is returned from FilterTubeDeactivated and is used to iterate over the raw logs and unpacked data for TubeDeactivated events raised by the AssetRegistry contract.
type AssetRegistryTubeDeactivatedIterator struct {
	Event *AssetRegistryTubeDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetRegistryTubeDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryTubeDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetRegistryTubeDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetRegistryTubeDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryTubeDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryTubeDeactivated represents a TubeDeactivated event raised by the AssetRegistry contract.
type AssetRegistryTubeDeactivated struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTubeDeactivated is a free log retrieval operation binding the contract event 0x8cd1ea4aeba4967f28b1f8ab83f16d14f78da3d2de8ac61fbc8e680ee01cb0db.
//
// Solidity: event TubeDeactivated(uint256 indexed id)
func (_AssetRegistry *AssetRegistryFilterer) FilterTubeDeactivated(opts *bind.FilterOpts, id []*big.Int) (*AssetRegistryTubeDeactivatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "TubeDeactivated", idRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryTubeDeactivatedIterator{contract: _AssetRegistry.contract, event: "TubeDeactivated", logs: logs, sub: sub}, nil
}

// WatchTubeDeactivated is a free log subscription operation binding the contract event 0x8cd1ea4aeba4967f28b1f8ab83f16d14f78da3d2de8ac61fbc8e680ee01cb0db.
//
// Solidity: event TubeDeactivated(uint256 indexed id)
func (_AssetRegistry *AssetRegistryFilterer) WatchTubeDeactivated(opts *bind.WatchOpts, sink chan<- *AssetRegistryTubeDeactivated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "TubeDeactivated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryTubeDeactivated)
				if err := _AssetRegistry.contract.UnpackLog(event, "TubeDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTubeDeactivated is a log parse operation binding the contract event 0x8cd1ea4aeba4967f28b1f8ab83f16d14f78da3d2de8ac61fbc8e680ee01cb0db.
//
// Solidity: event TubeDeactivated(uint256 indexed id)
func (_AssetRegistry *AssetRegistryFilterer) ParseTubeDeactivated(log types.Log) (*AssetRegistryTubeDeactivated, error) {
	event := new(AssetRegistryTubeDeactivated)
	if err := _AssetRegistry.contract.UnpackLog(event, "TubeDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
