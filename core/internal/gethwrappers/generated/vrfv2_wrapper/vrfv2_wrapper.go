// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrfv2_wrapper

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"
)

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

var VRFV2WrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_linkEthFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_coordinator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"WrapperFulfillmentFailed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COORDINATOR\",\"outputs\":[{\"internalType\":\"contractExtendedVRFCoordinatorV2Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK_ETH_FEED\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBSCRIPTION_ID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"}],\"name\":\"calculateRequestPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_requestGasPriceWei\",\"type\":\"uint256\"}],\"name\":\"estimateRequestPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"wrapperGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"coordinatorGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"wrapperPremiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"maxNumWords\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_callbacks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"callbackAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"requestGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"requestWeiPerUnitLink\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_configured\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_disabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_wrapperGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_coordinatorGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_wrapperPremiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_maxNumWords\",\"type\":\"uint8\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162002467380380620024678339810160408190526200003591620002c2565b8033806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620001f9565b5050506001600160601b0319606091821b811660805284821b811660a05283821b811660c0529082901b1660e0526040805163288688f960e21b815290516000916001600160a01b0384169163a21a23e49160048082019260209290919082900301818787803b1580156200013457600080fd5b505af115801562000149573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200016f91906200030c565b60c081901b6001600160c01b03191661010052604051631cd0704360e21b81526001600160401b03821660048201523060248201529091506001600160a01b03831690637341c10c90604401600060405180830381600087803b158015620001d657600080fd5b505af1158015620001eb573d6000803e3d6000fd5b50505050505050506200033e565b6001600160a01b038116331415620002545760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0381168114620002bd57600080fd5b919050565b600080600060608486031215620002d857600080fd5b620002e384620002a5565b9250620002f360208501620002a5565b91506200030360408501620002a5565b90509250925092565b6000602082840312156200031f57600080fd5b81516001600160401b03811681146200033757600080fd5b9392505050565b60805160601c60a05160601c60c05160601c60e05160601c6101005160c01c61208f620003d8600039600081816101810152610bd501526000818161026e01528181610b9601528181610ec901528181610faa01526110450152600081816103da015261164f015260008181610205015281816109f50152818161119701526114970152600081816104e5015261054d015261208f6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80637fb5d19d116100d8578063ad1783611161008c578063f2fde38b11610066578063f2fde38b1461049e578063f3fef3a3146104b1578063fc2a88c3146104c457600080fd5b8063ad178361146103d5578063c15ce4d7146103fc578063c3f909d41461040f57600080fd5b8063a3907d71116100bd578063a3907d71146103a8578063a4c0ed36146103b0578063a608a1e1146103c357600080fd5b80637fb5d19d146103775780638da5cb5b1461038a57600080fd5b80633b2bcbf11161012f57806348baa1c51161011457806348baa1c5146102b157806357a8070a1461035257806379ba50971461036f57600080fd5b80633b2bcbf1146102695780634306d3541461029057600080fd5b80631b6b6d23116101605780631b6b6d23146102005780631fe543e31461024c5780632f2770db1461026157600080fd5b8063030932bb1461017c578063181f5a77146101c1575b600080fd5b6101a37f000000000000000000000000000000000000000000000000000000000000000081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b604080518082018252601281527f56524656325772617070657220312e302e300000000000000000000000000000602082015290516101b89190611e4d565b6102277f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b8565b61025f61025a366004611b7c565b6104cd565b005b61025f61058d565b6102277f000000000000000000000000000000000000000000000000000000000000000081565b6102a361029e366004611c6b565b6105c3565b6040519081526020016101b8565b6103166102bf366004611b63565b60086020526000908152604090208054600182015460029092015473ffffffffffffffffffffffffffffffffffffffff8216927401000000000000000000000000000000000000000090920463ffffffff16919084565b6040805173ffffffffffffffffffffffffffffffffffffffff909516855263ffffffff90931660208501529183015260608201526080016101b8565b60035461035f9060ff1681565b60405190151581526020016101b8565b61025f6106ca565b6102a3610385366004611cd3565b6107c7565b60005473ffffffffffffffffffffffffffffffffffffffff16610227565b61025f6108cd565b61025f6103be366004611a42565b6108ff565b60035461035f90610100900460ff1681565b6102277f000000000000000000000000000000000000000000000000000000000000000081565b61025f61040a366004611da7565b610d84565b6004546005546006546007546040805194855263ffffffff80851660208701526401000000008504811691860191909152680100000000000000008404811660608601526c01000000000000000000000000840416608085015260ff700100000000000000000000000000000000909304831660a085015260c08401919091521660e0820152610100016101b8565b61025f6104ac3660046119fd565b61112f565b61025f6104bf366004611a18565b611143565b6102a360025481565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461057f576040517f1cf993f400000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660248201526044015b60405180910390fd5b6105898282611218565b5050565b610595611579565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16610100179055565b60035460009060ff16610632576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f77726170706572206973206e6f7420636f6e66696775726564000000000000006044820152606401610576565b600354610100900460ff16156106a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f777261707065722069732064697361626c6564000000000000000000000000006044820152606401610576565b60006106ae6115fc565b90506106c18363ffffffff163a83611770565b9150505b919050565b60015473ffffffffffffffffffffffffffffffffffffffff16331461074b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610576565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60035460009060ff16610836576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f77726170706572206973206e6f7420636f6e66696775726564000000000000006044820152606401610576565b600354610100900460ff16156108a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f777261707065722069732064697361626c6564000000000000000000000000006044820152606401610576565b60006108b26115fc565b90506108c58463ffffffff168483611770565b949350505050565b6108d5611579565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055565b60035460ff1661096b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f77726170706572206973206e6f7420636f6e66696775726564000000000000006044820152606401610576565b600354610100900460ff16156109dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f777261707065722069732064697361626c6564000000000000000000000000006044820152606401610576565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610a7c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c792063616c6c61626c652066726f6d204c494e4b0000000000000000006044820152606401610576565b60008080610a8c84860186611c88565b9250925092506000610a9c6115fc565b90506000610ab18563ffffffff163a84611770565b905080881015610b1d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f66656520746f6f206c6f770000000000000000000000000000000000000000006044820152606401610576565b60075460ff1663ffffffff84161115610b92576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6e756d576f72647320746f6f20686967680000000000000000000000000000006044820152606401610576565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635d3b1d306006547f000000000000000000000000000000000000000000000000000000000000000088600560089054906101000a900463ffffffff168b610c139190611f26565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b168152600481019490945267ffffffffffffffff909216602484015261ffff16604483015263ffffffff90811660648301528716608482015260a401602060405180830381600087803b158015610c9257600080fd5b505af1158015610ca6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cca9190611aeb565b6040805160808101825273ffffffffffffffffffffffffffffffffffffffff9c8d16815263ffffffff98891660208083019182523a83850190815260608401988952600086815260089092529390209151825491519e167fffffffffffffffff00000000000000000000000000000000000000000000000090911617740100000000000000000000000000000000000000009d9099169c909c02979097178b55955160018b01555050516002978801555050909355505050565b610d8c611579565b6005805460ff808616700100000000000000000000000000000000027fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff63ffffffff8981166c01000000000000000000000000027fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff918c166801000000000000000002919091167fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff909516949094179390931792909216919091179091556006839055600780549183167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00928316179055600380549091166001179055604080517fc3f909d4000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163c3f909d4916004828101926080929190829003018186803b158015610f0f57600080fd5b505afa158015610f23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f479190611b04565b50600580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff929092169190911790555050604080517f356dac7100000000000000000000000000000000000000000000000000000000815290517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169163356dac71916004808301926020929190829003018186803b15801561100557600080fd5b505afa158015611019573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103d9190611aeb565b6004819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635fbbc0d26040518163ffffffff1660e01b81526004016101206040518083038186803b1580156110aa57600080fd5b505afa1580156110be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110e29190611cf1565b50506005805463ffffffff909816640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff909816979097179096555050505050505050505050565b611137611579565b61114081611859565b50565b61114b611579565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390527f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb90604401602060405180830381600087803b1580156111db57600080fd5b505af11580156111ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112139190611ac9565b505050565b60008281526008602081815260408084208151608081018352815473ffffffffffffffffffffffffffffffffffffffff808216835263ffffffff74010000000000000000000000000000000000000000830416838701526001840180549584019590955260028401805460608501528a8952969095527fffffffffffffffff000000000000000000000000000000000000000000000000169091559084905592909155815116611324576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f72657175657374206e6f7420666f756e640000000000000000000000000000006044820152606401610576565b600080631fe543e360e01b8585604051602401611342929190611ec0565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152905060005a905060006113e1856020015163ffffffff1686600001518561194f565b905060005a6113f09084611feb565b90508161143e57855160405173ffffffffffffffffffffffffffffffffffffffff9091169089907fc551b83c151f2d1c7eeb938ac59008e0409f1c1dc1e2f112449d4d79b458902290600090a35b600061145d876020015163ffffffff1688604001518960600151611770565b905060006114748389604001518a60600151611770565b90508082111561156d57875173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906114c88486611feb565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff90921660048301526024820152604401602060405180830381600087803b15801561153357600080fd5b505af1158015611547573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061156b9190611ac9565b505b50505050505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146115fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610576565b565b600554604080517ffeaf968c000000000000000000000000000000000000000000000000000000008152905160009263ffffffff161515918391829173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163feaf968c9160048082019260a092909190829003018186803b15801561169657600080fd5b505afa1580156116aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116ce9190611e09565b5094509092508491505080156116f457506116e98242611feb565b60055463ffffffff16105b156116fe57506004545b6000811215611769576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e76616c6964204c494e4b20776569207072696365000000000000000000006044820152606401610576565b9392505050565b6005546000908190839063ffffffff6c0100000000000000000000000082048116916117aa91680100000000000000009091041688611f0e565b6117b49190611f0e565b6117c686670de0b6b3a7640000611fae565b6117d09190611fae565b6117da9190611f73565b60055490915060009060649061180790700100000000000000000000000000000000900460ff1682611f4e565b6118149060ff1684611fae565b61181e9190611f73565b60055490915060009061184490640100000000900463ffffffff1664e8d4a51000611fae565b61184e9083611f0e565b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff81163314156118d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610576565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60005a61138881101561196157600080fd5b61138881039050846040820482031161197957600080fd5b50823b61198557600080fd5b60008083516020850160008789f1949350505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146106c557600080fd5b805162ffffff811681146106c557600080fd5b803560ff811681146106c557600080fd5b805169ffffffffffffffffffff811681146106c557600080fd5b600060208284031215611a0f57600080fd5b6117698261199b565b60008060408385031215611a2b57600080fd5b611a348361199b565b946020939093013593505050565b60008060008060608587031215611a5857600080fd5b611a618561199b565b935060208501359250604085013567ffffffffffffffff80821115611a8557600080fd5b818701915087601f830112611a9957600080fd5b813581811115611aa857600080fd5b886020828501011115611aba57600080fd5b95989497505060200194505050565b600060208284031215611adb57600080fd5b8151801515811461176957600080fd5b600060208284031215611afd57600080fd5b5051919050565b60008060008060808587031215611b1a57600080fd5b8451611b2581612060565b6020860151909450611b3681612070565b6040860151909350611b4781612070565b6060860151909250611b5881612070565b939692955090935050565b600060208284031215611b7557600080fd5b5035919050565b60008060408385031215611b8f57600080fd5b8235915060208084013567ffffffffffffffff80821115611baf57600080fd5b818601915086601f830112611bc357600080fd5b813581811115611bd557611bd5612031565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f83011681018181108582111715611c1857611c18612031565b604052828152858101935084860182860187018b1015611c3757600080fd5b600095505b83861015611c5a578035855260019590950194938601938601611c3c565b508096505050505050509250929050565b600060208284031215611c7d57600080fd5b813561176981612070565b600080600060608486031215611c9d57600080fd5b8335611ca881612070565b92506020840135611cb881612060565b91506040840135611cc881612070565b809150509250925092565b60008060408385031215611ce657600080fd5b8235611a3481612070565b60008060008060008060008060006101208a8c031215611d1057600080fd5b8951611d1b81612070565b60208b0151909950611d2c81612070565b60408b0151909850611d3d81612070565b60608b0151909750611d4e81612070565b60808b0151909650611d5f81612070565b9450611d6d60a08b016119bf565b9350611d7b60c08b016119bf565b9250611d8960e08b016119bf565b9150611d986101008b016119bf565b90509295985092959850929598565b600080600080600060a08688031215611dbf57600080fd5b8535611dca81612070565b94506020860135611dda81612070565b9350611de8604087016119d2565b925060608601359150611dfd608087016119d2565b90509295509295909350565b600080600080600060a08688031215611e2157600080fd5b611e2a866119e3565b9450602086015193506040860151925060608601519150611dfd608087016119e3565b600060208083528351808285015260005b81811015611e7a57858101830151858201604001528201611e5e565b81811115611e8c576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b6000604082018483526020604081850152818551808452606086019150828701935060005b81811015611f0157845183529383019391830191600101611ee5565b5090979650505050505050565b60008219821115611f2157611f21612002565b500190565b600063ffffffff808316818516808303821115611f4557611f45612002565b01949350505050565b600060ff821660ff84168060ff03821115611f6b57611f6b612002565b019392505050565b600082611fa9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611fe657611fe6612002565b500290565b600082821015611ffd57611ffd612002565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61ffff8116811461114057600080fd5b63ffffffff8116811461114057600080fdfea164736f6c6343000806000a",
}

var VRFV2WrapperABI = VRFV2WrapperMetaData.ABI

var VRFV2WrapperBin = VRFV2WrapperMetaData.Bin

func DeployVRFV2Wrapper(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address, _linkEthFeed common.Address, _coordinator common.Address) (common.Address, *types.Transaction, *VRFV2Wrapper, error) {
	parsed, err := VRFV2WrapperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFV2WrapperBin), backend, _link, _linkEthFeed, _coordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFV2Wrapper{VRFV2WrapperCaller: VRFV2WrapperCaller{contract: contract}, VRFV2WrapperTransactor: VRFV2WrapperTransactor{contract: contract}, VRFV2WrapperFilterer: VRFV2WrapperFilterer{contract: contract}}, nil
}

type VRFV2Wrapper struct {
	address common.Address
	abi     abi.ABI
	VRFV2WrapperCaller
	VRFV2WrapperTransactor
	VRFV2WrapperFilterer
}

type VRFV2WrapperCaller struct {
	contract *bind.BoundContract
}

type VRFV2WrapperTransactor struct {
	contract *bind.BoundContract
}

type VRFV2WrapperFilterer struct {
	contract *bind.BoundContract
}

type VRFV2WrapperSession struct {
	Contract     *VRFV2Wrapper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFV2WrapperCallerSession struct {
	Contract *VRFV2WrapperCaller
	CallOpts bind.CallOpts
}

type VRFV2WrapperTransactorSession struct {
	Contract     *VRFV2WrapperTransactor
	TransactOpts bind.TransactOpts
}

type VRFV2WrapperRaw struct {
	Contract *VRFV2Wrapper
}

type VRFV2WrapperCallerRaw struct {
	Contract *VRFV2WrapperCaller
}

type VRFV2WrapperTransactorRaw struct {
	Contract *VRFV2WrapperTransactor
}

func NewVRFV2Wrapper(address common.Address, backend bind.ContractBackend) (*VRFV2Wrapper, error) {
	abi, err := abi.JSON(strings.NewReader(VRFV2WrapperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFV2Wrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFV2Wrapper{address: address, abi: abi, VRFV2WrapperCaller: VRFV2WrapperCaller{contract: contract}, VRFV2WrapperTransactor: VRFV2WrapperTransactor{contract: contract}, VRFV2WrapperFilterer: VRFV2WrapperFilterer{contract: contract}}, nil
}

func NewVRFV2WrapperCaller(address common.Address, caller bind.ContractCaller) (*VRFV2WrapperCaller, error) {
	contract, err := bindVRFV2Wrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperCaller{contract: contract}, nil
}

func NewVRFV2WrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFV2WrapperTransactor, error) {
	contract, err := bindVRFV2Wrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperTransactor{contract: contract}, nil
}

func NewVRFV2WrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFV2WrapperFilterer, error) {
	contract, err := bindVRFV2Wrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperFilterer{contract: contract}, nil
}

func bindVRFV2Wrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFV2WrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFV2Wrapper *VRFV2WrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV2Wrapper.Contract.VRFV2WrapperCaller.contract.Call(opts, result, method, params...)
}

func (_VRFV2Wrapper *VRFV2WrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.VRFV2WrapperTransactor.contract.Transfer(opts)
}

func (_VRFV2Wrapper *VRFV2WrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.VRFV2WrapperTransactor.contract.Transact(opts, method, params...)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV2Wrapper.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.contract.Transfer(opts)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.contract.Transact(opts, method, params...)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) COORDINATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "COORDINATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) COORDINATOR() (common.Address, error) {
	return _VRFV2Wrapper.Contract.COORDINATOR(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) COORDINATOR() (common.Address, error) {
	return _VRFV2Wrapper.Contract.COORDINATOR(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) LINK() (common.Address, error) {
	return _VRFV2Wrapper.Contract.LINK(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) LINK() (common.Address, error) {
	return _VRFV2Wrapper.Contract.LINK(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) LINKETHFEED(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "LINK_ETH_FEED")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) LINKETHFEED() (common.Address, error) {
	return _VRFV2Wrapper.Contract.LINKETHFEED(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) LINKETHFEED() (common.Address, error) {
	return _VRFV2Wrapper.Contract.LINKETHFEED(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) SUBSCRIPTIONID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "SUBSCRIPTION_ID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) SUBSCRIPTIONID() (uint64, error) {
	return _VRFV2Wrapper.Contract.SUBSCRIPTIONID(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) SUBSCRIPTIONID() (uint64, error) {
	return _VRFV2Wrapper.Contract.SUBSCRIPTIONID(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) CalculateRequestPrice(opts *bind.CallOpts, _callbackGasLimit uint32) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "calculateRequestPrice", _callbackGasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) CalculateRequestPrice(_callbackGasLimit uint32) (*big.Int, error) {
	return _VRFV2Wrapper.Contract.CalculateRequestPrice(&_VRFV2Wrapper.CallOpts, _callbackGasLimit)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) CalculateRequestPrice(_callbackGasLimit uint32) (*big.Int, error) {
	return _VRFV2Wrapper.Contract.CalculateRequestPrice(&_VRFV2Wrapper.CallOpts, _callbackGasLimit)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) EstimateRequestPrice(opts *bind.CallOpts, _callbackGasLimit uint32, _requestGasPriceWei *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "estimateRequestPrice", _callbackGasLimit, _requestGasPriceWei)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) EstimateRequestPrice(_callbackGasLimit uint32, _requestGasPriceWei *big.Int) (*big.Int, error) {
	return _VRFV2Wrapper.Contract.EstimateRequestPrice(&_VRFV2Wrapper.CallOpts, _callbackGasLimit, _requestGasPriceWei)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) EstimateRequestPrice(_callbackGasLimit uint32, _requestGasPriceWei *big.Int) (*big.Int, error) {
	return _VRFV2Wrapper.Contract.EstimateRequestPrice(&_VRFV2Wrapper.CallOpts, _callbackGasLimit, _requestGasPriceWei)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) GetConfig(opts *bind.CallOpts) (GetConfig,

	error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "getConfig")

	outstruct := new(GetConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FallbackWeiPerUnitLink = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StalenessSeconds = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeLinkPPM = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.WrapperGasOverhead = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.CoordinatorGasOverhead = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.WrapperPremiumPercentage = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.KeyHash = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.MaxNumWords = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) GetConfig() (GetConfig,

	error) {
	return _VRFV2Wrapper.Contract.GetConfig(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) GetConfig() (GetConfig,

	error) {
	return _VRFV2Wrapper.Contract.GetConfig(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) LastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) LastRequestId() (*big.Int, error) {
	return _VRFV2Wrapper.Contract.LastRequestId(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) LastRequestId() (*big.Int, error) {
	return _VRFV2Wrapper.Contract.LastRequestId(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) Owner() (common.Address, error) {
	return _VRFV2Wrapper.Contract.Owner(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) Owner() (common.Address, error) {
	return _VRFV2Wrapper.Contract.Owner(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) SCallbacks(opts *bind.CallOpts, arg0 *big.Int) (SCallbacks,

	error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "s_callbacks", arg0)

	outstruct := new(SCallbacks)
	if err != nil {
		return *outstruct, err
	}

	outstruct.CallbackAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CallbackGasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.RequestGasPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RequestWeiPerUnitLink = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) SCallbacks(arg0 *big.Int) (SCallbacks,

	error) {
	return _VRFV2Wrapper.Contract.SCallbacks(&_VRFV2Wrapper.CallOpts, arg0)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) SCallbacks(arg0 *big.Int) (SCallbacks,

	error) {
	return _VRFV2Wrapper.Contract.SCallbacks(&_VRFV2Wrapper.CallOpts, arg0)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) SConfigured(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "s_configured")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) SConfigured() (bool, error) {
	return _VRFV2Wrapper.Contract.SConfigured(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) SConfigured() (bool, error) {
	return _VRFV2Wrapper.Contract.SConfigured(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) SDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "s_disabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) SDisabled() (bool, error) {
	return _VRFV2Wrapper.Contract.SDisabled(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) SDisabled() (bool, error) {
	return _VRFV2Wrapper.Contract.SDisabled(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VRFV2Wrapper.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRFV2Wrapper *VRFV2WrapperSession) TypeAndVersion() (string, error) {
	return _VRFV2Wrapper.Contract.TypeAndVersion(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperCallerSession) TypeAndVersion() (string, error) {
	return _VRFV2Wrapper.Contract.TypeAndVersion(&_VRFV2Wrapper.CallOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2Wrapper.contract.Transact(opts, "acceptOwnership")
}

func (_VRFV2Wrapper *VRFV2WrapperSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.AcceptOwnership(&_VRFV2Wrapper.TransactOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.AcceptOwnership(&_VRFV2Wrapper.TransactOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactor) Disable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2Wrapper.contract.Transact(opts, "disable")
}

func (_VRFV2Wrapper *VRFV2WrapperSession) Disable() (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.Disable(&_VRFV2Wrapper.TransactOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactorSession) Disable() (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.Disable(&_VRFV2Wrapper.TransactOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactor) Enable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2Wrapper.contract.Transact(opts, "enable")
}

func (_VRFV2Wrapper *VRFV2WrapperSession) Enable() (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.Enable(&_VRFV2Wrapper.TransactOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactorSession) Enable() (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.Enable(&_VRFV2Wrapper.TransactOpts)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactor) OnTokenTransfer(opts *bind.TransactOpts, _sender common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _VRFV2Wrapper.contract.Transact(opts, "onTokenTransfer", _sender, _amount, _data)
}

func (_VRFV2Wrapper *VRFV2WrapperSession) OnTokenTransfer(_sender common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.OnTokenTransfer(&_VRFV2Wrapper.TransactOpts, _sender, _amount, _data)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactorSession) OnTokenTransfer(_sender common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.OnTokenTransfer(&_VRFV2Wrapper.TransactOpts, _sender, _amount, _data)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2Wrapper.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

func (_VRFV2Wrapper *VRFV2WrapperSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.RawFulfillRandomWords(&_VRFV2Wrapper.TransactOpts, requestId, randomWords)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.RawFulfillRandomWords(&_VRFV2Wrapper.TransactOpts, requestId, randomWords)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactor) SetConfig(opts *bind.TransactOpts, _wrapperGasOverhead uint32, _coordinatorGasOverhead uint32, _wrapperPremiumPercentage uint8, _keyHash [32]byte, _maxNumWords uint8) (*types.Transaction, error) {
	return _VRFV2Wrapper.contract.Transact(opts, "setConfig", _wrapperGasOverhead, _coordinatorGasOverhead, _wrapperPremiumPercentage, _keyHash, _maxNumWords)
}

func (_VRFV2Wrapper *VRFV2WrapperSession) SetConfig(_wrapperGasOverhead uint32, _coordinatorGasOverhead uint32, _wrapperPremiumPercentage uint8, _keyHash [32]byte, _maxNumWords uint8) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.SetConfig(&_VRFV2Wrapper.TransactOpts, _wrapperGasOverhead, _coordinatorGasOverhead, _wrapperPremiumPercentage, _keyHash, _maxNumWords)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactorSession) SetConfig(_wrapperGasOverhead uint32, _coordinatorGasOverhead uint32, _wrapperPremiumPercentage uint8, _keyHash [32]byte, _maxNumWords uint8) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.SetConfig(&_VRFV2Wrapper.TransactOpts, _wrapperGasOverhead, _coordinatorGasOverhead, _wrapperPremiumPercentage, _keyHash, _maxNumWords)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFV2Wrapper.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFV2Wrapper *VRFV2WrapperSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.TransferOwnership(&_VRFV2Wrapper.TransactOpts, to)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.TransferOwnership(&_VRFV2Wrapper.TransactOpts, to)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactor) Withdraw(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VRFV2Wrapper.contract.Transact(opts, "withdraw", _recipient, _amount)
}

func (_VRFV2Wrapper *VRFV2WrapperSession) Withdraw(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.Withdraw(&_VRFV2Wrapper.TransactOpts, _recipient, _amount)
}

func (_VRFV2Wrapper *VRFV2WrapperTransactorSession) Withdraw(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VRFV2Wrapper.Contract.Withdraw(&_VRFV2Wrapper.TransactOpts, _recipient, _amount)
}

type VRFV2WrapperOwnershipTransferRequestedIterator struct {
	Event *VRFV2WrapperOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2WrapperOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2WrapperOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFV2WrapperOwnershipTransferRequested)
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

func (it *VRFV2WrapperOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFV2WrapperOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2WrapperOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV2Wrapper *VRFV2WrapperFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2WrapperOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2Wrapper.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperOwnershipTransferRequestedIterator{contract: _VRFV2Wrapper.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFV2Wrapper *VRFV2WrapperFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2Wrapper.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2WrapperOwnershipTransferRequested)
				if err := _VRFV2Wrapper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFV2Wrapper *VRFV2WrapperFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFV2WrapperOwnershipTransferRequested, error) {
	event := new(VRFV2WrapperOwnershipTransferRequested)
	if err := _VRFV2Wrapper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV2WrapperOwnershipTransferredIterator struct {
	Event *VRFV2WrapperOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2WrapperOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2WrapperOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFV2WrapperOwnershipTransferred)
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

func (it *VRFV2WrapperOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFV2WrapperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2WrapperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV2Wrapper *VRFV2WrapperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2WrapperOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2Wrapper.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperOwnershipTransferredIterator{contract: _VRFV2Wrapper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFV2Wrapper *VRFV2WrapperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2Wrapper.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2WrapperOwnershipTransferred)
				if err := _VRFV2Wrapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFV2Wrapper *VRFV2WrapperFilterer) ParseOwnershipTransferred(log types.Log) (*VRFV2WrapperOwnershipTransferred, error) {
	event := new(VRFV2WrapperOwnershipTransferred)
	if err := _VRFV2Wrapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV2WrapperWrapperFulfillmentFailedIterator struct {
	Event *VRFV2WrapperWrapperFulfillmentFailed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2WrapperWrapperFulfillmentFailedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2WrapperWrapperFulfillmentFailed)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFV2WrapperWrapperFulfillmentFailed)
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

func (it *VRFV2WrapperWrapperFulfillmentFailedIterator) Error() error {
	return it.fail
}

func (it *VRFV2WrapperWrapperFulfillmentFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2WrapperWrapperFulfillmentFailed struct {
	RequestId *big.Int
	Consumer  common.Address
	Raw       types.Log
}

func (_VRFV2Wrapper *VRFV2WrapperFilterer) FilterWrapperFulfillmentFailed(opts *bind.FilterOpts, requestId []*big.Int, consumer []common.Address) (*VRFV2WrapperWrapperFulfillmentFailedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _VRFV2Wrapper.contract.FilterLogs(opts, "WrapperFulfillmentFailed", requestIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return &VRFV2WrapperWrapperFulfillmentFailedIterator{contract: _VRFV2Wrapper.contract, event: "WrapperFulfillmentFailed", logs: logs, sub: sub}, nil
}

func (_VRFV2Wrapper *VRFV2WrapperFilterer) WatchWrapperFulfillmentFailed(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperWrapperFulfillmentFailed, requestId []*big.Int, consumer []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _VRFV2Wrapper.contract.WatchLogs(opts, "WrapperFulfillmentFailed", requestIdRule, consumerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2WrapperWrapperFulfillmentFailed)
				if err := _VRFV2Wrapper.contract.UnpackLog(event, "WrapperFulfillmentFailed", log); err != nil {
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

func (_VRFV2Wrapper *VRFV2WrapperFilterer) ParseWrapperFulfillmentFailed(log types.Log) (*VRFV2WrapperWrapperFulfillmentFailed, error) {
	event := new(VRFV2WrapperWrapperFulfillmentFailed)
	if err := _VRFV2Wrapper.contract.UnpackLog(event, "WrapperFulfillmentFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetConfig struct {
	FallbackWeiPerUnitLink    *big.Int
	StalenessSeconds          uint32
	FulfillmentFlatFeeLinkPPM uint32
	WrapperGasOverhead        uint32
	CoordinatorGasOverhead    uint32
	WrapperPremiumPercentage  uint8
	KeyHash                   [32]byte
	MaxNumWords               uint8
}
type SCallbacks struct {
	CallbackAddress       common.Address
	CallbackGasLimit      uint32
	RequestGasPrice       *big.Int
	RequestWeiPerUnitLink *big.Int
}

func (_VRFV2Wrapper *VRFV2Wrapper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFV2Wrapper.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFV2Wrapper.ParseOwnershipTransferRequested(log)
	case _VRFV2Wrapper.abi.Events["OwnershipTransferred"].ID:
		return _VRFV2Wrapper.ParseOwnershipTransferred(log)
	case _VRFV2Wrapper.abi.Events["WrapperFulfillmentFailed"].ID:
		return _VRFV2Wrapper.ParseWrapperFulfillmentFailed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFV2WrapperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFV2WrapperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VRFV2WrapperWrapperFulfillmentFailed) Topic() common.Hash {
	return common.HexToHash("0xc551b83c151f2d1c7eeb938ac59008e0409f1c1dc1e2f112449d4d79b4589022")
}

func (_VRFV2Wrapper *VRFV2Wrapper) Address() common.Address {
	return _VRFV2Wrapper.address
}

type VRFV2WrapperInterface interface {
	COORDINATOR(opts *bind.CallOpts) (common.Address, error)

	LINK(opts *bind.CallOpts) (common.Address, error)

	LINKETHFEED(opts *bind.CallOpts) (common.Address, error)

	SUBSCRIPTIONID(opts *bind.CallOpts) (uint64, error)

	CalculateRequestPrice(opts *bind.CallOpts, _callbackGasLimit uint32) (*big.Int, error)

	EstimateRequestPrice(opts *bind.CallOpts, _callbackGasLimit uint32, _requestGasPriceWei *big.Int) (*big.Int, error)

	GetConfig(opts *bind.CallOpts) (GetConfig,

		error)

	LastRequestId(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SCallbacks(opts *bind.CallOpts, arg0 *big.Int) (SCallbacks,

		error)

	SConfigured(opts *bind.CallOpts) (bool, error)

	SDisabled(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Disable(opts *bind.TransactOpts) (*types.Transaction, error)

	Enable(opts *bind.TransactOpts) (*types.Transaction, error)

	OnTokenTransfer(opts *bind.TransactOpts, _sender common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error)

	RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _wrapperGasOverhead uint32, _coordinatorGasOverhead uint32, _wrapperPremiumPercentage uint8, _keyHash [32]byte, _maxNumWords uint8) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2WrapperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFV2WrapperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2WrapperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFV2WrapperOwnershipTransferred, error)

	FilterWrapperFulfillmentFailed(opts *bind.FilterOpts, requestId []*big.Int, consumer []common.Address) (*VRFV2WrapperWrapperFulfillmentFailedIterator, error)

	WatchWrapperFulfillmentFailed(opts *bind.WatchOpts, sink chan<- *VRFV2WrapperWrapperFulfillmentFailed, requestId []*big.Int, consumer []common.Address) (event.Subscription, error)

	ParseWrapperFulfillmentFailed(log types.Log) (*VRFV2WrapperWrapperFulfillmentFailed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
