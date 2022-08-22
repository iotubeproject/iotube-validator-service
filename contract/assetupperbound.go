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

// AssetUpperBoundABI is the input ABI used to generate the binding from.
const AssetUpperBoundABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"}],\"name\":\"UpperBoundSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetID\",\"type\":\"uint256\"}],\"name\":\"getUpperBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_upperBound\",\"type\":\"uint256\"}],\"name\":\"setUpperBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AssetUpperBound is an auto generated Go binding around an Ethereum contract.
type AssetUpperBound struct {
	AssetUpperBoundCaller     // Read-only binding to the contract
	AssetUpperBoundTransactor // Write-only binding to the contract
	AssetUpperBoundFilterer   // Log filterer for contract events
}

// AssetUpperBoundCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetUpperBoundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetUpperBoundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetUpperBoundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetUpperBoundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetUpperBoundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetUpperBoundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetUpperBoundSession struct {
	Contract     *AssetUpperBound  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetUpperBoundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetUpperBoundCallerSession struct {
	Contract *AssetUpperBoundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AssetUpperBoundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetUpperBoundTransactorSession struct {
	Contract     *AssetUpperBoundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AssetUpperBoundRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetUpperBoundRaw struct {
	Contract *AssetUpperBound // Generic contract binding to access the raw methods on
}

// AssetUpperBoundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetUpperBoundCallerRaw struct {
	Contract *AssetUpperBoundCaller // Generic read-only contract binding to access the raw methods on
}

// AssetUpperBoundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetUpperBoundTransactorRaw struct {
	Contract *AssetUpperBoundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetUpperBound creates a new instance of AssetUpperBound, bound to a specific deployed contract.
func NewAssetUpperBound(address common.Address, backend bind.ContractBackend) (*AssetUpperBound, error) {
	contract, err := bindAssetUpperBound(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetUpperBound{AssetUpperBoundCaller: AssetUpperBoundCaller{contract: contract}, AssetUpperBoundTransactor: AssetUpperBoundTransactor{contract: contract}, AssetUpperBoundFilterer: AssetUpperBoundFilterer{contract: contract}}, nil
}

// NewAssetUpperBoundCaller creates a new read-only instance of AssetUpperBound, bound to a specific deployed contract.
func NewAssetUpperBoundCaller(address common.Address, caller bind.ContractCaller) (*AssetUpperBoundCaller, error) {
	contract, err := bindAssetUpperBound(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetUpperBoundCaller{contract: contract}, nil
}

// NewAssetUpperBoundTransactor creates a new write-only instance of AssetUpperBound, bound to a specific deployed contract.
func NewAssetUpperBoundTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetUpperBoundTransactor, error) {
	contract, err := bindAssetUpperBound(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetUpperBoundTransactor{contract: contract}, nil
}

// NewAssetUpperBoundFilterer creates a new log filterer instance of AssetUpperBound, bound to a specific deployed contract.
func NewAssetUpperBoundFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetUpperBoundFilterer, error) {
	contract, err := bindAssetUpperBound(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetUpperBoundFilterer{contract: contract}, nil
}

// bindAssetUpperBound binds a generic wrapper to an already deployed contract.
func bindAssetUpperBound(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetUpperBoundABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetUpperBound *AssetUpperBoundRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetUpperBound.Contract.AssetUpperBoundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetUpperBound *AssetUpperBoundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetUpperBound.Contract.AssetUpperBoundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetUpperBound *AssetUpperBoundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetUpperBound.Contract.AssetUpperBoundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetUpperBound *AssetUpperBoundCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetUpperBound.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetUpperBound *AssetUpperBoundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetUpperBound.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetUpperBound *AssetUpperBoundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetUpperBound.Contract.contract.Transact(opts, method, params...)
}

// GetUpperBound is a free data retrieval call binding the contract method 0xd7876291.
//
// Solidity: function getUpperBound(uint256 _assetID) view returns(uint256)
func (_AssetUpperBound *AssetUpperBoundCaller) GetUpperBound(opts *bind.CallOpts, _assetID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AssetUpperBound.contract.Call(opts, &out, "getUpperBound", _assetID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUpperBound is a free data retrieval call binding the contract method 0xd7876291.
//
// Solidity: function getUpperBound(uint256 _assetID) view returns(uint256)
func (_AssetUpperBound *AssetUpperBoundSession) GetUpperBound(_assetID *big.Int) (*big.Int, error) {
	return _AssetUpperBound.Contract.GetUpperBound(&_AssetUpperBound.CallOpts, _assetID)
}

// GetUpperBound is a free data retrieval call binding the contract method 0xd7876291.
//
// Solidity: function getUpperBound(uint256 _assetID) view returns(uint256)
func (_AssetUpperBound *AssetUpperBoundCallerSession) GetUpperBound(_assetID *big.Int) (*big.Int, error) {
	return _AssetUpperBound.Contract.GetUpperBound(&_AssetUpperBound.CallOpts, _assetID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetUpperBound *AssetUpperBoundCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetUpperBound.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetUpperBound *AssetUpperBoundSession) Owner() (common.Address, error) {
	return _AssetUpperBound.Contract.Owner(&_AssetUpperBound.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetUpperBound *AssetUpperBoundCallerSession) Owner() (common.Address, error) {
	return _AssetUpperBound.Contract.Owner(&_AssetUpperBound.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AssetUpperBound *AssetUpperBoundTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetUpperBound.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AssetUpperBound *AssetUpperBoundSession) RenounceOwnership() (*types.Transaction, error) {
	return _AssetUpperBound.Contract.RenounceOwnership(&_AssetUpperBound.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AssetUpperBound *AssetUpperBoundTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AssetUpperBound.Contract.RenounceOwnership(&_AssetUpperBound.TransactOpts)
}

// SetUpperBound is a paid mutator transaction binding the contract method 0x36c8856b.
//
// Solidity: function setUpperBound(uint256 _assetID, uint256 _upperBound) returns()
func (_AssetUpperBound *AssetUpperBoundTransactor) SetUpperBound(opts *bind.TransactOpts, _assetID *big.Int, _upperBound *big.Int) (*types.Transaction, error) {
	return _AssetUpperBound.contract.Transact(opts, "setUpperBound", _assetID, _upperBound)
}

// SetUpperBound is a paid mutator transaction binding the contract method 0x36c8856b.
//
// Solidity: function setUpperBound(uint256 _assetID, uint256 _upperBound) returns()
func (_AssetUpperBound *AssetUpperBoundSession) SetUpperBound(_assetID *big.Int, _upperBound *big.Int) (*types.Transaction, error) {
	return _AssetUpperBound.Contract.SetUpperBound(&_AssetUpperBound.TransactOpts, _assetID, _upperBound)
}

// SetUpperBound is a paid mutator transaction binding the contract method 0x36c8856b.
//
// Solidity: function setUpperBound(uint256 _assetID, uint256 _upperBound) returns()
func (_AssetUpperBound *AssetUpperBoundTransactorSession) SetUpperBound(_assetID *big.Int, _upperBound *big.Int) (*types.Transaction, error) {
	return _AssetUpperBound.Contract.SetUpperBound(&_AssetUpperBound.TransactOpts, _assetID, _upperBound)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AssetUpperBound *AssetUpperBoundTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AssetUpperBound.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AssetUpperBound *AssetUpperBoundSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AssetUpperBound.Contract.TransferOwnership(&_AssetUpperBound.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AssetUpperBound *AssetUpperBoundTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AssetUpperBound.Contract.TransferOwnership(&_AssetUpperBound.TransactOpts, newOwner)
}

// AssetUpperBoundOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AssetUpperBound contract.
type AssetUpperBoundOwnershipTransferredIterator struct {
	Event *AssetUpperBoundOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AssetUpperBoundOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetUpperBoundOwnershipTransferred)
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
		it.Event = new(AssetUpperBoundOwnershipTransferred)
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
func (it *AssetUpperBoundOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetUpperBoundOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetUpperBoundOwnershipTransferred represents a OwnershipTransferred event raised by the AssetUpperBound contract.
type AssetUpperBoundOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AssetUpperBound *AssetUpperBoundFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AssetUpperBoundOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AssetUpperBound.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AssetUpperBoundOwnershipTransferredIterator{contract: _AssetUpperBound.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AssetUpperBound *AssetUpperBoundFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AssetUpperBoundOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AssetUpperBound.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetUpperBoundOwnershipTransferred)
				if err := _AssetUpperBound.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AssetUpperBound *AssetUpperBoundFilterer) ParseOwnershipTransferred(log types.Log) (*AssetUpperBoundOwnershipTransferred, error) {
	event := new(AssetUpperBoundOwnershipTransferred)
	if err := _AssetUpperBound.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetUpperBoundUpperBoundSetIterator is returned from FilterUpperBoundSet and is used to iterate over the raw logs and unpacked data for UpperBoundSet events raised by the AssetUpperBound contract.
type AssetUpperBoundUpperBoundSetIterator struct {
	Event *AssetUpperBoundUpperBoundSet // Event containing the contract specifics and raw log

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
func (it *AssetUpperBoundUpperBoundSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetUpperBoundUpperBoundSet)
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
		it.Event = new(AssetUpperBoundUpperBoundSet)
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
func (it *AssetUpperBoundUpperBoundSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetUpperBoundUpperBoundSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetUpperBoundUpperBoundSet represents a UpperBoundSet event raised by the AssetUpperBound contract.
type AssetUpperBoundUpperBoundSet struct {
	AssetID    *big.Int
	UpperBound *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpperBoundSet is a free log retrieval operation binding the contract event 0x4745108fe432eb4e630177b0117dad9218e79b2a93d7d87d375676c9d05f8234.
//
// Solidity: event UpperBoundSet(uint256 indexed assetID, uint256 upperBound)
func (_AssetUpperBound *AssetUpperBoundFilterer) FilterUpperBoundSet(opts *bind.FilterOpts, assetID []*big.Int) (*AssetUpperBoundUpperBoundSetIterator, error) {

	var assetIDRule []interface{}
	for _, assetIDItem := range assetID {
		assetIDRule = append(assetIDRule, assetIDItem)
	}

	logs, sub, err := _AssetUpperBound.contract.FilterLogs(opts, "UpperBoundSet", assetIDRule)
	if err != nil {
		return nil, err
	}
	return &AssetUpperBoundUpperBoundSetIterator{contract: _AssetUpperBound.contract, event: "UpperBoundSet", logs: logs, sub: sub}, nil
}

// WatchUpperBoundSet is a free log subscription operation binding the contract event 0x4745108fe432eb4e630177b0117dad9218e79b2a93d7d87d375676c9d05f8234.
//
// Solidity: event UpperBoundSet(uint256 indexed assetID, uint256 upperBound)
func (_AssetUpperBound *AssetUpperBoundFilterer) WatchUpperBoundSet(opts *bind.WatchOpts, sink chan<- *AssetUpperBoundUpperBoundSet, assetID []*big.Int) (event.Subscription, error) {

	var assetIDRule []interface{}
	for _, assetIDItem := range assetID {
		assetIDRule = append(assetIDRule, assetIDItem)
	}

	logs, sub, err := _AssetUpperBound.contract.WatchLogs(opts, "UpperBoundSet", assetIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetUpperBoundUpperBoundSet)
				if err := _AssetUpperBound.contract.UnpackLog(event, "UpperBoundSet", log); err != nil {
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

// ParseUpperBoundSet is a log parse operation binding the contract event 0x4745108fe432eb4e630177b0117dad9218e79b2a93d7d87d375676c9d05f8234.
//
// Solidity: event UpperBoundSet(uint256 indexed assetID, uint256 upperBound)
func (_AssetUpperBound *AssetUpperBoundFilterer) ParseUpperBoundSet(log types.Log) (*AssetUpperBoundUpperBoundSet, error) {
	event := new(AssetUpperBoundUpperBoundSet)
	if err := _AssetUpperBound.contract.UnpackLog(event, "UpperBoundSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
