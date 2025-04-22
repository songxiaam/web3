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

// StroeMetaData contains all meta data concerning the Stroe contract.
var StroeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161087838038061087883398181016040528101906100319190610193565b805f908161003f91906103ea565b50506104b9565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6100a58261005f565b810181811067ffffffffffffffff821117156100c4576100c361006f565b5b80604052505050565b5f6100d6610046565b90506100e2828261009c565b919050565b5f67ffffffffffffffff8211156101015761010061006f565b5b61010a8261005f565b9050602081019050919050565b8281835e5f83830152505050565b5f610137610132846100e7565b6100cd565b9050828152602081018484840111156101535761015261005b565b5b61015e848285610117565b509392505050565b5f82601f83011261017a57610179610057565b5b815161018a848260208601610125565b91505092915050565b5f602082840312156101a8576101a761004f565b5b5f82015167ffffffffffffffff8111156101c5576101c4610053565b5b6101d184828501610166565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061022857607f821691505b60208210810361023b5761023a6101e4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261029d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610262565b6102a78683610262565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102eb6102e66102e1846102bf565b6102c8565b6102bf565b9050919050565b5f819050919050565b610304836102d1565b610318610310826102f2565b84845461026e565b825550505050565b5f5f905090565b61032f610320565b61033a8184846102fb565b505050565b5b8181101561035d576103525f82610327565b600181019050610340565b5050565b601f8211156103a25761037381610241565b61037c84610253565b8101602085101561038b578190505b61039f61039785610253565b83018261033f565b50505b505050565b5f82821c905092915050565b5f6103c25f19846008026103a7565b1980831691505092915050565b5f6103da83836103b3565b9150826002028217905092915050565b6103f3826101da565b67ffffffffffffffff81111561040c5761040b61006f565b5b6104168254610211565b610421828285610361565b5f60209050601f831160018114610452575f8415610440578287015190505b61044a85826103cf565b8655506104b1565b601f19841661046086610241565b5f5b8281101561048757848901518255600182019150602085019450602081019050610462565b868310156104a457848901516104a0601f8916826103b3565b8355505b6001600288020188555050505b505050505050565b6103b2806104c65f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c806348f343f31461004357806354fd4d5014610073578063f56256c714610091575b5f5ffd5b61005d600480360381019061005891906101d7565b6100ad565b60405161006a9190610211565b60405180910390f35b61007b6100c2565b604051610088919061029a565b60405180910390f35b6100ab60048036038101906100a691906102ba565b61014d565b005b6001602052805f5260405f205f915090505481565b5f80546100ce90610325565b80601f01602080910402602001604051908101604052809291908181526020018280546100fa90610325565b80156101455780601f1061011c57610100808354040283529160200191610145565b820191905f5260205f20905b81548152906001019060200180831161012857829003601f168201915b505050505081565b8060015f8481526020019081526020015f20819055507fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d48282604051610194929190610355565b60405180910390a15050565b5f5ffd5b5f819050919050565b6101b6816101a4565b81146101c0575f5ffd5b50565b5f813590506101d1816101ad565b92915050565b5f602082840312156101ec576101eb6101a0565b5b5f6101f9848285016101c3565b91505092915050565b61020b816101a4565b82525050565b5f6020820190506102245f830184610202565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61026c8261022a565b6102768185610234565b9350610286818560208601610244565b61028f81610252565b840191505092915050565b5f6020820190508181035f8301526102b28184610262565b905092915050565b5f5f604083850312156102d0576102cf6101a0565b5b5f6102dd858286016101c3565b92505060206102ee858286016101c3565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061033c57607f821691505b60208210810361034f5761034e6102f8565b5b50919050565b5f6040820190506103685f830185610202565b6103756020830184610202565b939250505056fea2646970667358221220c542d1cb3abee80ff0fee491642fa4d1d935276c652e434a8ed33c230e45f71b64736f6c634300081d0033",
}

// StroeABI is the input ABI used to generate the binding from.
// Deprecated: Use StroeMetaData.ABI instead.
var StroeABI = StroeMetaData.ABI

// StroeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StroeMetaData.Bin instead.
var StroeBin = StroeMetaData.Bin

// DeployStroe deploys a new Ethereum contract, binding an instance of Stroe to it.
func DeployStroe(auth *bind.TransactOpts, backend bind.ContractBackend, _version string) (common.Address, *types.Transaction, *Stroe, error) {
	parsed, err := StroeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StroeBin), backend, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stroe{StroeCaller: StroeCaller{contract: contract}, StroeTransactor: StroeTransactor{contract: contract}, StroeFilterer: StroeFilterer{contract: contract}}, nil
}

// Stroe is an auto generated Go binding around an Ethereum contract.
type Stroe struct {
	StroeCaller     // Read-only binding to the contract
	StroeTransactor // Write-only binding to the contract
	StroeFilterer   // Log filterer for contract events
}

// StroeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StroeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StroeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StroeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StroeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StroeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StroeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StroeSession struct {
	Contract     *Stroe            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StroeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StroeCallerSession struct {
	Contract *StroeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StroeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StroeTransactorSession struct {
	Contract     *StroeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StroeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StroeRaw struct {
	Contract *Stroe // Generic contract binding to access the raw methods on
}

// StroeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StroeCallerRaw struct {
	Contract *StroeCaller // Generic read-only contract binding to access the raw methods on
}

// StroeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StroeTransactorRaw struct {
	Contract *StroeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStroe creates a new instance of Stroe, bound to a specific deployed contract.
func NewStroe(address common.Address, backend bind.ContractBackend) (*Stroe, error) {
	contract, err := bindStroe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stroe{StroeCaller: StroeCaller{contract: contract}, StroeTransactor: StroeTransactor{contract: contract}, StroeFilterer: StroeFilterer{contract: contract}}, nil
}

// NewStroeCaller creates a new read-only instance of Stroe, bound to a specific deployed contract.
func NewStroeCaller(address common.Address, caller bind.ContractCaller) (*StroeCaller, error) {
	contract, err := bindStroe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StroeCaller{contract: contract}, nil
}

// NewStroeTransactor creates a new write-only instance of Stroe, bound to a specific deployed contract.
func NewStroeTransactor(address common.Address, transactor bind.ContractTransactor) (*StroeTransactor, error) {
	contract, err := bindStroe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StroeTransactor{contract: contract}, nil
}

// NewStroeFilterer creates a new log filterer instance of Stroe, bound to a specific deployed contract.
func NewStroeFilterer(address common.Address, filterer bind.ContractFilterer) (*StroeFilterer, error) {
	contract, err := bindStroe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StroeFilterer{contract: contract}, nil
}

// bindStroe binds a generic wrapper to an already deployed contract.
func bindStroe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StroeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stroe *StroeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stroe.Contract.StroeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stroe *StroeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stroe.Contract.StroeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stroe *StroeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stroe.Contract.StroeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stroe *StroeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stroe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stroe *StroeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stroe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stroe *StroeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stroe.Contract.contract.Transact(opts, method, params...)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Stroe *StroeCaller) Items(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Stroe.contract.Call(opts, &out, "items", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Stroe *StroeSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Stroe.Contract.Items(&_Stroe.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Stroe *StroeCallerSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Stroe.Contract.Items(&_Stroe.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Stroe *StroeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stroe.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Stroe *StroeSession) Version() (string, error) {
	return _Stroe.Contract.Version(&_Stroe.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Stroe *StroeCallerSession) Version() (string, error) {
	return _Stroe.Contract.Version(&_Stroe.CallOpts)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Stroe *StroeTransactor) SetItem(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Stroe.contract.Transact(opts, "setItem", key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Stroe *StroeSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Stroe.Contract.SetItem(&_Stroe.TransactOpts, key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Stroe *StroeTransactorSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Stroe.Contract.SetItem(&_Stroe.TransactOpts, key, value)
}

// StroeItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Stroe contract.
type StroeItemSetIterator struct {
	Event *StroeItemSet // Event containing the contract specifics and raw log

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
func (it *StroeItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StroeItemSet)
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
		it.Event = new(StroeItemSet)
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
func (it *StroeItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StroeItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StroeItemSet represents a ItemSet event raised by the Stroe contract.
type StroeItemSet struct {
	Key   [32]byte
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Stroe *StroeFilterer) FilterItemSet(opts *bind.FilterOpts) (*StroeItemSetIterator, error) {

	logs, sub, err := _Stroe.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &StroeItemSetIterator{contract: _Stroe.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Stroe *StroeFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *StroeItemSet) (event.Subscription, error) {

	logs, sub, err := _Stroe.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StroeItemSet)
				if err := _Stroe.contract.UnpackLog(event, "ItemSet", log); err != nil {
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

// ParseItemSet is a log parse operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Stroe *StroeFilterer) ParseItemSet(log types.Log) (*StroeItemSet, error) {
	event := new(StroeItemSet)
	if err := _Stroe.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
