package ocrcommon

import (
	"time"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink/core/config"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/p2pkey"
)

type Config interface {
	LogSQL() bool
	EvmGasLimitDefault() uint64
	JobPipelineResultWriteQueueDepth() uint64
	OCRBlockchainTimeout() time.Duration
	OCRContractConfirmations() uint16
	OCRContractPollInterval() time.Duration
	OCRContractSubscribeInterval() time.Duration
	OCRContractTransmitterTransmitTimeout() time.Duration
	OCRDatabaseTimeout() time.Duration
	OCRDefaultTransactionQueueDepth() uint32
	OCRKeyBundleID() (string, error)
	OCRObservationGracePeriod() time.Duration
	OCRObservationTimeout() time.Duration
	OCRTraceLogging() bool
	OCRTransmitterAddress() (ethkey.EIP55Address, error)
	P2PBootstrapPeers() ([]string, error)
	P2PPeerID() p2pkey.PeerID
	P2PV2Bootstrappers() []commontypes.BootstrapperLocator
	FlagsContractAddress() string
	ChainType() config.ChainType
}
