// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

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

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"admins_\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MARK_DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RATE_DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPOT_DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"callMarks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"strikeLevel\",\"type\":\"uint8\"}],\"name\":\"latestRoundMark\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"}],\"name\":\"latestRoundMarks\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"uint256[11]\",\"name\":\"\",\"type\":\"uint256[11]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundRate\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundSpot\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"putMarks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundId\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAdmin_\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"spot_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate_\",\"type\":\"uint256\"},{\"internalType\":\"uint256[11]\",\"name\":\"callMarks_\",\"type\":\"uint256[11]\"},{\"internalType\":\"uint256[11]\",\"name\":\"putMarks_\",\"type\":\"uint256[11]\"}],\"name\":\"setLatestData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620017d7380380620017d783398181016040528101906200003791906200041a565b620000576200004b6200016060201b60201c565b6200016860201b60201c565b62000068336200016860201b60201c565b6001601b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060005b815181101562000158576001601b6000848481518110620000e957620000e86200046b565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080806200014f90620004d3565b915050620000c3565b505062000520565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620002908262000245565b810181811067ffffffffffffffff82111715620002b257620002b162000256565b5b80604052505050565b6000620002c76200022c565b9050620002d5828262000285565b919050565b600067ffffffffffffffff821115620002f857620002f762000256565b5b602082029050602081019050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200033b826200030e565b9050919050565b6200034d816200032e565b81146200035957600080fd5b50565b6000815190506200036d8162000342565b92915050565b60006200038a6200038484620002da565b620002bb565b90508083825260208201905060208402830185811115620003b057620003af62000309565b5b835b81811015620003dd5780620003c888826200035c565b845260208401935050602081019050620003b2565b5050509392505050565b600082601f830112620003ff57620003fe62000240565b5b81516200041184826020860162000373565b91505092915050565b60006020828403121562000433576200043262000236565b5b600082015167ffffffffffffffff8111156200045457620004536200023b565b5b6200046284828501620003e7565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b6000620004e082620004c9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036200051557620005146200049a565b5b600182019050919050565b6112a780620005306000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80638cd221c9116100ad578063d090714e11610071578063d090714e146102f3578063d0e224f614610323578063eb9f86a014610356578063f28d6f8714610372578063f2fde38b1461039057610121565b80638cd221c9146102485780638da5cb5b146102665780638de4ea4014610284578063b38017f1146102b7578063c20b721e146102d557610121565b80635a604c52116100f45780635a604c52146101c05780635ec81e29146101de57806360304e32146101ff5780636f265b9314610220578063715018a61461023e57610121565b806324d7806c146101265780632c4e722e146101565780634b0bddd2146101745780635494dfc214610190575b600080fd5b610140600480360381019061013b9190610ab6565b6103ac565b60405161014d9190610afe565b60405180910390f35b61015e6103cc565b60405161016b9190610b32565b60405180910390f35b61018e60048036038101906101899190610b79565b6103d2565b005b6101aa60048036038101906101a59190610be5565b610435565b6040516101b79190610b32565b60405180910390f35b6101c8610450565b6040516101d59190610c2e565b60405180910390f35b6101e6610455565b6040516101f69493929190610c6e565b60405180910390f35b61020761048a565b6040516102179493929190610c6e565b60405180910390f35b6102286104bf565b6040516102359190610b32565b60405180910390f35b6102466104c5565b005b6102506104d9565b60405161025d9190610cb3565b60405180910390f35b61026e6104f5565b60405161027b9190610cdd565b60405180910390f35b61029e60048036038101906102999190610cf8565b61051e565b6040516102ae9493929190610dd0565b60405180910390f35b6102bf61060c565b6040516102cc9190610b32565b60405180910390f35b6102dd610612565b6040516102ea9190610c2e565b60405180910390f35b61030d60048036038101906103089190610be5565b610617565b60405161031a9190610b32565b60405180910390f35b61033d60048036038101906103389190610e44565b610632565b60405161034d9493929190610c6e565b60405180910390f35b610370600480360381019061036b9190610fcb565b6106db565b005b61037a6107f7565b6040516103879190610c2e565b60405180910390f35b6103aa60048036038101906103a59190610ab6565b6107fc565b005b601b6020528060005260406000206000915054906101000a900460ff1681565b60025481565b6103da61087f565b80601b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600381600b811061044557600080fd5b016000915090505481565b600481565b600080600080601960009054906101000a900469ffffffffffffffffffff16600154601a546012935093509350935090919293565b600080600080601960009054906101000a900469ffffffffffffffffffff16600254601a546004935093509350935090919293565b60015481565b6104cd61087f565b6104d760006108fd565b565b601960009054906101000a900469ffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006105286109c9565b600080841561059d57601960009054906101000a900469ffffffffffffffffffff166003601a54601282600b806020026040519081016040528092919082600b8015610589576020028201915b815481526020019060010190808311610575575b505050505092509350935093509350610605565b601960009054906101000a900469ffffffffffffffffffff16600e601a54601282600b806020026040519081016040528092919082600b80156105f5576020028201915b8154815260200190600101908083116105e1575b5050505050925093509350935093505b9193509193565b601a5481565b601281565b600e81600b811061062757600080fd5b016000915090505481565b600080600080600b8560ff161061064857600080fd5b851561069257601960009054906101000a900469ffffffffffffffffffff1660038660ff16600b811061067e5761067d611034565b5b0154601a54601293509350935093506106d2565b601960009054906101000a900469ffffffffffffffffffff16600e8660ff16600b81106106c2576106c1611034565b5b0154601a54601293509350935093505b92959194509250565b601b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610767576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075e906110e6565b60405180910390fd5b6001601960009054906101000a900469ffffffffffffffffffff1661078c9190611135565b601960006101000a81548169ffffffffffffffffffff021916908369ffffffffffffffffffff16021790555042601a81905550836001819055508260028190555081600390600b6107de9291906109ec565b5080600e90600b6107f09291906109ec565b5050505050565b601281565b61080461087f565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610873576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086a906111e5565b60405180910390fd5b61087c816108fd565b50565b6108876109c1565b73ffffffffffffffffffffffffffffffffffffffff166108a56104f5565b73ffffffffffffffffffffffffffffffffffffffff16146108fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f290611251565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b604051806101600160405280600b90602082028036833780820191505090505090565b82600b8101928215610a1b579160200282015b82811115610a1a5782518255916020019190600101906109ff565b5b509050610a289190610a2c565b5090565b5b80821115610a45576000816000905550600101610a2d565b5090565b6000604051905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610a8382610a58565b9050919050565b610a9381610a78565b8114610a9e57600080fd5b50565b600081359050610ab081610a8a565b92915050565b600060208284031215610acc57610acb610a53565b5b6000610ada84828501610aa1565b91505092915050565b60008115159050919050565b610af881610ae3565b82525050565b6000602082019050610b136000830184610aef565b92915050565b6000819050919050565b610b2c81610b19565b82525050565b6000602082019050610b476000830184610b23565b92915050565b610b5681610ae3565b8114610b6157600080fd5b50565b600081359050610b7381610b4d565b92915050565b60008060408385031215610b9057610b8f610a53565b5b6000610b9e85828601610aa1565b9250506020610baf85828601610b64565b9150509250929050565b610bc281610b19565b8114610bcd57600080fd5b50565b600081359050610bdf81610bb9565b92915050565b600060208284031215610bfb57610bfa610a53565b5b6000610c0984828501610bd0565b91505092915050565b600060ff82169050919050565b610c2881610c12565b82525050565b6000602082019050610c436000830184610c1f565b92915050565b600069ffffffffffffffffffff82169050919050565b610c6881610c49565b82525050565b6000608082019050610c836000830187610c5f565b610c906020830186610b23565b610c9d6040830185610b23565b610caa6060830184610c1f565b95945050505050565b6000602082019050610cc86000830184610c5f565b92915050565b610cd781610a78565b82525050565b6000602082019050610cf26000830184610cce565b92915050565b600060208284031215610d0e57610d0d610a53565b5b6000610d1c84828501610b64565b91505092915050565b6000600b9050919050565b600081905092915050565b6000819050919050565b610d4e81610b19565b82525050565b6000610d608383610d45565b60208301905092915050565b6000602082019050919050565b610d8281610d25565b610d8c8184610d30565b9250610d9782610d3b565b8060005b83811015610dc8578151610daf8782610d54565b9650610dba83610d6c565b925050600181019050610d9b565b505050505050565b60006101c082019050610de66000830187610c5f565b610df36020830186610d79565b610e01610180830185610b23565b610e0f6101a0830184610c1f565b95945050505050565b610e2181610c12565b8114610e2c57600080fd5b50565b600081359050610e3e81610e18565b92915050565b60008060408385031215610e5b57610e5a610a53565b5b6000610e6985828601610b64565b9250506020610e7a85828601610e2f565b9150509250929050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610ed282610e89565b810181811067ffffffffffffffff82111715610ef157610ef0610e9a565b5b80604052505050565b6000610f04610a49565b9050610f108282610ec9565b919050565b600067ffffffffffffffff821115610f3057610f2f610e9a565b5b602082029050919050565b600080fd5b6000610f53610f4e84610f15565b610efa565b90508060208402830185811115610f6d57610f6c610f3b565b5b835b81811015610f965780610f828882610bd0565b845260208401935050602081019050610f6f565b5050509392505050565b600082601f830112610fb557610fb4610e84565b5b600b610fc2848285610f40565b91505092915050565b6000806000806103008587031215610fe657610fe5610a53565b5b6000610ff487828801610bd0565b945050602061100587828801610bd0565b935050604061101687828801610fa0565b9250506101a061102887828801610fa0565b91505092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825260208201905092915050565b7f6f6e6c7941646d696e3a2063616c6c6572206973206e6f7420616e2061646d6960008201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b60006110d0602183611063565b91506110db82611074565b604082019050919050565b600060208201905081810360008301526110ff816110c3565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061114082610c49565b915061114b83610c49565b9250828201905069ffffffffffffffffffff81111561116d5761116c611106565b5b92915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006111cf602683611063565b91506111da82611173565b604082019050919050565b600060208201905081810360008301526111fe816111c2565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b600061123b602083611063565b915061124682611205565b602082019050919050565b6000602082019050818103600083015261126a8161122e565b905091905056fea2646970667358221220e91b83f7e6da6a7934dd1fb11f385d3b4e466b6305db506524d5f629c3a324b864736f6c63430008110033",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// OracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleMetaData.Bin instead.
var OracleBin = OracleMetaData.Bin

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend, admins_ []common.Address) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleBin), backend, admins_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// MARKDECIMALS is a free data retrieval call binding the contract method 0xc20b721e.
//
// Solidity: function MARK_DECIMALS() view returns(uint8)
func (_Oracle *OracleCaller) MARKDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "MARK_DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MARKDECIMALS is a free data retrieval call binding the contract method 0xc20b721e.
//
// Solidity: function MARK_DECIMALS() view returns(uint8)
func (_Oracle *OracleSession) MARKDECIMALS() (uint8, error) {
	return _Oracle.Contract.MARKDECIMALS(&_Oracle.CallOpts)
}

// MARKDECIMALS is a free data retrieval call binding the contract method 0xc20b721e.
//
// Solidity: function MARK_DECIMALS() view returns(uint8)
func (_Oracle *OracleCallerSession) MARKDECIMALS() (uint8, error) {
	return _Oracle.Contract.MARKDECIMALS(&_Oracle.CallOpts)
}

// RATEDECIMALS is a free data retrieval call binding the contract method 0x5a604c52.
//
// Solidity: function RATE_DECIMALS() view returns(uint8)
func (_Oracle *OracleCaller) RATEDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "RATE_DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RATEDECIMALS is a free data retrieval call binding the contract method 0x5a604c52.
//
// Solidity: function RATE_DECIMALS() view returns(uint8)
func (_Oracle *OracleSession) RATEDECIMALS() (uint8, error) {
	return _Oracle.Contract.RATEDECIMALS(&_Oracle.CallOpts)
}

// RATEDECIMALS is a free data retrieval call binding the contract method 0x5a604c52.
//
// Solidity: function RATE_DECIMALS() view returns(uint8)
func (_Oracle *OracleCallerSession) RATEDECIMALS() (uint8, error) {
	return _Oracle.Contract.RATEDECIMALS(&_Oracle.CallOpts)
}

// SPOTDECIMALS is a free data retrieval call binding the contract method 0xf28d6f87.
//
// Solidity: function SPOT_DECIMALS() view returns(uint8)
func (_Oracle *OracleCaller) SPOTDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "SPOT_DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SPOTDECIMALS is a free data retrieval call binding the contract method 0xf28d6f87.
//
// Solidity: function SPOT_DECIMALS() view returns(uint8)
func (_Oracle *OracleSession) SPOTDECIMALS() (uint8, error) {
	return _Oracle.Contract.SPOTDECIMALS(&_Oracle.CallOpts)
}

// SPOTDECIMALS is a free data retrieval call binding the contract method 0xf28d6f87.
//
// Solidity: function SPOT_DECIMALS() view returns(uint8)
func (_Oracle *OracleCallerSession) SPOTDECIMALS() (uint8, error) {
	return _Oracle.Contract.SPOTDECIMALS(&_Oracle.CallOpts)
}

// CallMarks is a free data retrieval call binding the contract method 0x5494dfc2.
//
// Solidity: function callMarks(uint256 ) view returns(uint256)
func (_Oracle *OracleCaller) CallMarks(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "callMarks", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallMarks is a free data retrieval call binding the contract method 0x5494dfc2.
//
// Solidity: function callMarks(uint256 ) view returns(uint256)
func (_Oracle *OracleSession) CallMarks(arg0 *big.Int) (*big.Int, error) {
	return _Oracle.Contract.CallMarks(&_Oracle.CallOpts, arg0)
}

// CallMarks is a free data retrieval call binding the contract method 0x5494dfc2.
//
// Solidity: function callMarks(uint256 ) view returns(uint256)
func (_Oracle *OracleCallerSession) CallMarks(arg0 *big.Int) (*big.Int, error) {
	return _Oracle.Contract.CallMarks(&_Oracle.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Oracle *OracleCaller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Oracle *OracleSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _Oracle.Contract.IsAdmin(&_Oracle.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Oracle *OracleCallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _Oracle.Contract.IsAdmin(&_Oracle.CallOpts, arg0)
}

// LatestRoundMark is a free data retrieval call binding the contract method 0xd0e224f6.
//
// Solidity: function latestRoundMark(bool isCall, uint8 strikeLevel) view returns(uint80, uint256, uint256, uint8)
func (_Oracle *OracleCaller) LatestRoundMark(opts *bind.CallOpts, isCall bool, strikeLevel uint8) (*big.Int, *big.Int, *big.Int, uint8, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "latestRoundMark", isCall, strikeLevel)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return out0, out1, out2, out3, err

}

// LatestRoundMark is a free data retrieval call binding the contract method 0xd0e224f6.
//
// Solidity: function latestRoundMark(bool isCall, uint8 strikeLevel) view returns(uint80, uint256, uint256, uint8)
func (_Oracle *OracleSession) LatestRoundMark(isCall bool, strikeLevel uint8) (*big.Int, *big.Int, *big.Int, uint8, error) {
	return _Oracle.Contract.LatestRoundMark(&_Oracle.CallOpts, isCall, strikeLevel)
}

// LatestRoundMark is a free data retrieval call binding the contract method 0xd0e224f6.
//
// Solidity: function latestRoundMark(bool isCall, uint8 strikeLevel) view returns(uint80, uint256, uint256, uint8)
func (_Oracle *OracleCallerSession) LatestRoundMark(isCall bool, strikeLevel uint8) (*big.Int, *big.Int, *big.Int, uint8, error) {
	return _Oracle.Contract.LatestRoundMark(&_Oracle.CallOpts, isCall, strikeLevel)
}

// LatestRoundMarks is a free data retrieval call binding the contract method 0x8de4ea40.
//
// Solidity: function latestRoundMarks(bool isCall) view returns(uint80, uint256[11], uint256, uint8)
func (_Oracle *OracleCaller) LatestRoundMarks(opts *bind.CallOpts, isCall bool) (*big.Int, [11]*big.Int, *big.Int, uint8, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "latestRoundMarks", isCall)

	if err != nil {
		return *new(*big.Int), *new([11]*big.Int), *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([11]*big.Int)).(*[11]*big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return out0, out1, out2, out3, err

}

// LatestRoundMarks is a free data retrieval call binding the contract method 0x8de4ea40.
//
// Solidity: function latestRoundMarks(bool isCall) view returns(uint80, uint256[11], uint256, uint8)
func (_Oracle *OracleSession) LatestRoundMarks(isCall bool) (*big.Int, [11]*big.Int, *big.Int, uint8, error) {
	return _Oracle.Contract.LatestRoundMarks(&_Oracle.CallOpts, isCall)
}

// LatestRoundMarks is a free data retrieval call binding the contract method 0x8de4ea40.
//
// Solidity: function latestRoundMarks(bool isCall) view returns(uint80, uint256[11], uint256, uint8)
func (_Oracle *OracleCallerSession) LatestRoundMarks(isCall bool) (*big.Int, [11]*big.Int, *big.Int, uint8, error) {
	return _Oracle.Contract.LatestRoundMarks(&_Oracle.CallOpts, isCall)
}

// LatestRoundRate is a free data retrieval call binding the contract method 0x60304e32.
//
// Solidity: function latestRoundRate() view returns(uint80, uint256, uint256, uint8)
func (_Oracle *OracleCaller) LatestRoundRate(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, uint8, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "latestRoundRate")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return out0, out1, out2, out3, err

}

// LatestRoundRate is a free data retrieval call binding the contract method 0x60304e32.
//
// Solidity: function latestRoundRate() view returns(uint80, uint256, uint256, uint8)
func (_Oracle *OracleSession) LatestRoundRate() (*big.Int, *big.Int, *big.Int, uint8, error) {
	return _Oracle.Contract.LatestRoundRate(&_Oracle.CallOpts)
}

// LatestRoundRate is a free data retrieval call binding the contract method 0x60304e32.
//
// Solidity: function latestRoundRate() view returns(uint80, uint256, uint256, uint8)
func (_Oracle *OracleCallerSession) LatestRoundRate() (*big.Int, *big.Int, *big.Int, uint8, error) {
	return _Oracle.Contract.LatestRoundRate(&_Oracle.CallOpts)
}

// LatestRoundSpot is a free data retrieval call binding the contract method 0x5ec81e29.
//
// Solidity: function latestRoundSpot() view returns(uint80, uint256, uint256, uint8)
func (_Oracle *OracleCaller) LatestRoundSpot(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, uint8, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "latestRoundSpot")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return out0, out1, out2, out3, err

}

// LatestRoundSpot is a free data retrieval call binding the contract method 0x5ec81e29.
//
// Solidity: function latestRoundSpot() view returns(uint80, uint256, uint256, uint8)
func (_Oracle *OracleSession) LatestRoundSpot() (*big.Int, *big.Int, *big.Int, uint8, error) {
	return _Oracle.Contract.LatestRoundSpot(&_Oracle.CallOpts)
}

// LatestRoundSpot is a free data retrieval call binding the contract method 0x5ec81e29.
//
// Solidity: function latestRoundSpot() view returns(uint80, uint256, uint256, uint8)
func (_Oracle *OracleCallerSession) LatestRoundSpot() (*big.Int, *big.Int, *big.Int, uint8, error) {
	return _Oracle.Contract.LatestRoundSpot(&_Oracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleCallerSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// PutMarks is a free data retrieval call binding the contract method 0xd090714e.
//
// Solidity: function putMarks(uint256 ) view returns(uint256)
func (_Oracle *OracleCaller) PutMarks(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "putMarks", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PutMarks is a free data retrieval call binding the contract method 0xd090714e.
//
// Solidity: function putMarks(uint256 ) view returns(uint256)
func (_Oracle *OracleSession) PutMarks(arg0 *big.Int) (*big.Int, error) {
	return _Oracle.Contract.PutMarks(&_Oracle.CallOpts, arg0)
}

// PutMarks is a free data retrieval call binding the contract method 0xd090714e.
//
// Solidity: function putMarks(uint256 ) view returns(uint256)
func (_Oracle *OracleCallerSession) PutMarks(arg0 *big.Int) (*big.Int, error) {
	return _Oracle.Contract.PutMarks(&_Oracle.CallOpts, arg0)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Oracle *OracleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Oracle *OracleSession) Rate() (*big.Int, error) {
	return _Oracle.Contract.Rate(&_Oracle.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Oracle *OracleCallerSession) Rate() (*big.Int, error) {
	return _Oracle.Contract.Rate(&_Oracle.CallOpts)
}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_Oracle *OracleCaller) RoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "roundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_Oracle *OracleSession) RoundId() (*big.Int, error) {
	return _Oracle.Contract.RoundId(&_Oracle.CallOpts)
}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_Oracle *OracleCallerSession) RoundId() (*big.Int, error) {
	return _Oracle.Contract.RoundId(&_Oracle.CallOpts)
}

// RoundTimestamp is a free data retrieval call binding the contract method 0xb38017f1.
//
// Solidity: function roundTimestamp() view returns(uint256)
func (_Oracle *OracleCaller) RoundTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "roundTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundTimestamp is a free data retrieval call binding the contract method 0xb38017f1.
//
// Solidity: function roundTimestamp() view returns(uint256)
func (_Oracle *OracleSession) RoundTimestamp() (*big.Int, error) {
	return _Oracle.Contract.RoundTimestamp(&_Oracle.CallOpts)
}

// RoundTimestamp is a free data retrieval call binding the contract method 0xb38017f1.
//
// Solidity: function roundTimestamp() view returns(uint256)
func (_Oracle *OracleCallerSession) RoundTimestamp() (*big.Int, error) {
	return _Oracle.Contract.RoundTimestamp(&_Oracle.CallOpts)
}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(uint256)
func (_Oracle *OracleCaller) Spot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "spot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(uint256)
func (_Oracle *OracleSession) Spot() (*big.Int, error) {
	return _Oracle.Contract.Spot(&_Oracle.CallOpts)
}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(uint256)
func (_Oracle *OracleCallerSession) Spot() (*big.Int, error) {
	return _Oracle.Contract.Spot(&_Oracle.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oracle *OracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oracle *OracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oracle.Contract.RenounceOwnership(&_Oracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oracle *OracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oracle.Contract.RenounceOwnership(&_Oracle.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address account_, bool isAdmin_) returns()
func (_Oracle *OracleTransactor) SetAdmin(opts *bind.TransactOpts, account_ common.Address, isAdmin_ bool) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setAdmin", account_, isAdmin_)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address account_, bool isAdmin_) returns()
func (_Oracle *OracleSession) SetAdmin(account_ common.Address, isAdmin_ bool) (*types.Transaction, error) {
	return _Oracle.Contract.SetAdmin(&_Oracle.TransactOpts, account_, isAdmin_)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address account_, bool isAdmin_) returns()
func (_Oracle *OracleTransactorSession) SetAdmin(account_ common.Address, isAdmin_ bool) (*types.Transaction, error) {
	return _Oracle.Contract.SetAdmin(&_Oracle.TransactOpts, account_, isAdmin_)
}

// SetLatestData is a paid mutator transaction binding the contract method 0xeb9f86a0.
//
// Solidity: function setLatestData(uint256 spot_, uint256 rate_, uint256[11] callMarks_, uint256[11] putMarks_) returns()
func (_Oracle *OracleTransactor) SetLatestData(opts *bind.TransactOpts, spot_ *big.Int, rate_ *big.Int, callMarks_ [11]*big.Int, putMarks_ [11]*big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setLatestData", spot_, rate_, callMarks_, putMarks_)
}

// SetLatestData is a paid mutator transaction binding the contract method 0xeb9f86a0.
//
// Solidity: function setLatestData(uint256 spot_, uint256 rate_, uint256[11] callMarks_, uint256[11] putMarks_) returns()
func (_Oracle *OracleSession) SetLatestData(spot_ *big.Int, rate_ *big.Int, callMarks_ [11]*big.Int, putMarks_ [11]*big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetLatestData(&_Oracle.TransactOpts, spot_, rate_, callMarks_, putMarks_)
}

// SetLatestData is a paid mutator transaction binding the contract method 0xeb9f86a0.
//
// Solidity: function setLatestData(uint256 spot_, uint256 rate_, uint256[11] callMarks_, uint256[11] putMarks_) returns()
func (_Oracle *OracleTransactorSession) SetLatestData(spot_ *big.Int, rate_ *big.Int, callMarks_ [11]*big.Int, putMarks_ [11]*big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetLatestData(&_Oracle.TransactOpts, spot_, rate_, callMarks_, putMarks_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oracle *OracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oracle *OracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.TransferOwnership(&_Oracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oracle *OracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.TransferOwnership(&_Oracle.TransactOpts, newOwner)
}

// OracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oracle contract.
type OracleOwnershipTransferredIterator struct {
	Event *OracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleOwnershipTransferred)
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
		it.Event = new(OracleOwnershipTransferred)
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
func (it *OracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleOwnershipTransferred represents a OwnershipTransferred event raised by the Oracle contract.
type OracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oracle *OracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OracleOwnershipTransferredIterator{contract: _Oracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oracle *OracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleOwnershipTransferred)
				if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseOwnershipTransferred(log types.Log) (*OracleOwnershipTransferred, error) {
	event := new(OracleOwnershipTransferred)
	if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
