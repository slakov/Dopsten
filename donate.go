// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package donate

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DonateABI is the input ABI used to generate the binding from.
const DonateABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// Donate is an auto generated Go binding around an Ethereum contract.
type Donate struct {
	DonateCaller     // Read-only binding to the contract
	DonateTransactor // Write-only binding to the contract
	DonateFilterer   // Log filterer for contract events
}

// DonateCaller is an auto generated read-only Go binding around an Ethereum contract.
type DonateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DonateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DonateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DonateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DonateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DonateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DonateSession struct {
	Contract     *Donate           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DonateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DonateCallerSession struct {
	Contract *DonateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DonateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DonateTransactorSession struct {
	Contract     *DonateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DonateRaw is an auto generated low-level Go binding around an Ethereum contract.
type DonateRaw struct {
	Contract *Donate // Generic contract binding to access the raw methods on
}

// DonateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DonateCallerRaw struct {
	Contract *DonateCaller // Generic read-only contract binding to access the raw methods on
}

// DonateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DonateTransactorRaw struct {
	Contract *DonateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDonate creates a new instance of Donate, bound to a specific deployed contract.
func NewDonate(address common.Address, backend bind.ContractBackend) (*Donate, error) {
	contract, err := bindDonate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Donate{DonateCaller: DonateCaller{contract: contract}, DonateTransactor: DonateTransactor{contract: contract}, DonateFilterer: DonateFilterer{contract: contract}}, nil
}

// NewDonateCaller creates a new read-only instance of Donate, bound to a specific deployed contract.
func NewDonateCaller(address common.Address, caller bind.ContractCaller) (*DonateCaller, error) {
	contract, err := bindDonate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DonateCaller{contract: contract}, nil
}

// NewDonateTransactor creates a new write-only instance of Donate, bound to a specific deployed contract.
func NewDonateTransactor(address common.Address, transactor bind.ContractTransactor) (*DonateTransactor, error) {
	contract, err := bindDonate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DonateTransactor{contract: contract}, nil
}

// NewDonateFilterer creates a new log filterer instance of Donate, bound to a specific deployed contract.
func NewDonateFilterer(address common.Address, filterer bind.ContractFilterer) (*DonateFilterer, error) {
	contract, err := bindDonate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DonateFilterer{contract: contract}, nil
}

// bindDonate binds a generic wrapper to an already deployed contract.
func bindDonate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DonateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Donate *DonateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Donate.Contract.DonateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Donate *DonateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Donate.Contract.DonateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Donate *DonateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Donate.Contract.DonateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Donate *DonateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Donate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Donate *DonateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Donate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Donate *DonateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Donate.Contract.contract.Transact(opts, method, params...)
}

// AmountRaised is a free data retrieval call binding the contract method 0x7b3e5e7b.
//
// Solidity: function amountRaised() constant returns(uint256)
func (_Donate *DonateCaller) AmountRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Donate.contract.Call(opts, out, "amountRaised")
	return *ret0, err
}

// AmountRaised is a free data retrieval call binding the contract method 0x7b3e5e7b.
//
// Solidity: function amountRaised() constant returns(uint256)
func (_Donate *DonateSession) AmountRaised() (*big.Int, error) {
	return _Donate.Contract.AmountRaised(&_Donate.CallOpts)
}

// AmountRaised is a free data retrieval call binding the contract method 0x7b3e5e7b.
//
// Solidity: function amountRaised() constant returns(uint256)
func (_Donate *DonateCallerSession) AmountRaised() (*big.Int, error) {
	return _Donate.Contract.AmountRaised(&_Donate.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x12514bba.
//
// Solidity: function transfer(_value uint256) returns(success bool)
func (_Donate *DonateTransactor) Transfer(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Donate.contract.Transact(opts, "transfer", _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x12514bba.
//
// Solidity: function transfer(_value uint256) returns(success bool)
func (_Donate *DonateSession) Transfer(_value *big.Int) (*types.Transaction, error) {
	return _Donate.Contract.Transfer(&_Donate.TransactOpts, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x12514bba.
//
// Solidity: function transfer(_value uint256) returns(success bool)
func (_Donate *DonateTransactorSession) Transfer(_value *big.Int) (*types.Transaction, error) {
	return _Donate.Contract.Transfer(&_Donate.TransactOpts, _value)
}
