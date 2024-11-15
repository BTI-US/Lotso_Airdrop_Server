// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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
	_ = abi.ConvertType
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"airdrop\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BuyEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ExcludeFromBuyEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ExcludeFromMaxHolding\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"FeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"IncludeInBuyEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"IncludeInMaxHolding\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"TradingPermitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_buy_enabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_trading_permitted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buy_enabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enable_buying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"excludeAccountsFromBuyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"excludeAccountsFromMaxHolding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"excludeFromBuyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"excludeFromMaxHolding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"includeAccountsInBuyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"includeAccountsInMaxHolding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"includeInBuyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"includeInMaxHolding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxHolding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumLotso_Token.SwapAction\",\"name\":\"action\",\"type\":\"uint8\"}],\"name\":\"taxes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradingPermitted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Pair\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Pair\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// BuyEnabled is a free data retrieval call binding the contract method 0x52863f26.
//
// Solidity: function _buy_enabled() view returns(bool)
func (_Contracts *ContractsCaller) BuyEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "_buy_enabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BuyEnabled is a free data retrieval call binding the contract method 0x52863f26.
//
// Solidity: function _buy_enabled() view returns(bool)
func (_Contracts *ContractsCallerSession) BuyEnabled() (bool, error) {
	return _Contracts.Contract.BuyEnabled(&_Contracts.CallOpts)
}

// TradingPermitted is a free data retrieval call binding the contract method 0x8a4f3187.
//
// Solidity: function _trading_permitted() view returns(bool)
func (_Contracts *ContractsCaller) TradingPermitted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "_trading_permitted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TradingPermitted is a free data retrieval call binding the contract method 0x8a4f3187.
//
// Solidity: function _trading_permitted() view returns(bool)
func (_Contracts *ContractsCallerSession) TradingPermitted() (bool, error) {
	return _Contracts.Contract.TradingPermitted(&_Contracts.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Contracts *ContractsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Contracts *ContractsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Contracts *ContractsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contracts *ContractsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contracts *ContractsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contracts *ContractsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsCallerSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// MaxHolding is a free data retrieval call binding the contract method 0xd4288744.
//
// Solidity: function maxHolding(address account) view returns(uint256)
func (_Contracts *ContractsCaller) MaxHolding(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "maxHolding", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxHolding is a free data retrieval call binding the contract method 0xd4288744.
//
// Solidity: function maxHolding(address account) view returns(uint256)
func (_Contracts *ContractsSession) MaxHolding(account common.Address) (*big.Int, error) {
	return _Contracts.Contract.MaxHolding(&_Contracts.CallOpts, account)
}

// MaxHolding is a free data retrieval call binding the contract method 0xd4288744.
//
// Solidity: function maxHolding(address account) view returns(uint256)
func (_Contracts *ContractsCallerSession) MaxHolding(account common.Address) (*big.Int, error) {
	return _Contracts.Contract.MaxHolding(&_Contracts.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCallerSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCallerSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// Taxes is a free data retrieval call binding the contract method 0x6d86ce4c.
//
// Solidity: function taxes(uint8 action) view returns(uint256)
func (_Contracts *ContractsCaller) Taxes(opts *bind.CallOpts, action uint8) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "taxes", action)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Taxes is a free data retrieval call binding the contract method 0x6d86ce4c.
//
// Solidity: function taxes(uint8 action) view returns(uint256)
func (_Contracts *ContractsSession) Taxes(action uint8) (*big.Int, error) {
	return _Contracts.Contract.Taxes(&_Contracts.CallOpts, action)
}

// Taxes is a free data retrieval call binding the contract method 0x6d86ce4c.
//
// Solidity: function taxes(uint8 action) view returns(uint256)
func (_Contracts *ContractsCallerSession) Taxes(action uint8) (*big.Int, error) {
	return _Contracts.Contract.Taxes(&_Contracts.CallOpts, action)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Contracts *ContractsCaller) UniswapV2Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "uniswapV2Pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Contracts *ContractsSession) UniswapV2Pair() (common.Address, error) {
	return _Contracts.Contract.UniswapV2Pair(&_Contracts.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Contracts *ContractsCallerSession) UniswapV2Pair() (common.Address, error) {
	return _Contracts.Contract.UniswapV2Pair(&_Contracts.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Contracts *ContractsCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Contracts *ContractsSession) UniswapV2Router() (common.Address, error) {
	return _Contracts.Contract.UniswapV2Router(&_Contracts.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Contracts *ContractsCallerSession) UniswapV2Router() (common.Address, error) {
	return _Contracts.Contract.UniswapV2Router(&_Contracts.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Contracts *ContractsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Contracts *ContractsSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Contracts *ContractsTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, spender, value)
}

// BuyEnabled is a paid mutator transaction binding the contract method 0xd3b07bec.
//
// Solidity: function buy_enabled() returns(bool)
func (_Contracts *ContractsTransactor) BuyEnabled(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "buy_enabled")
}

// BuyEnabled is a paid mutator transaction binding the contract method 0xd3b07bec.
//
// Solidity: function buy_enabled() returns(bool)
func (_Contracts *ContractsTransactorSession) BuyEnabled() (*types.Transaction, error) {
	return _Contracts.Contract.BuyEnabled(&_Contracts.TransactOpts)
}

// EnableBuying is a paid mutator transaction binding the contract method 0x9c4cf130.
//
// Solidity: function enable_buying() returns()
func (_Contracts *ContractsTransactor) EnableBuying(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "enable_buying")
}

// EnableBuying is a paid mutator transaction binding the contract method 0x9c4cf130.
//
// Solidity: function enable_buying() returns()
func (_Contracts *ContractsSession) EnableBuying() (*types.Transaction, error) {
	return _Contracts.Contract.EnableBuying(&_Contracts.TransactOpts)
}

// EnableBuying is a paid mutator transaction binding the contract method 0x9c4cf130.
//
// Solidity: function enable_buying() returns()
func (_Contracts *ContractsTransactorSession) EnableBuying() (*types.Transaction, error) {
	return _Contracts.Contract.EnableBuying(&_Contracts.TransactOpts)
}

// ExcludeAccountsFromBuyEnabled is a paid mutator transaction binding the contract method 0x58070c49.
//
// Solidity: function excludeAccountsFromBuyEnabled(address[] accounts) returns()
func (_Contracts *ContractsTransactor) ExcludeAccountsFromBuyEnabled(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "excludeAccountsFromBuyEnabled", accounts)
}

// ExcludeAccountsFromBuyEnabled is a paid mutator transaction binding the contract method 0x58070c49.
//
// Solidity: function excludeAccountsFromBuyEnabled(address[] accounts) returns()
func (_Contracts *ContractsSession) ExcludeAccountsFromBuyEnabled(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ExcludeAccountsFromBuyEnabled(&_Contracts.TransactOpts, accounts)
}

// ExcludeAccountsFromBuyEnabled is a paid mutator transaction binding the contract method 0x58070c49.
//
// Solidity: function excludeAccountsFromBuyEnabled(address[] accounts) returns()
func (_Contracts *ContractsTransactorSession) ExcludeAccountsFromBuyEnabled(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ExcludeAccountsFromBuyEnabled(&_Contracts.TransactOpts, accounts)
}

// ExcludeAccountsFromMaxHolding is a paid mutator transaction binding the contract method 0x08e74558.
//
// Solidity: function excludeAccountsFromMaxHolding(address[] accounts) returns()
func (_Contracts *ContractsTransactor) ExcludeAccountsFromMaxHolding(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "excludeAccountsFromMaxHolding", accounts)
}

// ExcludeAccountsFromMaxHolding is a paid mutator transaction binding the contract method 0x08e74558.
//
// Solidity: function excludeAccountsFromMaxHolding(address[] accounts) returns()
func (_Contracts *ContractsSession) ExcludeAccountsFromMaxHolding(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ExcludeAccountsFromMaxHolding(&_Contracts.TransactOpts, accounts)
}

// ExcludeAccountsFromMaxHolding is a paid mutator transaction binding the contract method 0x08e74558.
//
// Solidity: function excludeAccountsFromMaxHolding(address[] accounts) returns()
func (_Contracts *ContractsTransactorSession) ExcludeAccountsFromMaxHolding(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ExcludeAccountsFromMaxHolding(&_Contracts.TransactOpts, accounts)
}

// ExcludeFromBuyEnabled is a paid mutator transaction binding the contract method 0xc9563c6d.
//
// Solidity: function excludeFromBuyEnabled(address account) returns()
func (_Contracts *ContractsTransactor) ExcludeFromBuyEnabled(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "excludeFromBuyEnabled", account)
}

// ExcludeFromBuyEnabled is a paid mutator transaction binding the contract method 0xc9563c6d.
//
// Solidity: function excludeFromBuyEnabled(address account) returns()
func (_Contracts *ContractsSession) ExcludeFromBuyEnabled(account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ExcludeFromBuyEnabled(&_Contracts.TransactOpts, account)
}

// ExcludeFromBuyEnabled is a paid mutator transaction binding the contract method 0xc9563c6d.
//
// Solidity: function excludeFromBuyEnabled(address account) returns()
func (_Contracts *ContractsTransactorSession) ExcludeFromBuyEnabled(account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ExcludeFromBuyEnabled(&_Contracts.TransactOpts, account)
}

// ExcludeFromMaxHolding is a paid mutator transaction binding the contract method 0xcafe3059.
//
// Solidity: function excludeFromMaxHolding(address account) returns()
func (_Contracts *ContractsTransactor) ExcludeFromMaxHolding(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "excludeFromMaxHolding", account)
}

// ExcludeFromMaxHolding is a paid mutator transaction binding the contract method 0xcafe3059.
//
// Solidity: function excludeFromMaxHolding(address account) returns()
func (_Contracts *ContractsSession) ExcludeFromMaxHolding(account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ExcludeFromMaxHolding(&_Contracts.TransactOpts, account)
}

// ExcludeFromMaxHolding is a paid mutator transaction binding the contract method 0xcafe3059.
//
// Solidity: function excludeFromMaxHolding(address account) returns()
func (_Contracts *ContractsTransactorSession) ExcludeFromMaxHolding(account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ExcludeFromMaxHolding(&_Contracts.TransactOpts, account)
}

// IncludeAccountsInBuyEnabled is a paid mutator transaction binding the contract method 0xd77f73db.
//
// Solidity: function includeAccountsInBuyEnabled(address[] accounts) returns()
func (_Contracts *ContractsTransactor) IncludeAccountsInBuyEnabled(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "includeAccountsInBuyEnabled", accounts)
}

// IncludeAccountsInBuyEnabled is a paid mutator transaction binding the contract method 0xd77f73db.
//
// Solidity: function includeAccountsInBuyEnabled(address[] accounts) returns()
func (_Contracts *ContractsSession) IncludeAccountsInBuyEnabled(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.IncludeAccountsInBuyEnabled(&_Contracts.TransactOpts, accounts)
}

// IncludeAccountsInBuyEnabled is a paid mutator transaction binding the contract method 0xd77f73db.
//
// Solidity: function includeAccountsInBuyEnabled(address[] accounts) returns()
func (_Contracts *ContractsTransactorSession) IncludeAccountsInBuyEnabled(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.IncludeAccountsInBuyEnabled(&_Contracts.TransactOpts, accounts)
}

// IncludeAccountsInMaxHolding is a paid mutator transaction binding the contract method 0xbf9888a5.
//
// Solidity: function includeAccountsInMaxHolding(address[] accounts) returns()
func (_Contracts *ContractsTransactor) IncludeAccountsInMaxHolding(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "includeAccountsInMaxHolding", accounts)
}

// IncludeAccountsInMaxHolding is a paid mutator transaction binding the contract method 0xbf9888a5.
//
// Solidity: function includeAccountsInMaxHolding(address[] accounts) returns()
func (_Contracts *ContractsSession) IncludeAccountsInMaxHolding(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.IncludeAccountsInMaxHolding(&_Contracts.TransactOpts, accounts)
}

// IncludeAccountsInMaxHolding is a paid mutator transaction binding the contract method 0xbf9888a5.
//
// Solidity: function includeAccountsInMaxHolding(address[] accounts) returns()
func (_Contracts *ContractsTransactorSession) IncludeAccountsInMaxHolding(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.IncludeAccountsInMaxHolding(&_Contracts.TransactOpts, accounts)
}

// IncludeInBuyEnabled is a paid mutator transaction binding the contract method 0x6927701f.
//
// Solidity: function includeInBuyEnabled(address account) returns()
func (_Contracts *ContractsTransactor) IncludeInBuyEnabled(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "includeInBuyEnabled", account)
}

// IncludeInBuyEnabled is a paid mutator transaction binding the contract method 0x6927701f.
//
// Solidity: function includeInBuyEnabled(address account) returns()
func (_Contracts *ContractsSession) IncludeInBuyEnabled(account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.IncludeInBuyEnabled(&_Contracts.TransactOpts, account)
}

// IncludeInBuyEnabled is a paid mutator transaction binding the contract method 0x6927701f.
//
// Solidity: function includeInBuyEnabled(address account) returns()
func (_Contracts *ContractsTransactorSession) IncludeInBuyEnabled(account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.IncludeInBuyEnabled(&_Contracts.TransactOpts, account)
}

// IncludeInMaxHolding is a paid mutator transaction binding the contract method 0x0928a377.
//
// Solidity: function includeInMaxHolding(address account) returns()
func (_Contracts *ContractsTransactor) IncludeInMaxHolding(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "includeInMaxHolding", account)
}

// IncludeInMaxHolding is a paid mutator transaction binding the contract method 0x0928a377.
//
// Solidity: function includeInMaxHolding(address account) returns()
func (_Contracts *ContractsSession) IncludeInMaxHolding(account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.IncludeInMaxHolding(&_Contracts.TransactOpts, account)
}

// IncludeInMaxHolding is a paid mutator transaction binding the contract method 0x0928a377.
//
// Solidity: function includeInMaxHolding(address account) returns()
func (_Contracts *ContractsTransactorSession) IncludeInMaxHolding(account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.IncludeInMaxHolding(&_Contracts.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0xb492c975.
//
// Solidity: function setFee(uint256[] fees) returns()
func (_Contracts *ContractsTransactor) SetFee(opts *bind.TransactOpts, fees []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setFee", fees)
}

// SetFee is a paid mutator transaction binding the contract method 0xb492c975.
//
// Solidity: function setFee(uint256[] fees) returns()
func (_Contracts *ContractsSession) SetFee(fees []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetFee(&_Contracts.TransactOpts, fees)
}

// SetFee is a paid mutator transaction binding the contract method 0xb492c975.
//
// Solidity: function setFee(uint256[] fees) returns()
func (_Contracts *ContractsTransactorSession) SetFee(fees []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetFee(&_Contracts.TransactOpts, fees)
}

// TradingPermitted is a paid mutator transaction binding the contract method 0x6772c962.
//
// Solidity: function tradingPermitted() returns(bool)
func (_Contracts *ContractsTransactor) TradingPermitted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "tradingPermitted")
}

// TradingPermitted is a paid mutator transaction binding the contract method 0x6772c962.
//
// Solidity: function tradingPermitted() returns(bool)
func (_Contracts *ContractsTransactorSession) TradingPermitted() (*types.Transaction, error) {
	return _Contracts.Contract.TradingPermitted(&_Contracts.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Contracts *ContractsTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Contracts *ContractsSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Contracts *ContractsTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Contracts *ContractsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Contracts *ContractsSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Contracts *ContractsTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactorSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// ContractsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contracts contract.
type ContractsApprovalIterator struct {
	Event *ContractsApproval // Event containing the contract specifics and raw log

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
func (it *ContractsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsApproval)
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
		it.Event = new(ContractsApproval)
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
func (it *ContractsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsApproval represents a Approval event raised by the Contracts contract.
type ContractsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsApprovalIterator{contract: _Contracts.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsApproval)
				if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) ParseApproval(log types.Log) (*ContractsApproval, error) {
	event := new(ContractsApproval)
	if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsBuyEnabledIterator is returned from FilterBuyEnabled and is used to iterate over the raw logs and unpacked data for BuyEnabled events raised by the Contracts contract.
type ContractsBuyEnabledIterator struct {
	Event *ContractsBuyEnabled // Event containing the contract specifics and raw log

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
func (it *ContractsBuyEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsBuyEnabled)
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
		it.Event = new(ContractsBuyEnabled)
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
func (it *ContractsBuyEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsBuyEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsBuyEnabled represents a BuyEnabled event raised by the Contracts contract.
type ContractsBuyEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBuyEnabled is a free log retrieval operation binding the contract event 0xf191476d5ac164e968301c74796f8041a91fd743dc9229e78a8629fa07a1890b.
//
// Solidity: event BuyEnabled()
func (_Contracts *ContractsFilterer) FilterBuyEnabled(opts *bind.FilterOpts) (*ContractsBuyEnabledIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "BuyEnabled")
	if err != nil {
		return nil, err
	}
	return &ContractsBuyEnabledIterator{contract: _Contracts.contract, event: "BuyEnabled", logs: logs, sub: sub}, nil
}

// WatchBuyEnabled is a free log subscription operation binding the contract event 0xf191476d5ac164e968301c74796f8041a91fd743dc9229e78a8629fa07a1890b.
//
// Solidity: event BuyEnabled()
func (_Contracts *ContractsFilterer) WatchBuyEnabled(opts *bind.WatchOpts, sink chan<- *ContractsBuyEnabled) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "BuyEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsBuyEnabled)
				if err := _Contracts.contract.UnpackLog(event, "BuyEnabled", log); err != nil {
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

// ParseBuyEnabled is a log parse operation binding the contract event 0xf191476d5ac164e968301c74796f8041a91fd743dc9229e78a8629fa07a1890b.
//
// Solidity: event BuyEnabled()
func (_Contracts *ContractsFilterer) ParseBuyEnabled(log types.Log) (*ContractsBuyEnabled, error) {
	event := new(ContractsBuyEnabled)
	if err := _Contracts.contract.UnpackLog(event, "BuyEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsExcludeFromBuyEnabledIterator is returned from FilterExcludeFromBuyEnabled and is used to iterate over the raw logs and unpacked data for ExcludeFromBuyEnabled events raised by the Contracts contract.
type ContractsExcludeFromBuyEnabledIterator struct {
	Event *ContractsExcludeFromBuyEnabled // Event containing the contract specifics and raw log

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
func (it *ContractsExcludeFromBuyEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsExcludeFromBuyEnabled)
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
		it.Event = new(ContractsExcludeFromBuyEnabled)
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
func (it *ContractsExcludeFromBuyEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsExcludeFromBuyEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsExcludeFromBuyEnabled represents a ExcludeFromBuyEnabled event raised by the Contracts contract.
type ContractsExcludeFromBuyEnabled struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExcludeFromBuyEnabled is a free log retrieval operation binding the contract event 0xc42c484b0d041db9b6f64e8ac17b1732bb306cdd621699b45d37db903d6b6d2c.
//
// Solidity: event ExcludeFromBuyEnabled(address arg0)
func (_Contracts *ContractsFilterer) FilterExcludeFromBuyEnabled(opts *bind.FilterOpts) (*ContractsExcludeFromBuyEnabledIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ExcludeFromBuyEnabled")
	if err != nil {
		return nil, err
	}
	return &ContractsExcludeFromBuyEnabledIterator{contract: _Contracts.contract, event: "ExcludeFromBuyEnabled", logs: logs, sub: sub}, nil
}

// WatchExcludeFromBuyEnabled is a free log subscription operation binding the contract event 0xc42c484b0d041db9b6f64e8ac17b1732bb306cdd621699b45d37db903d6b6d2c.
//
// Solidity: event ExcludeFromBuyEnabled(address arg0)
func (_Contracts *ContractsFilterer) WatchExcludeFromBuyEnabled(opts *bind.WatchOpts, sink chan<- *ContractsExcludeFromBuyEnabled) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ExcludeFromBuyEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsExcludeFromBuyEnabled)
				if err := _Contracts.contract.UnpackLog(event, "ExcludeFromBuyEnabled", log); err != nil {
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

// ParseExcludeFromBuyEnabled is a log parse operation binding the contract event 0xc42c484b0d041db9b6f64e8ac17b1732bb306cdd621699b45d37db903d6b6d2c.
//
// Solidity: event ExcludeFromBuyEnabled(address arg0)
func (_Contracts *ContractsFilterer) ParseExcludeFromBuyEnabled(log types.Log) (*ContractsExcludeFromBuyEnabled, error) {
	event := new(ContractsExcludeFromBuyEnabled)
	if err := _Contracts.contract.UnpackLog(event, "ExcludeFromBuyEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsExcludeFromMaxHoldingIterator is returned from FilterExcludeFromMaxHolding and is used to iterate over the raw logs and unpacked data for ExcludeFromMaxHolding events raised by the Contracts contract.
type ContractsExcludeFromMaxHoldingIterator struct {
	Event *ContractsExcludeFromMaxHolding // Event containing the contract specifics and raw log

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
func (it *ContractsExcludeFromMaxHoldingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsExcludeFromMaxHolding)
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
		it.Event = new(ContractsExcludeFromMaxHolding)
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
func (it *ContractsExcludeFromMaxHoldingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsExcludeFromMaxHoldingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsExcludeFromMaxHolding represents a ExcludeFromMaxHolding event raised by the Contracts contract.
type ContractsExcludeFromMaxHolding struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExcludeFromMaxHolding is a free log retrieval operation binding the contract event 0x1f6e28293d10080f233997eda6c3a55282eeaf225211575a55bc16b67b27f96d.
//
// Solidity: event ExcludeFromMaxHolding(address arg0)
func (_Contracts *ContractsFilterer) FilterExcludeFromMaxHolding(opts *bind.FilterOpts) (*ContractsExcludeFromMaxHoldingIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ExcludeFromMaxHolding")
	if err != nil {
		return nil, err
	}
	return &ContractsExcludeFromMaxHoldingIterator{contract: _Contracts.contract, event: "ExcludeFromMaxHolding", logs: logs, sub: sub}, nil
}

// WatchExcludeFromMaxHolding is a free log subscription operation binding the contract event 0x1f6e28293d10080f233997eda6c3a55282eeaf225211575a55bc16b67b27f96d.
//
// Solidity: event ExcludeFromMaxHolding(address arg0)
func (_Contracts *ContractsFilterer) WatchExcludeFromMaxHolding(opts *bind.WatchOpts, sink chan<- *ContractsExcludeFromMaxHolding) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ExcludeFromMaxHolding")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsExcludeFromMaxHolding)
				if err := _Contracts.contract.UnpackLog(event, "ExcludeFromMaxHolding", log); err != nil {
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

// ParseExcludeFromMaxHolding is a log parse operation binding the contract event 0x1f6e28293d10080f233997eda6c3a55282eeaf225211575a55bc16b67b27f96d.
//
// Solidity: event ExcludeFromMaxHolding(address arg0)
func (_Contracts *ContractsFilterer) ParseExcludeFromMaxHolding(log types.Log) (*ContractsExcludeFromMaxHolding, error) {
	event := new(ContractsExcludeFromMaxHolding)
	if err := _Contracts.contract.UnpackLog(event, "ExcludeFromMaxHolding", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsFeeSetIterator is returned from FilterFeeSet and is used to iterate over the raw logs and unpacked data for FeeSet events raised by the Contracts contract.
type ContractsFeeSetIterator struct {
	Event *ContractsFeeSet // Event containing the contract specifics and raw log

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
func (it *ContractsFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFeeSet)
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
		it.Event = new(ContractsFeeSet)
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
func (it *ContractsFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFeeSet represents a FeeSet event raised by the Contracts contract.
type ContractsFeeSet struct {
	Arg0 []*big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFeeSet is a free log retrieval operation binding the contract event 0x8bdf52d6634fa5eb46f4c9600b70e9cb52175d86385a372a20299138960cdb04.
//
// Solidity: event FeeSet(uint256[] arg0)
func (_Contracts *ContractsFilterer) FilterFeeSet(opts *bind.FilterOpts) (*ContractsFeeSetIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FeeSet")
	if err != nil {
		return nil, err
	}
	return &ContractsFeeSetIterator{contract: _Contracts.contract, event: "FeeSet", logs: logs, sub: sub}, nil
}

// WatchFeeSet is a free log subscription operation binding the contract event 0x8bdf52d6634fa5eb46f4c9600b70e9cb52175d86385a372a20299138960cdb04.
//
// Solidity: event FeeSet(uint256[] arg0)
func (_Contracts *ContractsFilterer) WatchFeeSet(opts *bind.WatchOpts, sink chan<- *ContractsFeeSet) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFeeSet)
				if err := _Contracts.contract.UnpackLog(event, "FeeSet", log); err != nil {
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

// ParseFeeSet is a log parse operation binding the contract event 0x8bdf52d6634fa5eb46f4c9600b70e9cb52175d86385a372a20299138960cdb04.
//
// Solidity: event FeeSet(uint256[] arg0)
func (_Contracts *ContractsFilterer) ParseFeeSet(log types.Log) (*ContractsFeeSet, error) {
	event := new(ContractsFeeSet)
	if err := _Contracts.contract.UnpackLog(event, "FeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsIncludeInBuyEnabledIterator is returned from FilterIncludeInBuyEnabled and is used to iterate over the raw logs and unpacked data for IncludeInBuyEnabled events raised by the Contracts contract.
type ContractsIncludeInBuyEnabledIterator struct {
	Event *ContractsIncludeInBuyEnabled // Event containing the contract specifics and raw log

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
func (it *ContractsIncludeInBuyEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsIncludeInBuyEnabled)
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
		it.Event = new(ContractsIncludeInBuyEnabled)
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
func (it *ContractsIncludeInBuyEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsIncludeInBuyEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsIncludeInBuyEnabled represents a IncludeInBuyEnabled event raised by the Contracts contract.
type ContractsIncludeInBuyEnabled struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterIncludeInBuyEnabled is a free log retrieval operation binding the contract event 0x00e7a6a058ff3c257bb0a9c2e25b8a0d82c1ad82361717a81c9670515405f7b8.
//
// Solidity: event IncludeInBuyEnabled(address arg0)
func (_Contracts *ContractsFilterer) FilterIncludeInBuyEnabled(opts *bind.FilterOpts) (*ContractsIncludeInBuyEnabledIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "IncludeInBuyEnabled")
	if err != nil {
		return nil, err
	}
	return &ContractsIncludeInBuyEnabledIterator{contract: _Contracts.contract, event: "IncludeInBuyEnabled", logs: logs, sub: sub}, nil
}

// WatchIncludeInBuyEnabled is a free log subscription operation binding the contract event 0x00e7a6a058ff3c257bb0a9c2e25b8a0d82c1ad82361717a81c9670515405f7b8.
//
// Solidity: event IncludeInBuyEnabled(address arg0)
func (_Contracts *ContractsFilterer) WatchIncludeInBuyEnabled(opts *bind.WatchOpts, sink chan<- *ContractsIncludeInBuyEnabled) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "IncludeInBuyEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsIncludeInBuyEnabled)
				if err := _Contracts.contract.UnpackLog(event, "IncludeInBuyEnabled", log); err != nil {
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

// ParseIncludeInBuyEnabled is a log parse operation binding the contract event 0x00e7a6a058ff3c257bb0a9c2e25b8a0d82c1ad82361717a81c9670515405f7b8.
//
// Solidity: event IncludeInBuyEnabled(address arg0)
func (_Contracts *ContractsFilterer) ParseIncludeInBuyEnabled(log types.Log) (*ContractsIncludeInBuyEnabled, error) {
	event := new(ContractsIncludeInBuyEnabled)
	if err := _Contracts.contract.UnpackLog(event, "IncludeInBuyEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsIncludeInMaxHoldingIterator is returned from FilterIncludeInMaxHolding and is used to iterate over the raw logs and unpacked data for IncludeInMaxHolding events raised by the Contracts contract.
type ContractsIncludeInMaxHoldingIterator struct {
	Event *ContractsIncludeInMaxHolding // Event containing the contract specifics and raw log

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
func (it *ContractsIncludeInMaxHoldingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsIncludeInMaxHolding)
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
		it.Event = new(ContractsIncludeInMaxHolding)
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
func (it *ContractsIncludeInMaxHoldingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsIncludeInMaxHoldingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsIncludeInMaxHolding represents a IncludeInMaxHolding event raised by the Contracts contract.
type ContractsIncludeInMaxHolding struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterIncludeInMaxHolding is a free log retrieval operation binding the contract event 0xd4a1693e17a8659123551fc8a66e4fefb47321e576eb276300f305bf1aa82d1a.
//
// Solidity: event IncludeInMaxHolding(address arg0)
func (_Contracts *ContractsFilterer) FilterIncludeInMaxHolding(opts *bind.FilterOpts) (*ContractsIncludeInMaxHoldingIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "IncludeInMaxHolding")
	if err != nil {
		return nil, err
	}
	return &ContractsIncludeInMaxHoldingIterator{contract: _Contracts.contract, event: "IncludeInMaxHolding", logs: logs, sub: sub}, nil
}

// WatchIncludeInMaxHolding is a free log subscription operation binding the contract event 0xd4a1693e17a8659123551fc8a66e4fefb47321e576eb276300f305bf1aa82d1a.
//
// Solidity: event IncludeInMaxHolding(address arg0)
func (_Contracts *ContractsFilterer) WatchIncludeInMaxHolding(opts *bind.WatchOpts, sink chan<- *ContractsIncludeInMaxHolding) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "IncludeInMaxHolding")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsIncludeInMaxHolding)
				if err := _Contracts.contract.UnpackLog(event, "IncludeInMaxHolding", log); err != nil {
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

// ParseIncludeInMaxHolding is a log parse operation binding the contract event 0xd4a1693e17a8659123551fc8a66e4fefb47321e576eb276300f305bf1aa82d1a.
//
// Solidity: event IncludeInMaxHolding(address arg0)
func (_Contracts *ContractsFilterer) ParseIncludeInMaxHolding(log types.Log) (*ContractsIncludeInMaxHolding, error) {
	event := new(ContractsIncludeInMaxHolding)
	if err := _Contracts.contract.UnpackLog(event, "IncludeInMaxHolding", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
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
		it.Event = new(ContractsOwnershipTransferred)
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
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Contracts contract.
type ContractsReceivedIterator struct {
	Event *ContractsReceived // Event containing the contract specifics and raw log

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
func (it *ContractsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsReceived)
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
		it.Event = new(ContractsReceived)
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
func (it *ContractsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsReceived represents a Received event raised by the Contracts contract.
type ContractsReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Contracts *ContractsFilterer) FilterReceived(opts *bind.FilterOpts) (*ContractsReceivedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &ContractsReceivedIterator{contract: _Contracts.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Contracts *ContractsFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *ContractsReceived) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsReceived)
				if err := _Contracts.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Contracts *ContractsFilterer) ParseReceived(log types.Log) (*ContractsReceived, error) {
	event := new(ContractsReceived)
	if err := _Contracts.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTradingPermittedIterator is returned from FilterTradingPermitted and is used to iterate over the raw logs and unpacked data for TradingPermitted events raised by the Contracts contract.
type ContractsTradingPermittedIterator struct {
	Event *ContractsTradingPermitted // Event containing the contract specifics and raw log

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
func (it *ContractsTradingPermittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTradingPermitted)
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
		it.Event = new(ContractsTradingPermitted)
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
func (it *ContractsTradingPermittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTradingPermittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTradingPermitted represents a TradingPermitted event raised by the Contracts contract.
type ContractsTradingPermitted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTradingPermitted is a free log retrieval operation binding the contract event 0xf260d48b6b88ef9460726001140923e2dce6a7ff7c9310a2a1df6d09ebabf00c.
//
// Solidity: event TradingPermitted()
func (_Contracts *ContractsFilterer) FilterTradingPermitted(opts *bind.FilterOpts) (*ContractsTradingPermittedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "TradingPermitted")
	if err != nil {
		return nil, err
	}
	return &ContractsTradingPermittedIterator{contract: _Contracts.contract, event: "TradingPermitted", logs: logs, sub: sub}, nil
}

// WatchTradingPermitted is a free log subscription operation binding the contract event 0xf260d48b6b88ef9460726001140923e2dce6a7ff7c9310a2a1df6d09ebabf00c.
//
// Solidity: event TradingPermitted()
func (_Contracts *ContractsFilterer) WatchTradingPermitted(opts *bind.WatchOpts, sink chan<- *ContractsTradingPermitted) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "TradingPermitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTradingPermitted)
				if err := _Contracts.contract.UnpackLog(event, "TradingPermitted", log); err != nil {
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

// ParseTradingPermitted is a log parse operation binding the contract event 0xf260d48b6b88ef9460726001140923e2dce6a7ff7c9310a2a1df6d09ebabf00c.
//
// Solidity: event TradingPermitted()
func (_Contracts *ContractsFilterer) ParseTradingPermitted(log types.Log) (*ContractsTradingPermitted, error) {
	event := new(ContractsTradingPermitted)
	if err := _Contracts.contract.UnpackLog(event, "TradingPermitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contracts contract.
type ContractsTransferIterator struct {
	Event *ContractsTransfer // Event containing the contract specifics and raw log

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
func (it *ContractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTransfer)
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
		it.Event = new(ContractsTransfer)
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
func (it *ContractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTransfer represents a Transfer event raised by the Contracts contract.
type ContractsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTransferIterator{contract: _Contracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTransfer)
				if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) ParseTransfer(log types.Log) (*ContractsTransfer, error) {
	event := new(ContractsTransfer)
	if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
