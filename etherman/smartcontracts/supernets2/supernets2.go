// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package supernets2

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

// Supernets2BatchData is an auto generated low-level Go binding around an user-defined struct.
type Supernets2BatchData struct {
	TransactionsHash   [32]byte
	GlobalExitRoot     [32]byte
	Timestamp          uint64
	MinForcedTimestamp uint64
}

// Supernets2ForcedBatchData is an auto generated low-level Go binding around an user-defined struct.
type Supernets2ForcedBatchData struct {
	Transactions       []byte
	GlobalExitRoot     [32]byte
	MinForcedTimestamp uint64
}

// Supernets2InitializePackedParameters is an auto generated low-level Go binding around an user-defined struct.
type Supernets2InitializePackedParameters struct {
	Admin                    common.Address
	TrustedSequencer         common.Address
	PendingStateTimeout      uint64
	TrustedAggregator        common.Address
	TrustedAggregatorTimeout uint64
}

// Supernets2MetaData contains all meta data concerning the Supernets2 contract.
var Supernets2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_matic\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"_rollupVerifier\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractISupernets2DataCommittee\",\"name\":\"_dataCommitteeAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_forkID\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ActivateForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"ConsolidatePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"OverridePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"provedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProveNonDeterministicPendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"SetMultiplierBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"SetPendingStateTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"SetTrustedAggregatorTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"SetVerifyBatchTimeTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"UpdateSupernets2Version\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequencedBatchNum\",\"type\":\"uint64\"}],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"}],\"name\":\"checkStateRootInsidePrime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"consolidatePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataCommitteeAddress\",\"outputs\":[{\"internalType\":\"contractISupernets2DataCommittee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maticAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forkID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"trustedAggregatorTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structSupernets2.InitializePackedParameters\",\"name\":\"initializePackedParameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"genesisRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isForcedBatchDisallowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"isPendingStateConsolidable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingState\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingStateConsolidated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matic\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierBatchFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"overridePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStateTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingStateTransitions\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"proveNonDeterministicPendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifierRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structSupernets2.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signaturesAndAddrs\",\"type\":\"bytes\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structSupernets2.ForcedBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"sequencedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"setMultiplierBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"setPendingStateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"setTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"setTrustedAggregatorTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"setVerifyBatchTimeTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregatorTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyBatchTimeTarget\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x61016060405234801562000011575f80fd5b5060405162006065380380620060658339810160408190526200003491620000a9565b6001600160a01b0396871660c05294861660805292851660a05290841660e052909216610100526001600160401b0391821661012052166101405262000147565b6001600160a01b03811681146200008a575f80fd5b50565b80516001600160401b0381168114620000a4575f80fd5b919050565b5f805f805f805f60e0888a031215620000c0575f80fd5b8751620000cd8162000075565b6020890151909750620000e08162000075565b6040890151909650620000f38162000075565b6060890151909550620001068162000075565b6080890151909450620001198162000075565b92506200012960a089016200008d565b91506200013960c089016200008d565b905092959891949750929550565b60805160a05160c05160e051610100516101205161014051615e476200021e5f395f81816106b401528181610e0501526131bc01525f81816108200152610ddb01525f81816105c1015261196f01525f81816107e601528181611bb90152818161384d0152614c9b01525f818161098b01528181610f77015281816111420152818161174e015281816121cd01528181613a2e015261490401525f8181610a38015281816140e0015261452e01525f81816108db01528181611b87015281816126b201528181613a0301526141cc0152615e475ff3fe608060405234801561000f575f80fd5b50600436106103b4575f3560e01c8063837a4738116101f5578063c754c7ed11610114578063e7a7ed02116100a9578063f14916d611610079578063f14916d614610a9a578063f2fde38b14610aad578063f851a44014610ac0578063f8b823e414610ae0575f80fd5b8063e7a7ed0214610a03578063e8bf92ed14610a33578063eaeb077b14610a5a578063ed6b010414610a6d575f80fd5b8063d2e129f9116100e4578063d2e129f9146109ad578063d8d1091b146109c0578063d939b315146109d3578063dbc16976146109fb575f80fd5b8063c754c7ed14610918578063c89e42df14610944578063cfa8ed4714610957578063d02103ca14610986575f80fd5b8063a3c573eb1161018a578063b4d63f581161015a578063b4d63f5814610870578063b6b0b097146108d6578063ba58ae39146108fd578063c0ed84e014610910575f80fd5b8063a3c573eb146107e1578063ada8f91914610808578063adc879e91461081b578063afd23cbe14610842575f80fd5b806399f5634e116101c557806399f5634e146107a05780639aa972a3146107a85780639c9f3dfe146107bb578063a066215c146107ce575f80fd5b8063837a4738146106d6578063841b24d71461074a5780638c3d73011461077a5780638da5cb5b14610782575f80fd5b8063458c0477116102e15780636046916911610276578063715018a611610246578063715018a6146106805780637215541a146106885780637fcb36531461069b578063831c7ead146106af575f80fd5b80636046916914610633578063621dd4111461063b5780636b8616ce1461064e5780636ff512cc1461066d575f80fd5b80634e487706116102b15780634e487706146105e35780635392c5e0146105f6578063542028d5146106235780635ec919581461062b575f80fd5b8063458c0477146105755780634a1a89a7146105895780634a910e6a146105a95780634df61d24146105bc575f80fd5b80632987898311610357578063394218e911610327578063394218e914610507578063423fa8561461051a578063438a53991461053a578063456052671461054d575f80fd5b806329878983146104a25780632b0006fa146104ce5780632c1f816a146104e1578063383b3be8146104f4575f80fd5b80631816b7e5116103925780631816b7e51461042157806319d8ac6114610436578063220d78991461044a578063267822471461045d575f80fd5b80630a0d9fbe146103b8578063107bf28c146103ef57806315064c9614610404575b5f80fd5b606f546103d190610100900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b6103f7610ae9565b6040516103e691906152b5565b606f546104119060ff1681565b60405190151581526020016103e6565b61043461042f3660046152ce565b610b75565b005b6073546103d19067ffffffffffffffff1681565b6103f7610458366004615306565b610c8d565b607b5461047d9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016103e6565b60745461047d9068010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b6104346104dc366004615366565b610e63565b6104346104ef3660046153ca565b61102d565b61041161050236600461543f565b611235565b61043461051536600461543f565b61128a565b6073546103d19068010000000000000000900467ffffffffffffffff1681565b6104346105483660046154c0565b61140e565b6073546103d190700100000000000000000000000000000000900467ffffffffffffffff1681565b6079546103d19067ffffffffffffffff1681565b6079546103d19068010000000000000000900467ffffffffffffffff1681565b6104346105b736600461543f565b611c74565b61047d7f000000000000000000000000000000000000000000000000000000000000000081565b6104346105f136600461543f565b611d27565b61061561060436600461543f565b60756020525f908152604090205481565b6040519081526020016103e6565b6103f7611eab565b610434611eb8565b610615611fb7565b610434610649366004615366565b611fcc565b61061561065c36600461543f565b60716020525f908152604090205481565b61043461067b366004615569565b61234a565b61043461241f565b61043461069636600461543f565b612432565b6074546103d19067ffffffffffffffff1681565b6103d17f000000000000000000000000000000000000000000000000000000000000000081565b61071e6106e4366004615582565b60786020525f908152604090208054600182015460029092015467ffffffffffffffff808316936801000000000000000090930416919084565b6040805167ffffffffffffffff95861681529490931660208501529183015260608201526080016103e6565b6079546103d1907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b61043461259f565b60335473ffffffffffffffffffffffffffffffffffffffff1661047d565b61061561266b565b6104346107b63660046153ca565b6127be565b6104346107c936600461543f565b61286e565b6104346107dc36600461543f565b6129ea565b61047d7f000000000000000000000000000000000000000000000000000000000000000081565b610434610816366004615569565b612af0565b6103d17f000000000000000000000000000000000000000000000000000000000000000081565b606f5461085d906901000000000000000000900461ffff1681565b60405161ffff90911681526020016103e6565b6108b061087e36600461543f565b60726020525f90815260409020805460019091015467ffffffffffffffff808216916801000000000000000090041683565b6040805193845267ffffffffffffffff92831660208501529116908201526060016103e6565b61047d7f000000000000000000000000000000000000000000000000000000000000000081565b61041161090b366004615582565b612bb4565b6103d1612c3c565b607b546103d19074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b610434610952366004615675565b612c8f565b606f5461047d906b010000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b61047d7f000000000000000000000000000000000000000000000000000000000000000081565b6104346109bb3660046156a7565b612d1c565b6104346109ce366004615752565b61325f565b6079546103d190700100000000000000000000000000000000900467ffffffffffffffff1681565b6104346137fa565b6073546103d1907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b61047d7f000000000000000000000000000000000000000000000000000000000000000081565b610434610a683660046157c1565b6138ce565b607b54610411907c0100000000000000000000000000000000000000000000000000000000900460ff1681565b610434610aa8366004615569565b613cbe565b610434610abb366004615569565b613d90565b607a5461047d9073ffffffffffffffffffffffffffffffffffffffff1681565b61061560705481565b60778054610af690615809565b80601f0160208091040260200160405190810160405280929190818152602001828054610b2290615809565b8015610b6d5780601f10610b4457610100808354040283529160200191610b6d565b820191905f5260205f20905b815481529060010190602001808311610b5057829003601f168201915b505050505081565b607a5473ffffffffffffffffffffffffffffffffffffffff163314610bc6576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88161ffff161080610bdf57506103ff8161ffff16115b15610c16576040517f4c2533c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffff0000ffffffffffffffffff16690100000000000000000061ffff8416908102919091179091556040519081527f7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5906020015b60405180910390a150565b67ffffffffffffffff8086165f818152607260205260408082205493881682529020546060929115801590610cc0575081155b15610cf7576040517f6818c29e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610d2e576040517f66385b5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d3784612bb4565b610d6d576040517f176b913c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080517fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003360601b166020820152603481019690965260548601929092527fffffffffffffffff00000000000000000000000000000000000000000000000060c098891b811660748701527f0000000000000000000000000000000000000000000000000000000000000000891b8116607c8701527f0000000000000000000000000000000000000000000000000000000000000000891b81166084870152608c86019490945260ac85015260cc840194909452509290931b90911660ec830152805180830360d401815260f4909201905290565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314610ec0576040517fbbcbbc0500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ece868686868686613e44565b607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169182179092555f9081526075602052604090208390556079541615610f4857607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d906024015f604051808303815f87803b158015610fcd575f80fd5b505af1158015610fdf573d5f803e3d5ffd5b505060405184815233925067ffffffffffffffff871691507fcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe906020015b60405180910390a3505050505050565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff16331461108a576040517fbbcbbc0500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611099878787878787876141ff565b607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169182179092555f908152607560205260409020839055607954161561111357607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d906024015f604051808303815f87803b158015611198575f80fd5b505af11580156111aa573d5f803e3d5ffd5b50506079805477ffffffffffffffffffffffffffffffffffffffffffffffff167a093a800000000000000000000000000000000000000000000000001790555050604051828152339067ffffffffffffffff8616907fcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf729060200160405180910390a350505050505050565b60795467ffffffffffffffff8281165f90815260786020526040812054909242926112789270010000000000000000000000000000000090920481169116615887565b67ffffffffffffffff16111592915050565b607a5473ffffffffffffffffffffffffffffffffffffffff1633146112db576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115611322576040517f1d06e87900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff166113915760795467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690821610611391576040517f401636df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6079805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527f1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a190602001610c82565b606f5460ff161561144b576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f546b010000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1633146114ab576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835f8190036114e6576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115611522576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff68010000000000000000820481165f81815260726020526040812054838516949293700100000000000000000000000000000000909304909216919082905b86811015611931575f8c8c83818110611588576115886158af565b90506080020180360381019061159e91906158dc565b606081015190915067ffffffffffffffff161561170b57846115bf8161594b565b9550505f815f01518260200151836060015160405160200161161993929190928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89165f908152607190935291205490915081146116a1576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8087165f908152607160205260408082209190915560608401519084015190821691161015611705576040517f7f7ab87200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50611806565b6020810151158015906117cf575060208101516040517f257b363200000000000000000000000000000000000000000000000000000000815260048101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063257b3632906024016020604051808303815f875af11580156117a9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117cd9190615971565b155b15611806576040517f73bd668d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8667ffffffffffffffff16816040015167ffffffffffffffff161080611839575042816040015167ffffffffffffffff16115b15611870576040517fea82791600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83815f0151826020015183604001518e6040516020016118fe9594939291909485526020850193909352604084019190915260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060808401919091521b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166068820152607c0190565b6040516020818303038152906040528051906020012093508060400151965050808061192990615988565b91505061156d565b506040517fc7a823e000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063c7a823e0906119a89085908c908c90600401615a06565b5f6040518083038186803b1580156119be575f80fd5b505afa1580156119d0573d5f803e3d5ffd5b5050505085846119e09190615887565b60735490945067ffffffffffffffff780100000000000000000000000000000000000000000000000090910481169084161115611a49576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611a548285615a28565b611a689067ffffffffffffffff1688615a49565b604080516060810182528581524267ffffffffffffffff908116602080840191825260738054680100000000000000009081900485168688019081528d86165f8181526072909552979093209551865592516001909501805492519585167fffffffffffffffffffffffffffffffff000000000000000000000000000000009384161795851684029590951790945583548c8416911617930292909217905590915082811690851614611b5d57607380547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8716021790555b611baf333083607054611b709190615a5c565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692919061462e565b611bb7614710565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b158015611c1c575f80fd5b505af1158015611c2e573d5f803e3d5ffd5b505060405167ffffffffffffffff881692507f303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce91505f90a2505050505050505050505050565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314611d1b57606f5460ff1615611cdc576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ce581611235565b611d1b576040517f0ce9e4a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d24816147bb565b50565b607a5473ffffffffffffffffffffffffffffffffffffffff163314611d78576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115611dbf576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff16611e2a57607b5467ffffffffffffffff74010000000000000000000000000000000000000000909104811690821610611e2a576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b90602001610c82565b60768054610af690615809565b607a5473ffffffffffffffffffffffffffffffffffffffff163314611f09576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b547c0100000000000000000000000000000000000000000000000000000000900460ff16611f65576040517ff6ba91a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690556040517f854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f905f90a1565b5f6070546064611fc79190615a5c565b905090565b606f5460ff1615612009576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795467ffffffffffffffff8581165f90815260726020526040902060010154429261205592780100000000000000000000000000000000000000000000000090910481169116615887565b67ffffffffffffffff161115612097576040517f8a0704d300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e86120a48686615a28565b67ffffffffffffffff1611156120e6576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6120f4868686868686613e44565b6120fd846149c8565b607954700100000000000000000000000000000000900467ffffffffffffffff165f0361223e57607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169182179092555f908152607560205260409020839055607954161561219e57607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d906024015f604051808303815f87803b158015612223575f80fd5b505af1158015612235573d5f803e3d5ffd5b5050505061230c565b612246614710565b6079805467ffffffffffffffff16905f61225f8361594b565b825467ffffffffffffffff9182166101009390930a92830292820219169190911790915560408051608081018252428316815287831660208083019182528284018981526060840189815260795487165f908152607890935294909120925183549251861668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169516949094171781559151600183015551600290910155505b604051828152339067ffffffffffffffff8616907f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f59669060200161101d565b607a5473ffffffffffffffffffffffffffffffffffffffff16331461239b576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fff0000000000000000000000000000000000000000ffffffffffffffffffffff166b01000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610c82565b612427614ba2565b6124305f614c23565b565b60335473ffffffffffffffffffffffffffffffffffffffff163314612597575f61245a612c3c565b90508067ffffffffffffffff168267ffffffffffffffff16116124a9576040517f812a372d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff68010000000000000000909104811690831611806124ee575067ffffffffffffffff8083165f9081526072602052604090206001015416155b15612525576040517f98c5c01400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8083165f9081526072602052604090206001015442916125539162093a809116615887565b67ffffffffffffffff161115612595576040517fd257555a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b611d24614c99565b607b5473ffffffffffffffffffffffffffffffffffffffff1633146125f0576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b54607a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f90819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa1580156126f7573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061271b9190615971565b90505f612726612c3c565b60735467ffffffffffffffff68010000000000000000820481169161277e9170010000000000000000000000000000000082048116917801000000000000000000000000000000000000000000000000900416615a28565b6127889190615887565b6127929190615a28565b67ffffffffffffffff169050805f036127ad575f9250505090565b6127b78183615aa0565b9250505090565b606f5460ff16156127fb576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61280a878787878787876141ff565b67ffffffffffffffff84165f908152607560209081526040918290205482519081529081018490527f1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010910160405180910390a1612865614c99565b50505050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff1633146128bf576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115612906576040517fcc96507000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff1661296d5760795467ffffffffffffffff70010000000000000000000000000000000090910481169082161061296d576040517f48a05a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607980547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c7590602001610c82565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612a3b576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620151808167ffffffffffffffff161115612a82576040517fe067dfe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff1661010067ffffffffffffffff8416908102919091179091556040519081527f1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c2890602001610c82565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612b41576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610c82565b5f67ffffffff0000000167ffffffffffffffff8316108015612beb575067ffffffff00000001604083901c67ffffffffffffffff16105b8015612c0c575067ffffffff00000001608083901c67ffffffffffffffff16105b8015612c23575067ffffffff0000000160c083901c105b15612c3057506001919050565b505f919050565b919050565b6079545f9067ffffffffffffffff1615612c7e575060795467ffffffffffffffff9081165f908152607860205260409020546801000000000000000090041690565b5060745467ffffffffffffffff1690565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612ce0576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6076612cec8282615b00565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610c8291906152b5565b5f54610100900460ff1615808015612d3a57505f54600160ff909116105b80612d535750303b158015612d5357505f5460ff166001145b612de4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015612e40575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b612e4d6020880188615569565b607a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055612ea26040880160208901615569565b606f805473ffffffffffffffffffffffffffffffffffffffff929092166b010000000000000000000000027fff0000000000000000000000000000000000000000ffffffffffffffffffffff909216919091179055612f076080880160608901615569565b6074805473ffffffffffffffffffffffffffffffffffffffff9290921668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff9092169190911790555f805260756020527ff9e3fbf150b7a0077118526f473c53cb4734f166167e2c6213e3567dd390b4ad8690556076612f918682615b00565b506077612f9e8582615b00565b5062093a80612fb36060890160408a0161543f565b67ffffffffffffffff161115612ff5576040517fcc96507000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613005606088016040890161543f565b6079805467ffffffffffffffff92909216700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff90921691909117905562093a8061306760a0890160808a0161543f565b67ffffffffffffffff1611156130a9576040517f1d06e87900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6130b960a088016080890161543f565b6079805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff939093169290920291909117905567016345785d8a0000607055606f80547fffffffffffffffffffffffffffffffffffffffffff00000000000000000000ff166a03ea000000000000070800179055607b80547fffffff000000000000000000ffffffffffffffffffffffffffffffffffffffff167c0100000000000697800000000000000000000000000000000000000000179055613198614d1c565b7f13e60734641bd97ab6e444f419277d9983e2c3eab9bb631a578ffa07756b63515f7f000000000000000000000000000000000000000000000000000000000000000085856040516131ed9493929190615c18565b60405180910390a18015612865575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050505050565b607b547c0100000000000000000000000000000000000000000000000000000000900460ff16156132bc576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff16156132f9576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f819003613334576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115613370576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff780100000000000000000000000000000000000000000000000082048116916133bb918491700100000000000000000000000000000000900416615c4f565b11156133f3576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff68010000000000000000820481165f8181526072602052604081205491937001000000000000000000000000000000009004909216915b84811015613698575f878783818110613451576134516158af565b90506020028101906134639190615c62565b61346c90615c9e565b9050836134788161594b565b825180516020918201208185015160408087015190519499509194505f936134d99386939101928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89165f90815260719093529120549091508114613561576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff86165f90815260716020526040812055613585600189615a49565b84036135f45742607b60149054906101000a900467ffffffffffffffff1684604001516135b29190615887565b67ffffffffffffffff1611156135f4576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020838101516040805192830188905282018490526060808301919091524260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016608083015233901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c01604051602081830303815290604052805190602001209450505050808061369090615988565b915050613436565b506136a38484615887565b6073805467ffffffffffffffff4281167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009092168217808455604080516060810182528781526020808201958652680100000000000000009384900485168284019081528589165f818152607290935284832093518455965160019390930180549151871686027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090921693871693909317179091558554938916700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff938602939093167fffffffffffffffff00000000000000000000000000000000ffffffffffffffff90941693909317919091179093559151929550917f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a49190a2505050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff16331461384b576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dbc169766040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156138b0575f80fd5b505af11580156138c2573d5f803e3d5ffd5b50505050612430614dbb565b607b547c0100000000000000000000000000000000000000000000000000000000900460ff161561392b576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff1615613968576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613971611fb7565b9050818111156139ad576040517f4732fdb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113888311156139e9576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613a2b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633308461462e565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613a95573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613ab99190615971565b60738054919250780100000000000000000000000000000000000000000000000090910467ffffffffffffffff16906018613af38361594b565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508484604051613b2a929190615d2a565b60408051918290038220602083015281018290527fffffffffffffffff0000000000000000000000000000000000000000000000004260c01b166060820152606801604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815291815281516020928301206073547801000000000000000000000000000000000000000000000000900467ffffffffffffffff165f9081526071909352912055323303613c58576073546040805183815233602082015260609181018290525f91810191909152780100000000000000000000000000000000000000000000000090910467ffffffffffffffff16907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319060800160405180910390a2613cb7565b607360189054906101000a900467ffffffffffffffff1667ffffffffffffffff167ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc93182338888604051613cae9493929190615d39565b60405180910390a25b5050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314613d0f576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607480547fffffffff0000000000000000000000000000000000000000ffffffffffffffff166801000000000000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca90602001610c82565b613d98614ba2565b73ffffffffffffffffffffffffffffffffffffffff8116613e3b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401612ddb565b611d2481614c23565b5f80613e4e612c3c565b905067ffffffffffffffff881615613f1d5760795467ffffffffffffffff9081169089161115613eaa576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8089165f908152607860205260409020600281015481549094509091898116680100000000000000009092041614613f17576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50613fbd565b67ffffffffffffffff87165f90815260756020526040902054915081613f6f576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168767ffffffffffffffff161115613fbd576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168667ffffffffffffffff161161400a576040517fb9b18f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6140188888888689610c8d565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160028360405161404c9190615d6e565b602060405180830381855afa158015614067573d5f803e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061408a9190615971565b6140949190615d7f565b6040805160208101825282815290517f9121da8a00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691639121da8a9161411691899190600401615d92565b602060405180830381865afa158015614131573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906141559190615dcc565b61418b576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6141f333614199858b615a28565b67ffffffffffffffff166141ab61266b565b6141b59190615a5c565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190614e49565b50505050505050505050565b5f67ffffffffffffffff8816156142cb5760795467ffffffffffffffff908116908916111561425a576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff8088165f9081526078602052604090206002810154815490928881166801000000000000000090920416146142c5576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50614366565b5067ffffffffffffffff85165f908152607560205260409020548061431c576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60745467ffffffffffffffff9081169087161115614366576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795467ffffffffffffffff908116908816118061439857508767ffffffffffffffff168767ffffffffffffffff1611155b806143bf575060795467ffffffffffffffff68010000000000000000909104811690881611155b156143f6576040517fbfa7079f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8781165f90815260786020526040902054680100000000000000009004811690861614614458576040517f32a2a77f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6144668787878588610c8d565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160028360405161449a9190615d6e565b602060405180830381855afa1580156144b5573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906144d89190615971565b6144e29190615d7f565b6040805160208101825282815290517f9121da8a00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691639121da8a9161456491889190600401615d92565b602060405180830381865afa15801561457f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906145a39190615dcc565b6145d9576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff89165f908152607860205260409020600201548590036141f3576040517fa47276bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261470a9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152614ea4565b50505050565b60795467ffffffffffffffff680100000000000000008204811691161115612430576079545f906147589068010000000000000000900467ffffffffffffffff166001615887565b905061476381611235565b15611d24576079545f9060029061478590849067ffffffffffffffff16615a28565b61478f9190615deb565b6147999083615887565b90506147a481611235565b156147b6576147b2816147bb565b5050565b6147b2825b60795467ffffffffffffffff6801000000000000000090910481169082161115806147f5575060795467ffffffffffffffff908116908216115b1561482c576040517fd086b70b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8181165f81815260786020908152604080832080546074805468010000000000000000928390049098167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000090981688179055600282015487865260759094529382902092909255607980547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff169390940292909217909255600182015490517f33d6247d00000000000000000000000000000000000000000000000000000000815260048101919091529091907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d906024015f604051808303815f87803b15801561495a575f80fd5b505af115801561496c573d5f803e3d5ffd5b505050508267ffffffffffffffff168167ffffffffffffffff167f328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e84600201546040516149bb91815260200190565b60405180910390a3505050565b5f6149d1612c3c565b9050815f806149e08484615a28565b606f5467ffffffffffffffff91821692505f91614a039161010090041642615a49565b90505b8467ffffffffffffffff168467ffffffffffffffff1614614a8d5767ffffffffffffffff8085165f9081526072602052604090206001810154909116821015614a6b57600181015468010000000000000000900467ffffffffffffffff169450614a87565b614a758686615a28565b67ffffffffffffffff16935050614a8d565b50614a06565b5f614a988484615a49565b905083811015614aef57808403600c8111614ab35780614ab6565b600c5b9050806103e80a81606f60099054906101000a900461ffff1661ffff160a6070540281614ae557614ae5615a73565b0460705550614b5e565b838103600c8111614b005780614b03565b600c5b90505f816103e80a82606f60099054906101000a900461ffff1661ffff160a670de0b6b3a76400000281614b3957614b39615a73565b04905080607054670de0b6b3a76400000281614b5757614b57615a73565b0460705550505b683635c9adc5dea000006070541115614b8357683635c9adc5dea00000607055612865565b633b9aca00607054101561286557633b9aca0060705550505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314612430576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401612ddb565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632072f6c56040518163ffffffff1660e01b81526004015f604051808303815f87803b158015614cfe575f80fd5b505af1158015614d10573d5f803e3d5ffd5b50505050612430614faf565b5f54610100900460ff16614db2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401612ddb565b61243033614c23565b606f5460ff16614df7576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052614e9f9084907fa9059cbb0000000000000000000000000000000000000000000000000000000090606401614688565b505050565b5f614f05826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166150419092919063ffffffff16565b805190915015614e9f5780806020019051810190614f239190615dcc565b614e9f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401612ddb565b606f5460ff1615614fec576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b606061504f84845f85615057565b949350505050565b6060824710156150e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401612ddb565b5f808673ffffffffffffffffffffffffffffffffffffffff1685876040516151119190615d6e565b5f6040518083038185875af1925050503d805f811461514b576040519150601f19603f3d011682016040523d82523d5f602084013e615150565b606091505b50915091506151618783838761516c565b979650505050505050565b606083156152015782515f036151fa5773ffffffffffffffffffffffffffffffffffffffff85163b6151fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401612ddb565b508161504f565b61504f83838151156152165781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ddb91906152b5565b5f5b8381101561526457818101518382015260200161524c565b50505f910152565b5f815180845261528381602086016020860161524a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f6152c7602083018461526c565b9392505050565b5f602082840312156152de575f80fd5b813561ffff811681146152c7575f80fd5b803567ffffffffffffffff81168114612c37575f80fd5b5f805f805f60a0868803121561531a575f80fd5b615323866152ef565b9450615331602087016152ef565b94979496505050506040830135926060810135926080909101359150565b806103008101831015615360575f80fd5b92915050565b5f805f805f806103a0878903121561537c575f80fd5b615385876152ef565b9550615393602088016152ef565b94506153a1604088016152ef565b935060608701359250608087013591506153be8860a0890161534f565b90509295509295509295565b5f805f805f805f6103c0888a0312156153e1575f80fd5b6153ea886152ef565b96506153f8602089016152ef565b9550615406604089016152ef565b9450615414606089016152ef565b93506080880135925060a088013591506154318960c08a0161534f565b905092959891949750929550565b5f6020828403121561544f575f80fd5b6152c7826152ef565b803573ffffffffffffffffffffffffffffffffffffffff81168114612c37575f80fd5b5f8083601f84011261548b575f80fd5b50813567ffffffffffffffff8111156154a2575f80fd5b6020830191508360208285010111156154b9575f80fd5b9250929050565b5f805f805f606086880312156154d4575f80fd5b853567ffffffffffffffff808211156154eb575f80fd5b818801915088601f8301126154fe575f80fd5b81358181111561550c575f80fd5b8960208260071b8501011115615520575f80fd5b6020830197508096505061553660208901615458565b9450604088013591508082111561554b575f80fd5b506155588882890161547b565b969995985093965092949392505050565b5f60208284031215615579575f80fd5b6152c782615458565b5f60208284031215615592575f80fd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f67ffffffffffffffff808411156155e0576155e0615599565b604051601f85017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561562657615626615599565b8160405280935085815286868601111561563e575f80fd5b858560208301375f602087830101525050509392505050565b5f82601f830112615666575f80fd5b6152c7838335602085016155c6565b5f60208284031215615685575f80fd5b813567ffffffffffffffff81111561569b575f80fd5b61504f84828501615657565b5f805f805f808688036101208112156156be575f80fd5b60a08112156156cb575f80fd5b5086955060a0870135945060c087013567ffffffffffffffff808211156156f0575f80fd5b6156fc8a838b01615657565b955060e0890135915080821115615711575f80fd5b61571d8a838b01615657565b9450610100890135915080821115615733575f80fd5b5061574089828a0161547b565b979a9699509497509295939492505050565b5f8060208385031215615763575f80fd5b823567ffffffffffffffff8082111561577a575f80fd5b818501915085601f83011261578d575f80fd5b81358181111561579b575f80fd5b8660208260051b85010111156157af575f80fd5b60209290920196919550909350505050565b5f805f604084860312156157d3575f80fd5b833567ffffffffffffffff8111156157e9575f80fd5b6157f58682870161547b565b909790965060209590950135949350505050565b600181811c9082168061581d57607f821691505b602082108103615854577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff8181168382160190808211156158a8576158a861585a565b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f608082840312156158ec575f80fd5b6040516080810181811067ffffffffffffffff8211171561590f5761590f615599565b8060405250823581526020830135602082015261592e604084016152ef565b604082015261593f606084016152ef565b60608201529392505050565b5f67ffffffffffffffff8083168181036159675761596761585a565b6001019392505050565b5f60208284031215615981575f80fd5b5051919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036159b8576159b861585a565b5060010190565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b838152604060208201525f615a1f6040830184866159bf565b95945050505050565b67ffffffffffffffff8281168282160390808211156158a8576158a861585a565b818103818111156153605761536061585a565b80820281158282048414176153605761536061585a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f82615aae57615aae615a73565b500490565b601f821115614e9f575f81815260208120601f850160051c81016020861015615ad95750805b601f850160051c820191505b81811015615af857828155600101615ae5565b505050505050565b815167ffffffffffffffff811115615b1a57615b1a615599565b615b2e81615b288454615809565b84615ab3565b602080601f831160018114615b80575f8415615b4a5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555615af8565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015615bcc57888601518255948401946001909101908401615bad565b5085821015615c0857878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b5f67ffffffffffffffff808716835280861660208401525060606040830152615c456060830184866159bf565b9695505050505050565b808201808211156153605761536061585a565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112615c94575f80fd5b9190910192915050565b5f60608236031215615cae575f80fd5b6040516060810167ffffffffffffffff8282108183111715615cd257615cd2615599565b816040528435915080821115615ce6575f80fd5b50830136601f820112615cf7575f80fd5b615d06368235602084016155c6565b82525060208301356020820152615d1f604084016152ef565b604082015292915050565b818382375f9101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201525f615c456060830184866159bf565b5f8251615c9481846020870161524a565b5f82615d8d57615d8d615a73565b500690565b6103208101610300808584378201835f5b6001811015615dc2578151835260209283019290910190600101615da3565b5050509392505050565b5f60208284031215615ddc575f80fd5b815180151581146152c7575f80fd5b5f67ffffffffffffffff80841680615e0557615e05615a73565b9216919091049291505056fea26469706673582212207087e7267ea85aa56f19949402dd5421ac541d59638e1709fc2d7480b6aa80ac64736f6c63430008140033",
}

// Supernets2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Supernets2MetaData.ABI instead.
var Supernets2ABI = Supernets2MetaData.ABI

// Supernets2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Supernets2MetaData.Bin instead.
var Supernets2Bin = Supernets2MetaData.Bin

// DeploySupernets2 deploys a new Ethereum contract, binding an instance of Supernets2 to it.
func DeploySupernets2(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _matic common.Address, _rollupVerifier common.Address, _bridgeAddress common.Address, _dataCommitteeAddress common.Address, _chainID uint64, _forkID uint64) (common.Address, *types.Transaction, *Supernets2, error) {
	parsed, err := Supernets2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Supernets2Bin), backend, _globalExitRootManager, _matic, _rollupVerifier, _bridgeAddress, _dataCommitteeAddress, _chainID, _forkID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Supernets2{Supernets2Caller: Supernets2Caller{contract: contract}, Supernets2Transactor: Supernets2Transactor{contract: contract}, Supernets2Filterer: Supernets2Filterer{contract: contract}}, nil
}

// Supernets2 is an auto generated Go binding around an Ethereum contract.
type Supernets2 struct {
	Supernets2Caller     // Read-only binding to the contract
	Supernets2Transactor // Write-only binding to the contract
	Supernets2Filterer   // Log filterer for contract events
}

// Supernets2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Supernets2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Supernets2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Supernets2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Supernets2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Supernets2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Supernets2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Supernets2Session struct {
	Contract     *Supernets2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Supernets2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Supernets2CallerSession struct {
	Contract *Supernets2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Supernets2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Supernets2TransactorSession struct {
	Contract     *Supernets2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Supernets2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Supernets2Raw struct {
	Contract *Supernets2 // Generic contract binding to access the raw methods on
}

// Supernets2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Supernets2CallerRaw struct {
	Contract *Supernets2Caller // Generic read-only contract binding to access the raw methods on
}

// Supernets2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Supernets2TransactorRaw struct {
	Contract *Supernets2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSupernets2 creates a new instance of Supernets2, bound to a specific deployed contract.
func NewSupernets2(address common.Address, backend bind.ContractBackend) (*Supernets2, error) {
	contract, err := bindSupernets2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Supernets2{Supernets2Caller: Supernets2Caller{contract: contract}, Supernets2Transactor: Supernets2Transactor{contract: contract}, Supernets2Filterer: Supernets2Filterer{contract: contract}}, nil
}

// NewSupernets2Caller creates a new read-only instance of Supernets2, bound to a specific deployed contract.
func NewSupernets2Caller(address common.Address, caller bind.ContractCaller) (*Supernets2Caller, error) {
	contract, err := bindSupernets2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Supernets2Caller{contract: contract}, nil
}

// NewSupernets2Transactor creates a new write-only instance of Supernets2, bound to a specific deployed contract.
func NewSupernets2Transactor(address common.Address, transactor bind.ContractTransactor) (*Supernets2Transactor, error) {
	contract, err := bindSupernets2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Supernets2Transactor{contract: contract}, nil
}

// NewSupernets2Filterer creates a new log filterer instance of Supernets2, bound to a specific deployed contract.
func NewSupernets2Filterer(address common.Address, filterer bind.ContractFilterer) (*Supernets2Filterer, error) {
	contract, err := bindSupernets2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Supernets2Filterer{contract: contract}, nil
}

// bindSupernets2 binds a generic wrapper to an already deployed contract.
func bindSupernets2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Supernets2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Supernets2 *Supernets2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Supernets2.Contract.Supernets2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Supernets2 *Supernets2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supernets2.Contract.Supernets2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Supernets2 *Supernets2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Supernets2.Contract.Supernets2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Supernets2 *Supernets2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Supernets2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Supernets2 *Supernets2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supernets2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Supernets2 *Supernets2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Supernets2.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Supernets2 *Supernets2Caller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Supernets2 *Supernets2Session) Admin() (common.Address, error) {
	return _Supernets2.Contract.Admin(&_Supernets2.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Supernets2 *Supernets2CallerSession) Admin() (common.Address, error) {
	return _Supernets2.Contract.Admin(&_Supernets2.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Supernets2 *Supernets2Caller) BatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "batchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Supernets2 *Supernets2Session) BatchFee() (*big.Int, error) {
	return _Supernets2.Contract.BatchFee(&_Supernets2.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Supernets2 *Supernets2CallerSession) BatchFee() (*big.Int, error) {
	return _Supernets2.Contract.BatchFee(&_Supernets2.CallOpts)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Supernets2 *Supernets2Caller) BatchNumToStateRoot(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "batchNumToStateRoot", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Supernets2 *Supernets2Session) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Supernets2.Contract.BatchNumToStateRoot(&_Supernets2.CallOpts, arg0)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Supernets2 *Supernets2CallerSession) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Supernets2.Contract.BatchNumToStateRoot(&_Supernets2.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Supernets2 *Supernets2Caller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Supernets2 *Supernets2Session) BridgeAddress() (common.Address, error) {
	return _Supernets2.Contract.BridgeAddress(&_Supernets2.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Supernets2 *Supernets2CallerSession) BridgeAddress() (common.Address, error) {
	return _Supernets2.Contract.BridgeAddress(&_Supernets2.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Supernets2 *Supernets2Caller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Supernets2 *Supernets2Session) CalculateRewardPerBatch() (*big.Int, error) {
	return _Supernets2.Contract.CalculateRewardPerBatch(&_Supernets2.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Supernets2 *Supernets2CallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Supernets2.Contract.CalculateRewardPerBatch(&_Supernets2.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Supernets2 *Supernets2Caller) ChainID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Supernets2 *Supernets2Session) ChainID() (uint64, error) {
	return _Supernets2.Contract.ChainID(&_Supernets2.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) ChainID() (uint64, error) {
	return _Supernets2.Contract.ChainID(&_Supernets2.CallOpts)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Supernets2 *Supernets2Caller) CheckStateRootInsidePrime(opts *bind.CallOpts, newStateRoot *big.Int) (bool, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "checkStateRootInsidePrime", newStateRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Supernets2 *Supernets2Session) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Supernets2.Contract.CheckStateRootInsidePrime(&_Supernets2.CallOpts, newStateRoot)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Supernets2 *Supernets2CallerSession) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Supernets2.Contract.CheckStateRootInsidePrime(&_Supernets2.CallOpts, newStateRoot)
}

// DataCommitteeAddress is a free data retrieval call binding the contract method 0x4df61d24.
//
// Solidity: function dataCommitteeAddress() view returns(address)
func (_Supernets2 *Supernets2Caller) DataCommitteeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "dataCommitteeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataCommitteeAddress is a free data retrieval call binding the contract method 0x4df61d24.
//
// Solidity: function dataCommitteeAddress() view returns(address)
func (_Supernets2 *Supernets2Session) DataCommitteeAddress() (common.Address, error) {
	return _Supernets2.Contract.DataCommitteeAddress(&_Supernets2.CallOpts)
}

// DataCommitteeAddress is a free data retrieval call binding the contract method 0x4df61d24.
//
// Solidity: function dataCommitteeAddress() view returns(address)
func (_Supernets2 *Supernets2CallerSession) DataCommitteeAddress() (common.Address, error) {
	return _Supernets2.Contract.DataCommitteeAddress(&_Supernets2.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Supernets2 *Supernets2Caller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Supernets2 *Supernets2Session) ForceBatchTimeout() (uint64, error) {
	return _Supernets2.Contract.ForceBatchTimeout(&_Supernets2.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) ForceBatchTimeout() (uint64, error) {
	return _Supernets2.Contract.ForceBatchTimeout(&_Supernets2.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Supernets2 *Supernets2Caller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Supernets2 *Supernets2Session) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Supernets2.Contract.ForcedBatches(&_Supernets2.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Supernets2 *Supernets2CallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Supernets2.Contract.ForcedBatches(&_Supernets2.CallOpts, arg0)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Supernets2 *Supernets2Caller) ForkID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "forkID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Supernets2 *Supernets2Session) ForkID() (uint64, error) {
	return _Supernets2.Contract.ForkID(&_Supernets2.CallOpts)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) ForkID() (uint64, error) {
	return _Supernets2.Contract.ForkID(&_Supernets2.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Supernets2 *Supernets2Caller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Supernets2 *Supernets2Session) GetForcedBatchFee() (*big.Int, error) {
	return _Supernets2.Contract.GetForcedBatchFee(&_Supernets2.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Supernets2 *Supernets2CallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Supernets2.Contract.GetForcedBatchFee(&_Supernets2.CallOpts)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Supernets2 *Supernets2Caller) GetInputSnarkBytes(opts *bind.CallOpts, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "getInputSnarkBytes", initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Supernets2 *Supernets2Session) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Supernets2.Contract.GetInputSnarkBytes(&_Supernets2.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Supernets2 *Supernets2CallerSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Supernets2.Contract.GetInputSnarkBytes(&_Supernets2.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Supernets2 *Supernets2Caller) GetLastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "getLastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Supernets2 *Supernets2Session) GetLastVerifiedBatch() (uint64, error) {
	return _Supernets2.Contract.GetLastVerifiedBatch(&_Supernets2.CallOpts)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) GetLastVerifiedBatch() (uint64, error) {
	return _Supernets2.Contract.GetLastVerifiedBatch(&_Supernets2.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Supernets2 *Supernets2Caller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Supernets2 *Supernets2Session) GlobalExitRootManager() (common.Address, error) {
	return _Supernets2.Contract.GlobalExitRootManager(&_Supernets2.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Supernets2 *Supernets2CallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Supernets2.Contract.GlobalExitRootManager(&_Supernets2.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Supernets2 *Supernets2Caller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Supernets2 *Supernets2Session) IsEmergencyState() (bool, error) {
	return _Supernets2.Contract.IsEmergencyState(&_Supernets2.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Supernets2 *Supernets2CallerSession) IsEmergencyState() (bool, error) {
	return _Supernets2.Contract.IsEmergencyState(&_Supernets2.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Supernets2 *Supernets2Caller) IsForcedBatchDisallowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "isForcedBatchDisallowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Supernets2 *Supernets2Session) IsForcedBatchDisallowed() (bool, error) {
	return _Supernets2.Contract.IsForcedBatchDisallowed(&_Supernets2.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Supernets2 *Supernets2CallerSession) IsForcedBatchDisallowed() (bool, error) {
	return _Supernets2.Contract.IsForcedBatchDisallowed(&_Supernets2.CallOpts)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Supernets2 *Supernets2Caller) IsPendingStateConsolidable(opts *bind.CallOpts, pendingStateNum uint64) (bool, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "isPendingStateConsolidable", pendingStateNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Supernets2 *Supernets2Session) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Supernets2.Contract.IsPendingStateConsolidable(&_Supernets2.CallOpts, pendingStateNum)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Supernets2 *Supernets2CallerSession) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Supernets2.Contract.IsPendingStateConsolidable(&_Supernets2.CallOpts, pendingStateNum)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastBatchSequenced() (uint64, error) {
	return _Supernets2.Contract.LastBatchSequenced(&_Supernets2.CallOpts)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastBatchSequenced() (uint64, error) {
	return _Supernets2.Contract.LastBatchSequenced(&_Supernets2.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastForceBatch() (uint64, error) {
	return _Supernets2.Contract.LastForceBatch(&_Supernets2.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastForceBatch() (uint64, error) {
	return _Supernets2.Contract.LastForceBatch(&_Supernets2.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastForceBatchSequenced() (uint64, error) {
	return _Supernets2.Contract.LastForceBatchSequenced(&_Supernets2.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Supernets2.Contract.LastForceBatchSequenced(&_Supernets2.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastPendingState(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastPendingState")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastPendingState() (uint64, error) {
	return _Supernets2.Contract.LastPendingState(&_Supernets2.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastPendingState() (uint64, error) {
	return _Supernets2.Contract.LastPendingState(&_Supernets2.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastPendingStateConsolidated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastPendingStateConsolidated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastPendingStateConsolidated() (uint64, error) {
	return _Supernets2.Contract.LastPendingStateConsolidated(&_Supernets2.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastPendingStateConsolidated() (uint64, error) {
	return _Supernets2.Contract.LastPendingStateConsolidated(&_Supernets2.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastTimestamp() (uint64, error) {
	return _Supernets2.Contract.LastTimestamp(&_Supernets2.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastTimestamp() (uint64, error) {
	return _Supernets2.Contract.LastTimestamp(&_Supernets2.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Supernets2 *Supernets2Caller) LastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "lastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Supernets2 *Supernets2Session) LastVerifiedBatch() (uint64, error) {
	return _Supernets2.Contract.LastVerifiedBatch(&_Supernets2.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) LastVerifiedBatch() (uint64, error) {
	return _Supernets2.Contract.LastVerifiedBatch(&_Supernets2.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Supernets2 *Supernets2Caller) Matic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "matic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Supernets2 *Supernets2Session) Matic() (common.Address, error) {
	return _Supernets2.Contract.Matic(&_Supernets2.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Supernets2 *Supernets2CallerSession) Matic() (common.Address, error) {
	return _Supernets2.Contract.Matic(&_Supernets2.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Supernets2 *Supernets2Caller) MultiplierBatchFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "multiplierBatchFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Supernets2 *Supernets2Session) MultiplierBatchFee() (uint16, error) {
	return _Supernets2.Contract.MultiplierBatchFee(&_Supernets2.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Supernets2 *Supernets2CallerSession) MultiplierBatchFee() (uint16, error) {
	return _Supernets2.Contract.MultiplierBatchFee(&_Supernets2.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Supernets2 *Supernets2Caller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Supernets2 *Supernets2Session) NetworkName() (string, error) {
	return _Supernets2.Contract.NetworkName(&_Supernets2.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Supernets2 *Supernets2CallerSession) NetworkName() (string, error) {
	return _Supernets2.Contract.NetworkName(&_Supernets2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Supernets2 *Supernets2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Supernets2 *Supernets2Session) Owner() (common.Address, error) {
	return _Supernets2.Contract.Owner(&_Supernets2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Supernets2 *Supernets2CallerSession) Owner() (common.Address, error) {
	return _Supernets2.Contract.Owner(&_Supernets2.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Supernets2 *Supernets2Caller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Supernets2 *Supernets2Session) PendingAdmin() (common.Address, error) {
	return _Supernets2.Contract.PendingAdmin(&_Supernets2.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Supernets2 *Supernets2CallerSession) PendingAdmin() (common.Address, error) {
	return _Supernets2.Contract.PendingAdmin(&_Supernets2.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Supernets2 *Supernets2Caller) PendingStateTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "pendingStateTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Supernets2 *Supernets2Session) PendingStateTimeout() (uint64, error) {
	return _Supernets2.Contract.PendingStateTimeout(&_Supernets2.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) PendingStateTimeout() (uint64, error) {
	return _Supernets2.Contract.PendingStateTimeout(&_Supernets2.CallOpts)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Supernets2 *Supernets2Caller) PendingStateTransitions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "pendingStateTransitions", arg0)

	outstruct := new(struct {
		Timestamp         uint64
		LastVerifiedBatch uint64
		ExitRoot          [32]byte
		StateRoot         [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.LastVerifiedBatch = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ExitRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StateRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Supernets2 *Supernets2Session) PendingStateTransitions(arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	return _Supernets2.Contract.PendingStateTransitions(&_Supernets2.CallOpts, arg0)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Supernets2 *Supernets2CallerSession) PendingStateTransitions(arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	return _Supernets2.Contract.PendingStateTransitions(&_Supernets2.CallOpts, arg0)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Supernets2 *Supernets2Caller) RollupVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "rollupVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Supernets2 *Supernets2Session) RollupVerifier() (common.Address, error) {
	return _Supernets2.Contract.RollupVerifier(&_Supernets2.CallOpts)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Supernets2 *Supernets2CallerSession) RollupVerifier() (common.Address, error) {
	return _Supernets2.Contract.RollupVerifier(&_Supernets2.CallOpts)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Supernets2 *Supernets2Caller) SequencedBatches(opts *bind.CallOpts, arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "sequencedBatches", arg0)

	outstruct := new(struct {
		AccInputHash               [32]byte
		SequencedTimestamp         uint64
		PreviousLastBatchSequenced uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccInputHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.SequencedTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.PreviousLastBatchSequenced = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Supernets2 *Supernets2Session) SequencedBatches(arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	return _Supernets2.Contract.SequencedBatches(&_Supernets2.CallOpts, arg0)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Supernets2 *Supernets2CallerSession) SequencedBatches(arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	return _Supernets2.Contract.SequencedBatches(&_Supernets2.CallOpts, arg0)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Supernets2 *Supernets2Caller) TrustedAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "trustedAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Supernets2 *Supernets2Session) TrustedAggregator() (common.Address, error) {
	return _Supernets2.Contract.TrustedAggregator(&_Supernets2.CallOpts)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Supernets2 *Supernets2CallerSession) TrustedAggregator() (common.Address, error) {
	return _Supernets2.Contract.TrustedAggregator(&_Supernets2.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Supernets2 *Supernets2Caller) TrustedAggregatorTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "trustedAggregatorTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Supernets2 *Supernets2Session) TrustedAggregatorTimeout() (uint64, error) {
	return _Supernets2.Contract.TrustedAggregatorTimeout(&_Supernets2.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Supernets2.Contract.TrustedAggregatorTimeout(&_Supernets2.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Supernets2 *Supernets2Caller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Supernets2 *Supernets2Session) TrustedSequencer() (common.Address, error) {
	return _Supernets2.Contract.TrustedSequencer(&_Supernets2.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Supernets2 *Supernets2CallerSession) TrustedSequencer() (common.Address, error) {
	return _Supernets2.Contract.TrustedSequencer(&_Supernets2.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Supernets2 *Supernets2Caller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Supernets2 *Supernets2Session) TrustedSequencerURL() (string, error) {
	return _Supernets2.Contract.TrustedSequencerURL(&_Supernets2.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Supernets2 *Supernets2CallerSession) TrustedSequencerURL() (string, error) {
	return _Supernets2.Contract.TrustedSequencerURL(&_Supernets2.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Supernets2 *Supernets2Caller) VerifyBatchTimeTarget(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Supernets2.contract.Call(opts, &out, "verifyBatchTimeTarget")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Supernets2 *Supernets2Session) VerifyBatchTimeTarget() (uint64, error) {
	return _Supernets2.Contract.VerifyBatchTimeTarget(&_Supernets2.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Supernets2 *Supernets2CallerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Supernets2.Contract.VerifyBatchTimeTarget(&_Supernets2.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Supernets2 *Supernets2Transactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Supernets2 *Supernets2Session) AcceptAdminRole() (*types.Transaction, error) {
	return _Supernets2.Contract.AcceptAdminRole(&_Supernets2.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Supernets2 *Supernets2TransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Supernets2.Contract.AcceptAdminRole(&_Supernets2.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Supernets2 *Supernets2Transactor) ActivateEmergencyState(opts *bind.TransactOpts, sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "activateEmergencyState", sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Supernets2 *Supernets2Session) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.ActivateEmergencyState(&_Supernets2.TransactOpts, sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Supernets2 *Supernets2TransactorSession) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.ActivateEmergencyState(&_Supernets2.TransactOpts, sequencedBatchNum)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Supernets2 *Supernets2Transactor) ActivateForceBatches(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "activateForceBatches")
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Supernets2 *Supernets2Session) ActivateForceBatches() (*types.Transaction, error) {
	return _Supernets2.Contract.ActivateForceBatches(&_Supernets2.TransactOpts)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Supernets2 *Supernets2TransactorSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Supernets2.Contract.ActivateForceBatches(&_Supernets2.TransactOpts)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Supernets2 *Supernets2Transactor) ConsolidatePendingState(opts *bind.TransactOpts, pendingStateNum uint64) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "consolidatePendingState", pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Supernets2 *Supernets2Session) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.ConsolidatePendingState(&_Supernets2.TransactOpts, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Supernets2 *Supernets2TransactorSession) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.ConsolidatePendingState(&_Supernets2.TransactOpts, pendingStateNum)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Supernets2 *Supernets2Transactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Supernets2 *Supernets2Session) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Supernets2.Contract.DeactivateEmergencyState(&_Supernets2.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Supernets2 *Supernets2TransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Supernets2.Contract.DeactivateEmergencyState(&_Supernets2.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Supernets2 *Supernets2Transactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "forceBatch", transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Supernets2 *Supernets2Session) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Supernets2.Contract.ForceBatch(&_Supernets2.TransactOpts, transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Supernets2 *Supernets2TransactorSession) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Supernets2.Contract.ForceBatch(&_Supernets2.TransactOpts, transactions, maticAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Supernets2 *Supernets2Transactor) Initialize(opts *bind.TransactOpts, initializePackedParameters Supernets2InitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "initialize", initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Supernets2 *Supernets2Session) Initialize(initializePackedParameters Supernets2InitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Supernets2.Contract.Initialize(&_Supernets2.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Supernets2 *Supernets2TransactorSession) Initialize(initializePackedParameters Supernets2InitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Supernets2.Contract.Initialize(&_Supernets2.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Supernets2 *Supernets2Transactor) OverridePendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "overridePendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Supernets2 *Supernets2Session) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Supernets2.Contract.OverridePendingState(&_Supernets2.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Supernets2 *Supernets2TransactorSession) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Supernets2.Contract.OverridePendingState(&_Supernets2.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Supernets2 *Supernets2Transactor) ProveNonDeterministicPendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "proveNonDeterministicPendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Supernets2 *Supernets2Session) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Supernets2.Contract.ProveNonDeterministicPendingState(&_Supernets2.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Supernets2 *Supernets2TransactorSession) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Supernets2.Contract.ProveNonDeterministicPendingState(&_Supernets2.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Supernets2 *Supernets2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Supernets2 *Supernets2Session) RenounceOwnership() (*types.Transaction, error) {
	return _Supernets2.Contract.RenounceOwnership(&_Supernets2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Supernets2 *Supernets2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Supernets2.Contract.RenounceOwnership(&_Supernets2.TransactOpts)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x438a5399.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes signaturesAndAddrs) returns()
func (_Supernets2 *Supernets2Transactor) SequenceBatches(opts *bind.TransactOpts, batches []Supernets2BatchData, l2Coinbase common.Address, signaturesAndAddrs []byte) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "sequenceBatches", batches, l2Coinbase, signaturesAndAddrs)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x438a5399.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes signaturesAndAddrs) returns()
func (_Supernets2 *Supernets2Session) SequenceBatches(batches []Supernets2BatchData, l2Coinbase common.Address, signaturesAndAddrs []byte) (*types.Transaction, error) {
	return _Supernets2.Contract.SequenceBatches(&_Supernets2.TransactOpts, batches, l2Coinbase, signaturesAndAddrs)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x438a5399.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes signaturesAndAddrs) returns()
func (_Supernets2 *Supernets2TransactorSession) SequenceBatches(batches []Supernets2BatchData, l2Coinbase common.Address, signaturesAndAddrs []byte) (*types.Transaction, error) {
	return _Supernets2.Contract.SequenceBatches(&_Supernets2.TransactOpts, batches, l2Coinbase, signaturesAndAddrs)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Supernets2 *Supernets2Transactor) SequenceForceBatches(opts *bind.TransactOpts, batches []Supernets2ForcedBatchData) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Supernets2 *Supernets2Session) SequenceForceBatches(batches []Supernets2ForcedBatchData) (*types.Transaction, error) {
	return _Supernets2.Contract.SequenceForceBatches(&_Supernets2.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Supernets2 *Supernets2TransactorSession) SequenceForceBatches(batches []Supernets2ForcedBatchData) (*types.Transaction, error) {
	return _Supernets2.Contract.SequenceForceBatches(&_Supernets2.TransactOpts, batches)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Supernets2 *Supernets2Transactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Supernets2 *Supernets2Session) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetForceBatchTimeout(&_Supernets2.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Supernets2 *Supernets2TransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetForceBatchTimeout(&_Supernets2.TransactOpts, newforceBatchTimeout)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Supernets2 *Supernets2Transactor) SetMultiplierBatchFee(opts *bind.TransactOpts, newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setMultiplierBatchFee", newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Supernets2 *Supernets2Session) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Supernets2.Contract.SetMultiplierBatchFee(&_Supernets2.TransactOpts, newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Supernets2 *Supernets2TransactorSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Supernets2.Contract.SetMultiplierBatchFee(&_Supernets2.TransactOpts, newMultiplierBatchFee)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Supernets2 *Supernets2Transactor) SetPendingStateTimeout(opts *bind.TransactOpts, newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setPendingStateTimeout", newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Supernets2 *Supernets2Session) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetPendingStateTimeout(&_Supernets2.TransactOpts, newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Supernets2 *Supernets2TransactorSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetPendingStateTimeout(&_Supernets2.TransactOpts, newPendingStateTimeout)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Supernets2 *Supernets2Transactor) SetTrustedAggregator(opts *bind.TransactOpts, newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setTrustedAggregator", newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Supernets2 *Supernets2Session) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedAggregator(&_Supernets2.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Supernets2 *Supernets2TransactorSession) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedAggregator(&_Supernets2.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Supernets2 *Supernets2Transactor) SetTrustedAggregatorTimeout(opts *bind.TransactOpts, newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setTrustedAggregatorTimeout", newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Supernets2 *Supernets2Session) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedAggregatorTimeout(&_Supernets2.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Supernets2 *Supernets2TransactorSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedAggregatorTimeout(&_Supernets2.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Supernets2 *Supernets2Transactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Supernets2 *Supernets2Session) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedSequencer(&_Supernets2.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Supernets2 *Supernets2TransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedSequencer(&_Supernets2.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Supernets2 *Supernets2Transactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Supernets2 *Supernets2Session) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedSequencerURL(&_Supernets2.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Supernets2 *Supernets2TransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Supernets2.Contract.SetTrustedSequencerURL(&_Supernets2.TransactOpts, newTrustedSequencerURL)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Supernets2 *Supernets2Transactor) SetVerifyBatchTimeTarget(opts *bind.TransactOpts, newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "setVerifyBatchTimeTarget", newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Supernets2 *Supernets2Session) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetVerifyBatchTimeTarget(&_Supernets2.TransactOpts, newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Supernets2 *Supernets2TransactorSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Supernets2.Contract.SetVerifyBatchTimeTarget(&_Supernets2.TransactOpts, newVerifyBatchTimeTarget)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Supernets2 *Supernets2Transactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Supernets2 *Supernets2Session) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.TransferAdminRole(&_Supernets2.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Supernets2 *Supernets2TransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.TransferAdminRole(&_Supernets2.TransactOpts, newPendingAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Supernets2 *Supernets2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Supernets2 *Supernets2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.TransferOwnership(&_Supernets2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Supernets2 *Supernets2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Supernets2.Contract.TransferOwnership(&_Supernets2.TransactOpts, newOwner)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Supernets2 *Supernets2Transactor) VerifyBatches(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "verifyBatches", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Supernets2 *Supernets2Session) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Supernets2.Contract.VerifyBatches(&_Supernets2.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Supernets2 *Supernets2TransactorSession) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Supernets2.Contract.VerifyBatches(&_Supernets2.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Supernets2 *Supernets2Transactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Supernets2.contract.Transact(opts, "verifyBatchesTrustedAggregator", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Supernets2 *Supernets2Session) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Supernets2.Contract.VerifyBatchesTrustedAggregator(&_Supernets2.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Supernets2 *Supernets2TransactorSession) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Supernets2.Contract.VerifyBatchesTrustedAggregator(&_Supernets2.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// Supernets2AcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Supernets2 contract.
type Supernets2AcceptAdminRoleIterator struct {
	Event *Supernets2AcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *Supernets2AcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2AcceptAdminRole)
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
		it.Event = new(Supernets2AcceptAdminRole)
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
func (it *Supernets2AcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2AcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2AcceptAdminRole represents a AcceptAdminRole event raised by the Supernets2 contract.
type Supernets2AcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Supernets2 *Supernets2Filterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*Supernets2AcceptAdminRoleIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &Supernets2AcceptAdminRoleIterator{contract: _Supernets2.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Supernets2 *Supernets2Filterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *Supernets2AcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2AcceptAdminRole)
				if err := _Supernets2.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Supernets2 *Supernets2Filterer) ParseAcceptAdminRole(log types.Log) (*Supernets2AcceptAdminRole, error) {
	event := new(Supernets2AcceptAdminRole)
	if err := _Supernets2.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2ActivateForceBatchesIterator is returned from FilterActivateForceBatches and is used to iterate over the raw logs and unpacked data for ActivateForceBatches events raised by the Supernets2 contract.
type Supernets2ActivateForceBatchesIterator struct {
	Event *Supernets2ActivateForceBatches // Event containing the contract specifics and raw log

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
func (it *Supernets2ActivateForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2ActivateForceBatches)
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
		it.Event = new(Supernets2ActivateForceBatches)
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
func (it *Supernets2ActivateForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2ActivateForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2ActivateForceBatches represents a ActivateForceBatches event raised by the Supernets2 contract.
type Supernets2ActivateForceBatches struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActivateForceBatches is a free log retrieval operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Supernets2 *Supernets2Filterer) FilterActivateForceBatches(opts *bind.FilterOpts) (*Supernets2ActivateForceBatchesIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return &Supernets2ActivateForceBatchesIterator{contract: _Supernets2.contract, event: "ActivateForceBatches", logs: logs, sub: sub}, nil
}

// WatchActivateForceBatches is a free log subscription operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Supernets2 *Supernets2Filterer) WatchActivateForceBatches(opts *bind.WatchOpts, sink chan<- *Supernets2ActivateForceBatches) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2ActivateForceBatches)
				if err := _Supernets2.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
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
func (_Supernets2 *Supernets2Filterer) ParseActivateForceBatches(log types.Log) (*Supernets2ActivateForceBatches, error) {
	event := new(Supernets2ActivateForceBatches)
	if err := _Supernets2.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2ConsolidatePendingStateIterator is returned from FilterConsolidatePendingState and is used to iterate over the raw logs and unpacked data for ConsolidatePendingState events raised by the Supernets2 contract.
type Supernets2ConsolidatePendingStateIterator struct {
	Event *Supernets2ConsolidatePendingState // Event containing the contract specifics and raw log

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
func (it *Supernets2ConsolidatePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2ConsolidatePendingState)
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
		it.Event = new(Supernets2ConsolidatePendingState)
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
func (it *Supernets2ConsolidatePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2ConsolidatePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2ConsolidatePendingState represents a ConsolidatePendingState event raised by the Supernets2 contract.
type Supernets2ConsolidatePendingState struct {
	NumBatch        uint64
	StateRoot       [32]byte
	PendingStateNum uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterConsolidatePendingState is a free log retrieval operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Supernets2 *Supernets2Filterer) FilterConsolidatePendingState(opts *bind.FilterOpts, numBatch []uint64, pendingStateNum []uint64) (*Supernets2ConsolidatePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2ConsolidatePendingStateIterator{contract: _Supernets2.contract, event: "ConsolidatePendingState", logs: logs, sub: sub}, nil
}

// WatchConsolidatePendingState is a free log subscription operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Supernets2 *Supernets2Filterer) WatchConsolidatePendingState(opts *bind.WatchOpts, sink chan<- *Supernets2ConsolidatePendingState, numBatch []uint64, pendingStateNum []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2ConsolidatePendingState)
				if err := _Supernets2.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
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

// ParseConsolidatePendingState is a log parse operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Supernets2 *Supernets2Filterer) ParseConsolidatePendingState(log types.Log) (*Supernets2ConsolidatePendingState, error) {
	event := new(Supernets2ConsolidatePendingState)
	if err := _Supernets2.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2EmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Supernets2 contract.
type Supernets2EmergencyStateActivatedIterator struct {
	Event *Supernets2EmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *Supernets2EmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2EmergencyStateActivated)
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
		it.Event = new(Supernets2EmergencyStateActivated)
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
func (it *Supernets2EmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2EmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2EmergencyStateActivated represents a EmergencyStateActivated event raised by the Supernets2 contract.
type Supernets2EmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Supernets2 *Supernets2Filterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*Supernets2EmergencyStateActivatedIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &Supernets2EmergencyStateActivatedIterator{contract: _Supernets2.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Supernets2 *Supernets2Filterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *Supernets2EmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2EmergencyStateActivated)
				if err := _Supernets2.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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

// ParseEmergencyStateActivated is a log parse operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Supernets2 *Supernets2Filterer) ParseEmergencyStateActivated(log types.Log) (*Supernets2EmergencyStateActivated, error) {
	event := new(Supernets2EmergencyStateActivated)
	if err := _Supernets2.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2EmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Supernets2 contract.
type Supernets2EmergencyStateDeactivatedIterator struct {
	Event *Supernets2EmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *Supernets2EmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2EmergencyStateDeactivated)
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
		it.Event = new(Supernets2EmergencyStateDeactivated)
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
func (it *Supernets2EmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2EmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2EmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Supernets2 contract.
type Supernets2EmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Supernets2 *Supernets2Filterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*Supernets2EmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &Supernets2EmergencyStateDeactivatedIterator{contract: _Supernets2.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Supernets2 *Supernets2Filterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *Supernets2EmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2EmergencyStateDeactivated)
				if err := _Supernets2.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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

// ParseEmergencyStateDeactivated is a log parse operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Supernets2 *Supernets2Filterer) ParseEmergencyStateDeactivated(log types.Log) (*Supernets2EmergencyStateDeactivated, error) {
	event := new(Supernets2EmergencyStateDeactivated)
	if err := _Supernets2.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2ForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Supernets2 contract.
type Supernets2ForceBatchIterator struct {
	Event *Supernets2ForceBatch // Event containing the contract specifics and raw log

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
func (it *Supernets2ForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2ForceBatch)
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
		it.Event = new(Supernets2ForceBatch)
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
func (it *Supernets2ForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2ForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2ForceBatch represents a ForceBatch event raised by the Supernets2 contract.
type Supernets2ForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Supernets2 *Supernets2Filterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*Supernets2ForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2ForceBatchIterator{contract: _Supernets2.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Supernets2 *Supernets2Filterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *Supernets2ForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2ForceBatch)
				if err := _Supernets2.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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
func (_Supernets2 *Supernets2Filterer) ParseForceBatch(log types.Log) (*Supernets2ForceBatch, error) {
	event := new(Supernets2ForceBatch)
	if err := _Supernets2.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Supernets2 contract.
type Supernets2InitializedIterator struct {
	Event *Supernets2Initialized // Event containing the contract specifics and raw log

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
func (it *Supernets2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2Initialized)
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
		it.Event = new(Supernets2Initialized)
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
func (it *Supernets2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2Initialized represents a Initialized event raised by the Supernets2 contract.
type Supernets2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Supernets2 *Supernets2Filterer) FilterInitialized(opts *bind.FilterOpts) (*Supernets2InitializedIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Supernets2InitializedIterator{contract: _Supernets2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Supernets2 *Supernets2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Supernets2Initialized) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2Initialized)
				if err := _Supernets2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Supernets2 *Supernets2Filterer) ParseInitialized(log types.Log) (*Supernets2Initialized, error) {
	event := new(Supernets2Initialized)
	if err := _Supernets2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2OverridePendingStateIterator is returned from FilterOverridePendingState and is used to iterate over the raw logs and unpacked data for OverridePendingState events raised by the Supernets2 contract.
type Supernets2OverridePendingStateIterator struct {
	Event *Supernets2OverridePendingState // Event containing the contract specifics and raw log

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
func (it *Supernets2OverridePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2OverridePendingState)
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
		it.Event = new(Supernets2OverridePendingState)
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
func (it *Supernets2OverridePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2OverridePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2OverridePendingState represents a OverridePendingState event raised by the Supernets2 contract.
type Supernets2OverridePendingState struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOverridePendingState is a free log retrieval operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) FilterOverridePendingState(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*Supernets2OverridePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2OverridePendingStateIterator{contract: _Supernets2.contract, event: "OverridePendingState", logs: logs, sub: sub}, nil
}

// WatchOverridePendingState is a free log subscription operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) WatchOverridePendingState(opts *bind.WatchOpts, sink chan<- *Supernets2OverridePendingState, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2OverridePendingState)
				if err := _Supernets2.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
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

// ParseOverridePendingState is a log parse operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) ParseOverridePendingState(log types.Log) (*Supernets2OverridePendingState, error) {
	event := new(Supernets2OverridePendingState)
	if err := _Supernets2.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Supernets2 contract.
type Supernets2OwnershipTransferredIterator struct {
	Event *Supernets2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Supernets2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2OwnershipTransferred)
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
		it.Event = new(Supernets2OwnershipTransferred)
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
func (it *Supernets2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2OwnershipTransferred represents a OwnershipTransferred event raised by the Supernets2 contract.
type Supernets2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Supernets2 *Supernets2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Supernets2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2OwnershipTransferredIterator{contract: _Supernets2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Supernets2 *Supernets2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Supernets2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2OwnershipTransferred)
				if err := _Supernets2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Supernets2 *Supernets2Filterer) ParseOwnershipTransferred(log types.Log) (*Supernets2OwnershipTransferred, error) {
	event := new(Supernets2OwnershipTransferred)
	if err := _Supernets2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2ProveNonDeterministicPendingStateIterator is returned from FilterProveNonDeterministicPendingState and is used to iterate over the raw logs and unpacked data for ProveNonDeterministicPendingState events raised by the Supernets2 contract.
type Supernets2ProveNonDeterministicPendingStateIterator struct {
	Event *Supernets2ProveNonDeterministicPendingState // Event containing the contract specifics and raw log

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
func (it *Supernets2ProveNonDeterministicPendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2ProveNonDeterministicPendingState)
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
		it.Event = new(Supernets2ProveNonDeterministicPendingState)
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
func (it *Supernets2ProveNonDeterministicPendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2ProveNonDeterministicPendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2ProveNonDeterministicPendingState represents a ProveNonDeterministicPendingState event raised by the Supernets2 contract.
type Supernets2ProveNonDeterministicPendingState struct {
	StoredStateRoot [32]byte
	ProvedStateRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProveNonDeterministicPendingState is a free log retrieval operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Supernets2 *Supernets2Filterer) FilterProveNonDeterministicPendingState(opts *bind.FilterOpts) (*Supernets2ProveNonDeterministicPendingStateIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return &Supernets2ProveNonDeterministicPendingStateIterator{contract: _Supernets2.contract, event: "ProveNonDeterministicPendingState", logs: logs, sub: sub}, nil
}

// WatchProveNonDeterministicPendingState is a free log subscription operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Supernets2 *Supernets2Filterer) WatchProveNonDeterministicPendingState(opts *bind.WatchOpts, sink chan<- *Supernets2ProveNonDeterministicPendingState) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2ProveNonDeterministicPendingState)
				if err := _Supernets2.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
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

// ParseProveNonDeterministicPendingState is a log parse operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Supernets2 *Supernets2Filterer) ParseProveNonDeterministicPendingState(log types.Log) (*Supernets2ProveNonDeterministicPendingState, error) {
	event := new(Supernets2ProveNonDeterministicPendingState)
	if err := _Supernets2.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Supernets2 contract.
type Supernets2SequenceBatchesIterator struct {
	Event *Supernets2SequenceBatches // Event containing the contract specifics and raw log

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
func (it *Supernets2SequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SequenceBatches)
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
		it.Event = new(Supernets2SequenceBatches)
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
func (it *Supernets2SequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SequenceBatches represents a SequenceBatches event raised by the Supernets2 contract.
type Supernets2SequenceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Supernets2 *Supernets2Filterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*Supernets2SequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2SequenceBatchesIterator{contract: _Supernets2.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Supernets2 *Supernets2Filterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *Supernets2SequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SequenceBatches)
				if err := _Supernets2.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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
func (_Supernets2 *Supernets2Filterer) ParseSequenceBatches(log types.Log) (*Supernets2SequenceBatches, error) {
	event := new(Supernets2SequenceBatches)
	if err := _Supernets2.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Supernets2 contract.
type Supernets2SequenceForceBatchesIterator struct {
	Event *Supernets2SequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *Supernets2SequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SequenceForceBatches)
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
		it.Event = new(Supernets2SequenceForceBatches)
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
func (it *Supernets2SequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SequenceForceBatches represents a SequenceForceBatches event raised by the Supernets2 contract.
type Supernets2SequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Supernets2 *Supernets2Filterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*Supernets2SequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2SequenceForceBatchesIterator{contract: _Supernets2.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Supernets2 *Supernets2Filterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *Supernets2SequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SequenceForceBatches)
				if err := _Supernets2.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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
func (_Supernets2 *Supernets2Filterer) ParseSequenceForceBatches(log types.Log) (*Supernets2SequenceForceBatches, error) {
	event := new(Supernets2SequenceForceBatches)
	if err := _Supernets2.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Supernets2 contract.
type Supernets2SetForceBatchTimeoutIterator struct {
	Event *Supernets2SetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *Supernets2SetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetForceBatchTimeout)
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
		it.Event = new(Supernets2SetForceBatchTimeout)
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
func (it *Supernets2SetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Supernets2 contract.
type Supernets2SetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Supernets2 *Supernets2Filterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*Supernets2SetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetForceBatchTimeoutIterator{contract: _Supernets2.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Supernets2 *Supernets2Filterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *Supernets2SetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetForceBatchTimeout)
				if err := _Supernets2.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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
func (_Supernets2 *Supernets2Filterer) ParseSetForceBatchTimeout(log types.Log) (*Supernets2SetForceBatchTimeout, error) {
	event := new(Supernets2SetForceBatchTimeout)
	if err := _Supernets2.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetMultiplierBatchFeeIterator is returned from FilterSetMultiplierBatchFee and is used to iterate over the raw logs and unpacked data for SetMultiplierBatchFee events raised by the Supernets2 contract.
type Supernets2SetMultiplierBatchFeeIterator struct {
	Event *Supernets2SetMultiplierBatchFee // Event containing the contract specifics and raw log

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
func (it *Supernets2SetMultiplierBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetMultiplierBatchFee)
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
		it.Event = new(Supernets2SetMultiplierBatchFee)
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
func (it *Supernets2SetMultiplierBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetMultiplierBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetMultiplierBatchFee represents a SetMultiplierBatchFee event raised by the Supernets2 contract.
type Supernets2SetMultiplierBatchFee struct {
	NewMultiplierBatchFee uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMultiplierBatchFee is a free log retrieval operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Supernets2 *Supernets2Filterer) FilterSetMultiplierBatchFee(opts *bind.FilterOpts) (*Supernets2SetMultiplierBatchFeeIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetMultiplierBatchFeeIterator{contract: _Supernets2.contract, event: "SetMultiplierBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetMultiplierBatchFee is a free log subscription operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Supernets2 *Supernets2Filterer) WatchSetMultiplierBatchFee(opts *bind.WatchOpts, sink chan<- *Supernets2SetMultiplierBatchFee) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetMultiplierBatchFee)
				if err := _Supernets2.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
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

// ParseSetMultiplierBatchFee is a log parse operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Supernets2 *Supernets2Filterer) ParseSetMultiplierBatchFee(log types.Log) (*Supernets2SetMultiplierBatchFee, error) {
	event := new(Supernets2SetMultiplierBatchFee)
	if err := _Supernets2.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetPendingStateTimeoutIterator is returned from FilterSetPendingStateTimeout and is used to iterate over the raw logs and unpacked data for SetPendingStateTimeout events raised by the Supernets2 contract.
type Supernets2SetPendingStateTimeoutIterator struct {
	Event *Supernets2SetPendingStateTimeout // Event containing the contract specifics and raw log

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
func (it *Supernets2SetPendingStateTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetPendingStateTimeout)
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
		it.Event = new(Supernets2SetPendingStateTimeout)
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
func (it *Supernets2SetPendingStateTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetPendingStateTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetPendingStateTimeout represents a SetPendingStateTimeout event raised by the Supernets2 contract.
type Supernets2SetPendingStateTimeout struct {
	NewPendingStateTimeout uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetPendingStateTimeout is a free log retrieval operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Supernets2 *Supernets2Filterer) FilterSetPendingStateTimeout(opts *bind.FilterOpts) (*Supernets2SetPendingStateTimeoutIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetPendingStateTimeoutIterator{contract: _Supernets2.contract, event: "SetPendingStateTimeout", logs: logs, sub: sub}, nil
}

// WatchSetPendingStateTimeout is a free log subscription operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Supernets2 *Supernets2Filterer) WatchSetPendingStateTimeout(opts *bind.WatchOpts, sink chan<- *Supernets2SetPendingStateTimeout) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetPendingStateTimeout)
				if err := _Supernets2.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
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

// ParseSetPendingStateTimeout is a log parse operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Supernets2 *Supernets2Filterer) ParseSetPendingStateTimeout(log types.Log) (*Supernets2SetPendingStateTimeout, error) {
	event := new(Supernets2SetPendingStateTimeout)
	if err := _Supernets2.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Supernets2 contract.
type Supernets2SetTrustedAggregatorIterator struct {
	Event *Supernets2SetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *Supernets2SetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetTrustedAggregator)
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
		it.Event = new(Supernets2SetTrustedAggregator)
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
func (it *Supernets2SetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetTrustedAggregator represents a SetTrustedAggregator event raised by the Supernets2 contract.
type Supernets2SetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Supernets2 *Supernets2Filterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*Supernets2SetTrustedAggregatorIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetTrustedAggregatorIterator{contract: _Supernets2.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Supernets2 *Supernets2Filterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *Supernets2SetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetTrustedAggregator)
				if err := _Supernets2.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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

// ParseSetTrustedAggregator is a log parse operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Supernets2 *Supernets2Filterer) ParseSetTrustedAggregator(log types.Log) (*Supernets2SetTrustedAggregator, error) {
	event := new(Supernets2SetTrustedAggregator)
	if err := _Supernets2.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetTrustedAggregatorTimeoutIterator is returned from FilterSetTrustedAggregatorTimeout and is used to iterate over the raw logs and unpacked data for SetTrustedAggregatorTimeout events raised by the Supernets2 contract.
type Supernets2SetTrustedAggregatorTimeoutIterator struct {
	Event *Supernets2SetTrustedAggregatorTimeout // Event containing the contract specifics and raw log

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
func (it *Supernets2SetTrustedAggregatorTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetTrustedAggregatorTimeout)
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
		it.Event = new(Supernets2SetTrustedAggregatorTimeout)
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
func (it *Supernets2SetTrustedAggregatorTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetTrustedAggregatorTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetTrustedAggregatorTimeout represents a SetTrustedAggregatorTimeout event raised by the Supernets2 contract.
type Supernets2SetTrustedAggregatorTimeout struct {
	NewTrustedAggregatorTimeout uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregatorTimeout is a free log retrieval operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Supernets2 *Supernets2Filterer) FilterSetTrustedAggregatorTimeout(opts *bind.FilterOpts) (*Supernets2SetTrustedAggregatorTimeoutIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetTrustedAggregatorTimeoutIterator{contract: _Supernets2.contract, event: "SetTrustedAggregatorTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregatorTimeout is a free log subscription operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Supernets2 *Supernets2Filterer) WatchSetTrustedAggregatorTimeout(opts *bind.WatchOpts, sink chan<- *Supernets2SetTrustedAggregatorTimeout) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetTrustedAggregatorTimeout)
				if err := _Supernets2.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
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

// ParseSetTrustedAggregatorTimeout is a log parse operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Supernets2 *Supernets2Filterer) ParseSetTrustedAggregatorTimeout(log types.Log) (*Supernets2SetTrustedAggregatorTimeout, error) {
	event := new(Supernets2SetTrustedAggregatorTimeout)
	if err := _Supernets2.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Supernets2 contract.
type Supernets2SetTrustedSequencerIterator struct {
	Event *Supernets2SetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *Supernets2SetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetTrustedSequencer)
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
		it.Event = new(Supernets2SetTrustedSequencer)
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
func (it *Supernets2SetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetTrustedSequencer represents a SetTrustedSequencer event raised by the Supernets2 contract.
type Supernets2SetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Supernets2 *Supernets2Filterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*Supernets2SetTrustedSequencerIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetTrustedSequencerIterator{contract: _Supernets2.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Supernets2 *Supernets2Filterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *Supernets2SetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetTrustedSequencer)
				if err := _Supernets2.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Supernets2 *Supernets2Filterer) ParseSetTrustedSequencer(log types.Log) (*Supernets2SetTrustedSequencer, error) {
	event := new(Supernets2SetTrustedSequencer)
	if err := _Supernets2.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Supernets2 contract.
type Supernets2SetTrustedSequencerURLIterator struct {
	Event *Supernets2SetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *Supernets2SetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetTrustedSequencerURL)
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
		it.Event = new(Supernets2SetTrustedSequencerURL)
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
func (it *Supernets2SetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Supernets2 contract.
type Supernets2SetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Supernets2 *Supernets2Filterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*Supernets2SetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetTrustedSequencerURLIterator{contract: _Supernets2.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Supernets2 *Supernets2Filterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *Supernets2SetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetTrustedSequencerURL)
				if err := _Supernets2.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Supernets2 *Supernets2Filterer) ParseSetTrustedSequencerURL(log types.Log) (*Supernets2SetTrustedSequencerURL, error) {
	event := new(Supernets2SetTrustedSequencerURL)
	if err := _Supernets2.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2SetVerifyBatchTimeTargetIterator is returned from FilterSetVerifyBatchTimeTarget and is used to iterate over the raw logs and unpacked data for SetVerifyBatchTimeTarget events raised by the Supernets2 contract.
type Supernets2SetVerifyBatchTimeTargetIterator struct {
	Event *Supernets2SetVerifyBatchTimeTarget // Event containing the contract specifics and raw log

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
func (it *Supernets2SetVerifyBatchTimeTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2SetVerifyBatchTimeTarget)
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
		it.Event = new(Supernets2SetVerifyBatchTimeTarget)
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
func (it *Supernets2SetVerifyBatchTimeTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2SetVerifyBatchTimeTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2SetVerifyBatchTimeTarget represents a SetVerifyBatchTimeTarget event raised by the Supernets2 contract.
type Supernets2SetVerifyBatchTimeTarget struct {
	NewVerifyBatchTimeTarget uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetVerifyBatchTimeTarget is a free log retrieval operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Supernets2 *Supernets2Filterer) FilterSetVerifyBatchTimeTarget(opts *bind.FilterOpts) (*Supernets2SetVerifyBatchTimeTargetIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return &Supernets2SetVerifyBatchTimeTargetIterator{contract: _Supernets2.contract, event: "SetVerifyBatchTimeTarget", logs: logs, sub: sub}, nil
}

// WatchSetVerifyBatchTimeTarget is a free log subscription operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Supernets2 *Supernets2Filterer) WatchSetVerifyBatchTimeTarget(opts *bind.WatchOpts, sink chan<- *Supernets2SetVerifyBatchTimeTarget) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2SetVerifyBatchTimeTarget)
				if err := _Supernets2.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
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

// ParseSetVerifyBatchTimeTarget is a log parse operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Supernets2 *Supernets2Filterer) ParseSetVerifyBatchTimeTarget(log types.Log) (*Supernets2SetVerifyBatchTimeTarget, error) {
	event := new(Supernets2SetVerifyBatchTimeTarget)
	if err := _Supernets2.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2TransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Supernets2 contract.
type Supernets2TransferAdminRoleIterator struct {
	Event *Supernets2TransferAdminRole // Event containing the contract specifics and raw log

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
func (it *Supernets2TransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2TransferAdminRole)
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
		it.Event = new(Supernets2TransferAdminRole)
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
func (it *Supernets2TransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2TransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2TransferAdminRole represents a TransferAdminRole event raised by the Supernets2 contract.
type Supernets2TransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Supernets2 *Supernets2Filterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*Supernets2TransferAdminRoleIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &Supernets2TransferAdminRoleIterator{contract: _Supernets2.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Supernets2 *Supernets2Filterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *Supernets2TransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2TransferAdminRole)
				if err := _Supernets2.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Supernets2 *Supernets2Filterer) ParseTransferAdminRole(log types.Log) (*Supernets2TransferAdminRole, error) {
	event := new(Supernets2TransferAdminRole)
	if err := _Supernets2.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2UpdateSupernets2VersionIterator is returned from FilterUpdateSupernets2Version and is used to iterate over the raw logs and unpacked data for UpdateSupernets2Version events raised by the Supernets2 contract.
type Supernets2UpdateSupernets2VersionIterator struct {
	Event *Supernets2UpdateSupernets2Version // Event containing the contract specifics and raw log

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
func (it *Supernets2UpdateSupernets2VersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2UpdateSupernets2Version)
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
		it.Event = new(Supernets2UpdateSupernets2Version)
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
func (it *Supernets2UpdateSupernets2VersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2UpdateSupernets2VersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2UpdateSupernets2Version represents a UpdateSupernets2Version event raised by the Supernets2 contract.
type Supernets2UpdateSupernets2Version struct {
	NumBatch uint64
	ForkID   uint64
	Version  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateSupernets2Version is a free log retrieval operation binding the contract event 0x13e60734641bd97ab6e444f419277d9983e2c3eab9bb631a578ffa07756b6351.
//
// Solidity: event UpdateSupernets2Version(uint64 numBatch, uint64 forkID, string version)
func (_Supernets2 *Supernets2Filterer) FilterUpdateSupernets2Version(opts *bind.FilterOpts) (*Supernets2UpdateSupernets2VersionIterator, error) {

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "UpdateSupernets2Version")
	if err != nil {
		return nil, err
	}
	return &Supernets2UpdateSupernets2VersionIterator{contract: _Supernets2.contract, event: "UpdateSupernets2Version", logs: logs, sub: sub}, nil
}

// WatchUpdateSupernets2Version is a free log subscription operation binding the contract event 0x13e60734641bd97ab6e444f419277d9983e2c3eab9bb631a578ffa07756b6351.
//
// Solidity: event UpdateSupernets2Version(uint64 numBatch, uint64 forkID, string version)
func (_Supernets2 *Supernets2Filterer) WatchUpdateSupernets2Version(opts *bind.WatchOpts, sink chan<- *Supernets2UpdateSupernets2Version) (event.Subscription, error) {

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "UpdateSupernets2Version")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2UpdateSupernets2Version)
				if err := _Supernets2.contract.UnpackLog(event, "UpdateSupernets2Version", log); err != nil {
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

// ParseUpdateSupernets2Version is a log parse operation binding the contract event 0x13e60734641bd97ab6e444f419277d9983e2c3eab9bb631a578ffa07756b6351.
//
// Solidity: event UpdateSupernets2Version(uint64 numBatch, uint64 forkID, string version)
func (_Supernets2 *Supernets2Filterer) ParseUpdateSupernets2Version(log types.Log) (*Supernets2UpdateSupernets2Version, error) {
	event := new(Supernets2UpdateSupernets2Version)
	if err := _Supernets2.contract.UnpackLog(event, "UpdateSupernets2Version", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2VerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Supernets2 contract.
type Supernets2VerifyBatchesIterator struct {
	Event *Supernets2VerifyBatches // Event containing the contract specifics and raw log

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
func (it *Supernets2VerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2VerifyBatches)
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
		it.Event = new(Supernets2VerifyBatches)
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
func (it *Supernets2VerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2VerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2VerifyBatches represents a VerifyBatches event raised by the Supernets2 contract.
type Supernets2VerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*Supernets2VerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2VerifyBatchesIterator{contract: _Supernets2.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *Supernets2VerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2VerifyBatches)
				if err := _Supernets2.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Supernets2 *Supernets2Filterer) ParseVerifyBatches(log types.Log) (*Supernets2VerifyBatches, error) {
	event := new(Supernets2VerifyBatches)
	if err := _Supernets2.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Supernets2VerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Supernets2 contract.
type Supernets2VerifyBatchesTrustedAggregatorIterator struct {
	Event *Supernets2VerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *Supernets2VerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Supernets2VerifyBatchesTrustedAggregator)
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
		it.Event = new(Supernets2VerifyBatchesTrustedAggregator)
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
func (it *Supernets2VerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Supernets2VerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Supernets2VerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Supernets2 contract.
type Supernets2VerifyBatchesTrustedAggregator struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatchesTrustedAggregator is a free log retrieval operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*Supernets2VerifyBatchesTrustedAggregatorIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Supernets2.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &Supernets2VerifyBatchesTrustedAggregatorIterator{contract: _Supernets2.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *Supernets2VerifyBatchesTrustedAggregator, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Supernets2.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Supernets2VerifyBatchesTrustedAggregator)
				if err := _Supernets2.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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

// ParseVerifyBatchesTrustedAggregator is a log parse operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Supernets2 *Supernets2Filterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*Supernets2VerifyBatchesTrustedAggregator, error) {
	event := new(Supernets2VerifyBatchesTrustedAggregator)
	if err := _Supernets2.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
