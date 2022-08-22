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

// ERC20TubeRouterABI is the input ABI used to generate the binding from.
const ERC20TubeRouterABI = "[{\"inputs\":[{\"internalType\":\"contractITube\",\"name\":\"_tube\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"targetTubeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RelayFeeReceipt\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_crosschainToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_crosschainToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"depositToWithToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"}],\"name\":\"relayFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToken\",\"type\":\"address\"}],\"name\":\"setFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setRelayFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tube\",\"outputs\":[{\"internalType\":\"contractITube\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20TubeRouter is an auto generated Go binding around an Ethereum contract.
type ERC20TubeRouter struct {
	ERC20TubeRouterCaller     // Read-only binding to the contract
	ERC20TubeRouterTransactor // Write-only binding to the contract
	ERC20TubeRouterFilterer   // Log filterer for contract events
}

// ERC20TubeRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TubeRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TubeRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TubeRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TubeRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TubeRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TubeRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TubeRouterSession struct {
	Contract     *ERC20TubeRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20TubeRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TubeRouterCallerSession struct {
	Contract *ERC20TubeRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ERC20TubeRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TubeRouterTransactorSession struct {
	Contract     *ERC20TubeRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC20TubeRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TubeRouterRaw struct {
	Contract *ERC20TubeRouter // Generic contract binding to access the raw methods on
}

// ERC20TubeRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TubeRouterCallerRaw struct {
	Contract *ERC20TubeRouterCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TubeRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TubeRouterTransactorRaw struct {
	Contract *ERC20TubeRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TubeRouter creates a new instance of ERC20TubeRouter, bound to a specific deployed contract.
func NewERC20TubeRouter(address common.Address, backend bind.ContractBackend) (*ERC20TubeRouter, error) {
	contract, err := bindERC20TubeRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeRouter{ERC20TubeRouterCaller: ERC20TubeRouterCaller{contract: contract}, ERC20TubeRouterTransactor: ERC20TubeRouterTransactor{contract: contract}, ERC20TubeRouterFilterer: ERC20TubeRouterFilterer{contract: contract}}, nil
}

// NewERC20TubeRouterCaller creates a new read-only instance of ERC20TubeRouter, bound to a specific deployed contract.
func NewERC20TubeRouterCaller(address common.Address, caller bind.ContractCaller) (*ERC20TubeRouterCaller, error) {
	contract, err := bindERC20TubeRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeRouterCaller{contract: contract}, nil
}

// NewERC20TubeRouterTransactor creates a new write-only instance of ERC20TubeRouter, bound to a specific deployed contract.
func NewERC20TubeRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TubeRouterTransactor, error) {
	contract, err := bindERC20TubeRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeRouterTransactor{contract: contract}, nil
}

// NewERC20TubeRouterFilterer creates a new log filterer instance of ERC20TubeRouter, bound to a specific deployed contract.
func NewERC20TubeRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TubeRouterFilterer, error) {
	contract, err := bindERC20TubeRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeRouterFilterer{contract: contract}, nil
}

// bindERC20TubeRouter binds a generic wrapper to an already deployed contract.
func bindERC20TubeRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20TubeRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TubeRouter *ERC20TubeRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TubeRouter.Contract.ERC20TubeRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TubeRouter *ERC20TubeRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.ERC20TubeRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TubeRouter *ERC20TubeRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.ERC20TubeRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TubeRouter *ERC20TubeRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TubeRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TubeRouter *ERC20TubeRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TubeRouter *ERC20TubeRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.contract.Transact(opts, method, params...)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_ERC20TubeRouter *ERC20TubeRouterCaller) FeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TubeRouter.contract.Call(opts, &out, "feeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_ERC20TubeRouter *ERC20TubeRouterSession) FeeToken() (common.Address, error) {
	return _ERC20TubeRouter.Contract.FeeToken(&_ERC20TubeRouter.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_ERC20TubeRouter *ERC20TubeRouterCallerSession) FeeToken() (common.Address, error) {
	return _ERC20TubeRouter.Contract.FeeToken(&_ERC20TubeRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TubeRouter *ERC20TubeRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TubeRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TubeRouter *ERC20TubeRouterSession) Owner() (common.Address, error) {
	return _ERC20TubeRouter.Contract.Owner(&_ERC20TubeRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TubeRouter *ERC20TubeRouterCallerSession) Owner() (common.Address, error) {
	return _ERC20TubeRouter.Contract.Owner(&_ERC20TubeRouter.CallOpts)
}

// RelayFee is a free data retrieval call binding the contract method 0xde20b448.
//
// Solidity: function relayFee(uint256 _tubeID) view returns(uint256)
func (_ERC20TubeRouter *ERC20TubeRouterCaller) RelayFee(opts *bind.CallOpts, _tubeID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TubeRouter.contract.Call(opts, &out, "relayFee", _tubeID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayFee is a free data retrieval call binding the contract method 0xde20b448.
//
// Solidity: function relayFee(uint256 _tubeID) view returns(uint256)
func (_ERC20TubeRouter *ERC20TubeRouterSession) RelayFee(_tubeID *big.Int) (*big.Int, error) {
	return _ERC20TubeRouter.Contract.RelayFee(&_ERC20TubeRouter.CallOpts, _tubeID)
}

// RelayFee is a free data retrieval call binding the contract method 0xde20b448.
//
// Solidity: function relayFee(uint256 _tubeID) view returns(uint256)
func (_ERC20TubeRouter *ERC20TubeRouterCallerSession) RelayFee(_tubeID *big.Int) (*big.Int, error) {
	return _ERC20TubeRouter.Contract.RelayFee(&_ERC20TubeRouter.CallOpts, _tubeID)
}

// Tube is a free data retrieval call binding the contract method 0xfaaee0a7.
//
// Solidity: function tube() view returns(address)
func (_ERC20TubeRouter *ERC20TubeRouterCaller) Tube(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TubeRouter.contract.Call(opts, &out, "tube")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tube is a free data retrieval call binding the contract method 0xfaaee0a7.
//
// Solidity: function tube() view returns(address)
func (_ERC20TubeRouter *ERC20TubeRouterSession) Tube() (common.Address, error) {
	return _ERC20TubeRouter.Contract.Tube(&_ERC20TubeRouter.CallOpts)
}

// Tube is a free data retrieval call binding the contract method 0xfaaee0a7.
//
// Solidity: function tube() view returns(address)
func (_ERC20TubeRouter *ERC20TubeRouterCallerSession) Tube() (common.Address, error) {
	return _ERC20TubeRouter.Contract.Tube(&_ERC20TubeRouter.CallOpts)
}

// DepositTo is a paid mutator transaction binding the contract method 0x5379a5cb.
//
// Solidity: function depositTo(address _crosschainToken, uint256 _amount, uint256 _tubeID, address _recipient) payable returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactor) DepositTo(opts *bind.TransactOpts, _crosschainToken common.Address, _amount *big.Int, _tubeID *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.contract.Transact(opts, "depositTo", _crosschainToken, _amount, _tubeID, _recipient)
}

// DepositTo is a paid mutator transaction binding the contract method 0x5379a5cb.
//
// Solidity: function depositTo(address _crosschainToken, uint256 _amount, uint256 _tubeID, address _recipient) payable returns()
func (_ERC20TubeRouter *ERC20TubeRouterSession) DepositTo(_crosschainToken common.Address, _amount *big.Int, _tubeID *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.DepositTo(&_ERC20TubeRouter.TransactOpts, _crosschainToken, _amount, _tubeID, _recipient)
}

// DepositTo is a paid mutator transaction binding the contract method 0x5379a5cb.
//
// Solidity: function depositTo(address _crosschainToken, uint256 _amount, uint256 _tubeID, address _recipient) payable returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactorSession) DepositTo(_crosschainToken common.Address, _amount *big.Int, _tubeID *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.DepositTo(&_ERC20TubeRouter.TransactOpts, _crosschainToken, _amount, _tubeID, _recipient)
}

// DepositToWithToken is a paid mutator transaction binding the contract method 0x97aabf24.
//
// Solidity: function depositToWithToken(address _crosschainToken, uint256 _amount, uint256 _tubeID, address _recipient) payable returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactor) DepositToWithToken(opts *bind.TransactOpts, _crosschainToken common.Address, _amount *big.Int, _tubeID *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.contract.Transact(opts, "depositToWithToken", _crosschainToken, _amount, _tubeID, _recipient)
}

// DepositToWithToken is a paid mutator transaction binding the contract method 0x97aabf24.
//
// Solidity: function depositToWithToken(address _crosschainToken, uint256 _amount, uint256 _tubeID, address _recipient) payable returns()
func (_ERC20TubeRouter *ERC20TubeRouterSession) DepositToWithToken(_crosschainToken common.Address, _amount *big.Int, _tubeID *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.DepositToWithToken(&_ERC20TubeRouter.TransactOpts, _crosschainToken, _amount, _tubeID, _recipient)
}

// DepositToWithToken is a paid mutator transaction binding the contract method 0x97aabf24.
//
// Solidity: function depositToWithToken(address _crosschainToken, uint256 _amount, uint256 _tubeID, address _recipient) payable returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactorSession) DepositToWithToken(_crosschainToken common.Address, _amount *big.Int, _tubeID *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.DepositToWithToken(&_ERC20TubeRouter.TransactOpts, _crosschainToken, _amount, _tubeID, _recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TubeRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TubeRouter *ERC20TubeRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.RenounceOwnership(&_ERC20TubeRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.RenounceOwnership(&_ERC20TubeRouter.TransactOpts)
}

// SetFeeToken is a paid mutator transaction binding the contract method 0x15cce224.
//
// Solidity: function setFeeToken(address _feeToken) returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactor) SetFeeToken(opts *bind.TransactOpts, _feeToken common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.contract.Transact(opts, "setFeeToken", _feeToken)
}

// SetFeeToken is a paid mutator transaction binding the contract method 0x15cce224.
//
// Solidity: function setFeeToken(address _feeToken) returns()
func (_ERC20TubeRouter *ERC20TubeRouterSession) SetFeeToken(_feeToken common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.SetFeeToken(&_ERC20TubeRouter.TransactOpts, _feeToken)
}

// SetFeeToken is a paid mutator transaction binding the contract method 0x15cce224.
//
// Solidity: function setFeeToken(address _feeToken) returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactorSession) SetFeeToken(_feeToken common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.SetFeeToken(&_ERC20TubeRouter.TransactOpts, _feeToken)
}

// SetRelayFee is a paid mutator transaction binding the contract method 0xb71cd60d.
//
// Solidity: function setRelayFee(uint256 _tubeID, uint256 _fee) returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactor) SetRelayFee(opts *bind.TransactOpts, _tubeID *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _ERC20TubeRouter.contract.Transact(opts, "setRelayFee", _tubeID, _fee)
}

// SetRelayFee is a paid mutator transaction binding the contract method 0xb71cd60d.
//
// Solidity: function setRelayFee(uint256 _tubeID, uint256 _fee) returns()
func (_ERC20TubeRouter *ERC20TubeRouterSession) SetRelayFee(_tubeID *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.SetRelayFee(&_ERC20TubeRouter.TransactOpts, _tubeID, _fee)
}

// SetRelayFee is a paid mutator transaction binding the contract method 0xb71cd60d.
//
// Solidity: function setRelayFee(uint256 _tubeID, uint256 _fee) returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactorSession) SetRelayFee(_tubeID *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.SetRelayFee(&_ERC20TubeRouter.TransactOpts, _tubeID, _fee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TubeRouter *ERC20TubeRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.TransferOwnership(&_ERC20TubeRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.TransferOwnership(&_ERC20TubeRouter.TransactOpts, newOwner)
}

// WithdrawCoin is a paid mutator transaction binding the contract method 0x5a73928f.
//
// Solidity: function withdrawCoin(address _to) returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactor) WithdrawCoin(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.contract.Transact(opts, "withdrawCoin", _to)
}

// WithdrawCoin is a paid mutator transaction binding the contract method 0x5a73928f.
//
// Solidity: function withdrawCoin(address _to) returns()
func (_ERC20TubeRouter *ERC20TubeRouterSession) WithdrawCoin(_to common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.WithdrawCoin(&_ERC20TubeRouter.TransactOpts, _to)
}

// WithdrawCoin is a paid mutator transaction binding the contract method 0x5a73928f.
//
// Solidity: function withdrawCoin(address _to) returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactorSession) WithdrawCoin(_to common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.WithdrawCoin(&_ERC20TubeRouter.TransactOpts, _to)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address _to, address _token) returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactor) WithdrawToken(opts *bind.TransactOpts, _to common.Address, _token common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.contract.Transact(opts, "withdrawToken", _to, _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address _to, address _token) returns()
func (_ERC20TubeRouter *ERC20TubeRouterSession) WithdrawToken(_to common.Address, _token common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.WithdrawToken(&_ERC20TubeRouter.TransactOpts, _to, _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address _to, address _token) returns()
func (_ERC20TubeRouter *ERC20TubeRouterTransactorSession) WithdrawToken(_to common.Address, _token common.Address) (*types.Transaction, error) {
	return _ERC20TubeRouter.Contract.WithdrawToken(&_ERC20TubeRouter.TransactOpts, _to, _token)
}

// ERC20TubeRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20TubeRouter contract.
type ERC20TubeRouterOwnershipTransferredIterator struct {
	Event *ERC20TubeRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20TubeRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TubeRouterOwnershipTransferred)
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
		it.Event = new(ERC20TubeRouterOwnershipTransferred)
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
func (it *ERC20TubeRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TubeRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TubeRouterOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20TubeRouter contract.
type ERC20TubeRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TubeRouter *ERC20TubeRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20TubeRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TubeRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeRouterOwnershipTransferredIterator{contract: _ERC20TubeRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TubeRouter *ERC20TubeRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20TubeRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TubeRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TubeRouterOwnershipTransferred)
				if err := _ERC20TubeRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20TubeRouter *ERC20TubeRouterFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20TubeRouterOwnershipTransferred, error) {
	event := new(ERC20TubeRouterOwnershipTransferred)
	if err := _ERC20TubeRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TubeRouterRelayFeeReceiptIterator is returned from FilterRelayFeeReceipt and is used to iterate over the raw logs and unpacked data for RelayFeeReceipt events raised by the ERC20TubeRouter contract.
type ERC20TubeRouterRelayFeeReceiptIterator struct {
	Event *ERC20TubeRouterRelayFeeReceipt // Event containing the contract specifics and raw log

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
func (it *ERC20TubeRouterRelayFeeReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TubeRouterRelayFeeReceipt)
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
		it.Event = new(ERC20TubeRouterRelayFeeReceipt)
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
func (it *ERC20TubeRouterRelayFeeReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TubeRouterRelayFeeReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TubeRouterRelayFeeReceipt represents a RelayFeeReceipt event raised by the ERC20TubeRouter contract.
type ERC20TubeRouterRelayFeeReceipt struct {
	User         common.Address
	Token        common.Address
	TargetTubeID *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelayFeeReceipt is a free log retrieval operation binding the contract event 0x6cf4b4ca0bfcad2a221b400542e1c4947041f7b3b106b613eab8a0fa824919d0.
//
// Solidity: event RelayFeeReceipt(address indexed user, address indexed token, uint256 indexed targetTubeID, uint256 amount)
func (_ERC20TubeRouter *ERC20TubeRouterFilterer) FilterRelayFeeReceipt(opts *bind.FilterOpts, user []common.Address, token []common.Address, targetTubeID []*big.Int) (*ERC20TubeRouterRelayFeeReceiptIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var targetTubeIDRule []interface{}
	for _, targetTubeIDItem := range targetTubeID {
		targetTubeIDRule = append(targetTubeIDRule, targetTubeIDItem)
	}

	logs, sub, err := _ERC20TubeRouter.contract.FilterLogs(opts, "RelayFeeReceipt", userRule, tokenRule, targetTubeIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeRouterRelayFeeReceiptIterator{contract: _ERC20TubeRouter.contract, event: "RelayFeeReceipt", logs: logs, sub: sub}, nil
}

// WatchRelayFeeReceipt is a free log subscription operation binding the contract event 0x6cf4b4ca0bfcad2a221b400542e1c4947041f7b3b106b613eab8a0fa824919d0.
//
// Solidity: event RelayFeeReceipt(address indexed user, address indexed token, uint256 indexed targetTubeID, uint256 amount)
func (_ERC20TubeRouter *ERC20TubeRouterFilterer) WatchRelayFeeReceipt(opts *bind.WatchOpts, sink chan<- *ERC20TubeRouterRelayFeeReceipt, user []common.Address, token []common.Address, targetTubeID []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var targetTubeIDRule []interface{}
	for _, targetTubeIDItem := range targetTubeID {
		targetTubeIDRule = append(targetTubeIDRule, targetTubeIDItem)
	}

	logs, sub, err := _ERC20TubeRouter.contract.WatchLogs(opts, "RelayFeeReceipt", userRule, tokenRule, targetTubeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TubeRouterRelayFeeReceipt)
				if err := _ERC20TubeRouter.contract.UnpackLog(event, "RelayFeeReceipt", log); err != nil {
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

// ParseRelayFeeReceipt is a log parse operation binding the contract event 0x6cf4b4ca0bfcad2a221b400542e1c4947041f7b3b106b613eab8a0fa824919d0.
//
// Solidity: event RelayFeeReceipt(address indexed user, address indexed token, uint256 indexed targetTubeID, uint256 amount)
func (_ERC20TubeRouter *ERC20TubeRouterFilterer) ParseRelayFeeReceipt(log types.Log) (*ERC20TubeRouterRelayFeeReceipt, error) {
	event := new(ERC20TubeRouterRelayFeeReceipt)
	if err := _ERC20TubeRouter.contract.UnpackLog(event, "RelayFeeReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
