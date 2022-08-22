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

// ValidatorRegistryABI is the input ABI used to generate the binding from.
const ValidatorRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"genre\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"Registration\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"files\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"genre\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getFile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_genre\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ValidatorRegistry is an auto generated Go binding around an Ethereum contract.
type ValidatorRegistry struct {
	ValidatorRegistryCaller     // Read-only binding to the contract
	ValidatorRegistryTransactor // Write-only binding to the contract
	ValidatorRegistryFilterer   // Log filterer for contract events
}

// ValidatorRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorRegistrySession struct {
	Contract     *ValidatorRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorRegistryCallerSession struct {
	Contract *ValidatorRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorRegistryTransactorSession struct {
	Contract     *ValidatorRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRegistryRaw struct {
	Contract *ValidatorRegistry // Generic contract binding to access the raw methods on
}

// ValidatorRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorRegistryCallerRaw struct {
	Contract *ValidatorRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorRegistryTransactorRaw struct {
	Contract *ValidatorRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorRegistry creates a new instance of ValidatorRegistry, bound to a specific deployed contract.
func NewValidatorRegistry(address common.Address, backend bind.ContractBackend) (*ValidatorRegistry, error) {
	contract, err := bindValidatorRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistry{ValidatorRegistryCaller: ValidatorRegistryCaller{contract: contract}, ValidatorRegistryTransactor: ValidatorRegistryTransactor{contract: contract}, ValidatorRegistryFilterer: ValidatorRegistryFilterer{contract: contract}}, nil
}

// NewValidatorRegistryCaller creates a new read-only instance of ValidatorRegistry, bound to a specific deployed contract.
func NewValidatorRegistryCaller(address common.Address, caller bind.ContractCaller) (*ValidatorRegistryCaller, error) {
	contract, err := bindValidatorRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryCaller{contract: contract}, nil
}

// NewValidatorRegistryTransactor creates a new write-only instance of ValidatorRegistry, bound to a specific deployed contract.
func NewValidatorRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorRegistryTransactor, error) {
	contract, err := bindValidatorRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryTransactor{contract: contract}, nil
}

// NewValidatorRegistryFilterer creates a new log filterer instance of ValidatorRegistry, bound to a specific deployed contract.
func NewValidatorRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorRegistryFilterer, error) {
	contract, err := bindValidatorRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryFilterer{contract: contract}, nil
}

// bindValidatorRegistry binds a generic wrapper to an already deployed contract.
func bindValidatorRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorRegistry *ValidatorRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorRegistry.Contract.ValidatorRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorRegistry *ValidatorRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.ValidatorRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorRegistry *ValidatorRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.ValidatorRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorRegistry *ValidatorRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorRegistry *ValidatorRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorRegistry *ValidatorRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.contract.Transact(opts, method, params...)
}

// Files is a free data retrieval call binding the contract method 0x39c637df.
//
// Solidity: function files(address ) view returns(address validator, uint256 genre, string uri)
func (_ValidatorRegistry *ValidatorRegistryCaller) Files(opts *bind.CallOpts, arg0 common.Address) (struct {
	Validator common.Address
	Genre     *big.Int
	Uri       string
}, error) {
	var out []interface{}
	err := _ValidatorRegistry.contract.Call(opts, &out, "files", arg0)

	outstruct := new(struct {
		Validator common.Address
		Genre     *big.Int
		Uri       string
	})

	outstruct.Validator = out[0].(common.Address)
	outstruct.Genre = out[1].(*big.Int)
	outstruct.Uri = out[2].(string)

	return *outstruct, err

}

// Files is a free data retrieval call binding the contract method 0x39c637df.
//
// Solidity: function files(address ) view returns(address validator, uint256 genre, string uri)
func (_ValidatorRegistry *ValidatorRegistrySession) Files(arg0 common.Address) (struct {
	Validator common.Address
	Genre     *big.Int
	Uri       string
}, error) {
	return _ValidatorRegistry.Contract.Files(&_ValidatorRegistry.CallOpts, arg0)
}

// Files is a free data retrieval call binding the contract method 0x39c637df.
//
// Solidity: function files(address ) view returns(address validator, uint256 genre, string uri)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) Files(arg0 common.Address) (struct {
	Validator common.Address
	Genre     *big.Int
	Uri       string
}, error) {
	return _ValidatorRegistry.Contract.Files(&_ValidatorRegistry.CallOpts, arg0)
}

// GetFile is a free data retrieval call binding the contract method 0x0b87c1e0.
//
// Solidity: function getFile(address _validator) view returns(uint256, string)
func (_ValidatorRegistry *ValidatorRegistryCaller) GetFile(opts *bind.CallOpts, _validator common.Address) (*big.Int, string, error) {
	var out []interface{}
	err := _ValidatorRegistry.contract.Call(opts, &out, "getFile", _validator)

	if err != nil {
		return *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetFile is a free data retrieval call binding the contract method 0x0b87c1e0.
//
// Solidity: function getFile(address _validator) view returns(uint256, string)
func (_ValidatorRegistry *ValidatorRegistrySession) GetFile(_validator common.Address) (*big.Int, string, error) {
	return _ValidatorRegistry.Contract.GetFile(&_ValidatorRegistry.CallOpts, _validator)
}

// GetFile is a free data retrieval call binding the contract method 0x0b87c1e0.
//
// Solidity: function getFile(address _validator) view returns(uint256, string)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) GetFile(_validator common.Address) (*big.Int, string, error) {
	return _ValidatorRegistry.Contract.GetFile(&_ValidatorRegistry.CallOpts, _validator)
}

// Register is a paid mutator transaction binding the contract method 0xa00fd3c8.
//
// Solidity: function register(uint256 _genre, string _uri) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactor) Register(opts *bind.TransactOpts, _genre *big.Int, _uri string) (*types.Transaction, error) {
	return _ValidatorRegistry.contract.Transact(opts, "register", _genre, _uri)
}

// Register is a paid mutator transaction binding the contract method 0xa00fd3c8.
//
// Solidity: function register(uint256 _genre, string _uri) returns()
func (_ValidatorRegistry *ValidatorRegistrySession) Register(_genre *big.Int, _uri string) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.Register(&_ValidatorRegistry.TransactOpts, _genre, _uri)
}

// Register is a paid mutator transaction binding the contract method 0xa00fd3c8.
//
// Solidity: function register(uint256 _genre, string _uri) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactorSession) Register(_genre *big.Int, _uri string) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.Register(&_ValidatorRegistry.TransactOpts, _genre, _uri)
}

// ValidatorRegistryRegistrationIterator is returned from FilterRegistration and is used to iterate over the raw logs and unpacked data for Registration events raised by the ValidatorRegistry contract.
type ValidatorRegistryRegistrationIterator struct {
	Event *ValidatorRegistryRegistration // Event containing the contract specifics and raw log

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
func (it *ValidatorRegistryRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRegistryRegistration)
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
		it.Event = new(ValidatorRegistryRegistration)
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
func (it *ValidatorRegistryRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRegistryRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRegistryRegistration represents a Registration event raised by the ValidatorRegistry contract.
type ValidatorRegistryRegistration struct {
	Validator common.Address
	Genre     *big.Int
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegistration is a free log retrieval operation binding the contract event 0xeee9ae735c82c3067b614d28068eed722045ff14935ee256c8654db2ff75c782.
//
// Solidity: event Registration(address indexed validator, uint256 indexed genre, string uri)
func (_ValidatorRegistry *ValidatorRegistryFilterer) FilterRegistration(opts *bind.FilterOpts, validator []common.Address, genre []*big.Int) (*ValidatorRegistryRegistrationIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var genreRule []interface{}
	for _, genreItem := range genre {
		genreRule = append(genreRule, genreItem)
	}

	logs, sub, err := _ValidatorRegistry.contract.FilterLogs(opts, "Registration", validatorRule, genreRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryRegistrationIterator{contract: _ValidatorRegistry.contract, event: "Registration", logs: logs, sub: sub}, nil
}

// WatchRegistration is a free log subscription operation binding the contract event 0xeee9ae735c82c3067b614d28068eed722045ff14935ee256c8654db2ff75c782.
//
// Solidity: event Registration(address indexed validator, uint256 indexed genre, string uri)
func (_ValidatorRegistry *ValidatorRegistryFilterer) WatchRegistration(opts *bind.WatchOpts, sink chan<- *ValidatorRegistryRegistration, validator []common.Address, genre []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var genreRule []interface{}
	for _, genreItem := range genre {
		genreRule = append(genreRule, genreItem)
	}

	logs, sub, err := _ValidatorRegistry.contract.WatchLogs(opts, "Registration", validatorRule, genreRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRegistryRegistration)
				if err := _ValidatorRegistry.contract.UnpackLog(event, "Registration", log); err != nil {
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

// ParseRegistration is a log parse operation binding the contract event 0xeee9ae735c82c3067b614d28068eed722045ff14935ee256c8654db2ff75c782.
//
// Solidity: event Registration(address indexed validator, uint256 indexed genre, string uri)
func (_ValidatorRegistry *ValidatorRegistryFilterer) ParseRegistration(log types.Log) (*ValidatorRegistryRegistration, error) {
	event := new(ValidatorRegistryRegistration)
	if err := _ValidatorRegistry.contract.UnpackLog(event, "Registration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
