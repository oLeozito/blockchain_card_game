// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rarity\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CardMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"CardTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"MatchRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"rarity\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameServer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMatchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"matches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scoreWinner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scoreLoser\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_rarity\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"mintCard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"playerInventory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_loser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sWin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sLose\",\"type\":\"uint256\"}],\"name\":\"recordMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cardId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferCard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506112f78061005d5f395ff3fe608060405234801561000f575f80fd5b5060043610610086575f3560e01c80638c4f7dae116100595780638c4f7dae146101295780638dc1076814610147578063940bdd0814610179578063f0918c291461019557610086565b8063167196831461008a57806329cd564c146100a65780634768d4ef146100c45780637ba03bb4146100f9575b5f80fd5b6100a4600480360381019061009f9190610b3d565b6101b1565b005b6100ae6103dd565b6040516100bb9190610bb8565b60405180910390f35b6100de60048036038101906100d99190610bd1565b610402565b6040516100f096959493929190610c0b565b60405180910390f35b610113600480360381019061010e9190610c6a565b610487565b6040516101209190610ca8565b60405180910390f35b6101316104b2565b60405161013e9190610ca8565b60405180910390f35b610161600480360381019061015c9190610bd1565b6104be565b60405161017093929190610d3b565b60405180910390f35b610193600480360381019061018e9190610d77565b610588565b005b6101af60048036038101906101aa9190610ddb565b6107a3565b005b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610240576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023790610e75565b60405180910390fd5b5f805f8581526020019081526020015f205f015414610294576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028b90610edd565b60405180910390fd5b60405180606001604052808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff168152505f808581526020019081526020015f205f820151815f015560208201518160010190816102f291906110f5565b506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505060015f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2083908060018154018082558091505060019003905f5260205f20015f90919091909150557fc1858a9e1260d2afcf7bd0ebec9aed42b8c7f31f5ec3f434dc8c443640b46a1f8383836040516103d093929190610d3b565b60405180910390a1505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60028181548110610411575f80fd5b905f5260205f2090600602015f91509050805f015490806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154908060050154905086565b6001602052815f5260405f2081815481106104a0575f80fd5b905f5260205f20015f91509150505481565b5f600280549050905090565b5f602052805f5260405f205f91509050805f0154908060010180546104e290610f28565b80601f016020809104026020016040519081016040528092919081815260200182805461050e90610f28565b80156105595780601f1061053057610100808354040283529160200191610559565b820191905f5260205f20905b81548152906001019060200180831161053c57829003601f168201915b505050505090806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610617576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060e90610e75565b60405180910390fd5b5f600160028054905061062a91906111f1565b905060026040518060c001604052808381526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815260200142815250908060018154018082558091505060019003905f5260205f2090600602015f909190919091505f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a0820151816005015550507f0e6331c94271f4ed03c803999e345d5f9085ef55c8ad8ec84860358f984be4d081868660405161079493929190611224565b60405180910390a15050505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610832576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082990610e75565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff165f808581526020019081526020015f206002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146108d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c9906112a3565b60405180910390fd5b805f808581526020019081526020015f206002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f159d2cfced3bfb288ec8d5c8fb510f89a619a192bf756a7a701c117b09b1082383838360405161095693929190611224565b60405180910390a1505050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61098681610974565b8114610990575f80fd5b50565b5f813590506109a18161097d565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6109f5826109af565b810181811067ffffffffffffffff82111715610a1457610a136109bf565b5b80604052505050565b5f610a26610963565b9050610a3282826109ec565b919050565b5f67ffffffffffffffff821115610a5157610a506109bf565b5b610a5a826109af565b9050602081019050919050565b828183375f83830152505050565b5f610a87610a8284610a37565b610a1d565b905082815260208101848484011115610aa357610aa26109ab565b5b610aae848285610a67565b509392505050565b5f82601f830112610aca57610ac96109a7565b5b8135610ada848260208601610a75565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b0c82610ae3565b9050919050565b610b1c81610b02565b8114610b26575f80fd5b50565b5f81359050610b3781610b13565b92915050565b5f805f60608486031215610b5457610b5361096c565b5b5f610b6186828701610993565b935050602084013567ffffffffffffffff811115610b8257610b81610970565b5b610b8e86828701610ab6565b9250506040610b9f86828701610b29565b9150509250925092565b610bb281610b02565b82525050565b5f602082019050610bcb5f830184610ba9565b92915050565b5f60208284031215610be657610be561096c565b5b5f610bf384828501610993565b91505092915050565b610c0581610974565b82525050565b5f60c082019050610c1e5f830189610bfc565b610c2b6020830188610ba9565b610c386040830187610ba9565b610c456060830186610bfc565b610c526080830185610bfc565b610c5f60a0830184610bfc565b979650505050505050565b5f8060408385031215610c8057610c7f61096c565b5b5f610c8d85828601610b29565b9250506020610c9e85828601610993565b9150509250929050565b5f602082019050610cbb5f830184610bfc565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610cf8578082015181840152602081019050610cdd565b5f8484015250505050565b5f610d0d82610cc1565b610d178185610ccb565b9350610d27818560208601610cdb565b610d30816109af565b840191505092915050565b5f606082019050610d4e5f830186610bfc565b8181036020830152610d608185610d03565b9050610d6f6040830184610ba9565b949350505050565b5f805f8060808587031215610d8f57610d8e61096c565b5b5f610d9c87828801610b29565b9450506020610dad87828801610b29565b9350506040610dbe87828801610993565b9250506060610dcf87828801610993565b91505092959194509250565b5f805f60608486031215610df257610df161096c565b5b5f610dff86828701610993565b9350506020610e1086828701610b29565b9250506040610e2186828701610b29565b9150509250925092565b7f4170656e6173207365727669646f7200000000000000000000000000000000005f82015250565b5f610e5f600f83610ccb565b9150610e6a82610e2b565b602082019050919050565b5f6020820190508181035f830152610e8c81610e53565b9050919050565b7f4361727461206a612065786973746521000000000000000000000000000000005f82015250565b5f610ec7601083610ccb565b9150610ed282610e93565b602082019050919050565b5f6020820190508181035f830152610ef481610ebb565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610f3f57607f821691505b602082108103610f5257610f51610efb565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610fb47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610f79565b610fbe8683610f79565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610ff9610ff4610fef84610974565b610fd6565b610974565b9050919050565b5f819050919050565b61101283610fdf565b61102661101e82611000565b848454610f85565b825550505050565b5f90565b61103a61102e565b611045818484611009565b505050565b5b818110156110685761105d5f82611032565b60018101905061104b565b5050565b601f8211156110ad5761107e81610f58565b61108784610f6a565b81016020851015611096578190505b6110aa6110a285610f6a565b83018261104a565b50505b505050565b5f82821c905092915050565b5f6110cd5f19846008026110b2565b1980831691505092915050565b5f6110e583836110be565b9150826002028217905092915050565b6110fe82610cc1565b67ffffffffffffffff811115611117576111166109bf565b5b6111218254610f28565b61112c82828561106c565b5f60209050601f83116001811461115d575f841561114b578287015190505b61115585826110da565b8655506111bc565b601f19841661116b86610f58565b5f5b828110156111925784890151825560018201915060208501945060208101905061116d565b868310156111af57848901516111ab601f8916826110be565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6111fb82610974565b915061120683610974565b925082820190508082111561121e5761121d6111c4565b5b92915050565b5f6060820190506112375f830186610bfc565b6112446020830185610ba9565b6112516040830184610ba9565b949350505050565b7f52656d6574656e7465206e616f206520646f6e6f0000000000000000000000005f82015250565b5f61128d601483610ccb565b915061129882611259565b602082019050919050565b5f6020820190508181035f8301526112ba81611281565b905091905056fea2646970667358221220e297bb90ae7b8d3d0f9dc0b82fd00b4d5bba1aab51adc6b42ee923f8df9c3c4b64736f6c63430008140033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
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

// Cards is a free data retrieval call binding the contract method 0x8dc10768.
//
// Solidity: function cards(uint256 ) view returns(uint256 id, string rarity, address owner)
func (_Contracts *ContractsCaller) Cards(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id     *big.Int
	Rarity string
	Owner  common.Address
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "cards", arg0)

	outstruct := new(struct {
		Id     *big.Int
		Rarity string
		Owner  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Rarity = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Cards is a free data retrieval call binding the contract method 0x8dc10768.
//
// Solidity: function cards(uint256 ) view returns(uint256 id, string rarity, address owner)
func (_Contracts *ContractsSession) Cards(arg0 *big.Int) (struct {
	Id     *big.Int
	Rarity string
	Owner  common.Address
}, error) {
	return _Contracts.Contract.Cards(&_Contracts.CallOpts, arg0)
}

// Cards is a free data retrieval call binding the contract method 0x8dc10768.
//
// Solidity: function cards(uint256 ) view returns(uint256 id, string rarity, address owner)
func (_Contracts *ContractsCallerSession) Cards(arg0 *big.Int) (struct {
	Id     *big.Int
	Rarity string
	Owner  common.Address
}, error) {
	return _Contracts.Contract.Cards(&_Contracts.CallOpts, arg0)
}

// GameServer is a free data retrieval call binding the contract method 0x29cd564c.
//
// Solidity: function gameServer() view returns(address)
func (_Contracts *ContractsCaller) GameServer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "gameServer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameServer is a free data retrieval call binding the contract method 0x29cd564c.
//
// Solidity: function gameServer() view returns(address)
func (_Contracts *ContractsSession) GameServer() (common.Address, error) {
	return _Contracts.Contract.GameServer(&_Contracts.CallOpts)
}

// GameServer is a free data retrieval call binding the contract method 0x29cd564c.
//
// Solidity: function gameServer() view returns(address)
func (_Contracts *ContractsCallerSession) GameServer() (common.Address, error) {
	return _Contracts.Contract.GameServer(&_Contracts.CallOpts)
}

// GetMatchCount is a free data retrieval call binding the contract method 0x8c4f7dae.
//
// Solidity: function getMatchCount() view returns(uint256)
func (_Contracts *ContractsCaller) GetMatchCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMatchCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMatchCount is a free data retrieval call binding the contract method 0x8c4f7dae.
//
// Solidity: function getMatchCount() view returns(uint256)
func (_Contracts *ContractsSession) GetMatchCount() (*big.Int, error) {
	return _Contracts.Contract.GetMatchCount(&_Contracts.CallOpts)
}

// GetMatchCount is a free data retrieval call binding the contract method 0x8c4f7dae.
//
// Solidity: function getMatchCount() view returns(uint256)
func (_Contracts *ContractsCallerSession) GetMatchCount() (*big.Int, error) {
	return _Contracts.Contract.GetMatchCount(&_Contracts.CallOpts)
}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(uint256 matchId, address winner, address loser, uint256 scoreWinner, uint256 scoreLoser, uint256 timestamp)
func (_Contracts *ContractsCaller) Matches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MatchId     *big.Int
	Winner      common.Address
	Loser       common.Address
	ScoreWinner *big.Int
	ScoreLoser  *big.Int
	Timestamp   *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "matches", arg0)

	outstruct := new(struct {
		MatchId     *big.Int
		Winner      common.Address
		Loser       common.Address
		ScoreWinner *big.Int
		ScoreLoser  *big.Int
		Timestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MatchId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Winner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Loser = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ScoreWinner = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ScoreLoser = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(uint256 matchId, address winner, address loser, uint256 scoreWinner, uint256 scoreLoser, uint256 timestamp)
func (_Contracts *ContractsSession) Matches(arg0 *big.Int) (struct {
	MatchId     *big.Int
	Winner      common.Address
	Loser       common.Address
	ScoreWinner *big.Int
	ScoreLoser  *big.Int
	Timestamp   *big.Int
}, error) {
	return _Contracts.Contract.Matches(&_Contracts.CallOpts, arg0)
}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(uint256 matchId, address winner, address loser, uint256 scoreWinner, uint256 scoreLoser, uint256 timestamp)
func (_Contracts *ContractsCallerSession) Matches(arg0 *big.Int) (struct {
	MatchId     *big.Int
	Winner      common.Address
	Loser       common.Address
	ScoreWinner *big.Int
	ScoreLoser  *big.Int
	Timestamp   *big.Int
}, error) {
	return _Contracts.Contract.Matches(&_Contracts.CallOpts, arg0)
}

// PlayerInventory is a free data retrieval call binding the contract method 0x7ba03bb4.
//
// Solidity: function playerInventory(address , uint256 ) view returns(uint256)
func (_Contracts *ContractsCaller) PlayerInventory(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "playerInventory", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlayerInventory is a free data retrieval call binding the contract method 0x7ba03bb4.
//
// Solidity: function playerInventory(address , uint256 ) view returns(uint256)
func (_Contracts *ContractsSession) PlayerInventory(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.PlayerInventory(&_Contracts.CallOpts, arg0, arg1)
}

// PlayerInventory is a free data retrieval call binding the contract method 0x7ba03bb4.
//
// Solidity: function playerInventory(address , uint256 ) view returns(uint256)
func (_Contracts *ContractsCallerSession) PlayerInventory(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.PlayerInventory(&_Contracts.CallOpts, arg0, arg1)
}

// MintCard is a paid mutator transaction binding the contract method 0x16719683.
//
// Solidity: function mintCard(uint256 _id, string _rarity, address _owner) returns()
func (_Contracts *ContractsTransactor) MintCard(opts *bind.TransactOpts, _id *big.Int, _rarity string, _owner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "mintCard", _id, _rarity, _owner)
}

// MintCard is a paid mutator transaction binding the contract method 0x16719683.
//
// Solidity: function mintCard(uint256 _id, string _rarity, address _owner) returns()
func (_Contracts *ContractsSession) MintCard(_id *big.Int, _rarity string, _owner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.MintCard(&_Contracts.TransactOpts, _id, _rarity, _owner)
}

// MintCard is a paid mutator transaction binding the contract method 0x16719683.
//
// Solidity: function mintCard(uint256 _id, string _rarity, address _owner) returns()
func (_Contracts *ContractsTransactorSession) MintCard(_id *big.Int, _rarity string, _owner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.MintCard(&_Contracts.TransactOpts, _id, _rarity, _owner)
}

// RecordMatch is a paid mutator transaction binding the contract method 0x940bdd08.
//
// Solidity: function recordMatch(address _winner, address _loser, uint256 _sWin, uint256 _sLose) returns()
func (_Contracts *ContractsTransactor) RecordMatch(opts *bind.TransactOpts, _winner common.Address, _loser common.Address, _sWin *big.Int, _sLose *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "recordMatch", _winner, _loser, _sWin, _sLose)
}

// RecordMatch is a paid mutator transaction binding the contract method 0x940bdd08.
//
// Solidity: function recordMatch(address _winner, address _loser, uint256 _sWin, uint256 _sLose) returns()
func (_Contracts *ContractsSession) RecordMatch(_winner common.Address, _loser common.Address, _sWin *big.Int, _sLose *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RecordMatch(&_Contracts.TransactOpts, _winner, _loser, _sWin, _sLose)
}

// RecordMatch is a paid mutator transaction binding the contract method 0x940bdd08.
//
// Solidity: function recordMatch(address _winner, address _loser, uint256 _sWin, uint256 _sLose) returns()
func (_Contracts *ContractsTransactorSession) RecordMatch(_winner common.Address, _loser common.Address, _sWin *big.Int, _sLose *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RecordMatch(&_Contracts.TransactOpts, _winner, _loser, _sWin, _sLose)
}

// TransferCard is a paid mutator transaction binding the contract method 0xf0918c29.
//
// Solidity: function transferCard(uint256 _cardId, address _from, address _to) returns()
func (_Contracts *ContractsTransactor) TransferCard(opts *bind.TransactOpts, _cardId *big.Int, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferCard", _cardId, _from, _to)
}

// TransferCard is a paid mutator transaction binding the contract method 0xf0918c29.
//
// Solidity: function transferCard(uint256 _cardId, address _from, address _to) returns()
func (_Contracts *ContractsSession) TransferCard(_cardId *big.Int, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferCard(&_Contracts.TransactOpts, _cardId, _from, _to)
}

// TransferCard is a paid mutator transaction binding the contract method 0xf0918c29.
//
// Solidity: function transferCard(uint256 _cardId, address _from, address _to) returns()
func (_Contracts *ContractsTransactorSession) TransferCard(_cardId *big.Int, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferCard(&_Contracts.TransactOpts, _cardId, _from, _to)
}

// ContractsCardMintedIterator is returned from FilterCardMinted and is used to iterate over the raw logs and unpacked data for CardMinted events raised by the Contracts contract.
type ContractsCardMintedIterator struct {
	Event *ContractsCardMinted // Event containing the contract specifics and raw log

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
func (it *ContractsCardMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCardMinted)
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
		it.Event = new(ContractsCardMinted)
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
func (it *ContractsCardMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCardMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCardMinted represents a CardMinted event raised by the Contracts contract.
type ContractsCardMinted struct {
	Id     *big.Int
	Rarity string
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCardMinted is a free log retrieval operation binding the contract event 0xc1858a9e1260d2afcf7bd0ebec9aed42b8c7f31f5ec3f434dc8c443640b46a1f.
//
// Solidity: event CardMinted(uint256 id, string rarity, address owner)
func (_Contracts *ContractsFilterer) FilterCardMinted(opts *bind.FilterOpts) (*ContractsCardMintedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CardMinted")
	if err != nil {
		return nil, err
	}
	return &ContractsCardMintedIterator{contract: _Contracts.contract, event: "CardMinted", logs: logs, sub: sub}, nil
}

// WatchCardMinted is a free log subscription operation binding the contract event 0xc1858a9e1260d2afcf7bd0ebec9aed42b8c7f31f5ec3f434dc8c443640b46a1f.
//
// Solidity: event CardMinted(uint256 id, string rarity, address owner)
func (_Contracts *ContractsFilterer) WatchCardMinted(opts *bind.WatchOpts, sink chan<- *ContractsCardMinted) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CardMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCardMinted)
				if err := _Contracts.contract.UnpackLog(event, "CardMinted", log); err != nil {
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

// ParseCardMinted is a log parse operation binding the contract event 0xc1858a9e1260d2afcf7bd0ebec9aed42b8c7f31f5ec3f434dc8c443640b46a1f.
//
// Solidity: event CardMinted(uint256 id, string rarity, address owner)
func (_Contracts *ContractsFilterer) ParseCardMinted(log types.Log) (*ContractsCardMinted, error) {
	event := new(ContractsCardMinted)
	if err := _Contracts.contract.UnpackLog(event, "CardMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCardTransferredIterator is returned from FilterCardTransferred and is used to iterate over the raw logs and unpacked data for CardTransferred events raised by the Contracts contract.
type ContractsCardTransferredIterator struct {
	Event *ContractsCardTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsCardTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCardTransferred)
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
		it.Event = new(ContractsCardTransferred)
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
func (it *ContractsCardTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCardTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCardTransferred represents a CardTransferred event raised by the Contracts contract.
type ContractsCardTransferred struct {
	Id   *big.Int
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCardTransferred is a free log retrieval operation binding the contract event 0x159d2cfced3bfb288ec8d5c8fb510f89a619a192bf756a7a701c117b09b10823.
//
// Solidity: event CardTransferred(uint256 id, address from, address to)
func (_Contracts *ContractsFilterer) FilterCardTransferred(opts *bind.FilterOpts) (*ContractsCardTransferredIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CardTransferred")
	if err != nil {
		return nil, err
	}
	return &ContractsCardTransferredIterator{contract: _Contracts.contract, event: "CardTransferred", logs: logs, sub: sub}, nil
}

// WatchCardTransferred is a free log subscription operation binding the contract event 0x159d2cfced3bfb288ec8d5c8fb510f89a619a192bf756a7a701c117b09b10823.
//
// Solidity: event CardTransferred(uint256 id, address from, address to)
func (_Contracts *ContractsFilterer) WatchCardTransferred(opts *bind.WatchOpts, sink chan<- *ContractsCardTransferred) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CardTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCardTransferred)
				if err := _Contracts.contract.UnpackLog(event, "CardTransferred", log); err != nil {
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

// ParseCardTransferred is a log parse operation binding the contract event 0x159d2cfced3bfb288ec8d5c8fb510f89a619a192bf756a7a701c117b09b10823.
//
// Solidity: event CardTransferred(uint256 id, address from, address to)
func (_Contracts *ContractsFilterer) ParseCardTransferred(log types.Log) (*ContractsCardTransferred, error) {
	event := new(ContractsCardTransferred)
	if err := _Contracts.contract.UnpackLog(event, "CardTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMatchRecordedIterator is returned from FilterMatchRecorded and is used to iterate over the raw logs and unpacked data for MatchRecorded events raised by the Contracts contract.
type ContractsMatchRecordedIterator struct {
	Event *ContractsMatchRecorded // Event containing the contract specifics and raw log

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
func (it *ContractsMatchRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMatchRecorded)
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
		it.Event = new(ContractsMatchRecorded)
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
func (it *ContractsMatchRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMatchRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMatchRecorded represents a MatchRecorded event raised by the Contracts contract.
type ContractsMatchRecorded struct {
	MatchId *big.Int
	Winner  common.Address
	Loser   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMatchRecorded is a free log retrieval operation binding the contract event 0x0e6331c94271f4ed03c803999e345d5f9085ef55c8ad8ec84860358f984be4d0.
//
// Solidity: event MatchRecorded(uint256 matchId, address winner, address loser)
func (_Contracts *ContractsFilterer) FilterMatchRecorded(opts *bind.FilterOpts) (*ContractsMatchRecordedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MatchRecorded")
	if err != nil {
		return nil, err
	}
	return &ContractsMatchRecordedIterator{contract: _Contracts.contract, event: "MatchRecorded", logs: logs, sub: sub}, nil
}

// WatchMatchRecorded is a free log subscription operation binding the contract event 0x0e6331c94271f4ed03c803999e345d5f9085ef55c8ad8ec84860358f984be4d0.
//
// Solidity: event MatchRecorded(uint256 matchId, address winner, address loser)
func (_Contracts *ContractsFilterer) WatchMatchRecorded(opts *bind.WatchOpts, sink chan<- *ContractsMatchRecorded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MatchRecorded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMatchRecorded)
				if err := _Contracts.contract.UnpackLog(event, "MatchRecorded", log); err != nil {
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

// ParseMatchRecorded is a log parse operation binding the contract event 0x0e6331c94271f4ed03c803999e345d5f9085ef55c8ad8ec84860358f984be4d0.
//
// Solidity: event MatchRecorded(uint256 matchId, address winner, address loser)
func (_Contracts *ContractsFilterer) ParseMatchRecorded(log types.Log) (*ContractsMatchRecorded, error) {
	event := new(ContractsMatchRecorded)
	if err := _Contracts.contract.UnpackLog(event, "MatchRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
