// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test

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

// StartupProfile is an auto generated low-level Go binding around an user-defined struct.
type StartupProfile struct {
	Name       string
	Mode       uint8
	Logo       string
	Mission    string
	Overview   string
	IsValidate bool
}

// TestMetaData contains all meta data concerning the Test contract.
var TestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"enumStartup.Mode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"logo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mission\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"overview\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isValidate\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structStartup.Profile\",\"name\":\"startUp\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msg\",\"type\":\"address\"}],\"name\":\"created\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"enumStartup.Mode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"logo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mission\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"overview\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isValidate\",\"type\":\"bool\"}],\"internalType\":\"structStartup.Profile\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"newStartup\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"startups\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"enumStartup.Mode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"logo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mission\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"overview\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isValidate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"suicide0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TestABI is the input ABI used to generate the binding from.
// Deprecated: Use TestMetaData.ABI instead.
var TestABI = TestMetaData.ABI

// Test is an auto generated Go binding around an Ethereum contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
	TestFilterer   // Log filterer for contract events
}

// TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// NewTestFilterer creates a new log filterer instance of Test, bound to a specific deployed contract.
func NewTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TestFilterer, error) {
	contract, err := bindTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestFilterer{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.contract.Transact(opts, method, params...)
}

// Startups is a free data retrieval call binding the contract method 0xf86025c3.
//
// Solidity: function startups(string ) view returns(string name, uint8 mode, string logo, string mission, string overview, bool isValidate)
func (_Test *TestCaller) Startups(opts *bind.CallOpts, arg0 string) (struct {
	Name       string
	Mode       uint8
	Logo       string
	Mission    string
	Overview   string
	IsValidate bool
}, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "startups", arg0)

	outstruct := new(struct {
		Name       string
		Mode       uint8
		Logo       string
		Mission    string
		Overview   string
		IsValidate bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Mode = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Logo = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Mission = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Overview = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.IsValidate = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Startups is a free data retrieval call binding the contract method 0xf86025c3.
//
// Solidity: function startups(string ) view returns(string name, uint8 mode, string logo, string mission, string overview, bool isValidate)
func (_Test *TestSession) Startups(arg0 string) (struct {
	Name       string
	Mode       uint8
	Logo       string
	Mission    string
	Overview   string
	IsValidate bool
}, error) {
	return _Test.Contract.Startups(&_Test.CallOpts, arg0)
}

// Startups is a free data retrieval call binding the contract method 0xf86025c3.
//
// Solidity: function startups(string ) view returns(string name, uint8 mode, string logo, string mission, string overview, bool isValidate)
func (_Test *TestCallerSession) Startups(arg0 string) (struct {
	Name       string
	Mode       uint8
	Logo       string
	Mission    string
	Overview   string
	IsValidate bool
}, error) {
	return _Test.Contract.Startups(&_Test.CallOpts, arg0)
}

// NewStartup is a paid mutator transaction binding the contract method 0xbaa143ed.
//
// Solidity: function newStartup((string,uint8,string,string,string,bool) p) payable returns()
func (_Test *TestTransactor) NewStartup(opts *bind.TransactOpts, p StartupProfile) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "newStartup", p)
}

// NewStartup is a paid mutator transaction binding the contract method 0xbaa143ed.
//
// Solidity: function newStartup((string,uint8,string,string,string,bool) p) payable returns()
func (_Test *TestSession) NewStartup(p StartupProfile) (*types.Transaction, error) {
	return _Test.Contract.NewStartup(&_Test.TransactOpts, p)
}

// NewStartup is a paid mutator transaction binding the contract method 0xbaa143ed.
//
// Solidity: function newStartup((string,uint8,string,string,string,bool) p) payable returns()
func (_Test *TestTransactorSession) NewStartup(p StartupProfile) (*types.Transaction, error) {
	return _Test.Contract.NewStartup(&_Test.TransactOpts, p)
}

// Suicide0 is a paid mutator transaction binding the contract method 0x14da447f.
//
// Solidity: function suicide0(address receiver) returns()
func (_Test *TestTransactor) Suicide0(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "suicide0", receiver)
}

// Suicide0 is a paid mutator transaction binding the contract method 0x14da447f.
//
// Solidity: function suicide0(address receiver) returns()
func (_Test *TestSession) Suicide0(receiver common.Address) (*types.Transaction, error) {
	return _Test.Contract.Suicide0(&_Test.TransactOpts, receiver)
}

// Suicide0 is a paid mutator transaction binding the contract method 0x14da447f.
//
// Solidity: function suicide0(address receiver) returns()
func (_Test *TestTransactorSession) Suicide0(receiver common.Address) (*types.Transaction, error) {
	return _Test.Contract.Suicide0(&_Test.TransactOpts, receiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Test *TestTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Test *TestSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Test.Contract.TransferOwnership(&_Test.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Test *TestTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Test.Contract.TransferOwnership(&_Test.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Test *TestTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Test.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Test *TestSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Test.Contract.Fallback(&_Test.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Test *TestTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Test.Contract.Fallback(&_Test.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Test *TestTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Test *TestSession) Receive() (*types.Transaction, error) {
	return _Test.Contract.Receive(&_Test.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Test *TestTransactorSession) Receive() (*types.Transaction, error) {
	return _Test.Contract.Receive(&_Test.TransactOpts)
}

// TestOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Test contract.
type TestOwnershipTransferredIterator struct {
	Event *TestOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOwnershipTransferred)
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
		it.Event = new(TestOwnershipTransferred)
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
func (it *TestOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestOwnershipTransferred represents a OwnershipTransferred event raised by the Test contract.
type TestOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Test *TestFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Test.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestOwnershipTransferredIterator{contract: _Test.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Test *TestFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Test.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestOwnershipTransferred)
				if err := _Test.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Test *TestFilterer) ParseOwnershipTransferred(log types.Log) (*TestOwnershipTransferred, error) {
	event := new(TestOwnershipTransferred)
	if err := _Test.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the Test contract.
type TestCreatedIterator struct {
	Event *TestCreated // Event containing the contract specifics and raw log

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
func (it *TestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestCreated)
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
		it.Event = new(TestCreated)
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
func (it *TestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestCreated represents a Created event raised by the Test contract.
type TestCreated struct {
	Name    string
	StartUp StartupProfile
	Msg     common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x822c16987e5c88fd1ec8ce2935c0b5daf646231496d234745d143a9b62673973.
//
// Solidity: event created(string name, (string,uint8,string,string,string,bool) startUp, address msg)
func (_Test *TestFilterer) FilterCreated(opts *bind.FilterOpts) (*TestCreatedIterator, error) {

	logs, sub, err := _Test.contract.FilterLogs(opts, "created")
	if err != nil {
		return nil, err
	}
	return &TestCreatedIterator{contract: _Test.contract, event: "created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x822c16987e5c88fd1ec8ce2935c0b5daf646231496d234745d143a9b62673973.
//
// Solidity: event created(string name, (string,uint8,string,string,string,bool) startUp, address msg)
func (_Test *TestFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *TestCreated) (event.Subscription, error) {

	logs, sub, err := _Test.contract.WatchLogs(opts, "created")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestCreated)
				if err := _Test.contract.UnpackLog(event, "created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0x822c16987e5c88fd1ec8ce2935c0b5daf646231496d234745d143a9b62673973.
//
// Solidity: event created(string name, (string,uint8,string,string,string,bool) startUp, address msg)
func (_Test *TestFilterer) ParseCreated(log types.Log) (*TestCreated, error) {
	event := new(TestCreated)
	if err := _Test.contract.UnpackLog(event, "created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
