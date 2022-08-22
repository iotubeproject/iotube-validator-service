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

// TestimonyDAOABI is the input ABI used to generate the binding from.
const TestimonyDAOABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"witness\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"testimony\",\"type\":\"bytes\"}],\"name\":\"Testimony\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"addTestimony\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestimonyDAO is an auto generated Go binding around an Ethereum contract.
type TestimonyDAO struct {
	TestimonyDAOCaller     // Read-only binding to the contract
	TestimonyDAOTransactor // Write-only binding to the contract
	TestimonyDAOFilterer   // Log filterer for contract events
}

// TestimonyDAOCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestimonyDAOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestimonyDAOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestimonyDAOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestimonyDAOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestimonyDAOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestimonyDAOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestimonyDAOSession struct {
	Contract     *TestimonyDAO     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestimonyDAOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestimonyDAOCallerSession struct {
	Contract *TestimonyDAOCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TestimonyDAOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestimonyDAOTransactorSession struct {
	Contract     *TestimonyDAOTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TestimonyDAORaw is an auto generated low-level Go binding around an Ethereum contract.
type TestimonyDAORaw struct {
	Contract *TestimonyDAO // Generic contract binding to access the raw methods on
}

// TestimonyDAOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestimonyDAOCallerRaw struct {
	Contract *TestimonyDAOCaller // Generic read-only contract binding to access the raw methods on
}

// TestimonyDAOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestimonyDAOTransactorRaw struct {
	Contract *TestimonyDAOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestimonyDAO creates a new instance of TestimonyDAO, bound to a specific deployed contract.
func NewTestimonyDAO(address common.Address, backend bind.ContractBackend) (*TestimonyDAO, error) {
	contract, err := bindTestimonyDAO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestimonyDAO{TestimonyDAOCaller: TestimonyDAOCaller{contract: contract}, TestimonyDAOTransactor: TestimonyDAOTransactor{contract: contract}, TestimonyDAOFilterer: TestimonyDAOFilterer{contract: contract}}, nil
}

// NewTestimonyDAOCaller creates a new read-only instance of TestimonyDAO, bound to a specific deployed contract.
func NewTestimonyDAOCaller(address common.Address, caller bind.ContractCaller) (*TestimonyDAOCaller, error) {
	contract, err := bindTestimonyDAO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestimonyDAOCaller{contract: contract}, nil
}

// NewTestimonyDAOTransactor creates a new write-only instance of TestimonyDAO, bound to a specific deployed contract.
func NewTestimonyDAOTransactor(address common.Address, transactor bind.ContractTransactor) (*TestimonyDAOTransactor, error) {
	contract, err := bindTestimonyDAO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestimonyDAOTransactor{contract: contract}, nil
}

// NewTestimonyDAOFilterer creates a new log filterer instance of TestimonyDAO, bound to a specific deployed contract.
func NewTestimonyDAOFilterer(address common.Address, filterer bind.ContractFilterer) (*TestimonyDAOFilterer, error) {
	contract, err := bindTestimonyDAO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestimonyDAOFilterer{contract: contract}, nil
}

// bindTestimonyDAO binds a generic wrapper to an already deployed contract.
func bindTestimonyDAO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestimonyDAOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestimonyDAO *TestimonyDAORaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestimonyDAO.Contract.TestimonyDAOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestimonyDAO *TestimonyDAORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestimonyDAO.Contract.TestimonyDAOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestimonyDAO *TestimonyDAORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestimonyDAO.Contract.TestimonyDAOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestimonyDAO *TestimonyDAOCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestimonyDAO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestimonyDAO *TestimonyDAOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestimonyDAO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestimonyDAO *TestimonyDAOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestimonyDAO.Contract.contract.Transact(opts, method, params...)
}

// AddTestimony is a paid mutator transaction binding the contract method 0x513bd021.
//
// Solidity: function addTestimony(bytes32 key, bytes value) returns()
func (_TestimonyDAO *TestimonyDAOTransactor) AddTestimony(opts *bind.TransactOpts, key [32]byte, value []byte) (*types.Transaction, error) {
	return _TestimonyDAO.contract.Transact(opts, "addTestimony", key, value)
}

// AddTestimony is a paid mutator transaction binding the contract method 0x513bd021.
//
// Solidity: function addTestimony(bytes32 key, bytes value) returns()
func (_TestimonyDAO *TestimonyDAOSession) AddTestimony(key [32]byte, value []byte) (*types.Transaction, error) {
	return _TestimonyDAO.Contract.AddTestimony(&_TestimonyDAO.TransactOpts, key, value)
}

// AddTestimony is a paid mutator transaction binding the contract method 0x513bd021.
//
// Solidity: function addTestimony(bytes32 key, bytes value) returns()
func (_TestimonyDAO *TestimonyDAOTransactorSession) AddTestimony(key [32]byte, value []byte) (*types.Transaction, error) {
	return _TestimonyDAO.Contract.AddTestimony(&_TestimonyDAO.TransactOpts, key, value)
}

// TestimonyDAOTestimonyIterator is returned from FilterTestimony and is used to iterate over the raw logs and unpacked data for Testimony events raised by the TestimonyDAO contract.
type TestimonyDAOTestimonyIterator struct {
	Event *TestimonyDAOTestimony // Event containing the contract specifics and raw log

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
func (it *TestimonyDAOTestimonyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestimonyDAOTestimony)
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
		it.Event = new(TestimonyDAOTestimony)
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
func (it *TestimonyDAOTestimonyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestimonyDAOTestimonyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestimonyDAOTestimony represents a Testimony event raised by the TestimonyDAO contract.
type TestimonyDAOTestimony struct {
	Witness   common.Address
	Key       [32]byte
	Testimony []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTestimony is a free log retrieval operation binding the contract event 0x738e5ba8c1c11648ac3cb5c81364c2ed48f5af4b2937b323a0e1f4cce15801cf.
//
// Solidity: event Testimony(address indexed witness, bytes32 indexed key, bytes testimony)
func (_TestimonyDAO *TestimonyDAOFilterer) FilterTestimony(opts *bind.FilterOpts, witness []common.Address, key [][32]byte) (*TestimonyDAOTestimonyIterator, error) {

	var witnessRule []interface{}
	for _, witnessItem := range witness {
		witnessRule = append(witnessRule, witnessItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _TestimonyDAO.contract.FilterLogs(opts, "Testimony", witnessRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &TestimonyDAOTestimonyIterator{contract: _TestimonyDAO.contract, event: "Testimony", logs: logs, sub: sub}, nil
}

// WatchTestimony is a free log subscription operation binding the contract event 0x738e5ba8c1c11648ac3cb5c81364c2ed48f5af4b2937b323a0e1f4cce15801cf.
//
// Solidity: event Testimony(address indexed witness, bytes32 indexed key, bytes testimony)
func (_TestimonyDAO *TestimonyDAOFilterer) WatchTestimony(opts *bind.WatchOpts, sink chan<- *TestimonyDAOTestimony, witness []common.Address, key [][32]byte) (event.Subscription, error) {

	var witnessRule []interface{}
	for _, witnessItem := range witness {
		witnessRule = append(witnessRule, witnessItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _TestimonyDAO.contract.WatchLogs(opts, "Testimony", witnessRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestimonyDAOTestimony)
				if err := _TestimonyDAO.contract.UnpackLog(event, "Testimony", log); err != nil {
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

// ParseTestimony is a log parse operation binding the contract event 0x738e5ba8c1c11648ac3cb5c81364c2ed48f5af4b2937b323a0e1f4cce15801cf.
//
// Solidity: event Testimony(address indexed witness, bytes32 indexed key, bytes testimony)
func (_TestimonyDAO *TestimonyDAOFilterer) ParseTestimony(log types.Log) (*TestimonyDAOTestimony, error) {
	event := new(TestimonyDAOTestimony)
	if err := _TestimonyDAO.contract.UnpackLog(event, "Testimony", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
