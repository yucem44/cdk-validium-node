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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceWithDataAvailabilityNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ActivateForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDataCommittee\",\"type\":\"address\"}],\"name\":\"SetDataCommittee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SwitchSequenceWithDataAvailability\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataCommittee\",\"outputs\":[{\"internalType\":\"contractICDKDataCommittee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isForcedBatchAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSequenceWithDataAvailabilityAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonRollupBase.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonDataComittee.ValidiumBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataAvailabilityMessage\",\"type\":\"bytes\"}],\"name\":\"sequenceBatchesDataCommittee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonRollupBase.ForcedBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICDKDataCommittee\",\"name\":\"newDataCommittee\",\"type\":\"address\"}],\"name\":\"setDataCommittee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchSequenceWithDataAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200545238038062005452833981016040819052620000359162000071565b6001600160a01b0393841660a052918316608052821660c0521660e052620000d9565b6001600160a01b03811681146200006e57600080fd5b50565b600080600080608085870312156200008857600080fd5b8451620000958162000058565b6020860151909450620000a88162000058565b6040860151909350620000bb8162000058565b6060860151909250620000ce8162000058565b939692955090935050565b60805160a05160c05160e05161527b620001d76000396000818161055101528181610a4f01528181610bbc01528181610f5d01528181611516015281816123ae01528181612477015281816124990152818161263001528181612bd001528181612c9e01528181612d8b0152818161370a0152818161378b015281816137ad01526138d70152600081816106f8015281816111b7015281816119af01528181611ab70152818161257101526138180152600081816107d70152818161136c0152818161211001528181612ed6015261342d01526000818161081c015281816108e701528181612401015281816125470152612eaa015261527b6000f3fe608060405234801561001057600080fd5b506004361061031f5760003560e01c80636ff512cc116101a7578063c754c7ed116100ee578063d8d1091b11610097578063eaeb077b11610071578063eaeb077b1461085e578063f35dda4714610871578063f851a4401461087957600080fd5b8063d8d1091b14610804578063e46761c414610817578063e7a7ed021461083e57600080fd5b8063cfa8ed47116100c8578063cfa8ed47146107b2578063d02103ca146107d2578063d7bc90ff146107f957600080fd5b8063c754c7ed14610767578063c7fffd4b14610797578063c89e42df1461079f57600080fd5b8063a3c573eb11610150578063ada8f9191161012a578063ada8f91914610735578063b0afe15414610748578063b81fcafd1461075457600080fd5b8063a3c573eb146106f3578063a652f26c1461071a578063aa3587d31461072d57600080fd5b80638c3d7301116101815780638c3d7301146106bd57806398cf4bd0146106c55780639e001877146106d857600080fd5b80636ff512cc1461065b578063712570221461066e5780637a5460c51461068157600080fd5b8063456052671161026b57806356e2943511610214578063676870d2116101ee578063676870d21461062a5780636b8616ce146106325780636e05d2cd1461065257600080fd5b806356e29435146105ef5780635e9145c91461060f5780635ec919581461062257600080fd5b80634e487706116102455780634e4877061461059857806352bdeb6d146105ab578063542028d5146105e757600080fd5b8063456052671461052457806349b7b8021461054c5780634c21fef31461057357600080fd5b806326782247116102cd5780633c351e10116102a75780633c351e10146104695780633cbc795b1461048e57806340b5de6c146104cc57600080fd5b806326782247146103f25780632b35b3ac1461043757806332c2d1531461045457600080fd5b8063107bf28c116102fe578063107bf28c146103a357806311e892d4146103ab57806319d8ac61146103c557600080fd5b8062d0295d14610324578063035089631461033f57806305835f371461035a575b600080fd5b61032c61089f565b6040519081526020015b60405180910390f35b610347602081565b60405161ffff9091168152602001610336565b6103966040518060400160405280600881526020017f80808401c9c3809400000000000000000000000000000000000000000000000081525081565b60405161033691906142a6565b6103966109bf565b6103b360f981565b60405160ff9091168152602001610336565b6007546103d99067ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610336565b6001546104129073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610336565b6008546104449060ff1681565b6040519015158152602001610336565b610467610462366004614306565b610a4d565b005b60085461041290610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6008546104b7907501000000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610336565b6104f37fff0000000000000000000000000000000000000000000000000000000000000081565b6040517fff000000000000000000000000000000000000000000000000000000000000009091168152602001610336565b6007546103d990700100000000000000000000000000000000900467ffffffffffffffff1681565b6104127f000000000000000000000000000000000000000000000000000000000000000081565b6009546104449074010000000000000000000000000000000000000000900460ff1681565b6104676105a6366004614348565b610b1c565b6103966040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525081565b610396610d36565b6009546104129073ffffffffffffffffffffffffffffffffffffffff1681565b61046761061d3660046143b1565b610d43565b610467610da7565b610347601f81565b61032c610640366004614348565b60066020526000908152604090205481565b61032c60055481565b6104676106693660046143fd565b610e91565b61046761067c36600461456f565b610f5b565b6103966040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525081565b610467611711565b6104676106d33660046143fd565b6117e4565b61041273a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b6104127f000000000000000000000000000000000000000000000000000000000000000081565b61039661072836600461461c565b6118ae565b610467611c93565b6104676107433660046143fd565b611d60565b61032c6405ca1ab1e081565b6104676107623660046146d3565b611e2a565b6007546103d9907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b6103b360e481565b6104676107ad366004614785565b612781565b6002546104129073ffffffffffffffffffffffffffffffffffffffff1681565b6104127f000000000000000000000000000000000000000000000000000000000000000081565b61032c635ca1ab1e81565b6104676108123660046147ba565b612814565b6104127f000000000000000000000000000000000000000000000000000000000000000081565b6007546103d99068010000000000000000900467ffffffffffffffff1681565b61046761086c3660046147fc565b612d4b565b6103b3601b81565b6000546104129062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600090819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561092e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109529190614848565b6007549091506000906109919067ffffffffffffffff700100000000000000000000000000000000820481169168010000000000000000900416614890565b67ffffffffffffffff169050806000036109ae5760009250505090565b6109b881836148b8565b9250505090565b600480546109cc906148f3565b80601f01602080910402602001604051908101604052809291908181526020018280546109f8906148f3565b8015610a455780601f10610a1a57610100808354040283529160200191610a45565b820191906000526020600020905b815481529060010190602001808311610a2857829003601f168201915b505050505081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610abc576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596684604051610b0f91815260200190565b60405180910390a3505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610b73576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115610bba576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c25573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c499190614946565b610cb25760075467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690821610610cb2576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b600380546109cc906148f3565b60095474010000000000000000000000000000000000000000900460ff16610d97576040517f821935b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610da283838361313a565b505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610dfe576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085460ff1615610e3b576040517ff6ba91a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f90600090a1565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610ee8576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610d2b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314610fca576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610100900460ff1615808015610fea5750600054600160ff909116105b806110045750303b158015611004575060005460ff166001145b611095576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156110f357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b606073ffffffffffffffffffffffffffffffffffffffff8516156113085761111a8561399e565b61112386613aae565b61112c87613bb5565b60405160200161113e93929190614968565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f318aee3d00000000000000000000000000000000000000000000000000000000825273ffffffffffffffffffffffffffffffffffffffff878116600484015290925060009182917f0000000000000000000000000000000000000000000000000000000000000000169063318aee3d9060240160408051808303816000875af11580156111ff573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061122391906149a1565b915091508163ffffffff166000146112c057600880547fffffffffffffff000000000000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416027fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff1617750100000000000000000000000000000000000000000063ffffffff851602179055611305565b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8a16021790555b50505b600854600090611355908890610100810473ffffffffffffffffffffffffffffffffffffffff16907501000000000000000000000000000000000000000000900463ffffffff16856118ae565b9050600081805190602001209050600042905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113f99190614848565b60408051600060208201819052918101869052606080820184905260c086901b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528e901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608882015291925090609c01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152908290528051602090910120600780547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff871617905560058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303816000875af1158015611574573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061159891906149db565b508c600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550886003908161162a9190614a46565b5060046116378982614a46565b5062069780600760186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e60405161169793929190614b60565b60405180910390a1505050505050801561170857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611762576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600154600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff16331461183b576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f76d12ecc41da43ceb240d60fdd4d9fc0baae9793060e7b592abde7ece5ee965190602001610d2b565b6060600085858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa6000876040516024016118e296959493929190614b9f565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff7000000000000000000000000000000000000000000000000000000001790528351909150606090600003611a335760f9601f83516119779190614c02565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525060e487604051602001611a1d9796959493929190614c1d565b6040516020818303038152906040529050611b37565b815161ffff1015611a70576040517f248b8f8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160f9611a7f602083614c02565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b90000000000000000000000000000000000000000000000000000000000008152508588604051602001611b249796959493929190614d00565b6040516020818303038152906040529150505b805160208083019190912060408051600080825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa158015611b98573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611c10576040517fcd16196600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604051600090611c569084906405ca1ab1e090635ca1ab1e90601b907fff0000000000000000000000000000000000000000000000000000000000000090602001614de3565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529450505050505b949350505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314611cea576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600980547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff8116740100000000000000000000000000000000000000009182900460ff16159091021790556040517ff32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f4190600090a1565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314611db7576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610d2b565b60025473ffffffffffffffffffffffffffffffffffffffff163314611e7b576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836000819003611eb7576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115611ef3576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460055467ffffffffffffffff80831692700100000000000000000000000000000000900416908160005b858110156122e35760008b8b83818110611f3c57611f3c614e3f565b905060800201803603810190611f529190614e6e565b805160608201519192509067ffffffffffffffff16156120cd5785611f7681614ec3565b96505060008183602001518460600151604051602001611fce93929190928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a16600090815260069093529120549091508114612057576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826060015167ffffffffffffffff16836040015167ffffffffffffffff1610156120ad576040517f7f7ab87200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff86166000908152600660205260408120556121cb565b602082015115801590612194575060208201516040517f257b363200000000000000000000000000000000000000000000000000000000815260048101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063257b3632906024016020604051808303816000875af115801561216e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121929190614848565b155b156121cb576040517f73bd668d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8667ffffffffffffffff16826040015167ffffffffffffffff1610806121fe575042826040015167ffffffffffffffff16115b15612235576040517fea82791600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602082810151604080850151815193840189905290830184905260608084019290925260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528c901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c0160405160208183030381529060405280519060200120945081604001519650505080806122db90614eea565b915050611f20565b5060075467ffffffffffffffff680100000000000000009091048116908416111561233a576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169190911790915560058390558590828116908516146124715760006123918386614890565b90506123a767ffffffffffffffff821683614f22565b91506124287f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff166123e061089f565b6123ea9190614f35565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190613ca8565b50600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8716021790555b61256f337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612502573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125269190614848565b6125309190614f35565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016929190613d7c565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156125d757600080fd5b505af11580156125eb573d6000803e3d6000fd5b50506040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8916600482015260248101869052600092507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169150639a908e73906044016020604051808303816000875af115801561268f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126b391906149db565b6009546040517fc7a823e000000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff169063c7a823e09061270e9087908d908d90600401614f95565b60006040518083038186803b15801561272657600080fd5b505afa15801561273a573d6000803e3d6000fd5b505060405167ffffffffffffffff841692507f303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce9150600090a2505050505050505050505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff1633146127d8576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60036127e48282614a46565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610d2b91906142a6565b60085460ff16612850576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600081900361288c576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156128c8576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075467ffffffffffffffff680100000000000000008204811691612903918491700100000000000000000000000000000000900416614fb8565b111561293b576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460055470010000000000000000000000000000000090910467ffffffffffffffff169060005b83811015612bca57600086868381811061298057612980614e3f565b90506020028101906129929190614fcb565b61299b90615009565b9050836129a781614ec3565b82518051602091820120818501516040808701519051949950919450600093612a099386939101928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8916600090815260069093529120549091508114612a92576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8616600090815260066020526040812055612ab7600188614f22565b8403612b265742600760189054906101000a900467ffffffffffffffff168460400151612ae49190615087565b67ffffffffffffffff161115612b26576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020838101516040805192830188905282018490526060808301919091524260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016608083015233901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c016040516020818303038152906040528051906020012094505050508080612bc290614eea565b915050612964565b50612bf87f0000000000000000000000000000000000000000000000000000000000000000846123e061089f565b60058190556007805467ffffffffffffffff848116700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffff000000000000000090921642821617919091179091556040517f9a908e7300000000000000000000000000000000000000000000000000000000815290841660048201526024810182905260009073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639a908e73906044016020604051808303816000875af1158015612ce7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d0b91906149db565b60405190915067ffffffffffffffff8216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a490600090a2505050505050565b60085460ff16612d87576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa158015612df4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e189190614848565b905081811115612e54576040517f2354600f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611388831115612e90576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612ed273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084613d7c565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f639190614848565b600780549192506801000000000000000090910467ffffffffffffffff16906008612f8d83614ec3565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508484604051612fc49291906150a8565b60408051918290038220602083015281018290527fffffffffffffffff0000000000000000000000000000000000000000000000004260c01b166060820152606801604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152918152815160209283012060075468010000000000000000900467ffffffffffffffff16600090815260069093529120553233036130d4576007546040805183815233602082015260609181018290526000918101919091526801000000000000000090910467ffffffffffffffff16907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319060800160405180910390a2613133565b600760089054906101000a900467ffffffffffffffff1667ffffffffffffffff167ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9318233888860405161312a94939291906150b8565b60405180910390a25b5050505050565b60025473ffffffffffffffffffffffffffffffffffffffff16331461318b576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160008190036131c7576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115613203576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460055467ffffffffffffffff80831692700100000000000000000000000000000000900416908160005b8581101561363f57600089898381811061324c5761324c614e3f565b905060200281019061325e91906150f8565b6132679061512c565b8051805160209091012060608201519192509067ffffffffffffffff16156133ea578561329381614ec3565b965050600081836020015184606001516040516020016132eb93929190928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a16600090815260069093529120549091508114613374576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826060015167ffffffffffffffff16836040015167ffffffffffffffff1610156133ca576040517f7f7ab87200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff8616600090815260066020526040812055613527565b6020820151158015906134b1575060208201516040517f257b363200000000000000000000000000000000000000000000000000000000815260048101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063257b3632906024016020604051808303816000875af115801561348b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134af9190614848565b155b156134e8576040517f73bd668d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151516201d4c01015613527576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8667ffffffffffffffff16826040015167ffffffffffffffff16108061355a575042826040015167ffffffffffffffff16115b15613591576040517fea82791600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602082810151604080850151815193840189905290830184905260608084019290925260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528a901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c01604051602081830303815290604052805190602001209450816040015196505050808061363790614eea565b915050613230565b5060075467ffffffffffffffff6801000000000000000090910481169084161115613696576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169190911790915560058390558590828116908516146137855760006136ed8386614890565b905061370367ffffffffffffffff821683614f22565b915061373c7f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff166123e061089f565b50600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8716021790555b613816337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612502573d6000803e3d6000fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561387e57600080fd5b505af1158015613892573d6000803e3d6000fd5b50506040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8916600482015260248101869052600092507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169150639a908e73906044016020604051808303816000875af1158015613936573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061395a91906149db565b60405190915067ffffffffffffffff8216907f303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce90600090a250505050505050505050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f06fdde03000000000000000000000000000000000000000000000000000000001790529051606091600091829173ffffffffffffffffffffffffffffffffffffffff861691613a2091906151a2565b600060405180830381855afa9150503d8060008114613a5b576040519150601f19603f3d011682016040523d82523d6000602084013e613a60565b606091505b509150915081613aa5576040518060400160405280600781526020017f4e4f5f4e414d4500000000000000000000000000000000000000000000000000815250611c8b565b611c8b81613de0565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f95d89b41000000000000000000000000000000000000000000000000000000001790529051606091600091829173ffffffffffffffffffffffffffffffffffffffff861691613b3091906151a2565b600060405180830381855afa9150503d8060008114613b6b576040519150601f19603f3d011682016040523d82523d6000602084013e613b70565b606091505b509150915081613aa5576040518060400160405280600981526020017f4e4f5f53594d424f4c0000000000000000000000000000000000000000000000815250611c8b565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f313ce5670000000000000000000000000000000000000000000000000000000017905290516000918291829173ffffffffffffffffffffffffffffffffffffffff861691613c3691906151a2565b600060405180830381855afa9150503d8060008114613c71576040519150601f19603f3d011682016040523d82523d6000602084013e613c76565b606091505b5091509150818015613c89575080516020145b613c94576012611c8b565b80806020019051810190611c8b91906151b4565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610da29084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613fbb565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052613dda9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401613cfa565b50505050565b60606040825110613e055781806020019051810190613dff91906151d7565b92915050565b8151602003613f7d5760005b602081108015613e585750828181518110613e2e57613e2e614e3f565b01602001517fff000000000000000000000000000000000000000000000000000000000000001615155b15613e6f5780613e6781614eea565b915050613e11565b80600003613eb257505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e4700000000000000000000000000006020820152919050565b60008167ffffffffffffffff811115613ecd57613ecd61442c565b6040519080825280601f01601f191660200182016040528015613ef7576020820181803683370190505b50905060005b82811015613f7557848181518110613f1757613f17614e3f565b602001015160f81c60f81b828281518110613f3457613f34614e3f565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080613f6d81614eea565b915050613efd565b509392505050565b505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e470000000000000000000000000000602082015290565b919050565b600061401d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166140c79092919063ffffffff16565b805190915015610da2578080602001905181019061403b9190614946565b610da2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161108c565b6060611c8b8484600085856000808673ffffffffffffffffffffffffffffffffffffffff1685876040516140fb91906151a2565b60006040518083038185875af1925050503d8060008114614138576040519150601f19603f3d011682016040523d82523d6000602084013e61413d565b606091505b509150915061414e87838387614159565b979650505050505050565b606083156141ef5782516000036141e85773ffffffffffffffffffffffffffffffffffffffff85163b6141e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161108c565b5081611c8b565b611c8b83838151156142045781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108c91906142a6565b60005b8381101561425357818101518382015260200161423b565b50506000910152565b60008151808452614274816020860160208601614238565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006142b9602083018461425c565b9392505050565b67ffffffffffffffff811681146142d657600080fd5b50565b73ffffffffffffffffffffffffffffffffffffffff811681146142d657600080fd5b8035613fb6816142d9565b60008060006060848603121561431b57600080fd5b8335614326816142c0565b925060208401359150604084013561433d816142d9565b809150509250925092565b60006020828403121561435a57600080fd5b81356142b9816142c0565b60008083601f84011261437757600080fd5b50813567ffffffffffffffff81111561438f57600080fd5b6020830191508360208260051b85010111156143aa57600080fd5b9250929050565b6000806000604084860312156143c657600080fd5b833567ffffffffffffffff8111156143dd57600080fd5b6143e986828701614365565b909450925050602084013561433d816142d9565b60006020828403121561440f57600080fd5b81356142b9816142d9565b63ffffffff811681146142d657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff8111828210171561447e5761447e61442c565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156144cb576144cb61442c565b604052919050565b600067ffffffffffffffff8211156144ed576144ed61442c565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261452a57600080fd5b813561453d614538826144d3565b614484565b81815284602083860101111561455257600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c0878903121561458857600080fd5b8635614593816142d9565b955060208701356145a3816142d9565b945060408701356145b38161441a565b935060608701356145c3816142d9565b9250608087013567ffffffffffffffff808211156145e057600080fd5b6145ec8a838b01614519565b935060a089013591508082111561460257600080fd5b5061460f89828a01614519565b9150509295509295509295565b6000806000806080858703121561463257600080fd5b843561463d8161441a565b9350602085013561464d816142d9565b9250604085013561465d8161441a565b9150606085013567ffffffffffffffff81111561467957600080fd5b61468587828801614519565b91505092959194509250565b60008083601f8401126146a357600080fd5b50813567ffffffffffffffff8111156146bb57600080fd5b6020830191508360208285010111156143aa57600080fd5b6000806000806000606086880312156146eb57600080fd5b853567ffffffffffffffff8082111561470357600080fd5b818801915088601f83011261471757600080fd5b81358181111561472657600080fd5b8960208260071b850101111561473b57600080fd5b60208301975080965050614751602089016142fb565b9450604088013591508082111561476757600080fd5b5061477488828901614691565b969995985093965092949392505050565b60006020828403121561479757600080fd5b813567ffffffffffffffff8111156147ae57600080fd5b611c8b84828501614519565b600080602083850312156147cd57600080fd5b823567ffffffffffffffff8111156147e457600080fd5b6147f085828601614365565b90969095509350505050565b60008060006040848603121561481157600080fd5b833567ffffffffffffffff81111561482857600080fd5b61483486828701614691565b909790965060209590950135949350505050565b60006020828403121561485a57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff8281168282160390808211156148b1576148b1614861565b5092915050565b6000826148ee577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600181811c9082168061490757607f821691505b602082108103614940577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60006020828403121561495857600080fd5b815180151581146142b957600080fd5b60608152600061497b606083018661425c565b828103602084015261498d818661425c565b91505060ff83166040830152949350505050565b600080604083850312156149b457600080fd5b82516149bf8161441a565b60208401519092506149d0816142d9565b809150509250929050565b6000602082840312156149ed57600080fd5b81516142b9816142c0565b601f821115610da257600081815260208120601f850160051c81016020861015614a1f5750805b601f850160051c820191505b81811015614a3e57828155600101614a2b565b505050505050565b815167ffffffffffffffff811115614a6057614a6061442c565b614a7481614a6e84546148f3565b846149f8565b602080601f831160018114614ac75760008415614a915750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555614a3e565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015614b1457888601518255948401946001909101908401614af5565b5085821015614b5057878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b606081526000614b73606083018661425c565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b600063ffffffff808916835273ffffffffffffffffffffffffffffffffffffffff8089166020850152818816604085015280871660608501528086166080850152505060c060a0830152614bf660c083018461425c565b98975050505050505050565b61ffff8181168382160190808211156148b1576148b1614861565b60007fff00000000000000000000000000000000000000000000000000000000000000808a60f81b1683527fffff0000000000000000000000000000000000000000000000000000000000008960f01b1660018401528751614c86816003860160208c01614238565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614cc9816017840160208b01614238565b808201915050818660f81b16601782015284519150614cef826018830160208801614238565b016018019998505050505050505050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b16815260007fffff000000000000000000000000000000000000000000000000000000000000808960f01b1660018401528751614d69816003860160208c01614238565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614dac816017840160208b01614238565b808201915050818660f01b16601782015284519150614dd2826019830160208801614238565b016019019998505050505050505050565b60008651614df5818460208b01614238565b9190910194855250602084019290925260f81b7fff000000000000000000000000000000000000000000000000000000000000009081166040840152166041820152604201919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060808284031215614e8057600080fd5b614e8861445b565b82358152602083013560208201526040830135614ea4816142c0565b60408201526060830135614eb7816142c0565b60608201529392505050565b600067ffffffffffffffff808316818103614ee057614ee0614861565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614f1b57614f1b614861565b5060010190565b81810381811115613dff57613dff614861565b8082028115828204841417613dff57613dff614861565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b838152604060208201526000614faf604083018486614f4c565b95945050505050565b80820180821115613dff57613dff614861565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112614fff57600080fd5b9190910192915050565b60006060823603121561501b57600080fd5b6040516060810167ffffffffffffffff828210818311171561503f5761503f61442c565b81604052843591508082111561505457600080fd5b5061506136828601614519565b82525060208301356020820152604083013561507c816142c0565b604082015292915050565b67ffffffffffffffff8181168382160190808211156148b1576148b1614861565b8183823760009101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff841660208201526060604082015260006150ee606083018486614f4c565b9695505050505050565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112614fff57600080fd5b60006080823603121561513e57600080fd5b61514661445b565b823567ffffffffffffffff81111561515d57600080fd5b61516936828601614519565b825250602083013560208201526040830135615184816142c0565b60408201526060830135615197816142c0565b606082015292915050565b60008251614fff818460208701614238565b6000602082840312156151c657600080fd5b815160ff811681146142b957600080fd5b6000602082840312156151e957600080fd5b815167ffffffffffffffff81111561520057600080fd5b8201601f8101841361521157600080fd5b805161521f614538826144d3565b81815285602083850101111561523457600080fd5b614faf82602083016020860161423856fea2646970667358221220c64096e69e927a9fd12b606638c7ffb879aac38d9c988e3a39e8c1b1d6b1092664736f6c63430008140033",
}

// CdkvalidiumABI is the input ABI used to generate the binding from.
// Deprecated: Use CdkvalidiumMetaData.ABI instead.
var CdkvalidiumABI = CdkvalidiumMetaData.ABI

// CdkvalidiumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CdkvalidiumMetaData.Bin instead.
var CdkvalidiumBin = CdkvalidiumMetaData.Bin

// DeployCdkvalidium deploys a new Ethereum contract, binding an instance of Cdkvalidium to it.
func DeployCdkvalidium(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *Cdkvalidium, error) {
	parsed, err := CdkvalidiumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CdkvalidiumBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
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

// SetDataCommittee is a paid mutator transaction binding the contract method 0x98cf4bd0.
//
// Solidity: function setDataCommittee(address newDataCommittee) returns()
func (_Cdkvalidium *CdkvalidiumTransactor) SetDataCommittee(opts *bind.TransactOpts, newDataCommittee common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.contract.Transact(opts, "setDataCommittee", newDataCommittee)
}

// SetDataCommittee is a paid mutator transaction binding the contract method 0x98cf4bd0.
//
// Solidity: function setDataCommittee(address newDataCommittee) returns()
func (_Cdkvalidium *CdkvalidiumSession) SetDataCommittee(newDataCommittee common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SetDataCommittee(&_Cdkvalidium.TransactOpts, newDataCommittee)
}

// SetDataCommittee is a paid mutator transaction binding the contract method 0x98cf4bd0.
//
// Solidity: function setDataCommittee(address newDataCommittee) returns()
func (_Cdkvalidium *CdkvalidiumTransactorSession) SetDataCommittee(newDataCommittee common.Address) (*types.Transaction, error) {
	return _Cdkvalidium.Contract.SetDataCommittee(&_Cdkvalidium.TransactOpts, newDataCommittee)
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

// CdkvalidiumSetDataCommitteeIterator is returned from FilterSetDataCommittee and is used to iterate over the raw logs and unpacked data for SetDataCommittee events raised by the Cdkvalidium contract.
type CdkvalidiumSetDataCommitteeIterator struct {
	Event *CdkvalidiumSetDataCommittee // Event containing the contract specifics and raw log

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
func (it *CdkvalidiumSetDataCommitteeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkvalidiumSetDataCommittee)
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
		it.Event = new(CdkvalidiumSetDataCommittee)
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
func (it *CdkvalidiumSetDataCommitteeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkvalidiumSetDataCommitteeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkvalidiumSetDataCommittee represents a SetDataCommittee event raised by the Cdkvalidium contract.
type CdkvalidiumSetDataCommittee struct {
	NewDataCommittee common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetDataCommittee is a free log retrieval operation binding the contract event 0x76d12ecc41da43ceb240d60fdd4d9fc0baae9793060e7b592abde7ece5ee9651.
//
// Solidity: event SetDataCommittee(address newDataCommittee)
func (_Cdkvalidium *CdkvalidiumFilterer) FilterSetDataCommittee(opts *bind.FilterOpts) (*CdkvalidiumSetDataCommitteeIterator, error) {

	logs, sub, err := _Cdkvalidium.contract.FilterLogs(opts, "SetDataCommittee")
	if err != nil {
		return nil, err
	}
	return &CdkvalidiumSetDataCommitteeIterator{contract: _Cdkvalidium.contract, event: "SetDataCommittee", logs: logs, sub: sub}, nil
}

// WatchSetDataCommittee is a free log subscription operation binding the contract event 0x76d12ecc41da43ceb240d60fdd4d9fc0baae9793060e7b592abde7ece5ee9651.
//
// Solidity: event SetDataCommittee(address newDataCommittee)
func (_Cdkvalidium *CdkvalidiumFilterer) WatchSetDataCommittee(opts *bind.WatchOpts, sink chan<- *CdkvalidiumSetDataCommittee) (event.Subscription, error) {

	logs, sub, err := _Cdkvalidium.contract.WatchLogs(opts, "SetDataCommittee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkvalidiumSetDataCommittee)
				if err := _Cdkvalidium.contract.UnpackLog(event, "SetDataCommittee", log); err != nil {
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

// ParseSetDataCommittee is a log parse operation binding the contract event 0x76d12ecc41da43ceb240d60fdd4d9fc0baae9793060e7b592abde7ece5ee9651.
//
// Solidity: event SetDataCommittee(address newDataCommittee)
func (_Cdkvalidium *CdkvalidiumFilterer) ParseSetDataCommittee(log types.Log) (*CdkvalidiumSetDataCommittee, error) {
	event := new(CdkvalidiumSetDataCommittee)
	if err := _Cdkvalidium.contract.UnpackLog(event, "SetDataCommittee", log); err != nil {
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
