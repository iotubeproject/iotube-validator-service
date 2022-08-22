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

// ERC20TubeTubeInfo is an auto generated low-level Go binding around an user-defined struct.
type ERC20TubeTubeInfo struct {
	Rate    *big.Int
	Enabled bool
}

// ERC20TubeABI is the input ABI used to generate the binding from.
const ERC20TubeABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"},{\"internalType\":\"contractILedger\",\"name\":\"_ledger\",\"type\":\"address\"},{\"internalType\":\"contractILord\",\"name\":\"_lord\",\"type\":\"address\"},{\"internalType\":\"contractIVerifier\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_safe\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"EmergencyOperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lord\",\"type\":\"address\"}],\"name\":\"LordUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"targetTubeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Receipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"safe\",\"type\":\"address\"}],\"name\":\"SafeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"Settled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tubeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"TubeInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetTubeID\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetTubeID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"}],\"name\":\"destinationTubeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structERC20Tube.TubeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcTubeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"genKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"isSettled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ledger\",\"outputs\":[{\"internalType\":\"contractILedger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lord\",\"outputs\":[{\"internalType\":\"contractILord\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"safe\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tubeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setDestinationTube\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setEmergencyOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILord\",\"name\":\"_lord\",\"type\":\"address\"}],\"name\":\"setLord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_safe\",\"type\":\"address\"}],\"name\":\"setSafe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tubeID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcTubeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20Tube is an auto generated Go binding around an Ethereum contract.
type ERC20Tube struct {
	ERC20TubeCaller     // Read-only binding to the contract
	ERC20TubeTransactor // Write-only binding to the contract
	ERC20TubeFilterer   // Log filterer for contract events
}

// ERC20TubeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TubeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TubeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TubeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TubeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TubeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TubeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TubeSession struct {
	Contract     *ERC20Tube        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20TubeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TubeCallerSession struct {
	Contract *ERC20TubeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ERC20TubeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TubeTransactorSession struct {
	Contract     *ERC20TubeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20TubeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TubeRaw struct {
	Contract *ERC20Tube // Generic contract binding to access the raw methods on
}

// ERC20TubeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TubeCallerRaw struct {
	Contract *ERC20TubeCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TubeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TubeTransactorRaw struct {
	Contract *ERC20TubeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Tube creates a new instance of ERC20Tube, bound to a specific deployed contract.
func NewERC20Tube(address common.Address, backend bind.ContractBackend) (*ERC20Tube, error) {
	contract, err := bindERC20Tube(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Tube{ERC20TubeCaller: ERC20TubeCaller{contract: contract}, ERC20TubeTransactor: ERC20TubeTransactor{contract: contract}, ERC20TubeFilterer: ERC20TubeFilterer{contract: contract}}, nil
}

// NewERC20TubeCaller creates a new read-only instance of ERC20Tube, bound to a specific deployed contract.
func NewERC20TubeCaller(address common.Address, caller bind.ContractCaller) (*ERC20TubeCaller, error) {
	contract, err := bindERC20Tube(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeCaller{contract: contract}, nil
}

// NewERC20TubeTransactor creates a new write-only instance of ERC20Tube, bound to a specific deployed contract.
func NewERC20TubeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TubeTransactor, error) {
	contract, err := bindERC20Tube(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeTransactor{contract: contract}, nil
}

// NewERC20TubeFilterer creates a new log filterer instance of ERC20Tube, bound to a specific deployed contract.
func NewERC20TubeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TubeFilterer, error) {
	contract, err := bindERC20Tube(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeFilterer{contract: contract}, nil
}

// bindERC20Tube binds a generic wrapper to an already deployed contract.
func bindERC20Tube(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20TubeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Tube *ERC20TubeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Tube.Contract.ERC20TubeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Tube *ERC20TubeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Tube.Contract.ERC20TubeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Tube *ERC20TubeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Tube.Contract.ERC20TubeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Tube *ERC20TubeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Tube.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Tube *ERC20TubeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Tube.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Tube *ERC20TubeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Tube.Contract.contract.Transact(opts, method, params...)
}

// DestinationTubeInfo is a free data retrieval call binding the contract method 0x8b1ef73a.
//
// Solidity: function destinationTubeInfo(uint256 _tubeID) view returns((uint256,bool))
func (_ERC20Tube *ERC20TubeCaller) DestinationTubeInfo(opts *bind.CallOpts, _tubeID *big.Int) (ERC20TubeTubeInfo, error) {
	var out []interface{}
	err := _ERC20Tube.contract.Call(opts, &out, "destinationTubeInfo", _tubeID)

	if err != nil {
		return *new(ERC20TubeTubeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC20TubeTubeInfo)).(*ERC20TubeTubeInfo)

	return out0, err

}

// DestinationTubeInfo is a free data retrieval call binding the contract method 0x8b1ef73a.
//
// Solidity: function destinationTubeInfo(uint256 _tubeID) view returns((uint256,bool))
func (_ERC20Tube *ERC20TubeSession) DestinationTubeInfo(_tubeID *big.Int) (ERC20TubeTubeInfo, error) {
	return _ERC20Tube.Contract.DestinationTubeInfo(&_ERC20Tube.CallOpts, _tubeID)
}

// DestinationTubeInfo is a free data retrieval call binding the contract method 0x8b1ef73a.
//
// Solidity: function destinationTubeInfo(uint256 _tubeID) view returns((uint256,bool))
func (_ERC20Tube *ERC20TubeCallerSession) DestinationTubeInfo(_tubeID *big.Int) (ERC20TubeTubeInfo, error) {
	return _ERC20Tube.Contract.DestinationTubeInfo(&_ERC20Tube.CallOpts, _tubeID)
}

// GenKey is a free data retrieval call binding the contract method 0xc015cdf2.
//
// Solidity: function genKey(uint256 _srcTubeID, uint256 _nonce, address _token, address _recipient, uint256 _amount) view returns(bytes32)
func (_ERC20Tube *ERC20TubeCaller) GenKey(opts *bind.CallOpts, _srcTubeID *big.Int, _nonce *big.Int, _token common.Address, _recipient common.Address, _amount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Tube.contract.Call(opts, &out, "genKey", _srcTubeID, _nonce, _token, _recipient, _amount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GenKey is a free data retrieval call binding the contract method 0xc015cdf2.
//
// Solidity: function genKey(uint256 _srcTubeID, uint256 _nonce, address _token, address _recipient, uint256 _amount) view returns(bytes32)
func (_ERC20Tube *ERC20TubeSession) GenKey(_srcTubeID *big.Int, _nonce *big.Int, _token common.Address, _recipient common.Address, _amount *big.Int) ([32]byte, error) {
	return _ERC20Tube.Contract.GenKey(&_ERC20Tube.CallOpts, _srcTubeID, _nonce, _token, _recipient, _amount)
}

// GenKey is a free data retrieval call binding the contract method 0xc015cdf2.
//
// Solidity: function genKey(uint256 _srcTubeID, uint256 _nonce, address _token, address _recipient, uint256 _amount) view returns(bytes32)
func (_ERC20Tube *ERC20TubeCallerSession) GenKey(_srcTubeID *big.Int, _nonce *big.Int, _token common.Address, _recipient common.Address, _amount *big.Int) ([32]byte, error) {
	return _ERC20Tube.Contract.GenKey(&_ERC20Tube.CallOpts, _srcTubeID, _nonce, _token, _recipient, _amount)
}

// IsSettled is a free data retrieval call binding the contract method 0xbd07f3c9.
//
// Solidity: function isSettled(bytes32 key) view returns(bool)
func (_ERC20Tube *ERC20TubeCaller) IsSettled(opts *bind.CallOpts, key [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC20Tube.contract.Call(opts, &out, "isSettled", key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSettled is a free data retrieval call binding the contract method 0xbd07f3c9.
//
// Solidity: function isSettled(bytes32 key) view returns(bool)
func (_ERC20Tube *ERC20TubeSession) IsSettled(key [32]byte) (bool, error) {
	return _ERC20Tube.Contract.IsSettled(&_ERC20Tube.CallOpts, key)
}

// IsSettled is a free data retrieval call binding the contract method 0xbd07f3c9.
//
// Solidity: function isSettled(bytes32 key) view returns(bool)
func (_ERC20Tube *ERC20TubeCallerSession) IsSettled(key [32]byte) (bool, error) {
	return _ERC20Tube.Contract.IsSettled(&_ERC20Tube.CallOpts, key)
}

// Ledger is a free data retrieval call binding the contract method 0x56397c35.
//
// Solidity: function ledger() view returns(address)
func (_ERC20Tube *ERC20TubeCaller) Ledger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Tube.contract.Call(opts, &out, "ledger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ledger is a free data retrieval call binding the contract method 0x56397c35.
//
// Solidity: function ledger() view returns(address)
func (_ERC20Tube *ERC20TubeSession) Ledger() (common.Address, error) {
	return _ERC20Tube.Contract.Ledger(&_ERC20Tube.CallOpts)
}

// Ledger is a free data retrieval call binding the contract method 0x56397c35.
//
// Solidity: function ledger() view returns(address)
func (_ERC20Tube *ERC20TubeCallerSession) Ledger() (common.Address, error) {
	return _ERC20Tube.Contract.Ledger(&_ERC20Tube.CallOpts)
}

// Lord is a free data retrieval call binding the contract method 0x78353d9a.
//
// Solidity: function lord() view returns(address)
func (_ERC20Tube *ERC20TubeCaller) Lord(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Tube.contract.Call(opts, &out, "lord")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lord is a free data retrieval call binding the contract method 0x78353d9a.
//
// Solidity: function lord() view returns(address)
func (_ERC20Tube *ERC20TubeSession) Lord() (common.Address, error) {
	return _ERC20Tube.Contract.Lord(&_ERC20Tube.CallOpts)
}

// Lord is a free data retrieval call binding the contract method 0x78353d9a.
//
// Solidity: function lord() view returns(address)
func (_ERC20Tube *ERC20TubeCallerSession) Lord() (common.Address, error) {
	return _ERC20Tube.Contract.Lord(&_ERC20Tube.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_ERC20Tube *ERC20TubeCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Tube.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_ERC20Tube *ERC20TubeSession) Nonce() (*big.Int, error) {
	return _ERC20Tube.Contract.Nonce(&_ERC20Tube.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_ERC20Tube *ERC20TubeCallerSession) Nonce() (*big.Int, error) {
	return _ERC20Tube.Contract.Nonce(&_ERC20Tube.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Tube *ERC20TubeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Tube.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Tube *ERC20TubeSession) Owner() (common.Address, error) {
	return _ERC20Tube.Contract.Owner(&_ERC20Tube.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Tube *ERC20TubeCallerSession) Owner() (common.Address, error) {
	return _ERC20Tube.Contract.Owner(&_ERC20Tube.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Tube *ERC20TubeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Tube.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Tube *ERC20TubeSession) Paused() (bool, error) {
	return _ERC20Tube.Contract.Paused(&_ERC20Tube.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Tube *ERC20TubeCallerSession) Paused() (bool, error) {
	return _ERC20Tube.Contract.Paused(&_ERC20Tube.CallOpts)
}

// Safe is a free data retrieval call binding the contract method 0x186f0354.
//
// Solidity: function safe() view returns(address)
func (_ERC20Tube *ERC20TubeCaller) Safe(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Tube.contract.Call(opts, &out, "safe")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Safe is a free data retrieval call binding the contract method 0x186f0354.
//
// Solidity: function safe() view returns(address)
func (_ERC20Tube *ERC20TubeSession) Safe() (common.Address, error) {
	return _ERC20Tube.Contract.Safe(&_ERC20Tube.CallOpts)
}

// Safe is a free data retrieval call binding the contract method 0x186f0354.
//
// Solidity: function safe() view returns(address)
func (_ERC20Tube *ERC20TubeCallerSession) Safe() (common.Address, error) {
	return _ERC20Tube.Contract.Safe(&_ERC20Tube.CallOpts)
}

// TubeID is a free data retrieval call binding the contract method 0x87622b93.
//
// Solidity: function tubeID() view returns(uint256)
func (_ERC20Tube *ERC20TubeCaller) TubeID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Tube.contract.Call(opts, &out, "tubeID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TubeID is a free data retrieval call binding the contract method 0x87622b93.
//
// Solidity: function tubeID() view returns(uint256)
func (_ERC20Tube *ERC20TubeSession) TubeID() (*big.Int, error) {
	return _ERC20Tube.Contract.TubeID(&_ERC20Tube.CallOpts)
}

// TubeID is a free data retrieval call binding the contract method 0x87622b93.
//
// Solidity: function tubeID() view returns(uint256)
func (_ERC20Tube *ERC20TubeCallerSession) TubeID() (*big.Int, error) {
	return _ERC20Tube.Contract.TubeID(&_ERC20Tube.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ERC20Tube *ERC20TubeCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Tube.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ERC20Tube *ERC20TubeSession) Verifier() (common.Address, error) {
	return _ERC20Tube.Contract.Verifier(&_ERC20Tube.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ERC20Tube *ERC20TubeCallerSession) Verifier() (common.Address, error) {
	return _ERC20Tube.Contract.Verifier(&_ERC20Tube.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address _token, uint256 _amount, uint256 _targetTubeID) returns()
func (_ERC20Tube *ERC20TubeTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _targetTubeID *big.Int) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "deposit", _token, _amount, _targetTubeID)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address _token, uint256 _amount, uint256 _targetTubeID) returns()
func (_ERC20Tube *ERC20TubeSession) Deposit(_token common.Address, _amount *big.Int, _targetTubeID *big.Int) (*types.Transaction, error) {
	return _ERC20Tube.Contract.Deposit(&_ERC20Tube.TransactOpts, _token, _amount, _targetTubeID)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address _token, uint256 _amount, uint256 _targetTubeID) returns()
func (_ERC20Tube *ERC20TubeTransactorSession) Deposit(_token common.Address, _amount *big.Int, _targetTubeID *big.Int) (*types.Transaction, error) {
	return _ERC20Tube.Contract.Deposit(&_ERC20Tube.TransactOpts, _token, _amount, _targetTubeID)
}

// DepositTo is a paid mutator transaction binding the contract method 0x5379a5cb.
//
// Solidity: function depositTo(address _token, uint256 _amount, uint256 _targetTubeID, address _to) returns()
func (_ERC20Tube *ERC20TubeTransactor) DepositTo(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _targetTubeID *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "depositTo", _token, _amount, _targetTubeID, _to)
}

// DepositTo is a paid mutator transaction binding the contract method 0x5379a5cb.
//
// Solidity: function depositTo(address _token, uint256 _amount, uint256 _targetTubeID, address _to) returns()
func (_ERC20Tube *ERC20TubeSession) DepositTo(_token common.Address, _amount *big.Int, _targetTubeID *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.DepositTo(&_ERC20Tube.TransactOpts, _token, _amount, _targetTubeID, _to)
}

// DepositTo is a paid mutator transaction binding the contract method 0x5379a5cb.
//
// Solidity: function depositTo(address _token, uint256 _amount, uint256 _targetTubeID, address _to) returns()
func (_ERC20Tube *ERC20TubeTransactorSession) DepositTo(_token common.Address, _amount *big.Int, _targetTubeID *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.DepositTo(&_ERC20Tube.TransactOpts, _token, _amount, _targetTubeID, _to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Tube *ERC20TubeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Tube *ERC20TubeSession) Pause() (*types.Transaction, error) {
	return _ERC20Tube.Contract.Pause(&_ERC20Tube.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Tube *ERC20TubeTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20Tube.Contract.Pause(&_ERC20Tube.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Tube *ERC20TubeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Tube *ERC20TubeSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20Tube.Contract.RenounceOwnership(&_ERC20Tube.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Tube *ERC20TubeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20Tube.Contract.RenounceOwnership(&_ERC20Tube.TransactOpts)
}

// SetDestinationTube is a paid mutator transaction binding the contract method 0xa6d1ebe1.
//
// Solidity: function setDestinationTube(uint256 _tubeID, uint256 _feeRate, bool _enabled) returns()
func (_ERC20Tube *ERC20TubeTransactor) SetDestinationTube(opts *bind.TransactOpts, _tubeID *big.Int, _feeRate *big.Int, _enabled bool) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "setDestinationTube", _tubeID, _feeRate, _enabled)
}

// SetDestinationTube is a paid mutator transaction binding the contract method 0xa6d1ebe1.
//
// Solidity: function setDestinationTube(uint256 _tubeID, uint256 _feeRate, bool _enabled) returns()
func (_ERC20Tube *ERC20TubeSession) SetDestinationTube(_tubeID *big.Int, _feeRate *big.Int, _enabled bool) (*types.Transaction, error) {
	return _ERC20Tube.Contract.SetDestinationTube(&_ERC20Tube.TransactOpts, _tubeID, _feeRate, _enabled)
}

// SetDestinationTube is a paid mutator transaction binding the contract method 0xa6d1ebe1.
//
// Solidity: function setDestinationTube(uint256 _tubeID, uint256 _feeRate, bool _enabled) returns()
func (_ERC20Tube *ERC20TubeTransactorSession) SetDestinationTube(_tubeID *big.Int, _feeRate *big.Int, _enabled bool) (*types.Transaction, error) {
	return _ERC20Tube.Contract.SetDestinationTube(&_ERC20Tube.TransactOpts, _tubeID, _feeRate, _enabled)
}

// SetEmergencyOperator is a paid mutator transaction binding the contract method 0x0958654c.
//
// Solidity: function setEmergencyOperator(address _operator) returns()
func (_ERC20Tube *ERC20TubeTransactor) SetEmergencyOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "setEmergencyOperator", _operator)
}

// SetEmergencyOperator is a paid mutator transaction binding the contract method 0x0958654c.
//
// Solidity: function setEmergencyOperator(address _operator) returns()
func (_ERC20Tube *ERC20TubeSession) SetEmergencyOperator(_operator common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.SetEmergencyOperator(&_ERC20Tube.TransactOpts, _operator)
}

// SetEmergencyOperator is a paid mutator transaction binding the contract method 0x0958654c.
//
// Solidity: function setEmergencyOperator(address _operator) returns()
func (_ERC20Tube *ERC20TubeTransactorSession) SetEmergencyOperator(_operator common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.SetEmergencyOperator(&_ERC20Tube.TransactOpts, _operator)
}

// SetLord is a paid mutator transaction binding the contract method 0xa2d49351.
//
// Solidity: function setLord(address _lord) returns()
func (_ERC20Tube *ERC20TubeTransactor) SetLord(opts *bind.TransactOpts, _lord common.Address) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "setLord", _lord)
}

// SetLord is a paid mutator transaction binding the contract method 0xa2d49351.
//
// Solidity: function setLord(address _lord) returns()
func (_ERC20Tube *ERC20TubeSession) SetLord(_lord common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.SetLord(&_ERC20Tube.TransactOpts, _lord)
}

// SetLord is a paid mutator transaction binding the contract method 0xa2d49351.
//
// Solidity: function setLord(address _lord) returns()
func (_ERC20Tube *ERC20TubeTransactorSession) SetLord(_lord common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.SetLord(&_ERC20Tube.TransactOpts, _lord)
}

// SetSafe is a paid mutator transaction binding the contract method 0x5db0cb94.
//
// Solidity: function setSafe(address _safe) returns()
func (_ERC20Tube *ERC20TubeTransactor) SetSafe(opts *bind.TransactOpts, _safe common.Address) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "setSafe", _safe)
}

// SetSafe is a paid mutator transaction binding the contract method 0x5db0cb94.
//
// Solidity: function setSafe(address _safe) returns()
func (_ERC20Tube *ERC20TubeSession) SetSafe(_safe common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.SetSafe(&_ERC20Tube.TransactOpts, _safe)
}

// SetSafe is a paid mutator transaction binding the contract method 0x5db0cb94.
//
// Solidity: function setSafe(address _safe) returns()
func (_ERC20Tube *ERC20TubeTransactorSession) SetSafe(_safe common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.SetSafe(&_ERC20Tube.TransactOpts, _safe)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Tube *ERC20TubeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Tube *ERC20TubeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.TransferOwnership(&_ERC20Tube.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Tube *ERC20TubeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.TransferOwnership(&_ERC20Tube.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20Tube *ERC20TubeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20Tube *ERC20TubeSession) Unpause() (*types.Transaction, error) {
	return _ERC20Tube.Contract.Unpause(&_ERC20Tube.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20Tube *ERC20TubeTransactorSession) Unpause() (*types.Transaction, error) {
	return _ERC20Tube.Contract.Unpause(&_ERC20Tube.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x277b5c7a.
//
// Solidity: function withdraw(uint256 _srcTubeID, uint256 _nonce, address _token, address _recipient, uint256 _amount, bytes _signatures) returns()
func (_ERC20Tube *ERC20TubeTransactor) Withdraw(opts *bind.TransactOpts, _srcTubeID *big.Int, _nonce *big.Int, _token common.Address, _recipient common.Address, _amount *big.Int, _signatures []byte) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "withdraw", _srcTubeID, _nonce, _token, _recipient, _amount, _signatures)
}

// Withdraw is a paid mutator transaction binding the contract method 0x277b5c7a.
//
// Solidity: function withdraw(uint256 _srcTubeID, uint256 _nonce, address _token, address _recipient, uint256 _amount, bytes _signatures) returns()
func (_ERC20Tube *ERC20TubeSession) Withdraw(_srcTubeID *big.Int, _nonce *big.Int, _token common.Address, _recipient common.Address, _amount *big.Int, _signatures []byte) (*types.Transaction, error) {
	return _ERC20Tube.Contract.Withdraw(&_ERC20Tube.TransactOpts, _srcTubeID, _nonce, _token, _recipient, _amount, _signatures)
}

// Withdraw is a paid mutator transaction binding the contract method 0x277b5c7a.
//
// Solidity: function withdraw(uint256 _srcTubeID, uint256 _nonce, address _token, address _recipient, uint256 _amount, bytes _signatures) returns()
func (_ERC20Tube *ERC20TubeTransactorSession) Withdraw(_srcTubeID *big.Int, _nonce *big.Int, _token common.Address, _recipient common.Address, _amount *big.Int, _signatures []byte) (*types.Transaction, error) {
	return _ERC20Tube.Contract.Withdraw(&_ERC20Tube.TransactOpts, _srcTubeID, _nonce, _token, _recipient, _amount, _signatures)
}

// WithdrawCoin is a paid mutator transaction binding the contract method 0x5a73928f.
//
// Solidity: function withdrawCoin(address _to) returns()
func (_ERC20Tube *ERC20TubeTransactor) WithdrawCoin(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "withdrawCoin", _to)
}

// WithdrawCoin is a paid mutator transaction binding the contract method 0x5a73928f.
//
// Solidity: function withdrawCoin(address _to) returns()
func (_ERC20Tube *ERC20TubeSession) WithdrawCoin(_to common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.WithdrawCoin(&_ERC20Tube.TransactOpts, _to)
}

// WithdrawCoin is a paid mutator transaction binding the contract method 0x5a73928f.
//
// Solidity: function withdrawCoin(address _to) returns()
func (_ERC20Tube *ERC20TubeTransactorSession) WithdrawCoin(_to common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.WithdrawCoin(&_ERC20Tube.TransactOpts, _to)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address _to, address _token) returns()
func (_ERC20Tube *ERC20TubeTransactor) WithdrawToken(opts *bind.TransactOpts, _to common.Address, _token common.Address) (*types.Transaction, error) {
	return _ERC20Tube.contract.Transact(opts, "withdrawToken", _to, _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address _to, address _token) returns()
func (_ERC20Tube *ERC20TubeSession) WithdrawToken(_to common.Address, _token common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.WithdrawToken(&_ERC20Tube.TransactOpts, _to, _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address _to, address _token) returns()
func (_ERC20Tube *ERC20TubeTransactorSession) WithdrawToken(_to common.Address, _token common.Address) (*types.Transaction, error) {
	return _ERC20Tube.Contract.WithdrawToken(&_ERC20Tube.TransactOpts, _to, _token)
}

// ERC20TubeEmergencyOperatorSetIterator is returned from FilterEmergencyOperatorSet and is used to iterate over the raw logs and unpacked data for EmergencyOperatorSet events raised by the ERC20Tube contract.
type ERC20TubeEmergencyOperatorSetIterator struct {
	Event *ERC20TubeEmergencyOperatorSet // Event containing the contract specifics and raw log

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
func (it *ERC20TubeEmergencyOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TubeEmergencyOperatorSet)
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
		it.Event = new(ERC20TubeEmergencyOperatorSet)
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
func (it *ERC20TubeEmergencyOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TubeEmergencyOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TubeEmergencyOperatorSet represents a EmergencyOperatorSet event raised by the ERC20Tube contract.
type ERC20TubeEmergencyOperatorSet struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEmergencyOperatorSet is a free log retrieval operation binding the contract event 0x9a174d577c7665e63298da21ff932b957bb0dc51d46920b0cc1e32f22437d8fd.
//
// Solidity: event EmergencyOperatorSet(address indexed operator)
func (_ERC20Tube *ERC20TubeFilterer) FilterEmergencyOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ERC20TubeEmergencyOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC20Tube.contract.FilterLogs(opts, "EmergencyOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeEmergencyOperatorSetIterator{contract: _ERC20Tube.contract, event: "EmergencyOperatorSet", logs: logs, sub: sub}, nil
}

// WatchEmergencyOperatorSet is a free log subscription operation binding the contract event 0x9a174d577c7665e63298da21ff932b957bb0dc51d46920b0cc1e32f22437d8fd.
//
// Solidity: event EmergencyOperatorSet(address indexed operator)
func (_ERC20Tube *ERC20TubeFilterer) WatchEmergencyOperatorSet(opts *bind.WatchOpts, sink chan<- *ERC20TubeEmergencyOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC20Tube.contract.WatchLogs(opts, "EmergencyOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TubeEmergencyOperatorSet)
				if err := _ERC20Tube.contract.UnpackLog(event, "EmergencyOperatorSet", log); err != nil {
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

// ParseEmergencyOperatorSet is a log parse operation binding the contract event 0x9a174d577c7665e63298da21ff932b957bb0dc51d46920b0cc1e32f22437d8fd.
//
// Solidity: event EmergencyOperatorSet(address indexed operator)
func (_ERC20Tube *ERC20TubeFilterer) ParseEmergencyOperatorSet(log types.Log) (*ERC20TubeEmergencyOperatorSet, error) {
	event := new(ERC20TubeEmergencyOperatorSet)
	if err := _ERC20Tube.contract.UnpackLog(event, "EmergencyOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TubeLordUpdatedIterator is returned from FilterLordUpdated and is used to iterate over the raw logs and unpacked data for LordUpdated events raised by the ERC20Tube contract.
type ERC20TubeLordUpdatedIterator struct {
	Event *ERC20TubeLordUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20TubeLordUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TubeLordUpdated)
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
		it.Event = new(ERC20TubeLordUpdated)
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
func (it *ERC20TubeLordUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TubeLordUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TubeLordUpdated represents a LordUpdated event raised by the ERC20Tube contract.
type ERC20TubeLordUpdated struct {
	Lord common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLordUpdated is a free log retrieval operation binding the contract event 0x709dc3c1241091107656b907d68eb690a999b90156fd0814e832ca64b3df8486.
//
// Solidity: event LordUpdated(address indexed lord)
func (_ERC20Tube *ERC20TubeFilterer) FilterLordUpdated(opts *bind.FilterOpts, lord []common.Address) (*ERC20TubeLordUpdatedIterator, error) {

	var lordRule []interface{}
	for _, lordItem := range lord {
		lordRule = append(lordRule, lordItem)
	}

	logs, sub, err := _ERC20Tube.contract.FilterLogs(opts, "LordUpdated", lordRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeLordUpdatedIterator{contract: _ERC20Tube.contract, event: "LordUpdated", logs: logs, sub: sub}, nil
}

// WatchLordUpdated is a free log subscription operation binding the contract event 0x709dc3c1241091107656b907d68eb690a999b90156fd0814e832ca64b3df8486.
//
// Solidity: event LordUpdated(address indexed lord)
func (_ERC20Tube *ERC20TubeFilterer) WatchLordUpdated(opts *bind.WatchOpts, sink chan<- *ERC20TubeLordUpdated, lord []common.Address) (event.Subscription, error) {

	var lordRule []interface{}
	for _, lordItem := range lord {
		lordRule = append(lordRule, lordItem)
	}

	logs, sub, err := _ERC20Tube.contract.WatchLogs(opts, "LordUpdated", lordRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TubeLordUpdated)
				if err := _ERC20Tube.contract.UnpackLog(event, "LordUpdated", log); err != nil {
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

// ParseLordUpdated is a log parse operation binding the contract event 0x709dc3c1241091107656b907d68eb690a999b90156fd0814e832ca64b3df8486.
//
// Solidity: event LordUpdated(address indexed lord)
func (_ERC20Tube *ERC20TubeFilterer) ParseLordUpdated(log types.Log) (*ERC20TubeLordUpdated, error) {
	event := new(ERC20TubeLordUpdated)
	if err := _ERC20Tube.contract.UnpackLog(event, "LordUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TubeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20Tube contract.
type ERC20TubeOwnershipTransferredIterator struct {
	Event *ERC20TubeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20TubeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TubeOwnershipTransferred)
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
		it.Event = new(ERC20TubeOwnershipTransferred)
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
func (it *ERC20TubeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TubeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TubeOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20Tube contract.
type ERC20TubeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20Tube *ERC20TubeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20TubeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20Tube.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeOwnershipTransferredIterator{contract: _ERC20Tube.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20Tube *ERC20TubeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20TubeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20Tube.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TubeOwnershipTransferred)
				if err := _ERC20Tube.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20Tube *ERC20TubeFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20TubeOwnershipTransferred, error) {
	event := new(ERC20TubeOwnershipTransferred)
	if err := _ERC20Tube.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TubePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20Tube contract.
type ERC20TubePausedIterator struct {
	Event *ERC20TubePaused // Event containing the contract specifics and raw log

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
func (it *ERC20TubePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TubePaused)
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
		it.Event = new(ERC20TubePaused)
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
func (it *ERC20TubePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TubePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TubePaused represents a Paused event raised by the ERC20Tube contract.
type ERC20TubePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Tube *ERC20TubeFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20TubePausedIterator, error) {

	logs, sub, err := _ERC20Tube.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20TubePausedIterator{contract: _ERC20Tube.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Tube *ERC20TubeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20TubePaused) (event.Subscription, error) {

	logs, sub, err := _ERC20Tube.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TubePaused)
				if err := _ERC20Tube.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Tube *ERC20TubeFilterer) ParsePaused(log types.Log) (*ERC20TubePaused, error) {
	event := new(ERC20TubePaused)
	if err := _ERC20Tube.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TubeReceiptIterator is returned from FilterReceipt and is used to iterate over the raw logs and unpacked data for Receipt events raised by the ERC20Tube contract.
type ERC20TubeReceiptIterator struct {
	Event *ERC20TubeReceipt // Event containing the contract specifics and raw log

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
func (it *ERC20TubeReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TubeReceipt)
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
		it.Event = new(ERC20TubeReceipt)
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
func (it *ERC20TubeReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TubeReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TubeReceipt represents a Receipt event raised by the ERC20Tube contract.
type ERC20TubeReceipt struct {
	Nonce        *big.Int
	Sender       common.Address
	Token        common.Address
	Amount       *big.Int
	TargetTubeID *big.Int
	Recipient    common.Address
	Fee          *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterReceipt is a free log retrieval operation binding the contract event 0x9cc019a8c1ed04f1d296410e07c9f2c92849081a15a108e6289724e97953d5d3.
//
// Solidity: event Receipt(uint256 indexed nonce, address sender, address indexed token, uint256 amount, uint256 indexed targetTubeID, address recipient, uint256 fee)
func (_ERC20Tube *ERC20TubeFilterer) FilterReceipt(opts *bind.FilterOpts, nonce []*big.Int, token []common.Address, targetTubeID []*big.Int) (*ERC20TubeReceiptIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var targetTubeIDRule []interface{}
	for _, targetTubeIDItem := range targetTubeID {
		targetTubeIDRule = append(targetTubeIDRule, targetTubeIDItem)
	}

	logs, sub, err := _ERC20Tube.contract.FilterLogs(opts, "Receipt", nonceRule, tokenRule, targetTubeIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeReceiptIterator{contract: _ERC20Tube.contract, event: "Receipt", logs: logs, sub: sub}, nil
}

// WatchReceipt is a free log subscription operation binding the contract event 0x9cc019a8c1ed04f1d296410e07c9f2c92849081a15a108e6289724e97953d5d3.
//
// Solidity: event Receipt(uint256 indexed nonce, address sender, address indexed token, uint256 amount, uint256 indexed targetTubeID, address recipient, uint256 fee)
func (_ERC20Tube *ERC20TubeFilterer) WatchReceipt(opts *bind.WatchOpts, sink chan<- *ERC20TubeReceipt, nonce []*big.Int, token []common.Address, targetTubeID []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var targetTubeIDRule []interface{}
	for _, targetTubeIDItem := range targetTubeID {
		targetTubeIDRule = append(targetTubeIDRule, targetTubeIDItem)
	}

	logs, sub, err := _ERC20Tube.contract.WatchLogs(opts, "Receipt", nonceRule, tokenRule, targetTubeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TubeReceipt)
				if err := _ERC20Tube.contract.UnpackLog(event, "Receipt", log); err != nil {
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

// ParseReceipt is a log parse operation binding the contract event 0x9cc019a8c1ed04f1d296410e07c9f2c92849081a15a108e6289724e97953d5d3.
//
// Solidity: event Receipt(uint256 indexed nonce, address sender, address indexed token, uint256 amount, uint256 indexed targetTubeID, address recipient, uint256 fee)
func (_ERC20Tube *ERC20TubeFilterer) ParseReceipt(log types.Log) (*ERC20TubeReceipt, error) {
	event := new(ERC20TubeReceipt)
	if err := _ERC20Tube.contract.UnpackLog(event, "Receipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TubeSafeUpdatedIterator is returned from FilterSafeUpdated and is used to iterate over the raw logs and unpacked data for SafeUpdated events raised by the ERC20Tube contract.
type ERC20TubeSafeUpdatedIterator struct {
	Event *ERC20TubeSafeUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20TubeSafeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TubeSafeUpdated)
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
		it.Event = new(ERC20TubeSafeUpdated)
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
func (it *ERC20TubeSafeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TubeSafeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TubeSafeUpdated represents a SafeUpdated event raised by the ERC20Tube contract.
type ERC20TubeSafeUpdated struct {
	Safe common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSafeUpdated is a free log retrieval operation binding the contract event 0x0ba726854996c35c3b65654380d718bfc2515d517bc26790fc8051b1d8da5485.
//
// Solidity: event SafeUpdated(address indexed safe)
func (_ERC20Tube *ERC20TubeFilterer) FilterSafeUpdated(opts *bind.FilterOpts, safe []common.Address) (*ERC20TubeSafeUpdatedIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _ERC20Tube.contract.FilterLogs(opts, "SafeUpdated", safeRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeSafeUpdatedIterator{contract: _ERC20Tube.contract, event: "SafeUpdated", logs: logs, sub: sub}, nil
}

// WatchSafeUpdated is a free log subscription operation binding the contract event 0x0ba726854996c35c3b65654380d718bfc2515d517bc26790fc8051b1d8da5485.
//
// Solidity: event SafeUpdated(address indexed safe)
func (_ERC20Tube *ERC20TubeFilterer) WatchSafeUpdated(opts *bind.WatchOpts, sink chan<- *ERC20TubeSafeUpdated, safe []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _ERC20Tube.contract.WatchLogs(opts, "SafeUpdated", safeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TubeSafeUpdated)
				if err := _ERC20Tube.contract.UnpackLog(event, "SafeUpdated", log); err != nil {
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

// ParseSafeUpdated is a log parse operation binding the contract event 0x0ba726854996c35c3b65654380d718bfc2515d517bc26790fc8051b1d8da5485.
//
// Solidity: event SafeUpdated(address indexed safe)
func (_ERC20Tube *ERC20TubeFilterer) ParseSafeUpdated(log types.Log) (*ERC20TubeSafeUpdated, error) {
	event := new(ERC20TubeSafeUpdated)
	if err := _ERC20Tube.contract.UnpackLog(event, "SafeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TubeSettledIterator is returned from FilterSettled and is used to iterate over the raw logs and unpacked data for Settled events raised by the ERC20Tube contract.
type ERC20TubeSettledIterator struct {
	Event *ERC20TubeSettled // Event containing the contract specifics and raw log

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
func (it *ERC20TubeSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TubeSettled)
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
		it.Event = new(ERC20TubeSettled)
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
func (it *ERC20TubeSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TubeSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TubeSettled represents a Settled event raised by the ERC20Tube contract.
type ERC20TubeSettled struct {
	Key        [32]byte
	Validators []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettled is a free log retrieval operation binding the contract event 0xe24922ac8cf2a1430fb8c7ce6a9125d7db5076a1eb2cefced90e82d6fcb24db0.
//
// Solidity: event Settled(bytes32 indexed key, address[] validators)
func (_ERC20Tube *ERC20TubeFilterer) FilterSettled(opts *bind.FilterOpts, key [][32]byte) (*ERC20TubeSettledIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _ERC20Tube.contract.FilterLogs(opts, "Settled", keyRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TubeSettledIterator{contract: _ERC20Tube.contract, event: "Settled", logs: logs, sub: sub}, nil
}

// WatchSettled is a free log subscription operation binding the contract event 0xe24922ac8cf2a1430fb8c7ce6a9125d7db5076a1eb2cefced90e82d6fcb24db0.
//
// Solidity: event Settled(bytes32 indexed key, address[] validators)
func (_ERC20Tube *ERC20TubeFilterer) WatchSettled(opts *bind.WatchOpts, sink chan<- *ERC20TubeSettled, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _ERC20Tube.contract.WatchLogs(opts, "Settled", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TubeSettled)
				if err := _ERC20Tube.contract.UnpackLog(event, "Settled", log); err != nil {
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

// ParseSettled is a log parse operation binding the contract event 0xe24922ac8cf2a1430fb8c7ce6a9125d7db5076a1eb2cefced90e82d6fcb24db0.
//
// Solidity: event Settled(bytes32 indexed key, address[] validators)
func (_ERC20Tube *ERC20TubeFilterer) ParseSettled(log types.Log) (*ERC20TubeSettled, error) {
	event := new(ERC20TubeSettled)
	if err := _ERC20Tube.contract.UnpackLog(event, "Settled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TubeTubeInfoUpdatedIterator is returned from FilterTubeInfoUpdated and is used to iterate over the raw logs and unpacked data for TubeInfoUpdated events raised by the ERC20Tube contract.
type ERC20TubeTubeInfoUpdatedIterator struct {
	Event *ERC20TubeTubeInfoUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20TubeTubeInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TubeTubeInfoUpdated)
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
		it.Event = new(ERC20TubeTubeInfoUpdated)
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
func (it *ERC20TubeTubeInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TubeTubeInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TubeTubeInfoUpdated represents a TubeInfoUpdated event raised by the ERC20Tube contract.
type ERC20TubeTubeInfoUpdated struct {
	TubeID  *big.Int
	FeeRate *big.Int
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTubeInfoUpdated is a free log retrieval operation binding the contract event 0x3d86dc5b78959889a7cc5db345a7c77a4d1b7a65a1db9de105cac31ea379320b.
//
// Solidity: event TubeInfoUpdated(uint256 tubeID, uint256 feeRate, bool enabled)
func (_ERC20Tube *ERC20TubeFilterer) FilterTubeInfoUpdated(opts *bind.FilterOpts) (*ERC20TubeTubeInfoUpdatedIterator, error) {

	logs, sub, err := _ERC20Tube.contract.FilterLogs(opts, "TubeInfoUpdated")
	if err != nil {
		return nil, err
	}
	return &ERC20TubeTubeInfoUpdatedIterator{contract: _ERC20Tube.contract, event: "TubeInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchTubeInfoUpdated is a free log subscription operation binding the contract event 0x3d86dc5b78959889a7cc5db345a7c77a4d1b7a65a1db9de105cac31ea379320b.
//
// Solidity: event TubeInfoUpdated(uint256 tubeID, uint256 feeRate, bool enabled)
func (_ERC20Tube *ERC20TubeFilterer) WatchTubeInfoUpdated(opts *bind.WatchOpts, sink chan<- *ERC20TubeTubeInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _ERC20Tube.contract.WatchLogs(opts, "TubeInfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TubeTubeInfoUpdated)
				if err := _ERC20Tube.contract.UnpackLog(event, "TubeInfoUpdated", log); err != nil {
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

// ParseTubeInfoUpdated is a log parse operation binding the contract event 0x3d86dc5b78959889a7cc5db345a7c77a4d1b7a65a1db9de105cac31ea379320b.
//
// Solidity: event TubeInfoUpdated(uint256 tubeID, uint256 feeRate, bool enabled)
func (_ERC20Tube *ERC20TubeFilterer) ParseTubeInfoUpdated(log types.Log) (*ERC20TubeTubeInfoUpdated, error) {
	event := new(ERC20TubeTubeInfoUpdated)
	if err := _ERC20Tube.contract.UnpackLog(event, "TubeInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TubeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20Tube contract.
type ERC20TubeUnpausedIterator struct {
	Event *ERC20TubeUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20TubeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TubeUnpaused)
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
		it.Event = new(ERC20TubeUnpaused)
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
func (it *ERC20TubeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TubeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TubeUnpaused represents a Unpaused event raised by the ERC20Tube contract.
type ERC20TubeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Tube *ERC20TubeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20TubeUnpausedIterator, error) {

	logs, sub, err := _ERC20Tube.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20TubeUnpausedIterator{contract: _ERC20Tube.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Tube *ERC20TubeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20TubeUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20Tube.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TubeUnpaused)
				if err := _ERC20Tube.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Tube *ERC20TubeFilterer) ParseUnpaused(log types.Log) (*ERC20TubeUnpaused, error) {
	event := new(ERC20TubeUnpaused)
	if err := _ERC20Tube.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
