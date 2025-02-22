// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mportal

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

// INttManagerNttManagerPeer is an auto generated low-level Go binding around an user-defined struct.
type INttManagerNttManagerPeer struct {
	PeerAddress   [32]byte
	TokenDecimals uint8
}

// IRateLimiterInboundQueuedTransfer is an auto generated low-level Go binding around an user-defined struct.
type IRateLimiterInboundQueuedTransfer struct {
	Amount      *big.Int
	TxTimestamp uint64
	Recipient   common.Address
}

// IRateLimiterOutboundQueuedTransfer is an auto generated low-level Go binding around an user-defined struct.
type IRateLimiterOutboundQueuedTransfer struct {
	Recipient               [32]byte
	RefundAddress           [32]byte
	Amount                  *big.Int
	TxTimestamp             uint64
	RecipientChain          uint16
	Sender                  common.Address
	TransceiverInstructions []byte
}

// IRateLimiterRateLimitParams is an auto generated low-level Go binding around an user-defined struct.
type IRateLimiterRateLimitParams struct {
	Limit           *big.Int
	CurrentCapacity *big.Int
	LastTxTimestamp uint64
}

// TransceiverRegistryTransceiverInfo is an auto generated low-level Go binding around an user-defined struct.
type TransceiverRegistryTransceiverInfo struct {
	Registered bool
	Enabled    bool
	Index      uint8
}

// TransceiverStructsNttManagerMessage is an auto generated low-level Go binding around an user-defined struct.
type TransceiverStructsNttManagerMessage struct {
	Id      [32]byte
	Sender  [32]byte
	Payload []byte
}

// AbiMetaData contains all meta data concerning the Abi contract.
var AbiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"registrar_\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"chainId_\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceDiff\",\"type\":\"uint256\"}],\"name\":\"BurnAmountDifferentThanBalanceDiff\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerNotTransceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"CancellerNotSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"TrimmedAmount\",\"name\":\"newCurrentCapacity\",\"type\":\"uint72\"},{\"internalType\":\"TrimmedAmount\",\"name\":\"newLimit\",\"type\":\"uint72\"}],\"name\":\"CapacityCannotExceedLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"providedPayment\",\"type\":\"uint256\"}],\"name\":\"DeliveryPaymentTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transceiver\",\"type\":\"address\"}],\"name\":\"DisabledTransceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EarningCannotBeReenabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EarningIsDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EarningIsEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"InboundQueuedTransferNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"transferTimestamp\",\"type\":\"uint256\"}],\"name\":\"InboundQueuedTransferStillQueued\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"}],\"name\":\"InvalidDestinationChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"evmChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidFork\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"mode\",\"type\":\"uint8\"}],\"name\":\"InvalidMode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"InvalidPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"InvalidPayloadLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"prefix\",\"type\":\"bytes4\"}],\"name\":\"InvalidPayloadPrefix\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"peerAddress\",\"type\":\"bytes32\"}],\"name\":\"InvalidPeer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPeerChainIdZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPeerDecimals\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPeerSameChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPeerZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRefundAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"targetChain\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"thisChain\",\"type\":\"uint16\"}],\"name\":\"InvalidTargetChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTransceiverZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IsApprovedEarner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"encodedLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"MessageNotApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoEnabledTransceivers\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transceiver\",\"type\":\"address\"}],\"name\":\"NonRegisteredTransceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotApprovedEarner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentCapacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NotEnoughCapacity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplemented\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMigrating\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyDelegateCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"queueSequence\",\"type\":\"uint64\"}],\"name\":\"OutboundQueuedTransferNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"queueSequence\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"transferTimestamp\",\"type\":\"uint256\"}],\"name\":\"OutboundQueuedTransferStillQueued\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"}],\"name\":\"PeerNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"RefundFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequireContractIsNotPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequireContractIsPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"retrieved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"registered\",\"type\":\"uint256\"}],\"name\":\"RetrievedIncorrectRegisteredTransceivers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaticcallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transceivers\",\"type\":\"uint256\"}],\"name\":\"ThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyTransceivers\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nttManagerMessageHash\",\"type\":\"bytes32\"}],\"name\":\"TransceiverAlreadyAttestedToMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transceiver\",\"type\":\"address\"}],\"name\":\"TransceiverAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dust\",\"type\":\"uint256\"}],\"name\":\"TransferAmountHasDust\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Uint64Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UndefinedRateLimiting\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expectedOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UnexpectedDeployer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"destinationToken\",\"type\":\"bytes32\"}],\"name\":\"UnsupportedBridgingPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroDestinationToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroMToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroRegistrar\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroSourceToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroThreshold\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mToken\",\"type\":\"bytes32\"}],\"name\":\"DestinationMTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"index\",\"type\":\"uint128\"}],\"name\":\"EarningDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"index\",\"type\":\"uint128\"}],\"name\":\"EarningEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"InboundTransferLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"InboundTransferQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"index\",\"type\":\"uint128\"}],\"name\":\"MTokenIndexSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"index\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MTokenReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destinationToken\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"index\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MTokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceNttManager\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"MessageAlreadyExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"MessageAttestedTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"notPaused\",\"type\":\"bool\"}],\"name\":\"NotPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OutboundTransferCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"OutboundTransferLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"queueSequence\",\"type\":\"uint64\"}],\"name\":\"OutboundTransferQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentCapacity\",\"type\":\"uint256\"}],\"name\":\"OutboundTransferRateLimited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPauser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"}],\"name\":\"PauserTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"chainId_\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldPeerContract\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"oldPeerDecimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peerContract\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"peerDecimals\",\"type\":\"uint8\"}],\"name\":\"PeerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"RegistrarKeySent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"listName\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"RegistrarListStatusSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationToken\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"name\":\"SupportedBridgingPathSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"oldThreshold\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transceiversNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"TransceiverAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"TransceiverRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"TransferRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"refundAddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"recipientChain\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"msgSequence\",\"type\":\"uint64\"}],\"name\":\"TransferSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"TransferSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationWrappedToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WrapFailed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_TRANSCEIVER_INSTRUCTIONS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NTT_MANAGER_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"sourceNttManagerAddress\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structTransceiverStructs.NttManagerMessage\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"attestationReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"cancelOutboundQueuedTransfer\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"completeInboundQueuedTransfer\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"completeOutboundQueuedTransfer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentIndex\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"}],\"name\":\"destinationMToken\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"mToken\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableEarning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableEarningIndex\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableEarning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"sourceNttManagerAddress\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structTransceiverStructs.NttManagerMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"getCurrentInboundCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentOutboundCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"getInboundLimitParams\",\"outputs\":[{\"components\":[{\"internalType\":\"TrimmedAmount\",\"name\":\"limit\",\"type\":\"uint72\"},{\"internalType\":\"TrimmedAmount\",\"name\":\"currentCapacity\",\"type\":\"uint72\"},{\"internalType\":\"uint64\",\"name\":\"lastTxTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIRateLimiter.RateLimitParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getInboundQueuedTransfer\",\"outputs\":[{\"components\":[{\"internalType\":\"TrimmedAmount\",\"name\":\"amount\",\"type\":\"uint72\"},{\"internalType\":\"uint64\",\"name\":\"txTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIRateLimiter.InboundQueuedTransfer\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMigratesImmutables\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMode\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOutboundLimitParams\",\"outputs\":[{\"components\":[{\"internalType\":\"TrimmedAmount\",\"name\":\"limit\",\"type\":\"uint72\"},{\"internalType\":\"TrimmedAmount\",\"name\":\"currentCapacity\",\"type\":\"uint72\"},{\"internalType\":\"uint64\",\"name\":\"lastTxTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIRateLimiter.RateLimitParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getOutboundQueuedTransfer\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"refundAddress\",\"type\":\"bytes32\"},{\"internalType\":\"TrimmedAmount\",\"name\":\"amount\",\"type\":\"uint72\"},{\"internalType\":\"uint64\",\"name\":\"txTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"recipientChain\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"transceiverInstructions\",\"type\":\"bytes\"}],\"internalType\":\"structIRateLimiter.OutboundQueuedTransfer\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"chainId_\",\"type\":\"uint16\"}],\"name\":\"getPeer\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"peerAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structINttManager.NttManagerPeer\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransceiverInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"internalType\":\"structTransceiverRegistry.TransceiverInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransceivers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isMessageApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isMessageExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"messageAttestations\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mode\",\"outputs\":[{\"internalType\":\"enumIManagerBase.Mode\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextMessageSequence\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"recipientChain\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"transceiverInstructions\",\"type\":\"bytes\"}],\"name\":\"quoteDeliveryPrice\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateLimitDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrar\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transceiver\",\"type\":\"address\"}],\"name\":\"removeTransceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"destinationChainId_\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"refundAddress_\",\"type\":\"bytes32\"}],\"name\":\"sendMTokenIndex\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId_\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"destinationChainId_\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"refundAddress_\",\"type\":\"bytes32\"}],\"name\":\"sendRegistrarKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId_\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"destinationChainId_\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"listName_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"refundAddress_\",\"type\":\"bytes32\"}],\"name\":\"sendRegistrarListStatus\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId_\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"destinationChainId_\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"mToken_\",\"type\":\"bytes32\"}],\"name\":\"setDestinationMToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"chainId_\",\"type\":\"uint16\"}],\"name\":\"setInboundLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setOutboundLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"peerChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"peerContract\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"inboundLimit\",\"type\":\"uint256\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken_\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId_\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"destinationToken_\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"supported_\",\"type\":\"bool\"}],\"name\":\"setSupportedBridgingPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transceiver\",\"type\":\"address\"}],\"name\":\"setTransceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"destinationToken\",\"type\":\"bytes32\"}],\"name\":\"supportedBridgingPath\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"transceiverAttestedToMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"recipientChain\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"recipientChain\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"refundAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"shouldQueue\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"transceiverInstructions\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sourceToken_\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId_\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"destinationToken_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"refundAddress_\",\"type\":\"bytes32\"}],\"name\":\"transferMLikeToken\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequence_\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"}],\"name\":\"transferPauserCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wasEarningEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AbiABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiMetaData.ABI instead.
var AbiABI = AbiMetaData.ABI

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTTRANSCEIVERINSTRUCTIONS is a free data retrieval call binding the contract method 0x74707407.
//
// Solidity: function DEFAULT_TRANSCEIVER_INSTRUCTIONS() view returns(bytes)
func (_Abi *AbiCaller) DEFAULTTRANSCEIVERINSTRUCTIONS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "DEFAULT_TRANSCEIVER_INSTRUCTIONS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DEFAULTTRANSCEIVERINSTRUCTIONS is a free data retrieval call binding the contract method 0x74707407.
//
// Solidity: function DEFAULT_TRANSCEIVER_INSTRUCTIONS() view returns(bytes)
func (_Abi *AbiSession) DEFAULTTRANSCEIVERINSTRUCTIONS() ([]byte, error) {
	return _Abi.Contract.DEFAULTTRANSCEIVERINSTRUCTIONS(&_Abi.CallOpts)
}

// DEFAULTTRANSCEIVERINSTRUCTIONS is a free data retrieval call binding the contract method 0x74707407.
//
// Solidity: function DEFAULT_TRANSCEIVER_INSTRUCTIONS() view returns(bytes)
func (_Abi *AbiCallerSession) DEFAULTTRANSCEIVERINSTRUCTIONS() ([]byte, error) {
	return _Abi.Contract.DEFAULTTRANSCEIVERINSTRUCTIONS(&_Abi.CallOpts)
}

// NTTMANAGERVERSION is a free data retrieval call binding the contract method 0xc0b07bde.
//
// Solidity: function NTT_MANAGER_VERSION() view returns(string)
func (_Abi *AbiCaller) NTTMANAGERVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "NTT_MANAGER_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NTTMANAGERVERSION is a free data retrieval call binding the contract method 0xc0b07bde.
//
// Solidity: function NTT_MANAGER_VERSION() view returns(string)
func (_Abi *AbiSession) NTTMANAGERVERSION() (string, error) {
	return _Abi.Contract.NTTMANAGERVERSION(&_Abi.CallOpts)
}

// NTTMANAGERVERSION is a free data retrieval call binding the contract method 0xc0b07bde.
//
// Solidity: function NTT_MANAGER_VERSION() view returns(string)
func (_Abi *AbiCallerSession) NTTMANAGERVERSION() (string, error) {
	return _Abi.Contract.NTTMANAGERVERSION(&_Abi.CallOpts)
}

// CancelOutboundQueuedTransfer is a free data retrieval call binding the contract method 0xf7514fbc.
//
// Solidity: function cancelOutboundQueuedTransfer(uint64 ) view returns()
func (_Abi *AbiCaller) CancelOutboundQueuedTransfer(opts *bind.CallOpts, arg0 uint64) error {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "cancelOutboundQueuedTransfer", arg0)

	if err != nil {
		return err
	}

	return err

}

// CancelOutboundQueuedTransfer is a free data retrieval call binding the contract method 0xf7514fbc.
//
// Solidity: function cancelOutboundQueuedTransfer(uint64 ) view returns()
func (_Abi *AbiSession) CancelOutboundQueuedTransfer(arg0 uint64) error {
	return _Abi.Contract.CancelOutboundQueuedTransfer(&_Abi.CallOpts, arg0)
}

// CancelOutboundQueuedTransfer is a free data retrieval call binding the contract method 0xf7514fbc.
//
// Solidity: function cancelOutboundQueuedTransfer(uint64 ) view returns()
func (_Abi *AbiCallerSession) CancelOutboundQueuedTransfer(arg0 uint64) error {
	return _Abi.Contract.CancelOutboundQueuedTransfer(&_Abi.CallOpts, arg0)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Abi *AbiCaller) ChainId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Abi *AbiSession) ChainId() (uint16, error) {
	return _Abi.Contract.ChainId(&_Abi.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Abi *AbiCallerSession) ChainId() (uint16, error) {
	return _Abi.Contract.ChainId(&_Abi.CallOpts)
}

// CompleteInboundQueuedTransfer is a free data retrieval call binding the contract method 0x8413bcba.
//
// Solidity: function completeInboundQueuedTransfer(bytes32 ) view returns()
func (_Abi *AbiCaller) CompleteInboundQueuedTransfer(opts *bind.CallOpts, arg0 [32]byte) error {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "completeInboundQueuedTransfer", arg0)

	if err != nil {
		return err
	}

	return err

}

// CompleteInboundQueuedTransfer is a free data retrieval call binding the contract method 0x8413bcba.
//
// Solidity: function completeInboundQueuedTransfer(bytes32 ) view returns()
func (_Abi *AbiSession) CompleteInboundQueuedTransfer(arg0 [32]byte) error {
	return _Abi.Contract.CompleteInboundQueuedTransfer(&_Abi.CallOpts, arg0)
}

// CompleteInboundQueuedTransfer is a free data retrieval call binding the contract method 0x8413bcba.
//
// Solidity: function completeInboundQueuedTransfer(bytes32 ) view returns()
func (_Abi *AbiCallerSession) CompleteInboundQueuedTransfer(arg0 [32]byte) error {
	return _Abi.Contract.CompleteInboundQueuedTransfer(&_Abi.CallOpts, arg0)
}

// CurrentIndex is a free data retrieval call binding the contract method 0x26987b60.
//
// Solidity: function currentIndex() view returns(uint128)
func (_Abi *AbiCaller) CurrentIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "currentIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentIndex is a free data retrieval call binding the contract method 0x26987b60.
//
// Solidity: function currentIndex() view returns(uint128)
func (_Abi *AbiSession) CurrentIndex() (*big.Int, error) {
	return _Abi.Contract.CurrentIndex(&_Abi.CallOpts)
}

// CurrentIndex is a free data retrieval call binding the contract method 0x26987b60.
//
// Solidity: function currentIndex() view returns(uint128)
func (_Abi *AbiCallerSession) CurrentIndex() (*big.Int, error) {
	return _Abi.Contract.CurrentIndex(&_Abi.CallOpts)
}

// DestinationMToken is a free data retrieval call binding the contract method 0xab5d44d1.
//
// Solidity: function destinationMToken(uint16 destinationChainId) view returns(bytes32 mToken)
func (_Abi *AbiCaller) DestinationMToken(opts *bind.CallOpts, destinationChainId uint16) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "destinationMToken", destinationChainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DestinationMToken is a free data retrieval call binding the contract method 0xab5d44d1.
//
// Solidity: function destinationMToken(uint16 destinationChainId) view returns(bytes32 mToken)
func (_Abi *AbiSession) DestinationMToken(destinationChainId uint16) ([32]byte, error) {
	return _Abi.Contract.DestinationMToken(&_Abi.CallOpts, destinationChainId)
}

// DestinationMToken is a free data retrieval call binding the contract method 0xab5d44d1.
//
// Solidity: function destinationMToken(uint16 destinationChainId) view returns(bytes32 mToken)
func (_Abi *AbiCallerSession) DestinationMToken(destinationChainId uint16) ([32]byte, error) {
	return _Abi.Contract.DestinationMToken(&_Abi.CallOpts, destinationChainId)
}

// DisableEarningIndex is a free data retrieval call binding the contract method 0xf22e63ac.
//
// Solidity: function disableEarningIndex() view returns(uint128)
func (_Abi *AbiCaller) DisableEarningIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "disableEarningIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisableEarningIndex is a free data retrieval call binding the contract method 0xf22e63ac.
//
// Solidity: function disableEarningIndex() view returns(uint128)
func (_Abi *AbiSession) DisableEarningIndex() (*big.Int, error) {
	return _Abi.Contract.DisableEarningIndex(&_Abi.CallOpts)
}

// DisableEarningIndex is a free data retrieval call binding the contract method 0xf22e63ac.
//
// Solidity: function disableEarningIndex() view returns(uint128)
func (_Abi *AbiCallerSession) DisableEarningIndex() (*big.Int, error) {
	return _Abi.Contract.DisableEarningIndex(&_Abi.CallOpts)
}

// GetCurrentInboundCapacity is a free data retrieval call binding the contract method 0x02717250.
//
// Solidity: function getCurrentInboundCapacity(uint16 ) pure returns(uint256)
func (_Abi *AbiCaller) GetCurrentInboundCapacity(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getCurrentInboundCapacity", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentInboundCapacity is a free data retrieval call binding the contract method 0x02717250.
//
// Solidity: function getCurrentInboundCapacity(uint16 ) pure returns(uint256)
func (_Abi *AbiSession) GetCurrentInboundCapacity(arg0 uint16) (*big.Int, error) {
	return _Abi.Contract.GetCurrentInboundCapacity(&_Abi.CallOpts, arg0)
}

// GetCurrentInboundCapacity is a free data retrieval call binding the contract method 0x02717250.
//
// Solidity: function getCurrentInboundCapacity(uint16 ) pure returns(uint256)
func (_Abi *AbiCallerSession) GetCurrentInboundCapacity(arg0 uint16) (*big.Int, error) {
	return _Abi.Contract.GetCurrentInboundCapacity(&_Abi.CallOpts, arg0)
}

// GetCurrentOutboundCapacity is a free data retrieval call binding the contract method 0xf5cfec18.
//
// Solidity: function getCurrentOutboundCapacity() pure returns(uint256)
func (_Abi *AbiCaller) GetCurrentOutboundCapacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getCurrentOutboundCapacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentOutboundCapacity is a free data retrieval call binding the contract method 0xf5cfec18.
//
// Solidity: function getCurrentOutboundCapacity() pure returns(uint256)
func (_Abi *AbiSession) GetCurrentOutboundCapacity() (*big.Int, error) {
	return _Abi.Contract.GetCurrentOutboundCapacity(&_Abi.CallOpts)
}

// GetCurrentOutboundCapacity is a free data retrieval call binding the contract method 0xf5cfec18.
//
// Solidity: function getCurrentOutboundCapacity() pure returns(uint256)
func (_Abi *AbiCallerSession) GetCurrentOutboundCapacity() (*big.Int, error) {
	return _Abi.Contract.GetCurrentOutboundCapacity(&_Abi.CallOpts)
}

// GetInboundLimitParams is a free data retrieval call binding the contract method 0xd788c147.
//
// Solidity: function getInboundLimitParams(uint16 ) pure returns((uint72,uint72,uint64))
func (_Abi *AbiCaller) GetInboundLimitParams(opts *bind.CallOpts, arg0 uint16) (IRateLimiterRateLimitParams, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getInboundLimitParams", arg0)

	if err != nil {
		return *new(IRateLimiterRateLimitParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IRateLimiterRateLimitParams)).(*IRateLimiterRateLimitParams)

	return out0, err

}

// GetInboundLimitParams is a free data retrieval call binding the contract method 0xd788c147.
//
// Solidity: function getInboundLimitParams(uint16 ) pure returns((uint72,uint72,uint64))
func (_Abi *AbiSession) GetInboundLimitParams(arg0 uint16) (IRateLimiterRateLimitParams, error) {
	return _Abi.Contract.GetInboundLimitParams(&_Abi.CallOpts, arg0)
}

// GetInboundLimitParams is a free data retrieval call binding the contract method 0xd788c147.
//
// Solidity: function getInboundLimitParams(uint16 ) pure returns((uint72,uint72,uint64))
func (_Abi *AbiCallerSession) GetInboundLimitParams(arg0 uint16) (IRateLimiterRateLimitParams, error) {
	return _Abi.Contract.GetInboundLimitParams(&_Abi.CallOpts, arg0)
}

// GetInboundQueuedTransfer is a free data retrieval call binding the contract method 0xfd96063c.
//
// Solidity: function getInboundQueuedTransfer(bytes32 ) pure returns((uint72,uint64,address))
func (_Abi *AbiCaller) GetInboundQueuedTransfer(opts *bind.CallOpts, arg0 [32]byte) (IRateLimiterInboundQueuedTransfer, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getInboundQueuedTransfer", arg0)

	if err != nil {
		return *new(IRateLimiterInboundQueuedTransfer), err
	}

	out0 := *abi.ConvertType(out[0], new(IRateLimiterInboundQueuedTransfer)).(*IRateLimiterInboundQueuedTransfer)

	return out0, err

}

// GetInboundQueuedTransfer is a free data retrieval call binding the contract method 0xfd96063c.
//
// Solidity: function getInboundQueuedTransfer(bytes32 ) pure returns((uint72,uint64,address))
func (_Abi *AbiSession) GetInboundQueuedTransfer(arg0 [32]byte) (IRateLimiterInboundQueuedTransfer, error) {
	return _Abi.Contract.GetInboundQueuedTransfer(&_Abi.CallOpts, arg0)
}

// GetInboundQueuedTransfer is a free data retrieval call binding the contract method 0xfd96063c.
//
// Solidity: function getInboundQueuedTransfer(bytes32 ) pure returns((uint72,uint64,address))
func (_Abi *AbiCallerSession) GetInboundQueuedTransfer(arg0 [32]byte) (IRateLimiterInboundQueuedTransfer, error) {
	return _Abi.Contract.GetInboundQueuedTransfer(&_Abi.CallOpts, arg0)
}

// GetMigratesImmutables is a free data retrieval call binding the contract method 0x689f90c3.
//
// Solidity: function getMigratesImmutables() view returns(bool)
func (_Abi *AbiCaller) GetMigratesImmutables(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getMigratesImmutables")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMigratesImmutables is a free data retrieval call binding the contract method 0x689f90c3.
//
// Solidity: function getMigratesImmutables() view returns(bool)
func (_Abi *AbiSession) GetMigratesImmutables() (bool, error) {
	return _Abi.Contract.GetMigratesImmutables(&_Abi.CallOpts)
}

// GetMigratesImmutables is a free data retrieval call binding the contract method 0x689f90c3.
//
// Solidity: function getMigratesImmutables() view returns(bool)
func (_Abi *AbiCallerSession) GetMigratesImmutables() (bool, error) {
	return _Abi.Contract.GetMigratesImmutables(&_Abi.CallOpts)
}

// GetMode is a free data retrieval call binding the contract method 0x4b4fd03b.
//
// Solidity: function getMode() view returns(uint8)
func (_Abi *AbiCaller) GetMode(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getMode")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetMode is a free data retrieval call binding the contract method 0x4b4fd03b.
//
// Solidity: function getMode() view returns(uint8)
func (_Abi *AbiSession) GetMode() (uint8, error) {
	return _Abi.Contract.GetMode(&_Abi.CallOpts)
}

// GetMode is a free data retrieval call binding the contract method 0x4b4fd03b.
//
// Solidity: function getMode() view returns(uint8)
func (_Abi *AbiCallerSession) GetMode() (uint8, error) {
	return _Abi.Contract.GetMode(&_Abi.CallOpts)
}

// GetOutboundLimitParams is a free data retrieval call binding the contract method 0x86e11ffa.
//
// Solidity: function getOutboundLimitParams() pure returns((uint72,uint72,uint64))
func (_Abi *AbiCaller) GetOutboundLimitParams(opts *bind.CallOpts) (IRateLimiterRateLimitParams, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getOutboundLimitParams")

	if err != nil {
		return *new(IRateLimiterRateLimitParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IRateLimiterRateLimitParams)).(*IRateLimiterRateLimitParams)

	return out0, err

}

// GetOutboundLimitParams is a free data retrieval call binding the contract method 0x86e11ffa.
//
// Solidity: function getOutboundLimitParams() pure returns((uint72,uint72,uint64))
func (_Abi *AbiSession) GetOutboundLimitParams() (IRateLimiterRateLimitParams, error) {
	return _Abi.Contract.GetOutboundLimitParams(&_Abi.CallOpts)
}

// GetOutboundLimitParams is a free data retrieval call binding the contract method 0x86e11ffa.
//
// Solidity: function getOutboundLimitParams() pure returns((uint72,uint72,uint64))
func (_Abi *AbiCallerSession) GetOutboundLimitParams() (IRateLimiterRateLimitParams, error) {
	return _Abi.Contract.GetOutboundLimitParams(&_Abi.CallOpts)
}

// GetOutboundQueuedTransfer is a free data retrieval call binding the contract method 0x1f97c9a8.
//
// Solidity: function getOutboundQueuedTransfer(uint64 ) pure returns((bytes32,bytes32,uint72,uint64,uint16,address,bytes))
func (_Abi *AbiCaller) GetOutboundQueuedTransfer(opts *bind.CallOpts, arg0 uint64) (IRateLimiterOutboundQueuedTransfer, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getOutboundQueuedTransfer", arg0)

	if err != nil {
		return *new(IRateLimiterOutboundQueuedTransfer), err
	}

	out0 := *abi.ConvertType(out[0], new(IRateLimiterOutboundQueuedTransfer)).(*IRateLimiterOutboundQueuedTransfer)

	return out0, err

}

// GetOutboundQueuedTransfer is a free data retrieval call binding the contract method 0x1f97c9a8.
//
// Solidity: function getOutboundQueuedTransfer(uint64 ) pure returns((bytes32,bytes32,uint72,uint64,uint16,address,bytes))
func (_Abi *AbiSession) GetOutboundQueuedTransfer(arg0 uint64) (IRateLimiterOutboundQueuedTransfer, error) {
	return _Abi.Contract.GetOutboundQueuedTransfer(&_Abi.CallOpts, arg0)
}

// GetOutboundQueuedTransfer is a free data retrieval call binding the contract method 0x1f97c9a8.
//
// Solidity: function getOutboundQueuedTransfer(uint64 ) pure returns((bytes32,bytes32,uint72,uint64,uint16,address,bytes))
func (_Abi *AbiCallerSession) GetOutboundQueuedTransfer(arg0 uint64) (IRateLimiterOutboundQueuedTransfer, error) {
	return _Abi.Contract.GetOutboundQueuedTransfer(&_Abi.CallOpts, arg0)
}

// GetPeer is a free data retrieval call binding the contract method 0xc128d170.
//
// Solidity: function getPeer(uint16 chainId_) view returns((bytes32,uint8))
func (_Abi *AbiCaller) GetPeer(opts *bind.CallOpts, chainId_ uint16) (INttManagerNttManagerPeer, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getPeer", chainId_)

	if err != nil {
		return *new(INttManagerNttManagerPeer), err
	}

	out0 := *abi.ConvertType(out[0], new(INttManagerNttManagerPeer)).(*INttManagerNttManagerPeer)

	return out0, err

}

// GetPeer is a free data retrieval call binding the contract method 0xc128d170.
//
// Solidity: function getPeer(uint16 chainId_) view returns((bytes32,uint8))
func (_Abi *AbiSession) GetPeer(chainId_ uint16) (INttManagerNttManagerPeer, error) {
	return _Abi.Contract.GetPeer(&_Abi.CallOpts, chainId_)
}

// GetPeer is a free data retrieval call binding the contract method 0xc128d170.
//
// Solidity: function getPeer(uint16 chainId_) view returns((bytes32,uint8))
func (_Abi *AbiCallerSession) GetPeer(chainId_ uint16) (INttManagerNttManagerPeer, error) {
	return _Abi.Contract.GetPeer(&_Abi.CallOpts, chainId_)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint8)
func (_Abi *AbiCaller) GetThreshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint8)
func (_Abi *AbiSession) GetThreshold() (uint8, error) {
	return _Abi.Contract.GetThreshold(&_Abi.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint8)
func (_Abi *AbiCallerSession) GetThreshold() (uint8, error) {
	return _Abi.Contract.GetThreshold(&_Abi.CallOpts)
}

// GetTransceiverInfo is a free data retrieval call binding the contract method 0xb150fc55.
//
// Solidity: function getTransceiverInfo() view returns((bool,bool,uint8)[])
func (_Abi *AbiCaller) GetTransceiverInfo(opts *bind.CallOpts) ([]TransceiverRegistryTransceiverInfo, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getTransceiverInfo")

	if err != nil {
		return *new([]TransceiverRegistryTransceiverInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransceiverRegistryTransceiverInfo)).(*[]TransceiverRegistryTransceiverInfo)

	return out0, err

}

// GetTransceiverInfo is a free data retrieval call binding the contract method 0xb150fc55.
//
// Solidity: function getTransceiverInfo() view returns((bool,bool,uint8)[])
func (_Abi *AbiSession) GetTransceiverInfo() ([]TransceiverRegistryTransceiverInfo, error) {
	return _Abi.Contract.GetTransceiverInfo(&_Abi.CallOpts)
}

// GetTransceiverInfo is a free data retrieval call binding the contract method 0xb150fc55.
//
// Solidity: function getTransceiverInfo() view returns((bool,bool,uint8)[])
func (_Abi *AbiCallerSession) GetTransceiverInfo() ([]TransceiverRegistryTransceiverInfo, error) {
	return _Abi.Contract.GetTransceiverInfo(&_Abi.CallOpts)
}

// GetTransceivers is a free data retrieval call binding the contract method 0xb4d591bb.
//
// Solidity: function getTransceivers() pure returns(address[] result)
func (_Abi *AbiCaller) GetTransceivers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getTransceivers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTransceivers is a free data retrieval call binding the contract method 0xb4d591bb.
//
// Solidity: function getTransceivers() pure returns(address[] result)
func (_Abi *AbiSession) GetTransceivers() ([]common.Address, error) {
	return _Abi.Contract.GetTransceivers(&_Abi.CallOpts)
}

// GetTransceivers is a free data retrieval call binding the contract method 0xb4d591bb.
//
// Solidity: function getTransceivers() pure returns(address[] result)
func (_Abi *AbiCallerSession) GetTransceivers() ([]common.Address, error) {
	return _Abi.Contract.GetTransceivers(&_Abi.CallOpts)
}

// IsMessageApproved is a free data retrieval call binding the contract method 0x0677df54.
//
// Solidity: function isMessageApproved(bytes32 digest) view returns(bool)
func (_Abi *AbiCaller) IsMessageApproved(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "isMessageApproved", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessageApproved is a free data retrieval call binding the contract method 0x0677df54.
//
// Solidity: function isMessageApproved(bytes32 digest) view returns(bool)
func (_Abi *AbiSession) IsMessageApproved(digest [32]byte) (bool, error) {
	return _Abi.Contract.IsMessageApproved(&_Abi.CallOpts, digest)
}

// IsMessageApproved is a free data retrieval call binding the contract method 0x0677df54.
//
// Solidity: function isMessageApproved(bytes32 digest) view returns(bool)
func (_Abi *AbiCallerSession) IsMessageApproved(digest [32]byte) (bool, error) {
	return _Abi.Contract.IsMessageApproved(&_Abi.CallOpts, digest)
}

// IsMessageExecuted is a free data retrieval call binding the contract method 0x396c16b7.
//
// Solidity: function isMessageExecuted(bytes32 digest) view returns(bool)
func (_Abi *AbiCaller) IsMessageExecuted(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "isMessageExecuted", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessageExecuted is a free data retrieval call binding the contract method 0x396c16b7.
//
// Solidity: function isMessageExecuted(bytes32 digest) view returns(bool)
func (_Abi *AbiSession) IsMessageExecuted(digest [32]byte) (bool, error) {
	return _Abi.Contract.IsMessageExecuted(&_Abi.CallOpts, digest)
}

// IsMessageExecuted is a free data retrieval call binding the contract method 0x396c16b7.
//
// Solidity: function isMessageExecuted(bytes32 digest) view returns(bool)
func (_Abi *AbiCallerSession) IsMessageExecuted(digest [32]byte) (bool, error) {
	return _Abi.Contract.IsMessageExecuted(&_Abi.CallOpts, digest)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Abi *AbiCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Abi *AbiSession) IsPaused() (bool, error) {
	return _Abi.Contract.IsPaused(&_Abi.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Abi *AbiCallerSession) IsPaused() (bool, error) {
	return _Abi.Contract.IsPaused(&_Abi.CallOpts)
}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() view returns(address)
func (_Abi *AbiCaller) MToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "mToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() view returns(address)
func (_Abi *AbiSession) MToken() (common.Address, error) {
	return _Abi.Contract.MToken(&_Abi.CallOpts)
}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() view returns(address)
func (_Abi *AbiCallerSession) MToken() (common.Address, error) {
	return _Abi.Contract.MToken(&_Abi.CallOpts)
}

// MessageAttestations is a free data retrieval call binding the contract method 0x89c619dd.
//
// Solidity: function messageAttestations(bytes32 digest) view returns(uint8 count)
func (_Abi *AbiCaller) MessageAttestations(opts *bind.CallOpts, digest [32]byte) (uint8, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "messageAttestations", digest)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MessageAttestations is a free data retrieval call binding the contract method 0x89c619dd.
//
// Solidity: function messageAttestations(bytes32 digest) view returns(uint8 count)
func (_Abi *AbiSession) MessageAttestations(digest [32]byte) (uint8, error) {
	return _Abi.Contract.MessageAttestations(&_Abi.CallOpts, digest)
}

// MessageAttestations is a free data retrieval call binding the contract method 0x89c619dd.
//
// Solidity: function messageAttestations(bytes32 digest) view returns(uint8 count)
func (_Abi *AbiCallerSession) MessageAttestations(digest [32]byte) (uint8, error) {
	return _Abi.Contract.MessageAttestations(&_Abi.CallOpts, digest)
}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(uint8)
func (_Abi *AbiCaller) Mode(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "mode")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(uint8)
func (_Abi *AbiSession) Mode() (uint8, error) {
	return _Abi.Contract.Mode(&_Abi.CallOpts)
}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(uint8)
func (_Abi *AbiCallerSession) Mode() (uint8, error) {
	return _Abi.Contract.Mode(&_Abi.CallOpts)
}

// NextMessageSequence is a free data retrieval call binding the contract method 0x23d75e31.
//
// Solidity: function nextMessageSequence() view returns(uint64)
func (_Abi *AbiCaller) NextMessageSequence(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "nextMessageSequence")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextMessageSequence is a free data retrieval call binding the contract method 0x23d75e31.
//
// Solidity: function nextMessageSequence() view returns(uint64)
func (_Abi *AbiSession) NextMessageSequence() (uint64, error) {
	return _Abi.Contract.NextMessageSequence(&_Abi.CallOpts)
}

// NextMessageSequence is a free data retrieval call binding the contract method 0x23d75e31.
//
// Solidity: function nextMessageSequence() view returns(uint64)
func (_Abi *AbiCallerSession) NextMessageSequence() (uint64, error) {
	return _Abi.Contract.NextMessageSequence(&_Abi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiSession) Owner() (common.Address, error) {
	return _Abi.Contract.Owner(&_Abi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiCallerSession) Owner() (common.Address, error) {
	return _Abi.Contract.Owner(&_Abi.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Abi *AbiCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Abi *AbiSession) Pauser() (common.Address, error) {
	return _Abi.Contract.Pauser(&_Abi.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Abi *AbiCallerSession) Pauser() (common.Address, error) {
	return _Abi.Contract.Pauser(&_Abi.CallOpts)
}

// QuoteDeliveryPrice is a free data retrieval call binding the contract method 0x9057412d.
//
// Solidity: function quoteDeliveryPrice(uint16 recipientChain, bytes transceiverInstructions) view returns(uint256[], uint256)
func (_Abi *AbiCaller) QuoteDeliveryPrice(opts *bind.CallOpts, recipientChain uint16, transceiverInstructions []byte) ([]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "quoteDeliveryPrice", recipientChain, transceiverInstructions)

	if err != nil {
		return *new([]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// QuoteDeliveryPrice is a free data retrieval call binding the contract method 0x9057412d.
//
// Solidity: function quoteDeliveryPrice(uint16 recipientChain, bytes transceiverInstructions) view returns(uint256[], uint256)
func (_Abi *AbiSession) QuoteDeliveryPrice(recipientChain uint16, transceiverInstructions []byte) ([]*big.Int, *big.Int, error) {
	return _Abi.Contract.QuoteDeliveryPrice(&_Abi.CallOpts, recipientChain, transceiverInstructions)
}

// QuoteDeliveryPrice is a free data retrieval call binding the contract method 0x9057412d.
//
// Solidity: function quoteDeliveryPrice(uint16 recipientChain, bytes transceiverInstructions) view returns(uint256[], uint256)
func (_Abi *AbiCallerSession) QuoteDeliveryPrice(recipientChain uint16, transceiverInstructions []byte) ([]*big.Int, *big.Int, error) {
	return _Abi.Contract.QuoteDeliveryPrice(&_Abi.CallOpts, recipientChain, transceiverInstructions)
}

// RateLimitDuration is a free data retrieval call binding the contract method 0x74aa7bfc.
//
// Solidity: function rateLimitDuration() view returns(uint64)
func (_Abi *AbiCaller) RateLimitDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "rateLimitDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RateLimitDuration is a free data retrieval call binding the contract method 0x74aa7bfc.
//
// Solidity: function rateLimitDuration() view returns(uint64)
func (_Abi *AbiSession) RateLimitDuration() (uint64, error) {
	return _Abi.Contract.RateLimitDuration(&_Abi.CallOpts)
}

// RateLimitDuration is a free data retrieval call binding the contract method 0x74aa7bfc.
//
// Solidity: function rateLimitDuration() view returns(uint64)
func (_Abi *AbiCallerSession) RateLimitDuration() (uint64, error) {
	return _Abi.Contract.RateLimitDuration(&_Abi.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_Abi *AbiCaller) Registrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "registrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_Abi *AbiSession) Registrar() (common.Address, error) {
	return _Abi.Contract.Registrar(&_Abi.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_Abi *AbiCallerSession) Registrar() (common.Address, error) {
	return _Abi.Contract.Registrar(&_Abi.CallOpts)
}

// SupportedBridgingPath is a free data retrieval call binding the contract method 0xcd9d5694.
//
// Solidity: function supportedBridgingPath(address sourceToken, uint16 destinationChainId, bytes32 destinationToken) view returns(bool supported)
func (_Abi *AbiCaller) SupportedBridgingPath(opts *bind.CallOpts, sourceToken common.Address, destinationChainId uint16, destinationToken [32]byte) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "supportedBridgingPath", sourceToken, destinationChainId, destinationToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportedBridgingPath is a free data retrieval call binding the contract method 0xcd9d5694.
//
// Solidity: function supportedBridgingPath(address sourceToken, uint16 destinationChainId, bytes32 destinationToken) view returns(bool supported)
func (_Abi *AbiSession) SupportedBridgingPath(sourceToken common.Address, destinationChainId uint16, destinationToken [32]byte) (bool, error) {
	return _Abi.Contract.SupportedBridgingPath(&_Abi.CallOpts, sourceToken, destinationChainId, destinationToken)
}

// SupportedBridgingPath is a free data retrieval call binding the contract method 0xcd9d5694.
//
// Solidity: function supportedBridgingPath(address sourceToken, uint16 destinationChainId, bytes32 destinationToken) view returns(bool supported)
func (_Abi *AbiCallerSession) SupportedBridgingPath(sourceToken common.Address, destinationChainId uint16, destinationToken [32]byte) (bool, error) {
	return _Abi.Contract.SupportedBridgingPath(&_Abi.CallOpts, sourceToken, destinationChainId, destinationToken)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Abi *AbiCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Abi *AbiSession) Token() (common.Address, error) {
	return _Abi.Contract.Token(&_Abi.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Abi *AbiCallerSession) Token() (common.Address, error) {
	return _Abi.Contract.Token(&_Abi.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_Abi *AbiCaller) TokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "tokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_Abi *AbiSession) TokenDecimals() (uint8, error) {
	return _Abi.Contract.TokenDecimals(&_Abi.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_Abi *AbiCallerSession) TokenDecimals() (uint8, error) {
	return _Abi.Contract.TokenDecimals(&_Abi.CallOpts)
}

// TransceiverAttestedToMessage is a free data retrieval call binding the contract method 0x8e3ba8c9.
//
// Solidity: function transceiverAttestedToMessage(bytes32 digest, uint8 index) view returns(bool)
func (_Abi *AbiCaller) TransceiverAttestedToMessage(opts *bind.CallOpts, digest [32]byte, index uint8) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "transceiverAttestedToMessage", digest, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransceiverAttestedToMessage is a free data retrieval call binding the contract method 0x8e3ba8c9.
//
// Solidity: function transceiverAttestedToMessage(bytes32 digest, uint8 index) view returns(bool)
func (_Abi *AbiSession) TransceiverAttestedToMessage(digest [32]byte, index uint8) (bool, error) {
	return _Abi.Contract.TransceiverAttestedToMessage(&_Abi.CallOpts, digest, index)
}

// TransceiverAttestedToMessage is a free data retrieval call binding the contract method 0x8e3ba8c9.
//
// Solidity: function transceiverAttestedToMessage(bytes32 digest, uint8 index) view returns(bool)
func (_Abi *AbiCallerSession) TransceiverAttestedToMessage(digest [32]byte, index uint8) (bool, error) {
	return _Abi.Contract.TransceiverAttestedToMessage(&_Abi.CallOpts, digest, index)
}

// WasEarningEnabled is a free data retrieval call binding the contract method 0x8a4a1017.
//
// Solidity: function wasEarningEnabled() view returns(bool)
func (_Abi *AbiCaller) WasEarningEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "wasEarningEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasEarningEnabled is a free data retrieval call binding the contract method 0x8a4a1017.
//
// Solidity: function wasEarningEnabled() view returns(bool)
func (_Abi *AbiSession) WasEarningEnabled() (bool, error) {
	return _Abi.Contract.WasEarningEnabled(&_Abi.CallOpts)
}

// WasEarningEnabled is a free data retrieval call binding the contract method 0x8a4a1017.
//
// Solidity: function wasEarningEnabled() view returns(bool)
func (_Abi *AbiCallerSession) WasEarningEnabled() (bool, error) {
	return _Abi.Contract.WasEarningEnabled(&_Abi.CallOpts)
}

// AttestationReceived is a paid mutator transaction binding the contract method 0x9d782454.
//
// Solidity: function attestationReceived(uint16 sourceChainId, bytes32 sourceNttManagerAddress, (bytes32,bytes32,bytes) payload) returns()
func (_Abi *AbiTransactor) AttestationReceived(opts *bind.TransactOpts, sourceChainId uint16, sourceNttManagerAddress [32]byte, payload TransceiverStructsNttManagerMessage) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "attestationReceived", sourceChainId, sourceNttManagerAddress, payload)
}

// AttestationReceived is a paid mutator transaction binding the contract method 0x9d782454.
//
// Solidity: function attestationReceived(uint16 sourceChainId, bytes32 sourceNttManagerAddress, (bytes32,bytes32,bytes) payload) returns()
func (_Abi *AbiSession) AttestationReceived(sourceChainId uint16, sourceNttManagerAddress [32]byte, payload TransceiverStructsNttManagerMessage) (*types.Transaction, error) {
	return _Abi.Contract.AttestationReceived(&_Abi.TransactOpts, sourceChainId, sourceNttManagerAddress, payload)
}

// AttestationReceived is a paid mutator transaction binding the contract method 0x9d782454.
//
// Solidity: function attestationReceived(uint16 sourceChainId, bytes32 sourceNttManagerAddress, (bytes32,bytes32,bytes) payload) returns()
func (_Abi *AbiTransactorSession) AttestationReceived(sourceChainId uint16, sourceNttManagerAddress [32]byte, payload TransceiverStructsNttManagerMessage) (*types.Transaction, error) {
	return _Abi.Contract.AttestationReceived(&_Abi.TransactOpts, sourceChainId, sourceNttManagerAddress, payload)
}

// CompleteOutboundQueuedTransfer is a paid mutator transaction binding the contract method 0x97c35146.
//
// Solidity: function completeOutboundQueuedTransfer(uint64 ) payable returns(uint64)
func (_Abi *AbiTransactor) CompleteOutboundQueuedTransfer(opts *bind.TransactOpts, arg0 uint64) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "completeOutboundQueuedTransfer", arg0)
}

// CompleteOutboundQueuedTransfer is a paid mutator transaction binding the contract method 0x97c35146.
//
// Solidity: function completeOutboundQueuedTransfer(uint64 ) payable returns(uint64)
func (_Abi *AbiSession) CompleteOutboundQueuedTransfer(arg0 uint64) (*types.Transaction, error) {
	return _Abi.Contract.CompleteOutboundQueuedTransfer(&_Abi.TransactOpts, arg0)
}

// CompleteOutboundQueuedTransfer is a paid mutator transaction binding the contract method 0x97c35146.
//
// Solidity: function completeOutboundQueuedTransfer(uint64 ) payable returns(uint64)
func (_Abi *AbiTransactorSession) CompleteOutboundQueuedTransfer(arg0 uint64) (*types.Transaction, error) {
	return _Abi.Contract.CompleteOutboundQueuedTransfer(&_Abi.TransactOpts, arg0)
}

// DisableEarning is a paid mutator transaction binding the contract method 0xa8afc01f.
//
// Solidity: function disableEarning() returns()
func (_Abi *AbiTransactor) DisableEarning(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "disableEarning")
}

// DisableEarning is a paid mutator transaction binding the contract method 0xa8afc01f.
//
// Solidity: function disableEarning() returns()
func (_Abi *AbiSession) DisableEarning() (*types.Transaction, error) {
	return _Abi.Contract.DisableEarning(&_Abi.TransactOpts)
}

// DisableEarning is a paid mutator transaction binding the contract method 0xa8afc01f.
//
// Solidity: function disableEarning() returns()
func (_Abi *AbiTransactorSession) DisableEarning() (*types.Transaction, error) {
	return _Abi.Contract.DisableEarning(&_Abi.TransactOpts)
}

// EnableEarning is a paid mutator transaction binding the contract method 0xc967891a.
//
// Solidity: function enableEarning() returns()
func (_Abi *AbiTransactor) EnableEarning(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "enableEarning")
}

// EnableEarning is a paid mutator transaction binding the contract method 0xc967891a.
//
// Solidity: function enableEarning() returns()
func (_Abi *AbiSession) EnableEarning() (*types.Transaction, error) {
	return _Abi.Contract.EnableEarning(&_Abi.TransactOpts)
}

// EnableEarning is a paid mutator transaction binding the contract method 0xc967891a.
//
// Solidity: function enableEarning() returns()
func (_Abi *AbiTransactorSession) EnableEarning() (*types.Transaction, error) {
	return _Abi.Contract.EnableEarning(&_Abi.TransactOpts)
}

// ExecuteMsg is a paid mutator transaction binding the contract method 0xda4856a1.
//
// Solidity: function executeMsg(uint16 sourceChainId, bytes32 sourceNttManagerAddress, (bytes32,bytes32,bytes) message) returns()
func (_Abi *AbiTransactor) ExecuteMsg(opts *bind.TransactOpts, sourceChainId uint16, sourceNttManagerAddress [32]byte, message TransceiverStructsNttManagerMessage) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "executeMsg", sourceChainId, sourceNttManagerAddress, message)
}

// ExecuteMsg is a paid mutator transaction binding the contract method 0xda4856a1.
//
// Solidity: function executeMsg(uint16 sourceChainId, bytes32 sourceNttManagerAddress, (bytes32,bytes32,bytes) message) returns()
func (_Abi *AbiSession) ExecuteMsg(sourceChainId uint16, sourceNttManagerAddress [32]byte, message TransceiverStructsNttManagerMessage) (*types.Transaction, error) {
	return _Abi.Contract.ExecuteMsg(&_Abi.TransactOpts, sourceChainId, sourceNttManagerAddress, message)
}

// ExecuteMsg is a paid mutator transaction binding the contract method 0xda4856a1.
//
// Solidity: function executeMsg(uint16 sourceChainId, bytes32 sourceNttManagerAddress, (bytes32,bytes32,bytes) message) returns()
func (_Abi *AbiTransactorSession) ExecuteMsg(sourceChainId uint16, sourceNttManagerAddress [32]byte, message TransceiverStructsNttManagerMessage) (*types.Transaction, error) {
	return _Abi.Contract.ExecuteMsg(&_Abi.TransactOpts, sourceChainId, sourceNttManagerAddress, message)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_Abi *AbiTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_Abi *AbiSession) Initialize() (*types.Transaction, error) {
	return _Abi.Contract.Initialize(&_Abi.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_Abi *AbiTransactorSession) Initialize() (*types.Transaction, error) {
	return _Abi.Contract.Initialize(&_Abi.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Abi *AbiTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Abi *AbiSession) Migrate() (*types.Transaction, error) {
	return _Abi.Contract.Migrate(&_Abi.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Abi *AbiTransactorSession) Migrate() (*types.Transaction, error) {
	return _Abi.Contract.Migrate(&_Abi.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Abi *AbiTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Abi *AbiSession) Pause() (*types.Transaction, error) {
	return _Abi.Contract.Pause(&_Abi.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Abi *AbiTransactorSession) Pause() (*types.Transaction, error) {
	return _Abi.Contract.Pause(&_Abi.TransactOpts)
}

// RemoveTransceiver is a paid mutator transaction binding the contract method 0x9f86029c.
//
// Solidity: function removeTransceiver(address transceiver) returns()
func (_Abi *AbiTransactor) RemoveTransceiver(opts *bind.TransactOpts, transceiver common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "removeTransceiver", transceiver)
}

// RemoveTransceiver is a paid mutator transaction binding the contract method 0x9f86029c.
//
// Solidity: function removeTransceiver(address transceiver) returns()
func (_Abi *AbiSession) RemoveTransceiver(transceiver common.Address) (*types.Transaction, error) {
	return _Abi.Contract.RemoveTransceiver(&_Abi.TransactOpts, transceiver)
}

// RemoveTransceiver is a paid mutator transaction binding the contract method 0x9f86029c.
//
// Solidity: function removeTransceiver(address transceiver) returns()
func (_Abi *AbiTransactorSession) RemoveTransceiver(transceiver common.Address) (*types.Transaction, error) {
	return _Abi.Contract.RemoveTransceiver(&_Abi.TransactOpts, transceiver)
}

// SendMTokenIndex is a paid mutator transaction binding the contract method 0xf6f61b3a.
//
// Solidity: function sendMTokenIndex(uint16 destinationChainId_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Abi *AbiTransactor) SendMTokenIndex(opts *bind.TransactOpts, destinationChainId_ uint16, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "sendMTokenIndex", destinationChainId_, refundAddress_)
}

// SendMTokenIndex is a paid mutator transaction binding the contract method 0xf6f61b3a.
//
// Solidity: function sendMTokenIndex(uint16 destinationChainId_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Abi *AbiSession) SendMTokenIndex(destinationChainId_ uint16, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.SendMTokenIndex(&_Abi.TransactOpts, destinationChainId_, refundAddress_)
}

// SendMTokenIndex is a paid mutator transaction binding the contract method 0xf6f61b3a.
//
// Solidity: function sendMTokenIndex(uint16 destinationChainId_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Abi *AbiTransactorSession) SendMTokenIndex(destinationChainId_ uint16, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.SendMTokenIndex(&_Abi.TransactOpts, destinationChainId_, refundAddress_)
}

// SendRegistrarKey is a paid mutator transaction binding the contract method 0x1f33899e.
//
// Solidity: function sendRegistrarKey(uint16 destinationChainId_, bytes32 key_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Abi *AbiTransactor) SendRegistrarKey(opts *bind.TransactOpts, destinationChainId_ uint16, key_ [32]byte, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "sendRegistrarKey", destinationChainId_, key_, refundAddress_)
}

// SendRegistrarKey is a paid mutator transaction binding the contract method 0x1f33899e.
//
// Solidity: function sendRegistrarKey(uint16 destinationChainId_, bytes32 key_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Abi *AbiSession) SendRegistrarKey(destinationChainId_ uint16, key_ [32]byte, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.SendRegistrarKey(&_Abi.TransactOpts, destinationChainId_, key_, refundAddress_)
}

// SendRegistrarKey is a paid mutator transaction binding the contract method 0x1f33899e.
//
// Solidity: function sendRegistrarKey(uint16 destinationChainId_, bytes32 key_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Abi *AbiTransactorSession) SendRegistrarKey(destinationChainId_ uint16, key_ [32]byte, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.SendRegistrarKey(&_Abi.TransactOpts, destinationChainId_, key_, refundAddress_)
}

// SendRegistrarListStatus is a paid mutator transaction binding the contract method 0x76deb363.
//
// Solidity: function sendRegistrarListStatus(uint16 destinationChainId_, bytes32 listName_, address account_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Abi *AbiTransactor) SendRegistrarListStatus(opts *bind.TransactOpts, destinationChainId_ uint16, listName_ [32]byte, account_ common.Address, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "sendRegistrarListStatus", destinationChainId_, listName_, account_, refundAddress_)
}

// SendRegistrarListStatus is a paid mutator transaction binding the contract method 0x76deb363.
//
// Solidity: function sendRegistrarListStatus(uint16 destinationChainId_, bytes32 listName_, address account_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Abi *AbiSession) SendRegistrarListStatus(destinationChainId_ uint16, listName_ [32]byte, account_ common.Address, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.SendRegistrarListStatus(&_Abi.TransactOpts, destinationChainId_, listName_, account_, refundAddress_)
}

// SendRegistrarListStatus is a paid mutator transaction binding the contract method 0x76deb363.
//
// Solidity: function sendRegistrarListStatus(uint16 destinationChainId_, bytes32 listName_, address account_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Abi *AbiTransactorSession) SendRegistrarListStatus(destinationChainId_ uint16, listName_ [32]byte, account_ common.Address, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.SendRegistrarListStatus(&_Abi.TransactOpts, destinationChainId_, listName_, account_, refundAddress_)
}

// SetDestinationMToken is a paid mutator transaction binding the contract method 0x76a4a539.
//
// Solidity: function setDestinationMToken(uint16 destinationChainId_, bytes32 mToken_) returns()
func (_Abi *AbiTransactor) SetDestinationMToken(opts *bind.TransactOpts, destinationChainId_ uint16, mToken_ [32]byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setDestinationMToken", destinationChainId_, mToken_)
}

// SetDestinationMToken is a paid mutator transaction binding the contract method 0x76a4a539.
//
// Solidity: function setDestinationMToken(uint16 destinationChainId_, bytes32 mToken_) returns()
func (_Abi *AbiSession) SetDestinationMToken(destinationChainId_ uint16, mToken_ [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.SetDestinationMToken(&_Abi.TransactOpts, destinationChainId_, mToken_)
}

// SetDestinationMToken is a paid mutator transaction binding the contract method 0x76a4a539.
//
// Solidity: function setDestinationMToken(uint16 destinationChainId_, bytes32 mToken_) returns()
func (_Abi *AbiTransactorSession) SetDestinationMToken(destinationChainId_ uint16, mToken_ [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.SetDestinationMToken(&_Abi.TransactOpts, destinationChainId_, mToken_)
}

// SetInboundLimit is a paid mutator transaction binding the contract method 0x186ce612.
//
// Solidity: function setInboundLimit(uint256 limit, uint16 chainId_) returns()
func (_Abi *AbiTransactor) SetInboundLimit(opts *bind.TransactOpts, limit *big.Int, chainId_ uint16) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setInboundLimit", limit, chainId_)
}

// SetInboundLimit is a paid mutator transaction binding the contract method 0x186ce612.
//
// Solidity: function setInboundLimit(uint256 limit, uint16 chainId_) returns()
func (_Abi *AbiSession) SetInboundLimit(limit *big.Int, chainId_ uint16) (*types.Transaction, error) {
	return _Abi.Contract.SetInboundLimit(&_Abi.TransactOpts, limit, chainId_)
}

// SetInboundLimit is a paid mutator transaction binding the contract method 0x186ce612.
//
// Solidity: function setInboundLimit(uint256 limit, uint16 chainId_) returns()
func (_Abi *AbiTransactorSession) SetInboundLimit(limit *big.Int, chainId_ uint16) (*types.Transaction, error) {
	return _Abi.Contract.SetInboundLimit(&_Abi.TransactOpts, limit, chainId_)
}

// SetOutboundLimit is a paid mutator transaction binding the contract method 0x19017175.
//
// Solidity: function setOutboundLimit(uint256 limit) returns()
func (_Abi *AbiTransactor) SetOutboundLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setOutboundLimit", limit)
}

// SetOutboundLimit is a paid mutator transaction binding the contract method 0x19017175.
//
// Solidity: function setOutboundLimit(uint256 limit) returns()
func (_Abi *AbiSession) SetOutboundLimit(limit *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetOutboundLimit(&_Abi.TransactOpts, limit)
}

// SetOutboundLimit is a paid mutator transaction binding the contract method 0x19017175.
//
// Solidity: function setOutboundLimit(uint256 limit) returns()
func (_Abi *AbiTransactorSession) SetOutboundLimit(limit *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetOutboundLimit(&_Abi.TransactOpts, limit)
}

// SetPeer is a paid mutator transaction binding the contract method 0x7c918634.
//
// Solidity: function setPeer(uint16 peerChainId, bytes32 peerContract, uint8 decimals, uint256 inboundLimit) returns()
func (_Abi *AbiTransactor) SetPeer(opts *bind.TransactOpts, peerChainId uint16, peerContract [32]byte, decimals uint8, inboundLimit *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setPeer", peerChainId, peerContract, decimals, inboundLimit)
}

// SetPeer is a paid mutator transaction binding the contract method 0x7c918634.
//
// Solidity: function setPeer(uint16 peerChainId, bytes32 peerContract, uint8 decimals, uint256 inboundLimit) returns()
func (_Abi *AbiSession) SetPeer(peerChainId uint16, peerContract [32]byte, decimals uint8, inboundLimit *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetPeer(&_Abi.TransactOpts, peerChainId, peerContract, decimals, inboundLimit)
}

// SetPeer is a paid mutator transaction binding the contract method 0x7c918634.
//
// Solidity: function setPeer(uint16 peerChainId, bytes32 peerContract, uint8 decimals, uint256 inboundLimit) returns()
func (_Abi *AbiTransactorSession) SetPeer(peerChainId uint16, peerContract [32]byte, decimals uint8, inboundLimit *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetPeer(&_Abi.TransactOpts, peerChainId, peerContract, decimals, inboundLimit)
}

// SetSupportedBridgingPath is a paid mutator transaction binding the contract method 0x1482cb4a.
//
// Solidity: function setSupportedBridgingPath(address sourceToken_, uint16 destinationChainId_, bytes32 destinationToken_, bool supported_) returns()
func (_Abi *AbiTransactor) SetSupportedBridgingPath(opts *bind.TransactOpts, sourceToken_ common.Address, destinationChainId_ uint16, destinationToken_ [32]byte, supported_ bool) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setSupportedBridgingPath", sourceToken_, destinationChainId_, destinationToken_, supported_)
}

// SetSupportedBridgingPath is a paid mutator transaction binding the contract method 0x1482cb4a.
//
// Solidity: function setSupportedBridgingPath(address sourceToken_, uint16 destinationChainId_, bytes32 destinationToken_, bool supported_) returns()
func (_Abi *AbiSession) SetSupportedBridgingPath(sourceToken_ common.Address, destinationChainId_ uint16, destinationToken_ [32]byte, supported_ bool) (*types.Transaction, error) {
	return _Abi.Contract.SetSupportedBridgingPath(&_Abi.TransactOpts, sourceToken_, destinationChainId_, destinationToken_, supported_)
}

// SetSupportedBridgingPath is a paid mutator transaction binding the contract method 0x1482cb4a.
//
// Solidity: function setSupportedBridgingPath(address sourceToken_, uint16 destinationChainId_, bytes32 destinationToken_, bool supported_) returns()
func (_Abi *AbiTransactorSession) SetSupportedBridgingPath(sourceToken_ common.Address, destinationChainId_ uint16, destinationToken_ [32]byte, supported_ bool) (*types.Transaction, error) {
	return _Abi.Contract.SetSupportedBridgingPath(&_Abi.TransactOpts, sourceToken_, destinationChainId_, destinationToken_, supported_)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xe5a98603.
//
// Solidity: function setThreshold(uint8 threshold) returns()
func (_Abi *AbiTransactor) SetThreshold(opts *bind.TransactOpts, threshold uint8) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setThreshold", threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xe5a98603.
//
// Solidity: function setThreshold(uint8 threshold) returns()
func (_Abi *AbiSession) SetThreshold(threshold uint8) (*types.Transaction, error) {
	return _Abi.Contract.SetThreshold(&_Abi.TransactOpts, threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xe5a98603.
//
// Solidity: function setThreshold(uint8 threshold) returns()
func (_Abi *AbiTransactorSession) SetThreshold(threshold uint8) (*types.Transaction, error) {
	return _Abi.Contract.SetThreshold(&_Abi.TransactOpts, threshold)
}

// SetTransceiver is a paid mutator transaction binding the contract method 0x203e4a9b.
//
// Solidity: function setTransceiver(address transceiver) returns()
func (_Abi *AbiTransactor) SetTransceiver(opts *bind.TransactOpts, transceiver common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setTransceiver", transceiver)
}

// SetTransceiver is a paid mutator transaction binding the contract method 0x203e4a9b.
//
// Solidity: function setTransceiver(address transceiver) returns()
func (_Abi *AbiSession) SetTransceiver(transceiver common.Address) (*types.Transaction, error) {
	return _Abi.Contract.SetTransceiver(&_Abi.TransactOpts, transceiver)
}

// SetTransceiver is a paid mutator transaction binding the contract method 0x203e4a9b.
//
// Solidity: function setTransceiver(address transceiver) returns()
func (_Abi *AbiTransactorSession) SetTransceiver(transceiver common.Address) (*types.Transaction, error) {
	return _Abi.Contract.SetTransceiver(&_Abi.TransactOpts, transceiver)
}

// Transfer is a paid mutator transaction binding the contract method 0x961b94d0.
//
// Solidity: function transfer(uint256 amount, uint16 recipientChain, bytes32 recipient) payable returns(uint64)
func (_Abi *AbiTransactor) Transfer(opts *bind.TransactOpts, amount *big.Int, recipientChain uint16, recipient [32]byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transfer", amount, recipientChain, recipient)
}

// Transfer is a paid mutator transaction binding the contract method 0x961b94d0.
//
// Solidity: function transfer(uint256 amount, uint16 recipientChain, bytes32 recipient) payable returns(uint64)
func (_Abi *AbiSession) Transfer(amount *big.Int, recipientChain uint16, recipient [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.Transfer(&_Abi.TransactOpts, amount, recipientChain, recipient)
}

// Transfer is a paid mutator transaction binding the contract method 0x961b94d0.
//
// Solidity: function transfer(uint256 amount, uint16 recipientChain, bytes32 recipient) payable returns(uint64)
func (_Abi *AbiTransactorSession) Transfer(amount *big.Int, recipientChain uint16, recipient [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.Transfer(&_Abi.TransactOpts, amount, recipientChain, recipient)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xb293f97f.
//
// Solidity: function transfer(uint256 amount, uint16 recipientChain, bytes32 recipient, bytes32 refundAddress, bool shouldQueue, bytes transceiverInstructions) payable returns(uint64)
func (_Abi *AbiTransactor) Transfer0(opts *bind.TransactOpts, amount *big.Int, recipientChain uint16, recipient [32]byte, refundAddress [32]byte, shouldQueue bool, transceiverInstructions []byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transfer0", amount, recipientChain, recipient, refundAddress, shouldQueue, transceiverInstructions)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xb293f97f.
//
// Solidity: function transfer(uint256 amount, uint16 recipientChain, bytes32 recipient, bytes32 refundAddress, bool shouldQueue, bytes transceiverInstructions) payable returns(uint64)
func (_Abi *AbiSession) Transfer0(amount *big.Int, recipientChain uint16, recipient [32]byte, refundAddress [32]byte, shouldQueue bool, transceiverInstructions []byte) (*types.Transaction, error) {
	return _Abi.Contract.Transfer0(&_Abi.TransactOpts, amount, recipientChain, recipient, refundAddress, shouldQueue, transceiverInstructions)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xb293f97f.
//
// Solidity: function transfer(uint256 amount, uint16 recipientChain, bytes32 recipient, bytes32 refundAddress, bool shouldQueue, bytes transceiverInstructions) payable returns(uint64)
func (_Abi *AbiTransactorSession) Transfer0(amount *big.Int, recipientChain uint16, recipient [32]byte, refundAddress [32]byte, shouldQueue bool, transceiverInstructions []byte) (*types.Transaction, error) {
	return _Abi.Contract.Transfer0(&_Abi.TransactOpts, amount, recipientChain, recipient, refundAddress, shouldQueue, transceiverInstructions)
}

// TransferMLikeToken is a paid mutator transaction binding the contract method 0x68df33c6.
//
// Solidity: function transferMLikeToken(uint256 amount_, address sourceToken_, uint16 destinationChainId_, bytes32 destinationToken_, bytes32 recipient_, bytes32 refundAddress_) payable returns(uint64 sequence_)
func (_Abi *AbiTransactor) TransferMLikeToken(opts *bind.TransactOpts, amount_ *big.Int, sourceToken_ common.Address, destinationChainId_ uint16, destinationToken_ [32]byte, recipient_ [32]byte, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transferMLikeToken", amount_, sourceToken_, destinationChainId_, destinationToken_, recipient_, refundAddress_)
}

// TransferMLikeToken is a paid mutator transaction binding the contract method 0x68df33c6.
//
// Solidity: function transferMLikeToken(uint256 amount_, address sourceToken_, uint16 destinationChainId_, bytes32 destinationToken_, bytes32 recipient_, bytes32 refundAddress_) payable returns(uint64 sequence_)
func (_Abi *AbiSession) TransferMLikeToken(amount_ *big.Int, sourceToken_ common.Address, destinationChainId_ uint16, destinationToken_ [32]byte, recipient_ [32]byte, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.TransferMLikeToken(&_Abi.TransactOpts, amount_, sourceToken_, destinationChainId_, destinationToken_, recipient_, refundAddress_)
}

// TransferMLikeToken is a paid mutator transaction binding the contract method 0x68df33c6.
//
// Solidity: function transferMLikeToken(uint256 amount_, address sourceToken_, uint16 destinationChainId_, bytes32 destinationToken_, bytes32 recipient_, bytes32 refundAddress_) payable returns(uint64 sequence_)
func (_Abi *AbiTransactorSession) TransferMLikeToken(amount_ *big.Int, sourceToken_ common.Address, destinationChainId_ uint16, destinationToken_ [32]byte, recipient_ [32]byte, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.TransferMLikeToken(&_Abi.TransactOpts, amount_, sourceToken_, destinationChainId_, destinationToken_, recipient_, refundAddress_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferOwnership(&_Abi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferOwnership(&_Abi.TransactOpts, newOwner)
}

// TransferPauserCapability is a paid mutator transaction binding the contract method 0x036de8af.
//
// Solidity: function transferPauserCapability(address newPauser) returns()
func (_Abi *AbiTransactor) TransferPauserCapability(opts *bind.TransactOpts, newPauser common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transferPauserCapability", newPauser)
}

// TransferPauserCapability is a paid mutator transaction binding the contract method 0x036de8af.
//
// Solidity: function transferPauserCapability(address newPauser) returns()
func (_Abi *AbiSession) TransferPauserCapability(newPauser common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferPauserCapability(&_Abi.TransactOpts, newPauser)
}

// TransferPauserCapability is a paid mutator transaction binding the contract method 0x036de8af.
//
// Solidity: function transferPauserCapability(address newPauser) returns()
func (_Abi *AbiTransactorSession) TransferPauserCapability(newPauser common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferPauserCapability(&_Abi.TransactOpts, newPauser)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Abi *AbiTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Abi *AbiSession) Unpause() (*types.Transaction, error) {
	return _Abi.Contract.Unpause(&_Abi.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Abi *AbiTransactorSession) Unpause() (*types.Transaction, error) {
	return _Abi.Contract.Unpause(&_Abi.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newImplementation) returns()
func (_Abi *AbiTransactor) Upgrade(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "upgrade", newImplementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newImplementation) returns()
func (_Abi *AbiSession) Upgrade(newImplementation common.Address) (*types.Transaction, error) {
	return _Abi.Contract.Upgrade(&_Abi.TransactOpts, newImplementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newImplementation) returns()
func (_Abi *AbiTransactorSession) Upgrade(newImplementation common.Address) (*types.Transaction, error) {
	return _Abi.Contract.Upgrade(&_Abi.TransactOpts, newImplementation)
}

// AbiAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Abi contract.
type AbiAdminChangedIterator struct {
	Event *AbiAdminChanged // Event containing the contract specifics and raw log

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
func (it *AbiAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiAdminChanged)
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
		it.Event = new(AbiAdminChanged)
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
func (it *AbiAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiAdminChanged represents a AdminChanged event raised by the Abi contract.
type AbiAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Abi *AbiFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AbiAdminChangedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AbiAdminChangedIterator{contract: _Abi.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Abi *AbiFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AbiAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiAdminChanged)
				if err := _Abi.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Abi *AbiFilterer) ParseAdminChanged(log types.Log) (*AbiAdminChanged, error) {
	event := new(AbiAdminChanged)
	if err := _Abi.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Abi contract.
type AbiBeaconUpgradedIterator struct {
	Event *AbiBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AbiBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiBeaconUpgraded)
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
		it.Event = new(AbiBeaconUpgraded)
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
func (it *AbiBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiBeaconUpgraded represents a BeaconUpgraded event raised by the Abi contract.
type AbiBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Abi *AbiFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AbiBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AbiBeaconUpgradedIterator{contract: _Abi.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Abi *AbiFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AbiBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiBeaconUpgraded)
				if err := _Abi.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Abi *AbiFilterer) ParseBeaconUpgraded(log types.Log) (*AbiBeaconUpgraded, error) {
	event := new(AbiBeaconUpgraded)
	if err := _Abi.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiDestinationMTokenSetIterator is returned from FilterDestinationMTokenSet and is used to iterate over the raw logs and unpacked data for DestinationMTokenSet events raised by the Abi contract.
type AbiDestinationMTokenSetIterator struct {
	Event *AbiDestinationMTokenSet // Event containing the contract specifics and raw log

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
func (it *AbiDestinationMTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiDestinationMTokenSet)
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
		it.Event = new(AbiDestinationMTokenSet)
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
func (it *AbiDestinationMTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiDestinationMTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiDestinationMTokenSet represents a DestinationMTokenSet event raised by the Abi contract.
type AbiDestinationMTokenSet struct {
	DestinationChainId uint16
	MToken             [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDestinationMTokenSet is a free log retrieval operation binding the contract event 0xe6658eeba76934cef78dc5f420ee40406e4876e8d3e767f7b700c63daf1bf0ba.
//
// Solidity: event DestinationMTokenSet(uint16 indexed destinationChainId, bytes32 mToken)
func (_Abi *AbiFilterer) FilterDestinationMTokenSet(opts *bind.FilterOpts, destinationChainId []uint16) (*AbiDestinationMTokenSetIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "DestinationMTokenSet", destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &AbiDestinationMTokenSetIterator{contract: _Abi.contract, event: "DestinationMTokenSet", logs: logs, sub: sub}, nil
}

// WatchDestinationMTokenSet is a free log subscription operation binding the contract event 0xe6658eeba76934cef78dc5f420ee40406e4876e8d3e767f7b700c63daf1bf0ba.
//
// Solidity: event DestinationMTokenSet(uint16 indexed destinationChainId, bytes32 mToken)
func (_Abi *AbiFilterer) WatchDestinationMTokenSet(opts *bind.WatchOpts, sink chan<- *AbiDestinationMTokenSet, destinationChainId []uint16) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "DestinationMTokenSet", destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiDestinationMTokenSet)
				if err := _Abi.contract.UnpackLog(event, "DestinationMTokenSet", log); err != nil {
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

// ParseDestinationMTokenSet is a log parse operation binding the contract event 0xe6658eeba76934cef78dc5f420ee40406e4876e8d3e767f7b700c63daf1bf0ba.
//
// Solidity: event DestinationMTokenSet(uint16 indexed destinationChainId, bytes32 mToken)
func (_Abi *AbiFilterer) ParseDestinationMTokenSet(log types.Log) (*AbiDestinationMTokenSet, error) {
	event := new(AbiDestinationMTokenSet)
	if err := _Abi.contract.UnpackLog(event, "DestinationMTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiEarningDisabledIterator is returned from FilterEarningDisabled and is used to iterate over the raw logs and unpacked data for EarningDisabled events raised by the Abi contract.
type AbiEarningDisabledIterator struct {
	Event *AbiEarningDisabled // Event containing the contract specifics and raw log

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
func (it *AbiEarningDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiEarningDisabled)
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
		it.Event = new(AbiEarningDisabled)
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
func (it *AbiEarningDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiEarningDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiEarningDisabled represents a EarningDisabled event raised by the Abi contract.
type AbiEarningDisabled struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEarningDisabled is a free log retrieval operation binding the contract event 0xee580fdb4da10ea17aa673e6f5c8c2370b4166d6a94bc88900e5a96d0589e3ce.
//
// Solidity: event EarningDisabled(uint128 index)
func (_Abi *AbiFilterer) FilterEarningDisabled(opts *bind.FilterOpts) (*AbiEarningDisabledIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "EarningDisabled")
	if err != nil {
		return nil, err
	}
	return &AbiEarningDisabledIterator{contract: _Abi.contract, event: "EarningDisabled", logs: logs, sub: sub}, nil
}

// WatchEarningDisabled is a free log subscription operation binding the contract event 0xee580fdb4da10ea17aa673e6f5c8c2370b4166d6a94bc88900e5a96d0589e3ce.
//
// Solidity: event EarningDisabled(uint128 index)
func (_Abi *AbiFilterer) WatchEarningDisabled(opts *bind.WatchOpts, sink chan<- *AbiEarningDisabled) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "EarningDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiEarningDisabled)
				if err := _Abi.contract.UnpackLog(event, "EarningDisabled", log); err != nil {
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

// ParseEarningDisabled is a log parse operation binding the contract event 0xee580fdb4da10ea17aa673e6f5c8c2370b4166d6a94bc88900e5a96d0589e3ce.
//
// Solidity: event EarningDisabled(uint128 index)
func (_Abi *AbiFilterer) ParseEarningDisabled(log types.Log) (*AbiEarningDisabled, error) {
	event := new(AbiEarningDisabled)
	if err := _Abi.contract.UnpackLog(event, "EarningDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiEarningEnabledIterator is returned from FilterEarningEnabled and is used to iterate over the raw logs and unpacked data for EarningEnabled events raised by the Abi contract.
type AbiEarningEnabledIterator struct {
	Event *AbiEarningEnabled // Event containing the contract specifics and raw log

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
func (it *AbiEarningEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiEarningEnabled)
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
		it.Event = new(AbiEarningEnabled)
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
func (it *AbiEarningEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiEarningEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiEarningEnabled represents a EarningEnabled event raised by the Abi contract.
type AbiEarningEnabled struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEarningEnabled is a free log retrieval operation binding the contract event 0x5098de6eb11dbd1127cf4dcd5e960e3944d48a7570b9b1939cff715cb35c5a18.
//
// Solidity: event EarningEnabled(uint128 index)
func (_Abi *AbiFilterer) FilterEarningEnabled(opts *bind.FilterOpts) (*AbiEarningEnabledIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "EarningEnabled")
	if err != nil {
		return nil, err
	}
	return &AbiEarningEnabledIterator{contract: _Abi.contract, event: "EarningEnabled", logs: logs, sub: sub}, nil
}

// WatchEarningEnabled is a free log subscription operation binding the contract event 0x5098de6eb11dbd1127cf4dcd5e960e3944d48a7570b9b1939cff715cb35c5a18.
//
// Solidity: event EarningEnabled(uint128 index)
func (_Abi *AbiFilterer) WatchEarningEnabled(opts *bind.WatchOpts, sink chan<- *AbiEarningEnabled) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "EarningEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiEarningEnabled)
				if err := _Abi.contract.UnpackLog(event, "EarningEnabled", log); err != nil {
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

// ParseEarningEnabled is a log parse operation binding the contract event 0x5098de6eb11dbd1127cf4dcd5e960e3944d48a7570b9b1939cff715cb35c5a18.
//
// Solidity: event EarningEnabled(uint128 index)
func (_Abi *AbiFilterer) ParseEarningEnabled(log types.Log) (*AbiEarningEnabled, error) {
	event := new(AbiEarningEnabled)
	if err := _Abi.contract.UnpackLog(event, "EarningEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiInboundTransferLimitUpdatedIterator is returned from FilterInboundTransferLimitUpdated and is used to iterate over the raw logs and unpacked data for InboundTransferLimitUpdated events raised by the Abi contract.
type AbiInboundTransferLimitUpdatedIterator struct {
	Event *AbiInboundTransferLimitUpdated // Event containing the contract specifics and raw log

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
func (it *AbiInboundTransferLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiInboundTransferLimitUpdated)
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
		it.Event = new(AbiInboundTransferLimitUpdated)
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
func (it *AbiInboundTransferLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiInboundTransferLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiInboundTransferLimitUpdated represents a InboundTransferLimitUpdated event raised by the Abi contract.
type AbiInboundTransferLimitUpdated struct {
	ChainId  uint16
	OldLimit *big.Int
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInboundTransferLimitUpdated is a free log retrieval operation binding the contract event 0x739ed886fd81a3ddc9f4b327ab69152e513cd45b26fda0c73660eaca8e119301.
//
// Solidity: event InboundTransferLimitUpdated(uint16 indexed chainId, uint256 oldLimit, uint256 newLimit)
func (_Abi *AbiFilterer) FilterInboundTransferLimitUpdated(opts *bind.FilterOpts, chainId []uint16) (*AbiInboundTransferLimitUpdatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "InboundTransferLimitUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &AbiInboundTransferLimitUpdatedIterator{contract: _Abi.contract, event: "InboundTransferLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchInboundTransferLimitUpdated is a free log subscription operation binding the contract event 0x739ed886fd81a3ddc9f4b327ab69152e513cd45b26fda0c73660eaca8e119301.
//
// Solidity: event InboundTransferLimitUpdated(uint16 indexed chainId, uint256 oldLimit, uint256 newLimit)
func (_Abi *AbiFilterer) WatchInboundTransferLimitUpdated(opts *bind.WatchOpts, sink chan<- *AbiInboundTransferLimitUpdated, chainId []uint16) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "InboundTransferLimitUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiInboundTransferLimitUpdated)
				if err := _Abi.contract.UnpackLog(event, "InboundTransferLimitUpdated", log); err != nil {
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

// ParseInboundTransferLimitUpdated is a log parse operation binding the contract event 0x739ed886fd81a3ddc9f4b327ab69152e513cd45b26fda0c73660eaca8e119301.
//
// Solidity: event InboundTransferLimitUpdated(uint16 indexed chainId, uint256 oldLimit, uint256 newLimit)
func (_Abi *AbiFilterer) ParseInboundTransferLimitUpdated(log types.Log) (*AbiInboundTransferLimitUpdated, error) {
	event := new(AbiInboundTransferLimitUpdated)
	if err := _Abi.contract.UnpackLog(event, "InboundTransferLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiInboundTransferQueuedIterator is returned from FilterInboundTransferQueued and is used to iterate over the raw logs and unpacked data for InboundTransferQueued events raised by the Abi contract.
type AbiInboundTransferQueuedIterator struct {
	Event *AbiInboundTransferQueued // Event containing the contract specifics and raw log

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
func (it *AbiInboundTransferQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiInboundTransferQueued)
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
		it.Event = new(AbiInboundTransferQueued)
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
func (it *AbiInboundTransferQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiInboundTransferQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiInboundTransferQueued represents a InboundTransferQueued event raised by the Abi contract.
type AbiInboundTransferQueued struct {
	Digest [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInboundTransferQueued is a free log retrieval operation binding the contract event 0x7f63c9251d82a933210c2b6d0b0f116252c3c116788120e64e8e8215df6f3162.
//
// Solidity: event InboundTransferQueued(bytes32 digest)
func (_Abi *AbiFilterer) FilterInboundTransferQueued(opts *bind.FilterOpts) (*AbiInboundTransferQueuedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "InboundTransferQueued")
	if err != nil {
		return nil, err
	}
	return &AbiInboundTransferQueuedIterator{contract: _Abi.contract, event: "InboundTransferQueued", logs: logs, sub: sub}, nil
}

// WatchInboundTransferQueued is a free log subscription operation binding the contract event 0x7f63c9251d82a933210c2b6d0b0f116252c3c116788120e64e8e8215df6f3162.
//
// Solidity: event InboundTransferQueued(bytes32 digest)
func (_Abi *AbiFilterer) WatchInboundTransferQueued(opts *bind.WatchOpts, sink chan<- *AbiInboundTransferQueued) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "InboundTransferQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiInboundTransferQueued)
				if err := _Abi.contract.UnpackLog(event, "InboundTransferQueued", log); err != nil {
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

// ParseInboundTransferQueued is a log parse operation binding the contract event 0x7f63c9251d82a933210c2b6d0b0f116252c3c116788120e64e8e8215df6f3162.
//
// Solidity: event InboundTransferQueued(bytes32 digest)
func (_Abi *AbiFilterer) ParseInboundTransferQueued(log types.Log) (*AbiInboundTransferQueued, error) {
	event := new(AbiInboundTransferQueued)
	if err := _Abi.contract.UnpackLog(event, "InboundTransferQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Abi contract.
type AbiInitializedIterator struct {
	Event *AbiInitialized // Event containing the contract specifics and raw log

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
func (it *AbiInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiInitialized)
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
		it.Event = new(AbiInitialized)
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
func (it *AbiInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiInitialized represents a Initialized event raised by the Abi contract.
type AbiInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Abi *AbiFilterer) FilterInitialized(opts *bind.FilterOpts) (*AbiInitializedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AbiInitializedIterator{contract: _Abi.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Abi *AbiFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AbiInitialized) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiInitialized)
				if err := _Abi.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Abi *AbiFilterer) ParseInitialized(log types.Log) (*AbiInitialized, error) {
	event := new(AbiInitialized)
	if err := _Abi.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiMTokenIndexSentIterator is returned from FilterMTokenIndexSent and is used to iterate over the raw logs and unpacked data for MTokenIndexSent events raised by the Abi contract.
type AbiMTokenIndexSentIterator struct {
	Event *AbiMTokenIndexSent // Event containing the contract specifics and raw log

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
func (it *AbiMTokenIndexSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiMTokenIndexSent)
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
		it.Event = new(AbiMTokenIndexSent)
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
func (it *AbiMTokenIndexSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiMTokenIndexSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiMTokenIndexSent represents a MTokenIndexSent event raised by the Abi contract.
type AbiMTokenIndexSent struct {
	DestinationChainId uint16
	MessageId          [32]byte
	Index              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMTokenIndexSent is a free log retrieval operation binding the contract event 0x3e1e343484a686e91a59f00b76fa98aaf17cb636b12f75a6654f432f45e64c14.
//
// Solidity: event MTokenIndexSent(uint16 indexed destinationChainId, bytes32 messageId, uint128 index)
func (_Abi *AbiFilterer) FilterMTokenIndexSent(opts *bind.FilterOpts, destinationChainId []uint16) (*AbiMTokenIndexSentIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "MTokenIndexSent", destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &AbiMTokenIndexSentIterator{contract: _Abi.contract, event: "MTokenIndexSent", logs: logs, sub: sub}, nil
}

// WatchMTokenIndexSent is a free log subscription operation binding the contract event 0x3e1e343484a686e91a59f00b76fa98aaf17cb636b12f75a6654f432f45e64c14.
//
// Solidity: event MTokenIndexSent(uint16 indexed destinationChainId, bytes32 messageId, uint128 index)
func (_Abi *AbiFilterer) WatchMTokenIndexSent(opts *bind.WatchOpts, sink chan<- *AbiMTokenIndexSent, destinationChainId []uint16) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "MTokenIndexSent", destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiMTokenIndexSent)
				if err := _Abi.contract.UnpackLog(event, "MTokenIndexSent", log); err != nil {
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

// ParseMTokenIndexSent is a log parse operation binding the contract event 0x3e1e343484a686e91a59f00b76fa98aaf17cb636b12f75a6654f432f45e64c14.
//
// Solidity: event MTokenIndexSent(uint16 indexed destinationChainId, bytes32 messageId, uint128 index)
func (_Abi *AbiFilterer) ParseMTokenIndexSent(log types.Log) (*AbiMTokenIndexSent, error) {
	event := new(AbiMTokenIndexSent)
	if err := _Abi.contract.UnpackLog(event, "MTokenIndexSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiMTokenReceivedIterator is returned from FilterMTokenReceived and is used to iterate over the raw logs and unpacked data for MTokenReceived events raised by the Abi contract.
type AbiMTokenReceivedIterator struct {
	Event *AbiMTokenReceived // Event containing the contract specifics and raw log

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
func (it *AbiMTokenReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiMTokenReceived)
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
		it.Event = new(AbiMTokenReceived)
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
func (it *AbiMTokenReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiMTokenReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiMTokenReceived represents a MTokenReceived event raised by the Abi contract.
type AbiMTokenReceived struct {
	SourceChainId    uint16
	DestinationToken common.Address
	Sender           [32]byte
	Recipient        common.Address
	Amount           *big.Int
	Index            *big.Int
	MessageId        [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMTokenReceived is a free log retrieval operation binding the contract event 0x1fa93583b8ef840b2aa9730b35f5454180f063fd8d3ffec60122d9e1f567c7be.
//
// Solidity: event MTokenReceived(uint16 sourceChainId, address indexed destinationToken, bytes32 indexed sender, address indexed recipient, uint256 amount, uint128 index, bytes32 messageId)
func (_Abi *AbiFilterer) FilterMTokenReceived(opts *bind.FilterOpts, destinationToken []common.Address, sender [][32]byte, recipient []common.Address) (*AbiMTokenReceivedIterator, error) {

	var destinationTokenRule []interface{}
	for _, destinationTokenItem := range destinationToken {
		destinationTokenRule = append(destinationTokenRule, destinationTokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "MTokenReceived", destinationTokenRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &AbiMTokenReceivedIterator{contract: _Abi.contract, event: "MTokenReceived", logs: logs, sub: sub}, nil
}

// WatchMTokenReceived is a free log subscription operation binding the contract event 0x1fa93583b8ef840b2aa9730b35f5454180f063fd8d3ffec60122d9e1f567c7be.
//
// Solidity: event MTokenReceived(uint16 sourceChainId, address indexed destinationToken, bytes32 indexed sender, address indexed recipient, uint256 amount, uint128 index, bytes32 messageId)
func (_Abi *AbiFilterer) WatchMTokenReceived(opts *bind.WatchOpts, sink chan<- *AbiMTokenReceived, destinationToken []common.Address, sender [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var destinationTokenRule []interface{}
	for _, destinationTokenItem := range destinationToken {
		destinationTokenRule = append(destinationTokenRule, destinationTokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "MTokenReceived", destinationTokenRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiMTokenReceived)
				if err := _Abi.contract.UnpackLog(event, "MTokenReceived", log); err != nil {
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

// ParseMTokenReceived is a log parse operation binding the contract event 0x1fa93583b8ef840b2aa9730b35f5454180f063fd8d3ffec60122d9e1f567c7be.
//
// Solidity: event MTokenReceived(uint16 sourceChainId, address indexed destinationToken, bytes32 indexed sender, address indexed recipient, uint256 amount, uint128 index, bytes32 messageId)
func (_Abi *AbiFilterer) ParseMTokenReceived(log types.Log) (*AbiMTokenReceived, error) {
	event := new(AbiMTokenReceived)
	if err := _Abi.contract.UnpackLog(event, "MTokenReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiMTokenSentIterator is returned from FilterMTokenSent and is used to iterate over the raw logs and unpacked data for MTokenSent events raised by the Abi contract.
type AbiMTokenSentIterator struct {
	Event *AbiMTokenSent // Event containing the contract specifics and raw log

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
func (it *AbiMTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiMTokenSent)
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
		it.Event = new(AbiMTokenSent)
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
func (it *AbiMTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiMTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiMTokenSent represents a MTokenSent event raised by the Abi contract.
type AbiMTokenSent struct {
	SourceToken        common.Address
	DestinationChainId uint16
	DestinationToken   [32]byte
	Sender             common.Address
	Recipient          [32]byte
	Amount             *big.Int
	Index              *big.Int
	MessageId          [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMTokenSent is a free log retrieval operation binding the contract event 0x98195478444bfed374ee6b111345e4c2088ade272cebe3ab76ed80bb170e400e.
//
// Solidity: event MTokenSent(address indexed sourceToken, uint16 destinationChainId, bytes32 destinationToken, address indexed sender, bytes32 indexed recipient, uint256 amount, uint128 index, bytes32 messageId)
func (_Abi *AbiFilterer) FilterMTokenSent(opts *bind.FilterOpts, sourceToken []common.Address, sender []common.Address, recipient [][32]byte) (*AbiMTokenSentIterator, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "MTokenSent", sourceTokenRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &AbiMTokenSentIterator{contract: _Abi.contract, event: "MTokenSent", logs: logs, sub: sub}, nil
}

// WatchMTokenSent is a free log subscription operation binding the contract event 0x98195478444bfed374ee6b111345e4c2088ade272cebe3ab76ed80bb170e400e.
//
// Solidity: event MTokenSent(address indexed sourceToken, uint16 destinationChainId, bytes32 destinationToken, address indexed sender, bytes32 indexed recipient, uint256 amount, uint128 index, bytes32 messageId)
func (_Abi *AbiFilterer) WatchMTokenSent(opts *bind.WatchOpts, sink chan<- *AbiMTokenSent, sourceToken []common.Address, sender []common.Address, recipient [][32]byte) (event.Subscription, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "MTokenSent", sourceTokenRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiMTokenSent)
				if err := _Abi.contract.UnpackLog(event, "MTokenSent", log); err != nil {
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

// ParseMTokenSent is a log parse operation binding the contract event 0x98195478444bfed374ee6b111345e4c2088ade272cebe3ab76ed80bb170e400e.
//
// Solidity: event MTokenSent(address indexed sourceToken, uint16 destinationChainId, bytes32 destinationToken, address indexed sender, bytes32 indexed recipient, uint256 amount, uint128 index, bytes32 messageId)
func (_Abi *AbiFilterer) ParseMTokenSent(log types.Log) (*AbiMTokenSent, error) {
	event := new(AbiMTokenSent)
	if err := _Abi.contract.UnpackLog(event, "MTokenSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiMessageAlreadyExecutedIterator is returned from FilterMessageAlreadyExecuted and is used to iterate over the raw logs and unpacked data for MessageAlreadyExecuted events raised by the Abi contract.
type AbiMessageAlreadyExecutedIterator struct {
	Event *AbiMessageAlreadyExecuted // Event containing the contract specifics and raw log

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
func (it *AbiMessageAlreadyExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiMessageAlreadyExecuted)
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
		it.Event = new(AbiMessageAlreadyExecuted)
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
func (it *AbiMessageAlreadyExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiMessageAlreadyExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiMessageAlreadyExecuted represents a MessageAlreadyExecuted event raised by the Abi contract.
type AbiMessageAlreadyExecuted struct {
	SourceNttManager [32]byte
	MsgHash          [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMessageAlreadyExecuted is a free log retrieval operation binding the contract event 0x4069dff8c9df7e38d2867c0910bd96fd61787695e5380281148c04932d02bef2.
//
// Solidity: event MessageAlreadyExecuted(bytes32 indexed sourceNttManager, bytes32 indexed msgHash)
func (_Abi *AbiFilterer) FilterMessageAlreadyExecuted(opts *bind.FilterOpts, sourceNttManager [][32]byte, msgHash [][32]byte) (*AbiMessageAlreadyExecutedIterator, error) {

	var sourceNttManagerRule []interface{}
	for _, sourceNttManagerItem := range sourceNttManager {
		sourceNttManagerRule = append(sourceNttManagerRule, sourceNttManagerItem)
	}
	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "MessageAlreadyExecuted", sourceNttManagerRule, msgHashRule)
	if err != nil {
		return nil, err
	}
	return &AbiMessageAlreadyExecutedIterator{contract: _Abi.contract, event: "MessageAlreadyExecuted", logs: logs, sub: sub}, nil
}

// WatchMessageAlreadyExecuted is a free log subscription operation binding the contract event 0x4069dff8c9df7e38d2867c0910bd96fd61787695e5380281148c04932d02bef2.
//
// Solidity: event MessageAlreadyExecuted(bytes32 indexed sourceNttManager, bytes32 indexed msgHash)
func (_Abi *AbiFilterer) WatchMessageAlreadyExecuted(opts *bind.WatchOpts, sink chan<- *AbiMessageAlreadyExecuted, sourceNttManager [][32]byte, msgHash [][32]byte) (event.Subscription, error) {

	var sourceNttManagerRule []interface{}
	for _, sourceNttManagerItem := range sourceNttManager {
		sourceNttManagerRule = append(sourceNttManagerRule, sourceNttManagerItem)
	}
	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "MessageAlreadyExecuted", sourceNttManagerRule, msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiMessageAlreadyExecuted)
				if err := _Abi.contract.UnpackLog(event, "MessageAlreadyExecuted", log); err != nil {
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

// ParseMessageAlreadyExecuted is a log parse operation binding the contract event 0x4069dff8c9df7e38d2867c0910bd96fd61787695e5380281148c04932d02bef2.
//
// Solidity: event MessageAlreadyExecuted(bytes32 indexed sourceNttManager, bytes32 indexed msgHash)
func (_Abi *AbiFilterer) ParseMessageAlreadyExecuted(log types.Log) (*AbiMessageAlreadyExecuted, error) {
	event := new(AbiMessageAlreadyExecuted)
	if err := _Abi.contract.UnpackLog(event, "MessageAlreadyExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiMessageAttestedToIterator is returned from FilterMessageAttestedTo and is used to iterate over the raw logs and unpacked data for MessageAttestedTo events raised by the Abi contract.
type AbiMessageAttestedToIterator struct {
	Event *AbiMessageAttestedTo // Event containing the contract specifics and raw log

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
func (it *AbiMessageAttestedToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiMessageAttestedTo)
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
		it.Event = new(AbiMessageAttestedTo)
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
func (it *AbiMessageAttestedToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiMessageAttestedToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiMessageAttestedTo represents a MessageAttestedTo event raised by the Abi contract.
type AbiMessageAttestedTo struct {
	Digest      [32]byte
	Transceiver common.Address
	Index       uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageAttestedTo is a free log retrieval operation binding the contract event 0x35a2101eaac94b493e0dfca061f9a7f087913fde8678e7cde0aca9897edba0e5.
//
// Solidity: event MessageAttestedTo(bytes32 digest, address transceiver, uint8 index)
func (_Abi *AbiFilterer) FilterMessageAttestedTo(opts *bind.FilterOpts) (*AbiMessageAttestedToIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "MessageAttestedTo")
	if err != nil {
		return nil, err
	}
	return &AbiMessageAttestedToIterator{contract: _Abi.contract, event: "MessageAttestedTo", logs: logs, sub: sub}, nil
}

// WatchMessageAttestedTo is a free log subscription operation binding the contract event 0x35a2101eaac94b493e0dfca061f9a7f087913fde8678e7cde0aca9897edba0e5.
//
// Solidity: event MessageAttestedTo(bytes32 digest, address transceiver, uint8 index)
func (_Abi *AbiFilterer) WatchMessageAttestedTo(opts *bind.WatchOpts, sink chan<- *AbiMessageAttestedTo) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "MessageAttestedTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiMessageAttestedTo)
				if err := _Abi.contract.UnpackLog(event, "MessageAttestedTo", log); err != nil {
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

// ParseMessageAttestedTo is a log parse operation binding the contract event 0x35a2101eaac94b493e0dfca061f9a7f087913fde8678e7cde0aca9897edba0e5.
//
// Solidity: event MessageAttestedTo(bytes32 digest, address transceiver, uint8 index)
func (_Abi *AbiFilterer) ParseMessageAttestedTo(log types.Log) (*AbiMessageAttestedTo, error) {
	event := new(AbiMessageAttestedTo)
	if err := _Abi.contract.UnpackLog(event, "MessageAttestedTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiNotPausedIterator is returned from FilterNotPaused and is used to iterate over the raw logs and unpacked data for NotPaused events raised by the Abi contract.
type AbiNotPausedIterator struct {
	Event *AbiNotPaused // Event containing the contract specifics and raw log

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
func (it *AbiNotPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiNotPaused)
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
		it.Event = new(AbiNotPaused)
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
func (it *AbiNotPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiNotPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiNotPaused represents a NotPaused event raised by the Abi contract.
type AbiNotPaused struct {
	NotPaused bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNotPaused is a free log retrieval operation binding the contract event 0xe11c2112add17fb763d3bd59f63b10429c3e11373da4fb8ef6725107a2fdc4b0.
//
// Solidity: event NotPaused(bool notPaused)
func (_Abi *AbiFilterer) FilterNotPaused(opts *bind.FilterOpts) (*AbiNotPausedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "NotPaused")
	if err != nil {
		return nil, err
	}
	return &AbiNotPausedIterator{contract: _Abi.contract, event: "NotPaused", logs: logs, sub: sub}, nil
}

// WatchNotPaused is a free log subscription operation binding the contract event 0xe11c2112add17fb763d3bd59f63b10429c3e11373da4fb8ef6725107a2fdc4b0.
//
// Solidity: event NotPaused(bool notPaused)
func (_Abi *AbiFilterer) WatchNotPaused(opts *bind.WatchOpts, sink chan<- *AbiNotPaused) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "NotPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiNotPaused)
				if err := _Abi.contract.UnpackLog(event, "NotPaused", log); err != nil {
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

// ParseNotPaused is a log parse operation binding the contract event 0xe11c2112add17fb763d3bd59f63b10429c3e11373da4fb8ef6725107a2fdc4b0.
//
// Solidity: event NotPaused(bool notPaused)
func (_Abi *AbiFilterer) ParseNotPaused(log types.Log) (*AbiNotPaused, error) {
	event := new(AbiNotPaused)
	if err := _Abi.contract.UnpackLog(event, "NotPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiOutboundTransferCancelledIterator is returned from FilterOutboundTransferCancelled and is used to iterate over the raw logs and unpacked data for OutboundTransferCancelled events raised by the Abi contract.
type AbiOutboundTransferCancelledIterator struct {
	Event *AbiOutboundTransferCancelled // Event containing the contract specifics and raw log

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
func (it *AbiOutboundTransferCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiOutboundTransferCancelled)
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
		it.Event = new(AbiOutboundTransferCancelled)
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
func (it *AbiOutboundTransferCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiOutboundTransferCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiOutboundTransferCancelled represents a OutboundTransferCancelled event raised by the Abi contract.
type AbiOutboundTransferCancelled struct {
	Sequence  *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOutboundTransferCancelled is a free log retrieval operation binding the contract event 0xf80e572ae1b63e2449629b6c7d783add85c36473926f216077f17ee002bcfd07.
//
// Solidity: event OutboundTransferCancelled(uint256 sequence, address recipient, uint256 amount)
func (_Abi *AbiFilterer) FilterOutboundTransferCancelled(opts *bind.FilterOpts) (*AbiOutboundTransferCancelledIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "OutboundTransferCancelled")
	if err != nil {
		return nil, err
	}
	return &AbiOutboundTransferCancelledIterator{contract: _Abi.contract, event: "OutboundTransferCancelled", logs: logs, sub: sub}, nil
}

// WatchOutboundTransferCancelled is a free log subscription operation binding the contract event 0xf80e572ae1b63e2449629b6c7d783add85c36473926f216077f17ee002bcfd07.
//
// Solidity: event OutboundTransferCancelled(uint256 sequence, address recipient, uint256 amount)
func (_Abi *AbiFilterer) WatchOutboundTransferCancelled(opts *bind.WatchOpts, sink chan<- *AbiOutboundTransferCancelled) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "OutboundTransferCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiOutboundTransferCancelled)
				if err := _Abi.contract.UnpackLog(event, "OutboundTransferCancelled", log); err != nil {
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

// ParseOutboundTransferCancelled is a log parse operation binding the contract event 0xf80e572ae1b63e2449629b6c7d783add85c36473926f216077f17ee002bcfd07.
//
// Solidity: event OutboundTransferCancelled(uint256 sequence, address recipient, uint256 amount)
func (_Abi *AbiFilterer) ParseOutboundTransferCancelled(log types.Log) (*AbiOutboundTransferCancelled, error) {
	event := new(AbiOutboundTransferCancelled)
	if err := _Abi.contract.UnpackLog(event, "OutboundTransferCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiOutboundTransferLimitUpdatedIterator is returned from FilterOutboundTransferLimitUpdated and is used to iterate over the raw logs and unpacked data for OutboundTransferLimitUpdated events raised by the Abi contract.
type AbiOutboundTransferLimitUpdatedIterator struct {
	Event *AbiOutboundTransferLimitUpdated // Event containing the contract specifics and raw log

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
func (it *AbiOutboundTransferLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiOutboundTransferLimitUpdated)
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
		it.Event = new(AbiOutboundTransferLimitUpdated)
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
func (it *AbiOutboundTransferLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiOutboundTransferLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiOutboundTransferLimitUpdated represents a OutboundTransferLimitUpdated event raised by the Abi contract.
type AbiOutboundTransferLimitUpdated struct {
	OldLimit *big.Int
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOutboundTransferLimitUpdated is a free log retrieval operation binding the contract event 0x7e3b0fc388be9d36273f66210aed83be975df3a9adfffa4c734033f498f362cd.
//
// Solidity: event OutboundTransferLimitUpdated(uint256 oldLimit, uint256 newLimit)
func (_Abi *AbiFilterer) FilterOutboundTransferLimitUpdated(opts *bind.FilterOpts) (*AbiOutboundTransferLimitUpdatedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "OutboundTransferLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &AbiOutboundTransferLimitUpdatedIterator{contract: _Abi.contract, event: "OutboundTransferLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchOutboundTransferLimitUpdated is a free log subscription operation binding the contract event 0x7e3b0fc388be9d36273f66210aed83be975df3a9adfffa4c734033f498f362cd.
//
// Solidity: event OutboundTransferLimitUpdated(uint256 oldLimit, uint256 newLimit)
func (_Abi *AbiFilterer) WatchOutboundTransferLimitUpdated(opts *bind.WatchOpts, sink chan<- *AbiOutboundTransferLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "OutboundTransferLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiOutboundTransferLimitUpdated)
				if err := _Abi.contract.UnpackLog(event, "OutboundTransferLimitUpdated", log); err != nil {
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

// ParseOutboundTransferLimitUpdated is a log parse operation binding the contract event 0x7e3b0fc388be9d36273f66210aed83be975df3a9adfffa4c734033f498f362cd.
//
// Solidity: event OutboundTransferLimitUpdated(uint256 oldLimit, uint256 newLimit)
func (_Abi *AbiFilterer) ParseOutboundTransferLimitUpdated(log types.Log) (*AbiOutboundTransferLimitUpdated, error) {
	event := new(AbiOutboundTransferLimitUpdated)
	if err := _Abi.contract.UnpackLog(event, "OutboundTransferLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiOutboundTransferQueuedIterator is returned from FilterOutboundTransferQueued and is used to iterate over the raw logs and unpacked data for OutboundTransferQueued events raised by the Abi contract.
type AbiOutboundTransferQueuedIterator struct {
	Event *AbiOutboundTransferQueued // Event containing the contract specifics and raw log

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
func (it *AbiOutboundTransferQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiOutboundTransferQueued)
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
		it.Event = new(AbiOutboundTransferQueued)
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
func (it *AbiOutboundTransferQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiOutboundTransferQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiOutboundTransferQueued represents a OutboundTransferQueued event raised by the Abi contract.
type AbiOutboundTransferQueued struct {
	QueueSequence uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutboundTransferQueued is a free log retrieval operation binding the contract event 0x69add1952a6a6b9cb86f04d05f0cb605cbb469a50ae916139d34495a9991481f.
//
// Solidity: event OutboundTransferQueued(uint64 queueSequence)
func (_Abi *AbiFilterer) FilterOutboundTransferQueued(opts *bind.FilterOpts) (*AbiOutboundTransferQueuedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "OutboundTransferQueued")
	if err != nil {
		return nil, err
	}
	return &AbiOutboundTransferQueuedIterator{contract: _Abi.contract, event: "OutboundTransferQueued", logs: logs, sub: sub}, nil
}

// WatchOutboundTransferQueued is a free log subscription operation binding the contract event 0x69add1952a6a6b9cb86f04d05f0cb605cbb469a50ae916139d34495a9991481f.
//
// Solidity: event OutboundTransferQueued(uint64 queueSequence)
func (_Abi *AbiFilterer) WatchOutboundTransferQueued(opts *bind.WatchOpts, sink chan<- *AbiOutboundTransferQueued) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "OutboundTransferQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiOutboundTransferQueued)
				if err := _Abi.contract.UnpackLog(event, "OutboundTransferQueued", log); err != nil {
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

// ParseOutboundTransferQueued is a log parse operation binding the contract event 0x69add1952a6a6b9cb86f04d05f0cb605cbb469a50ae916139d34495a9991481f.
//
// Solidity: event OutboundTransferQueued(uint64 queueSequence)
func (_Abi *AbiFilterer) ParseOutboundTransferQueued(log types.Log) (*AbiOutboundTransferQueued, error) {
	event := new(AbiOutboundTransferQueued)
	if err := _Abi.contract.UnpackLog(event, "OutboundTransferQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiOutboundTransferRateLimitedIterator is returned from FilterOutboundTransferRateLimited and is used to iterate over the raw logs and unpacked data for OutboundTransferRateLimited events raised by the Abi contract.
type AbiOutboundTransferRateLimitedIterator struct {
	Event *AbiOutboundTransferRateLimited // Event containing the contract specifics and raw log

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
func (it *AbiOutboundTransferRateLimitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiOutboundTransferRateLimited)
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
		it.Event = new(AbiOutboundTransferRateLimited)
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
func (it *AbiOutboundTransferRateLimitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiOutboundTransferRateLimitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiOutboundTransferRateLimited represents a OutboundTransferRateLimited event raised by the Abi contract.
type AbiOutboundTransferRateLimited struct {
	Sender          common.Address
	Sequence        uint64
	Amount          *big.Int
	CurrentCapacity *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOutboundTransferRateLimited is a free log retrieval operation binding the contract event 0xf33512b84e24a49905c26c6991942fc5a9652411769fc1e448f967cdb049f08a.
//
// Solidity: event OutboundTransferRateLimited(address indexed sender, uint64 sequence, uint256 amount, uint256 currentCapacity)
func (_Abi *AbiFilterer) FilterOutboundTransferRateLimited(opts *bind.FilterOpts, sender []common.Address) (*AbiOutboundTransferRateLimitedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "OutboundTransferRateLimited", senderRule)
	if err != nil {
		return nil, err
	}
	return &AbiOutboundTransferRateLimitedIterator{contract: _Abi.contract, event: "OutboundTransferRateLimited", logs: logs, sub: sub}, nil
}

// WatchOutboundTransferRateLimited is a free log subscription operation binding the contract event 0xf33512b84e24a49905c26c6991942fc5a9652411769fc1e448f967cdb049f08a.
//
// Solidity: event OutboundTransferRateLimited(address indexed sender, uint64 sequence, uint256 amount, uint256 currentCapacity)
func (_Abi *AbiFilterer) WatchOutboundTransferRateLimited(opts *bind.WatchOpts, sink chan<- *AbiOutboundTransferRateLimited, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "OutboundTransferRateLimited", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiOutboundTransferRateLimited)
				if err := _Abi.contract.UnpackLog(event, "OutboundTransferRateLimited", log); err != nil {
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

// ParseOutboundTransferRateLimited is a log parse operation binding the contract event 0xf33512b84e24a49905c26c6991942fc5a9652411769fc1e448f967cdb049f08a.
//
// Solidity: event OutboundTransferRateLimited(address indexed sender, uint64 sequence, uint256 amount, uint256 currentCapacity)
func (_Abi *AbiFilterer) ParseOutboundTransferRateLimited(log types.Log) (*AbiOutboundTransferRateLimited, error) {
	event := new(AbiOutboundTransferRateLimited)
	if err := _Abi.contract.UnpackLog(event, "OutboundTransferRateLimited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Abi contract.
type AbiOwnershipTransferredIterator struct {
	Event *AbiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AbiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiOwnershipTransferred)
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
		it.Event = new(AbiOwnershipTransferred)
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
func (it *AbiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiOwnershipTransferred represents a OwnershipTransferred event raised by the Abi contract.
type AbiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AbiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AbiOwnershipTransferredIterator{contract: _Abi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AbiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiOwnershipTransferred)
				if err := _Abi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Abi *AbiFilterer) ParseOwnershipTransferred(log types.Log) (*AbiOwnershipTransferred, error) {
	event := new(AbiOwnershipTransferred)
	if err := _Abi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Abi contract.
type AbiPausedIterator struct {
	Event *AbiPaused // Event containing the contract specifics and raw log

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
func (it *AbiPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiPaused)
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
		it.Event = new(AbiPaused)
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
func (it *AbiPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiPaused represents a Paused event raised by the Abi contract.
type AbiPaused struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool paused)
func (_Abi *AbiFilterer) FilterPaused(opts *bind.FilterOpts) (*AbiPausedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AbiPausedIterator{contract: _Abi.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool paused)
func (_Abi *AbiFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AbiPaused) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiPaused)
				if err := _Abi.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool paused)
func (_Abi *AbiFilterer) ParsePaused(log types.Log) (*AbiPaused, error) {
	event := new(AbiPaused)
	if err := _Abi.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiPauserTransferredIterator is returned from FilterPauserTransferred and is used to iterate over the raw logs and unpacked data for PauserTransferred events raised by the Abi contract.
type AbiPauserTransferredIterator struct {
	Event *AbiPauserTransferred // Event containing the contract specifics and raw log

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
func (it *AbiPauserTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiPauserTransferred)
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
		it.Event = new(AbiPauserTransferred)
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
func (it *AbiPauserTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiPauserTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiPauserTransferred represents a PauserTransferred event raised by the Abi contract.
type AbiPauserTransferred struct {
	OldPauser common.Address
	NewPauser common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPauserTransferred is a free log retrieval operation binding the contract event 0x51c4874e0f23f262e04a38c51751336dde72126d67f53eb672aaff02996b3ef6.
//
// Solidity: event PauserTransferred(address indexed oldPauser, address indexed newPauser)
func (_Abi *AbiFilterer) FilterPauserTransferred(opts *bind.FilterOpts, oldPauser []common.Address, newPauser []common.Address) (*AbiPauserTransferredIterator, error) {

	var oldPauserRule []interface{}
	for _, oldPauserItem := range oldPauser {
		oldPauserRule = append(oldPauserRule, oldPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "PauserTransferred", oldPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return &AbiPauserTransferredIterator{contract: _Abi.contract, event: "PauserTransferred", logs: logs, sub: sub}, nil
}

// WatchPauserTransferred is a free log subscription operation binding the contract event 0x51c4874e0f23f262e04a38c51751336dde72126d67f53eb672aaff02996b3ef6.
//
// Solidity: event PauserTransferred(address indexed oldPauser, address indexed newPauser)
func (_Abi *AbiFilterer) WatchPauserTransferred(opts *bind.WatchOpts, sink chan<- *AbiPauserTransferred, oldPauser []common.Address, newPauser []common.Address) (event.Subscription, error) {

	var oldPauserRule []interface{}
	for _, oldPauserItem := range oldPauser {
		oldPauserRule = append(oldPauserRule, oldPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "PauserTransferred", oldPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiPauserTransferred)
				if err := _Abi.contract.UnpackLog(event, "PauserTransferred", log); err != nil {
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

// ParsePauserTransferred is a log parse operation binding the contract event 0x51c4874e0f23f262e04a38c51751336dde72126d67f53eb672aaff02996b3ef6.
//
// Solidity: event PauserTransferred(address indexed oldPauser, address indexed newPauser)
func (_Abi *AbiFilterer) ParsePauserTransferred(log types.Log) (*AbiPauserTransferred, error) {
	event := new(AbiPauserTransferred)
	if err := _Abi.contract.UnpackLog(event, "PauserTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiPeerUpdatedIterator is returned from FilterPeerUpdated and is used to iterate over the raw logs and unpacked data for PeerUpdated events raised by the Abi contract.
type AbiPeerUpdatedIterator struct {
	Event *AbiPeerUpdated // Event containing the contract specifics and raw log

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
func (it *AbiPeerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiPeerUpdated)
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
		it.Event = new(AbiPeerUpdated)
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
func (it *AbiPeerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiPeerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiPeerUpdated represents a PeerUpdated event raised by the Abi contract.
type AbiPeerUpdated struct {
	ChainId         uint16
	OldPeerContract [32]byte
	OldPeerDecimals uint8
	PeerContract    [32]byte
	PeerDecimals    uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPeerUpdated is a free log retrieval operation binding the contract event 0x1456404e7f41f35c3daac941bb50bad417a66275c3040061b4287d787719599d.
//
// Solidity: event PeerUpdated(uint16 indexed chainId_, bytes32 oldPeerContract, uint8 oldPeerDecimals, bytes32 peerContract, uint8 peerDecimals)
func (_Abi *AbiFilterer) FilterPeerUpdated(opts *bind.FilterOpts, chainId_ []uint16) (*AbiPeerUpdatedIterator, error) {

	var chainId_Rule []interface{}
	for _, chainId_Item := range chainId_ {
		chainId_Rule = append(chainId_Rule, chainId_Item)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "PeerUpdated", chainId_Rule)
	if err != nil {
		return nil, err
	}
	return &AbiPeerUpdatedIterator{contract: _Abi.contract, event: "PeerUpdated", logs: logs, sub: sub}, nil
}

// WatchPeerUpdated is a free log subscription operation binding the contract event 0x1456404e7f41f35c3daac941bb50bad417a66275c3040061b4287d787719599d.
//
// Solidity: event PeerUpdated(uint16 indexed chainId_, bytes32 oldPeerContract, uint8 oldPeerDecimals, bytes32 peerContract, uint8 peerDecimals)
func (_Abi *AbiFilterer) WatchPeerUpdated(opts *bind.WatchOpts, sink chan<- *AbiPeerUpdated, chainId_ []uint16) (event.Subscription, error) {

	var chainId_Rule []interface{}
	for _, chainId_Item := range chainId_ {
		chainId_Rule = append(chainId_Rule, chainId_Item)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "PeerUpdated", chainId_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiPeerUpdated)
				if err := _Abi.contract.UnpackLog(event, "PeerUpdated", log); err != nil {
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

// ParsePeerUpdated is a log parse operation binding the contract event 0x1456404e7f41f35c3daac941bb50bad417a66275c3040061b4287d787719599d.
//
// Solidity: event PeerUpdated(uint16 indexed chainId_, bytes32 oldPeerContract, uint8 oldPeerDecimals, bytes32 peerContract, uint8 peerDecimals)
func (_Abi *AbiFilterer) ParsePeerUpdated(log types.Log) (*AbiPeerUpdated, error) {
	event := new(AbiPeerUpdated)
	if err := _Abi.contract.UnpackLog(event, "PeerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRegistrarKeySentIterator is returned from FilterRegistrarKeySent and is used to iterate over the raw logs and unpacked data for RegistrarKeySent events raised by the Abi contract.
type AbiRegistrarKeySentIterator struct {
	Event *AbiRegistrarKeySent // Event containing the contract specifics and raw log

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
func (it *AbiRegistrarKeySentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRegistrarKeySent)
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
		it.Event = new(AbiRegistrarKeySent)
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
func (it *AbiRegistrarKeySentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRegistrarKeySentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRegistrarKeySent represents a RegistrarKeySent event raised by the Abi contract.
type AbiRegistrarKeySent struct {
	DestinationChainId uint16
	MessageId          [32]byte
	Key                [32]byte
	Value              [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRegistrarKeySent is a free log retrieval operation binding the contract event 0x4e419c12ebf7388707567369b500f54671c9323b834077c5391f6054c1169544.
//
// Solidity: event RegistrarKeySent(uint16 indexed destinationChainId, bytes32 messageId, bytes32 indexed key, bytes32 value)
func (_Abi *AbiFilterer) FilterRegistrarKeySent(opts *bind.FilterOpts, destinationChainId []uint16, key [][32]byte) (*AbiRegistrarKeySentIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "RegistrarKeySent", destinationChainIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &AbiRegistrarKeySentIterator{contract: _Abi.contract, event: "RegistrarKeySent", logs: logs, sub: sub}, nil
}

// WatchRegistrarKeySent is a free log subscription operation binding the contract event 0x4e419c12ebf7388707567369b500f54671c9323b834077c5391f6054c1169544.
//
// Solidity: event RegistrarKeySent(uint16 indexed destinationChainId, bytes32 messageId, bytes32 indexed key, bytes32 value)
func (_Abi *AbiFilterer) WatchRegistrarKeySent(opts *bind.WatchOpts, sink chan<- *AbiRegistrarKeySent, destinationChainId []uint16, key [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "RegistrarKeySent", destinationChainIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRegistrarKeySent)
				if err := _Abi.contract.UnpackLog(event, "RegistrarKeySent", log); err != nil {
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

// ParseRegistrarKeySent is a log parse operation binding the contract event 0x4e419c12ebf7388707567369b500f54671c9323b834077c5391f6054c1169544.
//
// Solidity: event RegistrarKeySent(uint16 indexed destinationChainId, bytes32 messageId, bytes32 indexed key, bytes32 value)
func (_Abi *AbiFilterer) ParseRegistrarKeySent(log types.Log) (*AbiRegistrarKeySent, error) {
	event := new(AbiRegistrarKeySent)
	if err := _Abi.contract.UnpackLog(event, "RegistrarKeySent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRegistrarListStatusSentIterator is returned from FilterRegistrarListStatusSent and is used to iterate over the raw logs and unpacked data for RegistrarListStatusSent events raised by the Abi contract.
type AbiRegistrarListStatusSentIterator struct {
	Event *AbiRegistrarListStatusSent // Event containing the contract specifics and raw log

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
func (it *AbiRegistrarListStatusSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRegistrarListStatusSent)
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
		it.Event = new(AbiRegistrarListStatusSent)
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
func (it *AbiRegistrarListStatusSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRegistrarListStatusSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRegistrarListStatusSent represents a RegistrarListStatusSent event raised by the Abi contract.
type AbiRegistrarListStatusSent struct {
	DestinationChainId uint16
	MessageId          [32]byte
	ListName           [32]byte
	Account            common.Address
	Status             bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRegistrarListStatusSent is a free log retrieval operation binding the contract event 0xe4170b1dad8af8275fc8c0a7ec8ccd5f05e5139db178ce9a9d6566696e02b832.
//
// Solidity: event RegistrarListStatusSent(uint16 indexed destinationChainId, bytes32 messageId, bytes32 indexed listName, address indexed account, bool status)
func (_Abi *AbiFilterer) FilterRegistrarListStatusSent(opts *bind.FilterOpts, destinationChainId []uint16, listName [][32]byte, account []common.Address) (*AbiRegistrarListStatusSentIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var listNameRule []interface{}
	for _, listNameItem := range listName {
		listNameRule = append(listNameRule, listNameItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "RegistrarListStatusSent", destinationChainIdRule, listNameRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AbiRegistrarListStatusSentIterator{contract: _Abi.contract, event: "RegistrarListStatusSent", logs: logs, sub: sub}, nil
}

// WatchRegistrarListStatusSent is a free log subscription operation binding the contract event 0xe4170b1dad8af8275fc8c0a7ec8ccd5f05e5139db178ce9a9d6566696e02b832.
//
// Solidity: event RegistrarListStatusSent(uint16 indexed destinationChainId, bytes32 messageId, bytes32 indexed listName, address indexed account, bool status)
func (_Abi *AbiFilterer) WatchRegistrarListStatusSent(opts *bind.WatchOpts, sink chan<- *AbiRegistrarListStatusSent, destinationChainId []uint16, listName [][32]byte, account []common.Address) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var listNameRule []interface{}
	for _, listNameItem := range listName {
		listNameRule = append(listNameRule, listNameItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "RegistrarListStatusSent", destinationChainIdRule, listNameRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRegistrarListStatusSent)
				if err := _Abi.contract.UnpackLog(event, "RegistrarListStatusSent", log); err != nil {
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

// ParseRegistrarListStatusSent is a log parse operation binding the contract event 0xe4170b1dad8af8275fc8c0a7ec8ccd5f05e5139db178ce9a9d6566696e02b832.
//
// Solidity: event RegistrarListStatusSent(uint16 indexed destinationChainId, bytes32 messageId, bytes32 indexed listName, address indexed account, bool status)
func (_Abi *AbiFilterer) ParseRegistrarListStatusSent(log types.Log) (*AbiRegistrarListStatusSent, error) {
	event := new(AbiRegistrarListStatusSent)
	if err := _Abi.contract.UnpackLog(event, "RegistrarListStatusSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiSupportedBridgingPathSetIterator is returned from FilterSupportedBridgingPathSet and is used to iterate over the raw logs and unpacked data for SupportedBridgingPathSet events raised by the Abi contract.
type AbiSupportedBridgingPathSetIterator struct {
	Event *AbiSupportedBridgingPathSet // Event containing the contract specifics and raw log

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
func (it *AbiSupportedBridgingPathSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiSupportedBridgingPathSet)
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
		it.Event = new(AbiSupportedBridgingPathSet)
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
func (it *AbiSupportedBridgingPathSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiSupportedBridgingPathSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiSupportedBridgingPathSet represents a SupportedBridgingPathSet event raised by the Abi contract.
type AbiSupportedBridgingPathSet struct {
	SourceToken        common.Address
	DestinationChainId uint16
	DestinationToken   [32]byte
	Supported          bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSupportedBridgingPathSet is a free log retrieval operation binding the contract event 0xeb8377c150d1438121333afa2027232605a400b1208ae6f9cfa156438e453f6d.
//
// Solidity: event SupportedBridgingPathSet(address indexed sourceToken, uint16 indexed destinationChainId, bytes32 indexed destinationToken, bool supported)
func (_Abi *AbiFilterer) FilterSupportedBridgingPathSet(opts *bind.FilterOpts, sourceToken []common.Address, destinationChainId []uint16, destinationToken [][32]byte) (*AbiSupportedBridgingPathSetIterator, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}
	var destinationTokenRule []interface{}
	for _, destinationTokenItem := range destinationToken {
		destinationTokenRule = append(destinationTokenRule, destinationTokenItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "SupportedBridgingPathSet", sourceTokenRule, destinationChainIdRule, destinationTokenRule)
	if err != nil {
		return nil, err
	}
	return &AbiSupportedBridgingPathSetIterator{contract: _Abi.contract, event: "SupportedBridgingPathSet", logs: logs, sub: sub}, nil
}

// WatchSupportedBridgingPathSet is a free log subscription operation binding the contract event 0xeb8377c150d1438121333afa2027232605a400b1208ae6f9cfa156438e453f6d.
//
// Solidity: event SupportedBridgingPathSet(address indexed sourceToken, uint16 indexed destinationChainId, bytes32 indexed destinationToken, bool supported)
func (_Abi *AbiFilterer) WatchSupportedBridgingPathSet(opts *bind.WatchOpts, sink chan<- *AbiSupportedBridgingPathSet, sourceToken []common.Address, destinationChainId []uint16, destinationToken [][32]byte) (event.Subscription, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}
	var destinationTokenRule []interface{}
	for _, destinationTokenItem := range destinationToken {
		destinationTokenRule = append(destinationTokenRule, destinationTokenItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "SupportedBridgingPathSet", sourceTokenRule, destinationChainIdRule, destinationTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiSupportedBridgingPathSet)
				if err := _Abi.contract.UnpackLog(event, "SupportedBridgingPathSet", log); err != nil {
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

// ParseSupportedBridgingPathSet is a log parse operation binding the contract event 0xeb8377c150d1438121333afa2027232605a400b1208ae6f9cfa156438e453f6d.
//
// Solidity: event SupportedBridgingPathSet(address indexed sourceToken, uint16 indexed destinationChainId, bytes32 indexed destinationToken, bool supported)
func (_Abi *AbiFilterer) ParseSupportedBridgingPathSet(log types.Log) (*AbiSupportedBridgingPathSet, error) {
	event := new(AbiSupportedBridgingPathSet)
	if err := _Abi.contract.UnpackLog(event, "SupportedBridgingPathSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the Abi contract.
type AbiThresholdChangedIterator struct {
	Event *AbiThresholdChanged // Event containing the contract specifics and raw log

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
func (it *AbiThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiThresholdChanged)
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
		it.Event = new(AbiThresholdChanged)
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
func (it *AbiThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiThresholdChanged represents a ThresholdChanged event raised by the Abi contract.
type AbiThresholdChanged struct {
	OldThreshold uint8
	Threshold    uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x2a855b929b9a53c6fb5b5ed248b27e502b709c088e036a5aa17620c8fc5085a9.
//
// Solidity: event ThresholdChanged(uint8 oldThreshold, uint8 threshold)
func (_Abi *AbiFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*AbiThresholdChangedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &AbiThresholdChangedIterator{contract: _Abi.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x2a855b929b9a53c6fb5b5ed248b27e502b709c088e036a5aa17620c8fc5085a9.
//
// Solidity: event ThresholdChanged(uint8 oldThreshold, uint8 threshold)
func (_Abi *AbiFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *AbiThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiThresholdChanged)
				if err := _Abi.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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

// ParseThresholdChanged is a log parse operation binding the contract event 0x2a855b929b9a53c6fb5b5ed248b27e502b709c088e036a5aa17620c8fc5085a9.
//
// Solidity: event ThresholdChanged(uint8 oldThreshold, uint8 threshold)
func (_Abi *AbiFilterer) ParseThresholdChanged(log types.Log) (*AbiThresholdChanged, error) {
	event := new(AbiThresholdChanged)
	if err := _Abi.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiTransceiverAddedIterator is returned from FilterTransceiverAdded and is used to iterate over the raw logs and unpacked data for TransceiverAdded events raised by the Abi contract.
type AbiTransceiverAddedIterator struct {
	Event *AbiTransceiverAdded // Event containing the contract specifics and raw log

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
func (it *AbiTransceiverAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiTransceiverAdded)
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
		it.Event = new(AbiTransceiverAdded)
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
func (it *AbiTransceiverAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiTransceiverAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiTransceiverAdded represents a TransceiverAdded event raised by the Abi contract.
type AbiTransceiverAdded struct {
	Transceiver     common.Address
	TransceiversNum *big.Int
	Threshold       uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransceiverAdded is a free log retrieval operation binding the contract event 0xf05962b5774c658e85ed80c91a75af9d66d2af2253dda480f90bce78aff5eda5.
//
// Solidity: event TransceiverAdded(address transceiver, uint256 transceiversNum, uint8 threshold)
func (_Abi *AbiFilterer) FilterTransceiverAdded(opts *bind.FilterOpts) (*AbiTransceiverAddedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "TransceiverAdded")
	if err != nil {
		return nil, err
	}
	return &AbiTransceiverAddedIterator{contract: _Abi.contract, event: "TransceiverAdded", logs: logs, sub: sub}, nil
}

// WatchTransceiverAdded is a free log subscription operation binding the contract event 0xf05962b5774c658e85ed80c91a75af9d66d2af2253dda480f90bce78aff5eda5.
//
// Solidity: event TransceiverAdded(address transceiver, uint256 transceiversNum, uint8 threshold)
func (_Abi *AbiFilterer) WatchTransceiverAdded(opts *bind.WatchOpts, sink chan<- *AbiTransceiverAdded) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "TransceiverAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiTransceiverAdded)
				if err := _Abi.contract.UnpackLog(event, "TransceiverAdded", log); err != nil {
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

// ParseTransceiverAdded is a log parse operation binding the contract event 0xf05962b5774c658e85ed80c91a75af9d66d2af2253dda480f90bce78aff5eda5.
//
// Solidity: event TransceiverAdded(address transceiver, uint256 transceiversNum, uint8 threshold)
func (_Abi *AbiFilterer) ParseTransceiverAdded(log types.Log) (*AbiTransceiverAdded, error) {
	event := new(AbiTransceiverAdded)
	if err := _Abi.contract.UnpackLog(event, "TransceiverAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiTransceiverRemovedIterator is returned from FilterTransceiverRemoved and is used to iterate over the raw logs and unpacked data for TransceiverRemoved events raised by the Abi contract.
type AbiTransceiverRemovedIterator struct {
	Event *AbiTransceiverRemoved // Event containing the contract specifics and raw log

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
func (it *AbiTransceiverRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiTransceiverRemoved)
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
		it.Event = new(AbiTransceiverRemoved)
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
func (it *AbiTransceiverRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiTransceiverRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiTransceiverRemoved represents a TransceiverRemoved event raised by the Abi contract.
type AbiTransceiverRemoved struct {
	Transceiver common.Address
	Threshold   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransceiverRemoved is a free log retrieval operation binding the contract event 0x697a3853515b88013ad432f29f53d406debc9509ed6d9313dcfe115250fcd18f.
//
// Solidity: event TransceiverRemoved(address transceiver, uint8 threshold)
func (_Abi *AbiFilterer) FilterTransceiverRemoved(opts *bind.FilterOpts) (*AbiTransceiverRemovedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "TransceiverRemoved")
	if err != nil {
		return nil, err
	}
	return &AbiTransceiverRemovedIterator{contract: _Abi.contract, event: "TransceiverRemoved", logs: logs, sub: sub}, nil
}

// WatchTransceiverRemoved is a free log subscription operation binding the contract event 0x697a3853515b88013ad432f29f53d406debc9509ed6d9313dcfe115250fcd18f.
//
// Solidity: event TransceiverRemoved(address transceiver, uint8 threshold)
func (_Abi *AbiFilterer) WatchTransceiverRemoved(opts *bind.WatchOpts, sink chan<- *AbiTransceiverRemoved) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "TransceiverRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiTransceiverRemoved)
				if err := _Abi.contract.UnpackLog(event, "TransceiverRemoved", log); err != nil {
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

// ParseTransceiverRemoved is a log parse operation binding the contract event 0x697a3853515b88013ad432f29f53d406debc9509ed6d9313dcfe115250fcd18f.
//
// Solidity: event TransceiverRemoved(address transceiver, uint8 threshold)
func (_Abi *AbiFilterer) ParseTransceiverRemoved(log types.Log) (*AbiTransceiverRemoved, error) {
	event := new(AbiTransceiverRemoved)
	if err := _Abi.contract.UnpackLog(event, "TransceiverRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiTransferRedeemedIterator is returned from FilterTransferRedeemed and is used to iterate over the raw logs and unpacked data for TransferRedeemed events raised by the Abi contract.
type AbiTransferRedeemedIterator struct {
	Event *AbiTransferRedeemed // Event containing the contract specifics and raw log

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
func (it *AbiTransferRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiTransferRedeemed)
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
		it.Event = new(AbiTransferRedeemed)
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
func (it *AbiTransferRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiTransferRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiTransferRedeemed represents a TransferRedeemed event raised by the Abi contract.
type AbiTransferRedeemed struct {
	Digest [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferRedeemed is a free log retrieval operation binding the contract event 0x504e6efe18ab9eed10dc6501a417f5b12a2f7f2b1593aed9b89f9bce3cf29a91.
//
// Solidity: event TransferRedeemed(bytes32 indexed digest)
func (_Abi *AbiFilterer) FilterTransferRedeemed(opts *bind.FilterOpts, digest [][32]byte) (*AbiTransferRedeemedIterator, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "TransferRedeemed", digestRule)
	if err != nil {
		return nil, err
	}
	return &AbiTransferRedeemedIterator{contract: _Abi.contract, event: "TransferRedeemed", logs: logs, sub: sub}, nil
}

// WatchTransferRedeemed is a free log subscription operation binding the contract event 0x504e6efe18ab9eed10dc6501a417f5b12a2f7f2b1593aed9b89f9bce3cf29a91.
//
// Solidity: event TransferRedeemed(bytes32 indexed digest)
func (_Abi *AbiFilterer) WatchTransferRedeemed(opts *bind.WatchOpts, sink chan<- *AbiTransferRedeemed, digest [][32]byte) (event.Subscription, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "TransferRedeemed", digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiTransferRedeemed)
				if err := _Abi.contract.UnpackLog(event, "TransferRedeemed", log); err != nil {
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

// ParseTransferRedeemed is a log parse operation binding the contract event 0x504e6efe18ab9eed10dc6501a417f5b12a2f7f2b1593aed9b89f9bce3cf29a91.
//
// Solidity: event TransferRedeemed(bytes32 indexed digest)
func (_Abi *AbiFilterer) ParseTransferRedeemed(log types.Log) (*AbiTransferRedeemed, error) {
	event := new(AbiTransferRedeemed)
	if err := _Abi.contract.UnpackLog(event, "TransferRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiTransferSentIterator is returned from FilterTransferSent and is used to iterate over the raw logs and unpacked data for TransferSent events raised by the Abi contract.
type AbiTransferSentIterator struct {
	Event *AbiTransferSent // Event containing the contract specifics and raw log

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
func (it *AbiTransferSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiTransferSent)
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
		it.Event = new(AbiTransferSent)
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
func (it *AbiTransferSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiTransferSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiTransferSent represents a TransferSent event raised by the Abi contract.
type AbiTransferSent struct {
	Recipient      [32]byte
	RefundAddress  [32]byte
	Amount         *big.Int
	Fee            *big.Int
	RecipientChain uint16
	MsgSequence    uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransferSent is a free log retrieval operation binding the contract event 0xe54e51e42099622516fa3b48e9733581c9dbdcb771cafb093f745a0532a35982.
//
// Solidity: event TransferSent(bytes32 indexed recipient, bytes32 indexed refundAddress, uint256 amount, uint256 fee, uint16 recipientChain, uint64 msgSequence)
func (_Abi *AbiFilterer) FilterTransferSent(opts *bind.FilterOpts, recipient [][32]byte, refundAddress [][32]byte) (*AbiTransferSentIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var refundAddressRule []interface{}
	for _, refundAddressItem := range refundAddress {
		refundAddressRule = append(refundAddressRule, refundAddressItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "TransferSent", recipientRule, refundAddressRule)
	if err != nil {
		return nil, err
	}
	return &AbiTransferSentIterator{contract: _Abi.contract, event: "TransferSent", logs: logs, sub: sub}, nil
}

// WatchTransferSent is a free log subscription operation binding the contract event 0xe54e51e42099622516fa3b48e9733581c9dbdcb771cafb093f745a0532a35982.
//
// Solidity: event TransferSent(bytes32 indexed recipient, bytes32 indexed refundAddress, uint256 amount, uint256 fee, uint16 recipientChain, uint64 msgSequence)
func (_Abi *AbiFilterer) WatchTransferSent(opts *bind.WatchOpts, sink chan<- *AbiTransferSent, recipient [][32]byte, refundAddress [][32]byte) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var refundAddressRule []interface{}
	for _, refundAddressItem := range refundAddress {
		refundAddressRule = append(refundAddressRule, refundAddressItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "TransferSent", recipientRule, refundAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiTransferSent)
				if err := _Abi.contract.UnpackLog(event, "TransferSent", log); err != nil {
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

// ParseTransferSent is a log parse operation binding the contract event 0xe54e51e42099622516fa3b48e9733581c9dbdcb771cafb093f745a0532a35982.
//
// Solidity: event TransferSent(bytes32 indexed recipient, bytes32 indexed refundAddress, uint256 amount, uint256 fee, uint16 recipientChain, uint64 msgSequence)
func (_Abi *AbiFilterer) ParseTransferSent(log types.Log) (*AbiTransferSent, error) {
	event := new(AbiTransferSent)
	if err := _Abi.contract.UnpackLog(event, "TransferSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiTransferSent0Iterator is returned from FilterTransferSent0 and is used to iterate over the raw logs and unpacked data for TransferSent0 events raised by the Abi contract.
type AbiTransferSent0Iterator struct {
	Event *AbiTransferSent0 // Event containing the contract specifics and raw log

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
func (it *AbiTransferSent0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiTransferSent0)
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
		it.Event = new(AbiTransferSent0)
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
func (it *AbiTransferSent0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiTransferSent0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiTransferSent0 represents a TransferSent0 event raised by the Abi contract.
type AbiTransferSent0 struct {
	Digest [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferSent0 is a free log retrieval operation binding the contract event 0x3e6ae56314c6da8b461d872f41c6d0bb69317b9d0232805aaccfa45df1a16fa0.
//
// Solidity: event TransferSent(bytes32 indexed digest)
func (_Abi *AbiFilterer) FilterTransferSent0(opts *bind.FilterOpts, digest [][32]byte) (*AbiTransferSent0Iterator, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "TransferSent0", digestRule)
	if err != nil {
		return nil, err
	}
	return &AbiTransferSent0Iterator{contract: _Abi.contract, event: "TransferSent0", logs: logs, sub: sub}, nil
}

// WatchTransferSent0 is a free log subscription operation binding the contract event 0x3e6ae56314c6da8b461d872f41c6d0bb69317b9d0232805aaccfa45df1a16fa0.
//
// Solidity: event TransferSent(bytes32 indexed digest)
func (_Abi *AbiFilterer) WatchTransferSent0(opts *bind.WatchOpts, sink chan<- *AbiTransferSent0, digest [][32]byte) (event.Subscription, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "TransferSent0", digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiTransferSent0)
				if err := _Abi.contract.UnpackLog(event, "TransferSent0", log); err != nil {
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

// ParseTransferSent0 is a log parse operation binding the contract event 0x3e6ae56314c6da8b461d872f41c6d0bb69317b9d0232805aaccfa45df1a16fa0.
//
// Solidity: event TransferSent(bytes32 indexed digest)
func (_Abi *AbiFilterer) ParseTransferSent0(log types.Log) (*AbiTransferSent0, error) {
	event := new(AbiTransferSent0)
	if err := _Abi.contract.UnpackLog(event, "TransferSent0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Abi contract.
type AbiUpgradedIterator struct {
	Event *AbiUpgraded // Event containing the contract specifics and raw log

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
func (it *AbiUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiUpgraded)
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
		it.Event = new(AbiUpgraded)
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
func (it *AbiUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiUpgraded represents a Upgraded event raised by the Abi contract.
type AbiUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Abi *AbiFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AbiUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AbiUpgradedIterator{contract: _Abi.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Abi *AbiFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AbiUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiUpgraded)
				if err := _Abi.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Abi *AbiFilterer) ParseUpgraded(log types.Log) (*AbiUpgraded, error) {
	event := new(AbiUpgraded)
	if err := _Abi.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiWrapFailedIterator is returned from FilterWrapFailed and is used to iterate over the raw logs and unpacked data for WrapFailed events raised by the Abi contract.
type AbiWrapFailedIterator struct {
	Event *AbiWrapFailed // Event containing the contract specifics and raw log

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
func (it *AbiWrapFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiWrapFailed)
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
		it.Event = new(AbiWrapFailed)
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
func (it *AbiWrapFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiWrapFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiWrapFailed represents a WrapFailed event raised by the Abi contract.
type AbiWrapFailed struct {
	DestinationWrappedToken common.Address
	Recipient               common.Address
	Amount                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterWrapFailed is a free log retrieval operation binding the contract event 0x5e484dc77b9161908da1d2f0da5131cdaabac752ae7f0dd633ec8905627b5189.
//
// Solidity: event WrapFailed(address indexed destinationWrappedToken, address indexed recipient, uint256 amount)
func (_Abi *AbiFilterer) FilterWrapFailed(opts *bind.FilterOpts, destinationWrappedToken []common.Address, recipient []common.Address) (*AbiWrapFailedIterator, error) {

	var destinationWrappedTokenRule []interface{}
	for _, destinationWrappedTokenItem := range destinationWrappedToken {
		destinationWrappedTokenRule = append(destinationWrappedTokenRule, destinationWrappedTokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "WrapFailed", destinationWrappedTokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &AbiWrapFailedIterator{contract: _Abi.contract, event: "WrapFailed", logs: logs, sub: sub}, nil
}

// WatchWrapFailed is a free log subscription operation binding the contract event 0x5e484dc77b9161908da1d2f0da5131cdaabac752ae7f0dd633ec8905627b5189.
//
// Solidity: event WrapFailed(address indexed destinationWrappedToken, address indexed recipient, uint256 amount)
func (_Abi *AbiFilterer) WatchWrapFailed(opts *bind.WatchOpts, sink chan<- *AbiWrapFailed, destinationWrappedToken []common.Address, recipient []common.Address) (event.Subscription, error) {

	var destinationWrappedTokenRule []interface{}
	for _, destinationWrappedTokenItem := range destinationWrappedToken {
		destinationWrappedTokenRule = append(destinationWrappedTokenRule, destinationWrappedTokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "WrapFailed", destinationWrappedTokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiWrapFailed)
				if err := _Abi.contract.UnpackLog(event, "WrapFailed", log); err != nil {
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

// ParseWrapFailed is a log parse operation binding the contract event 0x5e484dc77b9161908da1d2f0da5131cdaabac752ae7f0dd633ec8905627b5189.
//
// Solidity: event WrapFailed(address indexed destinationWrappedToken, address indexed recipient, uint256 amount)
func (_Abi *AbiFilterer) ParseWrapFailed(log types.Log) (*AbiWrapFailed, error) {
	event := new(AbiWrapFailed)
	if err := _Abi.contract.UnpackLog(event, "WrapFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
