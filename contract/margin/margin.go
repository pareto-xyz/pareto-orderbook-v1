// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package margin

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

// DerivativeOption is an auto generated low-level Go binding around an user-defined struct.
type DerivativeOption struct {
	IsCall      bool
	StrikeLevel uint8
	Strike      *big.Int
	Expiry      *big.Int
	Underlying  uint8
	Decimals    uint8
}

// DerivativeOrder is an auto generated low-level Go binding around an user-defined struct.
type DerivativeOrder struct {
	Id         string
	Buyer      common.Address
	Seller     common.Address
	TradePrice *big.Int
	Quantity   *big.Int
	Option     DerivativeOption
}

// DerivativePositionParams is an auto generated low-level Go binding around an user-defined struct.
type DerivativePositionParams struct {
	Id            string
	Buyer         common.Address
	Seller        common.Address
	TradePrice    *big.Int
	Quantity      *big.Int
	IsCall        bool
	StrikeLevel   uint8
	Underlying    uint8
	IsBuyerMaker  bool
	IsSellerMaker bool
}

// MarginMetaData contains all meta data concerning the Margin contract.
var MarginMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bankruptUser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"counterparty\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bankruptcyAmount\",\"type\":\"uint256\"}],\"name\":\"BankruptcyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxBalance\",\"type\":\"uint256\"}],\"name\":\"MaxBalanceCapEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perc\",\"type\":\"uint256\"}],\"name\":\"MinMarginPercEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumDerivative.Underlying\",\"name\":\"underlying\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minQuantity\",\"type\":\"uint256\"}],\"name\":\"NewUnderlyingEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tradePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"enumDerivative.Underlying\",\"name\":\"underlying\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDerivative.StrikeLevel\",\"name\":\"strikeLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"RecordPositionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numPositions\",\"type\":\"uint256\"}],\"name\":\"SettlementEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"TogglePauseEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"discounted\",\"type\":\"uint256\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"enumDerivative.Underlying\",\"name\":\"underlying\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minQuantity\",\"type\":\"uint256\"}],\"name\":\"activateUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeExpiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"addKeepers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tradePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"},{\"internalType\":\"enumDerivative.StrikeLevel\",\"name\":\"strikeLevel\",\"type\":\"uint8\"},{\"internalType\":\"enumDerivative.Underlying\",\"name\":\"underlying\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isBuyerMaker\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSellerMaker\",\"type\":\"bool\"}],\"internalType\":\"structDerivative.PositionParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"addPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useInitialMargin\",\"type\":\"bool\"}],\"name\":\"checkMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"useInitialMargin\",\"type\":\"bool\"}],\"name\":\"checkMarginBatch\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"diffs\",\"type\":\"int256[]\"},{\"internalType\":\"bool[]\",\"name\":\"satisfieds\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curRound\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useInitialMargin\",\"type\":\"bool\"}],\"name\":\"getMargin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"onlyLoss\",\"type\":\"bool\"}],\"name\":\"getPayoff\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDerivative.Underlying\",\"name\":\"underlying\",\"type\":\"uint8\"}],\"name\":\"getPositions\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tradePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"},{\"internalType\":\"enumDerivative.StrikeLevel\",\"name\":\"strikeLevel\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"enumDerivative.Underlying\",\"name\":\"underlying\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structDerivative.Option\",\"name\":\"option\",\"type\":\"tuple\"}],\"internalType\":\"structDerivative.Order[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDerivative.Underlying\",\"name\":\"underlying\",\"type\":\"uint8\"}],\"name\":\"getStrikes\",\"outputs\":[{\"internalType\":\"uint256[11]\",\"name\":\"\",\"type\":\"uint256[11]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdc_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"insurance_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient_\",\"type\":\"address\"},{\"internalType\":\"enumDerivative.Underlying\",\"name\":\"underlying_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"oracle_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minQuantity_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insurance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDerivative.Underlying\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"isActiveUnderlying\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBalanceCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minMarginPerc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDerivative.Underlying\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"minQuantityPerUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"removeKeepers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"roundUsers\",\"type\":\"address[]\"}],\"name\":\"rollover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDerivative.Underlying\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roundStrikes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newInsurance\",\"type\":\"address\"}],\"name\":\"setInsurance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxBalance\",\"type\":\"uint256\"}],\"name\":\"setMaxBalanceCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perc\",\"type\":\"uint256\"}],\"name\":\"setMinMarginPerc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDerivative.Underlying\",\"name\":\"underlying\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minQuantity\",\"type\":\"uint256\"}],\"name\":\"setMinQuantity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDerivative.Underlying\",\"name\":\"underlying\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff168152503480156200004457600080fd5b5060805161d65f6200007d6000396000818161313e015281816131cc01528181613312015281816133a00152613450015261d65f6000f3fe6080604052600436106102725760003560e01c80637a594a7c1161014f578063b40751cc116100c1578063e74b981b1161007a578063e74b981b1461092d578063ea7e9daa14610956578063f2fde38b1461097f578063f555a419146109a8578063fe6bb409146109d1578063ff7f29da146109fa57610272565b8063b40751cc1461081d578063b6b55f2514610848578063bd6b798214610871578063c4ae31681461089c578063c955bb69146108b3578063cc4fdcd0146108f057610272565b80638da5cb5b116101135780638da5cb5b146106f85780639b96eece14610723578063a1dc656b14610760578063a665a63714610789578063a9520b00146107c7578063b3f75843146107f257610272565b80637a594a7c146106275780637f64978314610664578063846c260f1461068d578063853828b6146106b657806389cf3204146106cd57610272565b80632f865568116101e85780634f1ef286116101ac5780634f1ef2861461054e57806352d1902d1461056a578063548db174146105955780635c563dbe146105be5780636cd7751f146105e7578063715018a61461061057610272565b80632f8655681461047d578063314a83c9146104a65780633659cfe6146104cf5780633e413bee146104f8578063469048401461052357610272565b806312065fe01161023a57806312065fe0146103485780631b784a47146103735780631febd86d1461039c57806323880e78146103da5780632db0788e146104175780632e1a7d4d1461045457610272565b806301b4a0431461027757806303c4d3d6146102a057806305c4d1cb146102cb578063062c4878146102f457806311da60b414610331575b600080fd5b34801561028357600080fd5b5061029e600480360381019061029991906197cd565b610a37565b005b3480156102ac57600080fd5b506102b5610dfe565b6040516102c29190619869565b60405180910390f35b3480156102d757600080fd5b506102f260048036038101906102ed91906198e9565b610e05565b005b34801561030057600080fd5b5061031b60048036038101906103169190619936565b6112ae565b6040516103289190619cce565b60405180910390f35b34801561033d57600080fd5b50610346611797565b005b34801561035457600080fd5b5061035d612026565b60405161036a9190619869565b60405180910390f35b34801561037f57600080fd5b5061039a60048036038101906103959190619cf0565b61206e565b005b3480156103a857600080fd5b506103c360048036038101906103be9190619d5c565b612176565b6040516103d1929190619dc4565b60405180910390f35b3480156103e657600080fd5b5061040160048036038101906103fc9190619d5c565b612210565b60405161040e9190619ded565b60405180910390f35b34801561042357600080fd5b5061043e60048036038101906104399190619936565b612749565b60405161044b9190619e08565b60405180910390f35b34801561046057600080fd5b5061047b60048036038101906104769190619e23565b61276a565b005b34801561048957600080fd5b506104a4600480360381019061049f9190619e50565b612a22565b005b3480156104b257600080fd5b506104cd60048036038101906104c891906198e9565b612fd9565b005b3480156104db57600080fd5b506104f660048036038101906104f19190619e50565b61313c565b005b34801561050457600080fd5b5061050d6132c4565b60405161051a9190619e8c565b60405180910390f35b34801561052f57600080fd5b506105386132ea565b6040516105459190619e8c565b60405180910390f35b61056860048036038101906105639190619fd7565b613310565b005b34801561057657600080fd5b5061057f61344c565b60405161058c919061a04c565b60405180910390f35b3480156105a157600080fd5b506105bc60048036038101906105b791906198e9565b613505565b005b3480156105ca57600080fd5b506105e560048036038101906105e0919061a067565b613667565b005b3480156105f357600080fd5b5061060e60048036038101906106099190619e50565b613777565b005b34801561061c57600080fd5b50610625613853565b005b34801561063357600080fd5b5061064e60048036038101906106499190619d5c565b613867565b60405161065b9190619869565b60405180910390f35b34801561067057600080fd5b5061068b600480360381019061068691906198e9565b6142f0565b005b34801561069957600080fd5b506106b460048036038101906106af919061a0cc565b614453565b005b3480156106c257600080fd5b506106cb615522565b005b3480156106d957600080fd5b506106e261578b565b6040516106ef9190619e8c565b60405180910390f35b34801561070457600080fd5b5061070d6157b1565b60405161071a9190619e8c565b60405180910390f35b34801561072f57600080fd5b5061074a60048036038101906107459190619e50565b6157db565b6040516107579190619869565b60405180910390f35b34801561076c57600080fd5b5061078760048036038101906107829190619e23565b615825565b005b34801561079557600080fd5b506107b060048036038101906107ab919061a115565b6158cb565b6040516107be92919061a2e2565b60405180910390f35b3480156107d357600080fd5b506107dc615978565b6040516107e9919061a328565b60405180910390f35b3480156107fe57600080fd5b50610807615a10565b604051610814919061a328565b60405180910390f35b34801561082957600080fd5b50610832615a23565b60405161083f9190619869565b60405180910390f35b34801561085457600080fd5b5061086f600480360381019061086a9190619e23565b615a29565b005b34801561087d57600080fd5b50610886615cb8565b6040516108939190619869565b60405180910390f35b3480156108a857600080fd5b506108b1615cbf565b005b3480156108bf57600080fd5b506108da60048036038101906108d59190619936565b615d53565b6040516108e79190619869565b60405180910390f35b3480156108fc57600080fd5b506109176004803603810190610912919061a067565b615d6b565b6040516109249190619869565b60405180910390f35b34801561093957600080fd5b50610954600480360381019061094f9190619e50565b615d94565b005b34801561096257600080fd5b5061097d600480360381019061097891906198e9565b615e70565b005b34801561098b57600080fd5b506109a660048036038101906109a19190619e50565b615fd2565b005b3480156109b457600080fd5b506109cf60048036038101906109ca9190619e23565b616055565b005b3480156109dd57600080fd5b506109f860048036038101906109f3919061a343565b6160f8565b005b348015610a0657600080fd5b50610a216004803603810190610a1c9190619936565b6161e8565b604051610a2e919061a432565b60405180910390f35b60008060019054906101000a900460ff16159050808015610a685750600160008054906101000a900460ff1660ff16105b80610a955750610a773061626b565b158015610a945750600160008054906101000a900460ff1660ff16145b5b610ad4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610acb9061a4d1565b60405180910390fd5b60016000806101000a81548160ff021916908360ff1602179055508015610b11576001600060016101000a81548160ff0219169083151502179055505b600073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1603610b4a57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1603610b8357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1603610bbc57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610bf557600080fd5b60008211610c0257600080fd5b8660fb60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508560fc60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508460fd60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610ccd61628e565b610cd56162e7565b60016101086000610ce46157b1565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160fd60146101000a81548160ff021916908360ff160217905550606461010081905550610d62615978565b600a610d6e919061a653565b6107d0610d7b919061a69e565b60fe81905550610d8a42616340565b61010181905550610d9c848484616379565b8015610df55760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610dec919061a725565b60405180910390a15b50505050505050565b6101015481565b600260655403610e4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e419061a78c565b60405180910390fd5b600260658190555061010860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610edf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed69061a81e565b60405180910390fd5b61010360009054906101000a900460ff1615610f30576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f279061a88a565b60405180910390fd5b426101015410610f75576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6c9061a8f6565b60405180910390fd5b61010360019054906101000a900460ff16610fc5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fbc9061a988565b60405180910390fd5b6000610101549050610fd681616340565b61010181905550600160fd60148282829054906101000a900460ff16610ffc919061a9a8565b92506101000a81548160ff021916908360ff160217905550600061010360016101000a81548160ff02191690831515021790555060005b600460ff1681101561110657600081600381111561105457611053619a58565b5b9050610107600082600381111561106e5761106d619a58565b5b60038111156110805761107f619a58565b5b815260200190815260200160002060009054906101000a900460ff16156110f2576110aa8161659b565b61010e60008360038111156110c2576110c1619a58565b5b60038111156110d4576110d3619a58565b5b815260200190815260200160002090600b6110f09291906193f6565b505b5080806110fe9061a9dd565b915050611033565b5061010a60006111169190619436565b60005b838390508110156112a05761010b600085858481811061113c5761113b61aa25565b5b90506020020160208101906111519190619e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000611197919061945a565b61010c60008585848181106111af576111ae61aa25565b5b90506020020160208101906111c49190619e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061120a9190619482565b600061010d60008686858181106112245761122361aa25565b5b90506020020160208101906112399190619e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548161ffff021916908361ffff16021790555080806112989061a9dd565b915050611119565b505060016065819055505050565b6060600061010d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900461ffff1661ffff1667ffffffffffffffff81111561131e5761131d619eac565b5b60405190808252806020026020018201604052801561135757816020015b6113446194aa565b81526020019060019003908161133c5790505b5090506000805b61010b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905081101561178c5761010c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081815481106113fb576113fa61aa25565b5b90600052602060002090602091828204019190069054906101000a900460ff161561177957600061010a61010b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083815481106114775761147661aa25565b5b90600052602060002090601091828204019190066002029054906101000a900461ffff1661ffff16815481106114b0576114af61aa25565b5b9060005260206000209060090201905060008160040154036114d25750611779565b8560038111156114e5576114e4619a58565b5b8160050160030160009054906101000a900460ff16600381111561150c5761150b619a58565b5b146115175750611779565b806040518060c00160405290816000820180546115339061aa83565b80601f016020809104026020016040519081016040528092919081815260200182805461155f9061aa83565b80156115ac5780601f10611581576101008083540402835291602001916115ac565b820191906000526020600020905b81548152906001019060200180831161158f57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582016040518060c00160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff16600a8111156116c5576116c4619a58565b5b600a8111156116d7576116d6619a58565b5b815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff16600381111561171457611713619a58565b5b600381111561172657611725619a58565b5b81526020016003820160019054906101000a900460ff1660ff1660ff16815250508152505084848151811061175e5761175d61aa25565b5b602002602001018190525082806117749061a9dd565b935050505b80806117849061a9dd565b91505061135e565b508192505050919050565b6002606554036117dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117d39061a78c565b60405180910390fd5b60026065819055504261010154111561182a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118219061ab26565b60405180910390fd5b61010360019054906101000a900460ff161561187b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118729061abb8565b60405180910390fd5b60005b61010a80549050811015611f9a57600061010a82815481106118a3576118a261aa25565b5b90600052602060002090600902016040518060c00160405290816000820180546118cc9061aa83565b80601f01602080910402602001604051908101604052809291908181526020018280546118f89061aa83565b80156119455780601f1061191a57610100808354040283529160200191611945565b820191906000526020600020905b81548152906001019060200180831161192857829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582016040518060c00160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff16600a811115611a5e57611a5d619a58565b5b600a811115611a7057611a6f619a58565b5b815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff166003811115611aad57611aac619a58565b5b6003811115611abf57611abe619a58565b5b81526020016003820160019054906101000a900460ff1660ff1660ff16815250508152505090506000816080015103611af85750611f87565b6000611b0b8260a0015160800151616c2b565b90506000611b1e83602001518385616dca565b90506000836060015182611b32919061abd8565b90506000808212611b47578460400151611b4d565b84602001515b90506000808312611b62578560200151611b68565b85604001515b9050600080841215611b835783611b7e9061ac1b565b611b85565b835b90508061010960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410611c81578061010960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611c1e919061ac63565b925050819055508061010960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611c75919061ac97565b92505081905550611f7f565b600061010960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008183611cd4919061ac63565b905080610109600060fc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410611e5a578261010960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611d8f919061ac97565b9250508190555080610109600060fc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611e08919061ac63565b92505081905550600061010960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611f7c565b8261010960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611eaa919061ac97565b92505081905550600061010960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550806101046000828254611f0a919061ac97565b925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f4d546f30d16f3c74d3123fc4c51ff5535dc33bb0819abeb9d8e9e097a72fefbb8361010454604051611f7392919061accb565b60405180910390a35b50505b505050505050505b8080611f929061a9dd565b91505061187e565b50600161010360016101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f348b1bdb7ff7d772beba31e8d7cb57d1c2b108eaa1e3ac5ebce592c036e328d060fd60149054906101000a900460ff1661010a8054905060405161201492919061acf4565b60405180910390a26001606581905550565b600061010960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b612076616fa2565b610107600083600381111561208e5761208d619a58565b5b60038111156120a05761209f619a58565b5b815260200190815260200160002060009054906101000a900460ff166120fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120f29061ad8f565b60405180910390fd5b80610106600084600381111561211457612113619a58565b5b600381111561212657612125619a58565b5b815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600080600061010960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060006121ca8686613867565b905060006121d9876001612210565b905060008282856121ea919061adaf565b6121f4919061abd8565b9050600080821215905081819650965050505050509250929050565b60008061010d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900461ffff1661ffff16036122745760009050612743565b61227c619512565b60005b61010b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508110156126c65761010c60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020818154811061231c5761231b61aa25565b5b90600052602060002090602091828204019190069054906101000a900460ff16156126b357600061010a61010b60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083815481106123985761239761aa25565b5b90600052602060002090601091828204019190066002029054906101000a900461ffff1661ffff16815481106123d1576123d061aa25565b5b90600052602060002090600902016040518060c00160405290816000820180546123fa9061aa83565b80601f01602080910402602001604051908101604052809291908181526020018280546124269061aa83565b80156124735780601f1061244857610100808354040283529160200191612473565b820191906000526020600020905b81548152906001019060200180831161245657829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582016040518060c00160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff16600a81111561258c5761258b619a58565b5b600a81111561259e5761259d619a58565b5b815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff1660038111156125db576125da619a58565b5b60038111156125ed576125ec619a58565b5b81526020016003820160019054906101000a900460ff1660ff1660ff1681525050815250509050600081608001510361262657506126b3565b60008160a0015160200151600a81111561264357612642619a58565b5b905060006126588360a0015160800151616c2b565b90506000612667898386616dca565b905080868460ff16600b81106126805761267f61aa25565b5b602002015161268f919061adaf565b868460ff16600b81106126a5576126a461aa25565b5b602002018181525050505050505b80806126be9061a9dd565b91505061227f565b50600080600090505b600b60ff1681101561273c5760008382600b81106126f0576126ef61aa25565b5b60200201511280156126ff5750845b612729578281600b81106127165761271561aa25565b5b602002015182612726919061adaf565b91505b80806127349061a9dd565b9150506126cf565b5080925050505b92915050565b6101076020528060005260406000206000915054906101000a900460ff1681565b6002606554036127af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127a69061a78c565b60405180910390fd5b6002606581905550600081116127fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127f19061ae3f565b60405180910390fd5b61010960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481111561287d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128749061aeab565b60405180910390fd5b60006101045461010554612891919061ac97565b61010554836128a0919061a69e565b6128aa919061aefa565b90508161010960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546128fc919061ac63565b92505081905550808261290f919061ac63565b6101046000828254612921919061ac63565b925050819055506000612935336000612176565b91505080612978576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161296f9061af77565b60405180910390fd5b6129c5338360fb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166170209092919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167f5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab8484604051612a0d92919061accb565b60405180910390a25050600160658190555050565b600260655403612a67576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a5e9061a78c565b60405180910390fd5b6002606581905550600061010d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900461ffff1661ffff1611612b04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612afb9061afe3565b60405180910390fd5b6000612b11826000612176565b9150508015612b55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b4c9061b075565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1603612bc3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bba9061b107565b60405180910390fd5b60005b61010b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050811015612dc75761010c60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208181548110612c6357612c6261aa25565b5b90600052602060002090602091828204019190069054906101000a900460ff1615612db4578273ffffffffffffffffffffffffffffffffffffffff1661010a61010b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208381548110612cf457612cf361aa25565b5b90600052602060002090601091828204019190066002029054906101000a900461ffff1661ffff1681548110612d2d57612d2c61aa25565b5b906000526020600020906009020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160315612db4576000612d883385846170a6565b905080612d955750612db4565b612da0846000612176565b9050809350508215612db25750612dc7565b505b8080612dbf9061a9dd565b915050612bc6565b5060005b61010b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050811015612fcc5761010c60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208181548110612e6857612e6761aa25565b5b90600052602060002090602091828204019190069054906101000a900460ff1615612fb9578273ffffffffffffffffffffffffffffffffffffffff1661010a61010b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208381548110612ef957612ef861aa25565b5b90600052602060002090601091828204019190066002029054906101000a900461ffff1661ffff1681548110612f3257612f3161aa25565b5b906000526020600020906009020160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160315612fb9576000612f8d3385846170a6565b905080612f9a5750612fb9565b612fa5846000612176565b9050809350508215612fb75750612fcc565b505b8080612fc49061a9dd565b915050612dcb565b5050600160658190555050565b612fe1616fa2565b60005b828290508110156131375761010860008484848181106130075761300661aa25565b5b905060200201602081019061301c9190619e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156130a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161309b9061b173565b60405180910390fd5b600161010860008585858181106130be576130bd61aa25565b5b90506020020160208101906130d39190619e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550808061312f9061a9dd565b915050612fe4565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16036131ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131c19061b205565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166132096179c9565b73ffffffffffffffffffffffffffffffffffffffff161461325f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132569061b297565b60405180910390fd5b61326881617a20565b6132c181600067ffffffffffffffff81111561328757613286619eac565b5b6040519080825280601f01601f1916602001820160405280156132b95781602001600182028036833780820191505090505b506000617a2b565b50565b60fb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60fd60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff160361339e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133959061b205565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166133dd6179c9565b73ffffffffffffffffffffffffffffffffffffffff1614613433576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161342a9061b297565b60405180910390fd5b61343c82617a20565b61344882826001617a2b565b5050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16146134dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134d39061b329565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b61350d616fa2565b60005b828290508110156136625761010260008484848181106135335761353261aa25565b5b90506020020160208101906135489190619e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166135cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135c69061b3bb565b60405180910390fd5b600061010260008585858181106135e9576135e861aa25565b5b90506020020160208101906135fe9190619e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550808061365a9061a9dd565b915050613510565b505050565b61366f616fa2565b610107600083600381111561368757613686619a58565b5b600381111561369957613698619a58565b5b815260200190815260200160002060009054906101000a900460ff166136f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136eb9061b44d565b60405180910390fd5b60008111613737576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161372e9061b4df565b60405180910390fd5b8060ff600084600381111561374f5761374e619a58565b5b600381111561376157613760619a58565b5b8152602001908152602001600020819055505050565b61377f616fa2565b8073ffffffffffffffffffffffffffffffffffffffff1660fc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361380f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138069061b571565b60405180910390fd5b8060fc60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61385b616fa2565b6138656000617b99565b565b60008061010b60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050036138bc57600090506142ea565b600080600090505b61010b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508110156142e45761010c60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081815481106139615761396061aa25565b5b90600052602060002090602091828204019190069054906101000a900460ff16156142d157600061010a61010b60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083815481106139dd576139dc61aa25565b5b90600052602060002090601091828204019190066002029054906101000a900461ffff1661ffff1681548110613a1657613a1561aa25565b5b90600052602060002090600902016040518060c0016040529081600082018054613a3f9061aa83565b80601f0160208091040260200160405190810160405280929190818152602001828054613a6b9061aa83565b8015613ab85780601f10613a8d57610100808354040283529160200191613ab8565b820191906000526020600020905b815481529060010190602001808311613a9b57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582016040518060c00160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff16600a811115613bd157613bd0619a58565b5b600a811115613be357613be2619a58565b5b815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff166003811115613c2057613c1f619a58565b5b6003811115613c3257613c31619a58565b5b81526020016003820160019054906101000a900460ff1660ff1660ff168152505081525050905060008160a0015190506000613c6d82617c5f565b90506000836080015103613c83575050506142d1565b600080846020015173ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff161480613cf35750846040015173ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff16145b613d32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d299061b603565b60405180910390fd5b846020015173ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff1603613d8057846080015182613d79919061ac97565b9150613d93565b846080015181613d90919061ac97565b90505b60006001905060005b61010b60008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508110156141ec5780881480613e64575061010c60008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208181548110613e4257613e4161aa25565b5b90600052602060002090602091828204019190069054906101000a900460ff16155b6141d957600061010a61010b60008f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208381548110613ebf57613ebe61aa25565b5b90600052602060002090601091828204019190066002029054906101000a900461ffff1661ffff1681548110613ef857613ef761aa25565b5b90600052602060002090600902016040518060c0016040529081600082018054613f219061aa83565b80601f0160208091040260200160405190810160405280929190818152602001828054613f4d9061aa83565b8015613f9a5780601f10613f6f57610100808354040283529160200191613f9a565b820191906000526020600020905b815481529060010190602001808311613f7d57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582016040518060c00160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff16600a8111156140b3576140b2619a58565b5b600a8111156140c5576140c4619a58565b5b815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff16600381111561410257614101619a58565b5b600381111561411457614113619a58565b5b81526020016003820160019054906101000a900460ff1660ff1660ff1681525050815250509050600061414a8260a00151617c5f565b9050600082608001510361415f5750506141d9565b8087036141d657816020015173ffffffffffffffffffffffffffffffffffffffff168e73ffffffffffffffffffffffffffffffffffffffff16036141b4578160800151866141ad919061ac97565b95506141c7565b8160800151856141c4919061ac97565b94505b83806141d29061a9dd565b9450505b50505b80806141e49061a9dd565b915050613d9c565b5060008083851061420e578385614203919061ac63565b915060019050614221565b848461421a919061ac63565b9150600090505b60008211156142c85760006142398860800151616c2b565b9050600061425489608001518a600001518b60200151617ca5565b905060008e156142755761426e83858c8561010054617e65565b9050614288565b61428583858c8561010054618075565b90505b6004600a614296919061a653565b866142a1919061a69e565b81866142ad919061a69e565b6142b7919061aefa565b8d6142c2919061ac97565b9c505050505b50505050505050505b80806142dc9061a9dd565b9150506138c4565b50809150505b92915050565b6142f8616fa2565b60005b8282905081101561444e57610102600084848481811061431e5761431d61aa25565b5b90506020020160208101906143339190619e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156143bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016143b29061b695565b60405180910390fd5b600161010260008585858181106143d5576143d461aa25565b5b90506020020160208101906143ea9190619e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080806144469061a9dd565b9150506142fb565b505050565b600260655403614498576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161448f9061a78c565b60405180910390fd5b600260658190555061010860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661452d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016145249061a81e565b60405180910390fd5b6000816060013511614574576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161456b9061b727565b60405180910390fd5b60008160800135116145bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016145b29061b7b9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1661010660008360e00160208101906145eb9190619936565b60038111156145fd576145fc619a58565b5b600381111561460f5761460e619a58565b5b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603614694576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161468b9061b84b565b60405180910390fd5b8060400160208101906146a79190619e50565b73ffffffffffffffffffffffffffffffffffffffff168160200160208101906146d09190619e50565b73ffffffffffffffffffffffffffffffffffffffff1603614726576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161471d9061b8dd565b60405180910390fd5b6000614730615978565b9050600061010e60008460e001602081019061474c9190619936565b600381111561475e5761475d619a58565b5b60038111156147705761476f619a58565b5b81526020019081526020016000208360c0016020810190614791919061b922565b600a8111156147a3576147a2619a58565b5b60ff16600b81106147b7576147b661aa25565b5b01549050600081116147fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016147f59061b9c1565b60405180910390fd5b60006040518060c0016040528085806000019061481b919061b9f0565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018560200160208101906148769190619e50565b73ffffffffffffffffffffffffffffffffffffffff1681526020018560400160208101906148a49190619e50565b73ffffffffffffffffffffffffffffffffffffffff16815260200185606001358152602001856080013581526020016040518060c001604052808760a00160208101906148f1919061ba53565b151581526020018760c001602081019061490b919061b922565b600a81111561491d5761491c619a58565b5b81526020018581526020016101015481526020018760e00160208101906149449190619936565b600381111561495657614955619a58565b5b81526020018660ff16815250815250905060ff60008560e001602081019061497e9190619936565b60038111156149905761498f619a58565b5b60038111156149a2576149a1619a58565b5b815260200190815260200160002054816080015110156149f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016149ee9061bacc565b60405180910390fd5b600080614a2b8387610100016020810190614a12919061ba53565b88610120016020810190614a26919061ba53565b61828d565b91509150816101096000856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015614ab6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614aad9061bb5e565b60405180910390fd5b806101096000856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015614b3d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614b349061bbf0565b60405180910390fd5b816101096000856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254614b91919061ac63565b92505081905550806101096000856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254614bec919061ac63565b9250508190555061010a8390806001815401808255809150506001900390600052602060002090600902016000909190919091506000820151816000019081614c35919061bdb2565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a08201518160050160008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548160ff0219169083600a811115614d2c57614d2b619a58565b5b0217905550604082015181600101556060820151816002015560808201518160030160006101000a81548160ff02191690836003811115614d7057614d6f619a58565b5b021790555060a08201518160030160016101000a81548160ff021916908360ff160217905550505050506000600161010a80549050614daf919061ac63565b905061010c6000856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905061010b6000866020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905014614e4b57600080fd5b61010c6000856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905061010b6000866040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905014614ee557600080fd5b61010d6000856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900461ffff1661ffff1661010b6000866020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490501015614f8f57600080fd5b61010d6000856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900461ffff1661ffff1661010b6000866040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050101561503957600080fd5b61010b6000856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190806001815401808255809150506001900390600052602060002090601091828204019190066002029091909190916101000a81548161ffff021916908361ffff16021790555061010b6000856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190806001815401808255809150506001900390600052602060002090601091828204019190066002029091909190916101000a81548161ffff021916908361ffff16021790555061010c6000856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600190806001815401808255809150506001900390600052602060002090602091828204019190069091909190916101000a81548160ff02191690831515021790555061010c6000856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600190806001815401808255809150506001900390600052602060002090602091828204019190069091909190916101000a81548160ff02191690831515021790555061010d6000856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081819054906101000a900461ffff16809291906152c19061be92565b91906101000a81548161ffff021916908361ffff1602179055505061010d6000856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081819054906101000a900461ffff168092919061533e9061be92565b91906101000a81548161ffff021916908361ffff16021790555050600061536a85602001516000612176565b915050600061537e86604001516000612176565b915050816153c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016153b89061bf2e565b60405180910390fd5b80615401576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016153f89061bfc0565b60405180910390fd5b61547b60fd60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168587615433919061ac97565b60fb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166170209092919063ffffffff16565b7f2bea814922af00322b50dbedc22798d1e1c00ddd12bbc8fb60e2ca3d6fe7660e8980600001906154ac919061b9f0565b8b606001358c608001358d60a00160208101906154c9919061ba53565b8e60e00160208101906154dc9190619936565b8f60c00160208101906154ef919061b922565b6101015460405161550798979695949392919061c02b565b60405180910390a15050505050505050600160658190555050565b600260655403615567576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161555e9061a78c565b60405180910390fd5b6002606581905550600061010960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600081116155f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016155ee9061c0ef565b60405180910390fd5b6000610104546101055461560b919061ac97565b610105548361561a919061a69e565b615624919061aefa565b9050600061010960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508082615678919061ac63565b610104600082825461568a919061ac63565b92505081905550600061569e336000612176565b915050806156e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016156d89061c15b565b60405180910390fd5b61572e338360fb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166170209092919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167f5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab848460405161577692919061accb565b60405180910390a25050506001606581905550565b60fc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600061010960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b61582d616fa2565b612710811115615872576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016158699061c1ed565b60405180910390fd5b80610100819055503373ffffffffffffffffffffffffffffffffffffffff167fc4d2cbc4040c3ef45136b4540fbaaf123b01b278d3323ff57d3a3758c8a312ce826040516158c09190619869565b60405180910390a250565b60608060005b8585905081101561596f576000806159108888858181106158f5576158f461aa25565b5b905060200201602081019061590a9190619e50565b87612176565b91509150818584815181106159285761592761aa25565b5b602002602001018181525050808484815181106159485761594761aa25565b5b602002602001019015159081151581525050505080806159679061a9dd565b9150506158d1565b50935093915050565b600060fb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156159e7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615a0b919061c239565b905090565b60fd60149054906101000a900460ff1681565b60fe5481565b600260655403615a6e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615a659061a78c565b60405180910390fd5b600260658190555060008111615ab9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615ab09061c2b2565b60405180910390fd5b60fc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614615bb95760fe548161010960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054615b5d919061ac97565b1115615b9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615b959061c344565b60405180910390fd5b806101056000828254615bb1919061ac97565b925050819055505b8061010960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254615c09919061ac97565b92505081905550615c5f33308360fb60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166183fa909392919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167f2d8a08b6430a894aea608bcaa6013d5d3e263bc49110605e4d4ba76930ae5c2982604051615ca59190619869565b60405180910390a2600160658190555050565b6101005481565b615cc7616fa2565b61010360009054906101000a900460ff161561010360006101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f66f8a7ac2d1f7da12f55df6017dd3e4785a2c733138e53bc548a977cbee1099861010360009054906101000a900460ff16604051615d499190619e08565b60405180910390a2565b60ff6020528060005260406000206000915090505481565b61010e60205281600052604060002081600b8110615d8857600080fd5b01600091509150505481565b615d9c616fa2565b8073ffffffffffffffffffffffffffffffffffffffff1660fd60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603615e2c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615e239061c3d6565b60405180910390fd5b8060fd60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b615e78616fa2565b60005b82829050811015615fcd576101086000848484818110615e9e57615e9d61aa25565b5b9050602002016020810190615eb39190619e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16615f3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615f319061c442565b60405180910390fd5b60006101086000858585818110615f5457615f5361aa25565b5b9050602002016020810190615f699190619e50565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080615fc59061a9dd565b915050615e7b565b505050565b615fda616fa2565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603616049576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016160409061c4d4565b60405180910390fd5b61605281617b99565b50565b61605d616fa2565b600081116160a0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016160979061c540565b60405180910390fd5b8060fe819055503373ffffffffffffffffffffffffffffffffffffffff167f43bf10be81662116415358cada6e4928c87d70e8c74de1d25886dd63940c215d826040516160ed9190619869565b60405180910390a250565b616100616fa2565b610107600084600381111561611857616117619a58565b5b600381111561612a57616129619a58565b5b815260200190815260200160002060009054906101000a900460ff1615616186576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161617d9061c5d2565b60405180910390fd5b616191838383616379565b3373ffffffffffffffffffffffffffffffffffffffff167fb074d74b8b193c1b3395bc3dbb1b97e42880495c0f2708754d80d19921d14f958484846040516161db9392919061c5f2565b60405180910390a2505050565b6161f0619535565b61010e600083600381111561620857616207619a58565b5b600381111561621a57616219619a58565b5b8152602001908152602001600020600b806020026040519081016040528092919082600b801561625f576020028201915b81548152602001906001019080831161624b575b50505050509050919050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff166162dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016162d49061c69b565b60405180910390fd5b6162e5618483565b565b600060019054906101000a900460ff16616336576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161632d9061c69b565b60405180910390fd5b61633e6184dc565b565b600062093a8082616351919061ac97565b421115616368576163614261853d565b9050616374565b6163718261853d565b90505b919050565b610107600084600381111561639157616390619a58565b5b60038111156163a3576163a2619a58565b5b815260200190815260200160002060009054906101000a900460ff16156163ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016163f69061c72d565b60405180910390fd5b60008111616442576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016164399061c7bf565b60405180910390fd5b81610106600085600381111561645b5761645a619a58565b5b600381111561646d5761646c619a58565b5b815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060ff60008560038111156164d1576164d0619a58565b5b60038111156164e3576164e2619a58565b5b8152602001908152602001600020819055506164fe8361659b565b61010e600085600381111561651657616515619a58565b5b600381111561652857616527619a58565b5b815260200190815260200160002090600b6165449291906193f6565b506001610107600085600381111561655f5761655e619a58565b5b600381111561657157616570619a58565b5b815260200190815260200160002060006101000a81548160ff021916908315150217905550505050565b6165a3619535565b4261010154116165e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016165df9061c851565b60405180910390fd5b60006165f2615978565b905060006165ff84616c2b565b905060028261660e919061c871565b600a61661a919061a653565b81101561665c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016166539061c918565b60405180910390fd5b600682616669919061a9a8565b600a616675919061a653565b8111156166b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016166ae9061c9aa565b60405180910390fd5b600080600080604051806104200160405280606464ffffffffff16815260200160c864ffffffffff16815260200161012c64ffffffffff16815260200161025864ffffffffff1681526020016103e864ffffffffff1681526020016107d064ffffffffff168152602001610bb864ffffffffff16815260200161177064ffffffffff16815260200161271064ffffffffff168152602001614e2064ffffffffff16815260200161753064ffffffffff16815260200161ea6064ffffffffff168152602001620186a064ffffffffff16815260200162030d4064ffffffffff16815260200162061a8064ffffffffff168152602001620aae6064ffffffffff168152602001620f424064ffffffffff168152602001621e848064ffffffffff168152602001623d090064ffffffffff168152602001626acfc064ffffffffff1681526020016298968064ffffffffff1681526020016301312d0064ffffffffff1681526020016302625a0064ffffffffff16815260200163042c1d8064ffffffffff1681526020016305f5e10064ffffffffff168152602001630bebc20064ffffffffff1681526020016317d7840064ffffffffff168152602001632faf080064ffffffffff168152602001633b9aca0064ffffffffff16815260200163b2d05e0064ffffffffff16815260200164012a05f20064ffffffffff1681526020016401dcd6500064ffffffffff1681526020016402540be40064ffffffffff1681525090506000604051806104000160405280600963ffffffff168152602001601463ffffffff168152602001601e63ffffffff168152602001603263ffffffff168152602001605a63ffffffff16815260200160c863ffffffff16815260200161012c63ffffffff16815260200161025863ffffffff1681526020016103e863ffffffff1681526020016107d063ffffffff168152602001610bb863ffffffff16815260200161177063ffffffff16815260200161271063ffffffff168152602001614e2063ffffffff16815260200161753063ffffffff16815260200161ea6063ffffffff168152602001620186a063ffffffff16815260200162030d4063ffffffff16815260200162061a8063ffffffff168152602001620927c063ffffffff168152602001620f424063ffffffff168152602001621e848063ffffffff168152602001623d090063ffffffff168152602001626acfc063ffffffff1681526020016298968063ffffffff1681526020016301312d0063ffffffff1681526020016302625a0063ffffffff16815260200163042c1d8063ffffffff1681526020016305f5e10063ffffffff168152602001630bebc20063ffffffff1681526020016317d7840063ffffffff1681526020016329b9270063ffffffff16815250905060005b6020811015616c1e57600488616ad2919061c871565b600a616ade919061a653565b838260218110616af157616af061aa25565b5b602002015164ffffffffff16616b07919061a69e565b9550600488616b16919061c871565b600a616b22919061a653565b83600183616b30919061ac97565b60218110616b4157616b4061aa25565b5b602002015164ffffffffff16616b57919061a69e565b9450600488616b66919061c871565b600a616b72919061a653565b828260208110616b8557616b8461aa25565b5b602002015163ffffffff16616b9a919061a69e565b9350858710158015616bab57508487105b15616c0b5760005b600b60ff16811015616c05578481616bcb919061a69e565b87616bd6919061ac97565b8a82600b8110616be957616be861aa25565b5b6020020181815250508080616bfd9061a9dd565b915050616bb3565b50616c1e565b8080616c169061a9dd565b915050616abc565b5050505050505050919050565b60008073ffffffffffffffffffffffffffffffffffffffff166101066000846003811115616c5c57616c5b619a58565b5b6003811115616c6e57616c6d619a58565b5b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603616cf3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616cea9061ca16565b60405180910390fd5b60006101066000846003811115616d0d57616d0c619a58565b5b6003811115616d1f57616d1e619a58565b5b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635ec81e296040518163ffffffff1660e01b8152600401608060405180830381865afa158015616d98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616dbc919061ca8d565b505091505080915050919050565b6000816020015173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480616e395750816040015173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16145b616e78576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616e6f9061cb66565b60405180910390fd5b6000826020015173ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614616eb8576000616ebb565b60015b905060008360a00151905060008460a001516000015115616f03578160400151616ef28360400151886185ee90919063ffffffff16565b616efc919061abd8565b9050616f28565b85616f1b8784604001516185ee90919063ffffffff16565b616f25919061abd8565b90505b846060015181616f38919061abd8565b935082616f6e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84616f6b919061cb86565b93505b6004600a616f7c919061a653565b856080015185616f8c919061cb86565b616f96919061cbfe565b93505050509392505050565b616faa618608565b73ffffffffffffffffffffffffffffffffffffffff16616fc86157b1565b73ffffffffffffffffffffffffffffffffffffffff161461701e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016170159061ccb4565b60405180910390fd5b565b6170a18363a9059cbb60e01b848460405160240161703f92919061ccd4565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050618610565b505050565b60008061010a61010b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002084815481106170fe576170fd61aa25565b5b90600052602060002090601091828204019190066002029054906101000a900461ffff1661ffff16815481106171375761713661aa25565b5b9060005260206000209060090201905060008473ffffffffffffffffffffffffffffffffffffffff168260010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146171a75760006171aa565b60015b905060006171cb8360050160030160009054906101000a900460ff16616c2b565b905060006172168460050160030160009054906101000a900460ff168560050160000160009054906101000a900460ff168660050160000160019054906101000a900460ff16617ca5565b9050600061747c8560010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684876040518060c001604052908160008201805461725f9061aa83565b80601f016020809104026020016040519081016040528092919081815260200182805461728b9061aa83565b80156172d85780601f106172ad576101008083540402835291602001916172d8565b820191906000526020600020905b8154815290600101906020018083116172bb57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582016040518060c00160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff16600a8111156173f1576173f0619a58565b5b600a81111561740357617402619a58565b5b815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff1660038111156174405761743f619a58565b5b600381111561745257617451619a58565b5b81526020016003820160019054906101000a900460ff1660ff1660ff168152505081525050616dca565b9050600084617494578161748f9061ac1b565b617496565b815b9050600081126174af57600096505050505050506179c2565b6000856174e0578660010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16617506565b8660020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff165b90508073ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff16036176ab57600084836175489061ac1b565b617552919061abd8565b905060008112156175a857600080821261756c5781617577565b816175769061ac1b565b5b905060006175878e8e848f6186d7565b9050806175a15760009a50505050505050505050506179c2565b50506176a5565b6000886004018190555061010d60008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081819054906101000a900461ffff16809291906176109061ccfd565b91906101000a81548161ffff021916908361ffff1602179055505061010d60008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081819054906101000a900461ffff16809291906176899061ccfd565b91906101000a81548161ffff021916908361ffff160217905550505b506176d3565b60006176b98c8c878d6186d7565b9050806176d1576000985050505050505050506179c2565b505b60006177b586888a6005016040518060c00160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff16600a81111561772957617728619a58565b5b600a81111561773b5761773a619a58565b5b815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff16600381111561777857617777619a58565b5b600381111561778a57617789619a58565b5b81526020016003820160019054906101000a900460ff1660ff1660ff16815250508861010054618075565b905060646023826177c6919061a69e565b6177d0919061aefa565b61010960008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015617852576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016178499061cd98565b60405180910390fd5b6064602382617861919061a69e565b61786b919061aefa565b61010960008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546178ba919061ac63565b9250508190555060646019826178d0919061a69e565b6178da919061aefa565b61010960008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254617929919061ac97565b92505081905550600a8161793d919061aefa565b610109600060fc60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546179ae919061ac97565b925050819055506001985050505050505050505b9392505050565b60006179f77f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b618fba565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b617a28616fa2565b50565b617a577f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b618fc4565b60000160009054906101000a900460ff1615617a7b57617a7683618fce565b617b94565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015617ae357506040513d601f19601f82011682018060405250810190617ae0919061cde4565b60015b617b22576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401617b199061ce83565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114617b87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401617b7e9061cf15565b60405180910390fd5b50617b93838383619087565b5b505050565b600060c960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160c960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60008160000151826080015183604001518460600151604051602001617c88949392919061cfb5565b604051602081830303815290604052805190602001209050919050565b60008073ffffffffffffffffffffffffffffffffffffffff166101066000866003811115617cd657617cd5619a58565b5b6003811115617ce857617ce7619a58565b5b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603617d6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401617d649061d04f565b60405180910390fd5b6101066000856003811115617d8557617d84619a58565b5b6003811115617d9757617d96619a58565b5b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e224f68484600a811115617df557617df4619a58565b5b6040518363ffffffff1660e01b8152600401617e1292919061d06f565b608060405180830381865afa158015617e2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617e53919061ca8d565b90919250905050809150509392505050565b600042846060015111617ead576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401617ea49061d10a565b60405180910390fd5b8415617ee557600a617ed4600a85617ec5919061a69e565b886190b390919063ffffffff16565b617ede919061aefa565b9050618051565b836000015115617f94576000868560400151617f01919061abd8565b90506000811215617f1157600090505b6000816103e8617f21919061cb86565b8860c8617f2e919061cb86565b617f38919061abd8565b9050600088607d617f49919061a69e565b90506000821215617f69576103e881617f62919061aefa565b9350617f8c565b6103e8617f7f82846185ee90919063ffffffff16565b617f89919061aefa565b93505b505050618050565b6000846040015187617fa6919061abd8565b90506000811215617fb657600090505b6000816103e8617fc6919061cb86565b8860c8617fd3919061cb86565b617fdd919061abd8565b9050600088607d617fee919061a69e565b905060008083121561800257819050618018565b61801582846185ee90919063ffffffff16565b90505b6103e861803f89604001516101f4618030919061a69e565b836190b390919063ffffffff16565b618049919061aefa565b9450505050505b5b600061805d87846190cc565b90508181111561806b578091505b5095945050505050565b6000428460600151116180bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016180b49061d19c565b60405180910390fd5b8415618103576103e86180f26103e8856180d7919061a69e565b6041896180e4919061a69e565b6190b390919063ffffffff16565b6180fc919061aefa565b9050618269565b8360000151156181af57600086856040015161811f919061abd8565b9050600081121561812f57600090505b600081606461813e919061cb86565b88600a61814b919061cb86565b618155919061abd8565b90506000886008618166919061a69e565b905060008212156181855760648161817e919061aefa565b93506181a7565b606461819a82846185ee90919063ffffffff16565b6181a4919061aefa565b93505b505050618268565b60008460400151876181c1919061abd8565b905060008112156181d157600090505b60008160646181e0919061cb86565b88600a6181ed919061cb86565b6181f7919061abd8565b90506000886008618208919061a69e565b905060008083121561821c57819050618232565b61822f82846185ee90919063ffffffff16565b90505b606461825789604001516032618248919061a69e565b836190b390919063ffffffff16565b618261919061aefa565b9450505050505b5b600061827587846190cc565b905081811115618283578091505b5095945050505050565b600080600061829b866190ef565b905060008660600151905060006127106182ce8460066182bb919061a69e565b846103e86182c9919061a69e565b6190b3565b6182d8919061aefa565b905060006127106183028560036182ef919061a69e565b856103e86182fd919061a69e565b6190b3565b61830c919061aefa565b9050871561837a5761010260008a6020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166183705780618373565b60005b955061837e565b8195505b86156183ea5761010260008a6040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166183e057806183e3565b60005b94506183ee565b8194505b50505050935093915050565b61847d846323b872dd60e01b85858560405160240161841b9392919061d1bc565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050618610565b50505050565b600060019054906101000a900460ff166184d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016184c99061c69b565b60405180910390fd5b6001606581905550565b600060019054906101000a900460ff1661852b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016185229061c69b565b60405180910390fd5b61853b618536618608565b617b99565b565b600080600760046201518085618553919061aefa565b61855d919061ac97565b618567919061d1f3565b9050600062015180600783600c61857e919061ac63565b618588919061d1f3565b618592919061a69e565b8461859d919061ac97565b9050600061708062015180836185b3919061d1f3565b836185be919061ac63565b6185c8919061ac97565b90508085106185e35762093a80816185e0919061ac97565b90505b809350505050919050565b6000818310156185fe5781618600565b825b905092915050565b600033905090565b6000618672826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166191309092919063ffffffff16565b90506000815111156186d25780806020019051810190618692919061d239565b6186d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016186c89061d2d8565b60405180910390fd5b5b505050565b60008061010b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020838154811061872c5761872b61aa25565b5b90600052602060002090601091828204019190066002029054906101000a900461ffff169050600061010a8261ffff168154811061876d5761876c61aa25565b5b9060005260206000209060090201905060008673ffffffffffffffffffffffffffffffffffffffff168260010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146187dd5760006187e0565b60015b90508561010960008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156188365760009350505050618fb2565b8561010960008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254618886919061ac63565b925050819055508561010960008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546188dd919061ac97565b9250508190555061010b60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208390806001815401808255809150506001900390600052602060002090601091828204019190066002029091909190916101000a81548161ffff021916908361ffff16021790555061010c60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600190806001815401808255809150506001900390600052602060002090602091828204019190069091909190916101000a81548160ff02191690831515021790555061010d60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081819054906101000a900461ffff1680929190618a4d9061be92565b91906101000a81548161ffff021916908361ffff16021790555050600061010c60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208681548110618abc57618abb61aa25565b5b90600052602060002090602091828204019190066101000a81548160ff02191690831515021790555061010d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081819054906101000a900461ffff1680929190618b439061ccfd565b91906101000a81548161ffff021916908361ffff160217905550508015618bac57878260010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550618bf0565b878260020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b618bfb886000612176565b90508094505083618fae5761010b60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480618c5657618c5561d2f8565b5b60019003818190600052602060002090601091828204019190066002026101000a81549061ffff0219169055905561010c60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480618cd457618cd361d2f8565b5b60019003818190600052602060002090602091828204019190066101000a81549060ff0219169055905561010d60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081819054906101000a900461ffff1680929190618d5c9061ccfd565b91906101000a81548161ffff021916908361ffff16021790555050600161010c60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208681548110618dcb57618dca61aa25565b5b90600052602060002090602091828204019190066101000a81548160ff02191690831515021790555061010d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081819054906101000a900461ffff1680929190618e529061be92565b91906101000a81548161ffff021916908361ffff160217905550508015618ebb57868260010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550618eff565b868260020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b8561010960008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254618f4f919061ac97565b925050819055508561010960008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254618fa6919061ac63565b925050819055505b5050505b949350505050565b6000819050919050565b6000819050919050565b618fd78161626b565b619016576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161900d9061d399565b60405180910390fd5b806190437f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b618fba565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61909083619148565b60008251118061909d5750805b156190ae576190ac8383619197565b505b505050565b60008183106190c257816190c4565b825b905092915050565b600061271082846190dd919061a69e565b6190e7919061aefa565b905092915050565b60006004600a6190ff919061a653565b6191108360a0015160800151616c2b565b836080015161911f919061a69e565b619129919061aefa565b9050919050565b606061913f848460008561927b565b90509392505050565b61915181618fce565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b60606191a28361626b565b6191e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016191d89061d42b565b60405180910390fd5b6000808473ffffffffffffffffffffffffffffffffffffffff1684604051619209919061d492565b600060405180830381855af49150503d8060008114619244576040519150601f19603f3d011682016040523d82523d6000602084013e619249565b606091505b5091509150619271828260405180606001604052806027815260200161d6036027913961938f565b9250505092915050565b6060824710156192c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016192b79061d51b565b60405180910390fd5b6192c98561626b565b619308576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016192ff9061d587565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051619331919061d492565b60006040518083038185875af1925050503d806000811461936e576040519150601f19603f3d011682016040523d82523d6000602084013e619373565b606091505b509150915061938382828661938f565b92505050949350505050565b6060831561939f578290506193ef565b6000835111156193b25782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016193e6919061d5e0565b60405180910390fd5b9392505050565b82600b8101928215619425579160200282015b82811115619424578251825591602001919060010190619409565b5b5090506194329190619558565b5090565b50805460008255600902906000526020600020908101906194579190619575565b50565b50805460008255600f01601090049060005260206000209081019061947f9190619558565b50565b50805460008255601f0160209004906000526020600020908101906194a79190619558565b50565b6040518060c0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815260200161950c619661565b81525090565b604051806101600160405280600b90602082028036833780820191505090505090565b604051806101600160405280600b90602082028036833780820191505090505090565b5b80821115619571576000816000905550600101619559565b5090565b5b8082111561965d576000808201600061958f91906196c0565b6001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556003820160009055600482016000905560058201600080820160006101000a81549060ff02191690556000820160016101000a81549060ff0219169055600182016000905560028201600090556003820160006101000a81549060ff02191690556003820160016101000a81549060ff0219169055505050600901619576565b5090565b6040518060c001604052806000151581526020016000600a81111561968957619688619a58565b5b81526020016000815260200160008152602001600060038111156196b0576196af619a58565b5b8152602001600060ff1681525090565b5080546196cc9061aa83565b6000825580601f106196de57506196fd565b601f0160209004906000526020600020908101906196fc9190619558565b5b50565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061973f82619714565b9050919050565b61974f81619734565b811461975a57600080fd5b50565b60008135905061976c81619746565b92915050565b6004811061977f57600080fd5b50565b60008135905061979181619772565b92915050565b6000819050919050565b6197aa81619797565b81146197b557600080fd5b50565b6000813590506197c7816197a1565b92915050565b60008060008060008060c087890312156197ea576197e961970a565b5b60006197f889828a0161975d565b965050602061980989828a0161975d565b955050604061981a89828a0161975d565b945050606061982b89828a01619782565b935050608061983c89828a0161975d565b92505060a061984d89828a016197b8565b9150509295509295509295565b61986381619797565b82525050565b600060208201905061987e600083018461985a565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126198a9576198a8619884565b5b8235905067ffffffffffffffff8111156198c6576198c5619889565b5b6020830191508360208202830111156198e2576198e161988e565b5b9250929050565b60008060208385031215619900576198ff61970a565b5b600083013567ffffffffffffffff81111561991e5761991d61970f565b5b61992a85828601619893565b92509250509250929050565b60006020828403121561994c5761994b61970a565b5b600061995a84828501619782565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156199c95780820151818401526020810190506199ae565b60008484015250505050565b6000601f19601f8301169050919050565b60006199f18261998f565b6199fb818561999a565b9350619a0b8185602086016199ab565b619a14816199d5565b840191505092915050565b619a2881619734565b82525050565b619a3781619797565b82525050565b60008115159050919050565b619a5281619a3d565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600b8110619a9857619a97619a58565b5b50565b6000819050619aa982619a87565b919050565b6000619ab982619a9b565b9050919050565b619ac981619aae565b82525050565b60048110619ae057619adf619a58565b5b50565b6000819050619af182619acf565b919050565b6000619b0182619ae3565b9050919050565b619b1181619af6565b82525050565b600060ff82169050919050565b619b2d81619b17565b82525050565b60c082016000820151619b496000850182619a49565b506020820151619b5c6020850182619ac0565b506040820151619b6f6040850182619a2e565b506060820151619b826060850182619a2e565b506080820151619b956080850182619b08565b5060a0820151619ba860a0850182619b24565b50505050565b6000610160830160008301518482036000860152619bcc82826199e6565b9150506020830151619be16020860182619a1f565b506040830151619bf46040860182619a1f565b506060830151619c076060860182619a2e565b506080830151619c1a6080860182619a2e565b5060a0830151619c2d60a0860182619b33565b508091505092915050565b6000619c448383619bae565b905092915050565b6000602082019050919050565b6000619c6482619963565b619c6e818561996e565b935083602082028501619c808561997f565b8060005b85811015619cbc5784840389528151619c9d8582619c38565b9450619ca883619c4c565b925060208a01995050600181019050619c84565b50829750879550505050505092915050565b60006020820190508181036000830152619ce88184619c59565b905092915050565b60008060408385031215619d0757619d0661970a565b5b6000619d1585828601619782565b9250506020619d268582860161975d565b9150509250929050565b619d3981619a3d565b8114619d4457600080fd5b50565b600081359050619d5681619d30565b92915050565b60008060408385031215619d7357619d7261970a565b5b6000619d818582860161975d565b9250506020619d9285828601619d47565b9150509250929050565b6000819050919050565b619daf81619d9c565b82525050565b619dbe81619a3d565b82525050565b6000604082019050619dd96000830185619da6565b619de66020830184619db5565b9392505050565b6000602082019050619e026000830184619da6565b92915050565b6000602082019050619e1d6000830184619db5565b92915050565b600060208284031215619e3957619e3861970a565b5b6000619e47848285016197b8565b91505092915050565b600060208284031215619e6657619e6561970a565b5b6000619e748482850161975d565b91505092915050565b619e8681619734565b82525050565b6000602082019050619ea16000830184619e7d565b92915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b619ee4826199d5565b810181811067ffffffffffffffff82111715619f0357619f02619eac565b5b80604052505050565b6000619f16619700565b9050619f228282619edb565b919050565b600067ffffffffffffffff821115619f4257619f41619eac565b5b619f4b826199d5565b9050602081019050919050565b82818337600083830152505050565b6000619f7a619f7584619f27565b619f0c565b905082815260208101848484011115619f9657619f95619ea7565b5b619fa1848285619f58565b509392505050565b600082601f830112619fbe57619fbd619884565b5b8135619fce848260208601619f67565b91505092915050565b60008060408385031215619fee57619fed61970a565b5b6000619ffc8582860161975d565b925050602083013567ffffffffffffffff81111561a01d5761a01c61970f565b5b61a02985828601619fa9565b9150509250929050565b6000819050919050565b61a0468161a033565b82525050565b600060208201905061a061600083018461a03d565b92915050565b6000806040838503121561a07e5761a07d61970a565b5b600061a08c85828601619782565b925050602061a09d858286016197b8565b9150509250929050565b600080fd5b6000610140828403121561a0c35761a0c261a0a7565b5b81905092915050565b60006020828403121561a0e25761a0e161970a565b5b600082013567ffffffffffffffff81111561a1005761a0ff61970f565b5b61a10c8482850161a0ac565b91505092915050565b60008060006040848603121561a12e5761a12d61970a565b5b600084013567ffffffffffffffff81111561a14c5761a14b61970f565b5b61a15886828701619893565b9350935050602061a16b86828701619d47565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61a1aa81619d9c565b82525050565b600061a1bc838361a1a1565b60208301905092915050565b6000602082019050919050565b600061a1e08261a175565b61a1ea818561a180565b935061a1f58361a191565b8060005b8381101561a22657815161a20d888261a1b0565b975061a2188361a1c8565b92505060018101905061a1f9565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600061a26b8383619a49565b60208301905092915050565b6000602082019050919050565b600061a28f8261a233565b61a299818561a23e565b935061a2a48361a24f565b8060005b8381101561a2d557815161a2bc888261a25f565b975061a2c78361a277565b92505060018101905061a2a8565b5085935050505092915050565b6000604082019050818103600083015261a2fc818561a1d5565b9050818103602083015261a310818461a284565b90509392505050565b61a32281619b17565b82525050565b600060208201905061a33d600083018461a319565b92915050565b60008060006060848603121561a35c5761a35b61970a565b5b600061a36a86828701619782565b935050602061a37b8682870161975d565b925050604061a38c868287016197b8565b9150509250925092565b6000600b9050919050565b600081905092915050565b6000819050919050565b600061a3c28383619a2e565b60208301905092915050565b6000602082019050919050565b61a3e48161a396565b61a3ee818461a3a1565b925061a3f98261a3ac565b8060005b8381101561a42a57815161a411878261a3b6565b965061a41c8361a3ce565b92505060018101905061a3fd565b505050505050565b60006101608201905061a448600083018461a3db565b92915050565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b600061a4bb602e8361a44e565b915061a4c68261a45f565b604082019050919050565b6000602082019050818103600083015261a4ea8161a4ae565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b600185111561a5775780860481111561a5535761a55261a4f1565b5b600185161561a5625780820291505b808102905061a5708561a520565b945061a537565b94509492505050565b60008261a590576001905061a64c565b8161a59e576000905061a64c565b816001811461a5b4576002811461a5be5761a5ed565b600191505061a64c565b60ff84111561a5d05761a5cf61a4f1565b5b8360020a91508482111561a5e75761a5e661a4f1565b5b5061a64c565b5060208310610133831016604e8410600b841016171561a6225782820a90508381111561a61d5761a61c61a4f1565b5b61a64c565b61a62f848484600161a52d565b9250905081840481111561a6465761a64561a4f1565b5b81810290505b9392505050565b600061a65e82619797565b915061a66983619b17565b925061a6967fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461a580565b905092915050565b600061a6a982619797565b915061a6b483619797565b925082820261a6c281619797565b9150828204841483151761a6d95761a6d861a4f1565b5b5092915050565b6000819050919050565b6000819050919050565b600061a70f61a70a61a7058461a6e0565b61a6ea565b619b17565b9050919050565b61a71f8161a6f4565b82525050565b600060208201905061a73a600083018461a716565b92915050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b600061a776601f8361a44e565b915061a7818261a740565b602082019050919050565b6000602082019050818103600083015261a7a58161a769565b9050919050565b7f6f6e6c794b65657065723a2063616c6c6572206973206e6f742061206b65657060008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b600061a80860228361a44e565b915061a8138261a7ac565b604082019050919050565b6000602082019050818103600083015261a8378161a7fb565b9050919050565b7f726f6c6c6f7665723a20636f6e74726163742070617573656400000000000000600082015250565b600061a87460198361a44e565b915061a87f8261a83e565b602082019050919050565b6000602082019050818103600083015261a8a38161a867565b9050919050565b7f726f6c6c6f7665723a20746f6f206561726c7900000000000000000000000000600082015250565b600061a8e060138361a44e565b915061a8eb8261a8aa565b602082019050919050565b6000602082019050818103600083015261a90f8161a8d3565b9050919050565b7f726f6c6c6f7665723a20706c6561736520736574746c65206c61737420726f7560008201527f6e64206669727374000000000000000000000000000000000000000000000000602082015250565b600061a97260288361a44e565b915061a97d8261a916565b604082019050919050565b6000602082019050818103600083015261a9a18161a965565b9050919050565b600061a9b382619b17565b915061a9be83619b17565b9250828201905060ff81111561a9d75761a9d661a4f1565b5b92915050565b600061a9e882619797565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361aa1a5761aa1961a4f1565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061aa9b57607f821691505b60208210810361aaae5761aaad61aa54565b5b50919050565b7f736574746c653a20657870697279206d75737420626520696e2074686520706160008201527f7374000000000000000000000000000000000000000000000000000000000000602082015250565b600061ab1060228361a44e565b915061ab1b8261aab4565b604082019050919050565b6000602082019050818103600083015261ab3f8161ab03565b9050919050565b7f736574746c653a20616c726561647920736574746c6564207468697320726f7560008201527f6e64000000000000000000000000000000000000000000000000000000000000602082015250565b600061aba260228361a44e565b915061abad8261ab46565b604082019050919050565b6000602082019050818103600083015261abd18161ab95565b9050919050565b600061abe382619d9c565b915061abee83619d9c565b925082820390508181126000841216828213600085121516171561ac155761ac1461a4f1565b5b92915050565b600061ac2682619d9c565b91507f8000000000000000000000000000000000000000000000000000000000000000820361ac585761ac5761a4f1565b5b816000039050919050565b600061ac6e82619797565b915061ac7983619797565b925082820390508181111561ac915761ac9061a4f1565b5b92915050565b600061aca282619797565b915061acad83619797565b925082820190508082111561acc55761acc461a4f1565b5b92915050565b600060408201905061ace0600083018561985a565b61aced602083018461985a565b9392505050565b600060408201905061ad09600083018561a319565b61ad16602083018461985a565b9392505050565b7f7365744f7261636c653a20756e6465726c79696e67206d75737420616c72656160008201527f6479206265206163746976650000000000000000000000000000000000000000602082015250565b600061ad79602c8361a44e565b915061ad848261ad1d565b604082019050919050565b6000602082019050818103600083015261ada88161ad6c565b9050919050565b600061adba82619d9c565b915061adc583619d9c565b92508282019050828112156000831216838212600084121516171561aded5761adec61a4f1565b5b92915050565b7f77697468647261773a20616d6f756e74206d757374206265203e203000000000600082015250565b600061ae29601c8361a44e565b915061ae348261adf3565b602082019050919050565b6000602082019050818103600083015261ae588161ae1c565b9050919050565b7f77697468647261773a20616d6f756e74203e2062616c616e6365000000000000600082015250565b600061ae95601a8361a44e565b915061aea08261ae5f565b602082019050919050565b6000602082019050818103600083015261aec48161ae88565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061af0582619797565b915061af1083619797565b92508261af205761af1f61aecb565b5b828204905092915050565b7f77697468647261773a206d617267696e20636865636b206661696c6564000000600082015250565b600061af61601d8361a44e565b915061af6c8261af2b565b602082019050919050565b6000602082019050818103600083015261af908161af54565b9050919050565b7f6c69717569646174653a207573657220686173206e6f20706f736974696f6e73600082015250565b600061afcd60208361a44e565b915061afd88261af97565b602082019050919050565b6000602082019050818103600083015261affc8161afc0565b9050919050565b7f6c69717569646174653a207573657220706173736573206d617267696e20636860008201527f65636b0000000000000000000000000000000000000000000000000000000000602082015250565b600061b05f60238361a44e565b915061b06a8261b003565b604082019050919050565b6000602082019050818103600083015261b08e8161b052565b9050919050565b7f6c69717569646174653a2063616e6e6f74206c697175696461746520796f757260008201527f73656c6600000000000000000000000000000000000000000000000000000000602082015250565b600061b0f160248361a44e565b915061b0fc8261b095565b604082019050919050565b6000602082019050818103600083015261b1208161b0e4565b9050919050565b7f6164644b65657065723a20616c72656164792061206b65657065720000000000600082015250565b600061b15d601b8361a44e565b915061b1688261b127565b602082019050919050565b6000602082019050818103600083015261b18c8161b150565b9050919050565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b600061b1ef602c8361a44e565b915061b1fa8261b193565b604082019050919050565b6000602082019050818103600083015261b21e8161b1e2565b9050919050565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b600061b281602c8361a44e565b915061b28c8261b225565b604082019050919050565b6000602082019050818103600083015261b2b08161b274565b9050919050565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b600061b31360388361a44e565b915061b31e8261b2b7565b604082019050919050565b6000602082019050818103600083015261b3428161b306565b9050919050565b7f72656d6f766546726f6d57686974656c6973743a206e6f7420696e207768697460008201527f656c697374000000000000000000000000000000000000000000000000000000602082015250565b600061b3a560258361a44e565b915061b3b08261b349565b604082019050919050565b6000602082019050818103600083015261b3d48161b398565b9050919050565b7f7365744d696e5175616e746974793a20756e6465726c79696e67206d7573742060008201527f616c726561647920626520616374697665000000000000000000000000000000602082015250565b600061b43760318361a44e565b915061b4428261b3db565b604082019050919050565b6000602082019050818103600083015261b4668161b42a565b9050919050565b7f7365744d696e5175616e746974793a206d696e207175616e74697479206d757360008201527f74206265203e2030000000000000000000000000000000000000000000000000602082015250565b600061b4c960288361a44e565b915061b4d48261b46d565b604082019050919050565b6000602082019050818103600083015261b4f88161b4bc565b9050919050565b7f736574496e737572616e636546756e643a206d757374206265206e657720616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b600061b55b60258361a44e565b915061b5668261b4ff565b604082019050919050565b6000602082019050818103600083015261b58a8161b54e565b9050919050565b7f6765744d617267696e3a20747261646572206d7573742062652062757965722060008201527f6f722073656c6c65720000000000000000000000000000000000000000000000602082015250565b600061b5ed60298361a44e565b915061b5f88261b591565b604082019050919050565b6000602082019050818103600083015261b61c8161b5e0565b9050919050565b7f616464546f57686974656c6973743a20616c726561647920696e20776869746560008201527f6c69737400000000000000000000000000000000000000000000000000000000602082015250565b600061b67f60248361a44e565b915061b68a8261b623565b604082019050919050565b6000602082019050818103600083015261b6ae8161b672565b9050919050565b7f616464506f736974696f6e3a2074726164655072696365206d7573742062652060008201527f3e20300000000000000000000000000000000000000000000000000000000000602082015250565b600061b71160238361a44e565b915061b71c8261b6b5565b604082019050919050565b6000602082019050818103600083015261b7408161b704565b9050919050565b7f616464506f736974696f6e3a207175616e74697479206d757374206265203e2060008201527f3000000000000000000000000000000000000000000000000000000000000000602082015250565b600061b7a360218361a44e565b915061b7ae8261b747565b604082019050919050565b6000602082019050818103600083015261b7d28161b796565b9050919050565b7f616464506f736974696f6e3a206e6f206f7261636c6520666f7220756e64657260008201527f6c79696e67000000000000000000000000000000000000000000000000000000602082015250565b600061b83560258361a44e565b915061b8408261b7d9565b604082019050919050565b6000602082019050818103600083015261b8648161b828565b9050919050565b7f616464506f736974696f6e3a2063616e6e6f7420656e746572206120706f736960008201527f74696f6e207769746820796f757273656c660000000000000000000000000000602082015250565b600061b8c760328361a44e565b915061b8d28261b86b565b604082019050919050565b6000602082019050818103600083015261b8f68161b8ba565b9050919050565b600b811061b90a57600080fd5b50565b60008135905061b91c8161b8fd565b92915050565b60006020828403121561b9385761b93761970a565b5b600061b9468482850161b90d565b91505092915050565b7f616464506f736974696f6e3a206e6f20737472696b6520666f7220756e64657260008201527f6c79696e67000000000000000000000000000000000000000000000000000000602082015250565b600061b9ab60258361a44e565b915061b9b68261b94f565b604082019050919050565b6000602082019050818103600083015261b9da8161b99e565b9050919050565b600080fd5b600080fd5b600080fd5b6000808335600160200384360303811261ba0d5761ba0c61b9e1565b5b80840192508235915067ffffffffffffffff82111561ba2f5761ba2e61b9e6565b5b60208301925060018202360383131561ba4b5761ba4a61b9eb565b5b509250929050565b60006020828403121561ba695761ba6861970a565b5b600061ba7784828501619d47565b91505092915050565b7f616464506f736974696f6e3a2062656c6f77206d696e207175616e7469747900600082015250565b600061bab6601f8361a44e565b915061bac18261ba80565b602082019050919050565b6000602082019050818103600083015261bae58161baa9565b9050919050565b7f616464506f736974696f6e3a2062757965722063616e6e6f742070617920666560008201527f6573000000000000000000000000000000000000000000000000000000000000602082015250565b600061bb4860228361a44e565b915061bb538261baec565b604082019050919050565b6000602082019050818103600083015261bb778161bb3b565b9050919050565b7f616464506f736974696f6e3a2073656c6c65722063616e6e6f7420706179206660008201527f6565730000000000000000000000000000000000000000000000000000000000602082015250565b600061bbda60238361a44e565b915061bbe58261bb7e565b604082019050919050565b6000602082019050818103600083015261bc098161bbcd565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261bc727fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261bc35565b61bc7c868361bc35565b95508019841693508086168417925050509392505050565b600061bcaf61bcaa61bca584619797565b61a6ea565b619797565b9050919050565b6000819050919050565b61bcc98361bc94565b61bcdd61bcd58261bcb6565b84845461bc42565b825550505050565b600090565b61bcf261bce5565b61bcfd81848461bcc0565b505050565b5b8181101561bd215761bd1660008261bcea565b60018101905061bd03565b5050565b601f82111561bd665761bd378161bc10565b61bd408461bc25565b8101602085101561bd4f578190505b61bd6361bd5b8561bc25565b83018261bd02565b50505b505050565b600082821c905092915050565b600061bd896000198460080261bd6b565b1980831691505092915050565b600061bda2838361bd78565b9150826002028217905092915050565b61bdbb8261998f565b67ffffffffffffffff81111561bdd45761bdd3619eac565b5b61bdde825461aa83565b61bde982828561bd25565b600060209050601f83116001811461be1c576000841561be0a578287015190505b61be14858261bd96565b86555061be7c565b601f19841661be2a8661bc10565b60005b8281101561be525784890151825560018201915060208501945060208101905061be2d565b8683101561be6f578489015161be6b601f89168261bd78565b8355505b6001600288020188555050505b505050505050565b600061ffff82169050919050565b600061be9d8261be84565b915061ffff820361beb15761beb061a4f1565b5b600182019050919050565b7f616464506f736974696f6e3a206275796572206661696c6564206d617267696e60008201527f20636865636b0000000000000000000000000000000000000000000000000000602082015250565b600061bf1860268361a44e565b915061bf238261bebc565b604082019050919050565b6000602082019050818103600083015261bf478161bf0b565b9050919050565b7f616464506f736974696f6e3a2073656c6c6572206661696c6564206d6172676960008201527f6e20636865636b00000000000000000000000000000000000000000000000000602082015250565b600061bfaa60278361a44e565b915061bfb58261bf4e565b604082019050919050565b6000602082019050818103600083015261bfd98161bf9d565b9050919050565b600061bfec838561a44e565b935061bff9838584619f58565b61c002836199d5565b840190509392505050565b61c01681619af6565b82525050565b61c02581619aae565b82525050565b600060e082019050818103600083015261c046818a8c61bfe0565b905061c055602083018961985a565b61c062604083018861985a565b61c06f6060830187619db5565b61c07c608083018661c00d565b61c08960a083018561c01c565b61c09660c083018461985a565b9998505050505050505050565b7f7769746864726177416c6c3a20656d7074792062616c616e6365000000000000600082015250565b600061c0d9601a8361a44e565b915061c0e48261c0a3565b602082019050919050565b6000602082019050818103600083015261c1088161c0cc565b9050919050565b7f7769746864726177416c6c3a206d617267696e20636865636b206661696c6564600082015250565b600061c14560208361a44e565b915061c1508261c10f565b602082019050919050565b6000602082019050818103600083015261c1748161c138565b9050919050565b7f7365744d696e4d617267696e506572633a206d757374206265203c3d2031302a60008201527f2a34000000000000000000000000000000000000000000000000000000000000602082015250565b600061c1d760228361a44e565b915061c1e28261c17b565b604082019050919050565b6000602082019050818103600083015261c2068161c1ca565b9050919050565b61c21681619b17565b811461c22157600080fd5b50565b60008151905061c2338161c20d565b92915050565b60006020828403121561c24f5761c24e61970a565b5b600061c25d8482850161c224565b91505092915050565b7f6465706f7369743a2060616d6f756e7460206d757374206265203e2030000000600082015250565b600061c29c601d8361a44e565b915061c2a78261c266565b602082019050919050565b6000602082019050818103600083015261c2cb8161c28f565b9050919050565b7f6465706f7369743a2065786365656473206d6178696d756d2062616c616e636560008201527f2063617000000000000000000000000000000000000000000000000000000000602082015250565b600061c32e60248361a44e565b915061c3398261c2d2565b604082019050919050565b6000602082019050818103600083015261c35d8161c321565b9050919050565b7f736574466565526563697069656e743a206d757374206265206e65772066656560008201527f20726563697069656e7400000000000000000000000000000000000000000000602082015250565b600061c3c0602a8361a44e565b915061c3cb8261c364565b604082019050919050565b6000602082019050818103600083015261c3ef8161c3b3565b9050919050565b7f72656d6f76654b65657065723a206e6f742061206b6565706572000000000000600082015250565b600061c42c601a8361a44e565b915061c4378261c3f6565b602082019050919050565b6000602082019050818103600083015261c45b8161c41f565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061c4be60268361a44e565b915061c4c98261c462565b604082019050919050565b6000602082019050818103600083015261c4ed8161c4b1565b9050919050565b7f7365744d617842616c616e63654361703a206d757374206265203e2030000000600082015250565b600061c52a601d8361a44e565b915061c5358261c4f4565b602082019050919050565b6000602082019050818103600083015261c5598161c51d565b9050919050565b7f6163746976617465556e6465726c79696e673a20756e6465726c79696e67206d60008201527f757374206e6f7420796574206265206163746976650000000000000000000000602082015250565b600061c5bc60358361a44e565b915061c5c78261c560565b604082019050919050565b6000602082019050818103600083015261c5eb8161c5af565b9050919050565b600060608201905061c607600083018661c00d565b61c6146020830185619e7d565b61c621604083018461985a565b949350505050565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b600061c685602b8361a44e565b915061c6908261c629565b604082019050919050565b6000602082019050818103600083015261c6b48161c678565b9050919050565b7f6e6577556e6465726c79696e673a20756e6465726c79696e6720616c7265616460008201527f7920616374697665000000000000000000000000000000000000000000000000602082015250565b600061c71760288361a44e565b915061c7228261c6bb565b604082019050919050565b6000602082019050818103600083015261c7468161c70a565b9050919050565b7f6e6577556e6465726c79696e673a206d6178206e6f74696f6e616c206d75737460008201527f206265203e203000000000000000000000000000000000000000000000000000602082015250565b600061c7a960278361a44e565b915061c7b48261c74d565b604082019050919050565b6000602082019050818103600083015261c7d88161c79c565b9050919050565b7f676574537472696b654d656e753a2065787069727920696e207468652070617360008201527f7400000000000000000000000000000000000000000000000000000000000000602082015250565b600061c83b60218361a44e565b915061c8468261c7df565b604082019050919050565b6000602082019050818103600083015261c86a8161c82e565b9050919050565b600061c87c82619b17565b915061c88783619b17565b9250828203905060ff81111561c8a05761c89f61a4f1565b5b92915050565b7f676574537472696b654d656e753a2053706f7420707269636520746f6f20736d60008201527f616c6c0000000000000000000000000000000000000000000000000000000000602082015250565b600061c90260238361a44e565b915061c90d8261c8a6565b604082019050919050565b6000602082019050818103600083015261c9318161c8f5565b9050919050565b7f676574537472696b654d656e753a2053706f7420707269636520746f6f206c6160008201527f7267650000000000000000000000000000000000000000000000000000000000602082015250565b600061c99460238361a44e565b915061c99f8261c938565b604082019050919050565b6000602082019050818103600083015261c9c38161c987565b9050919050565b7f67657453706f743a206d697373696e67206f7261636c65000000000000000000600082015250565b600061ca0060178361a44e565b915061ca0b8261c9ca565b602082019050919050565b6000602082019050818103600083015261ca2f8161c9f3565b9050919050565b600069ffffffffffffffffffff82169050919050565b61ca558161ca36565b811461ca6057600080fd5b50565b60008151905061ca728161ca4c565b92915050565b60008151905061ca87816197a1565b92915050565b6000806000806080858703121561caa75761caa661970a565b5b600061cab58782880161ca63565b945050602061cac68782880161ca78565b935050604061cad78782880161ca78565b925050606061cae88782880161c224565b91505092959194509250565b7f676574496e697469616c4d617267696e3a20747261646572206d75737420626560008201527f206275796572206f722073656c6c657200000000000000000000000000000000602082015250565b600061cb5060308361a44e565b915061cb5b8261caf4565b604082019050919050565b6000602082019050818103600083015261cb7f8161cb43565b9050919050565b600061cb9182619d9c565b915061cb9c83619d9c565b925082820261cbaa81619d9c565b91507f8000000000000000000000000000000000000000000000000000000000000000841460008412161561cbe25761cbe161a4f1565b5b828205841483151761cbf75761cbf661a4f1565b5b5092915050565b600061cc0982619d9c565b915061cc1483619d9c565b92508261cc245761cc2361aecb565b5b600160000383147f80000000000000000000000000000000000000000000000000000000000000008314161561cc5d5761cc5c61a4f1565b5b828205905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b600061cc9e60208361a44e565b915061cca98261cc68565b602082019050919050565b6000602082019050818103600083015261cccd8161cc91565b9050919050565b600060408201905061cce96000830185619e7d565b61ccf6602083018461985a565b9392505050565b600061cd088261be84565b91506000820361cd1b5761cd1a61a4f1565b5b600182039050919050565b7f6c69717569646174653a20757365722063616e6e6f742070617920726577617260008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b600061cd8260218361a44e565b915061cd8d8261cd26565b604082019050919050565b6000602082019050818103600083015261cdb18161cd75565b9050919050565b61cdc18161a033565b811461cdcc57600080fd5b50565b60008151905061cdde8161cdb8565b92915050565b60006020828403121561cdfa5761cdf961970a565b5b600061ce088482850161cdcf565b91505092915050565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b600061ce6d602e8361a44e565b915061ce788261ce11565b604082019050919050565b6000602082019050818103600083015261ce9c8161ce60565b9050919050565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b600061ceff60298361a44e565b915061cf0a8261cea3565b604082019050919050565b6000602082019050818103600083015261cf2e8161cef2565b9050919050565b60008160f81b9050919050565b600061cf4d8261cf35565b9050919050565b600061cf5f8261cf42565b9050919050565b61cf7761cf7282619a3d565b61cf54565b82525050565b61cf8e61cf8982619af6565b61cf42565b82525050565b6000819050919050565b61cfaf61cfaa82619797565b61cf94565b82525050565b600061cfc1828761cf66565b60018201915061cfd1828661cf7d565b60018201915061cfe1828561cf9e565b60208201915061cff1828461cf9e565b60208201915081905095945050505050565b7f6765744d61726b3a206d697373696e67206f7261636c65000000000000000000600082015250565b600061d03960178361a44e565b915061d0448261d003565b602082019050919050565b6000602082019050818103600083015261d0688161d02c565b9050919050565b600060408201905061d0846000830185619db5565b61d091602083018461a319565b9392505050565b7f676574496e697469616c4d617267696e3a206f7074696f6e206973206578706960008201527f7265640000000000000000000000000000000000000000000000000000000000602082015250565b600061d0f460238361a44e565b915061d0ff8261d098565b604082019050919050565b6000602082019050818103600083015261d1238161d0e7565b9050919050565b7f6765744d61696e7461696e656e63654d617267696e3a206f7074696f6e20697360008201527f2065787069726564000000000000000000000000000000000000000000000000602082015250565b600061d18660288361a44e565b915061d1918261d12a565b604082019050919050565b6000602082019050818103600083015261d1b58161d179565b9050919050565b600060608201905061d1d16000830186619e7d565b61d1de6020830185619e7d565b61d1eb604083018461985a565b949350505050565b600061d1fe82619797565b915061d20983619797565b92508261d2195761d21861aecb565b5b828206905092915050565b60008151905061d23381619d30565b92915050565b60006020828403121561d24f5761d24e61970a565b5b600061d25d8482850161d224565b91505092915050565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b600061d2c2602a8361a44e565b915061d2cd8261d266565b604082019050919050565b6000602082019050818103600083015261d2f18161d2b5565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b600061d383602d8361a44e565b915061d38e8261d327565b604082019050919050565b6000602082019050818103600083015261d3b28161d376565b9050919050565b7f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60008201527f6e74726163740000000000000000000000000000000000000000000000000000602082015250565b600061d41560268361a44e565b915061d4208261d3b9565b604082019050919050565b6000602082019050818103600083015261d4448161d408565b9050919050565b600081519050919050565b600081905092915050565b600061d46c8261d44b565b61d476818561d456565b935061d4868185602086016199ab565b80840191505092915050565b600061d49e828461d461565b915081905092915050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b600061d50560268361a44e565b915061d5108261d4a9565b604082019050919050565b6000602082019050818103600083015261d5348161d4f8565b9050919050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b600061d571601d8361a44e565b915061d57c8261d53b565b602082019050919050565b6000602082019050818103600083015261d5a08161d564565b9050919050565b600061d5b28261998f565b61d5bc818561a44e565b935061d5cc8185602086016199ab565b61d5d5816199d5565b840191505092915050565b6000602082019050818103600083015261d5fa818461d5a7565b90509291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212200ad461962149a9b28bfa67db2f38f59dc9b5764d254a7e39eb65ca98100a25d264736f6c63430008110033",
}

// MarginABI is the input ABI used to generate the binding from.
// Deprecated: Use MarginMetaData.ABI instead.
var MarginABI = MarginMetaData.ABI

// MarginBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarginMetaData.Bin instead.
var MarginBin = MarginMetaData.Bin

// DeployMargin deploys a new Ethereum contract, binding an instance of Margin to it.
func DeployMargin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Margin, error) {
	parsed, err := MarginMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarginBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Margin{MarginCaller: MarginCaller{contract: contract}, MarginTransactor: MarginTransactor{contract: contract}, MarginFilterer: MarginFilterer{contract: contract}}, nil
}

// Margin is an auto generated Go binding around an Ethereum contract.
type Margin struct {
	MarginCaller     // Read-only binding to the contract
	MarginTransactor // Write-only binding to the contract
	MarginFilterer   // Log filterer for contract events
}

// MarginCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarginCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarginTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarginTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarginFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarginFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarginSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarginSession struct {
	Contract     *Margin           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarginCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarginCallerSession struct {
	Contract *MarginCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarginTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarginTransactorSession struct {
	Contract     *MarginTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarginRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarginRaw struct {
	Contract *Margin // Generic contract binding to access the raw methods on
}

// MarginCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarginCallerRaw struct {
	Contract *MarginCaller // Generic read-only contract binding to access the raw methods on
}

// MarginTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarginTransactorRaw struct {
	Contract *MarginTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMargin creates a new instance of Margin, bound to a specific deployed contract.
func NewMargin(address common.Address, backend bind.ContractBackend) (*Margin, error) {
	contract, err := bindMargin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Margin{MarginCaller: MarginCaller{contract: contract}, MarginTransactor: MarginTransactor{contract: contract}, MarginFilterer: MarginFilterer{contract: contract}}, nil
}

// NewMarginCaller creates a new read-only instance of Margin, bound to a specific deployed contract.
func NewMarginCaller(address common.Address, caller bind.ContractCaller) (*MarginCaller, error) {
	contract, err := bindMargin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarginCaller{contract: contract}, nil
}

// NewMarginTransactor creates a new write-only instance of Margin, bound to a specific deployed contract.
func NewMarginTransactor(address common.Address, transactor bind.ContractTransactor) (*MarginTransactor, error) {
	contract, err := bindMargin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarginTransactor{contract: contract}, nil
}

// NewMarginFilterer creates a new log filterer instance of Margin, bound to a specific deployed contract.
func NewMarginFilterer(address common.Address, filterer bind.ContractFilterer) (*MarginFilterer, error) {
	contract, err := bindMargin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarginFilterer{contract: contract}, nil
}

// bindMargin binds a generic wrapper to an already deployed contract.
func bindMargin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarginABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Margin *MarginRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Margin.Contract.MarginCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Margin *MarginRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Margin.Contract.MarginTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Margin *MarginRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Margin.Contract.MarginTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Margin *MarginCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Margin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Margin *MarginTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Margin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Margin *MarginTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Margin.Contract.contract.Transact(opts, method, params...)
}

// ActiveExpiry is a free data retrieval call binding the contract method 0x03c4d3d6.
//
// Solidity: function activeExpiry() view returns(uint256)
func (_Margin *MarginCaller) ActiveExpiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "activeExpiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveExpiry is a free data retrieval call binding the contract method 0x03c4d3d6.
//
// Solidity: function activeExpiry() view returns(uint256)
func (_Margin *MarginSession) ActiveExpiry() (*big.Int, error) {
	return _Margin.Contract.ActiveExpiry(&_Margin.CallOpts)
}

// ActiveExpiry is a free data retrieval call binding the contract method 0x03c4d3d6.
//
// Solidity: function activeExpiry() view returns(uint256)
func (_Margin *MarginCallerSession) ActiveExpiry() (*big.Int, error) {
	return _Margin.Contract.ActiveExpiry(&_Margin.CallOpts)
}

// CheckMargin is a free data retrieval call binding the contract method 0x1febd86d.
//
// Solidity: function checkMargin(address user, bool useInitialMargin) view returns(int256, bool)
func (_Margin *MarginCaller) CheckMargin(opts *bind.CallOpts, user common.Address, useInitialMargin bool) (*big.Int, bool, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "checkMargin", user, useInitialMargin)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// CheckMargin is a free data retrieval call binding the contract method 0x1febd86d.
//
// Solidity: function checkMargin(address user, bool useInitialMargin) view returns(int256, bool)
func (_Margin *MarginSession) CheckMargin(user common.Address, useInitialMargin bool) (*big.Int, bool, error) {
	return _Margin.Contract.CheckMargin(&_Margin.CallOpts, user, useInitialMargin)
}

// CheckMargin is a free data retrieval call binding the contract method 0x1febd86d.
//
// Solidity: function checkMargin(address user, bool useInitialMargin) view returns(int256, bool)
func (_Margin *MarginCallerSession) CheckMargin(user common.Address, useInitialMargin bool) (*big.Int, bool, error) {
	return _Margin.Contract.CheckMargin(&_Margin.CallOpts, user, useInitialMargin)
}

// CheckMarginBatch is a free data retrieval call binding the contract method 0xa665a637.
//
// Solidity: function checkMarginBatch(address[] users, bool useInitialMargin) view returns(int256[] diffs, bool[] satisfieds)
func (_Margin *MarginCaller) CheckMarginBatch(opts *bind.CallOpts, users []common.Address, useInitialMargin bool) (struct {
	Diffs      []*big.Int
	Satisfieds []bool
}, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "checkMarginBatch", users, useInitialMargin)

	outstruct := new(struct {
		Diffs      []*big.Int
		Satisfieds []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Diffs = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Satisfieds = *abi.ConvertType(out[1], new([]bool)).(*[]bool)

	return *outstruct, err

}

// CheckMarginBatch is a free data retrieval call binding the contract method 0xa665a637.
//
// Solidity: function checkMarginBatch(address[] users, bool useInitialMargin) view returns(int256[] diffs, bool[] satisfieds)
func (_Margin *MarginSession) CheckMarginBatch(users []common.Address, useInitialMargin bool) (struct {
	Diffs      []*big.Int
	Satisfieds []bool
}, error) {
	return _Margin.Contract.CheckMarginBatch(&_Margin.CallOpts, users, useInitialMargin)
}

// CheckMarginBatch is a free data retrieval call binding the contract method 0xa665a637.
//
// Solidity: function checkMarginBatch(address[] users, bool useInitialMargin) view returns(int256[] diffs, bool[] satisfieds)
func (_Margin *MarginCallerSession) CheckMarginBatch(users []common.Address, useInitialMargin bool) (struct {
	Diffs      []*big.Int
	Satisfieds []bool
}, error) {
	return _Margin.Contract.CheckMarginBatch(&_Margin.CallOpts, users, useInitialMargin)
}

// CurRound is a free data retrieval call binding the contract method 0xb3f75843.
//
// Solidity: function curRound() view returns(uint8)
func (_Margin *MarginCaller) CurRound(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "curRound")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CurRound is a free data retrieval call binding the contract method 0xb3f75843.
//
// Solidity: function curRound() view returns(uint8)
func (_Margin *MarginSession) CurRound() (uint8, error) {
	return _Margin.Contract.CurRound(&_Margin.CallOpts)
}

// CurRound is a free data retrieval call binding the contract method 0xb3f75843.
//
// Solidity: function curRound() view returns(uint8)
func (_Margin *MarginCallerSession) CurRound() (uint8, error) {
	return _Margin.Contract.CurRound(&_Margin.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Margin *MarginCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Margin *MarginSession) FeeRecipient() (common.Address, error) {
	return _Margin.Contract.FeeRecipient(&_Margin.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Margin *MarginCallerSession) FeeRecipient() (common.Address, error) {
	return _Margin.Contract.FeeRecipient(&_Margin.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Margin *MarginCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Margin *MarginSession) GetBalance() (*big.Int, error) {
	return _Margin.Contract.GetBalance(&_Margin.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Margin *MarginCallerSession) GetBalance() (*big.Int, error) {
	return _Margin.Contract.GetBalance(&_Margin.CallOpts)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address user) view returns(uint256)
func (_Margin *MarginCaller) GetBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "getBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address user) view returns(uint256)
func (_Margin *MarginSession) GetBalanceOf(user common.Address) (*big.Int, error) {
	return _Margin.Contract.GetBalanceOf(&_Margin.CallOpts, user)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address user) view returns(uint256)
func (_Margin *MarginCallerSession) GetBalanceOf(user common.Address) (*big.Int, error) {
	return _Margin.Contract.GetBalanceOf(&_Margin.CallOpts, user)
}

// GetCollateralDecimals is a free data retrieval call binding the contract method 0xa9520b00.
//
// Solidity: function getCollateralDecimals() view returns(uint8)
func (_Margin *MarginCaller) GetCollateralDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "getCollateralDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCollateralDecimals is a free data retrieval call binding the contract method 0xa9520b00.
//
// Solidity: function getCollateralDecimals() view returns(uint8)
func (_Margin *MarginSession) GetCollateralDecimals() (uint8, error) {
	return _Margin.Contract.GetCollateralDecimals(&_Margin.CallOpts)
}

// GetCollateralDecimals is a free data retrieval call binding the contract method 0xa9520b00.
//
// Solidity: function getCollateralDecimals() view returns(uint8)
func (_Margin *MarginCallerSession) GetCollateralDecimals() (uint8, error) {
	return _Margin.Contract.GetCollateralDecimals(&_Margin.CallOpts)
}

// GetMargin is a free data retrieval call binding the contract method 0x7a594a7c.
//
// Solidity: function getMargin(address user, bool useInitialMargin) view returns(uint256)
func (_Margin *MarginCaller) GetMargin(opts *bind.CallOpts, user common.Address, useInitialMargin bool) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "getMargin", user, useInitialMargin)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMargin is a free data retrieval call binding the contract method 0x7a594a7c.
//
// Solidity: function getMargin(address user, bool useInitialMargin) view returns(uint256)
func (_Margin *MarginSession) GetMargin(user common.Address, useInitialMargin bool) (*big.Int, error) {
	return _Margin.Contract.GetMargin(&_Margin.CallOpts, user, useInitialMargin)
}

// GetMargin is a free data retrieval call binding the contract method 0x7a594a7c.
//
// Solidity: function getMargin(address user, bool useInitialMargin) view returns(uint256)
func (_Margin *MarginCallerSession) GetMargin(user common.Address, useInitialMargin bool) (*big.Int, error) {
	return _Margin.Contract.GetMargin(&_Margin.CallOpts, user, useInitialMargin)
}

// GetPayoff is a free data retrieval call binding the contract method 0x23880e78.
//
// Solidity: function getPayoff(address user, bool onlyLoss) view returns(int256)
func (_Margin *MarginCaller) GetPayoff(opts *bind.CallOpts, user common.Address, onlyLoss bool) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "getPayoff", user, onlyLoss)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPayoff is a free data retrieval call binding the contract method 0x23880e78.
//
// Solidity: function getPayoff(address user, bool onlyLoss) view returns(int256)
func (_Margin *MarginSession) GetPayoff(user common.Address, onlyLoss bool) (*big.Int, error) {
	return _Margin.Contract.GetPayoff(&_Margin.CallOpts, user, onlyLoss)
}

// GetPayoff is a free data retrieval call binding the contract method 0x23880e78.
//
// Solidity: function getPayoff(address user, bool onlyLoss) view returns(int256)
func (_Margin *MarginCallerSession) GetPayoff(user common.Address, onlyLoss bool) (*big.Int, error) {
	return _Margin.Contract.GetPayoff(&_Margin.CallOpts, user, onlyLoss)
}

// GetPositions is a free data retrieval call binding the contract method 0x062c4878.
//
// Solidity: function getPositions(uint8 underlying) view returns((string,address,address,uint256,uint256,(bool,uint8,uint256,uint256,uint8,uint8))[])
func (_Margin *MarginCaller) GetPositions(opts *bind.CallOpts, underlying uint8) ([]DerivativeOrder, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "getPositions", underlying)

	if err != nil {
		return *new([]DerivativeOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]DerivativeOrder)).(*[]DerivativeOrder)

	return out0, err

}

// GetPositions is a free data retrieval call binding the contract method 0x062c4878.
//
// Solidity: function getPositions(uint8 underlying) view returns((string,address,address,uint256,uint256,(bool,uint8,uint256,uint256,uint8,uint8))[])
func (_Margin *MarginSession) GetPositions(underlying uint8) ([]DerivativeOrder, error) {
	return _Margin.Contract.GetPositions(&_Margin.CallOpts, underlying)
}

// GetPositions is a free data retrieval call binding the contract method 0x062c4878.
//
// Solidity: function getPositions(uint8 underlying) view returns((string,address,address,uint256,uint256,(bool,uint8,uint256,uint256,uint8,uint8))[])
func (_Margin *MarginCallerSession) GetPositions(underlying uint8) ([]DerivativeOrder, error) {
	return _Margin.Contract.GetPositions(&_Margin.CallOpts, underlying)
}

// GetStrikes is a free data retrieval call binding the contract method 0xff7f29da.
//
// Solidity: function getStrikes(uint8 underlying) view returns(uint256[11])
func (_Margin *MarginCaller) GetStrikes(opts *bind.CallOpts, underlying uint8) ([11]*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "getStrikes", underlying)

	if err != nil {
		return *new([11]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([11]*big.Int)).(*[11]*big.Int)

	return out0, err

}

// GetStrikes is a free data retrieval call binding the contract method 0xff7f29da.
//
// Solidity: function getStrikes(uint8 underlying) view returns(uint256[11])
func (_Margin *MarginSession) GetStrikes(underlying uint8) ([11]*big.Int, error) {
	return _Margin.Contract.GetStrikes(&_Margin.CallOpts, underlying)
}

// GetStrikes is a free data retrieval call binding the contract method 0xff7f29da.
//
// Solidity: function getStrikes(uint8 underlying) view returns(uint256[11])
func (_Margin *MarginCallerSession) GetStrikes(underlying uint8) ([11]*big.Int, error) {
	return _Margin.Contract.GetStrikes(&_Margin.CallOpts, underlying)
}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_Margin *MarginCaller) Insurance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "insurance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_Margin *MarginSession) Insurance() (common.Address, error) {
	return _Margin.Contract.Insurance(&_Margin.CallOpts)
}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_Margin *MarginCallerSession) Insurance() (common.Address, error) {
	return _Margin.Contract.Insurance(&_Margin.CallOpts)
}

// IsActiveUnderlying is a free data retrieval call binding the contract method 0x2db0788e.
//
// Solidity: function isActiveUnderlying(uint8 ) view returns(bool)
func (_Margin *MarginCaller) IsActiveUnderlying(opts *bind.CallOpts, arg0 uint8) (bool, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "isActiveUnderlying", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveUnderlying is a free data retrieval call binding the contract method 0x2db0788e.
//
// Solidity: function isActiveUnderlying(uint8 ) view returns(bool)
func (_Margin *MarginSession) IsActiveUnderlying(arg0 uint8) (bool, error) {
	return _Margin.Contract.IsActiveUnderlying(&_Margin.CallOpts, arg0)
}

// IsActiveUnderlying is a free data retrieval call binding the contract method 0x2db0788e.
//
// Solidity: function isActiveUnderlying(uint8 ) view returns(bool)
func (_Margin *MarginCallerSession) IsActiveUnderlying(arg0 uint8) (bool, error) {
	return _Margin.Contract.IsActiveUnderlying(&_Margin.CallOpts, arg0)
}

// MaxBalanceCap is a free data retrieval call binding the contract method 0xb40751cc.
//
// Solidity: function maxBalanceCap() view returns(uint256)
func (_Margin *MarginCaller) MaxBalanceCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "maxBalanceCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBalanceCap is a free data retrieval call binding the contract method 0xb40751cc.
//
// Solidity: function maxBalanceCap() view returns(uint256)
func (_Margin *MarginSession) MaxBalanceCap() (*big.Int, error) {
	return _Margin.Contract.MaxBalanceCap(&_Margin.CallOpts)
}

// MaxBalanceCap is a free data retrieval call binding the contract method 0xb40751cc.
//
// Solidity: function maxBalanceCap() view returns(uint256)
func (_Margin *MarginCallerSession) MaxBalanceCap() (*big.Int, error) {
	return _Margin.Contract.MaxBalanceCap(&_Margin.CallOpts)
}

// MinMarginPerc is a free data retrieval call binding the contract method 0xbd6b7982.
//
// Solidity: function minMarginPerc() view returns(uint256)
func (_Margin *MarginCaller) MinMarginPerc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "minMarginPerc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinMarginPerc is a free data retrieval call binding the contract method 0xbd6b7982.
//
// Solidity: function minMarginPerc() view returns(uint256)
func (_Margin *MarginSession) MinMarginPerc() (*big.Int, error) {
	return _Margin.Contract.MinMarginPerc(&_Margin.CallOpts)
}

// MinMarginPerc is a free data retrieval call binding the contract method 0xbd6b7982.
//
// Solidity: function minMarginPerc() view returns(uint256)
func (_Margin *MarginCallerSession) MinMarginPerc() (*big.Int, error) {
	return _Margin.Contract.MinMarginPerc(&_Margin.CallOpts)
}

// MinQuantityPerUnderlying is a free data retrieval call binding the contract method 0xc955bb69.
//
// Solidity: function minQuantityPerUnderlying(uint8 ) view returns(uint256)
func (_Margin *MarginCaller) MinQuantityPerUnderlying(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "minQuantityPerUnderlying", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuantityPerUnderlying is a free data retrieval call binding the contract method 0xc955bb69.
//
// Solidity: function minQuantityPerUnderlying(uint8 ) view returns(uint256)
func (_Margin *MarginSession) MinQuantityPerUnderlying(arg0 uint8) (*big.Int, error) {
	return _Margin.Contract.MinQuantityPerUnderlying(&_Margin.CallOpts, arg0)
}

// MinQuantityPerUnderlying is a free data retrieval call binding the contract method 0xc955bb69.
//
// Solidity: function minQuantityPerUnderlying(uint8 ) view returns(uint256)
func (_Margin *MarginCallerSession) MinQuantityPerUnderlying(arg0 uint8) (*big.Int, error) {
	return _Margin.Contract.MinQuantityPerUnderlying(&_Margin.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Margin *MarginCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Margin *MarginSession) Owner() (common.Address, error) {
	return _Margin.Contract.Owner(&_Margin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Margin *MarginCallerSession) Owner() (common.Address, error) {
	return _Margin.Contract.Owner(&_Margin.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Margin *MarginCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Margin *MarginSession) ProxiableUUID() ([32]byte, error) {
	return _Margin.Contract.ProxiableUUID(&_Margin.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Margin *MarginCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Margin.Contract.ProxiableUUID(&_Margin.CallOpts)
}

// RoundStrikes is a free data retrieval call binding the contract method 0xcc4fdcd0.
//
// Solidity: function roundStrikes(uint8 , uint256 ) view returns(uint256)
func (_Margin *MarginCaller) RoundStrikes(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "roundStrikes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundStrikes is a free data retrieval call binding the contract method 0xcc4fdcd0.
//
// Solidity: function roundStrikes(uint8 , uint256 ) view returns(uint256)
func (_Margin *MarginSession) RoundStrikes(arg0 uint8, arg1 *big.Int) (*big.Int, error) {
	return _Margin.Contract.RoundStrikes(&_Margin.CallOpts, arg0, arg1)
}

// RoundStrikes is a free data retrieval call binding the contract method 0xcc4fdcd0.
//
// Solidity: function roundStrikes(uint8 , uint256 ) view returns(uint256)
func (_Margin *MarginCallerSession) RoundStrikes(arg0 uint8, arg1 *big.Int) (*big.Int, error) {
	return _Margin.Contract.RoundStrikes(&_Margin.CallOpts, arg0, arg1)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Margin *MarginCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Margin.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Margin *MarginSession) Usdc() (common.Address, error) {
	return _Margin.Contract.Usdc(&_Margin.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Margin *MarginCallerSession) Usdc() (common.Address, error) {
	return _Margin.Contract.Usdc(&_Margin.CallOpts)
}

// ActivateUnderlying is a paid mutator transaction binding the contract method 0xfe6bb409.
//
// Solidity: function activateUnderlying(uint8 underlying, address oracle, uint256 minQuantity) returns()
func (_Margin *MarginTransactor) ActivateUnderlying(opts *bind.TransactOpts, underlying uint8, oracle common.Address, minQuantity *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "activateUnderlying", underlying, oracle, minQuantity)
}

// ActivateUnderlying is a paid mutator transaction binding the contract method 0xfe6bb409.
//
// Solidity: function activateUnderlying(uint8 underlying, address oracle, uint256 minQuantity) returns()
func (_Margin *MarginSession) ActivateUnderlying(underlying uint8, oracle common.Address, minQuantity *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.ActivateUnderlying(&_Margin.TransactOpts, underlying, oracle, minQuantity)
}

// ActivateUnderlying is a paid mutator transaction binding the contract method 0xfe6bb409.
//
// Solidity: function activateUnderlying(uint8 underlying, address oracle, uint256 minQuantity) returns()
func (_Margin *MarginTransactorSession) ActivateUnderlying(underlying uint8, oracle common.Address, minQuantity *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.ActivateUnderlying(&_Margin.TransactOpts, underlying, oracle, minQuantity)
}

// AddKeepers is a paid mutator transaction binding the contract method 0x314a83c9.
//
// Solidity: function addKeepers(address[] accounts) returns()
func (_Margin *MarginTransactor) AddKeepers(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "addKeepers", accounts)
}

// AddKeepers is a paid mutator transaction binding the contract method 0x314a83c9.
//
// Solidity: function addKeepers(address[] accounts) returns()
func (_Margin *MarginSession) AddKeepers(accounts []common.Address) (*types.Transaction, error) {
	return _Margin.Contract.AddKeepers(&_Margin.TransactOpts, accounts)
}

// AddKeepers is a paid mutator transaction binding the contract method 0x314a83c9.
//
// Solidity: function addKeepers(address[] accounts) returns()
func (_Margin *MarginTransactorSession) AddKeepers(accounts []common.Address) (*types.Transaction, error) {
	return _Margin.Contract.AddKeepers(&_Margin.TransactOpts, accounts)
}

// AddPosition is a paid mutator transaction binding the contract method 0x846c260f.
//
// Solidity: function addPosition((string,address,address,uint256,uint256,bool,uint8,uint8,bool,bool) params) returns()
func (_Margin *MarginTransactor) AddPosition(opts *bind.TransactOpts, params DerivativePositionParams) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "addPosition", params)
}

// AddPosition is a paid mutator transaction binding the contract method 0x846c260f.
//
// Solidity: function addPosition((string,address,address,uint256,uint256,bool,uint8,uint8,bool,bool) params) returns()
func (_Margin *MarginSession) AddPosition(params DerivativePositionParams) (*types.Transaction, error) {
	return _Margin.Contract.AddPosition(&_Margin.TransactOpts, params)
}

// AddPosition is a paid mutator transaction binding the contract method 0x846c260f.
//
// Solidity: function addPosition((string,address,address,uint256,uint256,bool,uint8,uint8,bool,bool) params) returns()
func (_Margin *MarginTransactorSession) AddPosition(params DerivativePositionParams) (*types.Transaction, error) {
	return _Margin.Contract.AddPosition(&_Margin.TransactOpts, params)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] accounts) returns()
func (_Margin *MarginTransactor) AddToWhitelist(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "addToWhitelist", accounts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] accounts) returns()
func (_Margin *MarginSession) AddToWhitelist(accounts []common.Address) (*types.Transaction, error) {
	return _Margin.Contract.AddToWhitelist(&_Margin.TransactOpts, accounts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] accounts) returns()
func (_Margin *MarginTransactorSession) AddToWhitelist(accounts []common.Address) (*types.Transaction, error) {
	return _Margin.Contract.AddToWhitelist(&_Margin.TransactOpts, accounts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Margin *MarginTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Margin *MarginSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.Deposit(&_Margin.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Margin *MarginTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.Deposit(&_Margin.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x01b4a043.
//
// Solidity: function initialize(address usdc_, address insurance_, address feeRecipient_, uint8 underlying_, address oracle_, uint256 minQuantity_) returns()
func (_Margin *MarginTransactor) Initialize(opts *bind.TransactOpts, usdc_ common.Address, insurance_ common.Address, feeRecipient_ common.Address, underlying_ uint8, oracle_ common.Address, minQuantity_ *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "initialize", usdc_, insurance_, feeRecipient_, underlying_, oracle_, minQuantity_)
}

// Initialize is a paid mutator transaction binding the contract method 0x01b4a043.
//
// Solidity: function initialize(address usdc_, address insurance_, address feeRecipient_, uint8 underlying_, address oracle_, uint256 minQuantity_) returns()
func (_Margin *MarginSession) Initialize(usdc_ common.Address, insurance_ common.Address, feeRecipient_ common.Address, underlying_ uint8, oracle_ common.Address, minQuantity_ *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.Initialize(&_Margin.TransactOpts, usdc_, insurance_, feeRecipient_, underlying_, oracle_, minQuantity_)
}

// Initialize is a paid mutator transaction binding the contract method 0x01b4a043.
//
// Solidity: function initialize(address usdc_, address insurance_, address feeRecipient_, uint8 underlying_, address oracle_, uint256 minQuantity_) returns()
func (_Margin *MarginTransactorSession) Initialize(usdc_ common.Address, insurance_ common.Address, feeRecipient_ common.Address, underlying_ uint8, oracle_ common.Address, minQuantity_ *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.Initialize(&_Margin.TransactOpts, usdc_, insurance_, feeRecipient_, underlying_, oracle_, minQuantity_)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address user) returns()
func (_Margin *MarginTransactor) Liquidate(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "liquidate", user)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address user) returns()
func (_Margin *MarginSession) Liquidate(user common.Address) (*types.Transaction, error) {
	return _Margin.Contract.Liquidate(&_Margin.TransactOpts, user)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address user) returns()
func (_Margin *MarginTransactorSession) Liquidate(user common.Address) (*types.Transaction, error) {
	return _Margin.Contract.Liquidate(&_Margin.TransactOpts, user)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] accounts) returns()
func (_Margin *MarginTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "removeFromWhitelist", accounts)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] accounts) returns()
func (_Margin *MarginSession) RemoveFromWhitelist(accounts []common.Address) (*types.Transaction, error) {
	return _Margin.Contract.RemoveFromWhitelist(&_Margin.TransactOpts, accounts)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] accounts) returns()
func (_Margin *MarginTransactorSession) RemoveFromWhitelist(accounts []common.Address) (*types.Transaction, error) {
	return _Margin.Contract.RemoveFromWhitelist(&_Margin.TransactOpts, accounts)
}

// RemoveKeepers is a paid mutator transaction binding the contract method 0xea7e9daa.
//
// Solidity: function removeKeepers(address[] accounts) returns()
func (_Margin *MarginTransactor) RemoveKeepers(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "removeKeepers", accounts)
}

// RemoveKeepers is a paid mutator transaction binding the contract method 0xea7e9daa.
//
// Solidity: function removeKeepers(address[] accounts) returns()
func (_Margin *MarginSession) RemoveKeepers(accounts []common.Address) (*types.Transaction, error) {
	return _Margin.Contract.RemoveKeepers(&_Margin.TransactOpts, accounts)
}

// RemoveKeepers is a paid mutator transaction binding the contract method 0xea7e9daa.
//
// Solidity: function removeKeepers(address[] accounts) returns()
func (_Margin *MarginTransactorSession) RemoveKeepers(accounts []common.Address) (*types.Transaction, error) {
	return _Margin.Contract.RemoveKeepers(&_Margin.TransactOpts, accounts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Margin *MarginTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Margin *MarginSession) RenounceOwnership() (*types.Transaction, error) {
	return _Margin.Contract.RenounceOwnership(&_Margin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Margin *MarginTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Margin.Contract.RenounceOwnership(&_Margin.TransactOpts)
}

// Rollover is a paid mutator transaction binding the contract method 0x05c4d1cb.
//
// Solidity: function rollover(address[] roundUsers) returns()
func (_Margin *MarginTransactor) Rollover(opts *bind.TransactOpts, roundUsers []common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "rollover", roundUsers)
}

// Rollover is a paid mutator transaction binding the contract method 0x05c4d1cb.
//
// Solidity: function rollover(address[] roundUsers) returns()
func (_Margin *MarginSession) Rollover(roundUsers []common.Address) (*types.Transaction, error) {
	return _Margin.Contract.Rollover(&_Margin.TransactOpts, roundUsers)
}

// Rollover is a paid mutator transaction binding the contract method 0x05c4d1cb.
//
// Solidity: function rollover(address[] roundUsers) returns()
func (_Margin *MarginTransactorSession) Rollover(roundUsers []common.Address) (*types.Transaction, error) {
	return _Margin.Contract.Rollover(&_Margin.TransactOpts, roundUsers)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Margin *MarginTransactor) SetFeeRecipient(opts *bind.TransactOpts, newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "setFeeRecipient", newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Margin *MarginSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Margin.Contract.SetFeeRecipient(&_Margin.TransactOpts, newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Margin *MarginTransactorSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Margin.Contract.SetFeeRecipient(&_Margin.TransactOpts, newFeeRecipient)
}

// SetInsurance is a paid mutator transaction binding the contract method 0x6cd7751f.
//
// Solidity: function setInsurance(address newInsurance) returns()
func (_Margin *MarginTransactor) SetInsurance(opts *bind.TransactOpts, newInsurance common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "setInsurance", newInsurance)
}

// SetInsurance is a paid mutator transaction binding the contract method 0x6cd7751f.
//
// Solidity: function setInsurance(address newInsurance) returns()
func (_Margin *MarginSession) SetInsurance(newInsurance common.Address) (*types.Transaction, error) {
	return _Margin.Contract.SetInsurance(&_Margin.TransactOpts, newInsurance)
}

// SetInsurance is a paid mutator transaction binding the contract method 0x6cd7751f.
//
// Solidity: function setInsurance(address newInsurance) returns()
func (_Margin *MarginTransactorSession) SetInsurance(newInsurance common.Address) (*types.Transaction, error) {
	return _Margin.Contract.SetInsurance(&_Margin.TransactOpts, newInsurance)
}

// SetMaxBalanceCap is a paid mutator transaction binding the contract method 0xf555a419.
//
// Solidity: function setMaxBalanceCap(uint256 maxBalance) returns()
func (_Margin *MarginTransactor) SetMaxBalanceCap(opts *bind.TransactOpts, maxBalance *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "setMaxBalanceCap", maxBalance)
}

// SetMaxBalanceCap is a paid mutator transaction binding the contract method 0xf555a419.
//
// Solidity: function setMaxBalanceCap(uint256 maxBalance) returns()
func (_Margin *MarginSession) SetMaxBalanceCap(maxBalance *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.SetMaxBalanceCap(&_Margin.TransactOpts, maxBalance)
}

// SetMaxBalanceCap is a paid mutator transaction binding the contract method 0xf555a419.
//
// Solidity: function setMaxBalanceCap(uint256 maxBalance) returns()
func (_Margin *MarginTransactorSession) SetMaxBalanceCap(maxBalance *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.SetMaxBalanceCap(&_Margin.TransactOpts, maxBalance)
}

// SetMinMarginPerc is a paid mutator transaction binding the contract method 0xa1dc656b.
//
// Solidity: function setMinMarginPerc(uint256 perc) returns()
func (_Margin *MarginTransactor) SetMinMarginPerc(opts *bind.TransactOpts, perc *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "setMinMarginPerc", perc)
}

// SetMinMarginPerc is a paid mutator transaction binding the contract method 0xa1dc656b.
//
// Solidity: function setMinMarginPerc(uint256 perc) returns()
func (_Margin *MarginSession) SetMinMarginPerc(perc *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.SetMinMarginPerc(&_Margin.TransactOpts, perc)
}

// SetMinMarginPerc is a paid mutator transaction binding the contract method 0xa1dc656b.
//
// Solidity: function setMinMarginPerc(uint256 perc) returns()
func (_Margin *MarginTransactorSession) SetMinMarginPerc(perc *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.SetMinMarginPerc(&_Margin.TransactOpts, perc)
}

// SetMinQuantity is a paid mutator transaction binding the contract method 0x5c563dbe.
//
// Solidity: function setMinQuantity(uint8 underlying, uint256 minQuantity) returns()
func (_Margin *MarginTransactor) SetMinQuantity(opts *bind.TransactOpts, underlying uint8, minQuantity *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "setMinQuantity", underlying, minQuantity)
}

// SetMinQuantity is a paid mutator transaction binding the contract method 0x5c563dbe.
//
// Solidity: function setMinQuantity(uint8 underlying, uint256 minQuantity) returns()
func (_Margin *MarginSession) SetMinQuantity(underlying uint8, minQuantity *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.SetMinQuantity(&_Margin.TransactOpts, underlying, minQuantity)
}

// SetMinQuantity is a paid mutator transaction binding the contract method 0x5c563dbe.
//
// Solidity: function setMinQuantity(uint8 underlying, uint256 minQuantity) returns()
func (_Margin *MarginTransactorSession) SetMinQuantity(underlying uint8, minQuantity *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.SetMinQuantity(&_Margin.TransactOpts, underlying, minQuantity)
}

// SetOracle is a paid mutator transaction binding the contract method 0x1b784a47.
//
// Solidity: function setOracle(uint8 underlying, address oracle) returns()
func (_Margin *MarginTransactor) SetOracle(opts *bind.TransactOpts, underlying uint8, oracle common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "setOracle", underlying, oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x1b784a47.
//
// Solidity: function setOracle(uint8 underlying, address oracle) returns()
func (_Margin *MarginSession) SetOracle(underlying uint8, oracle common.Address) (*types.Transaction, error) {
	return _Margin.Contract.SetOracle(&_Margin.TransactOpts, underlying, oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x1b784a47.
//
// Solidity: function setOracle(uint8 underlying, address oracle) returns()
func (_Margin *MarginTransactorSession) SetOracle(underlying uint8, oracle common.Address) (*types.Transaction, error) {
	return _Margin.Contract.SetOracle(&_Margin.TransactOpts, underlying, oracle)
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() returns()
func (_Margin *MarginTransactor) Settle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "settle")
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() returns()
func (_Margin *MarginSession) Settle() (*types.Transaction, error) {
	return _Margin.Contract.Settle(&_Margin.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() returns()
func (_Margin *MarginTransactorSession) Settle() (*types.Transaction, error) {
	return _Margin.Contract.Settle(&_Margin.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_Margin *MarginTransactor) TogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "togglePause")
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_Margin *MarginSession) TogglePause() (*types.Transaction, error) {
	return _Margin.Contract.TogglePause(&_Margin.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_Margin *MarginTransactorSession) TogglePause() (*types.Transaction, error) {
	return _Margin.Contract.TogglePause(&_Margin.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Margin *MarginTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Margin *MarginSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Margin.Contract.TransferOwnership(&_Margin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Margin *MarginTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Margin.Contract.TransferOwnership(&_Margin.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Margin *MarginTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Margin *MarginSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Margin.Contract.UpgradeTo(&_Margin.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Margin *MarginTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Margin.Contract.UpgradeTo(&_Margin.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Margin *MarginTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Margin *MarginSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Margin.Contract.UpgradeToAndCall(&_Margin.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Margin *MarginTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Margin.Contract.UpgradeToAndCall(&_Margin.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Margin *MarginTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Margin *MarginSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.Withdraw(&_Margin.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Margin *MarginTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Margin.Contract.Withdraw(&_Margin.TransactOpts, amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Margin *MarginTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Margin.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Margin *MarginSession) WithdrawAll() (*types.Transaction, error) {
	return _Margin.Contract.WithdrawAll(&_Margin.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Margin *MarginTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Margin.Contract.WithdrawAll(&_Margin.TransactOpts)
}

// MarginAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Margin contract.
type MarginAdminChangedIterator struct {
	Event *MarginAdminChanged // Event containing the contract specifics and raw log

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
func (it *MarginAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginAdminChanged)
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
		it.Event = new(MarginAdminChanged)
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
func (it *MarginAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginAdminChanged represents a AdminChanged event raised by the Margin contract.
type MarginAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Margin *MarginFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*MarginAdminChangedIterator, error) {

	logs, sub, err := _Margin.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &MarginAdminChangedIterator{contract: _Margin.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Margin *MarginFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *MarginAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Margin.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginAdminChanged)
				if err := _Margin.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Margin *MarginFilterer) ParseAdminChanged(log types.Log) (*MarginAdminChanged, error) {
	event := new(MarginAdminChanged)
	if err := _Margin.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginBankruptcyEventIterator is returned from FilterBankruptcyEvent and is used to iterate over the raw logs and unpacked data for BankruptcyEvent events raised by the Margin contract.
type MarginBankruptcyEventIterator struct {
	Event *MarginBankruptcyEvent // Event containing the contract specifics and raw log

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
func (it *MarginBankruptcyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginBankruptcyEvent)
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
		it.Event = new(MarginBankruptcyEvent)
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
func (it *MarginBankruptcyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginBankruptcyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginBankruptcyEvent represents a BankruptcyEvent event raised by the Margin contract.
type MarginBankruptcyEvent struct {
	BankruptUser     common.Address
	Counterparty     common.Address
	Amount           *big.Int
	BankruptcyAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBankruptcyEvent is a free log retrieval operation binding the contract event 0x4d546f30d16f3c74d3123fc4c51ff5535dc33bb0819abeb9d8e9e097a72fefbb.
//
// Solidity: event BankruptcyEvent(address indexed bankruptUser, address indexed counterparty, uint256 amount, uint256 bankruptcyAmount)
func (_Margin *MarginFilterer) FilterBankruptcyEvent(opts *bind.FilterOpts, bankruptUser []common.Address, counterparty []common.Address) (*MarginBankruptcyEventIterator, error) {

	var bankruptUserRule []interface{}
	for _, bankruptUserItem := range bankruptUser {
		bankruptUserRule = append(bankruptUserRule, bankruptUserItem)
	}
	var counterpartyRule []interface{}
	for _, counterpartyItem := range counterparty {
		counterpartyRule = append(counterpartyRule, counterpartyItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "BankruptcyEvent", bankruptUserRule, counterpartyRule)
	if err != nil {
		return nil, err
	}
	return &MarginBankruptcyEventIterator{contract: _Margin.contract, event: "BankruptcyEvent", logs: logs, sub: sub}, nil
}

// WatchBankruptcyEvent is a free log subscription operation binding the contract event 0x4d546f30d16f3c74d3123fc4c51ff5535dc33bb0819abeb9d8e9e097a72fefbb.
//
// Solidity: event BankruptcyEvent(address indexed bankruptUser, address indexed counterparty, uint256 amount, uint256 bankruptcyAmount)
func (_Margin *MarginFilterer) WatchBankruptcyEvent(opts *bind.WatchOpts, sink chan<- *MarginBankruptcyEvent, bankruptUser []common.Address, counterparty []common.Address) (event.Subscription, error) {

	var bankruptUserRule []interface{}
	for _, bankruptUserItem := range bankruptUser {
		bankruptUserRule = append(bankruptUserRule, bankruptUserItem)
	}
	var counterpartyRule []interface{}
	for _, counterpartyItem := range counterparty {
		counterpartyRule = append(counterpartyRule, counterpartyItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "BankruptcyEvent", bankruptUserRule, counterpartyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginBankruptcyEvent)
				if err := _Margin.contract.UnpackLog(event, "BankruptcyEvent", log); err != nil {
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

// ParseBankruptcyEvent is a log parse operation binding the contract event 0x4d546f30d16f3c74d3123fc4c51ff5535dc33bb0819abeb9d8e9e097a72fefbb.
//
// Solidity: event BankruptcyEvent(address indexed bankruptUser, address indexed counterparty, uint256 amount, uint256 bankruptcyAmount)
func (_Margin *MarginFilterer) ParseBankruptcyEvent(log types.Log) (*MarginBankruptcyEvent, error) {
	event := new(MarginBankruptcyEvent)
	if err := _Margin.contract.UnpackLog(event, "BankruptcyEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Margin contract.
type MarginBeaconUpgradedIterator struct {
	Event *MarginBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *MarginBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginBeaconUpgraded)
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
		it.Event = new(MarginBeaconUpgraded)
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
func (it *MarginBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginBeaconUpgraded represents a BeaconUpgraded event raised by the Margin contract.
type MarginBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Margin *MarginFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*MarginBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &MarginBeaconUpgradedIterator{contract: _Margin.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Margin *MarginFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *MarginBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginBeaconUpgraded)
				if err := _Margin.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Margin *MarginFilterer) ParseBeaconUpgraded(log types.Log) (*MarginBeaconUpgraded, error) {
	event := new(MarginBeaconUpgraded)
	if err := _Margin.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the Margin contract.
type MarginDepositEventIterator struct {
	Event *MarginDepositEvent // Event containing the contract specifics and raw log

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
func (it *MarginDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginDepositEvent)
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
		it.Event = new(MarginDepositEvent)
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
func (it *MarginDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginDepositEvent represents a DepositEvent event raised by the Margin contract.
type MarginDepositEvent struct {
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0x2d8a08b6430a894aea608bcaa6013d5d3e263bc49110605e4d4ba76930ae5c29.
//
// Solidity: event DepositEvent(address indexed depositor, uint256 amount)
func (_Margin *MarginFilterer) FilterDepositEvent(opts *bind.FilterOpts, depositor []common.Address) (*MarginDepositEventIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "DepositEvent", depositorRule)
	if err != nil {
		return nil, err
	}
	return &MarginDepositEventIterator{contract: _Margin.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0x2d8a08b6430a894aea608bcaa6013d5d3e263bc49110605e4d4ba76930ae5c29.
//
// Solidity: event DepositEvent(address indexed depositor, uint256 amount)
func (_Margin *MarginFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *MarginDepositEvent, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "DepositEvent", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginDepositEvent)
				if err := _Margin.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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

// ParseDepositEvent is a log parse operation binding the contract event 0x2d8a08b6430a894aea608bcaa6013d5d3e263bc49110605e4d4ba76930ae5c29.
//
// Solidity: event DepositEvent(address indexed depositor, uint256 amount)
func (_Margin *MarginFilterer) ParseDepositEvent(log types.Log) (*MarginDepositEvent, error) {
	event := new(MarginDepositEvent)
	if err := _Margin.contract.UnpackLog(event, "DepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Margin contract.
type MarginInitializedIterator struct {
	Event *MarginInitialized // Event containing the contract specifics and raw log

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
func (it *MarginInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginInitialized)
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
		it.Event = new(MarginInitialized)
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
func (it *MarginInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginInitialized represents a Initialized event raised by the Margin contract.
type MarginInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Margin *MarginFilterer) FilterInitialized(opts *bind.FilterOpts) (*MarginInitializedIterator, error) {

	logs, sub, err := _Margin.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MarginInitializedIterator{contract: _Margin.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Margin *MarginFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MarginInitialized) (event.Subscription, error) {

	logs, sub, err := _Margin.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginInitialized)
				if err := _Margin.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Margin *MarginFilterer) ParseInitialized(log types.Log) (*MarginInitialized, error) {
	event := new(MarginInitialized)
	if err := _Margin.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginMaxBalanceCapEventIterator is returned from FilterMaxBalanceCapEvent and is used to iterate over the raw logs and unpacked data for MaxBalanceCapEvent events raised by the Margin contract.
type MarginMaxBalanceCapEventIterator struct {
	Event *MarginMaxBalanceCapEvent // Event containing the contract specifics and raw log

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
func (it *MarginMaxBalanceCapEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginMaxBalanceCapEvent)
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
		it.Event = new(MarginMaxBalanceCapEvent)
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
func (it *MarginMaxBalanceCapEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginMaxBalanceCapEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginMaxBalanceCapEvent represents a MaxBalanceCapEvent event raised by the Margin contract.
type MarginMaxBalanceCapEvent struct {
	Owner      common.Address
	MaxBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMaxBalanceCapEvent is a free log retrieval operation binding the contract event 0x43bf10be81662116415358cada6e4928c87d70e8c74de1d25886dd63940c215d.
//
// Solidity: event MaxBalanceCapEvent(address indexed owner, uint256 maxBalance)
func (_Margin *MarginFilterer) FilterMaxBalanceCapEvent(opts *bind.FilterOpts, owner []common.Address) (*MarginMaxBalanceCapEventIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "MaxBalanceCapEvent", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MarginMaxBalanceCapEventIterator{contract: _Margin.contract, event: "MaxBalanceCapEvent", logs: logs, sub: sub}, nil
}

// WatchMaxBalanceCapEvent is a free log subscription operation binding the contract event 0x43bf10be81662116415358cada6e4928c87d70e8c74de1d25886dd63940c215d.
//
// Solidity: event MaxBalanceCapEvent(address indexed owner, uint256 maxBalance)
func (_Margin *MarginFilterer) WatchMaxBalanceCapEvent(opts *bind.WatchOpts, sink chan<- *MarginMaxBalanceCapEvent, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "MaxBalanceCapEvent", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginMaxBalanceCapEvent)
				if err := _Margin.contract.UnpackLog(event, "MaxBalanceCapEvent", log); err != nil {
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

// ParseMaxBalanceCapEvent is a log parse operation binding the contract event 0x43bf10be81662116415358cada6e4928c87d70e8c74de1d25886dd63940c215d.
//
// Solidity: event MaxBalanceCapEvent(address indexed owner, uint256 maxBalance)
func (_Margin *MarginFilterer) ParseMaxBalanceCapEvent(log types.Log) (*MarginMaxBalanceCapEvent, error) {
	event := new(MarginMaxBalanceCapEvent)
	if err := _Margin.contract.UnpackLog(event, "MaxBalanceCapEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginMinMarginPercEventIterator is returned from FilterMinMarginPercEvent and is used to iterate over the raw logs and unpacked data for MinMarginPercEvent events raised by the Margin contract.
type MarginMinMarginPercEventIterator struct {
	Event *MarginMinMarginPercEvent // Event containing the contract specifics and raw log

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
func (it *MarginMinMarginPercEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginMinMarginPercEvent)
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
		it.Event = new(MarginMinMarginPercEvent)
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
func (it *MarginMinMarginPercEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginMinMarginPercEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginMinMarginPercEvent represents a MinMarginPercEvent event raised by the Margin contract.
type MarginMinMarginPercEvent struct {
	Owner common.Address
	Perc  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinMarginPercEvent is a free log retrieval operation binding the contract event 0xc4d2cbc4040c3ef45136b4540fbaaf123b01b278d3323ff57d3a3758c8a312ce.
//
// Solidity: event MinMarginPercEvent(address indexed owner, uint256 perc)
func (_Margin *MarginFilterer) FilterMinMarginPercEvent(opts *bind.FilterOpts, owner []common.Address) (*MarginMinMarginPercEventIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "MinMarginPercEvent", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MarginMinMarginPercEventIterator{contract: _Margin.contract, event: "MinMarginPercEvent", logs: logs, sub: sub}, nil
}

// WatchMinMarginPercEvent is a free log subscription operation binding the contract event 0xc4d2cbc4040c3ef45136b4540fbaaf123b01b278d3323ff57d3a3758c8a312ce.
//
// Solidity: event MinMarginPercEvent(address indexed owner, uint256 perc)
func (_Margin *MarginFilterer) WatchMinMarginPercEvent(opts *bind.WatchOpts, sink chan<- *MarginMinMarginPercEvent, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "MinMarginPercEvent", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginMinMarginPercEvent)
				if err := _Margin.contract.UnpackLog(event, "MinMarginPercEvent", log); err != nil {
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

// ParseMinMarginPercEvent is a log parse operation binding the contract event 0xc4d2cbc4040c3ef45136b4540fbaaf123b01b278d3323ff57d3a3758c8a312ce.
//
// Solidity: event MinMarginPercEvent(address indexed owner, uint256 perc)
func (_Margin *MarginFilterer) ParseMinMarginPercEvent(log types.Log) (*MarginMinMarginPercEvent, error) {
	event := new(MarginMinMarginPercEvent)
	if err := _Margin.contract.UnpackLog(event, "MinMarginPercEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginNewUnderlyingEventIterator is returned from FilterNewUnderlyingEvent and is used to iterate over the raw logs and unpacked data for NewUnderlyingEvent events raised by the Margin contract.
type MarginNewUnderlyingEventIterator struct {
	Event *MarginNewUnderlyingEvent // Event containing the contract specifics and raw log

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
func (it *MarginNewUnderlyingEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginNewUnderlyingEvent)
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
		it.Event = new(MarginNewUnderlyingEvent)
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
func (it *MarginNewUnderlyingEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginNewUnderlyingEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginNewUnderlyingEvent represents a NewUnderlyingEvent event raised by the Margin contract.
type MarginNewUnderlyingEvent struct {
	Owner       common.Address
	Underlying  uint8
	Oracle      common.Address
	MinQuantity *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewUnderlyingEvent is a free log retrieval operation binding the contract event 0xb074d74b8b193c1b3395bc3dbb1b97e42880495c0f2708754d80d19921d14f95.
//
// Solidity: event NewUnderlyingEvent(address indexed owner, uint8 underlying, address oracle, uint256 minQuantity)
func (_Margin *MarginFilterer) FilterNewUnderlyingEvent(opts *bind.FilterOpts, owner []common.Address) (*MarginNewUnderlyingEventIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "NewUnderlyingEvent", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MarginNewUnderlyingEventIterator{contract: _Margin.contract, event: "NewUnderlyingEvent", logs: logs, sub: sub}, nil
}

// WatchNewUnderlyingEvent is a free log subscription operation binding the contract event 0xb074d74b8b193c1b3395bc3dbb1b97e42880495c0f2708754d80d19921d14f95.
//
// Solidity: event NewUnderlyingEvent(address indexed owner, uint8 underlying, address oracle, uint256 minQuantity)
func (_Margin *MarginFilterer) WatchNewUnderlyingEvent(opts *bind.WatchOpts, sink chan<- *MarginNewUnderlyingEvent, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "NewUnderlyingEvent", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginNewUnderlyingEvent)
				if err := _Margin.contract.UnpackLog(event, "NewUnderlyingEvent", log); err != nil {
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

// ParseNewUnderlyingEvent is a log parse operation binding the contract event 0xb074d74b8b193c1b3395bc3dbb1b97e42880495c0f2708754d80d19921d14f95.
//
// Solidity: event NewUnderlyingEvent(address indexed owner, uint8 underlying, address oracle, uint256 minQuantity)
func (_Margin *MarginFilterer) ParseNewUnderlyingEvent(log types.Log) (*MarginNewUnderlyingEvent, error) {
	event := new(MarginNewUnderlyingEvent)
	if err := _Margin.contract.UnpackLog(event, "NewUnderlyingEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Margin contract.
type MarginOwnershipTransferredIterator struct {
	Event *MarginOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarginOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginOwnershipTransferred)
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
		it.Event = new(MarginOwnershipTransferred)
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
func (it *MarginOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginOwnershipTransferred represents a OwnershipTransferred event raised by the Margin contract.
type MarginOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Margin *MarginFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarginOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarginOwnershipTransferredIterator{contract: _Margin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Margin *MarginFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarginOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginOwnershipTransferred)
				if err := _Margin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Margin *MarginFilterer) ParseOwnershipTransferred(log types.Log) (*MarginOwnershipTransferred, error) {
	event := new(MarginOwnershipTransferred)
	if err := _Margin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginRecordPositionEventIterator is returned from FilterRecordPositionEvent and is used to iterate over the raw logs and unpacked data for RecordPositionEvent events raised by the Margin contract.
type MarginRecordPositionEventIterator struct {
	Event *MarginRecordPositionEvent // Event containing the contract specifics and raw log

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
func (it *MarginRecordPositionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginRecordPositionEvent)
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
		it.Event = new(MarginRecordPositionEvent)
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
func (it *MarginRecordPositionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginRecordPositionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginRecordPositionEvent represents a RecordPositionEvent event raised by the Margin contract.
type MarginRecordPositionEvent struct {
	Id          string
	TradePrice  *big.Int
	Quantity    *big.Int
	IsCall      bool
	Underlying  uint8
	StrikeLevel uint8
	Expiry      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRecordPositionEvent is a free log retrieval operation binding the contract event 0x2bea814922af00322b50dbedc22798d1e1c00ddd12bbc8fb60e2ca3d6fe7660e.
//
// Solidity: event RecordPositionEvent(string id, uint256 tradePrice, uint256 quantity, bool isCall, uint8 underlying, uint8 strikeLevel, uint256 expiry)
func (_Margin *MarginFilterer) FilterRecordPositionEvent(opts *bind.FilterOpts) (*MarginRecordPositionEventIterator, error) {

	logs, sub, err := _Margin.contract.FilterLogs(opts, "RecordPositionEvent")
	if err != nil {
		return nil, err
	}
	return &MarginRecordPositionEventIterator{contract: _Margin.contract, event: "RecordPositionEvent", logs: logs, sub: sub}, nil
}

// WatchRecordPositionEvent is a free log subscription operation binding the contract event 0x2bea814922af00322b50dbedc22798d1e1c00ddd12bbc8fb60e2ca3d6fe7660e.
//
// Solidity: event RecordPositionEvent(string id, uint256 tradePrice, uint256 quantity, bool isCall, uint8 underlying, uint8 strikeLevel, uint256 expiry)
func (_Margin *MarginFilterer) WatchRecordPositionEvent(opts *bind.WatchOpts, sink chan<- *MarginRecordPositionEvent) (event.Subscription, error) {

	logs, sub, err := _Margin.contract.WatchLogs(opts, "RecordPositionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginRecordPositionEvent)
				if err := _Margin.contract.UnpackLog(event, "RecordPositionEvent", log); err != nil {
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

// ParseRecordPositionEvent is a log parse operation binding the contract event 0x2bea814922af00322b50dbedc22798d1e1c00ddd12bbc8fb60e2ca3d6fe7660e.
//
// Solidity: event RecordPositionEvent(string id, uint256 tradePrice, uint256 quantity, bool isCall, uint8 underlying, uint8 strikeLevel, uint256 expiry)
func (_Margin *MarginFilterer) ParseRecordPositionEvent(log types.Log) (*MarginRecordPositionEvent, error) {
	event := new(MarginRecordPositionEvent)
	if err := _Margin.contract.UnpackLog(event, "RecordPositionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginSettlementEventIterator is returned from FilterSettlementEvent and is used to iterate over the raw logs and unpacked data for SettlementEvent events raised by the Margin contract.
type MarginSettlementEventIterator struct {
	Event *MarginSettlementEvent // Event containing the contract specifics and raw log

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
func (it *MarginSettlementEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginSettlementEvent)
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
		it.Event = new(MarginSettlementEvent)
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
func (it *MarginSettlementEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginSettlementEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginSettlementEvent represents a SettlementEvent event raised by the Margin contract.
type MarginSettlementEvent struct {
	Caller       common.Address
	Round        uint8
	NumPositions *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSettlementEvent is a free log retrieval operation binding the contract event 0x348b1bdb7ff7d772beba31e8d7cb57d1c2b108eaa1e3ac5ebce592c036e328d0.
//
// Solidity: event SettlementEvent(address indexed caller, uint8 round, uint256 numPositions)
func (_Margin *MarginFilterer) FilterSettlementEvent(opts *bind.FilterOpts, caller []common.Address) (*MarginSettlementEventIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "SettlementEvent", callerRule)
	if err != nil {
		return nil, err
	}
	return &MarginSettlementEventIterator{contract: _Margin.contract, event: "SettlementEvent", logs: logs, sub: sub}, nil
}

// WatchSettlementEvent is a free log subscription operation binding the contract event 0x348b1bdb7ff7d772beba31e8d7cb57d1c2b108eaa1e3ac5ebce592c036e328d0.
//
// Solidity: event SettlementEvent(address indexed caller, uint8 round, uint256 numPositions)
func (_Margin *MarginFilterer) WatchSettlementEvent(opts *bind.WatchOpts, sink chan<- *MarginSettlementEvent, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "SettlementEvent", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginSettlementEvent)
				if err := _Margin.contract.UnpackLog(event, "SettlementEvent", log); err != nil {
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

// ParseSettlementEvent is a log parse operation binding the contract event 0x348b1bdb7ff7d772beba31e8d7cb57d1c2b108eaa1e3ac5ebce592c036e328d0.
//
// Solidity: event SettlementEvent(address indexed caller, uint8 round, uint256 numPositions)
func (_Margin *MarginFilterer) ParseSettlementEvent(log types.Log) (*MarginSettlementEvent, error) {
	event := new(MarginSettlementEvent)
	if err := _Margin.contract.UnpackLog(event, "SettlementEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginTogglePauseEventIterator is returned from FilterTogglePauseEvent and is used to iterate over the raw logs and unpacked data for TogglePauseEvent events raised by the Margin contract.
type MarginTogglePauseEventIterator struct {
	Event *MarginTogglePauseEvent // Event containing the contract specifics and raw log

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
func (it *MarginTogglePauseEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginTogglePauseEvent)
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
		it.Event = new(MarginTogglePauseEvent)
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
func (it *MarginTogglePauseEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginTogglePauseEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginTogglePauseEvent represents a TogglePauseEvent event raised by the Margin contract.
type MarginTogglePauseEvent struct {
	Owner  common.Address
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTogglePauseEvent is a free log retrieval operation binding the contract event 0x66f8a7ac2d1f7da12f55df6017dd3e4785a2c733138e53bc548a977cbee10998.
//
// Solidity: event TogglePauseEvent(address indexed owner, bool paused)
func (_Margin *MarginFilterer) FilterTogglePauseEvent(opts *bind.FilterOpts, owner []common.Address) (*MarginTogglePauseEventIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "TogglePauseEvent", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MarginTogglePauseEventIterator{contract: _Margin.contract, event: "TogglePauseEvent", logs: logs, sub: sub}, nil
}

// WatchTogglePauseEvent is a free log subscription operation binding the contract event 0x66f8a7ac2d1f7da12f55df6017dd3e4785a2c733138e53bc548a977cbee10998.
//
// Solidity: event TogglePauseEvent(address indexed owner, bool paused)
func (_Margin *MarginFilterer) WatchTogglePauseEvent(opts *bind.WatchOpts, sink chan<- *MarginTogglePauseEvent, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "TogglePauseEvent", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginTogglePauseEvent)
				if err := _Margin.contract.UnpackLog(event, "TogglePauseEvent", log); err != nil {
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

// ParseTogglePauseEvent is a log parse operation binding the contract event 0x66f8a7ac2d1f7da12f55df6017dd3e4785a2c733138e53bc548a977cbee10998.
//
// Solidity: event TogglePauseEvent(address indexed owner, bool paused)
func (_Margin *MarginFilterer) ParseTogglePauseEvent(log types.Log) (*MarginTogglePauseEvent, error) {
	event := new(MarginTogglePauseEvent)
	if err := _Margin.contract.UnpackLog(event, "TogglePauseEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Margin contract.
type MarginUpgradedIterator struct {
	Event *MarginUpgraded // Event containing the contract specifics and raw log

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
func (it *MarginUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginUpgraded)
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
		it.Event = new(MarginUpgraded)
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
func (it *MarginUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginUpgraded represents a Upgraded event raised by the Margin contract.
type MarginUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Margin *MarginFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MarginUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MarginUpgradedIterator{contract: _Margin.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Margin *MarginFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MarginUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginUpgraded)
				if err := _Margin.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Margin *MarginFilterer) ParseUpgraded(log types.Log) (*MarginUpgraded, error) {
	event := new(MarginUpgraded)
	if err := _Margin.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarginWithdrawEventIterator is returned from FilterWithdrawEvent and is used to iterate over the raw logs and unpacked data for WithdrawEvent events raised by the Margin contract.
type MarginWithdrawEventIterator struct {
	Event *MarginWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *MarginWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarginWithdrawEvent)
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
		it.Event = new(MarginWithdrawEvent)
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
func (it *MarginWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarginWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarginWithdrawEvent represents a WithdrawEvent event raised by the Margin contract.
type MarginWithdrawEvent struct {
	User       common.Address
	Amount     *big.Int
	Discounted *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEvent is a free log retrieval operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: event WithdrawEvent(address indexed user, uint256 amount, uint256 discounted)
func (_Margin *MarginFilterer) FilterWithdrawEvent(opts *bind.FilterOpts, user []common.Address) (*MarginWithdrawEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Margin.contract.FilterLogs(opts, "WithdrawEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &MarginWithdrawEventIterator{contract: _Margin.contract, event: "WithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawEvent is a free log subscription operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: event WithdrawEvent(address indexed user, uint256 amount, uint256 discounted)
func (_Margin *MarginFilterer) WatchWithdrawEvent(opts *bind.WatchOpts, sink chan<- *MarginWithdrawEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Margin.contract.WatchLogs(opts, "WithdrawEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarginWithdrawEvent)
				if err := _Margin.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
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

// ParseWithdrawEvent is a log parse operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: event WithdrawEvent(address indexed user, uint256 amount, uint256 discounted)
func (_Margin *MarginFilterer) ParseWithdrawEvent(log types.Log) (*MarginWithdrawEvent, error) {
	event := new(MarginWithdrawEvent)
	if err := _Margin.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
