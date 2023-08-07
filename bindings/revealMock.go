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

// RevealAssignedAsset is an auto generated low-level Go binding around an user-defined struct.
type RevealAssignedAsset struct {
	LeafId uint16
	Value  *big.Int
}

// RevealMerkleTree is an auto generated low-level Go binding around an user-defined struct.
type RevealMerkleTree struct {
	Values []RevealAssignedAsset
	Proofs [][][32]byte
	Root   [32]byte
}

// RevealMockMetaData contains all meta data concerning the RevealMock contract.
var RevealMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"commitmentHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numActiveCollections\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"components\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"leafId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structReveal.AssignedAsset[]\",\"name\":\"values\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"proofs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structReveal.MerkleTree\",\"name\":\"tree\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_commitmentHash\",\"type\":\"bytes32\"}],\"name\":\"setCommitmentHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"setSalt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toAssign\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RevealMockABI is the input ABI used to generate the binding from.
// Deprecated: Use RevealMockMetaData.ABI instead.
var RevealMockABI = RevealMockMetaData.ABI

// RevealMock is an auto generated Go binding around an Ethereum contract.
type RevealMock struct {
	RevealMockCaller     // Read-only binding to the contract
	RevealMockTransactor // Write-only binding to the contract
	RevealMockFilterer   // Log filterer for contract events
}

// RevealMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type RevealMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevealMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RevealMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevealMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RevealMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevealMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RevealMockSession struct {
	Contract     *RevealMock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RevealMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RevealMockCallerSession struct {
	Contract *RevealMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RevealMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RevealMockTransactorSession struct {
	Contract     *RevealMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RevealMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type RevealMockRaw struct {
	Contract *RevealMock // Generic contract binding to access the raw methods on
}

// RevealMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RevealMockCallerRaw struct {
	Contract *RevealMockCaller // Generic read-only contract binding to access the raw methods on
}

// RevealMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RevealMockTransactorRaw struct {
	Contract *RevealMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRevealMock creates a new instance of RevealMock, bound to a specific deployed contract.
func NewRevealMock(address common.Address, backend bind.ContractBackend) (*RevealMock, error) {
	contract, err := bindRevealMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RevealMock{RevealMockCaller: RevealMockCaller{contract: contract}, RevealMockTransactor: RevealMockTransactor{contract: contract}, RevealMockFilterer: RevealMockFilterer{contract: contract}}, nil
}

// NewRevealMockCaller creates a new read-only instance of RevealMock, bound to a specific deployed contract.
func NewRevealMockCaller(address common.Address, caller bind.ContractCaller) (*RevealMockCaller, error) {
	contract, err := bindRevealMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RevealMockCaller{contract: contract}, nil
}

// NewRevealMockTransactor creates a new write-only instance of RevealMock, bound to a specific deployed contract.
func NewRevealMockTransactor(address common.Address, transactor bind.ContractTransactor) (*RevealMockTransactor, error) {
	contract, err := bindRevealMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RevealMockTransactor{contract: contract}, nil
}

// NewRevealMockFilterer creates a new log filterer instance of RevealMock, bound to a specific deployed contract.
func NewRevealMockFilterer(address common.Address, filterer bind.ContractFilterer) (*RevealMockFilterer, error) {
	contract, err := bindRevealMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RevealMockFilterer{contract: contract}, nil
}

// bindRevealMock binds a generic wrapper to an already deployed contract.
func bindRevealMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RevealMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevealMock *RevealMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevealMock.Contract.RevealMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevealMock *RevealMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevealMock.Contract.RevealMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevealMock *RevealMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevealMock.Contract.RevealMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevealMock *RevealMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevealMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevealMock *RevealMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevealMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevealMock *RevealMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevealMock.Contract.contract.Transact(opts, method, params...)
}

// CommitmentHash is a free data retrieval call binding the contract method 0xf6098ce8.
//
// Solidity: function commitmentHash() view returns(bytes32)
func (_RevealMock *RevealMockCaller) CommitmentHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RevealMock.contract.Call(opts, &out, "commitmentHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommitmentHash is a free data retrieval call binding the contract method 0xf6098ce8.
//
// Solidity: function commitmentHash() view returns(bytes32)
func (_RevealMock *RevealMockSession) CommitmentHash() ([32]byte, error) {
	return _RevealMock.Contract.CommitmentHash(&_RevealMock.CallOpts)
}

// CommitmentHash is a free data retrieval call binding the contract method 0xf6098ce8.
//
// Solidity: function commitmentHash() view returns(bytes32)
func (_RevealMock *RevealMockCallerSession) CommitmentHash() ([32]byte, error) {
	return _RevealMock.Contract.CommitmentHash(&_RevealMock.CallOpts)
}

// Depth is a free data retrieval call binding the contract method 0x631c56ef.
//
// Solidity: function depth() view returns(uint256)
func (_RevealMock *RevealMockCaller) Depth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RevealMock.contract.Call(opts, &out, "depth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Depth is a free data retrieval call binding the contract method 0x631c56ef.
//
// Solidity: function depth() view returns(uint256)
func (_RevealMock *RevealMockSession) Depth() (*big.Int, error) {
	return _RevealMock.Contract.Depth(&_RevealMock.CallOpts)
}

// Depth is a free data retrieval call binding the contract method 0x631c56ef.
//
// Solidity: function depth() view returns(uint256)
func (_RevealMock *RevealMockCallerSession) Depth() (*big.Int, error) {
	return _RevealMock.Contract.Depth(&_RevealMock.CallOpts)
}

// NumActiveCollections is a free data retrieval call binding the contract method 0x8f2702ef.
//
// Solidity: function numActiveCollections() view returns(uint16)
func (_RevealMock *RevealMockCaller) NumActiveCollections(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _RevealMock.contract.Call(opts, &out, "numActiveCollections")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// NumActiveCollections is a free data retrieval call binding the contract method 0x8f2702ef.
//
// Solidity: function numActiveCollections() view returns(uint16)
func (_RevealMock *RevealMockSession) NumActiveCollections() (uint16, error) {
	return _RevealMock.Contract.NumActiveCollections(&_RevealMock.CallOpts)
}

// NumActiveCollections is a free data retrieval call binding the contract method 0x8f2702ef.
//
// Solidity: function numActiveCollections() view returns(uint16)
func (_RevealMock *RevealMockCallerSession) NumActiveCollections() (uint16, error) {
	return _RevealMock.Contract.NumActiveCollections(&_RevealMock.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_RevealMock *RevealMockCaller) Salt(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RevealMock.contract.Call(opts, &out, "salt")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_RevealMock *RevealMockSession) Salt() ([32]byte, error) {
	return _RevealMock.Contract.Salt(&_RevealMock.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_RevealMock *RevealMockCallerSession) Salt() ([32]byte, error) {
	return _RevealMock.Contract.Salt(&_RevealMock.CallOpts)
}

// ToAssign is a free data retrieval call binding the contract method 0xcbb24791.
//
// Solidity: function toAssign() view returns(uint256)
func (_RevealMock *RevealMockCaller) ToAssign(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RevealMock.contract.Call(opts, &out, "toAssign")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToAssign is a free data retrieval call binding the contract method 0xcbb24791.
//
// Solidity: function toAssign() view returns(uint256)
func (_RevealMock *RevealMockSession) ToAssign() (*big.Int, error) {
	return _RevealMock.Contract.ToAssign(&_RevealMock.CallOpts)
}

// ToAssign is a free data retrieval call binding the contract method 0xcbb24791.
//
// Solidity: function toAssign() view returns(uint256)
func (_RevealMock *RevealMockCallerSession) ToAssign() (*big.Int, error) {
	return _RevealMock.Contract.ToAssign(&_RevealMock.CallOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0xf110a0fa.
//
// Solidity: function reveal(uint32 epoch, ((uint16,uint256)[],bytes32[][],bytes32) tree, bytes signature) returns()
func (_RevealMock *RevealMockTransactor) Reveal(opts *bind.TransactOpts, epoch uint32, tree RevealMerkleTree, signature []byte) (*types.Transaction, error) {
	return _RevealMock.contract.Transact(opts, "reveal", epoch, tree, signature)
}

// Reveal is a paid mutator transaction binding the contract method 0xf110a0fa.
//
// Solidity: function reveal(uint32 epoch, ((uint16,uint256)[],bytes32[][],bytes32) tree, bytes signature) returns()
func (_RevealMock *RevealMockSession) Reveal(epoch uint32, tree RevealMerkleTree, signature []byte) (*types.Transaction, error) {
	return _RevealMock.Contract.Reveal(&_RevealMock.TransactOpts, epoch, tree, signature)
}

// Reveal is a paid mutator transaction binding the contract method 0xf110a0fa.
//
// Solidity: function reveal(uint32 epoch, ((uint16,uint256)[],bytes32[][],bytes32) tree, bytes signature) returns()
func (_RevealMock *RevealMockTransactorSession) Reveal(epoch uint32, tree RevealMerkleTree, signature []byte) (*types.Transaction, error) {
	return _RevealMock.Contract.Reveal(&_RevealMock.TransactOpts, epoch, tree, signature)
}

// SetCommitmentHash is a paid mutator transaction binding the contract method 0x4c2cbde7.
//
// Solidity: function setCommitmentHash(bytes32 _commitmentHash) returns()
func (_RevealMock *RevealMockTransactor) SetCommitmentHash(opts *bind.TransactOpts, _commitmentHash [32]byte) (*types.Transaction, error) {
	return _RevealMock.contract.Transact(opts, "setCommitmentHash", _commitmentHash)
}

// SetCommitmentHash is a paid mutator transaction binding the contract method 0x4c2cbde7.
//
// Solidity: function setCommitmentHash(bytes32 _commitmentHash) returns()
func (_RevealMock *RevealMockSession) SetCommitmentHash(_commitmentHash [32]byte) (*types.Transaction, error) {
	return _RevealMock.Contract.SetCommitmentHash(&_RevealMock.TransactOpts, _commitmentHash)
}

// SetCommitmentHash is a paid mutator transaction binding the contract method 0x4c2cbde7.
//
// Solidity: function setCommitmentHash(bytes32 _commitmentHash) returns()
func (_RevealMock *RevealMockTransactorSession) SetCommitmentHash(_commitmentHash [32]byte) (*types.Transaction, error) {
	return _RevealMock.Contract.SetCommitmentHash(&_RevealMock.TransactOpts, _commitmentHash)
}

// SetSalt is a paid mutator transaction binding the contract method 0x611ef44f.
//
// Solidity: function setSalt(bytes32 _salt) returns()
func (_RevealMock *RevealMockTransactor) SetSalt(opts *bind.TransactOpts, _salt [32]byte) (*types.Transaction, error) {
	return _RevealMock.contract.Transact(opts, "setSalt", _salt)
}

// SetSalt is a paid mutator transaction binding the contract method 0x611ef44f.
//
// Solidity: function setSalt(bytes32 _salt) returns()
func (_RevealMock *RevealMockSession) SetSalt(_salt [32]byte) (*types.Transaction, error) {
	return _RevealMock.Contract.SetSalt(&_RevealMock.TransactOpts, _salt)
}

// SetSalt is a paid mutator transaction binding the contract method 0x611ef44f.
//
// Solidity: function setSalt(bytes32 _salt) returns()
func (_RevealMock *RevealMockTransactorSession) SetSalt(_salt [32]byte) (*types.Transaction, error) {
	return _RevealMock.Contract.SetSalt(&_RevealMock.TransactOpts, _salt)
}

