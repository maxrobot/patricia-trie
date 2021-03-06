// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PatriciaABI is the input ABI used to generate the binding from.
const PatriciaABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"insert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getNode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"getProof\",\"outputs\":[{\"name\":\"branchMask\",\"type\":\"uint256\"},{\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootEdge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"branchMask\",\"type\":\"uint256\"},{\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PatriciaBin is the compiled bytecode used for deploying new contracts.
const PatriciaBin = `608060405234801561001057600080fd5b50611792806100206000396000f300608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806320ba5b601461007d57806350c946fe1461012c578063693ac4fb146101b4578063a43914da14610279578063ebf0c717146102c2578063f7e498f6146102f5575b600080fd5b34801561008957600080fd5b5061012a600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506103ff565b005b34801561013857600080fd5b5061015b6004803603810190808035600019169060200190929190505050610620565b604051808781526020018660001916600019168152602001856000191660001916815260200184815260200183600019166000191681526020018260001916600019168152602001965050505050505060405180910390f35b3480156101c057600080fd5b5061021b600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610706565b6040518083815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015610264578082015181840152602081019050610249565b50505050905001935050505060405180910390f35b34801561028557600080fd5b5061028e610a5e565b6040518084815260200183600019166000191681526020018260001916600019168152602001935050505060405180910390f35b3480156102ce57600080fd5b506102d7610a86565b60405180826000191660001916815260200191505060405180910390f35b34801561030157600080fd5b506103fd6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192908035906020019092919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610a8c565b005b610407611567565b6000610411611584565b6040805190810160405280866040518082805190602001908083835b602083101515610452578051825260208201915060208101905060208303925061042d565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191681526020016101008152509250836040518082805190602001908083835b6020831015156104c657805182526020820191506020810190506020830392506104a1565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020915083600080846000191660001916815260200190815260200160002090805190602001906105239291906115a8565b5060006001026003600001546000191614801561054857506000600360010160010154145b1561056f5782816020018190525081816000019060001916908160001916815250506105cb565b6105c8600360408051908101604052908160008201546000191660001916815260200160018201604080519081016040529081600082015460001916600019168152602001600182015481525050815250508484610cfb565b90505b6105d481610e56565b6002816000191690555080600360008201518160000190600019169055602082015181600101600082015181600001906000191690556020820151816001015550509050505050505050565b6000806000806000806000600160008960001916600019168152602001908152602001600020905080600001600060028110151561065a57fe5b600302016001016001015481600001600060028110151561067757fe5b600302016001016000015482600001600060028110151561069457fe5b60030201600001548360000160016002811015156106ae57fe5b60030201600101600101548460000160016002811015156106cb57fe5b60030201600101600001548560000160016002811015156106e857fe5b60030201600001549650965096509650965096505091939550919395565b60006060610712611567565b61071a611584565b610722611628565b60008061072d611567565b610735611567565b600061073f611567565b600060408051908101604052808e6040518082805190602001908083835b602083101515610782578051825260208201915060208101905060208303925061075d565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916815260200161010081525099506003604080519081016040529081600082015460001916600019168152602001600182016040805190810160405290816000820154600019166000191681526020016001820154815250508152505098505b6001156109bf576108278a8a60200151610ea9565b94509450886020015160200151856020015114151561084557600080fd5b600084602001511415610857576109bf565b8460200151870196508660ff0360019060020a028c179b5060018701965061087e84610ed7565b9250925061090d600160008b6000015160001916600019168152602001908152602001600020600001846001036002811015156108b757fe5b600302016040805190810160405290816000820154600019166000191681526020016001820160408051908101604052908160008201546000191660001916815260200160018201548152505081525050610e56565b88878060010198506101008110151561092257fe5b60200201906000191690816000191681525050600160008a60000151600019166000191681526020019081526020016000206000018360028110151561096457fe5b6003020160408051908101604052908160008201546000191660001916815260200160018201604080519081016040529081600082015460001916600019168152602001600182015481525050815250509850819950610812565b6000861115610a4f57856040519080825280602002602001820160405280156109f75781602001602082028038833980820191505090505b509a50600090505b85811015610a4e57878161010081101515610a1657fe5b60200201518b82815181101515610a2957fe5b90602001906020020190600019169081600019168152505080806001019150506109ff565b5b50505050505050505050915091565b6000806000600360010160010154600360010160000154600360000154925092509250909192565b60025481565b610a94611567565b610a9c611584565b6000806000610aa961164d565b60408051908101604052808b6040518082805190602001908083835b602083101515610aea5780518252602082019150602081019050602083039250610ac5565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191681526020016101008152509550886040518082805190602001908083835b602083101515610b5e5780518252602082019150602081019050602083039250610b39565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902085600001906000191690816000191681525050600093505b600088141515610cc757610bb688610f44565b92508260019060020a021988169750610bd2868460ff036110b4565b809050866020018190528197505050610bee8560200151610ed7565b809050866020018190528193505050610c0685610e56565b8183600281101515610c1457fe5b602002019060001916908160001916815250508660018589510303815181101515610c3b57fe5b906020019060200201518183600103600281101515610c5657fe5b60200201906000191690816000191681525050806040518082600260200280838360005b83811015610c95578082015181840152602081019050610c7a565b505050509050019150506040518091039020856000019060001916908160001916815250508380600101945050610ba3565b858560200181905250610cd985610e56565b600019168b60001916141515610cee57600080fd5b5050505050505050505050565b610d03611584565b610d0b611567565b610d13611567565b600080610d1e611567565b610d2661166f565b896020015160200151896020015110151515610d4157600080fd5b610d4f898b60200151610ea9565b95509550600085602001511415610d6857879350610e2e565b8960200151602001518660200151101515610d9357610d8c8a60000151868a61117b565b9350610e2d565b610d9c85610ed7565b9250925060408051908101604052808960001916815260200183815250816000015184600281101515610dcb57fe5b602002018190525060408051908101604052808b60000151600019168152602001610e018c6020015160018a60200151016112b7565b815250816000015184600103600281101515610e1957fe5b6020020181905250610e2a8161130a565b93505b5b6040805190810160405280856000191681526020018781525096505050505050509392505050565b600081600001518260200151602001518360200151600001516040518084600019166000191681526020018381526020018260001916600019168152602001935050505060405180910390209050919050565b610eb1611567565b610eb9611567565b610ecc84610ec78587611421565b6110b4565b915091509250929050565b6000610ee1611567565b60008360200151111515610ef457600080fd5b60ff8360000151600019169060020a900460019004604080519081016040528060018660000151600019169060020a02600019168152602001600186602001510381525080905091509150915091565b6000806000806000808614151515610f5b57600080fd5b856001029350600092505b60208310156110025760007f0100000000000000000000000000000000000000000000000000000000000000028484601f03602081101515610fa457fe5b1a7f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916141515610ff557611002565b8280600101935050610f66565b8383601f0360208110151561101357fe5b1a7f0100000000000000000000000000000000000000000000000000000000000000027f0100000000000000000000000000000000000000000000000000000000000000900460ff16915060019050600094505b61010085101561109c57600081831614151561108a5784836008020194506110ab565b80810190508480600101955050611067565b600015156110a657fe5b600094505b50505050919050565b6110bc611567565b6110c4611567565b836020015183111580156110da57506101008311155b15156110e557600080fd5b82826020018181525050600083141561111557600060010282600001906000191690816000191681525050611142565b6001836101000360019060020a020360010219846000015116826000019060001916908160001916815250505b82846020015103816020018181525050828460000151600019169060020a02816000019060001916908160001916815250509250929050565b600061118561166f565b600061118f611567565b600186602001511115156111a257600080fd5b6001600088600019166000191681526020019081526020016000206020604051908101604052908160008201600280602002604051908101604052809291906000905b8282101561124f578382600302016040805190810160405290816000820154600019166000191681526020016001820160408051908101604052908160008201546000191660001916815260200160018201548152505081525050815260200190600101906111e5565b5050505081525050925061126286610ed7565b9150915061128783600001518360028110151561127b57fe5b60200201518287610cfb565b83600001518360028110151561129957fe5b60200201819052506112ab87846114b1565b93505050509392505050565b6112bf611567565b826020015182111515156112d257600080fd5b81836020015103816020018181525050818360000151600019169060020a028160000190600019169081600019168152505092915050565b600080611316836114f2565b90508260000151600060028110151561132b57fe5b6020020151600160008360001916600019168152602001908152602001600020600001600060028110151561135c57fe5b600302016000820151816000019060001916905560208201518160010160008201518160000190600019169055602082015181600101555050905050826000015160016002811015156113ab57fe5b602002015160016000836000191660001916815260200190815260200160002060000160016002811015156113dc57fe5b60030201600082015181600001906000191690556020820151816001016000820151816000019060001916905560208201518160010155505090505080915050919050565b6000806000808460200151866020015110611440578460200151611446565b85602001515b925084600001518660000151186001900491507f800000000000000000000000000000000000000000000000000000000000000090505b828410156114a8576000828216141515611496576114a8565b8182019150838060010194505061147d565b50505092915050565b6000600160008460001916600019168152602001908152602001600020600080820160006114df9190611689565b50506114ea8261130a565b905092915050565b60006115148260000151600060028110151561150a57fe5b6020020151610e56565b6115348360000151600160028110151561152a57fe5b6020020151610e56565b60405180836000191660001916815260200182600019166000191681526020019250505060405180910390209050919050565b604080519081016040528060008019168152602001600081525090565b606060405190810160405280600080191681526020016115a26116d2565b81525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106115e957805160ff1916838001178555611617565b82800160010185558215611617579182015b828111156116165782518255916020019190600101906115fb565b5b50905061162491906116ef565b5090565b6120006040519081016040528061010090602082028038833980820191505090505090565b6040805190810160405280600290602082028038833980820191505090505090565b60c060405190810160405280611683611714565b81525090565b5060008082016000905560018201600080820160009055600182016000905550505060030160008082016000905560018201600080820160009055600182016000905550505050565b604080519081016040528060008019168152602001600081525090565b61171191905b8082111561170d5760008160009055506001016116f5565b5090565b90565b60c0604051908101604052806002905b61172c611742565b8152602001906001900390816117245790505090565b606060405190810160405280600080191681526020016117606116d2565b815250905600a165627a7a72305820e1da700e0233fe4039457d4464632da485087da583fd01e841f873e38bedc17f0029`

// DeployPatricia deploys a new Ethereum contract, binding an instance of Patricia to it.
func DeployPatricia(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Patricia, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PatriciaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Patricia{PatriciaCaller: PatriciaCaller{contract: contract}, PatriciaTransactor: PatriciaTransactor{contract: contract}, PatriciaFilterer: PatriciaFilterer{contract: contract}}, nil
}

// Patricia is an auto generated Go binding around an Ethereum contract.
type Patricia struct {
	PatriciaCaller     // Read-only binding to the contract
	PatriciaTransactor // Write-only binding to the contract
	PatriciaFilterer   // Log filterer for contract events
}

// PatriciaCaller is an auto generated read-only Go binding around an Ethereum contract.
type PatriciaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PatriciaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PatriciaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PatriciaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PatriciaSession struct {
	Contract     *Patricia         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PatriciaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PatriciaCallerSession struct {
	Contract *PatriciaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PatriciaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PatriciaTransactorSession struct {
	Contract     *PatriciaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PatriciaRaw is an auto generated low-level Go binding around an Ethereum contract.
type PatriciaRaw struct {
	Contract *Patricia // Generic contract binding to access the raw methods on
}

// PatriciaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PatriciaCallerRaw struct {
	Contract *PatriciaCaller // Generic read-only contract binding to access the raw methods on
}

// PatriciaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PatriciaTransactorRaw struct {
	Contract *PatriciaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPatricia creates a new instance of Patricia, bound to a specific deployed contract.
func NewPatricia(address common.Address, backend bind.ContractBackend) (*Patricia, error) {
	contract, err := bindPatricia(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Patricia{PatriciaCaller: PatriciaCaller{contract: contract}, PatriciaTransactor: PatriciaTransactor{contract: contract}, PatriciaFilterer: PatriciaFilterer{contract: contract}}, nil
}

// NewPatriciaCaller creates a new read-only instance of Patricia, bound to a specific deployed contract.
func NewPatriciaCaller(address common.Address, caller bind.ContractCaller) (*PatriciaCaller, error) {
	contract, err := bindPatricia(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaCaller{contract: contract}, nil
}

// NewPatriciaTransactor creates a new write-only instance of Patricia, bound to a specific deployed contract.
func NewPatriciaTransactor(address common.Address, transactor bind.ContractTransactor) (*PatriciaTransactor, error) {
	contract, err := bindPatricia(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PatriciaTransactor{contract: contract}, nil
}

// NewPatriciaFilterer creates a new log filterer instance of Patricia, bound to a specific deployed contract.
func NewPatriciaFilterer(address common.Address, filterer bind.ContractFilterer) (*PatriciaFilterer, error) {
	contract, err := bindPatricia(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PatriciaFilterer{contract: contract}, nil
}

// bindPatricia binds a generic wrapper to an already deployed contract.
func bindPatricia(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PatriciaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Patricia *PatriciaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Patricia.Contract.PatriciaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Patricia *PatriciaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Patricia.Contract.PatriciaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Patricia *PatriciaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Patricia.Contract.PatriciaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Patricia *PatriciaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Patricia.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Patricia *PatriciaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Patricia.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Patricia *PatriciaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Patricia.Contract.contract.Transact(opts, method, params...)
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(uint256, bytes32, bytes32, uint256, bytes32, bytes32)
func (_Patricia *PatriciaCaller) GetNode(opts *bind.CallOpts, hash [32]byte) (*big.Int, [32]byte, [32]byte, *big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
		ret4 = new([32]byte)
		ret5 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _Patricia.contract.Call(opts, out, "getNode", hash)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(uint256, bytes32, bytes32, uint256, bytes32, bytes32)
func (_Patricia *PatriciaSession) GetNode(hash [32]byte) (*big.Int, [32]byte, [32]byte, *big.Int, [32]byte, [32]byte, error) {
	return _Patricia.Contract.GetNode(&_Patricia.CallOpts, hash)
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(hash bytes32) constant returns(uint256, bytes32, bytes32, uint256, bytes32, bytes32)
func (_Patricia *PatriciaCallerSession) GetNode(hash [32]byte) (*big.Int, [32]byte, [32]byte, *big.Int, [32]byte, [32]byte, error) {
	return _Patricia.Contract.GetNode(&_Patricia.CallOpts, hash)
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_Patricia *PatriciaCaller) GetProof(opts *bind.CallOpts, key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	ret := new(struct {
		BranchMask *big.Int
		Siblings   [][32]byte
	})
	out := ret
	err := _Patricia.contract.Call(opts, out, "getProof", key)
	return *ret, err
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_Patricia *PatriciaSession) GetProof(key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	return _Patricia.Contract.GetProof(&_Patricia.CallOpts, key)
}

// GetProof is a free data retrieval call binding the contract method 0x693ac4fb.
//
// Solidity: function getProof(key bytes) constant returns(branchMask uint256, _siblings bytes32[])
func (_Patricia *PatriciaCallerSession) GetProof(key []byte) (struct {
	BranchMask *big.Int
	Siblings   [][32]byte
}, error) {
	return _Patricia.Contract.GetProof(&_Patricia.CallOpts, key)
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(uint256, bytes32, bytes32)
func (_Patricia *PatriciaCaller) GetRootEdge(opts *bind.CallOpts) (*big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Patricia.contract.Call(opts, out, "getRootEdge")
	return *ret0, *ret1, *ret2, err
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(uint256, bytes32, bytes32)
func (_Patricia *PatriciaSession) GetRootEdge() (*big.Int, [32]byte, [32]byte, error) {
	return _Patricia.Contract.GetRootEdge(&_Patricia.CallOpts)
}

// GetRootEdge is a free data retrieval call binding the contract method 0xa43914da.
//
// Solidity: function getRootEdge() constant returns(uint256, bytes32, bytes32)
func (_Patricia *PatriciaCallerSession) GetRootEdge() (*big.Int, [32]byte, [32]byte, error) {
	return _Patricia.Contract.GetRootEdge(&_Patricia.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(bytes32)
func (_Patricia *PatriciaCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Patricia.contract.Call(opts, out, "root")
	return *ret0, err
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(bytes32)
func (_Patricia *PatriciaSession) Root() ([32]byte, error) {
	return _Patricia.Contract.Root(&_Patricia.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(bytes32)
func (_Patricia *PatriciaCallerSession) Root() ([32]byte, error) {
	return _Patricia.Contract.Root(&_Patricia.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns()
func (_Patricia *PatriciaCaller) VerifyProof(opts *bind.CallOpts, rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) error {
	var ()
	out := &[]interface{}{}
	err := _Patricia.contract.Call(opts, out, "verifyProof", rootHash, key, value, branchMask, siblings)
	return err
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns()
func (_Patricia *PatriciaSession) VerifyProof(rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) error {
	return _Patricia.Contract.VerifyProof(&_Patricia.CallOpts, rootHash, key, value, branchMask, siblings)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf7e498f6.
//
// Solidity: function verifyProof(rootHash bytes32, key bytes, value bytes, branchMask uint256, siblings bytes32[]) constant returns()
func (_Patricia *PatriciaCallerSession) VerifyProof(rootHash [32]byte, key []byte, value []byte, branchMask *big.Int, siblings [][32]byte) error {
	return _Patricia.Contract.VerifyProof(&_Patricia.CallOpts, rootHash, key, value, branchMask, siblings)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_Patricia *PatriciaTransactor) Insert(opts *bind.TransactOpts, key []byte, value []byte) (*types.Transaction, error) {
	return _Patricia.contract.Transact(opts, "insert", key, value)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_Patricia *PatriciaSession) Insert(key []byte, value []byte) (*types.Transaction, error) {
	return _Patricia.Contract.Insert(&_Patricia.TransactOpts, key, value)
}

// Insert is a paid mutator transaction binding the contract method 0x20ba5b60.
//
// Solidity: function insert(key bytes, value bytes) returns()
func (_Patricia *PatriciaTransactorSession) Insert(key []byte, value []byte) (*types.Transaction, error) {
	return _Patricia.Contract.Insert(&_Patricia.TransactOpts, key, value)
}
