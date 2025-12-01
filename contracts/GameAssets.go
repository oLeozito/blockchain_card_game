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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rarity\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CardMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"CardTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"MatchRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"rarity\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameServer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMatchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"matches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scoreWinner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scoreLoser\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_rarity\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"mintCard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"playerInventory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_loser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sWin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sLose\",\"type\":\"uint256\"}],\"name\":\"recordMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cardId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferCard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506117f88061005d5f395ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c80637ba03bb4116100645780637ba03bb4146101205780638c4f7dae146101505780638dc107681461016e578063940bdd08146101a0578063f8b2cb4f146101bc57610091565b8063167196831461009557806329cd564c146100b15780634768d4ef146100cf57806372f79d2414610104575b5f80fd5b6100af60048036038101906100aa9190610d81565b6101ec565b005b6100b9610418565b6040516100c69190610dfc565b60405180910390f35b6100e960048036038101906100e49190610e15565b61043d565b6040516100fb96959493929190610e4f565b60405180910390f35b61011e60048036038101906101199190610eae565b6104c2565b005b61013a60048036038101906101359190610eec565b610745565b6040516101479190610f2a565b60405180910390f35b610158610770565b6040516101659190610f2a565b60405180910390f35b61018860048036038101906101839190610e15565b61077c565b60405161019793929190610fbd565b60405180910390f35b6101ba60048036038101906101b59190610ff9565b610846565b005b6101d660048036038101906101d1919061105d565b610a61565b6040516101e39190610f2a565b60405180910390f35b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461027b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610272906110f8565b60405180910390fd5b5f805f8581526020019081526020015f205f0154146102cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c690611160565b60405180910390fd5b60405180606001604052808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff168152505f808581526020019081526020015f205f820151815f0155602082015181600101908161032d9190611378565b506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505060015f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2083908060018154018082558091505060019003905f5260205f20015f90919091909150557fc1858a9e1260d2afcf7bd0ebec9aed42b8c7f31f5ec3f434dc8c443640b46a1f83838360405161040b93929190610fbd565b60405180910390a1505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6002818154811061044c575f80fd5b905f5260205f2090600602015f91509050805f015490806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154908060050154905086565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610551576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610548906114b7565b60405180910390fd5b5f805f8481526020019081526020015f205f0154036105a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059c9061151f565b60405180910390fd5b5f805f8481526020019081526020015f206002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610649576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610640906115ad565b60405180910390fd5b815f808581526020019081526020015f206002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506106a48184610aaa565b60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2083908060018154018082558091505060019003905f5260205f20015f90919091909150557f159d2cfced3bfb288ec8d5c8fb510f89a619a192bf756a7a701c117b09b10823838284604051610738939291906115cb565b60405180910390a1505050565b6001602052815f5260405f20818154811061075e575f80fd5b905f5260205f20015f91509150505481565b5f600280549050905090565b5f602052805f5260405f205f91509050805f0154908060010180546107a0906111ab565b80601f01602080910402602001604051908101604052809291908181526020018280546107cc906111ab565b80156108175780601f106107ee57610100808354040283529160200191610817565b820191905f5260205f20905b8154815290600101906020018083116107fa57829003601f168201915b505050505090806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108cc90611670565b60405180910390fd5b5f60016002805490506108e891906116bb565b905060026040518060c001604052808381526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815260200142815250908060018154018082558091505060019003905f5260205f2090600602015f909190919091505f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a0820151816005015550507f0e6331c94271f4ed03c803999e345d5f9085ef55c8ad8ec84860358f984be4d0818686604051610a52939291906115cb565b60405180910390a15050505050565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490509050919050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f5b8180549050811015610ba15782828281548110610b0c57610b0b6116ee565b5b905f5260205f20015403610b8e578160018380549050610b2c919061171b565b81548110610b3d57610b3c6116ee565b5b905f5260205f200154828281548110610b5957610b586116ee565b5b905f5260205f20018190555081805480610b7657610b7561174e565b5b600190038181905f5260205f20015f90559055610ba1565b8080610b999061177b565b915050610aec565b50505050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b610bca81610bb8565b8114610bd4575f80fd5b50565b5f81359050610be581610bc1565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610c3982610bf3565b810181811067ffffffffffffffff82111715610c5857610c57610c03565b5b80604052505050565b5f610c6a610ba7565b9050610c768282610c30565b919050565b5f67ffffffffffffffff821115610c9557610c94610c03565b5b610c9e82610bf3565b9050602081019050919050565b828183375f83830152505050565b5f610ccb610cc684610c7b565b610c61565b905082815260208101848484011115610ce757610ce6610bef565b5b610cf2848285610cab565b509392505050565b5f82601f830112610d0e57610d0d610beb565b5b8135610d1e848260208601610cb9565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d5082610d27565b9050919050565b610d6081610d46565b8114610d6a575f80fd5b50565b5f81359050610d7b81610d57565b92915050565b5f805f60608486031215610d9857610d97610bb0565b5b5f610da586828701610bd7565b935050602084013567ffffffffffffffff811115610dc657610dc5610bb4565b5b610dd286828701610cfa565b9250506040610de386828701610d6d565b9150509250925092565b610df681610d46565b82525050565b5f602082019050610e0f5f830184610ded565b92915050565b5f60208284031215610e2a57610e29610bb0565b5b5f610e3784828501610bd7565b91505092915050565b610e4981610bb8565b82525050565b5f60c082019050610e625f830189610e40565b610e6f6020830188610ded565b610e7c6040830187610ded565b610e896060830186610e40565b610e966080830185610e40565b610ea360a0830184610e40565b979650505050505050565b5f8060408385031215610ec457610ec3610bb0565b5b5f610ed185828601610bd7565b9250506020610ee285828601610d6d565b9150509250929050565b5f8060408385031215610f0257610f01610bb0565b5b5f610f0f85828601610d6d565b9250506020610f2085828601610bd7565b9150509250929050565b5f602082019050610f3d5f830184610e40565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610f7a578082015181840152602081019050610f5f565b5f8484015250505050565b5f610f8f82610f43565b610f998185610f4d565b9350610fa9818560208601610f5d565b610fb281610bf3565b840191505092915050565b5f606082019050610fd05f830186610e40565b8181036020830152610fe28185610f85565b9050610ff16040830184610ded565b949350505050565b5f805f806080858703121561101157611010610bb0565b5b5f61101e87828801610d6d565b945050602061102f87828801610d6d565b935050604061104087828801610bd7565b925050606061105187828801610bd7565b91505092959194509250565b5f6020828403121561107257611071610bb0565b5b5f61107f84828501610d6d565b91505092915050565b7f4170656e6173207365727669646f7220706f64652063726961722063617274615f8201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b5f6110e2602183610f4d565b91506110ed82611088565b604082019050919050565b5f6020820190508181035f83015261110f816110d6565b9050919050565b7f4361727461206a612065786973746521000000000000000000000000000000005f82015250565b5f61114a601083610f4d565b915061115582611116565b602082019050919050565b5f6020820190508181035f8301526111778161113e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806111c257607f821691505b6020821081036111d5576111d461117e565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026112377fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826111fc565b61124186836111fc565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61127c61127761127284610bb8565b611259565b610bb8565b9050919050565b5f819050919050565b61129583611262565b6112a96112a182611283565b848454611208565b825550505050565b5f90565b6112bd6112b1565b6112c881848461128c565b505050565b5b818110156112eb576112e05f826112b5565b6001810190506112ce565b5050565b601f82111561133057611301816111db565b61130a846111ed565b81016020851015611319578190505b61132d611325856111ed565b8301826112cd565b50505b505050565b5f82821c905092915050565b5f6113505f1984600802611335565b1980831691505092915050565b5f6113688383611341565b9150826002028217905092915050565b61138182610f43565b67ffffffffffffffff81111561139a57611399610c03565b5b6113a482546111ab565b6113af8282856112ef565b5f60209050601f8311600181146113e0575f84156113ce578287015190505b6113d8858261135d565b86555061143f565b601f1984166113ee866111db565b5f5b82811015611415578489015182556001820191506020850194506020810190506113f0565b86831015611432578489015161142e601f891682611341565b8355505b6001600288020188555050505b505050505050565b7f4170656e6173207365727669646f7220706f6465206d6f7665722063617274615f8201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b5f6114a1602183610f4d565b91506114ac82611447565b604082019050919050565b5f6020820190508181035f8301526114ce81611495565b9050919050565b7f4361727461206e616f20657869737465000000000000000000000000000000005f82015250565b5f611509601083610f4d565b9150611514826114d5565b602082019050919050565b5f6020820190508181035f830152611536816114fd565b9050919050565b7f4e616f20706f6465207472616e7366657269722070617261207369206d65736d5f8201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b5f611597602183610f4d565b91506115a28261153d565b604082019050919050565b5f6020820190508181035f8301526115c48161158b565b9050919050565b5f6060820190506115de5f830186610e40565b6115eb6020830185610ded565b6115f86040830184610ded565b949350505050565b7f4170656e6173207365727669646f7220706f64652072656769737472617220705f8201527f6172746964617300000000000000000000000000000000000000000000000000602082015250565b5f61165a602783610f4d565b915061166582611600565b604082019050919050565b5f6020820190508181035f8301526116878161164e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6116c582610bb8565b91506116d083610bb8565b92508282019050808211156116e8576116e761168e565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f61172582610bb8565b915061173083610bb8565b92508282039050818111156117485761174761168e565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f61178582610bb8565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117b7576117b661168e565b5b60018201905091905056fea264697066735822122074b68f9583cf2cac3b1cea676da6a200875c2cd7e6e9002c5b72041468d10e9b64736f6c63430008140033",
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

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _owner) view returns(uint256)
func (_Contracts *ContractsCaller) GetBalance(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getBalance", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _owner) view returns(uint256)
func (_Contracts *ContractsSession) GetBalance(_owner common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetBalance(&_Contracts.CallOpts, _owner)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _owner) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetBalance(_owner common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetBalance(&_Contracts.CallOpts, _owner)
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

// TransferCard is a paid mutator transaction binding the contract method 0x72f79d24.
//
// Solidity: function transferCard(uint256 _cardId, address _to) returns()
func (_Contracts *ContractsTransactor) TransferCard(opts *bind.TransactOpts, _cardId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferCard", _cardId, _to)
}

// TransferCard is a paid mutator transaction binding the contract method 0x72f79d24.
//
// Solidity: function transferCard(uint256 _cardId, address _to) returns()
func (_Contracts *ContractsSession) TransferCard(_cardId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferCard(&_Contracts.TransactOpts, _cardId, _to)
}

// TransferCard is a paid mutator transaction binding the contract method 0x72f79d24.
//
// Solidity: function transferCard(uint256 _cardId, address _to) returns()
func (_Contracts *ContractsTransactorSession) TransferCard(_cardId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferCard(&_Contracts.TransactOpts, _cardId, _to)
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
