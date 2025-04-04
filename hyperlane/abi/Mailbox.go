// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hyperlaneabi

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

// HyperlaneMetaData contains all meta data concerning the Hyperlane contract.
var HyperlaneMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"DefaultHookSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DefaultIsmSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"DispatchId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"Process\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"ProcessId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"RequiredHookSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultHook\",\"outputs\":[{\"internalType\":\"contractIPostDispatchHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultIsm\",\"outputs\":[{\"internalType\":\"contractIInterchainSecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"delivered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"contractIPostDispatchHook\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"hookMetadata\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultIsm\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultHook\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_requiredHook\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestDispatchedId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"process\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"processedAt\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"processor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"contractIPostDispatchHook\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"quoteDispatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"quoteDispatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"defaultHookMetadata\",\"type\":\"bytes\"}],\"name\":\"quoteDispatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"recipientIsm\",\"outputs\":[{\"internalType\":\"contractIInterchainSecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredHook\",\"outputs\":[{\"internalType\":\"contractIPostDispatchHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hook\",\"type\":\"address\"}],\"name\":\"setDefaultHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_module\",\"type\":\"address\"}],\"name\":\"setDefaultIsm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hook\",\"type\":\"address\"}],\"name\":\"setRequiredHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HyperlaneABI is the input ABI used to generate the binding from.
// Deprecated: Use HyperlaneMetaData.ABI instead.
var HyperlaneABI = HyperlaneMetaData.ABI

// Hyperlane is an auto generated Go binding around an Ethereum contract.
type Hyperlane struct {
	HyperlaneCaller     // Read-only binding to the contract
	HyperlaneTransactor // Write-only binding to the contract
	HyperlaneFilterer   // Log filterer for contract events
}

// HyperlaneCaller is an auto generated read-only Go binding around an Ethereum contract.
type HyperlaneCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperlaneTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HyperlaneTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperlaneFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HyperlaneFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperlaneSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HyperlaneSession struct {
	Contract     *Hyperlane        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HyperlaneCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HyperlaneCallerSession struct {
	Contract *HyperlaneCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HyperlaneTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HyperlaneTransactorSession struct {
	Contract     *HyperlaneTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HyperlaneRaw is an auto generated low-level Go binding around an Ethereum contract.
type HyperlaneRaw struct {
	Contract *Hyperlane // Generic contract binding to access the raw methods on
}

// HyperlaneCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HyperlaneCallerRaw struct {
	Contract *HyperlaneCaller // Generic read-only contract binding to access the raw methods on
}

// HyperlaneTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HyperlaneTransactorRaw struct {
	Contract *HyperlaneTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHyperlane creates a new instance of Hyperlane, bound to a specific deployed contract.
func NewHyperlane(address common.Address, backend bind.ContractBackend) (*Hyperlane, error) {
	contract, err := bindHyperlane(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hyperlane{HyperlaneCaller: HyperlaneCaller{contract: contract}, HyperlaneTransactor: HyperlaneTransactor{contract: contract}, HyperlaneFilterer: HyperlaneFilterer{contract: contract}}, nil
}

// NewHyperlaneCaller creates a new read-only instance of Hyperlane, bound to a specific deployed contract.
func NewHyperlaneCaller(address common.Address, caller bind.ContractCaller) (*HyperlaneCaller, error) {
	contract, err := bindHyperlane(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HyperlaneCaller{contract: contract}, nil
}

// NewHyperlaneTransactor creates a new write-only instance of Hyperlane, bound to a specific deployed contract.
func NewHyperlaneTransactor(address common.Address, transactor bind.ContractTransactor) (*HyperlaneTransactor, error) {
	contract, err := bindHyperlane(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HyperlaneTransactor{contract: contract}, nil
}

// NewHyperlaneFilterer creates a new log filterer instance of Hyperlane, bound to a specific deployed contract.
func NewHyperlaneFilterer(address common.Address, filterer bind.ContractFilterer) (*HyperlaneFilterer, error) {
	contract, err := bindHyperlane(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HyperlaneFilterer{contract: contract}, nil
}

// bindHyperlane binds a generic wrapper to an already deployed contract.
func bindHyperlane(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HyperlaneMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hyperlane *HyperlaneRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hyperlane.Contract.HyperlaneCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hyperlane *HyperlaneRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyperlane.Contract.HyperlaneTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hyperlane *HyperlaneRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hyperlane.Contract.HyperlaneTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hyperlane *HyperlaneCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hyperlane.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hyperlane *HyperlaneTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyperlane.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hyperlane *HyperlaneTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hyperlane.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Hyperlane *HyperlaneCaller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Hyperlane *HyperlaneSession) VERSION() (uint8, error) {
	return _Hyperlane.Contract.VERSION(&_Hyperlane.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Hyperlane *HyperlaneCallerSession) VERSION() (uint8, error) {
	return _Hyperlane.Contract.VERSION(&_Hyperlane.CallOpts)
}

// DefaultHook is a free data retrieval call binding the contract method 0x3d1250b7.
//
// Solidity: function defaultHook() view returns(address)
func (_Hyperlane *HyperlaneCaller) DefaultHook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "defaultHook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultHook is a free data retrieval call binding the contract method 0x3d1250b7.
//
// Solidity: function defaultHook() view returns(address)
func (_Hyperlane *HyperlaneSession) DefaultHook() (common.Address, error) {
	return _Hyperlane.Contract.DefaultHook(&_Hyperlane.CallOpts)
}

// DefaultHook is a free data retrieval call binding the contract method 0x3d1250b7.
//
// Solidity: function defaultHook() view returns(address)
func (_Hyperlane *HyperlaneCallerSession) DefaultHook() (common.Address, error) {
	return _Hyperlane.Contract.DefaultHook(&_Hyperlane.CallOpts)
}

// DefaultIsm is a free data retrieval call binding the contract method 0x6e5f516e.
//
// Solidity: function defaultIsm() view returns(address)
func (_Hyperlane *HyperlaneCaller) DefaultIsm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "defaultIsm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultIsm is a free data retrieval call binding the contract method 0x6e5f516e.
//
// Solidity: function defaultIsm() view returns(address)
func (_Hyperlane *HyperlaneSession) DefaultIsm() (common.Address, error) {
	return _Hyperlane.Contract.DefaultIsm(&_Hyperlane.CallOpts)
}

// DefaultIsm is a free data retrieval call binding the contract method 0x6e5f516e.
//
// Solidity: function defaultIsm() view returns(address)
func (_Hyperlane *HyperlaneCallerSession) DefaultIsm() (common.Address, error) {
	return _Hyperlane.Contract.DefaultIsm(&_Hyperlane.CallOpts)
}

// Delivered is a free data retrieval call binding the contract method 0xe495f1d4.
//
// Solidity: function delivered(bytes32 _id) view returns(bool)
func (_Hyperlane *HyperlaneCaller) Delivered(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "delivered", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Delivered is a free data retrieval call binding the contract method 0xe495f1d4.
//
// Solidity: function delivered(bytes32 _id) view returns(bool)
func (_Hyperlane *HyperlaneSession) Delivered(_id [32]byte) (bool, error) {
	return _Hyperlane.Contract.Delivered(&_Hyperlane.CallOpts, _id)
}

// Delivered is a free data retrieval call binding the contract method 0xe495f1d4.
//
// Solidity: function delivered(bytes32 _id) view returns(bool)
func (_Hyperlane *HyperlaneCallerSession) Delivered(_id [32]byte) (bool, error) {
	return _Hyperlane.Contract.Delivered(&_Hyperlane.CallOpts, _id)
}

// DeployedBlock is a free data retrieval call binding the contract method 0x82ea7bfe.
//
// Solidity: function deployedBlock() view returns(uint256)
func (_Hyperlane *HyperlaneCaller) DeployedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "deployedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeployedBlock is a free data retrieval call binding the contract method 0x82ea7bfe.
//
// Solidity: function deployedBlock() view returns(uint256)
func (_Hyperlane *HyperlaneSession) DeployedBlock() (*big.Int, error) {
	return _Hyperlane.Contract.DeployedBlock(&_Hyperlane.CallOpts)
}

// DeployedBlock is a free data retrieval call binding the contract method 0x82ea7bfe.
//
// Solidity: function deployedBlock() view returns(uint256)
func (_Hyperlane *HyperlaneCallerSession) DeployedBlock() (*big.Int, error) {
	return _Hyperlane.Contract.DeployedBlock(&_Hyperlane.CallOpts)
}

// LatestDispatchedId is a free data retrieval call binding the contract method 0x134fbb4f.
//
// Solidity: function latestDispatchedId() view returns(bytes32)
func (_Hyperlane *HyperlaneCaller) LatestDispatchedId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "latestDispatchedId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestDispatchedId is a free data retrieval call binding the contract method 0x134fbb4f.
//
// Solidity: function latestDispatchedId() view returns(bytes32)
func (_Hyperlane *HyperlaneSession) LatestDispatchedId() ([32]byte, error) {
	return _Hyperlane.Contract.LatestDispatchedId(&_Hyperlane.CallOpts)
}

// LatestDispatchedId is a free data retrieval call binding the contract method 0x134fbb4f.
//
// Solidity: function latestDispatchedId() view returns(bytes32)
func (_Hyperlane *HyperlaneCallerSession) LatestDispatchedId() ([32]byte, error) {
	return _Hyperlane.Contract.LatestDispatchedId(&_Hyperlane.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Hyperlane *HyperlaneCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Hyperlane *HyperlaneSession) LocalDomain() (uint32, error) {
	return _Hyperlane.Contract.LocalDomain(&_Hyperlane.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Hyperlane *HyperlaneCallerSession) LocalDomain() (uint32, error) {
	return _Hyperlane.Contract.LocalDomain(&_Hyperlane.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_Hyperlane *HyperlaneCaller) Nonce(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_Hyperlane *HyperlaneSession) Nonce() (uint32, error) {
	return _Hyperlane.Contract.Nonce(&_Hyperlane.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_Hyperlane *HyperlaneCallerSession) Nonce() (uint32, error) {
	return _Hyperlane.Contract.Nonce(&_Hyperlane.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hyperlane *HyperlaneCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hyperlane *HyperlaneSession) Owner() (common.Address, error) {
	return _Hyperlane.Contract.Owner(&_Hyperlane.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hyperlane *HyperlaneCallerSession) Owner() (common.Address, error) {
	return _Hyperlane.Contract.Owner(&_Hyperlane.CallOpts)
}

// ProcessedAt is a free data retrieval call binding the contract method 0x07a2fda1.
//
// Solidity: function processedAt(bytes32 _id) view returns(uint48)
func (_Hyperlane *HyperlaneCaller) ProcessedAt(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "processedAt", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProcessedAt is a free data retrieval call binding the contract method 0x07a2fda1.
//
// Solidity: function processedAt(bytes32 _id) view returns(uint48)
func (_Hyperlane *HyperlaneSession) ProcessedAt(_id [32]byte) (*big.Int, error) {
	return _Hyperlane.Contract.ProcessedAt(&_Hyperlane.CallOpts, _id)
}

// ProcessedAt is a free data retrieval call binding the contract method 0x07a2fda1.
//
// Solidity: function processedAt(bytes32 _id) view returns(uint48)
func (_Hyperlane *HyperlaneCallerSession) ProcessedAt(_id [32]byte) (*big.Int, error) {
	return _Hyperlane.Contract.ProcessedAt(&_Hyperlane.CallOpts, _id)
}

// Processor is a free data retrieval call binding the contract method 0x5d1fe5a9.
//
// Solidity: function processor(bytes32 _id) view returns(address)
func (_Hyperlane *HyperlaneCaller) Processor(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "processor", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Processor is a free data retrieval call binding the contract method 0x5d1fe5a9.
//
// Solidity: function processor(bytes32 _id) view returns(address)
func (_Hyperlane *HyperlaneSession) Processor(_id [32]byte) (common.Address, error) {
	return _Hyperlane.Contract.Processor(&_Hyperlane.CallOpts, _id)
}

// Processor is a free data retrieval call binding the contract method 0x5d1fe5a9.
//
// Solidity: function processor(bytes32 _id) view returns(address)
func (_Hyperlane *HyperlaneCallerSession) Processor(_id [32]byte) (common.Address, error) {
	return _Hyperlane.Contract.Processor(&_Hyperlane.CallOpts, _id)
}

// QuoteDispatch is a free data retrieval call binding the contract method 0x81d2ea95.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes metadata, address hook) view returns(uint256 fee)
func (_Hyperlane *HyperlaneCaller) QuoteDispatch(opts *bind.CallOpts, destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, metadata []byte, hook common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "quoteDispatch", destinationDomain, recipientAddress, messageBody, metadata, hook)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteDispatch is a free data retrieval call binding the contract method 0x81d2ea95.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes metadata, address hook) view returns(uint256 fee)
func (_Hyperlane *HyperlaneSession) QuoteDispatch(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, metadata []byte, hook common.Address) (*big.Int, error) {
	return _Hyperlane.Contract.QuoteDispatch(&_Hyperlane.CallOpts, destinationDomain, recipientAddress, messageBody, metadata, hook)
}

// QuoteDispatch is a free data retrieval call binding the contract method 0x81d2ea95.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes metadata, address hook) view returns(uint256 fee)
func (_Hyperlane *HyperlaneCallerSession) QuoteDispatch(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, metadata []byte, hook common.Address) (*big.Int, error) {
	return _Hyperlane.Contract.QuoteDispatch(&_Hyperlane.CallOpts, destinationDomain, recipientAddress, messageBody, metadata, hook)
}

// QuoteDispatch0 is a free data retrieval call binding the contract method 0x9c42bd18.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody) view returns(uint256 fee)
func (_Hyperlane *HyperlaneCaller) QuoteDispatch0(opts *bind.CallOpts, destinationDomain uint32, recipientAddress [32]byte, messageBody []byte) (*big.Int, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "quoteDispatch0", destinationDomain, recipientAddress, messageBody)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteDispatch0 is a free data retrieval call binding the contract method 0x9c42bd18.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody) view returns(uint256 fee)
func (_Hyperlane *HyperlaneSession) QuoteDispatch0(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte) (*big.Int, error) {
	return _Hyperlane.Contract.QuoteDispatch0(&_Hyperlane.CallOpts, destinationDomain, recipientAddress, messageBody)
}

// QuoteDispatch0 is a free data retrieval call binding the contract method 0x9c42bd18.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody) view returns(uint256 fee)
func (_Hyperlane *HyperlaneCallerSession) QuoteDispatch0(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte) (*big.Int, error) {
	return _Hyperlane.Contract.QuoteDispatch0(&_Hyperlane.CallOpts, destinationDomain, recipientAddress, messageBody)
}

// QuoteDispatch1 is a free data retrieval call binding the contract method 0xf7ccd321.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes defaultHookMetadata) view returns(uint256 fee)
func (_Hyperlane *HyperlaneCaller) QuoteDispatch1(opts *bind.CallOpts, destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, defaultHookMetadata []byte) (*big.Int, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "quoteDispatch1", destinationDomain, recipientAddress, messageBody, defaultHookMetadata)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteDispatch1 is a free data retrieval call binding the contract method 0xf7ccd321.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes defaultHookMetadata) view returns(uint256 fee)
func (_Hyperlane *HyperlaneSession) QuoteDispatch1(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, defaultHookMetadata []byte) (*big.Int, error) {
	return _Hyperlane.Contract.QuoteDispatch1(&_Hyperlane.CallOpts, destinationDomain, recipientAddress, messageBody, defaultHookMetadata)
}

// QuoteDispatch1 is a free data retrieval call binding the contract method 0xf7ccd321.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes defaultHookMetadata) view returns(uint256 fee)
func (_Hyperlane *HyperlaneCallerSession) QuoteDispatch1(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, defaultHookMetadata []byte) (*big.Int, error) {
	return _Hyperlane.Contract.QuoteDispatch1(&_Hyperlane.CallOpts, destinationDomain, recipientAddress, messageBody, defaultHookMetadata)
}

// RecipientIsm is a free data retrieval call binding the contract method 0xe70f48ac.
//
// Solidity: function recipientIsm(address _recipient) view returns(address)
func (_Hyperlane *HyperlaneCaller) RecipientIsm(opts *bind.CallOpts, _recipient common.Address) (common.Address, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "recipientIsm", _recipient)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecipientIsm is a free data retrieval call binding the contract method 0xe70f48ac.
//
// Solidity: function recipientIsm(address _recipient) view returns(address)
func (_Hyperlane *HyperlaneSession) RecipientIsm(_recipient common.Address) (common.Address, error) {
	return _Hyperlane.Contract.RecipientIsm(&_Hyperlane.CallOpts, _recipient)
}

// RecipientIsm is a free data retrieval call binding the contract method 0xe70f48ac.
//
// Solidity: function recipientIsm(address _recipient) view returns(address)
func (_Hyperlane *HyperlaneCallerSession) RecipientIsm(_recipient common.Address) (common.Address, error) {
	return _Hyperlane.Contract.RecipientIsm(&_Hyperlane.CallOpts, _recipient)
}

// RequiredHook is a free data retrieval call binding the contract method 0xd6d08a09.
//
// Solidity: function requiredHook() view returns(address)
func (_Hyperlane *HyperlaneCaller) RequiredHook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hyperlane.contract.Call(opts, &out, "requiredHook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RequiredHook is a free data retrieval call binding the contract method 0xd6d08a09.
//
// Solidity: function requiredHook() view returns(address)
func (_Hyperlane *HyperlaneSession) RequiredHook() (common.Address, error) {
	return _Hyperlane.Contract.RequiredHook(&_Hyperlane.CallOpts)
}

// RequiredHook is a free data retrieval call binding the contract method 0xd6d08a09.
//
// Solidity: function requiredHook() view returns(address)
func (_Hyperlane *HyperlaneCallerSession) RequiredHook() (common.Address, error) {
	return _Hyperlane.Contract.RequiredHook(&_Hyperlane.CallOpts)
}

// Dispatch is a paid mutator transaction binding the contract method 0x10b83dc0.
//
// Solidity: function dispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes metadata, address hook) payable returns(bytes32)
func (_Hyperlane *HyperlaneTransactor) Dispatch(opts *bind.TransactOpts, destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, metadata []byte, hook common.Address) (*types.Transaction, error) {
	return _Hyperlane.contract.Transact(opts, "dispatch", destinationDomain, recipientAddress, messageBody, metadata, hook)
}

// Dispatch is a paid mutator transaction binding the contract method 0x10b83dc0.
//
// Solidity: function dispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes metadata, address hook) payable returns(bytes32)
func (_Hyperlane *HyperlaneSession) Dispatch(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, metadata []byte, hook common.Address) (*types.Transaction, error) {
	return _Hyperlane.Contract.Dispatch(&_Hyperlane.TransactOpts, destinationDomain, recipientAddress, messageBody, metadata, hook)
}

// Dispatch is a paid mutator transaction binding the contract method 0x10b83dc0.
//
// Solidity: function dispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes metadata, address hook) payable returns(bytes32)
func (_Hyperlane *HyperlaneTransactorSession) Dispatch(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, metadata []byte, hook common.Address) (*types.Transaction, error) {
	return _Hyperlane.Contract.Dispatch(&_Hyperlane.TransactOpts, destinationDomain, recipientAddress, messageBody, metadata, hook)
}

// Dispatch0 is a paid mutator transaction binding the contract method 0x48aee8d4.
//
// Solidity: function dispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes hookMetadata) payable returns(bytes32)
func (_Hyperlane *HyperlaneTransactor) Dispatch0(opts *bind.TransactOpts, destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, hookMetadata []byte) (*types.Transaction, error) {
	return _Hyperlane.contract.Transact(opts, "dispatch0", destinationDomain, recipientAddress, messageBody, hookMetadata)
}

// Dispatch0 is a paid mutator transaction binding the contract method 0x48aee8d4.
//
// Solidity: function dispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes hookMetadata) payable returns(bytes32)
func (_Hyperlane *HyperlaneSession) Dispatch0(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, hookMetadata []byte) (*types.Transaction, error) {
	return _Hyperlane.Contract.Dispatch0(&_Hyperlane.TransactOpts, destinationDomain, recipientAddress, messageBody, hookMetadata)
}

// Dispatch0 is a paid mutator transaction binding the contract method 0x48aee8d4.
//
// Solidity: function dispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes hookMetadata) payable returns(bytes32)
func (_Hyperlane *HyperlaneTransactorSession) Dispatch0(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, hookMetadata []byte) (*types.Transaction, error) {
	return _Hyperlane.Contract.Dispatch0(&_Hyperlane.TransactOpts, destinationDomain, recipientAddress, messageBody, hookMetadata)
}

// Dispatch1 is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) payable returns(bytes32)
func (_Hyperlane *HyperlaneTransactor) Dispatch1(opts *bind.TransactOpts, _destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Hyperlane.contract.Transact(opts, "dispatch1", _destinationDomain, _recipientAddress, _messageBody)
}

// Dispatch1 is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) payable returns(bytes32)
func (_Hyperlane *HyperlaneSession) Dispatch1(_destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Hyperlane.Contract.Dispatch1(&_Hyperlane.TransactOpts, _destinationDomain, _recipientAddress, _messageBody)
}

// Dispatch1 is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) payable returns(bytes32)
func (_Hyperlane *HyperlaneTransactorSession) Dispatch1(_destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Hyperlane.Contract.Dispatch1(&_Hyperlane.TransactOpts, _destinationDomain, _recipientAddress, _messageBody)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _owner, address _defaultIsm, address _defaultHook, address _requiredHook) returns()
func (_Hyperlane *HyperlaneTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _defaultIsm common.Address, _defaultHook common.Address, _requiredHook common.Address) (*types.Transaction, error) {
	return _Hyperlane.contract.Transact(opts, "initialize", _owner, _defaultIsm, _defaultHook, _requiredHook)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _owner, address _defaultIsm, address _defaultHook, address _requiredHook) returns()
func (_Hyperlane *HyperlaneSession) Initialize(_owner common.Address, _defaultIsm common.Address, _defaultHook common.Address, _requiredHook common.Address) (*types.Transaction, error) {
	return _Hyperlane.Contract.Initialize(&_Hyperlane.TransactOpts, _owner, _defaultIsm, _defaultHook, _requiredHook)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _owner, address _defaultIsm, address _defaultHook, address _requiredHook) returns()
func (_Hyperlane *HyperlaneTransactorSession) Initialize(_owner common.Address, _defaultIsm common.Address, _defaultHook common.Address, _requiredHook common.Address) (*types.Transaction, error) {
	return _Hyperlane.Contract.Initialize(&_Hyperlane.TransactOpts, _owner, _defaultIsm, _defaultHook, _requiredHook)
}

// Process is a paid mutator transaction binding the contract method 0x7c39d130.
//
// Solidity: function process(bytes _metadata, bytes _message) payable returns()
func (_Hyperlane *HyperlaneTransactor) Process(opts *bind.TransactOpts, _metadata []byte, _message []byte) (*types.Transaction, error) {
	return _Hyperlane.contract.Transact(opts, "process", _metadata, _message)
}

// Process is a paid mutator transaction binding the contract method 0x7c39d130.
//
// Solidity: function process(bytes _metadata, bytes _message) payable returns()
func (_Hyperlane *HyperlaneSession) Process(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _Hyperlane.Contract.Process(&_Hyperlane.TransactOpts, _metadata, _message)
}

// Process is a paid mutator transaction binding the contract method 0x7c39d130.
//
// Solidity: function process(bytes _metadata, bytes _message) payable returns()
func (_Hyperlane *HyperlaneTransactorSession) Process(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _Hyperlane.Contract.Process(&_Hyperlane.TransactOpts, _metadata, _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hyperlane *HyperlaneTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyperlane.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hyperlane *HyperlaneSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hyperlane.Contract.RenounceOwnership(&_Hyperlane.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hyperlane *HyperlaneTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hyperlane.Contract.RenounceOwnership(&_Hyperlane.TransactOpts)
}

// SetDefaultHook is a paid mutator transaction binding the contract method 0x99b04809.
//
// Solidity: function setDefaultHook(address _hook) returns()
func (_Hyperlane *HyperlaneTransactor) SetDefaultHook(opts *bind.TransactOpts, _hook common.Address) (*types.Transaction, error) {
	return _Hyperlane.contract.Transact(opts, "setDefaultHook", _hook)
}

// SetDefaultHook is a paid mutator transaction binding the contract method 0x99b04809.
//
// Solidity: function setDefaultHook(address _hook) returns()
func (_Hyperlane *HyperlaneSession) SetDefaultHook(_hook common.Address) (*types.Transaction, error) {
	return _Hyperlane.Contract.SetDefaultHook(&_Hyperlane.TransactOpts, _hook)
}

// SetDefaultHook is a paid mutator transaction binding the contract method 0x99b04809.
//
// Solidity: function setDefaultHook(address _hook) returns()
func (_Hyperlane *HyperlaneTransactorSession) SetDefaultHook(_hook common.Address) (*types.Transaction, error) {
	return _Hyperlane.Contract.SetDefaultHook(&_Hyperlane.TransactOpts, _hook)
}

// SetDefaultIsm is a paid mutator transaction binding the contract method 0xf794687a.
//
// Solidity: function setDefaultIsm(address _module) returns()
func (_Hyperlane *HyperlaneTransactor) SetDefaultIsm(opts *bind.TransactOpts, _module common.Address) (*types.Transaction, error) {
	return _Hyperlane.contract.Transact(opts, "setDefaultIsm", _module)
}

// SetDefaultIsm is a paid mutator transaction binding the contract method 0xf794687a.
//
// Solidity: function setDefaultIsm(address _module) returns()
func (_Hyperlane *HyperlaneSession) SetDefaultIsm(_module common.Address) (*types.Transaction, error) {
	return _Hyperlane.Contract.SetDefaultIsm(&_Hyperlane.TransactOpts, _module)
}

// SetDefaultIsm is a paid mutator transaction binding the contract method 0xf794687a.
//
// Solidity: function setDefaultIsm(address _module) returns()
func (_Hyperlane *HyperlaneTransactorSession) SetDefaultIsm(_module common.Address) (*types.Transaction, error) {
	return _Hyperlane.Contract.SetDefaultIsm(&_Hyperlane.TransactOpts, _module)
}

// SetRequiredHook is a paid mutator transaction binding the contract method 0x1426b7f4.
//
// Solidity: function setRequiredHook(address _hook) returns()
func (_Hyperlane *HyperlaneTransactor) SetRequiredHook(opts *bind.TransactOpts, _hook common.Address) (*types.Transaction, error) {
	return _Hyperlane.contract.Transact(opts, "setRequiredHook", _hook)
}

// SetRequiredHook is a paid mutator transaction binding the contract method 0x1426b7f4.
//
// Solidity: function setRequiredHook(address _hook) returns()
func (_Hyperlane *HyperlaneSession) SetRequiredHook(_hook common.Address) (*types.Transaction, error) {
	return _Hyperlane.Contract.SetRequiredHook(&_Hyperlane.TransactOpts, _hook)
}

// SetRequiredHook is a paid mutator transaction binding the contract method 0x1426b7f4.
//
// Solidity: function setRequiredHook(address _hook) returns()
func (_Hyperlane *HyperlaneTransactorSession) SetRequiredHook(_hook common.Address) (*types.Transaction, error) {
	return _Hyperlane.Contract.SetRequiredHook(&_Hyperlane.TransactOpts, _hook)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hyperlane *HyperlaneTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Hyperlane.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hyperlane *HyperlaneSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hyperlane.Contract.TransferOwnership(&_Hyperlane.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hyperlane *HyperlaneTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hyperlane.Contract.TransferOwnership(&_Hyperlane.TransactOpts, newOwner)
}

// HyperlaneDefaultHookSetIterator is returned from FilterDefaultHookSet and is used to iterate over the raw logs and unpacked data for DefaultHookSet events raised by the Hyperlane contract.
type HyperlaneDefaultHookSetIterator struct {
	Event *HyperlaneDefaultHookSet // Event containing the contract specifics and raw log

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
func (it *HyperlaneDefaultHookSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperlaneDefaultHookSet)
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
		it.Event = new(HyperlaneDefaultHookSet)
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
func (it *HyperlaneDefaultHookSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperlaneDefaultHookSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperlaneDefaultHookSet represents a DefaultHookSet event raised by the Hyperlane contract.
type HyperlaneDefaultHookSet struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDefaultHookSet is a free log retrieval operation binding the contract event 0x65a63e5066ee2fcdf9d32a7f1bf7ce71c76066f19d0609dddccd334ab87237d7.
//
// Solidity: event DefaultHookSet(address indexed hook)
func (_Hyperlane *HyperlaneFilterer) FilterDefaultHookSet(opts *bind.FilterOpts, hook []common.Address) (*HyperlaneDefaultHookSetIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _Hyperlane.contract.FilterLogs(opts, "DefaultHookSet", hookRule)
	if err != nil {
		return nil, err
	}
	return &HyperlaneDefaultHookSetIterator{contract: _Hyperlane.contract, event: "DefaultHookSet", logs: logs, sub: sub}, nil
}

// WatchDefaultHookSet is a free log subscription operation binding the contract event 0x65a63e5066ee2fcdf9d32a7f1bf7ce71c76066f19d0609dddccd334ab87237d7.
//
// Solidity: event DefaultHookSet(address indexed hook)
func (_Hyperlane *HyperlaneFilterer) WatchDefaultHookSet(opts *bind.WatchOpts, sink chan<- *HyperlaneDefaultHookSet, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _Hyperlane.contract.WatchLogs(opts, "DefaultHookSet", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperlaneDefaultHookSet)
				if err := _Hyperlane.contract.UnpackLog(event, "DefaultHookSet", log); err != nil {
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

// ParseDefaultHookSet is a log parse operation binding the contract event 0x65a63e5066ee2fcdf9d32a7f1bf7ce71c76066f19d0609dddccd334ab87237d7.
//
// Solidity: event DefaultHookSet(address indexed hook)
func (_Hyperlane *HyperlaneFilterer) ParseDefaultHookSet(log types.Log) (*HyperlaneDefaultHookSet, error) {
	event := new(HyperlaneDefaultHookSet)
	if err := _Hyperlane.contract.UnpackLog(event, "DefaultHookSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperlaneDefaultIsmSetIterator is returned from FilterDefaultIsmSet and is used to iterate over the raw logs and unpacked data for DefaultIsmSet events raised by the Hyperlane contract.
type HyperlaneDefaultIsmSetIterator struct {
	Event *HyperlaneDefaultIsmSet // Event containing the contract specifics and raw log

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
func (it *HyperlaneDefaultIsmSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperlaneDefaultIsmSet)
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
		it.Event = new(HyperlaneDefaultIsmSet)
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
func (it *HyperlaneDefaultIsmSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperlaneDefaultIsmSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperlaneDefaultIsmSet represents a DefaultIsmSet event raised by the Hyperlane contract.
type HyperlaneDefaultIsmSet struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultIsmSet is a free log retrieval operation binding the contract event 0xa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa.
//
// Solidity: event DefaultIsmSet(address indexed module)
func (_Hyperlane *HyperlaneFilterer) FilterDefaultIsmSet(opts *bind.FilterOpts, module []common.Address) (*HyperlaneDefaultIsmSetIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Hyperlane.contract.FilterLogs(opts, "DefaultIsmSet", moduleRule)
	if err != nil {
		return nil, err
	}
	return &HyperlaneDefaultIsmSetIterator{contract: _Hyperlane.contract, event: "DefaultIsmSet", logs: logs, sub: sub}, nil
}

// WatchDefaultIsmSet is a free log subscription operation binding the contract event 0xa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa.
//
// Solidity: event DefaultIsmSet(address indexed module)
func (_Hyperlane *HyperlaneFilterer) WatchDefaultIsmSet(opts *bind.WatchOpts, sink chan<- *HyperlaneDefaultIsmSet, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Hyperlane.contract.WatchLogs(opts, "DefaultIsmSet", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperlaneDefaultIsmSet)
				if err := _Hyperlane.contract.UnpackLog(event, "DefaultIsmSet", log); err != nil {
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

// ParseDefaultIsmSet is a log parse operation binding the contract event 0xa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa.
//
// Solidity: event DefaultIsmSet(address indexed module)
func (_Hyperlane *HyperlaneFilterer) ParseDefaultIsmSet(log types.Log) (*HyperlaneDefaultIsmSet, error) {
	event := new(HyperlaneDefaultIsmSet)
	if err := _Hyperlane.contract.UnpackLog(event, "DefaultIsmSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperlaneDispatchIterator is returned from FilterDispatch and is used to iterate over the raw logs and unpacked data for Dispatch events raised by the Hyperlane contract.
type HyperlaneDispatchIterator struct {
	Event *HyperlaneDispatch // Event containing the contract specifics and raw log

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
func (it *HyperlaneDispatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperlaneDispatch)
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
		it.Event = new(HyperlaneDispatch)
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
func (it *HyperlaneDispatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperlaneDispatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperlaneDispatch represents a Dispatch event raised by the Hyperlane contract.
type HyperlaneDispatch struct {
	Sender      common.Address
	Destination uint32
	Recipient   [32]byte
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDispatch is a free log retrieval operation binding the contract event 0x769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c814.
//
// Solidity: event Dispatch(address indexed sender, uint32 indexed destination, bytes32 indexed recipient, bytes message)
func (_Hyperlane *HyperlaneFilterer) FilterDispatch(opts *bind.FilterOpts, sender []common.Address, destination []uint32, recipient [][32]byte) (*HyperlaneDispatchIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Hyperlane.contract.FilterLogs(opts, "Dispatch", senderRule, destinationRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &HyperlaneDispatchIterator{contract: _Hyperlane.contract, event: "Dispatch", logs: logs, sub: sub}, nil
}

// WatchDispatch is a free log subscription operation binding the contract event 0x769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c814.
//
// Solidity: event Dispatch(address indexed sender, uint32 indexed destination, bytes32 indexed recipient, bytes message)
func (_Hyperlane *HyperlaneFilterer) WatchDispatch(opts *bind.WatchOpts, sink chan<- *HyperlaneDispatch, sender []common.Address, destination []uint32, recipient [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Hyperlane.contract.WatchLogs(opts, "Dispatch", senderRule, destinationRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperlaneDispatch)
				if err := _Hyperlane.contract.UnpackLog(event, "Dispatch", log); err != nil {
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

// ParseDispatch is a log parse operation binding the contract event 0x769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c814.
//
// Solidity: event Dispatch(address indexed sender, uint32 indexed destination, bytes32 indexed recipient, bytes message)
func (_Hyperlane *HyperlaneFilterer) ParseDispatch(log types.Log) (*HyperlaneDispatch, error) {
	event := new(HyperlaneDispatch)
	if err := _Hyperlane.contract.UnpackLog(event, "Dispatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperlaneDispatchIdIterator is returned from FilterDispatchId and is used to iterate over the raw logs and unpacked data for DispatchId events raised by the Hyperlane contract.
type HyperlaneDispatchIdIterator struct {
	Event *HyperlaneDispatchId // Event containing the contract specifics and raw log

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
func (it *HyperlaneDispatchIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperlaneDispatchId)
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
		it.Event = new(HyperlaneDispatchId)
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
func (it *HyperlaneDispatchIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperlaneDispatchIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperlaneDispatchId represents a DispatchId event raised by the Hyperlane contract.
type HyperlaneDispatchId struct {
	MessageId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDispatchId is a free log retrieval operation binding the contract event 0x788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a.
//
// Solidity: event DispatchId(bytes32 indexed messageId)
func (_Hyperlane *HyperlaneFilterer) FilterDispatchId(opts *bind.FilterOpts, messageId [][32]byte) (*HyperlaneDispatchIdIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Hyperlane.contract.FilterLogs(opts, "DispatchId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &HyperlaneDispatchIdIterator{contract: _Hyperlane.contract, event: "DispatchId", logs: logs, sub: sub}, nil
}

// WatchDispatchId is a free log subscription operation binding the contract event 0x788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a.
//
// Solidity: event DispatchId(bytes32 indexed messageId)
func (_Hyperlane *HyperlaneFilterer) WatchDispatchId(opts *bind.WatchOpts, sink chan<- *HyperlaneDispatchId, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Hyperlane.contract.WatchLogs(opts, "DispatchId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperlaneDispatchId)
				if err := _Hyperlane.contract.UnpackLog(event, "DispatchId", log); err != nil {
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

// ParseDispatchId is a log parse operation binding the contract event 0x788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a.
//
// Solidity: event DispatchId(bytes32 indexed messageId)
func (_Hyperlane *HyperlaneFilterer) ParseDispatchId(log types.Log) (*HyperlaneDispatchId, error) {
	event := new(HyperlaneDispatchId)
	if err := _Hyperlane.contract.UnpackLog(event, "DispatchId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperlaneInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Hyperlane contract.
type HyperlaneInitializedIterator struct {
	Event *HyperlaneInitialized // Event containing the contract specifics and raw log

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
func (it *HyperlaneInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperlaneInitialized)
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
		it.Event = new(HyperlaneInitialized)
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
func (it *HyperlaneInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperlaneInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperlaneInitialized represents a Initialized event raised by the Hyperlane contract.
type HyperlaneInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Hyperlane *HyperlaneFilterer) FilterInitialized(opts *bind.FilterOpts) (*HyperlaneInitializedIterator, error) {

	logs, sub, err := _Hyperlane.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HyperlaneInitializedIterator{contract: _Hyperlane.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Hyperlane *HyperlaneFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HyperlaneInitialized) (event.Subscription, error) {

	logs, sub, err := _Hyperlane.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperlaneInitialized)
				if err := _Hyperlane.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Hyperlane *HyperlaneFilterer) ParseInitialized(log types.Log) (*HyperlaneInitialized, error) {
	event := new(HyperlaneInitialized)
	if err := _Hyperlane.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperlaneOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Hyperlane contract.
type HyperlaneOwnershipTransferredIterator struct {
	Event *HyperlaneOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HyperlaneOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperlaneOwnershipTransferred)
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
		it.Event = new(HyperlaneOwnershipTransferred)
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
func (it *HyperlaneOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperlaneOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperlaneOwnershipTransferred represents a OwnershipTransferred event raised by the Hyperlane contract.
type HyperlaneOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hyperlane *HyperlaneFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HyperlaneOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hyperlane.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HyperlaneOwnershipTransferredIterator{contract: _Hyperlane.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hyperlane *HyperlaneFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HyperlaneOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hyperlane.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperlaneOwnershipTransferred)
				if err := _Hyperlane.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Hyperlane *HyperlaneFilterer) ParseOwnershipTransferred(log types.Log) (*HyperlaneOwnershipTransferred, error) {
	event := new(HyperlaneOwnershipTransferred)
	if err := _Hyperlane.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperlaneProcessIterator is returned from FilterProcess and is used to iterate over the raw logs and unpacked data for Process events raised by the Hyperlane contract.
type HyperlaneProcessIterator struct {
	Event *HyperlaneProcess // Event containing the contract specifics and raw log

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
func (it *HyperlaneProcessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperlaneProcess)
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
		it.Event = new(HyperlaneProcess)
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
func (it *HyperlaneProcessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperlaneProcessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperlaneProcess represents a Process event raised by the Hyperlane contract.
type HyperlaneProcess struct {
	Origin    uint32
	Sender    [32]byte
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcess is a free log retrieval operation binding the contract event 0x0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca6574.
//
// Solidity: event Process(uint32 indexed origin, bytes32 indexed sender, address indexed recipient)
func (_Hyperlane *HyperlaneFilterer) FilterProcess(opts *bind.FilterOpts, origin []uint32, sender [][32]byte, recipient []common.Address) (*HyperlaneProcessIterator, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Hyperlane.contract.FilterLogs(opts, "Process", originRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &HyperlaneProcessIterator{contract: _Hyperlane.contract, event: "Process", logs: logs, sub: sub}, nil
}

// WatchProcess is a free log subscription operation binding the contract event 0x0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca6574.
//
// Solidity: event Process(uint32 indexed origin, bytes32 indexed sender, address indexed recipient)
func (_Hyperlane *HyperlaneFilterer) WatchProcess(opts *bind.WatchOpts, sink chan<- *HyperlaneProcess, origin []uint32, sender [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Hyperlane.contract.WatchLogs(opts, "Process", originRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperlaneProcess)
				if err := _Hyperlane.contract.UnpackLog(event, "Process", log); err != nil {
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

// ParseProcess is a log parse operation binding the contract event 0x0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca6574.
//
// Solidity: event Process(uint32 indexed origin, bytes32 indexed sender, address indexed recipient)
func (_Hyperlane *HyperlaneFilterer) ParseProcess(log types.Log) (*HyperlaneProcess, error) {
	event := new(HyperlaneProcess)
	if err := _Hyperlane.contract.UnpackLog(event, "Process", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperlaneProcessIdIterator is returned from FilterProcessId and is used to iterate over the raw logs and unpacked data for ProcessId events raised by the Hyperlane contract.
type HyperlaneProcessIdIterator struct {
	Event *HyperlaneProcessId // Event containing the contract specifics and raw log

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
func (it *HyperlaneProcessIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperlaneProcessId)
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
		it.Event = new(HyperlaneProcessId)
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
func (it *HyperlaneProcessIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperlaneProcessIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperlaneProcessId represents a ProcessId event raised by the Hyperlane contract.
type HyperlaneProcessId struct {
	MessageId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessId is a free log retrieval operation binding the contract event 0x1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a9.
//
// Solidity: event ProcessId(bytes32 indexed messageId)
func (_Hyperlane *HyperlaneFilterer) FilterProcessId(opts *bind.FilterOpts, messageId [][32]byte) (*HyperlaneProcessIdIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Hyperlane.contract.FilterLogs(opts, "ProcessId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &HyperlaneProcessIdIterator{contract: _Hyperlane.contract, event: "ProcessId", logs: logs, sub: sub}, nil
}

// WatchProcessId is a free log subscription operation binding the contract event 0x1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a9.
//
// Solidity: event ProcessId(bytes32 indexed messageId)
func (_Hyperlane *HyperlaneFilterer) WatchProcessId(opts *bind.WatchOpts, sink chan<- *HyperlaneProcessId, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Hyperlane.contract.WatchLogs(opts, "ProcessId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperlaneProcessId)
				if err := _Hyperlane.contract.UnpackLog(event, "ProcessId", log); err != nil {
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

// ParseProcessId is a log parse operation binding the contract event 0x1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a9.
//
// Solidity: event ProcessId(bytes32 indexed messageId)
func (_Hyperlane *HyperlaneFilterer) ParseProcessId(log types.Log) (*HyperlaneProcessId, error) {
	event := new(HyperlaneProcessId)
	if err := _Hyperlane.contract.UnpackLog(event, "ProcessId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperlaneRequiredHookSetIterator is returned from FilterRequiredHookSet and is used to iterate over the raw logs and unpacked data for RequiredHookSet events raised by the Hyperlane contract.
type HyperlaneRequiredHookSetIterator struct {
	Event *HyperlaneRequiredHookSet // Event containing the contract specifics and raw log

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
func (it *HyperlaneRequiredHookSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperlaneRequiredHookSet)
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
		it.Event = new(HyperlaneRequiredHookSet)
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
func (it *HyperlaneRequiredHookSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperlaneRequiredHookSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperlaneRequiredHookSet represents a RequiredHookSet event raised by the Hyperlane contract.
type HyperlaneRequiredHookSet struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRequiredHookSet is a free log retrieval operation binding the contract event 0x329ec8e2438a73828ecf31a6568d7a91d7b1d79e342b0692914fd053d1a002b1.
//
// Solidity: event RequiredHookSet(address indexed hook)
func (_Hyperlane *HyperlaneFilterer) FilterRequiredHookSet(opts *bind.FilterOpts, hook []common.Address) (*HyperlaneRequiredHookSetIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _Hyperlane.contract.FilterLogs(opts, "RequiredHookSet", hookRule)
	if err != nil {
		return nil, err
	}
	return &HyperlaneRequiredHookSetIterator{contract: _Hyperlane.contract, event: "RequiredHookSet", logs: logs, sub: sub}, nil
}

// WatchRequiredHookSet is a free log subscription operation binding the contract event 0x329ec8e2438a73828ecf31a6568d7a91d7b1d79e342b0692914fd053d1a002b1.
//
// Solidity: event RequiredHookSet(address indexed hook)
func (_Hyperlane *HyperlaneFilterer) WatchRequiredHookSet(opts *bind.WatchOpts, sink chan<- *HyperlaneRequiredHookSet, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _Hyperlane.contract.WatchLogs(opts, "RequiredHookSet", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperlaneRequiredHookSet)
				if err := _Hyperlane.contract.UnpackLog(event, "RequiredHookSet", log); err != nil {
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

// ParseRequiredHookSet is a log parse operation binding the contract event 0x329ec8e2438a73828ecf31a6568d7a91d7b1d79e342b0692914fd053d1a002b1.
//
// Solidity: event RequiredHookSet(address indexed hook)
func (_Hyperlane *HyperlaneFilterer) ParseRequiredHookSet(log types.Log) (*HyperlaneRequiredHookSet, error) {
	event := new(HyperlaneRequiredHookSet)
	if err := _Hyperlane.contract.UnpackLog(event, "RequiredHookSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
