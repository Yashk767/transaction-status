// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GetSetContractMetaData contains all meta data concerning the GetSetContract contract.
var GetSetContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_testValue\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"}],\"name\":\"ValueChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GetSetContractABI is the input ABI used to generate the binding from.
// Deprecated: Use GetSetContractMetaData.ABI instead.
var GetSetContractABI = GetSetContractMetaData.ABI

// GetSetContract is an auto generated Go binding around an Ethereum contract.
type GetSetContract struct {
	GetSetContractCaller     // Read-only binding to the contract
	GetSetContractTransactor // Write-only binding to the contract
	GetSetContractFilterer   // Log filterer for contract events
}

// GetSetContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type GetSetContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetSetContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GetSetContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetSetContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GetSetContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetSetContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GetSetContractSession struct {
	Contract     *GetSetContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GetSetContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GetSetContractCallerSession struct {
	Contract *GetSetContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GetSetContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GetSetContractTransactorSession struct {
	Contract     *GetSetContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GetSetContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type GetSetContractRaw struct {
	Contract *GetSetContract // Generic contract binding to access the raw methods on
}

// GetSetContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GetSetContractCallerRaw struct {
	Contract *GetSetContractCaller // Generic read-only contract binding to access the raw methods on
}

// GetSetContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GetSetContractTransactorRaw struct {
	Contract *GetSetContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGetSetContract creates a new instance of GetSetContract, bound to a specific deployed contract.
func NewGetSetContract(address common.Address, backend bind.ContractBackend) (*GetSetContract, error) {
	contract, err := bindGetSetContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GetSetContract{GetSetContractCaller: GetSetContractCaller{contract: contract}, GetSetContractTransactor: GetSetContractTransactor{contract: contract}, GetSetContractFilterer: GetSetContractFilterer{contract: contract}}, nil
}

// NewGetSetContractCaller creates a new read-only instance of GetSetContract, bound to a specific deployed contract.
func NewGetSetContractCaller(address common.Address, caller bind.ContractCaller) (*GetSetContractCaller, error) {
	contract, err := bindGetSetContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GetSetContractCaller{contract: contract}, nil
}

// NewGetSetContractTransactor creates a new write-only instance of GetSetContract, bound to a specific deployed contract.
func NewGetSetContractTransactor(address common.Address, transactor bind.ContractTransactor) (*GetSetContractTransactor, error) {
	contract, err := bindGetSetContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GetSetContractTransactor{contract: contract}, nil
}

// NewGetSetContractFilterer creates a new log filterer instance of GetSetContract, bound to a specific deployed contract.
func NewGetSetContractFilterer(address common.Address, filterer bind.ContractFilterer) (*GetSetContractFilterer, error) {
	contract, err := bindGetSetContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GetSetContractFilterer{contract: contract}, nil
}

// bindGetSetContract binds a generic wrapper to an already deployed contract.
func bindGetSetContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GetSetContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetSetContract *GetSetContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GetSetContract.Contract.GetSetContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetSetContract *GetSetContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetSetContract.Contract.GetSetContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetSetContract *GetSetContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetSetContract.Contract.GetSetContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetSetContract *GetSetContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GetSetContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetSetContract *GetSetContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetSetContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetSetContract *GetSetContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetSetContract.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_GetSetContract *GetSetContractCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GetSetContract.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_GetSetContract *GetSetContractSession) Get() (*big.Int, error) {
	return _GetSetContract.Contract.Get(&_GetSetContract.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_GetSetContract *GetSetContractCallerSession) Get() (*big.Int, error) {
	return _GetSetContract.Contract.Get(&_GetSetContract.CallOpts)
}

// TestValue is a free data retrieval call binding the contract method 0x8af5de72.
//
// Solidity: function testValue() view returns(uint256)
func (_GetSetContract *GetSetContractCaller) TestValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GetSetContract.contract.Call(opts, &out, "testValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestValue is a free data retrieval call binding the contract method 0x8af5de72.
//
// Solidity: function testValue() view returns(uint256)
func (_GetSetContract *GetSetContractSession) TestValue() (*big.Int, error) {
	return _GetSetContract.Contract.TestValue(&_GetSetContract.CallOpts)
}

// TestValue is a free data retrieval call binding the contract method 0x8af5de72.
//
// Solidity: function testValue() view returns(uint256)
func (_GetSetContract *GetSetContractCallerSession) TestValue() (*big.Int, error) {
	return _GetSetContract.Contract.TestValue(&_GetSetContract.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 _testValue) returns()
func (_GetSetContract *GetSetContractTransactor) Set(opts *bind.TransactOpts, _testValue *big.Int) (*types.Transaction, error) {
	return _GetSetContract.contract.Transact(opts, "set", _testValue)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 _testValue) returns()
func (_GetSetContract *GetSetContractSession) Set(_testValue *big.Int) (*types.Transaction, error) {
	return _GetSetContract.Contract.Set(&_GetSetContract.TransactOpts, _testValue)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 _testValue) returns()
func (_GetSetContract *GetSetContractTransactorSession) Set(_testValue *big.Int) (*types.Transaction, error) {
	return _GetSetContract.Contract.Set(&_GetSetContract.TransactOpts, _testValue)
}

// GetSetContractValueChangedIterator is returned from FilterValueChanged and is used to iterate over the raw logs and unpacked data for ValueChanged events raised by the GetSetContract contract.
type GetSetContractValueChangedIterator struct {
	Event *GetSetContractValueChanged // Event containing the contract specifics and raw log

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
func (it *GetSetContractValueChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GetSetContractValueChanged)
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
		it.Event = new(GetSetContractValueChanged)
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
func (it *GetSetContractValueChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GetSetContractValueChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GetSetContractValueChanged represents a ValueChanged event raised by the GetSetContract contract.
type GetSetContractValueChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValueChanged is a free log retrieval operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 _newValue)
func (_GetSetContract *GetSetContractFilterer) FilterValueChanged(opts *bind.FilterOpts) (*GetSetContractValueChangedIterator, error) {

	logs, sub, err := _GetSetContract.contract.FilterLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}
	return &GetSetContractValueChangedIterator{contract: _GetSetContract.contract, event: "ValueChanged", logs: logs, sub: sub}, nil
}

// WatchValueChanged is a free log subscription operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 _newValue)
func (_GetSetContract *GetSetContractFilterer) WatchValueChanged(opts *bind.WatchOpts, sink chan<- *GetSetContractValueChanged) (event.Subscription, error) {

	logs, sub, err := _GetSetContract.contract.WatchLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GetSetContractValueChanged)
				if err := _GetSetContract.contract.UnpackLog(event, "ValueChanged", log); err != nil {
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

// ParseValueChanged is a log parse operation binding the contract event 0x93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c59.
//
// Solidity: event ValueChanged(uint256 _newValue)
func (_GetSetContract *GetSetContractFilterer) ParseValueChanged(log types.Log) (*GetSetContractValueChanged, error) {
	event := new(GetSetContractValueChanged)
	if err := _GetSetContract.contract.UnpackLog(event, "ValueChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
