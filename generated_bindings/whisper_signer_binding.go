// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shyft_contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ShyftNetwork/go-empyrean"
	"github.com/ShyftNetwork/go-empyrean/accounts/abi"
	"github.com/ShyftNetwork/go-empyrean/accounts/abi/bind"
	"github.com/ShyftNetwork/go-empyrean/common"
	"github.com/ShyftNetwork/go-empyrean/core/types"
	"github.com/ShyftNetwork/go-empyrean/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ValidSignersABI is the input ABI used to generate the binding from.
const ValidSignersABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"removeSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"isValidSigner\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_new_signer\",\"type\":\"address\"}],\"name\":\"addValidSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ValidSignersBin is the compiled bytecode used for deploying new contracts.
const ValidSignersBin = `0x608060405234801561001057600080fd5b5060018054600160a060020a0319163317905561020b806100326000396000f3fe60806040526004361061005b577c010000000000000000000000000000000000000000000000000000000060003504630e316ab781146100605780638da5cb5b14610095578063d5f50582146100c6578063de8ee8171461010d575b600080fd5b34801561006c57600080fd5b506100936004803603602081101561008357600080fd5b5035600160a060020a0316610140565b005b3480156100a157600080fd5b506100aa610177565b60408051600160a060020a039092168252519081900360200190f35b3480156100d257600080fd5b506100f9600480360360208110156100e957600080fd5b5035600160a060020a0316610186565b604080519115158252519081900360200190f35b34801561011957600080fd5b506100936004803603602081101561013057600080fd5b5035600160a060020a03166101a4565b600154600160a060020a031633141561005b57600160a060020a0381166000908152602081905260409020805460ff191690555b50565b600154600160a060020a031681565b600160a060020a031660009081526020819052604090205460ff1690565b600154600160a060020a031633141561005b57600160a060020a0381166000908152602081905260409020805460ff1916600117905561017456fea165627a7a723058204c3826c0149ce63e78e417db3a1ec89f1010254ca15f7594dd835908f85163df0029`

// DeployValidSigners deploys a new Ethereum contract, binding an instance of ValidSigners to it.
func DeployValidSigners(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidSigners, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidSignersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidSignersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidSigners{ValidSignersCaller: ValidSignersCaller{contract: contract}, ValidSignersTransactor: ValidSignersTransactor{contract: contract}, ValidSignersFilterer: ValidSignersFilterer{contract: contract}}, nil
}

// ValidSigners is an auto generated Go binding around an Ethereum contract.
type ValidSigners struct {
	ValidSignersCaller     // Read-only binding to the contract
	ValidSignersTransactor // Write-only binding to the contract
	ValidSignersFilterer   // Log filterer for contract events
}

// ValidSignersCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidSignersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidSignersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidSignersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidSignersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidSignersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidSignersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidSignersSession struct {
	Contract     *ValidSigners     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidSignersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidSignersCallerSession struct {
	Contract *ValidSignersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ValidSignersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidSignersTransactorSession struct {
	Contract     *ValidSignersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidSignersRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidSignersRaw struct {
	Contract *ValidSigners // Generic contract binding to access the raw methods on
}

// ValidSignersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidSignersCallerRaw struct {
	Contract *ValidSignersCaller // Generic read-only contract binding to access the raw methods on
}

// ValidSignersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidSignersTransactorRaw struct {
	Contract *ValidSignersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidSigners creates a new instance of ValidSigners, bound to a specific deployed contract.
func NewValidSigners(address common.Address, backend bind.ContractBackend) (*ValidSigners, error) {
	contract, err := bindValidSigners(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidSigners{ValidSignersCaller: ValidSignersCaller{contract: contract}, ValidSignersTransactor: ValidSignersTransactor{contract: contract}, ValidSignersFilterer: ValidSignersFilterer{contract: contract}}, nil
}

// NewValidSignersCaller creates a new read-only instance of ValidSigners, bound to a specific deployed contract.
func NewValidSignersCaller(address common.Address, caller bind.ContractCaller) (*ValidSignersCaller, error) {
	contract, err := bindValidSigners(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidSignersCaller{contract: contract}, nil
}

// NewValidSignersTransactor creates a new write-only instance of ValidSigners, bound to a specific deployed contract.
func NewValidSignersTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidSignersTransactor, error) {
	contract, err := bindValidSigners(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidSignersTransactor{contract: contract}, nil
}

// NewValidSignersFilterer creates a new log filterer instance of ValidSigners, bound to a specific deployed contract.
func NewValidSignersFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidSignersFilterer, error) {
	contract, err := bindValidSigners(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidSignersFilterer{contract: contract}, nil
}

// bindValidSigners binds a generic wrapper to an already deployed contract.
func bindValidSigners(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidSignersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidSigners *ValidSignersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidSigners.Contract.ValidSignersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidSigners *ValidSignersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidSigners.Contract.ValidSignersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidSigners *ValidSignersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidSigners.Contract.ValidSignersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidSigners *ValidSignersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidSigners.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidSigners *ValidSignersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidSigners.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidSigners *ValidSignersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidSigners.Contract.contract.Transact(opts, method, params...)
}

// IsValidSigner is a free data retrieval call binding the contract method 0xd5f50582.
//
// Solidity: function isValidSigner(_signer address) constant returns(result bool)
func (_ValidSigners *ValidSignersCaller) IsValidSigner(opts *bind.CallOpts, _signer common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidSigners.contract.Call(opts, out, "isValidSigner", _signer)
	return *ret0, err
}

// IsValidSigner is a free data retrieval call binding the contract method 0xd5f50582.
//
// Solidity: function isValidSigner(_signer address) constant returns(result bool)
func (_ValidSigners *ValidSignersSession) IsValidSigner(_signer common.Address) (bool, error) {
	return _ValidSigners.Contract.IsValidSigner(&_ValidSigners.CallOpts, _signer)
}

// IsValidSigner is a free data retrieval call binding the contract method 0xd5f50582.
//
// Solidity: function isValidSigner(_signer address) constant returns(result bool)
func (_ValidSigners *ValidSignersCallerSession) IsValidSigner(_signer common.Address) (bool, error) {
	return _ValidSigners.Contract.IsValidSigner(&_ValidSigners.CallOpts, _signer)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ValidSigners *ValidSignersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidSigners.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ValidSigners *ValidSignersSession) Owner() (common.Address, error) {
	return _ValidSigners.Contract.Owner(&_ValidSigners.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ValidSigners *ValidSignersCallerSession) Owner() (common.Address, error) {
	return _ValidSigners.Contract.Owner(&_ValidSigners.CallOpts)
}

// AddValidSigner is a paid mutator transaction binding the contract method 0xde8ee817.
//
// Solidity: function addValidSigner(_new_signer address) returns()
func (_ValidSigners *ValidSignersTransactor) AddValidSigner(opts *bind.TransactOpts, _new_signer common.Address) (*types.Transaction, error) {
	return _ValidSigners.contract.Transact(opts, "addValidSigner", _new_signer)
}

// AddValidSigner is a paid mutator transaction binding the contract method 0xde8ee817.
//
// Solidity: function addValidSigner(_new_signer address) returns()
func (_ValidSigners *ValidSignersSession) AddValidSigner(_new_signer common.Address) (*types.Transaction, error) {
	return _ValidSigners.Contract.AddValidSigner(&_ValidSigners.TransactOpts, _new_signer)
}

// AddValidSigner is a paid mutator transaction binding the contract method 0xde8ee817.
//
// Solidity: function addValidSigner(_new_signer address) returns()
func (_ValidSigners *ValidSignersTransactorSession) AddValidSigner(_new_signer common.Address) (*types.Transaction, error) {
	return _ValidSigners.Contract.AddValidSigner(&_ValidSigners.TransactOpts, _new_signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(_signer address) returns()
func (_ValidSigners *ValidSignersTransactor) RemoveSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _ValidSigners.contract.Transact(opts, "removeSigner", _signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(_signer address) returns()
func (_ValidSigners *ValidSignersSession) RemoveSigner(_signer common.Address) (*types.Transaction, error) {
	return _ValidSigners.Contract.RemoveSigner(&_ValidSigners.TransactOpts, _signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(_signer address) returns()
func (_ValidSigners *ValidSignersTransactorSession) RemoveSigner(_signer common.Address) (*types.Transaction, error) {
	return _ValidSigners.Contract.RemoveSigner(&_ValidSigners.TransactOpts, _signer)
}
