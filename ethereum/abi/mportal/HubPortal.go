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

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"mToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registrar_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId_\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_TRANSCEIVER_INSTRUCTIONS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NTT_MANAGER_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"attestationReceived\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sourceNttManagerAddress\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"tuple\",\"internalType\":\"structTransceiverStructs.NttManagerMessage\",\"components\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOutboundQueuedTransfer\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"completeInboundQueuedTransfer\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"completeOutboundQueuedTransfer\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"currentIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableEarning\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableEarningIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"enableEarning\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeMsg\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sourceNttManagerAddress\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structTransceiverStructs.NttManagerMessage\",\"components\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentInboundCapacity\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundCapacity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getInboundLimitParams\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRateLimiter.RateLimitParams\",\"components\":[{\"name\":\"limit\",\"type\":\"uint72\",\"internalType\":\"TrimmedAmount\"},{\"name\":\"currentCapacity\",\"type\":\"uint72\",\"internalType\":\"TrimmedAmount\"},{\"name\":\"lastTxTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getInboundQueuedTransfer\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRateLimiter.InboundQueuedTransfer\",\"components\":[{\"name\":\"amount\",\"type\":\"uint72\",\"internalType\":\"TrimmedAmount\"},{\"name\":\"txTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getMigratesImmutables\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOutboundLimitParams\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRateLimiter.RateLimitParams\",\"components\":[{\"name\":\"limit\",\"type\":\"uint72\",\"internalType\":\"TrimmedAmount\"},{\"name\":\"currentCapacity\",\"type\":\"uint72\",\"internalType\":\"TrimmedAmount\"},{\"name\":\"lastTxTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getOutboundQueuedTransfer\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRateLimiter.OutboundQueuedTransfer\",\"components\":[{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"refundAddress\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint72\",\"internalType\":\"TrimmedAmount\"},{\"name\":\"txTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"recipientChain\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transceiverInstructions\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getPeer\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structINttManager.NttManagerPeer\",\"components\":[{\"name\":\"peerAddress\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTransceiverInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structTransceiverRegistry.TransceiverInfo[]\",\"components\":[{\"name\":\"registered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTransceivers\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"isMessageApproved\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMessageExecuted\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageAttestations\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIManagerBase.Mode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextMessageSequence\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauser\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quoteDeliveryPrice\",\"inputs\":[{\"name\":\"recipientChain\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"transceiverInstructions\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rateLimitDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeTransceiver\",\"inputs\":[{\"name\":\"transceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendMTokenIndex\",\"inputs\":[{\"name\":\"destinationChainId_\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"refundAddress_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"messageId_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sendRegistrarKey\",\"inputs\":[{\"name\":\"destinationChainId_\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"key_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"refundAddress_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"messageId_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sendRegistrarListStatus\",\"inputs\":[{\"name\":\"destinationChainId_\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"listName_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"refundAddress_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"messageId_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setInboundLimit\",\"inputs\":[{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId_\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOutboundLimit\",\"inputs\":[{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPeer\",\"inputs\":[{\"name\":\"peerChainId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"peerContract\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"inboundLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThreshold\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTransceiver\",\"inputs\":[{\"name\":\"transceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transceiverAttestedToMessage\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientChain\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientChain\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"refundAddress\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shouldQueue\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"transceiverInstructions\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferPauserCapability\",\"inputs\":[{\"name\":\"newPauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgrade\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"wasEarningEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EarningDisabled\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EarningEnabled\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundTransferQueued\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MTokenIndexSent\",\"inputs\":[{\"name\":\"destinationChainId\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MTokenReceived\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MTokenSent\",\"inputs\":[{\"name\":\"destinationChainId\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageAlreadyExecuted\",\"inputs\":[{\"name\":\"sourceNttManager\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"msgHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageAttestedTo\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"transceiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NotPaused\",\"inputs\":[{\"name\":\"notPaused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundTransferCancelled\",\"inputs\":[{\"name\":\"sequence\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundTransferQueued\",\"inputs\":[{\"name\":\"queueSequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundTransferRateLimited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"currentCapacity\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"paused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserTransferred\",\"inputs\":[{\"name\":\"oldPauser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPauser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PeerUpdated\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"oldPeerContract\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"oldPeerDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"peerContract\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"peerDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrarKeySent\",\"inputs\":[{\"name\":\"destinationChainId\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrarListStatusSent\",\"inputs\":[{\"name\":\"destinationChainId\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"listName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdChanged\",\"inputs\":[{\"name\":\"oldThreshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransceiverAdded\",\"inputs\":[{\"name\":\"transceiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"transceiversNum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransceiverRemoved\",\"inputs\":[{\"name\":\"transceiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferRedeemed\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"refundAddress\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipientChain\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"msgSequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferSent\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BurnAmountDifferentThanBalanceDiff\",\"inputs\":[{\"name\":\"burnAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balanceDiff\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CallerNotTransceiver\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CancellerNotSender\",\"inputs\":[{\"name\":\"canceller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CapacityCannotExceedLimit\",\"inputs\":[{\"name\":\"newCurrentCapacity\",\"type\":\"uint72\",\"internalType\":\"TrimmedAmount\"},{\"name\":\"newLimit\",\"type\":\"uint72\",\"internalType\":\"TrimmedAmount\"}]},{\"type\":\"error\",\"name\":\"DeliveryPaymentTooLow\",\"inputs\":[{\"name\":\"requiredPayment\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"providedPayment\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"DisabledTransceiver\",\"inputs\":[{\"name\":\"transceiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EarningCannotBeReenabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EarningIsDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EarningIsEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InboundQueuedTransferNotFound\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InboundQueuedTransferStillQueued\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transferTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidFork\",\"inputs\":[{\"name\":\"evmChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidFork\",\"inputs\":[{\"name\":\"evmChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMode\",\"inputs\":[{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidPauser\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidPayloadLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidPayloadPrefix\",\"inputs\":[{\"name\":\"prefix\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidPeer\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"peerAddress\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidPeerChainIdZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPeerDecimals\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPeerSameChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPeerZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRecipient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRefundAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTargetChain\",\"inputs\":[{\"name\":\"targetChain\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"thisChain\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidTransceiverZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsApprovedEarner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[{\"name\":\"encodedLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MessageNotApproved\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NoEnabledTransceivers\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonRegisteredTransceiver\",\"inputs\":[{\"name\":\"transceiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotApprovedEarner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughCapacity\",\"inputs\":[{\"name\":\"currentCapacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotImplemented\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMigrating\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegateCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutboundQueuedTransferNotFound\",\"inputs\":[{\"name\":\"queueSequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OutboundQueuedTransferStillQueued\",\"inputs\":[{\"name\":\"queueSequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"transferTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PeerNotRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RefundFailed\",\"inputs\":[{\"name\":\"refundAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RequireContractIsNotPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequireContractIsPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RetrievedIncorrectRegisteredTransceivers\",\"inputs\":[{\"name\":\"retrieved\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"registered\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"StaticcallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ThresholdTooHigh\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transceivers\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TooManyTransceivers\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransceiverAlreadyAttestedToMessage\",\"inputs\":[{\"name\":\"nttManagerMessageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"TransceiverAlreadyEnabled\",\"inputs\":[{\"name\":\"transceiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TransferAmountHasDust\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dust\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Uint64Overflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UndefinedRateLimiting\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedDeployer\",\"inputs\":[{\"name\":\"expectedOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedMsgValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroMToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroRegistrar\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroThreshold\",\"inputs\":[]}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTTRANSCEIVERINSTRUCTIONS is a free data retrieval call binding the contract method 0x74707407.
//
// Solidity: function DEFAULT_TRANSCEIVER_INSTRUCTIONS() view returns(bytes)
func (_Bindings *BindingsCaller) DEFAULTTRANSCEIVERINSTRUCTIONS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "DEFAULT_TRANSCEIVER_INSTRUCTIONS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DEFAULTTRANSCEIVERINSTRUCTIONS is a free data retrieval call binding the contract method 0x74707407.
//
// Solidity: function DEFAULT_TRANSCEIVER_INSTRUCTIONS() view returns(bytes)
func (_Bindings *BindingsSession) DEFAULTTRANSCEIVERINSTRUCTIONS() ([]byte, error) {
	return _Bindings.Contract.DEFAULTTRANSCEIVERINSTRUCTIONS(&_Bindings.CallOpts)
}

// DEFAULTTRANSCEIVERINSTRUCTIONS is a free data retrieval call binding the contract method 0x74707407.
//
// Solidity: function DEFAULT_TRANSCEIVER_INSTRUCTIONS() view returns(bytes)
func (_Bindings *BindingsCallerSession) DEFAULTTRANSCEIVERINSTRUCTIONS() ([]byte, error) {
	return _Bindings.Contract.DEFAULTTRANSCEIVERINSTRUCTIONS(&_Bindings.CallOpts)
}

// NTTMANAGERVERSION is a free data retrieval call binding the contract method 0xc0b07bde.
//
// Solidity: function NTT_MANAGER_VERSION() view returns(string)
func (_Bindings *BindingsCaller) NTTMANAGERVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "NTT_MANAGER_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NTTMANAGERVERSION is a free data retrieval call binding the contract method 0xc0b07bde.
//
// Solidity: function NTT_MANAGER_VERSION() view returns(string)
func (_Bindings *BindingsSession) NTTMANAGERVERSION() (string, error) {
	return _Bindings.Contract.NTTMANAGERVERSION(&_Bindings.CallOpts)
}

// NTTMANAGERVERSION is a free data retrieval call binding the contract method 0xc0b07bde.
//
// Solidity: function NTT_MANAGER_VERSION() view returns(string)
func (_Bindings *BindingsCallerSession) NTTMANAGERVERSION() (string, error) {
	return _Bindings.Contract.NTTMANAGERVERSION(&_Bindings.CallOpts)
}

// CancelOutboundQueuedTransfer is a free data retrieval call binding the contract method 0xf7514fbc.
//
// Solidity: function cancelOutboundQueuedTransfer(uint64 ) view returns()
func (_Bindings *BindingsCaller) CancelOutboundQueuedTransfer(opts *bind.CallOpts, arg0 uint64) error {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "cancelOutboundQueuedTransfer", arg0)

	if err != nil {
		return err
	}

	return err

}

// CancelOutboundQueuedTransfer is a free data retrieval call binding the contract method 0xf7514fbc.
//
// Solidity: function cancelOutboundQueuedTransfer(uint64 ) view returns()
func (_Bindings *BindingsSession) CancelOutboundQueuedTransfer(arg0 uint64) error {
	return _Bindings.Contract.CancelOutboundQueuedTransfer(&_Bindings.CallOpts, arg0)
}

// CancelOutboundQueuedTransfer is a free data retrieval call binding the contract method 0xf7514fbc.
//
// Solidity: function cancelOutboundQueuedTransfer(uint64 ) view returns()
func (_Bindings *BindingsCallerSession) CancelOutboundQueuedTransfer(arg0 uint64) error {
	return _Bindings.Contract.CancelOutboundQueuedTransfer(&_Bindings.CallOpts, arg0)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Bindings *BindingsCaller) ChainId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Bindings *BindingsSession) ChainId() (uint16, error) {
	return _Bindings.Contract.ChainId(&_Bindings.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Bindings *BindingsCallerSession) ChainId() (uint16, error) {
	return _Bindings.Contract.ChainId(&_Bindings.CallOpts)
}

// CompleteInboundQueuedTransfer is a free data retrieval call binding the contract method 0x8413bcba.
//
// Solidity: function completeInboundQueuedTransfer(bytes32 ) view returns()
func (_Bindings *BindingsCaller) CompleteInboundQueuedTransfer(opts *bind.CallOpts, arg0 [32]byte) error {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "completeInboundQueuedTransfer", arg0)

	if err != nil {
		return err
	}

	return err

}

// CompleteInboundQueuedTransfer is a free data retrieval call binding the contract method 0x8413bcba.
//
// Solidity: function completeInboundQueuedTransfer(bytes32 ) view returns()
func (_Bindings *BindingsSession) CompleteInboundQueuedTransfer(arg0 [32]byte) error {
	return _Bindings.Contract.CompleteInboundQueuedTransfer(&_Bindings.CallOpts, arg0)
}

// CompleteInboundQueuedTransfer is a free data retrieval call binding the contract method 0x8413bcba.
//
// Solidity: function completeInboundQueuedTransfer(bytes32 ) view returns()
func (_Bindings *BindingsCallerSession) CompleteInboundQueuedTransfer(arg0 [32]byte) error {
	return _Bindings.Contract.CompleteInboundQueuedTransfer(&_Bindings.CallOpts, arg0)
}

// CurrentIndex is a free data retrieval call binding the contract method 0x26987b60.
//
// Solidity: function currentIndex() view returns(uint128)
func (_Bindings *BindingsCaller) CurrentIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "currentIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentIndex is a free data retrieval call binding the contract method 0x26987b60.
//
// Solidity: function currentIndex() view returns(uint128)
func (_Bindings *BindingsSession) CurrentIndex() (*big.Int, error) {
	return _Bindings.Contract.CurrentIndex(&_Bindings.CallOpts)
}

// CurrentIndex is a free data retrieval call binding the contract method 0x26987b60.
//
// Solidity: function currentIndex() view returns(uint128)
func (_Bindings *BindingsCallerSession) CurrentIndex() (*big.Int, error) {
	return _Bindings.Contract.CurrentIndex(&_Bindings.CallOpts)
}

// DisableEarningIndex is a free data retrieval call binding the contract method 0xf22e63ac.
//
// Solidity: function disableEarningIndex() view returns(uint128)
func (_Bindings *BindingsCaller) DisableEarningIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "disableEarningIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisableEarningIndex is a free data retrieval call binding the contract method 0xf22e63ac.
//
// Solidity: function disableEarningIndex() view returns(uint128)
func (_Bindings *BindingsSession) DisableEarningIndex() (*big.Int, error) {
	return _Bindings.Contract.DisableEarningIndex(&_Bindings.CallOpts)
}

// DisableEarningIndex is a free data retrieval call binding the contract method 0xf22e63ac.
//
// Solidity: function disableEarningIndex() view returns(uint128)
func (_Bindings *BindingsCallerSession) DisableEarningIndex() (*big.Int, error) {
	return _Bindings.Contract.DisableEarningIndex(&_Bindings.CallOpts)
}

// GetCurrentInboundCapacity is a free data retrieval call binding the contract method 0x02717250.
//
// Solidity: function getCurrentInboundCapacity(uint16 ) pure returns(uint256)
func (_Bindings *BindingsCaller) GetCurrentInboundCapacity(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getCurrentInboundCapacity", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentInboundCapacity is a free data retrieval call binding the contract method 0x02717250.
//
// Solidity: function getCurrentInboundCapacity(uint16 ) pure returns(uint256)
func (_Bindings *BindingsSession) GetCurrentInboundCapacity(arg0 uint16) (*big.Int, error) {
	return _Bindings.Contract.GetCurrentInboundCapacity(&_Bindings.CallOpts, arg0)
}

// GetCurrentInboundCapacity is a free data retrieval call binding the contract method 0x02717250.
//
// Solidity: function getCurrentInboundCapacity(uint16 ) pure returns(uint256)
func (_Bindings *BindingsCallerSession) GetCurrentInboundCapacity(arg0 uint16) (*big.Int, error) {
	return _Bindings.Contract.GetCurrentInboundCapacity(&_Bindings.CallOpts, arg0)
}

// GetCurrentOutboundCapacity is a free data retrieval call binding the contract method 0xf5cfec18.
//
// Solidity: function getCurrentOutboundCapacity() pure returns(uint256)
func (_Bindings *BindingsCaller) GetCurrentOutboundCapacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getCurrentOutboundCapacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentOutboundCapacity is a free data retrieval call binding the contract method 0xf5cfec18.
//
// Solidity: function getCurrentOutboundCapacity() pure returns(uint256)
func (_Bindings *BindingsSession) GetCurrentOutboundCapacity() (*big.Int, error) {
	return _Bindings.Contract.GetCurrentOutboundCapacity(&_Bindings.CallOpts)
}

// GetCurrentOutboundCapacity is a free data retrieval call binding the contract method 0xf5cfec18.
//
// Solidity: function getCurrentOutboundCapacity() pure returns(uint256)
func (_Bindings *BindingsCallerSession) GetCurrentOutboundCapacity() (*big.Int, error) {
	return _Bindings.Contract.GetCurrentOutboundCapacity(&_Bindings.CallOpts)
}

// GetInboundLimitParams is a free data retrieval call binding the contract method 0xd788c147.
//
// Solidity: function getInboundLimitParams(uint16 ) pure returns((uint72,uint72,uint64))
func (_Bindings *BindingsCaller) GetInboundLimitParams(opts *bind.CallOpts, arg0 uint16) (IRateLimiterRateLimitParams, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getInboundLimitParams", arg0)

	if err != nil {
		return *new(IRateLimiterRateLimitParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IRateLimiterRateLimitParams)).(*IRateLimiterRateLimitParams)

	return out0, err

}

// GetInboundLimitParams is a free data retrieval call binding the contract method 0xd788c147.
//
// Solidity: function getInboundLimitParams(uint16 ) pure returns((uint72,uint72,uint64))
func (_Bindings *BindingsSession) GetInboundLimitParams(arg0 uint16) (IRateLimiterRateLimitParams, error) {
	return _Bindings.Contract.GetInboundLimitParams(&_Bindings.CallOpts, arg0)
}

// GetInboundLimitParams is a free data retrieval call binding the contract method 0xd788c147.
//
// Solidity: function getInboundLimitParams(uint16 ) pure returns((uint72,uint72,uint64))
func (_Bindings *BindingsCallerSession) GetInboundLimitParams(arg0 uint16) (IRateLimiterRateLimitParams, error) {
	return _Bindings.Contract.GetInboundLimitParams(&_Bindings.CallOpts, arg0)
}

// GetInboundQueuedTransfer is a free data retrieval call binding the contract method 0xfd96063c.
//
// Solidity: function getInboundQueuedTransfer(bytes32 ) pure returns((uint72,uint64,address))
func (_Bindings *BindingsCaller) GetInboundQueuedTransfer(opts *bind.CallOpts, arg0 [32]byte) (IRateLimiterInboundQueuedTransfer, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getInboundQueuedTransfer", arg0)

	if err != nil {
		return *new(IRateLimiterInboundQueuedTransfer), err
	}

	out0 := *abi.ConvertType(out[0], new(IRateLimiterInboundQueuedTransfer)).(*IRateLimiterInboundQueuedTransfer)

	return out0, err

}

// GetInboundQueuedTransfer is a free data retrieval call binding the contract method 0xfd96063c.
//
// Solidity: function getInboundQueuedTransfer(bytes32 ) pure returns((uint72,uint64,address))
func (_Bindings *BindingsSession) GetInboundQueuedTransfer(arg0 [32]byte) (IRateLimiterInboundQueuedTransfer, error) {
	return _Bindings.Contract.GetInboundQueuedTransfer(&_Bindings.CallOpts, arg0)
}

// GetInboundQueuedTransfer is a free data retrieval call binding the contract method 0xfd96063c.
//
// Solidity: function getInboundQueuedTransfer(bytes32 ) pure returns((uint72,uint64,address))
func (_Bindings *BindingsCallerSession) GetInboundQueuedTransfer(arg0 [32]byte) (IRateLimiterInboundQueuedTransfer, error) {
	return _Bindings.Contract.GetInboundQueuedTransfer(&_Bindings.CallOpts, arg0)
}

// GetMigratesImmutables is a free data retrieval call binding the contract method 0x689f90c3.
//
// Solidity: function getMigratesImmutables() view returns(bool)
func (_Bindings *BindingsCaller) GetMigratesImmutables(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getMigratesImmutables")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMigratesImmutables is a free data retrieval call binding the contract method 0x689f90c3.
//
// Solidity: function getMigratesImmutables() view returns(bool)
func (_Bindings *BindingsSession) GetMigratesImmutables() (bool, error) {
	return _Bindings.Contract.GetMigratesImmutables(&_Bindings.CallOpts)
}

// GetMigratesImmutables is a free data retrieval call binding the contract method 0x689f90c3.
//
// Solidity: function getMigratesImmutables() view returns(bool)
func (_Bindings *BindingsCallerSession) GetMigratesImmutables() (bool, error) {
	return _Bindings.Contract.GetMigratesImmutables(&_Bindings.CallOpts)
}

// GetMode is a free data retrieval call binding the contract method 0x4b4fd03b.
//
// Solidity: function getMode() view returns(uint8)
func (_Bindings *BindingsCaller) GetMode(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getMode")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetMode is a free data retrieval call binding the contract method 0x4b4fd03b.
//
// Solidity: function getMode() view returns(uint8)
func (_Bindings *BindingsSession) GetMode() (uint8, error) {
	return _Bindings.Contract.GetMode(&_Bindings.CallOpts)
}

// GetMode is a free data retrieval call binding the contract method 0x4b4fd03b.
//
// Solidity: function getMode() view returns(uint8)
func (_Bindings *BindingsCallerSession) GetMode() (uint8, error) {
	return _Bindings.Contract.GetMode(&_Bindings.CallOpts)
}

// GetOutboundLimitParams is a free data retrieval call binding the contract method 0x86e11ffa.
//
// Solidity: function getOutboundLimitParams() pure returns((uint72,uint72,uint64))
func (_Bindings *BindingsCaller) GetOutboundLimitParams(opts *bind.CallOpts) (IRateLimiterRateLimitParams, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getOutboundLimitParams")

	if err != nil {
		return *new(IRateLimiterRateLimitParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IRateLimiterRateLimitParams)).(*IRateLimiterRateLimitParams)

	return out0, err

}

// GetOutboundLimitParams is a free data retrieval call binding the contract method 0x86e11ffa.
//
// Solidity: function getOutboundLimitParams() pure returns((uint72,uint72,uint64))
func (_Bindings *BindingsSession) GetOutboundLimitParams() (IRateLimiterRateLimitParams, error) {
	return _Bindings.Contract.GetOutboundLimitParams(&_Bindings.CallOpts)
}

// GetOutboundLimitParams is a free data retrieval call binding the contract method 0x86e11ffa.
//
// Solidity: function getOutboundLimitParams() pure returns((uint72,uint72,uint64))
func (_Bindings *BindingsCallerSession) GetOutboundLimitParams() (IRateLimiterRateLimitParams, error) {
	return _Bindings.Contract.GetOutboundLimitParams(&_Bindings.CallOpts)
}

// GetOutboundQueuedTransfer is a free data retrieval call binding the contract method 0x1f97c9a8.
//
// Solidity: function getOutboundQueuedTransfer(uint64 ) pure returns((bytes32,bytes32,uint72,uint64,uint16,address,bytes))
func (_Bindings *BindingsCaller) GetOutboundQueuedTransfer(opts *bind.CallOpts, arg0 uint64) (IRateLimiterOutboundQueuedTransfer, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getOutboundQueuedTransfer", arg0)

	if err != nil {
		return *new(IRateLimiterOutboundQueuedTransfer), err
	}

	out0 := *abi.ConvertType(out[0], new(IRateLimiterOutboundQueuedTransfer)).(*IRateLimiterOutboundQueuedTransfer)

	return out0, err

}

// GetOutboundQueuedTransfer is a free data retrieval call binding the contract method 0x1f97c9a8.
//
// Solidity: function getOutboundQueuedTransfer(uint64 ) pure returns((bytes32,bytes32,uint72,uint64,uint16,address,bytes))
func (_Bindings *BindingsSession) GetOutboundQueuedTransfer(arg0 uint64) (IRateLimiterOutboundQueuedTransfer, error) {
	return _Bindings.Contract.GetOutboundQueuedTransfer(&_Bindings.CallOpts, arg0)
}

// GetOutboundQueuedTransfer is a free data retrieval call binding the contract method 0x1f97c9a8.
//
// Solidity: function getOutboundQueuedTransfer(uint64 ) pure returns((bytes32,bytes32,uint72,uint64,uint16,address,bytes))
func (_Bindings *BindingsCallerSession) GetOutboundQueuedTransfer(arg0 uint64) (IRateLimiterOutboundQueuedTransfer, error) {
	return _Bindings.Contract.GetOutboundQueuedTransfer(&_Bindings.CallOpts, arg0)
}

// GetPeer is a free data retrieval call binding the contract method 0xc128d170.
//
// Solidity: function getPeer(uint16 chainId_) view returns((bytes32,uint8))
func (_Bindings *BindingsCaller) GetPeer(opts *bind.CallOpts, chainId_ uint16) (INttManagerNttManagerPeer, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getPeer", chainId_)

	if err != nil {
		return *new(INttManagerNttManagerPeer), err
	}

	out0 := *abi.ConvertType(out[0], new(INttManagerNttManagerPeer)).(*INttManagerNttManagerPeer)

	return out0, err

}

// GetPeer is a free data retrieval call binding the contract method 0xc128d170.
//
// Solidity: function getPeer(uint16 chainId_) view returns((bytes32,uint8))
func (_Bindings *BindingsSession) GetPeer(chainId_ uint16) (INttManagerNttManagerPeer, error) {
	return _Bindings.Contract.GetPeer(&_Bindings.CallOpts, chainId_)
}

// GetPeer is a free data retrieval call binding the contract method 0xc128d170.
//
// Solidity: function getPeer(uint16 chainId_) view returns((bytes32,uint8))
func (_Bindings *BindingsCallerSession) GetPeer(chainId_ uint16) (INttManagerNttManagerPeer, error) {
	return _Bindings.Contract.GetPeer(&_Bindings.CallOpts, chainId_)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint8)
func (_Bindings *BindingsCaller) GetThreshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint8)
func (_Bindings *BindingsSession) GetThreshold() (uint8, error) {
	return _Bindings.Contract.GetThreshold(&_Bindings.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint8)
func (_Bindings *BindingsCallerSession) GetThreshold() (uint8, error) {
	return _Bindings.Contract.GetThreshold(&_Bindings.CallOpts)
}

// GetTransceiverInfo is a free data retrieval call binding the contract method 0xb150fc55.
//
// Solidity: function getTransceiverInfo() view returns((bool,bool,uint8)[])
func (_Bindings *BindingsCaller) GetTransceiverInfo(opts *bind.CallOpts) ([]TransceiverRegistryTransceiverInfo, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getTransceiverInfo")

	if err != nil {
		return *new([]TransceiverRegistryTransceiverInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransceiverRegistryTransceiverInfo)).(*[]TransceiverRegistryTransceiverInfo)

	return out0, err

}

// GetTransceiverInfo is a free data retrieval call binding the contract method 0xb150fc55.
//
// Solidity: function getTransceiverInfo() view returns((bool,bool,uint8)[])
func (_Bindings *BindingsSession) GetTransceiverInfo() ([]TransceiverRegistryTransceiverInfo, error) {
	return _Bindings.Contract.GetTransceiverInfo(&_Bindings.CallOpts)
}

// GetTransceiverInfo is a free data retrieval call binding the contract method 0xb150fc55.
//
// Solidity: function getTransceiverInfo() view returns((bool,bool,uint8)[])
func (_Bindings *BindingsCallerSession) GetTransceiverInfo() ([]TransceiverRegistryTransceiverInfo, error) {
	return _Bindings.Contract.GetTransceiverInfo(&_Bindings.CallOpts)
}

// GetTransceivers is a free data retrieval call binding the contract method 0xb4d591bb.
//
// Solidity: function getTransceivers() pure returns(address[] result)
func (_Bindings *BindingsCaller) GetTransceivers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getTransceivers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTransceivers is a free data retrieval call binding the contract method 0xb4d591bb.
//
// Solidity: function getTransceivers() pure returns(address[] result)
func (_Bindings *BindingsSession) GetTransceivers() ([]common.Address, error) {
	return _Bindings.Contract.GetTransceivers(&_Bindings.CallOpts)
}

// GetTransceivers is a free data retrieval call binding the contract method 0xb4d591bb.
//
// Solidity: function getTransceivers() pure returns(address[] result)
func (_Bindings *BindingsCallerSession) GetTransceivers() ([]common.Address, error) {
	return _Bindings.Contract.GetTransceivers(&_Bindings.CallOpts)
}

// IsMessageApproved is a free data retrieval call binding the contract method 0x0677df54.
//
// Solidity: function isMessageApproved(bytes32 digest) view returns(bool)
func (_Bindings *BindingsCaller) IsMessageApproved(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isMessageApproved", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessageApproved is a free data retrieval call binding the contract method 0x0677df54.
//
// Solidity: function isMessageApproved(bytes32 digest) view returns(bool)
func (_Bindings *BindingsSession) IsMessageApproved(digest [32]byte) (bool, error) {
	return _Bindings.Contract.IsMessageApproved(&_Bindings.CallOpts, digest)
}

// IsMessageApproved is a free data retrieval call binding the contract method 0x0677df54.
//
// Solidity: function isMessageApproved(bytes32 digest) view returns(bool)
func (_Bindings *BindingsCallerSession) IsMessageApproved(digest [32]byte) (bool, error) {
	return _Bindings.Contract.IsMessageApproved(&_Bindings.CallOpts, digest)
}

// IsMessageExecuted is a free data retrieval call binding the contract method 0x396c16b7.
//
// Solidity: function isMessageExecuted(bytes32 digest) view returns(bool)
func (_Bindings *BindingsCaller) IsMessageExecuted(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isMessageExecuted", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessageExecuted is a free data retrieval call binding the contract method 0x396c16b7.
//
// Solidity: function isMessageExecuted(bytes32 digest) view returns(bool)
func (_Bindings *BindingsSession) IsMessageExecuted(digest [32]byte) (bool, error) {
	return _Bindings.Contract.IsMessageExecuted(&_Bindings.CallOpts, digest)
}

// IsMessageExecuted is a free data retrieval call binding the contract method 0x396c16b7.
//
// Solidity: function isMessageExecuted(bytes32 digest) view returns(bool)
func (_Bindings *BindingsCallerSession) IsMessageExecuted(digest [32]byte) (bool, error) {
	return _Bindings.Contract.IsMessageExecuted(&_Bindings.CallOpts, digest)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Bindings *BindingsCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Bindings *BindingsSession) IsPaused() (bool, error) {
	return _Bindings.Contract.IsPaused(&_Bindings.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Bindings *BindingsCallerSession) IsPaused() (bool, error) {
	return _Bindings.Contract.IsPaused(&_Bindings.CallOpts)
}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() view returns(address)
func (_Bindings *BindingsCaller) MToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "mToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() view returns(address)
func (_Bindings *BindingsSession) MToken() (common.Address, error) {
	return _Bindings.Contract.MToken(&_Bindings.CallOpts)
}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() view returns(address)
func (_Bindings *BindingsCallerSession) MToken() (common.Address, error) {
	return _Bindings.Contract.MToken(&_Bindings.CallOpts)
}

// MessageAttestations is a free data retrieval call binding the contract method 0x89c619dd.
//
// Solidity: function messageAttestations(bytes32 digest) view returns(uint8 count)
func (_Bindings *BindingsCaller) MessageAttestations(opts *bind.CallOpts, digest [32]byte) (uint8, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "messageAttestations", digest)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MessageAttestations is a free data retrieval call binding the contract method 0x89c619dd.
//
// Solidity: function messageAttestations(bytes32 digest) view returns(uint8 count)
func (_Bindings *BindingsSession) MessageAttestations(digest [32]byte) (uint8, error) {
	return _Bindings.Contract.MessageAttestations(&_Bindings.CallOpts, digest)
}

// MessageAttestations is a free data retrieval call binding the contract method 0x89c619dd.
//
// Solidity: function messageAttestations(bytes32 digest) view returns(uint8 count)
func (_Bindings *BindingsCallerSession) MessageAttestations(digest [32]byte) (uint8, error) {
	return _Bindings.Contract.MessageAttestations(&_Bindings.CallOpts, digest)
}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(uint8)
func (_Bindings *BindingsCaller) Mode(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "mode")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(uint8)
func (_Bindings *BindingsSession) Mode() (uint8, error) {
	return _Bindings.Contract.Mode(&_Bindings.CallOpts)
}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(uint8)
func (_Bindings *BindingsCallerSession) Mode() (uint8, error) {
	return _Bindings.Contract.Mode(&_Bindings.CallOpts)
}

// NextMessageSequence is a free data retrieval call binding the contract method 0x23d75e31.
//
// Solidity: function nextMessageSequence() view returns(uint64)
func (_Bindings *BindingsCaller) NextMessageSequence(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "nextMessageSequence")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextMessageSequence is a free data retrieval call binding the contract method 0x23d75e31.
//
// Solidity: function nextMessageSequence() view returns(uint64)
func (_Bindings *BindingsSession) NextMessageSequence() (uint64, error) {
	return _Bindings.Contract.NextMessageSequence(&_Bindings.CallOpts)
}

// NextMessageSequence is a free data retrieval call binding the contract method 0x23d75e31.
//
// Solidity: function nextMessageSequence() view returns(uint64)
func (_Bindings *BindingsCallerSession) NextMessageSequence() (uint64, error) {
	return _Bindings.Contract.NextMessageSequence(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCallerSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Bindings *BindingsCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Bindings *BindingsSession) Pauser() (common.Address, error) {
	return _Bindings.Contract.Pauser(&_Bindings.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Bindings *BindingsCallerSession) Pauser() (common.Address, error) {
	return _Bindings.Contract.Pauser(&_Bindings.CallOpts)
}

// QuoteDeliveryPrice is a free data retrieval call binding the contract method 0x9057412d.
//
// Solidity: function quoteDeliveryPrice(uint16 recipientChain, bytes transceiverInstructions) view returns(uint256[], uint256)
func (_Bindings *BindingsCaller) QuoteDeliveryPrice(opts *bind.CallOpts, recipientChain uint16, transceiverInstructions []byte) ([]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "quoteDeliveryPrice", recipientChain, transceiverInstructions)

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
func (_Bindings *BindingsSession) QuoteDeliveryPrice(recipientChain uint16, transceiverInstructions []byte) ([]*big.Int, *big.Int, error) {
	return _Bindings.Contract.QuoteDeliveryPrice(&_Bindings.CallOpts, recipientChain, transceiverInstructions)
}

// QuoteDeliveryPrice is a free data retrieval call binding the contract method 0x9057412d.
//
// Solidity: function quoteDeliveryPrice(uint16 recipientChain, bytes transceiverInstructions) view returns(uint256[], uint256)
func (_Bindings *BindingsCallerSession) QuoteDeliveryPrice(recipientChain uint16, transceiverInstructions []byte) ([]*big.Int, *big.Int, error) {
	return _Bindings.Contract.QuoteDeliveryPrice(&_Bindings.CallOpts, recipientChain, transceiverInstructions)
}

// RateLimitDuration is a free data retrieval call binding the contract method 0x74aa7bfc.
//
// Solidity: function rateLimitDuration() view returns(uint64)
func (_Bindings *BindingsCaller) RateLimitDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "rateLimitDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RateLimitDuration is a free data retrieval call binding the contract method 0x74aa7bfc.
//
// Solidity: function rateLimitDuration() view returns(uint64)
func (_Bindings *BindingsSession) RateLimitDuration() (uint64, error) {
	return _Bindings.Contract.RateLimitDuration(&_Bindings.CallOpts)
}

// RateLimitDuration is a free data retrieval call binding the contract method 0x74aa7bfc.
//
// Solidity: function rateLimitDuration() view returns(uint64)
func (_Bindings *BindingsCallerSession) RateLimitDuration() (uint64, error) {
	return _Bindings.Contract.RateLimitDuration(&_Bindings.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_Bindings *BindingsCaller) Registrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "registrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_Bindings *BindingsSession) Registrar() (common.Address, error) {
	return _Bindings.Contract.Registrar(&_Bindings.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_Bindings *BindingsCallerSession) Registrar() (common.Address, error) {
	return _Bindings.Contract.Registrar(&_Bindings.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bindings *BindingsCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bindings *BindingsSession) Token() (common.Address, error) {
	return _Bindings.Contract.Token(&_Bindings.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bindings *BindingsCallerSession) Token() (common.Address, error) {
	return _Bindings.Contract.Token(&_Bindings.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_Bindings *BindingsCaller) TokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "tokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_Bindings *BindingsSession) TokenDecimals() (uint8, error) {
	return _Bindings.Contract.TokenDecimals(&_Bindings.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_Bindings *BindingsCallerSession) TokenDecimals() (uint8, error) {
	return _Bindings.Contract.TokenDecimals(&_Bindings.CallOpts)
}

// TransceiverAttestedToMessage is a free data retrieval call binding the contract method 0x8e3ba8c9.
//
// Solidity: function transceiverAttestedToMessage(bytes32 digest, uint8 index) view returns(bool)
func (_Bindings *BindingsCaller) TransceiverAttestedToMessage(opts *bind.CallOpts, digest [32]byte, index uint8) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "transceiverAttestedToMessage", digest, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransceiverAttestedToMessage is a free data retrieval call binding the contract method 0x8e3ba8c9.
//
// Solidity: function transceiverAttestedToMessage(bytes32 digest, uint8 index) view returns(bool)
func (_Bindings *BindingsSession) TransceiverAttestedToMessage(digest [32]byte, index uint8) (bool, error) {
	return _Bindings.Contract.TransceiverAttestedToMessage(&_Bindings.CallOpts, digest, index)
}

// TransceiverAttestedToMessage is a free data retrieval call binding the contract method 0x8e3ba8c9.
//
// Solidity: function transceiverAttestedToMessage(bytes32 digest, uint8 index) view returns(bool)
func (_Bindings *BindingsCallerSession) TransceiverAttestedToMessage(digest [32]byte, index uint8) (bool, error) {
	return _Bindings.Contract.TransceiverAttestedToMessage(&_Bindings.CallOpts, digest, index)
}

// WasEarningEnabled is a free data retrieval call binding the contract method 0x8a4a1017.
//
// Solidity: function wasEarningEnabled() view returns(bool)
func (_Bindings *BindingsCaller) WasEarningEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "wasEarningEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasEarningEnabled is a free data retrieval call binding the contract method 0x8a4a1017.
//
// Solidity: function wasEarningEnabled() view returns(bool)
func (_Bindings *BindingsSession) WasEarningEnabled() (bool, error) {
	return _Bindings.Contract.WasEarningEnabled(&_Bindings.CallOpts)
}

// WasEarningEnabled is a free data retrieval call binding the contract method 0x8a4a1017.
//
// Solidity: function wasEarningEnabled() view returns(bool)
func (_Bindings *BindingsCallerSession) WasEarningEnabled() (bool, error) {
	return _Bindings.Contract.WasEarningEnabled(&_Bindings.CallOpts)
}

// AttestationReceived is a paid mutator transaction binding the contract method 0x9d782454.
//
// Solidity: function attestationReceived(uint16 sourceChainId, bytes32 sourceNttManagerAddress, (bytes32,bytes32,bytes) payload) returns()
func (_Bindings *BindingsTransactor) AttestationReceived(opts *bind.TransactOpts, sourceChainId uint16, sourceNttManagerAddress [32]byte, payload TransceiverStructsNttManagerMessage) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "attestationReceived", sourceChainId, sourceNttManagerAddress, payload)
}

// AttestationReceived is a paid mutator transaction binding the contract method 0x9d782454.
//
// Solidity: function attestationReceived(uint16 sourceChainId, bytes32 sourceNttManagerAddress, (bytes32,bytes32,bytes) payload) returns()
func (_Bindings *BindingsSession) AttestationReceived(sourceChainId uint16, sourceNttManagerAddress [32]byte, payload TransceiverStructsNttManagerMessage) (*types.Transaction, error) {
	return _Bindings.Contract.AttestationReceived(&_Bindings.TransactOpts, sourceChainId, sourceNttManagerAddress, payload)
}

// AttestationReceived is a paid mutator transaction binding the contract method 0x9d782454.
//
// Solidity: function attestationReceived(uint16 sourceChainId, bytes32 sourceNttManagerAddress, (bytes32,bytes32,bytes) payload) returns()
func (_Bindings *BindingsTransactorSession) AttestationReceived(sourceChainId uint16, sourceNttManagerAddress [32]byte, payload TransceiverStructsNttManagerMessage) (*types.Transaction, error) {
	return _Bindings.Contract.AttestationReceived(&_Bindings.TransactOpts, sourceChainId, sourceNttManagerAddress, payload)
}

// CompleteOutboundQueuedTransfer is a paid mutator transaction binding the contract method 0x97c35146.
//
// Solidity: function completeOutboundQueuedTransfer(uint64 ) payable returns(uint64)
func (_Bindings *BindingsTransactor) CompleteOutboundQueuedTransfer(opts *bind.TransactOpts, arg0 uint64) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "completeOutboundQueuedTransfer", arg0)
}

// CompleteOutboundQueuedTransfer is a paid mutator transaction binding the contract method 0x97c35146.
//
// Solidity: function completeOutboundQueuedTransfer(uint64 ) payable returns(uint64)
func (_Bindings *BindingsSession) CompleteOutboundQueuedTransfer(arg0 uint64) (*types.Transaction, error) {
	return _Bindings.Contract.CompleteOutboundQueuedTransfer(&_Bindings.TransactOpts, arg0)
}

// CompleteOutboundQueuedTransfer is a paid mutator transaction binding the contract method 0x97c35146.
//
// Solidity: function completeOutboundQueuedTransfer(uint64 ) payable returns(uint64)
func (_Bindings *BindingsTransactorSession) CompleteOutboundQueuedTransfer(arg0 uint64) (*types.Transaction, error) {
	return _Bindings.Contract.CompleteOutboundQueuedTransfer(&_Bindings.TransactOpts, arg0)
}

// DisableEarning is a paid mutator transaction binding the contract method 0xa8afc01f.
//
// Solidity: function disableEarning() returns()
func (_Bindings *BindingsTransactor) DisableEarning(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "disableEarning")
}

// DisableEarning is a paid mutator transaction binding the contract method 0xa8afc01f.
//
// Solidity: function disableEarning() returns()
func (_Bindings *BindingsSession) DisableEarning() (*types.Transaction, error) {
	return _Bindings.Contract.DisableEarning(&_Bindings.TransactOpts)
}

// DisableEarning is a paid mutator transaction binding the contract method 0xa8afc01f.
//
// Solidity: function disableEarning() returns()
func (_Bindings *BindingsTransactorSession) DisableEarning() (*types.Transaction, error) {
	return _Bindings.Contract.DisableEarning(&_Bindings.TransactOpts)
}

// EnableEarning is a paid mutator transaction binding the contract method 0xc967891a.
//
// Solidity: function enableEarning() returns()
func (_Bindings *BindingsTransactor) EnableEarning(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "enableEarning")
}

// EnableEarning is a paid mutator transaction binding the contract method 0xc967891a.
//
// Solidity: function enableEarning() returns()
func (_Bindings *BindingsSession) EnableEarning() (*types.Transaction, error) {
	return _Bindings.Contract.EnableEarning(&_Bindings.TransactOpts)
}

// EnableEarning is a paid mutator transaction binding the contract method 0xc967891a.
//
// Solidity: function enableEarning() returns()
func (_Bindings *BindingsTransactorSession) EnableEarning() (*types.Transaction, error) {
	return _Bindings.Contract.EnableEarning(&_Bindings.TransactOpts)
}

// ExecuteMsg is a paid mutator transaction binding the contract method 0xda4856a1.
//
// Solidity: function executeMsg(uint16 sourceChainId, bytes32 sourceNttManagerAddress, (bytes32,bytes32,bytes) message) returns()
func (_Bindings *BindingsTransactor) ExecuteMsg(opts *bind.TransactOpts, sourceChainId uint16, sourceNttManagerAddress [32]byte, message TransceiverStructsNttManagerMessage) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "executeMsg", sourceChainId, sourceNttManagerAddress, message)
}

// ExecuteMsg is a paid mutator transaction binding the contract method 0xda4856a1.
//
// Solidity: function executeMsg(uint16 sourceChainId, bytes32 sourceNttManagerAddress, (bytes32,bytes32,bytes) message) returns()
func (_Bindings *BindingsSession) ExecuteMsg(sourceChainId uint16, sourceNttManagerAddress [32]byte, message TransceiverStructsNttManagerMessage) (*types.Transaction, error) {
	return _Bindings.Contract.ExecuteMsg(&_Bindings.TransactOpts, sourceChainId, sourceNttManagerAddress, message)
}

// ExecuteMsg is a paid mutator transaction binding the contract method 0xda4856a1.
//
// Solidity: function executeMsg(uint16 sourceChainId, bytes32 sourceNttManagerAddress, (bytes32,bytes32,bytes) message) returns()
func (_Bindings *BindingsTransactorSession) ExecuteMsg(sourceChainId uint16, sourceNttManagerAddress [32]byte, message TransceiverStructsNttManagerMessage) (*types.Transaction, error) {
	return _Bindings.Contract.ExecuteMsg(&_Bindings.TransactOpts, sourceChainId, sourceNttManagerAddress, message)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_Bindings *BindingsTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_Bindings *BindingsSession) Initialize() (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_Bindings *BindingsTransactorSession) Initialize() (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Bindings *BindingsTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Bindings *BindingsSession) Migrate() (*types.Transaction, error) {
	return _Bindings.Contract.Migrate(&_Bindings.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Bindings *BindingsTransactorSession) Migrate() (*types.Transaction, error) {
	return _Bindings.Contract.Migrate(&_Bindings.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bindings *BindingsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bindings *BindingsSession) Pause() (*types.Transaction, error) {
	return _Bindings.Contract.Pause(&_Bindings.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bindings *BindingsTransactorSession) Pause() (*types.Transaction, error) {
	return _Bindings.Contract.Pause(&_Bindings.TransactOpts)
}

// RemoveTransceiver is a paid mutator transaction binding the contract method 0x9f86029c.
//
// Solidity: function removeTransceiver(address transceiver) returns()
func (_Bindings *BindingsTransactor) RemoveTransceiver(opts *bind.TransactOpts, transceiver common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "removeTransceiver", transceiver)
}

// RemoveTransceiver is a paid mutator transaction binding the contract method 0x9f86029c.
//
// Solidity: function removeTransceiver(address transceiver) returns()
func (_Bindings *BindingsSession) RemoveTransceiver(transceiver common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RemoveTransceiver(&_Bindings.TransactOpts, transceiver)
}

// RemoveTransceiver is a paid mutator transaction binding the contract method 0x9f86029c.
//
// Solidity: function removeTransceiver(address transceiver) returns()
func (_Bindings *BindingsTransactorSession) RemoveTransceiver(transceiver common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RemoveTransceiver(&_Bindings.TransactOpts, transceiver)
}

// SendMTokenIndex is a paid mutator transaction binding the contract method 0xf6f61b3a.
//
// Solidity: function sendMTokenIndex(uint16 destinationChainId_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Bindings *BindingsTransactor) SendMTokenIndex(opts *bind.TransactOpts, destinationChainId_ uint16, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "sendMTokenIndex", destinationChainId_, refundAddress_)
}

// SendMTokenIndex is a paid mutator transaction binding the contract method 0xf6f61b3a.
//
// Solidity: function sendMTokenIndex(uint16 destinationChainId_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Bindings *BindingsSession) SendMTokenIndex(destinationChainId_ uint16, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.SendMTokenIndex(&_Bindings.TransactOpts, destinationChainId_, refundAddress_)
}

// SendMTokenIndex is a paid mutator transaction binding the contract method 0xf6f61b3a.
//
// Solidity: function sendMTokenIndex(uint16 destinationChainId_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Bindings *BindingsTransactorSession) SendMTokenIndex(destinationChainId_ uint16, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.SendMTokenIndex(&_Bindings.TransactOpts, destinationChainId_, refundAddress_)
}

// SendRegistrarKey is a paid mutator transaction binding the contract method 0x1f33899e.
//
// Solidity: function sendRegistrarKey(uint16 destinationChainId_, bytes32 key_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Bindings *BindingsTransactor) SendRegistrarKey(opts *bind.TransactOpts, destinationChainId_ uint16, key_ [32]byte, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "sendRegistrarKey", destinationChainId_, key_, refundAddress_)
}

// SendRegistrarKey is a paid mutator transaction binding the contract method 0x1f33899e.
//
// Solidity: function sendRegistrarKey(uint16 destinationChainId_, bytes32 key_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Bindings *BindingsSession) SendRegistrarKey(destinationChainId_ uint16, key_ [32]byte, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.SendRegistrarKey(&_Bindings.TransactOpts, destinationChainId_, key_, refundAddress_)
}

// SendRegistrarKey is a paid mutator transaction binding the contract method 0x1f33899e.
//
// Solidity: function sendRegistrarKey(uint16 destinationChainId_, bytes32 key_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Bindings *BindingsTransactorSession) SendRegistrarKey(destinationChainId_ uint16, key_ [32]byte, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.SendRegistrarKey(&_Bindings.TransactOpts, destinationChainId_, key_, refundAddress_)
}

// SendRegistrarListStatus is a paid mutator transaction binding the contract method 0x76deb363.
//
// Solidity: function sendRegistrarListStatus(uint16 destinationChainId_, bytes32 listName_, address account_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Bindings *BindingsTransactor) SendRegistrarListStatus(opts *bind.TransactOpts, destinationChainId_ uint16, listName_ [32]byte, account_ common.Address, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "sendRegistrarListStatus", destinationChainId_, listName_, account_, refundAddress_)
}

// SendRegistrarListStatus is a paid mutator transaction binding the contract method 0x76deb363.
//
// Solidity: function sendRegistrarListStatus(uint16 destinationChainId_, bytes32 listName_, address account_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Bindings *BindingsSession) SendRegistrarListStatus(destinationChainId_ uint16, listName_ [32]byte, account_ common.Address, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.SendRegistrarListStatus(&_Bindings.TransactOpts, destinationChainId_, listName_, account_, refundAddress_)
}

// SendRegistrarListStatus is a paid mutator transaction binding the contract method 0x76deb363.
//
// Solidity: function sendRegistrarListStatus(uint16 destinationChainId_, bytes32 listName_, address account_, bytes32 refundAddress_) payable returns(bytes32 messageId_)
func (_Bindings *BindingsTransactorSession) SendRegistrarListStatus(destinationChainId_ uint16, listName_ [32]byte, account_ common.Address, refundAddress_ [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.SendRegistrarListStatus(&_Bindings.TransactOpts, destinationChainId_, listName_, account_, refundAddress_)
}

// SetInboundLimit is a paid mutator transaction binding the contract method 0x186ce612.
//
// Solidity: function setInboundLimit(uint256 limit, uint16 chainId_) returns()
func (_Bindings *BindingsTransactor) SetInboundLimit(opts *bind.TransactOpts, limit *big.Int, chainId_ uint16) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setInboundLimit", limit, chainId_)
}

// SetInboundLimit is a paid mutator transaction binding the contract method 0x186ce612.
//
// Solidity: function setInboundLimit(uint256 limit, uint16 chainId_) returns()
func (_Bindings *BindingsSession) SetInboundLimit(limit *big.Int, chainId_ uint16) (*types.Transaction, error) {
	return _Bindings.Contract.SetInboundLimit(&_Bindings.TransactOpts, limit, chainId_)
}

// SetInboundLimit is a paid mutator transaction binding the contract method 0x186ce612.
//
// Solidity: function setInboundLimit(uint256 limit, uint16 chainId_) returns()
func (_Bindings *BindingsTransactorSession) SetInboundLimit(limit *big.Int, chainId_ uint16) (*types.Transaction, error) {
	return _Bindings.Contract.SetInboundLimit(&_Bindings.TransactOpts, limit, chainId_)
}

// SetOutboundLimit is a paid mutator transaction binding the contract method 0x19017175.
//
// Solidity: function setOutboundLimit(uint256 limit) returns()
func (_Bindings *BindingsTransactor) SetOutboundLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setOutboundLimit", limit)
}

// SetOutboundLimit is a paid mutator transaction binding the contract method 0x19017175.
//
// Solidity: function setOutboundLimit(uint256 limit) returns()
func (_Bindings *BindingsSession) SetOutboundLimit(limit *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetOutboundLimit(&_Bindings.TransactOpts, limit)
}

// SetOutboundLimit is a paid mutator transaction binding the contract method 0x19017175.
//
// Solidity: function setOutboundLimit(uint256 limit) returns()
func (_Bindings *BindingsTransactorSession) SetOutboundLimit(limit *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetOutboundLimit(&_Bindings.TransactOpts, limit)
}

// SetPeer is a paid mutator transaction binding the contract method 0x7c918634.
//
// Solidity: function setPeer(uint16 peerChainId, bytes32 peerContract, uint8 decimals, uint256 inboundLimit) returns()
func (_Bindings *BindingsTransactor) SetPeer(opts *bind.TransactOpts, peerChainId uint16, peerContract [32]byte, decimals uint8, inboundLimit *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setPeer", peerChainId, peerContract, decimals, inboundLimit)
}

// SetPeer is a paid mutator transaction binding the contract method 0x7c918634.
//
// Solidity: function setPeer(uint16 peerChainId, bytes32 peerContract, uint8 decimals, uint256 inboundLimit) returns()
func (_Bindings *BindingsSession) SetPeer(peerChainId uint16, peerContract [32]byte, decimals uint8, inboundLimit *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetPeer(&_Bindings.TransactOpts, peerChainId, peerContract, decimals, inboundLimit)
}

// SetPeer is a paid mutator transaction binding the contract method 0x7c918634.
//
// Solidity: function setPeer(uint16 peerChainId, bytes32 peerContract, uint8 decimals, uint256 inboundLimit) returns()
func (_Bindings *BindingsTransactorSession) SetPeer(peerChainId uint16, peerContract [32]byte, decimals uint8, inboundLimit *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetPeer(&_Bindings.TransactOpts, peerChainId, peerContract, decimals, inboundLimit)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xe5a98603.
//
// Solidity: function setThreshold(uint8 threshold) returns()
func (_Bindings *BindingsTransactor) SetThreshold(opts *bind.TransactOpts, threshold uint8) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setThreshold", threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xe5a98603.
//
// Solidity: function setThreshold(uint8 threshold) returns()
func (_Bindings *BindingsSession) SetThreshold(threshold uint8) (*types.Transaction, error) {
	return _Bindings.Contract.SetThreshold(&_Bindings.TransactOpts, threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xe5a98603.
//
// Solidity: function setThreshold(uint8 threshold) returns()
func (_Bindings *BindingsTransactorSession) SetThreshold(threshold uint8) (*types.Transaction, error) {
	return _Bindings.Contract.SetThreshold(&_Bindings.TransactOpts, threshold)
}

// SetTransceiver is a paid mutator transaction binding the contract method 0x203e4a9b.
//
// Solidity: function setTransceiver(address transceiver) returns()
func (_Bindings *BindingsTransactor) SetTransceiver(opts *bind.TransactOpts, transceiver common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setTransceiver", transceiver)
}

// SetTransceiver is a paid mutator transaction binding the contract method 0x203e4a9b.
//
// Solidity: function setTransceiver(address transceiver) returns()
func (_Bindings *BindingsSession) SetTransceiver(transceiver common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetTransceiver(&_Bindings.TransactOpts, transceiver)
}

// SetTransceiver is a paid mutator transaction binding the contract method 0x203e4a9b.
//
// Solidity: function setTransceiver(address transceiver) returns()
func (_Bindings *BindingsTransactorSession) SetTransceiver(transceiver common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetTransceiver(&_Bindings.TransactOpts, transceiver)
}

// Transfer is a paid mutator transaction binding the contract method 0x961b94d0.
//
// Solidity: function transfer(uint256 amount, uint16 recipientChain, bytes32 recipient) payable returns(uint64)
func (_Bindings *BindingsTransactor) Transfer(opts *bind.TransactOpts, amount *big.Int, recipientChain uint16, recipient [32]byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transfer", amount, recipientChain, recipient)
}

// Transfer is a paid mutator transaction binding the contract method 0x961b94d0.
//
// Solidity: function transfer(uint256 amount, uint16 recipientChain, bytes32 recipient) payable returns(uint64)
func (_Bindings *BindingsSession) Transfer(amount *big.Int, recipientChain uint16, recipient [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.Transfer(&_Bindings.TransactOpts, amount, recipientChain, recipient)
}

// Transfer is a paid mutator transaction binding the contract method 0x961b94d0.
//
// Solidity: function transfer(uint256 amount, uint16 recipientChain, bytes32 recipient) payable returns(uint64)
func (_Bindings *BindingsTransactorSession) Transfer(amount *big.Int, recipientChain uint16, recipient [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.Transfer(&_Bindings.TransactOpts, amount, recipientChain, recipient)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xb293f97f.
//
// Solidity: function transfer(uint256 amount, uint16 recipientChain, bytes32 recipient, bytes32 refundAddress, bool shouldQueue, bytes transceiverInstructions) payable returns(uint64)
func (_Bindings *BindingsTransactor) Transfer0(opts *bind.TransactOpts, amount *big.Int, recipientChain uint16, recipient [32]byte, refundAddress [32]byte, shouldQueue bool, transceiverInstructions []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transfer0", amount, recipientChain, recipient, refundAddress, shouldQueue, transceiverInstructions)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xb293f97f.
//
// Solidity: function transfer(uint256 amount, uint16 recipientChain, bytes32 recipient, bytes32 refundAddress, bool shouldQueue, bytes transceiverInstructions) payable returns(uint64)
func (_Bindings *BindingsSession) Transfer0(amount *big.Int, recipientChain uint16, recipient [32]byte, refundAddress [32]byte, shouldQueue bool, transceiverInstructions []byte) (*types.Transaction, error) {
	return _Bindings.Contract.Transfer0(&_Bindings.TransactOpts, amount, recipientChain, recipient, refundAddress, shouldQueue, transceiverInstructions)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xb293f97f.
//
// Solidity: function transfer(uint256 amount, uint16 recipientChain, bytes32 recipient, bytes32 refundAddress, bool shouldQueue, bytes transceiverInstructions) payable returns(uint64)
func (_Bindings *BindingsTransactorSession) Transfer0(amount *big.Int, recipientChain uint16, recipient [32]byte, refundAddress [32]byte, shouldQueue bool, transceiverInstructions []byte) (*types.Transaction, error) {
	return _Bindings.Contract.Transfer0(&_Bindings.TransactOpts, amount, recipientChain, recipient, refundAddress, shouldQueue, transceiverInstructions)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// TransferPauserCapability is a paid mutator transaction binding the contract method 0x036de8af.
//
// Solidity: function transferPauserCapability(address newPauser) returns()
func (_Bindings *BindingsTransactor) TransferPauserCapability(opts *bind.TransactOpts, newPauser common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferPauserCapability", newPauser)
}

// TransferPauserCapability is a paid mutator transaction binding the contract method 0x036de8af.
//
// Solidity: function transferPauserCapability(address newPauser) returns()
func (_Bindings *BindingsSession) TransferPauserCapability(newPauser common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferPauserCapability(&_Bindings.TransactOpts, newPauser)
}

// TransferPauserCapability is a paid mutator transaction binding the contract method 0x036de8af.
//
// Solidity: function transferPauserCapability(address newPauser) returns()
func (_Bindings *BindingsTransactorSession) TransferPauserCapability(newPauser common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferPauserCapability(&_Bindings.TransactOpts, newPauser)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bindings *BindingsTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bindings *BindingsSession) Unpause() (*types.Transaction, error) {
	return _Bindings.Contract.Unpause(&_Bindings.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bindings *BindingsTransactorSession) Unpause() (*types.Transaction, error) {
	return _Bindings.Contract.Unpause(&_Bindings.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newImplementation) returns()
func (_Bindings *BindingsTransactor) Upgrade(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "upgrade", newImplementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newImplementation) returns()
func (_Bindings *BindingsSession) Upgrade(newImplementation common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Upgrade(&_Bindings.TransactOpts, newImplementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newImplementation) returns()
func (_Bindings *BindingsTransactorSession) Upgrade(newImplementation common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Upgrade(&_Bindings.TransactOpts, newImplementation)
}

// BindingsAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Bindings contract.
type BindingsAdminChangedIterator struct {
	Event *BindingsAdminChanged // Event containing the contract specifics and raw log

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
func (it *BindingsAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsAdminChanged)
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
		it.Event = new(BindingsAdminChanged)
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
func (it *BindingsAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsAdminChanged represents a AdminChanged event raised by the Bindings contract.
type BindingsAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Bindings *BindingsFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BindingsAdminChangedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BindingsAdminChangedIterator{contract: _Bindings.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Bindings *BindingsFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BindingsAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsAdminChanged)
				if err := _Bindings.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseAdminChanged(log types.Log) (*BindingsAdminChanged, error) {
	event := new(BindingsAdminChanged)
	if err := _Bindings.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Bindings contract.
type BindingsBeaconUpgradedIterator struct {
	Event *BindingsBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BindingsBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBeaconUpgraded)
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
		it.Event = new(BindingsBeaconUpgraded)
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
func (it *BindingsBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBeaconUpgraded represents a BeaconUpgraded event raised by the Bindings contract.
type BindingsBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Bindings *BindingsFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BindingsBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BindingsBeaconUpgradedIterator{contract: _Bindings.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Bindings *BindingsFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BindingsBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBeaconUpgraded)
				if err := _Bindings.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseBeaconUpgraded(log types.Log) (*BindingsBeaconUpgraded, error) {
	event := new(BindingsBeaconUpgraded)
	if err := _Bindings.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsEarningDisabledIterator is returned from FilterEarningDisabled and is used to iterate over the raw logs and unpacked data for EarningDisabled events raised by the Bindings contract.
type BindingsEarningDisabledIterator struct {
	Event *BindingsEarningDisabled // Event containing the contract specifics and raw log

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
func (it *BindingsEarningDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsEarningDisabled)
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
		it.Event = new(BindingsEarningDisabled)
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
func (it *BindingsEarningDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsEarningDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsEarningDisabled represents a EarningDisabled event raised by the Bindings contract.
type BindingsEarningDisabled struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEarningDisabled is a free log retrieval operation binding the contract event 0xee580fdb4da10ea17aa673e6f5c8c2370b4166d6a94bc88900e5a96d0589e3ce.
//
// Solidity: event EarningDisabled(uint128 index)
func (_Bindings *BindingsFilterer) FilterEarningDisabled(opts *bind.FilterOpts) (*BindingsEarningDisabledIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "EarningDisabled")
	if err != nil {
		return nil, err
	}
	return &BindingsEarningDisabledIterator{contract: _Bindings.contract, event: "EarningDisabled", logs: logs, sub: sub}, nil
}

// WatchEarningDisabled is a free log subscription operation binding the contract event 0xee580fdb4da10ea17aa673e6f5c8c2370b4166d6a94bc88900e5a96d0589e3ce.
//
// Solidity: event EarningDisabled(uint128 index)
func (_Bindings *BindingsFilterer) WatchEarningDisabled(opts *bind.WatchOpts, sink chan<- *BindingsEarningDisabled) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "EarningDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsEarningDisabled)
				if err := _Bindings.contract.UnpackLog(event, "EarningDisabled", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseEarningDisabled(log types.Log) (*BindingsEarningDisabled, error) {
	event := new(BindingsEarningDisabled)
	if err := _Bindings.contract.UnpackLog(event, "EarningDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsEarningEnabledIterator is returned from FilterEarningEnabled and is used to iterate over the raw logs and unpacked data for EarningEnabled events raised by the Bindings contract.
type BindingsEarningEnabledIterator struct {
	Event *BindingsEarningEnabled // Event containing the contract specifics and raw log

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
func (it *BindingsEarningEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsEarningEnabled)
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
		it.Event = new(BindingsEarningEnabled)
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
func (it *BindingsEarningEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsEarningEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsEarningEnabled represents a EarningEnabled event raised by the Bindings contract.
type BindingsEarningEnabled struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEarningEnabled is a free log retrieval operation binding the contract event 0x5098de6eb11dbd1127cf4dcd5e960e3944d48a7570b9b1939cff715cb35c5a18.
//
// Solidity: event EarningEnabled(uint128 index)
func (_Bindings *BindingsFilterer) FilterEarningEnabled(opts *bind.FilterOpts) (*BindingsEarningEnabledIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "EarningEnabled")
	if err != nil {
		return nil, err
	}
	return &BindingsEarningEnabledIterator{contract: _Bindings.contract, event: "EarningEnabled", logs: logs, sub: sub}, nil
}

// WatchEarningEnabled is a free log subscription operation binding the contract event 0x5098de6eb11dbd1127cf4dcd5e960e3944d48a7570b9b1939cff715cb35c5a18.
//
// Solidity: event EarningEnabled(uint128 index)
func (_Bindings *BindingsFilterer) WatchEarningEnabled(opts *bind.WatchOpts, sink chan<- *BindingsEarningEnabled) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "EarningEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsEarningEnabled)
				if err := _Bindings.contract.UnpackLog(event, "EarningEnabled", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseEarningEnabled(log types.Log) (*BindingsEarningEnabled, error) {
	event := new(BindingsEarningEnabled)
	if err := _Bindings.contract.UnpackLog(event, "EarningEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsInboundTransferQueuedIterator is returned from FilterInboundTransferQueued and is used to iterate over the raw logs and unpacked data for InboundTransferQueued events raised by the Bindings contract.
type BindingsInboundTransferQueuedIterator struct {
	Event *BindingsInboundTransferQueued // Event containing the contract specifics and raw log

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
func (it *BindingsInboundTransferQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsInboundTransferQueued)
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
		it.Event = new(BindingsInboundTransferQueued)
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
func (it *BindingsInboundTransferQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsInboundTransferQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsInboundTransferQueued represents a InboundTransferQueued event raised by the Bindings contract.
type BindingsInboundTransferQueued struct {
	Digest [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInboundTransferQueued is a free log retrieval operation binding the contract event 0x7f63c9251d82a933210c2b6d0b0f116252c3c116788120e64e8e8215df6f3162.
//
// Solidity: event InboundTransferQueued(bytes32 digest)
func (_Bindings *BindingsFilterer) FilterInboundTransferQueued(opts *bind.FilterOpts) (*BindingsInboundTransferQueuedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "InboundTransferQueued")
	if err != nil {
		return nil, err
	}
	return &BindingsInboundTransferQueuedIterator{contract: _Bindings.contract, event: "InboundTransferQueued", logs: logs, sub: sub}, nil
}

// WatchInboundTransferQueued is a free log subscription operation binding the contract event 0x7f63c9251d82a933210c2b6d0b0f116252c3c116788120e64e8e8215df6f3162.
//
// Solidity: event InboundTransferQueued(bytes32 digest)
func (_Bindings *BindingsFilterer) WatchInboundTransferQueued(opts *bind.WatchOpts, sink chan<- *BindingsInboundTransferQueued) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "InboundTransferQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsInboundTransferQueued)
				if err := _Bindings.contract.UnpackLog(event, "InboundTransferQueued", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseInboundTransferQueued(log types.Log) (*BindingsInboundTransferQueued, error) {
	event := new(BindingsInboundTransferQueued)
	if err := _Bindings.contract.UnpackLog(event, "InboundTransferQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bindings contract.
type BindingsInitializedIterator struct {
	Event *BindingsInitialized // Event containing the contract specifics and raw log

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
func (it *BindingsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsInitialized)
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
		it.Event = new(BindingsInitialized)
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
func (it *BindingsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsInitialized represents a Initialized event raised by the Bindings contract.
type BindingsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) FilterInitialized(opts *bind.FilterOpts) (*BindingsInitializedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BindingsInitializedIterator{contract: _Bindings.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BindingsInitialized) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsInitialized)
				if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseInitialized(log types.Log) (*BindingsInitialized, error) {
	event := new(BindingsInitialized)
	if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMTokenIndexSentIterator is returned from FilterMTokenIndexSent and is used to iterate over the raw logs and unpacked data for MTokenIndexSent events raised by the Bindings contract.
type BindingsMTokenIndexSentIterator struct {
	Event *BindingsMTokenIndexSent // Event containing the contract specifics and raw log

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
func (it *BindingsMTokenIndexSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMTokenIndexSent)
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
		it.Event = new(BindingsMTokenIndexSent)
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
func (it *BindingsMTokenIndexSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMTokenIndexSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMTokenIndexSent represents a MTokenIndexSent event raised by the Bindings contract.
type BindingsMTokenIndexSent struct {
	DestinationChainId uint16
	MessageId          [32]byte
	Index              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMTokenIndexSent is a free log retrieval operation binding the contract event 0x3e1e343484a686e91a59f00b76fa98aaf17cb636b12f75a6654f432f45e64c14.
//
// Solidity: event MTokenIndexSent(uint16 indexed destinationChainId, bytes32 messageId, uint128 index)
func (_Bindings *BindingsFilterer) FilterMTokenIndexSent(opts *bind.FilterOpts, destinationChainId []uint16) (*BindingsMTokenIndexSentIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MTokenIndexSent", destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMTokenIndexSentIterator{contract: _Bindings.contract, event: "MTokenIndexSent", logs: logs, sub: sub}, nil
}

// WatchMTokenIndexSent is a free log subscription operation binding the contract event 0x3e1e343484a686e91a59f00b76fa98aaf17cb636b12f75a6654f432f45e64c14.
//
// Solidity: event MTokenIndexSent(uint16 indexed destinationChainId, bytes32 messageId, uint128 index)
func (_Bindings *BindingsFilterer) WatchMTokenIndexSent(opts *bind.WatchOpts, sink chan<- *BindingsMTokenIndexSent, destinationChainId []uint16) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MTokenIndexSent", destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMTokenIndexSent)
				if err := _Bindings.contract.UnpackLog(event, "MTokenIndexSent", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseMTokenIndexSent(log types.Log) (*BindingsMTokenIndexSent, error) {
	event := new(BindingsMTokenIndexSent)
	if err := _Bindings.contract.UnpackLog(event, "MTokenIndexSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMTokenReceivedIterator is returned from FilterMTokenReceived and is used to iterate over the raw logs and unpacked data for MTokenReceived events raised by the Bindings contract.
type BindingsMTokenReceivedIterator struct {
	Event *BindingsMTokenReceived // Event containing the contract specifics and raw log

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
func (it *BindingsMTokenReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMTokenReceived)
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
		it.Event = new(BindingsMTokenReceived)
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
func (it *BindingsMTokenReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMTokenReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMTokenReceived represents a MTokenReceived event raised by the Bindings contract.
type BindingsMTokenReceived struct {
	SourceChainId uint16
	MessageId     [32]byte
	Sender        [32]byte
	Recipient     common.Address
	Amount        *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMTokenReceived is a free log retrieval operation binding the contract event 0xbd6146ab70befc9dc9a724a2fa48ae30b5a4dbe253e9e80e9337b659f3c1eb9a.
//
// Solidity: event MTokenReceived(uint16 indexed sourceChainId, bytes32 messageId, bytes32 indexed sender, address indexed recipient, uint256 amount, uint128 index)
func (_Bindings *BindingsFilterer) FilterMTokenReceived(opts *bind.FilterOpts, sourceChainId []uint16, sender [][32]byte, recipient []common.Address) (*BindingsMTokenReceivedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MTokenReceived", sourceChainIdRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMTokenReceivedIterator{contract: _Bindings.contract, event: "MTokenReceived", logs: logs, sub: sub}, nil
}

// WatchMTokenReceived is a free log subscription operation binding the contract event 0xbd6146ab70befc9dc9a724a2fa48ae30b5a4dbe253e9e80e9337b659f3c1eb9a.
//
// Solidity: event MTokenReceived(uint16 indexed sourceChainId, bytes32 messageId, bytes32 indexed sender, address indexed recipient, uint256 amount, uint128 index)
func (_Bindings *BindingsFilterer) WatchMTokenReceived(opts *bind.WatchOpts, sink chan<- *BindingsMTokenReceived, sourceChainId []uint16, sender [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MTokenReceived", sourceChainIdRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMTokenReceived)
				if err := _Bindings.contract.UnpackLog(event, "MTokenReceived", log); err != nil {
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

// ParseMTokenReceived is a log parse operation binding the contract event 0xbd6146ab70befc9dc9a724a2fa48ae30b5a4dbe253e9e80e9337b659f3c1eb9a.
//
// Solidity: event MTokenReceived(uint16 indexed sourceChainId, bytes32 messageId, bytes32 indexed sender, address indexed recipient, uint256 amount, uint128 index)
func (_Bindings *BindingsFilterer) ParseMTokenReceived(log types.Log) (*BindingsMTokenReceived, error) {
	event := new(BindingsMTokenReceived)
	if err := _Bindings.contract.UnpackLog(event, "MTokenReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMTokenSentIterator is returned from FilterMTokenSent and is used to iterate over the raw logs and unpacked data for MTokenSent events raised by the Bindings contract.
type BindingsMTokenSentIterator struct {
	Event *BindingsMTokenSent // Event containing the contract specifics and raw log

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
func (it *BindingsMTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMTokenSent)
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
		it.Event = new(BindingsMTokenSent)
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
func (it *BindingsMTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMTokenSent represents a MTokenSent event raised by the Bindings contract.
type BindingsMTokenSent struct {
	DestinationChainId uint16
	MessageId          [32]byte
	Sender             common.Address
	Recipient          [32]byte
	Amount             *big.Int
	Index              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMTokenSent is a free log retrieval operation binding the contract event 0x5b89200a06a504bee6b4e357fe58bee68910ddfd4bdf6bd54c233df75a380d30.
//
// Solidity: event MTokenSent(uint16 indexed destinationChainId, bytes32 messageId, address indexed sender, bytes32 indexed recipient, uint256 amount, uint128 index)
func (_Bindings *BindingsFilterer) FilterMTokenSent(opts *bind.FilterOpts, destinationChainId []uint16, sender []common.Address, recipient [][32]byte) (*BindingsMTokenSentIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MTokenSent", destinationChainIdRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMTokenSentIterator{contract: _Bindings.contract, event: "MTokenSent", logs: logs, sub: sub}, nil
}

// WatchMTokenSent is a free log subscription operation binding the contract event 0x5b89200a06a504bee6b4e357fe58bee68910ddfd4bdf6bd54c233df75a380d30.
//
// Solidity: event MTokenSent(uint16 indexed destinationChainId, bytes32 messageId, address indexed sender, bytes32 indexed recipient, uint256 amount, uint128 index)
func (_Bindings *BindingsFilterer) WatchMTokenSent(opts *bind.WatchOpts, sink chan<- *BindingsMTokenSent, destinationChainId []uint16, sender []common.Address, recipient [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MTokenSent", destinationChainIdRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMTokenSent)
				if err := _Bindings.contract.UnpackLog(event, "MTokenSent", log); err != nil {
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

// ParseMTokenSent is a log parse operation binding the contract event 0x5b89200a06a504bee6b4e357fe58bee68910ddfd4bdf6bd54c233df75a380d30.
//
// Solidity: event MTokenSent(uint16 indexed destinationChainId, bytes32 messageId, address indexed sender, bytes32 indexed recipient, uint256 amount, uint128 index)
func (_Bindings *BindingsFilterer) ParseMTokenSent(log types.Log) (*BindingsMTokenSent, error) {
	event := new(BindingsMTokenSent)
	if err := _Bindings.contract.UnpackLog(event, "MTokenSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMessageAlreadyExecutedIterator is returned from FilterMessageAlreadyExecuted and is used to iterate over the raw logs and unpacked data for MessageAlreadyExecuted events raised by the Bindings contract.
type BindingsMessageAlreadyExecutedIterator struct {
	Event *BindingsMessageAlreadyExecuted // Event containing the contract specifics and raw log

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
func (it *BindingsMessageAlreadyExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMessageAlreadyExecuted)
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
		it.Event = new(BindingsMessageAlreadyExecuted)
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
func (it *BindingsMessageAlreadyExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMessageAlreadyExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMessageAlreadyExecuted represents a MessageAlreadyExecuted event raised by the Bindings contract.
type BindingsMessageAlreadyExecuted struct {
	SourceNttManager [32]byte
	MsgHash          [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMessageAlreadyExecuted is a free log retrieval operation binding the contract event 0x4069dff8c9df7e38d2867c0910bd96fd61787695e5380281148c04932d02bef2.
//
// Solidity: event MessageAlreadyExecuted(bytes32 indexed sourceNttManager, bytes32 indexed msgHash)
func (_Bindings *BindingsFilterer) FilterMessageAlreadyExecuted(opts *bind.FilterOpts, sourceNttManager [][32]byte, msgHash [][32]byte) (*BindingsMessageAlreadyExecutedIterator, error) {

	var sourceNttManagerRule []interface{}
	for _, sourceNttManagerItem := range sourceNttManager {
		sourceNttManagerRule = append(sourceNttManagerRule, sourceNttManagerItem)
	}
	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MessageAlreadyExecuted", sourceNttManagerRule, msgHashRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMessageAlreadyExecutedIterator{contract: _Bindings.contract, event: "MessageAlreadyExecuted", logs: logs, sub: sub}, nil
}

// WatchMessageAlreadyExecuted is a free log subscription operation binding the contract event 0x4069dff8c9df7e38d2867c0910bd96fd61787695e5380281148c04932d02bef2.
//
// Solidity: event MessageAlreadyExecuted(bytes32 indexed sourceNttManager, bytes32 indexed msgHash)
func (_Bindings *BindingsFilterer) WatchMessageAlreadyExecuted(opts *bind.WatchOpts, sink chan<- *BindingsMessageAlreadyExecuted, sourceNttManager [][32]byte, msgHash [][32]byte) (event.Subscription, error) {

	var sourceNttManagerRule []interface{}
	for _, sourceNttManagerItem := range sourceNttManager {
		sourceNttManagerRule = append(sourceNttManagerRule, sourceNttManagerItem)
	}
	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MessageAlreadyExecuted", sourceNttManagerRule, msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMessageAlreadyExecuted)
				if err := _Bindings.contract.UnpackLog(event, "MessageAlreadyExecuted", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseMessageAlreadyExecuted(log types.Log) (*BindingsMessageAlreadyExecuted, error) {
	event := new(BindingsMessageAlreadyExecuted)
	if err := _Bindings.contract.UnpackLog(event, "MessageAlreadyExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMessageAttestedToIterator is returned from FilterMessageAttestedTo and is used to iterate over the raw logs and unpacked data for MessageAttestedTo events raised by the Bindings contract.
type BindingsMessageAttestedToIterator struct {
	Event *BindingsMessageAttestedTo // Event containing the contract specifics and raw log

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
func (it *BindingsMessageAttestedToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMessageAttestedTo)
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
		it.Event = new(BindingsMessageAttestedTo)
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
func (it *BindingsMessageAttestedToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMessageAttestedToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMessageAttestedTo represents a MessageAttestedTo event raised by the Bindings contract.
type BindingsMessageAttestedTo struct {
	Digest      [32]byte
	Transceiver common.Address
	Index       uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageAttestedTo is a free log retrieval operation binding the contract event 0x35a2101eaac94b493e0dfca061f9a7f087913fde8678e7cde0aca9897edba0e5.
//
// Solidity: event MessageAttestedTo(bytes32 digest, address transceiver, uint8 index)
func (_Bindings *BindingsFilterer) FilterMessageAttestedTo(opts *bind.FilterOpts) (*BindingsMessageAttestedToIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MessageAttestedTo")
	if err != nil {
		return nil, err
	}
	return &BindingsMessageAttestedToIterator{contract: _Bindings.contract, event: "MessageAttestedTo", logs: logs, sub: sub}, nil
}

// WatchMessageAttestedTo is a free log subscription operation binding the contract event 0x35a2101eaac94b493e0dfca061f9a7f087913fde8678e7cde0aca9897edba0e5.
//
// Solidity: event MessageAttestedTo(bytes32 digest, address transceiver, uint8 index)
func (_Bindings *BindingsFilterer) WatchMessageAttestedTo(opts *bind.WatchOpts, sink chan<- *BindingsMessageAttestedTo) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MessageAttestedTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMessageAttestedTo)
				if err := _Bindings.contract.UnpackLog(event, "MessageAttestedTo", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseMessageAttestedTo(log types.Log) (*BindingsMessageAttestedTo, error) {
	event := new(BindingsMessageAttestedTo)
	if err := _Bindings.contract.UnpackLog(event, "MessageAttestedTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsNotPausedIterator is returned from FilterNotPaused and is used to iterate over the raw logs and unpacked data for NotPaused events raised by the Bindings contract.
type BindingsNotPausedIterator struct {
	Event *BindingsNotPaused // Event containing the contract specifics and raw log

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
func (it *BindingsNotPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsNotPaused)
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
		it.Event = new(BindingsNotPaused)
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
func (it *BindingsNotPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsNotPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsNotPaused represents a NotPaused event raised by the Bindings contract.
type BindingsNotPaused struct {
	NotPaused bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNotPaused is a free log retrieval operation binding the contract event 0xe11c2112add17fb763d3bd59f63b10429c3e11373da4fb8ef6725107a2fdc4b0.
//
// Solidity: event NotPaused(bool notPaused)
func (_Bindings *BindingsFilterer) FilterNotPaused(opts *bind.FilterOpts) (*BindingsNotPausedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "NotPaused")
	if err != nil {
		return nil, err
	}
	return &BindingsNotPausedIterator{contract: _Bindings.contract, event: "NotPaused", logs: logs, sub: sub}, nil
}

// WatchNotPaused is a free log subscription operation binding the contract event 0xe11c2112add17fb763d3bd59f63b10429c3e11373da4fb8ef6725107a2fdc4b0.
//
// Solidity: event NotPaused(bool notPaused)
func (_Bindings *BindingsFilterer) WatchNotPaused(opts *bind.WatchOpts, sink chan<- *BindingsNotPaused) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "NotPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsNotPaused)
				if err := _Bindings.contract.UnpackLog(event, "NotPaused", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseNotPaused(log types.Log) (*BindingsNotPaused, error) {
	event := new(BindingsNotPaused)
	if err := _Bindings.contract.UnpackLog(event, "NotPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOutboundTransferCancelledIterator is returned from FilterOutboundTransferCancelled and is used to iterate over the raw logs and unpacked data for OutboundTransferCancelled events raised by the Bindings contract.
type BindingsOutboundTransferCancelledIterator struct {
	Event *BindingsOutboundTransferCancelled // Event containing the contract specifics and raw log

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
func (it *BindingsOutboundTransferCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOutboundTransferCancelled)
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
		it.Event = new(BindingsOutboundTransferCancelled)
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
func (it *BindingsOutboundTransferCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOutboundTransferCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOutboundTransferCancelled represents a OutboundTransferCancelled event raised by the Bindings contract.
type BindingsOutboundTransferCancelled struct {
	Sequence  *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOutboundTransferCancelled is a free log retrieval operation binding the contract event 0xf80e572ae1b63e2449629b6c7d783add85c36473926f216077f17ee002bcfd07.
//
// Solidity: event OutboundTransferCancelled(uint256 sequence, address recipient, uint256 amount)
func (_Bindings *BindingsFilterer) FilterOutboundTransferCancelled(opts *bind.FilterOpts) (*BindingsOutboundTransferCancelledIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OutboundTransferCancelled")
	if err != nil {
		return nil, err
	}
	return &BindingsOutboundTransferCancelledIterator{contract: _Bindings.contract, event: "OutboundTransferCancelled", logs: logs, sub: sub}, nil
}

// WatchOutboundTransferCancelled is a free log subscription operation binding the contract event 0xf80e572ae1b63e2449629b6c7d783add85c36473926f216077f17ee002bcfd07.
//
// Solidity: event OutboundTransferCancelled(uint256 sequence, address recipient, uint256 amount)
func (_Bindings *BindingsFilterer) WatchOutboundTransferCancelled(opts *bind.WatchOpts, sink chan<- *BindingsOutboundTransferCancelled) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OutboundTransferCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOutboundTransferCancelled)
				if err := _Bindings.contract.UnpackLog(event, "OutboundTransferCancelled", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseOutboundTransferCancelled(log types.Log) (*BindingsOutboundTransferCancelled, error) {
	event := new(BindingsOutboundTransferCancelled)
	if err := _Bindings.contract.UnpackLog(event, "OutboundTransferCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOutboundTransferQueuedIterator is returned from FilterOutboundTransferQueued and is used to iterate over the raw logs and unpacked data for OutboundTransferQueued events raised by the Bindings contract.
type BindingsOutboundTransferQueuedIterator struct {
	Event *BindingsOutboundTransferQueued // Event containing the contract specifics and raw log

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
func (it *BindingsOutboundTransferQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOutboundTransferQueued)
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
		it.Event = new(BindingsOutboundTransferQueued)
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
func (it *BindingsOutboundTransferQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOutboundTransferQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOutboundTransferQueued represents a OutboundTransferQueued event raised by the Bindings contract.
type BindingsOutboundTransferQueued struct {
	QueueSequence uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutboundTransferQueued is a free log retrieval operation binding the contract event 0x69add1952a6a6b9cb86f04d05f0cb605cbb469a50ae916139d34495a9991481f.
//
// Solidity: event OutboundTransferQueued(uint64 queueSequence)
func (_Bindings *BindingsFilterer) FilterOutboundTransferQueued(opts *bind.FilterOpts) (*BindingsOutboundTransferQueuedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OutboundTransferQueued")
	if err != nil {
		return nil, err
	}
	return &BindingsOutboundTransferQueuedIterator{contract: _Bindings.contract, event: "OutboundTransferQueued", logs: logs, sub: sub}, nil
}

// WatchOutboundTransferQueued is a free log subscription operation binding the contract event 0x69add1952a6a6b9cb86f04d05f0cb605cbb469a50ae916139d34495a9991481f.
//
// Solidity: event OutboundTransferQueued(uint64 queueSequence)
func (_Bindings *BindingsFilterer) WatchOutboundTransferQueued(opts *bind.WatchOpts, sink chan<- *BindingsOutboundTransferQueued) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OutboundTransferQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOutboundTransferQueued)
				if err := _Bindings.contract.UnpackLog(event, "OutboundTransferQueued", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseOutboundTransferQueued(log types.Log) (*BindingsOutboundTransferQueued, error) {
	event := new(BindingsOutboundTransferQueued)
	if err := _Bindings.contract.UnpackLog(event, "OutboundTransferQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOutboundTransferRateLimitedIterator is returned from FilterOutboundTransferRateLimited and is used to iterate over the raw logs and unpacked data for OutboundTransferRateLimited events raised by the Bindings contract.
type BindingsOutboundTransferRateLimitedIterator struct {
	Event *BindingsOutboundTransferRateLimited // Event containing the contract specifics and raw log

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
func (it *BindingsOutboundTransferRateLimitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOutboundTransferRateLimited)
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
		it.Event = new(BindingsOutboundTransferRateLimited)
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
func (it *BindingsOutboundTransferRateLimitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOutboundTransferRateLimitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOutboundTransferRateLimited represents a OutboundTransferRateLimited event raised by the Bindings contract.
type BindingsOutboundTransferRateLimited struct {
	Sender          common.Address
	Sequence        uint64
	Amount          *big.Int
	CurrentCapacity *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOutboundTransferRateLimited is a free log retrieval operation binding the contract event 0xf33512b84e24a49905c26c6991942fc5a9652411769fc1e448f967cdb049f08a.
//
// Solidity: event OutboundTransferRateLimited(address indexed sender, uint64 sequence, uint256 amount, uint256 currentCapacity)
func (_Bindings *BindingsFilterer) FilterOutboundTransferRateLimited(opts *bind.FilterOpts, sender []common.Address) (*BindingsOutboundTransferRateLimitedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OutboundTransferRateLimited", senderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOutboundTransferRateLimitedIterator{contract: _Bindings.contract, event: "OutboundTransferRateLimited", logs: logs, sub: sub}, nil
}

// WatchOutboundTransferRateLimited is a free log subscription operation binding the contract event 0xf33512b84e24a49905c26c6991942fc5a9652411769fc1e448f967cdb049f08a.
//
// Solidity: event OutboundTransferRateLimited(address indexed sender, uint64 sequence, uint256 amount, uint256 currentCapacity)
func (_Bindings *BindingsFilterer) WatchOutboundTransferRateLimited(opts *bind.WatchOpts, sink chan<- *BindingsOutboundTransferRateLimited, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OutboundTransferRateLimited", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOutboundTransferRateLimited)
				if err := _Bindings.contract.UnpackLog(event, "OutboundTransferRateLimited", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseOutboundTransferRateLimited(log types.Log) (*BindingsOutboundTransferRateLimited, error) {
	event := new(BindingsOutboundTransferRateLimited)
	if err := _Bindings.contract.UnpackLog(event, "OutboundTransferRateLimited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bindings contract.
type BindingsOwnershipTransferredIterator struct {
	Event *BindingsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BindingsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOwnershipTransferred)
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
		it.Event = new(BindingsOwnershipTransferred)
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
func (it *BindingsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOwnershipTransferred represents a OwnershipTransferred event raised by the Bindings contract.
type BindingsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOwnershipTransferredIterator{contract: _Bindings.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindingsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOwnershipTransferred)
				if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseOwnershipTransferred(log types.Log) (*BindingsOwnershipTransferred, error) {
	event := new(BindingsOwnershipTransferred)
	if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bindings contract.
type BindingsPausedIterator struct {
	Event *BindingsPaused // Event containing the contract specifics and raw log

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
func (it *BindingsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPaused)
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
		it.Event = new(BindingsPaused)
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
func (it *BindingsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPaused represents a Paused event raised by the Bindings contract.
type BindingsPaused struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool paused)
func (_Bindings *BindingsFilterer) FilterPaused(opts *bind.FilterOpts) (*BindingsPausedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BindingsPausedIterator{contract: _Bindings.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool paused)
func (_Bindings *BindingsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BindingsPaused) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPaused)
				if err := _Bindings.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParsePaused(log types.Log) (*BindingsPaused, error) {
	event := new(BindingsPaused)
	if err := _Bindings.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPauserTransferredIterator is returned from FilterPauserTransferred and is used to iterate over the raw logs and unpacked data for PauserTransferred events raised by the Bindings contract.
type BindingsPauserTransferredIterator struct {
	Event *BindingsPauserTransferred // Event containing the contract specifics and raw log

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
func (it *BindingsPauserTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPauserTransferred)
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
		it.Event = new(BindingsPauserTransferred)
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
func (it *BindingsPauserTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPauserTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPauserTransferred represents a PauserTransferred event raised by the Bindings contract.
type BindingsPauserTransferred struct {
	OldPauser common.Address
	NewPauser common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPauserTransferred is a free log retrieval operation binding the contract event 0x51c4874e0f23f262e04a38c51751336dde72126d67f53eb672aaff02996b3ef6.
//
// Solidity: event PauserTransferred(address indexed oldPauser, address indexed newPauser)
func (_Bindings *BindingsFilterer) FilterPauserTransferred(opts *bind.FilterOpts, oldPauser []common.Address, newPauser []common.Address) (*BindingsPauserTransferredIterator, error) {

	var oldPauserRule []interface{}
	for _, oldPauserItem := range oldPauser {
		oldPauserRule = append(oldPauserRule, oldPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "PauserTransferred", oldPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return &BindingsPauserTransferredIterator{contract: _Bindings.contract, event: "PauserTransferred", logs: logs, sub: sub}, nil
}

// WatchPauserTransferred is a free log subscription operation binding the contract event 0x51c4874e0f23f262e04a38c51751336dde72126d67f53eb672aaff02996b3ef6.
//
// Solidity: event PauserTransferred(address indexed oldPauser, address indexed newPauser)
func (_Bindings *BindingsFilterer) WatchPauserTransferred(opts *bind.WatchOpts, sink chan<- *BindingsPauserTransferred, oldPauser []common.Address, newPauser []common.Address) (event.Subscription, error) {

	var oldPauserRule []interface{}
	for _, oldPauserItem := range oldPauser {
		oldPauserRule = append(oldPauserRule, oldPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "PauserTransferred", oldPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPauserTransferred)
				if err := _Bindings.contract.UnpackLog(event, "PauserTransferred", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParsePauserTransferred(log types.Log) (*BindingsPauserTransferred, error) {
	event := new(BindingsPauserTransferred)
	if err := _Bindings.contract.UnpackLog(event, "PauserTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPeerUpdatedIterator is returned from FilterPeerUpdated and is used to iterate over the raw logs and unpacked data for PeerUpdated events raised by the Bindings contract.
type BindingsPeerUpdatedIterator struct {
	Event *BindingsPeerUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsPeerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPeerUpdated)
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
		it.Event = new(BindingsPeerUpdated)
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
func (it *BindingsPeerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPeerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPeerUpdated represents a PeerUpdated event raised by the Bindings contract.
type BindingsPeerUpdated struct {
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
func (_Bindings *BindingsFilterer) FilterPeerUpdated(opts *bind.FilterOpts, chainId_ []uint16) (*BindingsPeerUpdatedIterator, error) {

	var chainId_Rule []interface{}
	for _, chainId_Item := range chainId_ {
		chainId_Rule = append(chainId_Rule, chainId_Item)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "PeerUpdated", chainId_Rule)
	if err != nil {
		return nil, err
	}
	return &BindingsPeerUpdatedIterator{contract: _Bindings.contract, event: "PeerUpdated", logs: logs, sub: sub}, nil
}

// WatchPeerUpdated is a free log subscription operation binding the contract event 0x1456404e7f41f35c3daac941bb50bad417a66275c3040061b4287d787719599d.
//
// Solidity: event PeerUpdated(uint16 indexed chainId_, bytes32 oldPeerContract, uint8 oldPeerDecimals, bytes32 peerContract, uint8 peerDecimals)
func (_Bindings *BindingsFilterer) WatchPeerUpdated(opts *bind.WatchOpts, sink chan<- *BindingsPeerUpdated, chainId_ []uint16) (event.Subscription, error) {

	var chainId_Rule []interface{}
	for _, chainId_Item := range chainId_ {
		chainId_Rule = append(chainId_Rule, chainId_Item)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "PeerUpdated", chainId_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPeerUpdated)
				if err := _Bindings.contract.UnpackLog(event, "PeerUpdated", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParsePeerUpdated(log types.Log) (*BindingsPeerUpdated, error) {
	event := new(BindingsPeerUpdated)
	if err := _Bindings.contract.UnpackLog(event, "PeerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRegistrarKeySentIterator is returned from FilterRegistrarKeySent and is used to iterate over the raw logs and unpacked data for RegistrarKeySent events raised by the Bindings contract.
type BindingsRegistrarKeySentIterator struct {
	Event *BindingsRegistrarKeySent // Event containing the contract specifics and raw log

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
func (it *BindingsRegistrarKeySentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRegistrarKeySent)
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
		it.Event = new(BindingsRegistrarKeySent)
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
func (it *BindingsRegistrarKeySentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRegistrarKeySentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRegistrarKeySent represents a RegistrarKeySent event raised by the Bindings contract.
type BindingsRegistrarKeySent struct {
	DestinationChainId uint16
	MessageId          [32]byte
	Key                [32]byte
	Value              [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRegistrarKeySent is a free log retrieval operation binding the contract event 0x4e419c12ebf7388707567369b500f54671c9323b834077c5391f6054c1169544.
//
// Solidity: event RegistrarKeySent(uint16 indexed destinationChainId, bytes32 messageId, bytes32 indexed key, bytes32 value)
func (_Bindings *BindingsFilterer) FilterRegistrarKeySent(opts *bind.FilterOpts, destinationChainId []uint16, key [][32]byte) (*BindingsRegistrarKeySentIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RegistrarKeySent", destinationChainIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRegistrarKeySentIterator{contract: _Bindings.contract, event: "RegistrarKeySent", logs: logs, sub: sub}, nil
}

// WatchRegistrarKeySent is a free log subscription operation binding the contract event 0x4e419c12ebf7388707567369b500f54671c9323b834077c5391f6054c1169544.
//
// Solidity: event RegistrarKeySent(uint16 indexed destinationChainId, bytes32 messageId, bytes32 indexed key, bytes32 value)
func (_Bindings *BindingsFilterer) WatchRegistrarKeySent(opts *bind.WatchOpts, sink chan<- *BindingsRegistrarKeySent, destinationChainId []uint16, key [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RegistrarKeySent", destinationChainIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRegistrarKeySent)
				if err := _Bindings.contract.UnpackLog(event, "RegistrarKeySent", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseRegistrarKeySent(log types.Log) (*BindingsRegistrarKeySent, error) {
	event := new(BindingsRegistrarKeySent)
	if err := _Bindings.contract.UnpackLog(event, "RegistrarKeySent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRegistrarListStatusSentIterator is returned from FilterRegistrarListStatusSent and is used to iterate over the raw logs and unpacked data for RegistrarListStatusSent events raised by the Bindings contract.
type BindingsRegistrarListStatusSentIterator struct {
	Event *BindingsRegistrarListStatusSent // Event containing the contract specifics and raw log

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
func (it *BindingsRegistrarListStatusSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRegistrarListStatusSent)
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
		it.Event = new(BindingsRegistrarListStatusSent)
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
func (it *BindingsRegistrarListStatusSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRegistrarListStatusSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRegistrarListStatusSent represents a RegistrarListStatusSent event raised by the Bindings contract.
type BindingsRegistrarListStatusSent struct {
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
func (_Bindings *BindingsFilterer) FilterRegistrarListStatusSent(opts *bind.FilterOpts, destinationChainId []uint16, listName [][32]byte, account []common.Address) (*BindingsRegistrarListStatusSentIterator, error) {

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

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RegistrarListStatusSent", destinationChainIdRule, listNameRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRegistrarListStatusSentIterator{contract: _Bindings.contract, event: "RegistrarListStatusSent", logs: logs, sub: sub}, nil
}

// WatchRegistrarListStatusSent is a free log subscription operation binding the contract event 0xe4170b1dad8af8275fc8c0a7ec8ccd5f05e5139db178ce9a9d6566696e02b832.
//
// Solidity: event RegistrarListStatusSent(uint16 indexed destinationChainId, bytes32 messageId, bytes32 indexed listName, address indexed account, bool status)
func (_Bindings *BindingsFilterer) WatchRegistrarListStatusSent(opts *bind.WatchOpts, sink chan<- *BindingsRegistrarListStatusSent, destinationChainId []uint16, listName [][32]byte, account []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RegistrarListStatusSent", destinationChainIdRule, listNameRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRegistrarListStatusSent)
				if err := _Bindings.contract.UnpackLog(event, "RegistrarListStatusSent", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseRegistrarListStatusSent(log types.Log) (*BindingsRegistrarListStatusSent, error) {
	event := new(BindingsRegistrarListStatusSent)
	if err := _Bindings.contract.UnpackLog(event, "RegistrarListStatusSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the Bindings contract.
type BindingsThresholdChangedIterator struct {
	Event *BindingsThresholdChanged // Event containing the contract specifics and raw log

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
func (it *BindingsThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsThresholdChanged)
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
		it.Event = new(BindingsThresholdChanged)
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
func (it *BindingsThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsThresholdChanged represents a ThresholdChanged event raised by the Bindings contract.
type BindingsThresholdChanged struct {
	OldThreshold uint8
	Threshold    uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x2a855b929b9a53c6fb5b5ed248b27e502b709c088e036a5aa17620c8fc5085a9.
//
// Solidity: event ThresholdChanged(uint8 oldThreshold, uint8 threshold)
func (_Bindings *BindingsFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*BindingsThresholdChangedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &BindingsThresholdChangedIterator{contract: _Bindings.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x2a855b929b9a53c6fb5b5ed248b27e502b709c088e036a5aa17620c8fc5085a9.
//
// Solidity: event ThresholdChanged(uint8 oldThreshold, uint8 threshold)
func (_Bindings *BindingsFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *BindingsThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsThresholdChanged)
				if err := _Bindings.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseThresholdChanged(log types.Log) (*BindingsThresholdChanged, error) {
	event := new(BindingsThresholdChanged)
	if err := _Bindings.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsTransceiverAddedIterator is returned from FilterTransceiverAdded and is used to iterate over the raw logs and unpacked data for TransceiverAdded events raised by the Bindings contract.
type BindingsTransceiverAddedIterator struct {
	Event *BindingsTransceiverAdded // Event containing the contract specifics and raw log

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
func (it *BindingsTransceiverAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTransceiverAdded)
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
		it.Event = new(BindingsTransceiverAdded)
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
func (it *BindingsTransceiverAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTransceiverAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTransceiverAdded represents a TransceiverAdded event raised by the Bindings contract.
type BindingsTransceiverAdded struct {
	Transceiver     common.Address
	TransceiversNum *big.Int
	Threshold       uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransceiverAdded is a free log retrieval operation binding the contract event 0xf05962b5774c658e85ed80c91a75af9d66d2af2253dda480f90bce78aff5eda5.
//
// Solidity: event TransceiverAdded(address transceiver, uint256 transceiversNum, uint8 threshold)
func (_Bindings *BindingsFilterer) FilterTransceiverAdded(opts *bind.FilterOpts) (*BindingsTransceiverAddedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "TransceiverAdded")
	if err != nil {
		return nil, err
	}
	return &BindingsTransceiverAddedIterator{contract: _Bindings.contract, event: "TransceiverAdded", logs: logs, sub: sub}, nil
}

// WatchTransceiverAdded is a free log subscription operation binding the contract event 0xf05962b5774c658e85ed80c91a75af9d66d2af2253dda480f90bce78aff5eda5.
//
// Solidity: event TransceiverAdded(address transceiver, uint256 transceiversNum, uint8 threshold)
func (_Bindings *BindingsFilterer) WatchTransceiverAdded(opts *bind.WatchOpts, sink chan<- *BindingsTransceiverAdded) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "TransceiverAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTransceiverAdded)
				if err := _Bindings.contract.UnpackLog(event, "TransceiverAdded", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseTransceiverAdded(log types.Log) (*BindingsTransceiverAdded, error) {
	event := new(BindingsTransceiverAdded)
	if err := _Bindings.contract.UnpackLog(event, "TransceiverAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsTransceiverRemovedIterator is returned from FilterTransceiverRemoved and is used to iterate over the raw logs and unpacked data for TransceiverRemoved events raised by the Bindings contract.
type BindingsTransceiverRemovedIterator struct {
	Event *BindingsTransceiverRemoved // Event containing the contract specifics and raw log

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
func (it *BindingsTransceiverRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTransceiverRemoved)
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
		it.Event = new(BindingsTransceiverRemoved)
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
func (it *BindingsTransceiverRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTransceiverRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTransceiverRemoved represents a TransceiverRemoved event raised by the Bindings contract.
type BindingsTransceiverRemoved struct {
	Transceiver common.Address
	Threshold   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransceiverRemoved is a free log retrieval operation binding the contract event 0x697a3853515b88013ad432f29f53d406debc9509ed6d9313dcfe115250fcd18f.
//
// Solidity: event TransceiverRemoved(address transceiver, uint8 threshold)
func (_Bindings *BindingsFilterer) FilterTransceiverRemoved(opts *bind.FilterOpts) (*BindingsTransceiverRemovedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "TransceiverRemoved")
	if err != nil {
		return nil, err
	}
	return &BindingsTransceiverRemovedIterator{contract: _Bindings.contract, event: "TransceiverRemoved", logs: logs, sub: sub}, nil
}

// WatchTransceiverRemoved is a free log subscription operation binding the contract event 0x697a3853515b88013ad432f29f53d406debc9509ed6d9313dcfe115250fcd18f.
//
// Solidity: event TransceiverRemoved(address transceiver, uint8 threshold)
func (_Bindings *BindingsFilterer) WatchTransceiverRemoved(opts *bind.WatchOpts, sink chan<- *BindingsTransceiverRemoved) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "TransceiverRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTransceiverRemoved)
				if err := _Bindings.contract.UnpackLog(event, "TransceiverRemoved", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseTransceiverRemoved(log types.Log) (*BindingsTransceiverRemoved, error) {
	event := new(BindingsTransceiverRemoved)
	if err := _Bindings.contract.UnpackLog(event, "TransceiverRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsTransferRedeemedIterator is returned from FilterTransferRedeemed and is used to iterate over the raw logs and unpacked data for TransferRedeemed events raised by the Bindings contract.
type BindingsTransferRedeemedIterator struct {
	Event *BindingsTransferRedeemed // Event containing the contract specifics and raw log

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
func (it *BindingsTransferRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTransferRedeemed)
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
		it.Event = new(BindingsTransferRedeemed)
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
func (it *BindingsTransferRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTransferRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTransferRedeemed represents a TransferRedeemed event raised by the Bindings contract.
type BindingsTransferRedeemed struct {
	Digest [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferRedeemed is a free log retrieval operation binding the contract event 0x504e6efe18ab9eed10dc6501a417f5b12a2f7f2b1593aed9b89f9bce3cf29a91.
//
// Solidity: event TransferRedeemed(bytes32 indexed digest)
func (_Bindings *BindingsFilterer) FilterTransferRedeemed(opts *bind.FilterOpts, digest [][32]byte) (*BindingsTransferRedeemedIterator, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "TransferRedeemed", digestRule)
	if err != nil {
		return nil, err
	}
	return &BindingsTransferRedeemedIterator{contract: _Bindings.contract, event: "TransferRedeemed", logs: logs, sub: sub}, nil
}

// WatchTransferRedeemed is a free log subscription operation binding the contract event 0x504e6efe18ab9eed10dc6501a417f5b12a2f7f2b1593aed9b89f9bce3cf29a91.
//
// Solidity: event TransferRedeemed(bytes32 indexed digest)
func (_Bindings *BindingsFilterer) WatchTransferRedeemed(opts *bind.WatchOpts, sink chan<- *BindingsTransferRedeemed, digest [][32]byte) (event.Subscription, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "TransferRedeemed", digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTransferRedeemed)
				if err := _Bindings.contract.UnpackLog(event, "TransferRedeemed", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseTransferRedeemed(log types.Log) (*BindingsTransferRedeemed, error) {
	event := new(BindingsTransferRedeemed)
	if err := _Bindings.contract.UnpackLog(event, "TransferRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsTransferSentIterator is returned from FilterTransferSent and is used to iterate over the raw logs and unpacked data for TransferSent events raised by the Bindings contract.
type BindingsTransferSentIterator struct {
	Event *BindingsTransferSent // Event containing the contract specifics and raw log

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
func (it *BindingsTransferSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTransferSent)
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
		it.Event = new(BindingsTransferSent)
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
func (it *BindingsTransferSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTransferSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTransferSent represents a TransferSent event raised by the Bindings contract.
type BindingsTransferSent struct {
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
// Solidity: event TransferSent(bytes32 recipient, bytes32 refundAddress, uint256 amount, uint256 fee, uint16 recipientChain, uint64 msgSequence)
func (_Bindings *BindingsFilterer) FilterTransferSent(opts *bind.FilterOpts) (*BindingsTransferSentIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "TransferSent")
	if err != nil {
		return nil, err
	}
	return &BindingsTransferSentIterator{contract: _Bindings.contract, event: "TransferSent", logs: logs, sub: sub}, nil
}

// WatchTransferSent is a free log subscription operation binding the contract event 0xe54e51e42099622516fa3b48e9733581c9dbdcb771cafb093f745a0532a35982.
//
// Solidity: event TransferSent(bytes32 recipient, bytes32 refundAddress, uint256 amount, uint256 fee, uint16 recipientChain, uint64 msgSequence)
func (_Bindings *BindingsFilterer) WatchTransferSent(opts *bind.WatchOpts, sink chan<- *BindingsTransferSent) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "TransferSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTransferSent)
				if err := _Bindings.contract.UnpackLog(event, "TransferSent", log); err != nil {
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
// Solidity: event TransferSent(bytes32 recipient, bytes32 refundAddress, uint256 amount, uint256 fee, uint16 recipientChain, uint64 msgSequence)
func (_Bindings *BindingsFilterer) ParseTransferSent(log types.Log) (*BindingsTransferSent, error) {
	event := new(BindingsTransferSent)
	if err := _Bindings.contract.UnpackLog(event, "TransferSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsTransferSent0Iterator is returned from FilterTransferSent0 and is used to iterate over the raw logs and unpacked data for TransferSent0 events raised by the Bindings contract.
type BindingsTransferSent0Iterator struct {
	Event *BindingsTransferSent0 // Event containing the contract specifics and raw log

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
func (it *BindingsTransferSent0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTransferSent0)
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
		it.Event = new(BindingsTransferSent0)
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
func (it *BindingsTransferSent0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTransferSent0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTransferSent0 represents a TransferSent0 event raised by the Bindings contract.
type BindingsTransferSent0 struct {
	Digest [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferSent0 is a free log retrieval operation binding the contract event 0x3e6ae56314c6da8b461d872f41c6d0bb69317b9d0232805aaccfa45df1a16fa0.
//
// Solidity: event TransferSent(bytes32 indexed digest)
func (_Bindings *BindingsFilterer) FilterTransferSent0(opts *bind.FilterOpts, digest [][32]byte) (*BindingsTransferSent0Iterator, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "TransferSent0", digestRule)
	if err != nil {
		return nil, err
	}
	return &BindingsTransferSent0Iterator{contract: _Bindings.contract, event: "TransferSent0", logs: logs, sub: sub}, nil
}

// WatchTransferSent0 is a free log subscription operation binding the contract event 0x3e6ae56314c6da8b461d872f41c6d0bb69317b9d0232805aaccfa45df1a16fa0.
//
// Solidity: event TransferSent(bytes32 indexed digest)
func (_Bindings *BindingsFilterer) WatchTransferSent0(opts *bind.WatchOpts, sink chan<- *BindingsTransferSent0, digest [][32]byte) (event.Subscription, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "TransferSent0", digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTransferSent0)
				if err := _Bindings.contract.UnpackLog(event, "TransferSent0", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseTransferSent0(log types.Log) (*BindingsTransferSent0, error) {
	event := new(BindingsTransferSent0)
	if err := _Bindings.contract.UnpackLog(event, "TransferSent0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Bindings contract.
type BindingsUpgradedIterator struct {
	Event *BindingsUpgraded // Event containing the contract specifics and raw log

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
func (it *BindingsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUpgraded)
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
		it.Event = new(BindingsUpgraded)
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
func (it *BindingsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUpgraded represents a Upgraded event raised by the Bindings contract.
type BindingsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bindings *BindingsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BindingsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BindingsUpgradedIterator{contract: _Bindings.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bindings *BindingsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BindingsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUpgraded)
				if err := _Bindings.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseUpgraded(log types.Log) (*BindingsUpgraded, error) {
	event := new(BindingsUpgraded)
	if err := _Bindings.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
