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

// Myerc20tokenMetaData contains all meta data concerning the Myerc20token contract.
var Myerc20tokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]",
}

// Myerc20tokenABI is the input ABI used to generate the binding from.
// Deprecated: Use Myerc20tokenMetaData.ABI instead.
var Myerc20tokenABI = Myerc20tokenMetaData.ABI

// Myerc20token is an auto generated Go binding around an Ethereum contract.
type Myerc20token struct {
	Myerc20tokenCaller     // Read-only binding to the contract
	Myerc20tokenTransactor // Write-only binding to the contract
	Myerc20tokenFilterer   // Log filterer for contract events
}

// Myerc20tokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type Myerc20tokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Myerc20tokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Myerc20tokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Myerc20tokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Myerc20tokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Myerc20tokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Myerc20tokenSession struct {
	Contract     *Myerc20token     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Myerc20tokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Myerc20tokenCallerSession struct {
	Contract *Myerc20tokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Myerc20tokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Myerc20tokenTransactorSession struct {
	Contract     *Myerc20tokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Myerc20tokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type Myerc20tokenRaw struct {
	Contract *Myerc20token // Generic contract binding to access the raw methods on
}

// Myerc20tokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Myerc20tokenCallerRaw struct {
	Contract *Myerc20tokenCaller // Generic read-only contract binding to access the raw methods on
}

// Myerc20tokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Myerc20tokenTransactorRaw struct {
	Contract *Myerc20tokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyerc20token creates a new instance of Myerc20token, bound to a specific deployed contract.
func NewMyerc20token(address common.Address, backend bind.ContractBackend) (*Myerc20token, error) {
	contract, err := bindMyerc20token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Myerc20token{Myerc20tokenCaller: Myerc20tokenCaller{contract: contract}, Myerc20tokenTransactor: Myerc20tokenTransactor{contract: contract}, Myerc20tokenFilterer: Myerc20tokenFilterer{contract: contract}}, nil
}

// NewMyerc20tokenCaller creates a new read-only instance of Myerc20token, bound to a specific deployed contract.
func NewMyerc20tokenCaller(address common.Address, caller bind.ContractCaller) (*Myerc20tokenCaller, error) {
	contract, err := bindMyerc20token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Myerc20tokenCaller{contract: contract}, nil
}

// NewMyerc20tokenTransactor creates a new write-only instance of Myerc20token, bound to a specific deployed contract.
func NewMyerc20tokenTransactor(address common.Address, transactor bind.ContractTransactor) (*Myerc20tokenTransactor, error) {
	contract, err := bindMyerc20token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Myerc20tokenTransactor{contract: contract}, nil
}

// NewMyerc20tokenFilterer creates a new log filterer instance of Myerc20token, bound to a specific deployed contract.
func NewMyerc20tokenFilterer(address common.Address, filterer bind.ContractFilterer) (*Myerc20tokenFilterer, error) {
	contract, err := bindMyerc20token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Myerc20tokenFilterer{contract: contract}, nil
}

// bindMyerc20token binds a generic wrapper to an already deployed contract.
func bindMyerc20token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Myerc20tokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Myerc20token *Myerc20tokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Myerc20token.Contract.Myerc20tokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Myerc20token *Myerc20tokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Myerc20token.Contract.Myerc20tokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Myerc20token *Myerc20tokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Myerc20token.Contract.Myerc20tokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Myerc20token *Myerc20tokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Myerc20token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Myerc20token *Myerc20tokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Myerc20token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Myerc20token *Myerc20tokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Myerc20token.Contract.contract.Transact(opts, method, params...)
}

// Myerc20tokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Myerc20token contract.
type Myerc20tokenApprovalIterator struct {
	Event *Myerc20tokenApproval // Event containing the contract specifics and raw log

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
func (it *Myerc20tokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Myerc20tokenApproval)
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
		it.Event = new(Myerc20tokenApproval)
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
func (it *Myerc20tokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Myerc20tokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Myerc20tokenApproval represents a Approval event raised by the Myerc20token contract.
type Myerc20tokenApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Myerc20token *Myerc20tokenFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*Myerc20tokenApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Myerc20token.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Myerc20tokenApprovalIterator{contract: _Myerc20token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Myerc20token *Myerc20tokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Myerc20tokenApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Myerc20token.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Myerc20tokenApproval)
				if err := _Myerc20token.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Myerc20token *Myerc20tokenFilterer) ParseApproval(log types.Log) (*Myerc20tokenApproval, error) {
	event := new(Myerc20tokenApproval)
	if err := _Myerc20token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Myerc20tokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Myerc20token contract.
type Myerc20tokenTransferIterator struct {
	Event *Myerc20tokenTransfer // Event containing the contract specifics and raw log

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
func (it *Myerc20tokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Myerc20tokenTransfer)
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
		it.Event = new(Myerc20tokenTransfer)
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
func (it *Myerc20tokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Myerc20tokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Myerc20tokenTransfer represents a Transfer event raised by the Myerc20token contract.
type Myerc20tokenTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Myerc20token *Myerc20tokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Myerc20tokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Myerc20token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Myerc20tokenTransferIterator{contract: _Myerc20token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Myerc20token *Myerc20tokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Myerc20tokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Myerc20token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Myerc20tokenTransfer)
				if err := _Myerc20token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Myerc20token *Myerc20tokenFilterer) ParseTransfer(log types.Log) (*Myerc20tokenTransfer, error) {
	event := new(Myerc20tokenTransfer)
	if err := _Myerc20token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
