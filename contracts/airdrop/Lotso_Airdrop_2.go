// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package airdrop

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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInMerkle\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenAddressSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_airdrop\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_claim_times\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_total_claim_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040525f6005555f600655348015610017575f80fd5b50604051611292380380611292833981810160405281019061003991906103b8565b335f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100aa575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100a19190610405565b60405180910390fd5b6100b98161011760201b60201c565b50816002819055503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610110816101d860201b60201c565b505061041e565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6101e661026060201b60201c565b8060045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd82cad5fdc98633445b90f806f2e1a61a5409f92187ee9cd87f1da18c9069266816040516102559190610405565b60405180910390a150565b61026e6102f960201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1661029261030060201b60201c565b73ffffffffffffffffffffffffffffffffffffffff16146102f7576102bb6102f960201b60201c565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016102ee9190610405565b60405180910390fd5b565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f80fd5b5f819050919050565b61033d8161032b565b8114610347575f80fd5b50565b5f8151905061035881610334565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6103878261035e565b9050919050565b6103978161037d565b81146103a1575f80fd5b50565b5f815190506103b28161038e565b92915050565b5f80604083850312156103ce576103cd610327565b5b5f6103db8582860161034a565b92505060206103ec858286016103a4565b9150509250929050565b6103ff8161037d565b82525050565b5f6020820190506104185f8301846103f6565b92915050565b610e678061042b5f395ff3fe60806040526004361061009f575f3560e01c806372948db21161006357806372948db21461020257806373b2e80e1461022c5780638da5cb5b146102685780639d76ea5814610292578063b1000f23146102bc578063f2fde38b146102e657610144565b8063144fa6d7146101485780632eb4a7ab146101705780632f52ebb71461019a57806354f80248146101c2578063715018a6146101ec57610144565b366101445760015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc3490811502906040515f60405180830381858888f19350505050158015610108573d5f803e3d5ffd5b507f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874333460405161013a929190610a03565b60405180910390a1005b5f80fd5b348015610153575f80fd5b5061016e60048036038101906101699190610a5c565b61030e565b005b34801561017b575f80fd5b50610184610390565b6040516101919190610a9f565b60405180910390f35b3480156101a5575f80fd5b506101c060048036038101906101bb9190610b43565b610396565b005b3480156101cd575f80fd5b506101d6610687565b6040516101e39190610ba0565b60405180910390f35b3480156101f7575f80fd5b5061020061068d565b005b34801561020d575f80fd5b506102166106a0565b6040516102239190610ba0565b60405180910390f35b348015610237575f80fd5b50610252600480360381019061024d9190610a5c565b6106a6565b60405161025f9190610bd3565b60405180910390f35b348015610273575f80fd5b5061027c6106c3565b6040516102899190610bec565b60405180910390f35b34801561029d575f80fd5b506102a66106ea565b6040516102b39190610bec565b60405180910390f35b3480156102c7575f80fd5b506102d0610712565b6040516102dd9190610c25565b60405180910390f35b3480156102f1575f80fd5b5061030c60048036038101906103079190610a5c565b610737565b005b6103166107bb565b8060045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd82cad5fdc98633445b90f806f2e1a61a5409f92187ee9cd87f1da18c9069266816040516103859190610bec565b60405180910390a150565b60025481565b60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610417576040517f646cf55800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f338460405160200161042b929190610a03565b604051602081830303815290604052805190602001206040516020016104519190610c5e565b6040516020818303038152906040528051906020012090505f6104b78484808060200260200160405190810160405280939291908181526020018383602002808284375f81840152601f19601f8201169050808301925050505050505060025484610842565b9050806104f0576040517f8a585be200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060055f815461055390610ca5565b919050819055508460065f82825461056b9190610cec565b9250508190555060045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1633886040518463ffffffff1660e01b81526004016105f193929190610d7a565b6020604051808303815f875af115801561060d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106319190610dd9565b503373ffffffffffffffffffffffffffffffffffffffff167f47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4866040516106789190610ba0565b60405180910390a25050505050565b60065481565b6106956107bb565b61069e5f610858565b565b60055481565b6003602052805f5260405f205f915054906101000a900460ff1681565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61073f6107bb565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036107af575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016107a69190610bec565b60405180910390fd5b6107b881610858565b50565b6107c3610919565b73ffffffffffffffffffffffffffffffffffffffff166107e16106c3565b73ffffffffffffffffffffffffffffffffffffffff161461084057610804610919565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016108379190610bec565b60405180910390fd5b565b5f8261084e8584610920565b1490509392505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b5f808290505f5b8451811015610963576109548286838151811061094757610946610e04565b5b602002602001015161096e565b91508080600101915050610927565b508091505092915050565b5f818310610985576109808284610998565b610990565b61098f8383610998565b5b905092915050565b5f825f528160205260405f20905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6109d5826109ac565b9050919050565b6109e5816109cb565b82525050565b5f819050919050565b6109fd816109eb565b82525050565b5f604082019050610a165f8301856109dc565b610a2360208301846109f4565b9392505050565b5f80fd5b5f80fd5b610a3b816109cb565b8114610a45575f80fd5b50565b5f81359050610a5681610a32565b92915050565b5f60208284031215610a7157610a70610a2a565b5b5f610a7e84828501610a48565b91505092915050565b5f819050919050565b610a9981610a87565b82525050565b5f602082019050610ab25f830184610a90565b92915050565b610ac1816109eb565b8114610acb575f80fd5b50565b5f81359050610adc81610ab8565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112610b0357610b02610ae2565b5b8235905067ffffffffffffffff811115610b2057610b1f610ae6565b5b602083019150836020820283011115610b3c57610b3b610aea565b5b9250929050565b5f805f60408486031215610b5a57610b59610a2a565b5b5f610b6786828701610ace565b935050602084013567ffffffffffffffff811115610b8857610b87610a2e565b5b610b9486828701610aee565b92509250509250925092565b5f602082019050610bb35f8301846109f4565b92915050565b5f8115159050919050565b610bcd81610bb9565b82525050565b5f602082019050610be65f830184610bc4565b92915050565b5f602082019050610bff5f8301846109dc565b92915050565b5f610c0f826109ac565b9050919050565b610c1f81610c05565b82525050565b5f602082019050610c385f830184610c16565b92915050565b5f819050919050565b610c58610c5382610a87565b610c3e565b82525050565b5f610c698284610c47565b60208201915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610caf826109eb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610ce157610ce0610c78565b5b600182019050919050565b5f610cf6826109eb565b9150610d01836109eb565b9250828201905080821115610d1957610d18610c78565b5b92915050565b5f819050919050565b5f610d42610d3d610d38846109ac565b610d1f565b6109ac565b9050919050565b5f610d5382610d28565b9050919050565b5f610d6482610d49565b9050919050565b610d7481610d5a565b82525050565b5f606082019050610d8d5f830186610d6b565b610d9a60208301856109dc565b610da760408301846109f4565b949350505050565b610db881610bb9565b8114610dc2575f80fd5b50565b5f81519050610dd381610daf565b92915050565b5f60208284031215610dee57610ded610a2a565b5b5f610dfb84828501610dc5565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea264697066735822122027f9c955db175d153c5b0c8104b9909996aca7750d933c91d69417c03f6fca0b64736f6c63430008190033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _merkleRoot [32]byte, token common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _merkleRoot, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

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

// Airdrop is a free data retrieval call binding the contract method 0xb1000f23.
//
// Solidity: function _airdrop() view returns(address)
func (_Contracts *ContractsCaller) Airdrop(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "_airdrop")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Airdrop is a free data retrieval call binding the contract method 0xb1000f23.
//
// Solidity: function _airdrop() view returns(address)
func (_Contracts *ContractsSession) Airdrop() (common.Address, error) {
	return _Contracts.Contract.Airdrop(&_Contracts.CallOpts)
}

// Airdrop is a free data retrieval call binding the contract method 0xb1000f23.
//
// Solidity: function _airdrop() view returns(address)
func (_Contracts *ContractsCallerSession) Airdrop() (common.Address, error) {
	return _Contracts.Contract.Airdrop(&_Contracts.CallOpts)
}

// ClaimTimes is a free data retrieval call binding the contract method 0x72948db2.
//
// Solidity: function _claim_times() view returns(uint256)
func (_Contracts *ContractsCaller) ClaimTimes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "_claim_times")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimTimes is a free data retrieval call binding the contract method 0x72948db2.
//
// Solidity: function _claim_times() view returns(uint256)
func (_Contracts *ContractsSession) ClaimTimes() (*big.Int, error) {
	return _Contracts.Contract.ClaimTimes(&_Contracts.CallOpts)
}

// ClaimTimes is a free data retrieval call binding the contract method 0x72948db2.
//
// Solidity: function _claim_times() view returns(uint256)
func (_Contracts *ContractsCallerSession) ClaimTimes() (*big.Int, error) {
	return _Contracts.Contract.ClaimTimes(&_Contracts.CallOpts)
}

// TotalClaimAmount is a free data retrieval call binding the contract method 0x54f80248.
//
// Solidity: function _total_claim_amount() view returns(uint256)
func (_Contracts *ContractsCaller) TotalClaimAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "_total_claim_amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimAmount is a free data retrieval call binding the contract method 0x54f80248.
//
// Solidity: function _total_claim_amount() view returns(uint256)
func (_Contracts *ContractsSession) TotalClaimAmount() (*big.Int, error) {
	return _Contracts.Contract.TotalClaimAmount(&_Contracts.CallOpts)
}

// TotalClaimAmount is a free data retrieval call binding the contract method 0x54f80248.
//
// Solidity: function _total_claim_amount() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalClaimAmount() (*big.Int, error) {
	return _Contracts.Contract.TotalClaimAmount(&_Contracts.CallOpts)
}

// HasClaimed is a free data retrieval call binding the contract method 0x73b2e80e.
//
// Solidity: function hasClaimed(address ) view returns(bool)
func (_Contracts *ContractsCaller) HasClaimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "hasClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClaimed is a free data retrieval call binding the contract method 0x73b2e80e.
//
// Solidity: function hasClaimed(address ) view returns(bool)
func (_Contracts *ContractsSession) HasClaimed(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.HasClaimed(&_Contracts.CallOpts, arg0)
}

// HasClaimed is a free data retrieval call binding the contract method 0x73b2e80e.
//
// Solidity: function hasClaimed(address ) view returns(bool)
func (_Contracts *ContractsCallerSession) HasClaimed(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.HasClaimed(&_Contracts.CallOpts, arg0)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Contracts *ContractsCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Contracts *ContractsSession) MerkleRoot() ([32]byte, error) {
	return _Contracts.Contract.MerkleRoot(&_Contracts.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Contracts *ContractsCallerSession) MerkleRoot() ([32]byte, error) {
	return _Contracts.Contract.MerkleRoot(&_Contracts.CallOpts)
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

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Contracts *ContractsCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "tokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Contracts *ContractsSession) TokenAddress() (common.Address, error) {
	return _Contracts.Contract.TokenAddress(&_Contracts.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Contracts *ContractsCallerSession) TokenAddress() (common.Address, error) {
	return _Contracts.Contract.TokenAddress(&_Contracts.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x2f52ebb7.
//
// Solidity: function claim(uint256 amount, bytes32[] proof) returns()
func (_Contracts *ContractsTransactor) Claim(opts *bind.TransactOpts, amount *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "claim", amount, proof)
}

// Claim is a paid mutator transaction binding the contract method 0x2f52ebb7.
//
// Solidity: function claim(uint256 amount, bytes32[] proof) returns()
func (_Contracts *ContractsSession) Claim(amount *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Claim(&_Contracts.TransactOpts, amount, proof)
}

// Claim is a paid mutator transaction binding the contract method 0x2f52ebb7.
//
// Solidity: function claim(uint256 amount, bytes32[] proof) returns()
func (_Contracts *ContractsTransactorSession) Claim(amount *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Claim(&_Contracts.TransactOpts, amount, proof)
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

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address token) returns()
func (_Contracts *ContractsTransactor) SetToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setToken", token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address token) returns()
func (_Contracts *ContractsSession) SetToken(token common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetToken(&_Contracts.TransactOpts, token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address token) returns()
func (_Contracts *ContractsTransactorSession) SetToken(token common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetToken(&_Contracts.TransactOpts, token)
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

// ContractsClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Contracts contract.
type ContractsClaimIterator struct {
	Event *ContractsClaim // Event containing the contract specifics and raw log

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
func (it *ContractsClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsClaim)
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
		it.Event = new(ContractsClaim)
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
func (it *ContractsClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsClaim represents a Claim event raised by the Contracts contract.
type ContractsClaim struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed to, uint256 amount)
func (_Contracts *ContractsFilterer) FilterClaim(opts *bind.FilterOpts, to []common.Address) (*ContractsClaimIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Claim", toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsClaimIterator{contract: _Contracts.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed to, uint256 amount)
func (_Contracts *ContractsFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *ContractsClaim, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Claim", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsClaim)
				if err := _Contracts.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed to, uint256 amount)
func (_Contracts *ContractsFilterer) ParseClaim(log types.Log) (*ContractsClaim, error) {
	event := new(ContractsClaim)
	if err := _Contracts.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ContractsTokenAddressSetIterator is returned from FilterTokenAddressSet and is used to iterate over the raw logs and unpacked data for TokenAddressSet events raised by the Contracts contract.
type ContractsTokenAddressSetIterator struct {
	Event *ContractsTokenAddressSet // Event containing the contract specifics and raw log

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
func (it *ContractsTokenAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTokenAddressSet)
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
		it.Event = new(ContractsTokenAddressSet)
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
func (it *ContractsTokenAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTokenAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTokenAddressSet represents a TokenAddressSet event raised by the Contracts contract.
type ContractsTokenAddressSet struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTokenAddressSet is a free log retrieval operation binding the contract event 0xd82cad5fdc98633445b90f806f2e1a61a5409f92187ee9cd87f1da18c9069266.
//
// Solidity: event TokenAddressSet(address arg0)
func (_Contracts *ContractsFilterer) FilterTokenAddressSet(opts *bind.FilterOpts) (*ContractsTokenAddressSetIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "TokenAddressSet")
	if err != nil {
		return nil, err
	}
	return &ContractsTokenAddressSetIterator{contract: _Contracts.contract, event: "TokenAddressSet", logs: logs, sub: sub}, nil
}

// WatchTokenAddressSet is a free log subscription operation binding the contract event 0xd82cad5fdc98633445b90f806f2e1a61a5409f92187ee9cd87f1da18c9069266.
//
// Solidity: event TokenAddressSet(address arg0)
func (_Contracts *ContractsFilterer) WatchTokenAddressSet(opts *bind.WatchOpts, sink chan<- *ContractsTokenAddressSet) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "TokenAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTokenAddressSet)
				if err := _Contracts.contract.UnpackLog(event, "TokenAddressSet", log); err != nil {
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

// ParseTokenAddressSet is a log parse operation binding the contract event 0xd82cad5fdc98633445b90f806f2e1a61a5409f92187ee9cd87f1da18c9069266.
//
// Solidity: event TokenAddressSet(address arg0)
func (_Contracts *ContractsFilterer) ParseTokenAddressSet(log types.Log) (*ContractsTokenAddressSet, error) {
	event := new(ContractsTokenAddressSet)
	if err := _Contracts.contract.UnpackLog(event, "TokenAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
