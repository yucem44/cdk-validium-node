// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cdkvalidium

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

// PolygonDataComitteeValidiumBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonDataComitteeValidiumBatchData struct {
	TransactionsHash   [32]byte
	GlobalExitRoot     [32]byte
	Timestamp          uint64
	MinForcedTimestamp uint64
}

// PolygonRollupBaseBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonRollupBaseBatchData struct {
	Transactions       []byte
	GlobalExitRoot     [32]byte
	Timestamp          uint64
	MinForcedTimestamp uint64
}

// PolygonRollupBaseForcedBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonRollupBaseForcedBatchData struct {
	Transactions       []byte
	GlobalExitRoot     [32]byte
	MinForcedTimestamp uint64
}

// CdkvalidiumMetaData contains all meta data concerning the Cdkvalidium contract.
var CdkvalidiumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractICDKDataCommittee\",\"name\":\"_dataCommittee\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceWithDataAvailabilityNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ActivateForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SwitchSequenceWithDataAvailability\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataCommittee\",\"outputs\":[{\"internalType\":\"contractICDKDataCommittee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isForcedBatchAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSequenceWithDataAvailabilityAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonRollupBase.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonDataComittee.ValidiumBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataAvailabilityMessage\",\"type\":\"bytes\"}],\"name\":\"sequenceBatchesDataCommittee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonRollupBase.ForcedBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchSequenceWithDataAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61012060405234801562000011575f80fd5b50604051620052b8380380620052b8833981016040819052620000349162000076565b6001600160a01b0394851660a05292841660805290831660c052821660e0521661010052620000f2565b6001600160a01b038116811462000073575f80fd5b50565b5f805f805f60a086880312156200008b575f80fd5b855162000098816200005e565b6020870151909550620000ab816200005e565b6040870151909450620000be816200005e565b6060870151909350620000d1816200005e565b6080870151909250620000e4816200005e565b809150509295509295909350565b60805160a05160c05160e051610100516150bd620001fb5f395f81816105dc015261227e01525f818161053401528181610a2101528181610b8d01528181610f2e015281816114d60152818161233b0152818161240401528181612426015281816125b501528181612ac201528181612b8f01528181612c77015281816135e5015281816136660152818161368801526137aa01525f81816106d3015281816111820152818161189b015281816119a3015281816124fc01526136f101525f81816107b20152818161132f01528181611ff401528181612dbf015261330c01525f81816107f7015281816108c00152818161238e015281816124d20152612d9401526150bd5ff3fe608060405234801561000f575f80fd5b5060043610610303575f3560e01c80636e05d2cd1161019d578063c754c7ed116100e8578063d8d1091b11610093578063eaeb077b1161006e578063eaeb077b14610839578063f35dda471461084c578063f851a44014610854575f80fd5b8063d8d1091b146107df578063e46761c4146107f2578063e7a7ed0214610819575f80fd5b8063cfa8ed47116100c3578063cfa8ed471461078d578063d02103ca146107ad578063d7bc90ff146107d4575f80fd5b8063c754c7ed14610742578063c7fffd4b14610772578063c89e42df1461077a575f80fd5b8063a3c573eb11610148578063ada8f91911610123578063ada8f91914610710578063b0afe15414610723578063b81fcafd1461072f575f80fd5b8063a3c573eb146106ce578063a652f26c146106f5578063aa3587d314610708575f80fd5b80637a5460c5116101785780637a5460c51461066f5780638c3d7301146106ab5780639e001877146106b3575f80fd5b80636e05d2cd146106405780636ff512cc14610649578063712570221461065c575f80fd5b806340b5de6c1161025d578063542028d5116102085780635ec91958116101e35780635ec9195814610611578063676870d2146106195780636b8616ce14610621575f80fd5b8063542028d5146105cf57806356e29435146105d75780635e9145c9146105fe575f80fd5b80634c21fef3116102385780634c21fef3146105565780634e4877061461058057806352bdeb6d14610593575f80fd5b806340b5de6c146104af578063456052671461050757806349b7b8021461052f575f80fd5b806319d8ac61116102bd57806332c2d1531161029857806332c2d153146104375780633c351e101461044c5780633cbc795b14610471575f80fd5b806319d8ac61146103a857806326782247146103d55780632b35b3ac1461041a575f80fd5b806305835f37116102ed57806305835f371461033d578063107bf28c1461038657806311e892d41461038e575f80fd5b8062d0295d146103075780630350896314610322575b5f80fd5b61030f610879565b6040519081526020015b60405180910390f35b61032a602081565b60405161ffff9091168152602001610319565b6103796040518060400160405280600881526020017f80808401c9c3809400000000000000000000000000000000000000000000000081525081565b604051610319919061415a565b610379610993565b61039660f981565b60405160ff9091168152602001610319565b6007546103bc9067ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610319565b6001546103f59073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610319565b6008546104279060ff1681565b6040519015158152602001610319565b61044a6104453660046141b7565b610a1f565b005b6008546103f590610100900473ffffffffffffffffffffffffffffffffffffffff1681565b60085461049a907501000000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610319565b6104d67fff0000000000000000000000000000000000000000000000000000000000000081565b6040517fff000000000000000000000000000000000000000000000000000000000000009091168152602001610319565b6007546103bc90700100000000000000000000000000000000900467ffffffffffffffff1681565b6103f57f000000000000000000000000000000000000000000000000000000000000000081565b60085461042790790100000000000000000000000000000000000000000000000000900460ff1681565b61044a61058e3660046141f6565b610aee565b6103796040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525081565b610379610d05565b6103f57f000000000000000000000000000000000000000000000000000000000000000081565b61044a61060c366004614259565b610d12565b61044a610d7b565b61032a601f81565b61030f61062f3660046141f6565b60066020525f908152604090205481565b61030f60055481565b61044a6106573660046142a1565b610e63565b61044a61066a366004614409565b610f2c565b6103796040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525081565b61044a6116cb565b6103f573a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b6103f57f000000000000000000000000000000000000000000000000000000000000000081565b6103796107033660046144b0565b61179d565b61044a611b7b565b61044a61071e3660046142a1565b611c4b565b61030f6405ca1ab1e081565b61044a61073d36600461455f565b611d14565b6007546103bc907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b61039660e481565b61044a610788366004614608565b61267a565b6002546103f59073ffffffffffffffffffffffffffffffffffffffff1681565b6103f57f000000000000000000000000000000000000000000000000000000000000000081565b61030f635ca1ab1e81565b61044a6107ed36600461463a565b61270c565b6103f57f000000000000000000000000000000000000000000000000000000000000000081565b6007546103bc9068010000000000000000900467ffffffffffffffff1681565b61044a610847366004614679565b612c38565b610396601b81565b5f546103f59062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f90819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015610905573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061092991906146c1565b6007549091505f906109679067ffffffffffffffff700100000000000000000000000000000000820481169168010000000000000000900416614705565b67ffffffffffffffff169050805f03610982575f9250505090565b61098c818361472d565b9250505090565b600480546109a090614765565b80601f01602080910402602001604051908101604052809291908181526020018280546109cc90614765565b8015610a175780601f106109ee57610100808354040283529160200191610a17565b820191905f5260205f20905b8154815290600101906020018083116109fa57829003601f168201915b505050505081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610a8e576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596684604051610ae191815260200190565b60405180910390a3505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314610b44576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115610b8b576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bf4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c1891906147b6565b610c815760075467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690821610610c81576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b600380546109a090614765565b600854790100000000000000000000000000000000000000000000000000900460ff16610d6b576040517f821935b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d7683838361301f565b505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314610dd1576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085460ff1615610e0e576040517ff6ba91a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f905f90a1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314610eb9576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610cfa565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610f9b576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54610100900460ff1615808015610fb957505f54600160ff909116105b80610fd25750303b158015610fd257505f5460ff166001145b611063576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156110bf575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b606073ffffffffffffffffffffffffffffffffffffffff8516156112cf576110e68561386d565b6110ef86613979565b6110f887613a7c565b60405160200161110a939291906147d5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f318aee3d00000000000000000000000000000000000000000000000000000000825273ffffffffffffffffffffffffffffffffffffffff87811660048401529092505f9182917f0000000000000000000000000000000000000000000000000000000000000000169063318aee3d9060240160408051808303815f875af11580156111c7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111eb919061480d565b915091508163ffffffff165f1461128757600880547fffffffffffffff000000000000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416027fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff1617750100000000000000000000000000000000000000000063ffffffff8516021790556112cc565b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8a16021790555b50505b6008545f9061131b908890610100810473ffffffffffffffffffffffffffffffffffffffff16907501000000000000000000000000000000000000000000900463ffffffff168561179d565b90505f818051906020012090505f4290505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611396573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113ba91906146c1565b604080515f60208201819052918101869052606080820184905260c086901b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528e901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608882015291925090609c01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152908290528051602090910120600780547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff871617905560058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af1158015611531573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115559190614845565b508c5f60026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b60025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555088600390816115e591906148ad565b5060046115f289826148ad565b5062069780600760186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e604051611652939291906149c5565b60405180910390a150505050505080156116c2575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff16331461171c576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b60605f85858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa5f876040516024016117cf96959493929190614a03565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff70000000000000000000000000000000000000000000000000000000017905283519091506060905f0361191f5760f9601f83516118639190614a65565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525060e4876040516020016119099796959493929190614a80565b6040516020818303038152906040529050611a23565b815161ffff101561195c576040517f248b8f8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160f961196b602083614a65565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b90000000000000000000000000000000000000000000000000000000000008152508588604051602001611a109796959493929190614b62565b6040516020818303038152906040529150505b8051602080830191909120604080515f80825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa158015611a81573d5f803e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611af9576040517fcd16196600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040515f90611b3e9084906405ca1ab1e090635ca1ab1e90601b907fff0000000000000000000000000000000000000000000000000000000000000090602001614c44565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529450505050505b949350505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611bd1576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffff81167901000000000000000000000000000000000000000000000000009182900460ff16159091021790556040517ff32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41905f90a1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611ca1576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610cfa565b60025473ffffffffffffffffffffffffffffffffffffffff163314611d65576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835f819003611da0576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115611ddc576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460055467ffffffffffffffff8083169270010000000000000000000000000000000090041690815f5b858110156121c4575f8b8b83818110611e2357611e23614c9f565b905060800201803603810190611e399190614ccc565b805160608201519192509067ffffffffffffffff1615611fb15785611e5d81614d1f565b9650505f8183602001518460600151604051602001611eb493929190928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a165f90815260069093529120549091508114611f3c576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826060015167ffffffffffffffff16836040015167ffffffffffffffff161015611f92576040517f7f7ab87200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff86165f908152600660205260408120556120ac565b602082015115801590612075575060208201516040517f257b363200000000000000000000000000000000000000000000000000000000815260048101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063257b3632906024016020604051808303815f875af115801561204f573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061207391906146c1565b155b156120ac576040517f73bd668d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8667ffffffffffffffff16826040015167ffffffffffffffff1610806120df575042826040015167ffffffffffffffff16115b15612116576040517fea82791600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602082810151604080850151815193840189905290830184905260608084019290925260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528c901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c0160405160208183030381529060405280519060200120945081604001519650505080806121bc90614d45565b915050611e08565b5060075467ffffffffffffffff680100000000000000009091048116908416111561221b576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff861617905560058290556040517fc7a823e000000000000000000000000000000000000000000000000000000000815285907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063c7a823e0906122cc9086908c908c90600401614dc3565b5f6040518083038186803b1580156122e2575f80fd5b505afa1580156122f4573d5f803e3d5ffd5b505050508167ffffffffffffffff168467ffffffffffffffff16146123fe575f61231e8386614705565b905061233467ffffffffffffffff821683614de5565b91506123b57f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff1661236d610879565b6123779190614df8565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190613b6b565b50600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8716021790555b6124fa337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa15801561248d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124b191906146c1565b6124bb9190614df8565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016929190613c3f565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561255f575f80fd5b505af1158015612571573d5f803e3d5ffd5b50506040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff89166004820152602481018690525f92507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169150639a908e73906044016020604051808303815f875af1158015612611573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126359190614845565b60405190915067ffffffffffffffff8216907f303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce905f90a2505050505050505050505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146126d0576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60036126dc82826148ad565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610cfa919061415a565b60085460ff16612748576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f819003612783576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156127bf576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075467ffffffffffffffff6801000000000000000082048116916127fa918491700100000000000000000000000000000000900416614e0f565b1115612832576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460055470010000000000000000000000000000000090910467ffffffffffffffff16905f5b83811015612abc575f86868381811061287557612875614c9f565b90506020028101906128879190614e22565b61289090614e5e565b90508361289c81614d1f565b825180516020918201208185015160408087015190519499509194505f936128fd9386939101928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89165f90815260069093529120549091508114612985576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff86165f908152600660205260408120556129a9600188614de5565b8403612a185742600760189054906101000a900467ffffffffffffffff1684604001516129d69190614ed9565b67ffffffffffffffff161115612a18576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020838101516040805192830188905282018490526060808301919091524260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016608083015233901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c016040516020818303038152906040528051906020012094505050508080612ab490614d45565b91505061285a565b50612aea7f00000000000000000000000000000000000000000000000000000000000000008461236d610879565b60058190556007805467ffffffffffffffff848116700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffff000000000000000090921642821617919091179091556040517f9a908e730000000000000000000000000000000000000000000000000000000081529084166004820152602481018290525f9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639a908e73906044016020604051808303815f875af1158015612bd5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612bf99190614845565b60405190915067ffffffffffffffff8216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4905f90a2505050505050565b60085460ff16612c74576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa158015612cde573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d0291906146c1565b905081811115612d3e576040517f2354600f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611388831115612d7a576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612dbc73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084613c3f565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e26573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e4a91906146c1565b600780549192506801000000000000000090910467ffffffffffffffff16906008612e7483614d1f565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508484604051612eab929190614efa565b60408051918290038220602083015281018290527fffffffffffffffff0000000000000000000000000000000000000000000000004260c01b166060820152606801604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152918152815160209283012060075468010000000000000000900467ffffffffffffffff165f9081526006909352912055323303612fb9576007546040805183815233602082015260609181018290525f918101919091526801000000000000000090910467ffffffffffffffff16907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319060800160405180910390a2613018565b600760089054906101000a900467ffffffffffffffff1667ffffffffffffffff167ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9318233888860405161300f9493929190614f09565b60405180910390a25b5050505050565b60025473ffffffffffffffffffffffffffffffffffffffff163314613070576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f8190036130ab576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156130e7576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460055467ffffffffffffffff8083169270010000000000000000000000000000000090041690815f5b8581101561351b575f89898381811061312e5761312e614c9f565b90506020028101906131409190614f48565b61314990614f7a565b8051805160209091012060608201519192509067ffffffffffffffff16156132c9578561317581614d1f565b9650505f81836020015184606001516040516020016131cc93929190928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a165f90815260069093529120549091508114613254576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826060015167ffffffffffffffff16836040015167ffffffffffffffff1610156132aa576040517f7f7ab87200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff86165f90815260066020526040812055613403565b60208201511580159061338d575060208201516040517f257b363200000000000000000000000000000000000000000000000000000000815260048101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063257b3632906024016020604051808303815f875af1158015613367573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061338b91906146c1565b155b156133c4576040517f73bd668d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151516201d4c01015613403576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8667ffffffffffffffff16826040015167ffffffffffffffff161080613436575042826040015167ffffffffffffffff16115b1561346d576040517fea82791600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602082810151604080850151815193840189905290830184905260608084019290925260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528a901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c01604051602081830303815290604052805190602001209450816040015196505050808061351390614d45565b915050613113565b5060075467ffffffffffffffff6801000000000000000090910481169084161115613572576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff868116919091179091556005839055859082811690851614613660575f6135c88386614705565b90506135de67ffffffffffffffff821683614de5565b91506136177f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff1661236d610879565b50600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8716021790555b6136ef337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa15801561248d573d5f803e3d5ffd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b158015613754575f80fd5b505af1158015613766573d5f803e3d5ffd5b50506040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff89166004820152602481018690525f92507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169150639a908e73906044016020604051808303815f875af1158015613806573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061382a9190614845565b60405190915067ffffffffffffffff8216907f303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce905f90a250505050505050505050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f06fdde030000000000000000000000000000000000000000000000000000000017905290516060915f91829173ffffffffffffffffffffffffffffffffffffffff8616916138ee9190614fed565b5f60405180830381855afa9150503d805f8114613926576040519150601f19603f3d011682016040523d82523d5f602084013e61392b565b606091505b509150915081613970576040518060400160405280600781526020017f4e4f5f4e414d4500000000000000000000000000000000000000000000000000815250611b73565b611b7381613ca3565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f95d89b410000000000000000000000000000000000000000000000000000000017905290516060915f91829173ffffffffffffffffffffffffffffffffffffffff8616916139fa9190614fed565b5f60405180830381855afa9150503d805f8114613a32576040519150601f19603f3d011682016040523d82523d5f602084013e613a37565b606091505b509150915081613970576040518060400160405280600981526020017f4e4f5f53594d424f4c0000000000000000000000000000000000000000000000815250611b73565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f313ce5670000000000000000000000000000000000000000000000000000000017905290515f918291829173ffffffffffffffffffffffffffffffffffffffff861691613afc9190614fed565b5f60405180830381855afa9150503d805f8114613b34576040519150601f19603f3d011682016040523d82523d5f602084013e613b39565b606091505b5091509150818015613b4c575080516020145b613b57576012611b73565b80806020019051810190611b739190614ffe565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610d769084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613e79565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052613c9d9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401613bbd565b50505050565b60606040825110613cc85781806020019051810190613cc2919061501e565b92915050565b8151602003613e3b575f5b602081108015613d1a5750828181518110613cf057613cf0614c9f565b01602001517fff000000000000000000000000000000000000000000000000000000000000001615155b15613d315780613d2981614d45565b915050613cd3565b805f03613d7357505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e4700000000000000000000000000006020820152919050565b5f8167ffffffffffffffff811115613d8d57613d8d6142cd565b6040519080825280601f01601f191660200182016040528015613db7576020820181803683370190505b5090505f5b82811015613e3357848181518110613dd657613dd6614c9f565b602001015160f81c60f81b828281518110613df357613df3614c9f565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a90535080613e2b81614d45565b915050613dbc565b509392505050565b505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e470000000000000000000000000000602082015290565b919050565b5f613eda826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16613f849092919063ffffffff16565b805190915015610d765780806020019051810190613ef891906147b6565b610d76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161105a565b6060611b7384845f85855f808673ffffffffffffffffffffffffffffffffffffffff168587604051613fb69190614fed565b5f6040518083038185875af1925050503d805f8114613ff0576040519150601f19603f3d011682016040523d82523d5f602084013e613ff5565b606091505b509150915061400687838387614011565b979650505050505050565b606083156140a65782515f0361409f5773ffffffffffffffffffffffffffffffffffffffff85163b61409f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161105a565b5081611b73565b611b7383838151156140bb5781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161105a919061415a565b5f5b838110156141095781810151838201526020016140f1565b50505f910152565b5f81518084526141288160208601602086016140ef565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f61416c6020830184614111565b9392505050565b67ffffffffffffffff81168114614188575f80fd5b50565b73ffffffffffffffffffffffffffffffffffffffff81168114614188575f80fd5b8035613e748161418b565b5f805f606084860312156141c9575f80fd5b83356141d481614173565b92506020840135915060408401356141eb8161418b565b809150509250925092565b5f60208284031215614206575f80fd5b813561416c81614173565b5f8083601f840112614221575f80fd5b50813567ffffffffffffffff811115614238575f80fd5b6020830191508360208260051b8501011115614252575f80fd5b9250929050565b5f805f6040848603121561426b575f80fd5b833567ffffffffffffffff811115614281575f80fd5b61428d86828701614211565b90945092505060208401356141eb8161418b565b5f602082840312156142b1575f80fd5b813561416c8161418b565b63ffffffff81168114614188575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516080810167ffffffffffffffff8111828210171561431d5761431d6142cd565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561436a5761436a6142cd565b604052919050565b5f67ffffffffffffffff82111561438b5761438b6142cd565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f8301126143c6575f80fd5b81356143d96143d482614372565b614323565b8181528460208386010111156143ed575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f8060c0878903121561441e575f80fd5b86356144298161418b565b955060208701356144398161418b565b94506040870135614449816142bc565b935060608701356144598161418b565b9250608087013567ffffffffffffffff80821115614475575f80fd5b6144818a838b016143b7565b935060a0890135915080821115614496575f80fd5b506144a389828a016143b7565b9150509295509295509295565b5f805f80608085870312156144c3575f80fd5b84356144ce816142bc565b935060208501356144de8161418b565b925060408501356144ee816142bc565b9150606085013567ffffffffffffffff811115614509575f80fd5b614515878288016143b7565b91505092959194509250565b5f8083601f840112614531575f80fd5b50813567ffffffffffffffff811115614548575f80fd5b602083019150836020828501011115614252575f80fd5b5f805f805f60608688031215614573575f80fd5b853567ffffffffffffffff8082111561458a575f80fd5b818801915088601f83011261459d575f80fd5b8135818111156145ab575f80fd5b8960208260071b85010111156145bf575f80fd5b602083019750809650506145d5602089016141ac565b945060408801359150808211156145ea575f80fd5b506145f788828901614521565b969995985093965092949392505050565b5f60208284031215614618575f80fd5b813567ffffffffffffffff81111561462e575f80fd5b611b73848285016143b7565b5f806020838503121561464b575f80fd5b823567ffffffffffffffff811115614661575f80fd5b61466d85828601614211565b90969095509350505050565b5f805f6040848603121561468b575f80fd5b833567ffffffffffffffff8111156146a1575f80fd5b6146ad86828701614521565b909790965060209590950135949350505050565b5f602082840312156146d1575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff828116828216039080821115614726576147266146d8565b5092915050565b5f82614760577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b600181811c9082168061477957607f821691505b6020821081036147b0577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f602082840312156147c6575f80fd5b8151801515811461416c575f80fd5b606081525f6147e76060830186614111565b82810360208401526147f98186614111565b91505060ff83166040830152949350505050565b5f806040838503121561481e575f80fd5b8251614829816142bc565b602084015190925061483a8161418b565b809150509250929050565b5f60208284031215614855575f80fd5b815161416c81614173565b601f821115610d76575f81815260208120601f850160051c810160208610156148865750805b601f850160051c820191505b818110156148a557828155600101614892565b505050505050565b815167ffffffffffffffff8111156148c7576148c76142cd565b6148db816148d58454614765565b84614860565b602080601f83116001811461492d575f84156148f75750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556148a5565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156149795788860151825594840194600190910190840161495a565b50858210156149b557878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b606081525f6149d76060830186614111565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b5f63ffffffff808916835273ffffffffffffffffffffffffffffffffffffffff8089166020850152818816604085015280871660608501528086166080850152505060c060a0830152614a5960c0830184614111565b98975050505050505050565b61ffff818116838216019080821115614726576147266146d8565b5f7fff00000000000000000000000000000000000000000000000000000000000000808a60f81b1683527fffff0000000000000000000000000000000000000000000000000000000000008960f01b1660018401528751614ae8816003860160208c016140ef565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614b2b816017840160208b016140ef565b808201915050818660f81b16601782015284519150614b518260188301602088016140ef565b016018019998505050505050505050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b1681525f7fffff000000000000000000000000000000000000000000000000000000000000808960f01b1660018401528751614bca816003860160208c016140ef565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614c0d816017840160208b016140ef565b808201915050818660f01b16601782015284519150614c338260198301602088016140ef565b016019019998505050505050505050565b5f8651614c55818460208b016140ef565b9190910194855250602084019290925260f81b7fff000000000000000000000000000000000000000000000000000000000000009081166040840152166041820152604201919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60808284031215614cdc575f80fd5b614ce46142fa565b82358152602083013560208201526040830135614d0081614173565b60408201526060830135614d1381614173565b60608201529392505050565b5f67ffffffffffffffff808316818103614d3b57614d3b6146d8565b6001019392505050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614d7557614d756146d8565b5060010190565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b838152604060208201525f614ddc604083018486614d7c565b95945050505050565b81810381811115613cc257613cc26146d8565b8082028115828204841417613cc257613cc26146d8565b80820180821115613cc257613cc26146d8565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112614e54575f80fd5b9190910192915050565b5f60608236031215614e6e575f80fd5b6040516060810167ffffffffffffffff8282108183111715614e9257614e926142cd565b816040528435915080821115614ea6575f80fd5b50614eb3368286016143b7565b825250602083013560208201526040830135614ece81614173565b604082015292915050565b67ffffffffffffffff818116838216019080821115614726576147266146d8565b818382375f9101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201525f614f3e606083018486614d7c565b9695505050505050565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112614e54575f80fd5b5f60808236031215614f8a575f80fd5b614f926142fa565b823567ffffffffffffffff811115614fa8575f80fd5b614fb4368286016143b7565b825250602083013560208201526040830135614fcf81614173565b60408201526060830135614fe281614173565b606082015292915050565b5f8251614e548184602087016140ef565b5f6020828403121561500e575f80fd5b815160ff8116811461416c575f80fd5b5f6020828403121561502e575f80fd5b815167ffffffffffffffff811115615044575f80fd5b8201601f81018413615054575f80fd5b80516150626143d482614372565b818152856020838501011115615076575f80fd5b614ddc8260208301602086016140ef56fea26469706673582212202d810c42aa866d408b5990575b5893b8db5a6a7058b00217f497d295aabf656864736f6c63430008140033",
}

// CdkvalidiumABI is the input ABI used to generate the binding from.
// Deprecated: Use CdkvalidiumMetaData.ABI instead.
var CdkvalidiumABI = CdkvalidiumMetaData.ABI

// CdkvalidiumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CdkvalidiumMetaData.Bin instead.
var CdkvalidiumBin = CdkvalidiumMetaData.Bin

// DeployCdkvalidium deploys a new Ethereum contract, binding an instance of Cdkvalidium to it.
func DeployCdkvalidium(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address, _dataCommittee common.Address) (common.Address, *types.Transaction, *Cdkvalidium, error) {
	parsed, err := CdkvalidiumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CdkvalidiumBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager, _dataCommittee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cdkvalidium{CdkvalidiumCaller: CdkvalidiumCaller{contract: contract}, CdkvalidiumTransactor: CdkvalidiumTransactor{contract: contract}, CdkvalidiumFilterer: CdkvalidiumFilterer{contract: contract}}, nil
}

// Cdkvalidium is an auto generated Go binding around an Ethereum contract.
type Cdkvalidium struct {
	CdkvalidiumCaller     // Read-only binding to the contract
	CdkvalidiumTransactor // Write-only binding to the contract
	CdkvalidiumFilterer   // Log filterer for contract events
}

// CdkvalidiumCaller is an auto generated read-only Go binding around an Ethereum contract.
type CdkvalidiumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdkvalidiumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CdkvalidiumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdkvalidiumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CdkvalidiumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdkvalidiumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CdkvalidiumSession struct {
	Contract     *Cdkvalidium      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CdkvalidiumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CdkvalidiumCallerSession struct {
	Contract *CdkvalidiumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CdkvalidiumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CdkvalidiumTransactorSession struct {
	Contract     *CdkvalidiumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CdkvalidiumRaw is an auto generated low-level Go binding around an Ethereum contract.
type CdkvalidiumRaw struct {
	Contract *Cdkvalidium // Generic contract binding to access the raw methods on
}

// CdkvalidiumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CdkvalidiumCallerRaw struct {
	Contract *CdkvalidiumCaller // Generic read-only contract binding to access the raw methods on
}

// CdkvalidiumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CdkvalidiumTransactorRaw struct {
	Contract *CdkvalidiumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCdkvalidium creates a new instance of Cdkvalidium, bound to a specific deployed contract.
func NewCdkvalidium(address common.Address, backend bind.ContractBackend) (*Cdkvalidium, error) {
	contract, err := bindCdkvalidium(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cdkvalidium{CdkvalidiumCaller: CdkvalidiumCaller{contract: contract}, CdkvalidiumTransactor: CdkvalidiumTransactor{contract: contract}, CdkvalidiumFilterer: CdkvalidiumFilterer{contract: contract}}, nil
}

// NewCdkvalidiumCaller creates a new read-only instance of Cdkvalidium, bound to a specific deployed contract.
func NewCdkvalidiumCaller(address common.Address, caller bind.ContractCaller) (*CdkvalidiumCaller, error) {
	contract, err := bindCdkvalidium(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumCaller{contract: contract}, nil
}

// NewCdkvalidiumTransactor creates a new write-only instance of Cdkvalidium, bound to a specific deployed contract.
func NewCdkvalidiumTransactor(address common.Address, transactor bind.ContractTransactor) (*CdkvalidiumTransactor, error) {
	contract, err := bindCdkvalidium(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumTransactor{contract: contract}, nil
}

// NewCdkvalidiumFilterer creates a new log filterer instance of Cdkvalidium, bound to a specific deployed contract.
func NewCdkvalidiumFilterer(address common.Address, filterer bind.ContractFilterer) (*CdkvalidiumFilterer, error) {
	contract, err := bindCdkvalidium(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumFilterer{contract: contract}, nil
}

// bindCdkvalidium binds a generic wrapper to an already deployed contract.
func bindCdkvalidium(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CdkvalidiumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cdkvalidium *CdkvalidiumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cdkvalidium.Contract.CdkvalidiumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cdkvalidium *CdkvalidiumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.CdkvalidiumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cdkvalidium *CdkvalidiumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.CdkvalidiumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cdkvalidium *CdkvalidiumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cdkvalidium.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cdkvalidium *CdkvalidiumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cdkvalidium *CdkvalidiumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Cdkvalidium *CdkvalidiumCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Cdkvalidium *CdkvalidiumSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Cdkvalidium.Contract.GLOBALEXITROOTMANAGERL2(&_Cdkvalidium.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Cdkvalidium *CdkvalidiumCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Cdkvalidium.Contract.GLOBALEXITROOTMANAGERL2(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Cdkvalidium *CdkvalidiumCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Cdkvalidium *CdkvalidiumSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Cdkvalidium.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Cdkvalidium *CdkvalidiumCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Cdkvalidium.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Cdkvalidium *CdkvalidiumCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Cdkvalidium *CdkvalidiumSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Cdkvalidium.Contract.INITIALIZETXBRIDGEPARAMS(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Cdkvalidium *CdkvalidiumCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Cdkvalidium.Contract.INITIALIZETXBRIDGEPARAMS(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Cdkvalidium *CdkvalidiumCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Cdkvalidium *CdkvalidiumSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Cdkvalidium.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Cdkvalidium *CdkvalidiumCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Cdkvalidium.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Cdkvalidium *CdkvalidiumCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Cdkvalidium *CdkvalidiumSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Cdkvalidium.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Cdkvalidium *CdkvalidiumCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Cdkvalidium.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Cdkvalidium *CdkvalidiumCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Cdkvalidium *CdkvalidiumSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Cdkvalidium.Contract.INITIALIZETXCONSTANTBYTES(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Cdkvalidium *CdkvalidiumCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Cdkvalidium.Contract.INITIALIZETXCONSTANTBYTES(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Cdkvalidium *CdkvalidiumCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Cdkvalidium *CdkvalidiumSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Cdkvalidium.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Cdkvalidium *CdkvalidiumCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Cdkvalidium.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Cdkvalidium *CdkvalidiumCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Cdkvalidium *CdkvalidiumSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Cdkvalidium.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Cdkvalidium *CdkvalidiumCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Cdkvalidium.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Cdkvalidium *CdkvalidiumCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Cdkvalidium *CdkvalidiumSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Cdkvalidium.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Cdkvalidium.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Cdkvalidium *CdkvalidiumCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Cdkvalidium.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Cdkvalidium.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Cdkvalidium *CdkvalidiumCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Cdkvalidium *CdkvalidiumSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Cdkvalidium.Contract.SIGNATUREINITIALIZETXR(&_Cdkvalidium.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Cdkvalidium *CdkvalidiumCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Cdkvalidium.Contract.SIGNATUREINITIALIZETXR(&_Cdkvalidium.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Cdkvalidium *CdkvalidiumCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Cdkvalidium *CdkvalidiumSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Cdkvalidium.Contract.SIGNATUREINITIALIZETXS(&_Cdkvalidium.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Cdkvalidium *CdkvalidiumCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Cdkvalidium.Contract.SIGNATUREINITIALIZETXS(&_Cdkvalidium.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Cdkvalidium *CdkvalidiumCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Cdkvalidium *CdkvalidiumSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Cdkvalidium.Contract.SIGNATUREINITIALIZETXV(&_Cdkvalidium.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Cdkvalidium *CdkvalidiumCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Cdkvalidium.Contract.SIGNATUREINITIALIZETXV(&_Cdkvalidium.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Cdkvalidium *CdkvalidiumCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Cdkvalidium *CdkvalidiumSession) Admin() (common.Address, error) {
	return _Cdkvalidium.Contract.Admin(&_Cdkvalidium.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Cdkvalidium *CdkvalidiumCallerSession) Admin() (common.Address, error) {
	return _Cdkvalidium.Contract.Admin(&_Cdkvalidium.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Cdkvalidium *CdkvalidiumCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Cdkvalidium *CdkvalidiumSession) BridgeAddress() (common.Address, error) {
	return _Cdkvalidium.Contract.BridgeAddress(&_Cdkvalidium.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Cdkvalidium *CdkvalidiumCallerSession) BridgeAddress() (common.Address, error) {
	return _Cdkvalidium.Contract.BridgeAddress(&_Cdkvalidium.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Cdkvalidium *CdkvalidiumCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Cdkvalidium *CdkvalidiumSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Cdkvalidium.Contract.CalculatePolPerForceBatch(&_Cdkvalidium.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Cdkvalidium *CdkvalidiumCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Cdkvalidium.Contract.CalculatePolPerForceBatch(&_Cdkvalidium.CallOpts)
}

// DataCommittee is a free data retrieval call binding the contract method 0x56e29435.
//
// Solidity: function dataCommittee() view returns(address)
func (_Cdkvalidium *CdkvalidiumCaller) DataCommittee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "dataCommittee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataCommittee is a free data retrieval call binding the contract method 0x56e29435.
//
// Solidity: function dataCommittee() view returns(address)
func (_Cdkvalidium *CdkvalidiumSession) DataCommittee() (common.Address, error) {
	return _Cdkvalidium.Contract.DataCommittee(&_Cdkvalidium.CallOpts)
}

// DataCommittee is a free data retrieval call binding the contract method 0x56e29435.
//
// Solidity: function dataCommittee() view returns(address)
func (_Cdkvalidium *CdkvalidiumCallerSession) DataCommittee() (common.Address, error) {
	return _Cdkvalidium.Contract.DataCommittee(&_Cdkvalidium.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Cdkvalidium *CdkvalidiumCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Cdkvalidium *CdkvalidiumSession) ForceBatchTimeout() (uint64, error) {
	return _Cdkvalidium.Contract.ForceBatchTimeout(&_Cdkvalidium.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Cdkvalidium *CdkvalidiumCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Cdkvalidium.Contract.ForceBatchTimeout(&_Cdkvalidium.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Cdkvalidium *CdkvalidiumCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Cdkvalidium *CdkvalidiumSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Cdkvalidium.Contract.ForcedBatches(&_Cdkvalidium.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Cdkvalidium *CdkvalidiumCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Cdkvalidium.Contract.ForcedBatches(&_Cdkvalidium.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Cdkvalidium *CdkvalidiumCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Cdkvalidium *CdkvalidiumSession) GasTokenAddress() (common.Address, error) {
	return _Cdkvalidium.Contract.GasTokenAddress(&_Cdkvalidium.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Cdkvalidium *CdkvalidiumCallerSession) GasTokenAddress() (common.Address, error) {
	return _Cdkvalidium.Contract.GasTokenAddress(&_Cdkvalidium.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Cdkvalidium *CdkvalidiumCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Cdkvalidium *CdkvalidiumSession) GasTokenNetwork() (uint32, error) {
	return _Cdkvalidium.Contract.GasTokenNetwork(&_Cdkvalidium.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Cdkvalidium *CdkvalidiumCallerSession) GasTokenNetwork() (uint32, error) {
	return _Cdkvalidium.Contract.GasTokenNetwork(&_Cdkvalidium.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Cdkvalidium *CdkvalidiumCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Cdkvalidium *CdkvalidiumSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Cdkvalidium.Contract.GenerateInitializeTransaction(&_Cdkvalidium.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Cdkvalidium *CdkvalidiumCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Cdkvalidium.Contract.GenerateInitializeTransaction(&_Cdkvalidium.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Cdkvalidium *CdkvalidiumCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Cdkvalidium *CdkvalidiumSession) GlobalExitRootManager() (common.Address, error) {
	return _Cdkvalidium.Contract.GlobalExitRootManager(&_Cdkvalidium.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Cdkvalidium *CdkvalidiumCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Cdkvalidium.Contract.GlobalExitRootManager(&_Cdkvalidium.CallOpts)
}

// IsForcedBatchAllowed is a free data retrieval call binding the contract method 0x2b35b3ac.
//
// Solidity: function isForcedBatchAllowed() view returns(bool)
func (_Cdkvalidium *CdkvalidiumCaller) IsForcedBatchAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "isForcedBatchAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForcedBatchAllowed is a free data retrieval call binding the contract method 0x2b35b3ac.
//
// Solidity: function isForcedBatchAllowed() view returns(bool)
func (_Cdkvalidium *CdkvalidiumSession) IsForcedBatchAllowed() (bool, error) {
	return _Cdkvalidium.Contract.IsForcedBatchAllowed(&_Cdkvalidium.CallOpts)
}

// IsForcedBatchAllowed is a free data retrieval call binding the contract method 0x2b35b3ac.
//
// Solidity: function isForcedBatchAllowed() view returns(bool)
func (_Cdkvalidium *CdkvalidiumCallerSession) IsForcedBatchAllowed() (bool, error) {
	return _Cdkvalidium.Contract.IsForcedBatchAllowed(&_Cdkvalidium.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Cdkvalidium *CdkvalidiumCaller) IsSequenceWithDataAvailabilityAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "isSequenceWithDataAvailabilityAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Cdkvalidium *CdkvalidiumSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Cdkvalidium.Contract.IsSequenceWithDataAvailabilityAllowed(&_Cdkvalidium.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Cdkvalidium *CdkvalidiumCallerSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Cdkvalidium.Contract.IsSequenceWithDataAvailabilityAllowed(&_Cdkvalidium.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Cdkvalidium *CdkvalidiumCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Cdkvalidium *CdkvalidiumSession) LastAccInputHash() ([32]byte, error) {
	return _Cdkvalidium.Contract.LastAccInputHash(&_Cdkvalidium.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Cdkvalidium *CdkvalidiumCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Cdkvalidium.Contract.LastAccInputHash(&_Cdkvalidium.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Cdkvalidium *CdkvalidiumCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Cdkvalidium *CdkvalidiumSession) LastForceBatch() (uint64, error) {
	return _Cdkvalidium.Contract.LastForceBatch(&_Cdkvalidium.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Cdkvalidium *CdkvalidiumCallerSession) LastForceBatch() (uint64, error) {
	return _Cdkvalidium.Contract.LastForceBatch(&_Cdkvalidium.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Cdkvalidium *CdkvalidiumCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Cdkvalidium *CdkvalidiumSession) LastForceBatchSequenced() (uint64, error) {
	return _Cdkvalidium.Contract.LastForceBatchSequenced(&_Cdkvalidium.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Cdkvalidium *CdkvalidiumCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Cdkvalidium.Contract.LastForceBatchSequenced(&_Cdkvalidium.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Cdkvalidium *CdkvalidiumCaller) LastTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "lastTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Cdkvalidium *CdkvalidiumSession) LastTimestamp() (uint64, error) {
	return _Cdkvalidium.Contract.LastTimestamp(&_Cdkvalidium.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Cdkvalidium *CdkvalidiumCallerSession) LastTimestamp() (uint64, error) {
	return _Cdkvalidium.Contract.LastTimestamp(&_Cdkvalidium.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Cdkvalidium *CdkvalidiumCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Cdkvalidium *CdkvalidiumSession) NetworkName() (string, error) {
	return _Cdkvalidium.Contract.NetworkName(&_Cdkvalidium.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Cdkvalidium *CdkvalidiumCallerSession) NetworkName() (string, error) {
	return _Cdkvalidium.Contract.NetworkName(&_Cdkvalidium.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Cdkvalidium *CdkvalidiumCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Cdkvalidium *CdkvalidiumSession) PendingAdmin() (common.Address, error) {
	return _Cdkvalidium.Contract.PendingAdmin(&_Cdkvalidium.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Cdkvalidium *CdkvalidiumCallerSession) PendingAdmin() (common.Address, error) {
	return _Cdkvalidium.Contract.PendingAdmin(&_Cdkvalidium.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Cdkvalidium *CdkvalidiumCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Cdkvalidium *CdkvalidiumSession) Pol() (common.Address, error) {
	return _Cdkvalidium.Contract.Pol(&_Cdkvalidium.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Cdkvalidium *CdkvalidiumCallerSession) Pol() (common.Address, error) {
	return _Cdkvalidium.Contract.Pol(&_Cdkvalidium.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Cdkvalidium *CdkvalidiumCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Cdkvalidium *CdkvalidiumSession) RollupManager() (common.Address, error) {
	return _Cdkvalidium.Contract.RollupManager(&_Cdkvalidium.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Cdkvalidium *CdkvalidiumCallerSession) RollupManager() (common.Address, error) {
	return _Cdkvalidium.Contract.RollupManager(&_Cdkvalidium.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Cdkvalidium *CdkvalidiumCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Cdkvalidium *CdkvalidiumSession) TrustedSequencer() (common.Address, error) {
	return _Cdkvalidium.Contract.TrustedSequencer(&_Cdkvalidium.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Cdkvalidium *CdkvalidiumCallerSession) TrustedSequencer() (common.Address, error) {
	return _Cdkvalidium.Contract.TrustedSequencer(&_Cdkvalidium.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Cdkvalidium *CdkvalidiumCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cdkvalidium.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Cdkvalidium *CdkvalidiumSession) TrustedSequencerURL() (string, error) {
	return _Cdkvalidium.Contract.TrustedSequencerURL(&_Cdkvalidium.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Cdkvalidium *CdkvalidiumCallerSession) TrustedSequencerURL() (string, error) {
	return _Cdkvalidium.Contract.TrustedSequencerURL(&_Cdkvalidium.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Cdkvalidium *CdkvalidiumTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Cdkvalidium *CdkvalidiumSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Cdkvalidium.Contract.AcceptAdminRole(&_Cdkvalidium.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Cdkvalidium.Contract.AcceptAdminRole(&_Cdkvalidium.TransactOpts)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Cdkvalidium *CdkvalidiumTransactor) ActivateForceBatches(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "activateForceBatches")
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Cdkvalidium *CdkvalidiumSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Cdkvalidium.Contract.ActivateForceBatches(&_Cdkvalidium.TransactOpts)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Cdkvalidium.Contract.ActivateForceBatches(&_Cdkvalidium.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Cdkvalidium *CdkvalidiumTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Cdkvalidium *CdkvalidiumSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.ForceBatch(&_Cdkvalidium.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.ForceBatch(&_Cdkvalidium.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Cdkvalidium *CdkvalidiumTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Cdkvalidium *CdkvalidiumSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.Initialize(&_Cdkvalidium.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.Initialize(&_Cdkvalidium.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Cdkvalidium *CdkvalidiumTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Cdkvalidium *CdkvalidiumSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.OnVerifyBatches(&_Cdkvalidium.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.OnVerifyBatches(&_Cdkvalidium.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x5e9145c9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,uint64)[] batches, address l2Coinbase) returns()
func (_Cdkvalidium *CdkvalidiumTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "sequenceBatches", batches, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x5e9145c9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,uint64)[] batches, address l2Coinbase) returns()
func (_Cdkvalidium *CdkvalidiumSession) SequenceBatches(batches []PolygonRollupBaseBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SequenceBatches(&_Cdkvalidium.TransactOpts, batches, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x5e9145c9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,uint64)[] batches, address l2Coinbase) returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) SequenceBatches(batches []PolygonRollupBaseBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SequenceBatches(&_Cdkvalidium.TransactOpts, batches, l2Coinbase)
}

// SequenceBatchesDataCommittee is a paid mutator transaction binding the contract method 0xb81fcafd.
//
// Solidity: function sequenceBatchesDataCommittee((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Cdkvalidium *CdkvalidiumTransactor) SequenceBatchesDataCommittee(opts *bind.TransactOpts, batches []PolygonDataComitteeValidiumBatchData, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "sequenceBatchesDataCommittee", batches, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesDataCommittee is a paid mutator transaction binding the contract method 0xb81fcafd.
//
// Solidity: function sequenceBatchesDataCommittee((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Cdkvalidium *CdkvalidiumSession) SequenceBatchesDataCommittee(batches []PolygonDataComitteeValidiumBatchData, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SequenceBatchesDataCommittee(&_Cdkvalidium.TransactOpts, batches, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesDataCommittee is a paid mutator transaction binding the contract method 0xb81fcafd.
//
// Solidity: function sequenceBatchesDataCommittee((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) SequenceBatchesDataCommittee(batches []PolygonDataComitteeValidiumBatchData, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SequenceBatchesDataCommittee(&_Cdkvalidium.TransactOpts, batches, l2Coinbase, dataAvailabilityMessage)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Cdkvalidium *CdkvalidiumTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseForcedBatchData) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Cdkvalidium *CdkvalidiumSession) SequenceForceBatches(batches []PolygonRollupBaseForcedBatchData) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SequenceForceBatches(&_Cdkvalidium.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseForcedBatchData) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SequenceForceBatches(&_Cdkvalidium.TransactOpts, batches)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Cdkvalidium *CdkvalidiumTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Cdkvalidium *CdkvalidiumSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SetForceBatchTimeout(&_Cdkvalidium.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SetForceBatchTimeout(&_Cdkvalidium.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Cdkvalidium *CdkvalidiumTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Cdkvalidium *CdkvalidiumSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SetTrustedSequencer(&_Cdkvalidium.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SetTrustedSequencer(&_Cdkvalidium.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Cdkvalidium *CdkvalidiumTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Cdkvalidium *CdkvalidiumSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SetTrustedSequencerURL(&_Cdkvalidium.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SetTrustedSequencerURL(&_Cdkvalidium.TransactOpts, newTrustedSequencerURL)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0xaa3587d3.
//
// Solidity: function switchSequenceWithDataAvailability() returns()
func (_Cdkvalidium *CdkvalidiumTransactor) SwitchSequenceWithDataAvailability(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "switchSequenceWithDataAvailability")
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0xaa3587d3.
//
// Solidity: function switchSequenceWithDataAvailability() returns()
func (_Cdkvalidium *CdkvalidiumSession) SwitchSequenceWithDataAvailability() (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SwitchSequenceWithDataAvailability(&_Cdkvalidium.TransactOpts)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0xaa3587d3.
//
// Solidity: function switchSequenceWithDataAvailability() returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) SwitchSequenceWithDataAvailability() (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SwitchSequenceWithDataAvailability(&_Cdkvalidium.TransactOpts)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Cdkvalidium *CdkvalidiumTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Cdkvalidium *CdkvalidiumSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.TransferAdminRole(&_Cdkvalidium.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.TransferAdminRole(&_Cdkvalidium.TransactOpts, newPendingAdmin)
}

// CdkvalidiumAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Cdkvalidium contract.
type CdkvalidiumAcceptAdminRoleIterator struct {
	Event *CdkvalidiumAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumAcceptAdminRole)
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
		it.Event = new(CdkvalidiumAcceptAdminRole)
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
func (it *CdkvalidiumAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumAcceptAdminRole represents a AcceptAdminRole event raised by the Cdkvalidium contract.
type CdkvalidiumAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Cdkvalidium *CdkvalidiumFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*CdkvalidiumAcceptAdminRoleIterator, error) {

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumAcceptAdminRoleIterator{contract: _Cdkvalidium.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Cdkvalidium *CdkvalidiumFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *CdkvalidiumAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumAcceptAdminRole)
				if err := _Cdkvalidium.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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

// ParseAcceptAdminRole is a log parse operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Cdkvalidium *CdkvalidiumFilterer) ParseAcceptAdminRole(log types.Log) (*CdkvalidiumAcceptAdminRole, error) {
	event := new(CdkvalidiumAcceptAdminRole)
	if err := _Cdkvalidium.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkvalidiumActivateForceBatchesIterator is returned from FilterActivateForceBatches and is used to iterate over the raw logs and unpacked data for ActivateForceBatches events raised by the Cdkvalidium contract.
type CdkvalidiumActivateForceBatchesIterator struct {
	Event *CdkvalidiumActivateForceBatches // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumActivateForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumActivateForceBatches)
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
		it.Event = new(CdkvalidiumActivateForceBatches)
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
func (it *CdkvalidiumActivateForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumActivateForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumActivateForceBatches represents a ActivateForceBatches event raised by the Cdkvalidium contract.
type CdkvalidiumActivateForceBatches struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActivateForceBatches is a free log retrieval operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Cdkvalidium *CdkvalidiumFilterer) FilterActivateForceBatches(opts *bind.FilterOpts) (*CdkvalidiumActivateForceBatchesIterator, error) {

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumActivateForceBatchesIterator{contract: _Cdkvalidium.contract, event: "ActivateForceBatches", logs: logs, sub: sub}, nil
}

// WatchActivateForceBatches is a free log subscription operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Cdkvalidium *CdkvalidiumFilterer) WatchActivateForceBatches(opts *bind.WatchOpts, sink chan<- *CdkvalidiumActivateForceBatches) (event.Subscription, error) {

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumActivateForceBatches)
				if err := _Cdkvalidium.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
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

// ParseActivateForceBatches is a log parse operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Cdkvalidium *CdkvalidiumFilterer) ParseActivateForceBatches(log types.Log) (*CdkvalidiumActivateForceBatches, error) {
	event := new(CdkvalidiumActivateForceBatches)
	if err := _Cdkvalidium.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkvalidiumForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Cdkvalidium contract.
type CdkvalidiumForceBatchIterator struct {
	Event *CdkvalidiumForceBatch // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumForceBatch)
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
		it.Event = new(CdkvalidiumForceBatch)
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
func (it *CdkvalidiumForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumForceBatch represents a ForceBatch event raised by the Cdkvalidium contract.
type CdkvalidiumForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Cdkvalidium *CdkvalidiumFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*CdkvalidiumForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumForceBatchIterator{contract: _Cdkvalidium.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Cdkvalidium *CdkvalidiumFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *CdkvalidiumForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumForceBatch)
				if err := _Cdkvalidium.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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

// ParseForceBatch is a log parse operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Cdkvalidium *CdkvalidiumFilterer) ParseForceBatch(log types.Log) (*CdkvalidiumForceBatch, error) {
	event := new(CdkvalidiumForceBatch)
	if err := _Cdkvalidium.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkvalidiumInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the Cdkvalidium contract.
type CdkvalidiumInitialSequenceBatchesIterator struct {
	Event *CdkvalidiumInitialSequenceBatches // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumInitialSequenceBatches)
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
		it.Event = new(CdkvalidiumInitialSequenceBatches)
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
func (it *CdkvalidiumInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumInitialSequenceBatches represents a InitialSequenceBatches event raised by the Cdkvalidium contract.
type CdkvalidiumInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Cdkvalidium *CdkvalidiumFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*CdkvalidiumInitialSequenceBatchesIterator, error) {

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumInitialSequenceBatchesIterator{contract: _Cdkvalidium.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Cdkvalidium *CdkvalidiumFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *CdkvalidiumInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumInitialSequenceBatches)
				if err := _Cdkvalidium.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
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

// ParseInitialSequenceBatches is a log parse operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Cdkvalidium *CdkvalidiumFilterer) ParseInitialSequenceBatches(log types.Log) (*CdkvalidiumInitialSequenceBatches, error) {
	event := new(CdkvalidiumInitialSequenceBatches)
	if err := _Cdkvalidium.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkvalidiumInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Cdkvalidium contract.
type CdkvalidiumInitializedIterator struct {
	Event *CdkvalidiumInitialized // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumInitialized)
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
		it.Event = new(CdkvalidiumInitialized)
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
func (it *CdkvalidiumInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumInitialized represents a Initialized event raised by the Cdkvalidium contract.
type CdkvalidiumInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cdkvalidium *CdkvalidiumFilterer) FilterInitialized(opts *bind.FilterOpts) (*CdkvalidiumInitializedIterator, error) {

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumInitializedIterator{contract: _Cdkvalidium.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cdkvalidium *CdkvalidiumFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CdkvalidiumInitialized) (event.Subscription, error) {

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumInitialized)
				if err := _Cdkvalidium.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Cdkvalidium *CdkvalidiumFilterer) ParseInitialized(log types.Log) (*CdkvalidiumInitialized, error) {
	event := new(CdkvalidiumInitialized)
	if err := _Cdkvalidium.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkvalidiumSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Cdkvalidium contract.
type CdkvalidiumSequenceBatchesIterator struct {
	Event *CdkvalidiumSequenceBatches // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumSequenceBatches)
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
		it.Event = new(CdkvalidiumSequenceBatches)
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
func (it *CdkvalidiumSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumSequenceBatches represents a SequenceBatches event raised by the Cdkvalidium contract.
type CdkvalidiumSequenceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Cdkvalidium *CdkvalidiumFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*CdkvalidiumSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumSequenceBatchesIterator{contract: _Cdkvalidium.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Cdkvalidium *CdkvalidiumFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *CdkvalidiumSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumSequenceBatches)
				if err := _Cdkvalidium.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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

// ParseSequenceBatches is a log parse operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Cdkvalidium *CdkvalidiumFilterer) ParseSequenceBatches(log types.Log) (*CdkvalidiumSequenceBatches, error) {
	event := new(CdkvalidiumSequenceBatches)
	if err := _Cdkvalidium.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkvalidiumSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Cdkvalidium contract.
type CdkvalidiumSequenceForceBatchesIterator struct {
	Event *CdkvalidiumSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumSequenceForceBatches)
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
		it.Event = new(CdkvalidiumSequenceForceBatches)
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
func (it *CdkvalidiumSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumSequenceForceBatches represents a SequenceForceBatches event raised by the Cdkvalidium contract.
type CdkvalidiumSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Cdkvalidium *CdkvalidiumFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*CdkvalidiumSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumSequenceForceBatchesIterator{contract: _Cdkvalidium.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Cdkvalidium *CdkvalidiumFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *CdkvalidiumSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumSequenceForceBatches)
				if err := _Cdkvalidium.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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

// ParseSequenceForceBatches is a log parse operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Cdkvalidium *CdkvalidiumFilterer) ParseSequenceForceBatches(log types.Log) (*CdkvalidiumSequenceForceBatches, error) {
	event := new(CdkvalidiumSequenceForceBatches)
	if err := _Cdkvalidium.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkvalidiumSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Cdkvalidium contract.
type CdkvalidiumSetForceBatchTimeoutIterator struct {
	Event *CdkvalidiumSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumSetForceBatchTimeout)
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
		it.Event = new(CdkvalidiumSetForceBatchTimeout)
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
func (it *CdkvalidiumSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Cdkvalidium contract.
type CdkvalidiumSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Cdkvalidium *CdkvalidiumFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*CdkvalidiumSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumSetForceBatchTimeoutIterator{contract: _Cdkvalidium.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Cdkvalidium *CdkvalidiumFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *CdkvalidiumSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumSetForceBatchTimeout)
				if err := _Cdkvalidium.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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

// ParseSetForceBatchTimeout is a log parse operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Cdkvalidium *CdkvalidiumFilterer) ParseSetForceBatchTimeout(log types.Log) (*CdkvalidiumSetForceBatchTimeout, error) {
	event := new(CdkvalidiumSetForceBatchTimeout)
	if err := _Cdkvalidium.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkvalidiumSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Cdkvalidium contract.
type CdkvalidiumSetTrustedSequencerIterator struct {
	Event *CdkvalidiumSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumSetTrustedSequencer)
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
		it.Event = new(CdkvalidiumSetTrustedSequencer)
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
func (it *CdkvalidiumSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumSetTrustedSequencer represents a SetTrustedSequencer event raised by the Cdkvalidium contract.
type CdkvalidiumSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Cdkvalidium *CdkvalidiumFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*CdkvalidiumSetTrustedSequencerIterator, error) {

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumSetTrustedSequencerIterator{contract: _Cdkvalidium.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Cdkvalidium *CdkvalidiumFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *CdkvalidiumSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumSetTrustedSequencer)
				if err := _Cdkvalidium.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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

// ParseSetTrustedSequencer is a log parse operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Cdkvalidium *CdkvalidiumFilterer) ParseSetTrustedSequencer(log types.Log) (*CdkvalidiumSetTrustedSequencer, error) {
	event := new(CdkvalidiumSetTrustedSequencer)
	if err := _Cdkvalidium.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkvalidiumSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Cdkvalidium contract.
type CdkvalidiumSetTrustedSequencerURLIterator struct {
	Event *CdkvalidiumSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumSetTrustedSequencerURL)
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
		it.Event = new(CdkvalidiumSetTrustedSequencerURL)
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
func (it *CdkvalidiumSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Cdkvalidium contract.
type CdkvalidiumSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Cdkvalidium *CdkvalidiumFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*CdkvalidiumSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumSetTrustedSequencerURLIterator{contract: _Cdkvalidium.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Cdkvalidium *CdkvalidiumFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *CdkvalidiumSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumSetTrustedSequencerURL)
				if err := _Cdkvalidium.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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

// ParseSetTrustedSequencerURL is a log parse operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Cdkvalidium *CdkvalidiumFilterer) ParseSetTrustedSequencerURL(log types.Log) (*CdkvalidiumSetTrustedSequencerURL, error) {
	event := new(CdkvalidiumSetTrustedSequencerURL)
	if err := _Cdkvalidium.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkvalidiumSwitchSequenceWithDataAvailabilityIterator is returned from FilterSwitchSequenceWithDataAvailability and is used to iterate over the raw logs and unpacked data for SwitchSequenceWithDataAvailability events raised by the Cdkvalidium contract.
type CdkvalidiumSwitchSequenceWithDataAvailabilityIterator struct {
	Event *CdkvalidiumSwitchSequenceWithDataAvailability // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumSwitchSequenceWithDataAvailabilityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumSwitchSequenceWithDataAvailability)
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
		it.Event = new(CdkvalidiumSwitchSequenceWithDataAvailability)
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
func (it *CdkvalidiumSwitchSequenceWithDataAvailabilityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumSwitchSequenceWithDataAvailabilityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumSwitchSequenceWithDataAvailability represents a SwitchSequenceWithDataAvailability event raised by the Cdkvalidium contract.
type CdkvalidiumSwitchSequenceWithDataAvailability struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSwitchSequenceWithDataAvailability is a free log retrieval operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Cdkvalidium *CdkvalidiumFilterer) FilterSwitchSequenceWithDataAvailability(opts *bind.FilterOpts) (*CdkvalidiumSwitchSequenceWithDataAvailabilityIterator, error) {

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumSwitchSequenceWithDataAvailabilityIterator{contract: _Cdkvalidium.contract, event: "SwitchSequenceWithDataAvailability", logs: logs, sub: sub}, nil
}

// WatchSwitchSequenceWithDataAvailability is a free log subscription operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Cdkvalidium *CdkvalidiumFilterer) WatchSwitchSequenceWithDataAvailability(opts *bind.WatchOpts, sink chan<- *CdkvalidiumSwitchSequenceWithDataAvailability) (event.Subscription, error) {

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumSwitchSequenceWithDataAvailability)
				if err := _Cdkvalidium.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
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

// ParseSwitchSequenceWithDataAvailability is a log parse operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Cdkvalidium *CdkvalidiumFilterer) ParseSwitchSequenceWithDataAvailability(log types.Log) (*CdkvalidiumSwitchSequenceWithDataAvailability, error) {
	event := new(CdkvalidiumSwitchSequenceWithDataAvailability)
	if err := _Cdkvalidium.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkvalidiumTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Cdkvalidium contract.
type CdkvalidiumTransferAdminRoleIterator struct {
	Event *CdkvalidiumTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumTransferAdminRole)
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
		it.Event = new(CdkvalidiumTransferAdminRole)
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
func (it *CdkvalidiumTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumTransferAdminRole represents a TransferAdminRole event raised by the Cdkvalidium contract.
type CdkvalidiumTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Cdkvalidium *CdkvalidiumFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*CdkvalidiumTransferAdminRoleIterator, error) {

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumTransferAdminRoleIterator{contract: _Cdkvalidium.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Cdkvalidium *CdkvalidiumFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *CdkvalidiumTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumTransferAdminRole)
				if err := _Cdkvalidium.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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

// ParseTransferAdminRole is a log parse operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Cdkvalidium *CdkvalidiumFilterer) ParseTransferAdminRole(log types.Log) (*CdkvalidiumTransferAdminRole, error) {
	event := new(CdkvalidiumTransferAdminRole)
	if err := _Cdkvalidium.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkvalidiumVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Cdkvalidium contract.
type CdkvalidiumVerifyBatchesIterator struct {
	Event *CdkvalidiumVerifyBatches // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumVerifyBatches)
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
		it.Event = new(CdkvalidiumVerifyBatches)
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
func (it *CdkvalidiumVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumVerifyBatches represents a VerifyBatches event raised by the Cdkvalidium contract.
type CdkvalidiumVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Cdkvalidium *CdkvalidiumFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*CdkvalidiumVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumVerifyBatchesIterator{contract: _Cdkvalidium.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Cdkvalidium *CdkvalidiumFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *CdkvalidiumVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumVerifyBatches)
				if err := _Cdkvalidium.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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

// ParseVerifyBatches is a log parse operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Cdkvalidium *CdkvalidiumFilterer) ParseVerifyBatches(log types.Log) (*CdkvalidiumVerifyBatches, error) {
	event := new(CdkvalidiumVerifyBatches)
	if err := _Cdkvalidium.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
