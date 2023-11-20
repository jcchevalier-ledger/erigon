// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package ethconfig

import (
	"time"

	"github.com/c2h5oh/datasize"
	"github.com/ledgerwatch/erigon/erigon-lib/chain"
	libcommon "github.com/ledgerwatch/erigon/erigon-lib/common"
	"github.com/ledgerwatch/erigon/consensus/ethash/ethashcfg"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/eth/gasprice/gaspricecfg"

	"github.com/ledgerwatch/erigon/ethdb/prune"
	"github.com/ledgerwatch/erigon/params"
)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                        *types.Genesis `toml:",omitempty"`
		NetworkID                      uint64
		EthDiscoveryURLs               []string
		Prune                          prune.Mode
		BatchSize                      datasize.ByteSize
		ImportMode                     bool
		BadBlockHash                   libcommon.Hash
		Snapshot                       BlocksFreezing
		BlockDownloaderWindow          int
		ExternalSnapshotDownloaderAddr string
		Whitelist                      map[uint64]libcommon.Hash `toml:"-"`
		Miner                          params.MiningConfig
		Ethash                         ethashcfg.Config
		Clique                         params.ConsensusSnapshotConfig
		Aura                           chain.AuRaConfig
		TxPool                         DeprecatedTxPoolConfig
		GPO                            gaspricecfg.Config
		RPCGasCap                      uint64  `toml:",omitempty"`
		RPCTxFeeCap                    float64 `toml:",omitempty"`
		StateStream                    bool
		BodyDownloadTimeoutSeconds     int
		SyncLoopThrottle               time.Duration
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkID = c.NetworkID
	enc.EthDiscoveryURLs = c.EthDiscoveryURLs
	enc.Prune = c.Prune
	enc.BatchSize = c.BatchSize
	enc.ImportMode = c.ImportMode
	enc.BadBlockHash = c.BadBlockHash
	enc.Snapshot = c.Snapshot
	enc.ExternalSnapshotDownloaderAddr = c.ExternalSnapshotDownloaderAddr
	enc.Whitelist = c.Whitelist
	enc.Miner = c.Miner
	enc.Ethash = c.Ethash
	enc.Clique = c.Clique
	enc.Aura = c.Aura
	enc.TxPool = c.DeprecatedTxPool
	enc.GPO = c.GPO
	enc.RPCGasCap = c.RPCGasCap
	enc.RPCTxFeeCap = c.RPCTxFeeCap
	enc.StateStream = c.StateStream
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                        *types.Genesis `toml:",omitempty"`
		NetworkID                      *uint64
		EthDiscoveryURLs               []string
		Prune                          *prune.Mode
		BatchSize                      *datasize.ByteSize
		ImportMode                     *bool
		BadBlockHash                   *libcommon.Hash
		Snapshot                       *BlocksFreezing
		BlockDownloaderWindow          *int
		ExternalSnapshotDownloaderAddr *string
		Whitelist                      map[uint64]libcommon.Hash `toml:"-"`
		Miner                          *params.MiningConfig
		Ethash                         *ethashcfg.Config
		Clique                         *params.ConsensusSnapshotConfig
		Aura                           *chain.AuRaConfig
		TxPool                         *DeprecatedTxPoolConfig
		GPO                            *gaspricecfg.Config
		RPCGasCap                      *uint64  `toml:",omitempty"`
		RPCTxFeeCap                    *float64 `toml:",omitempty"`
		StateStream                    *bool
		BodyDownloadTimeoutSeconds     *int
		SyncLoopThrottle               *time.Duration
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkID != nil {
		c.NetworkID = *dec.NetworkID
	}
	if dec.EthDiscoveryURLs != nil {
		c.EthDiscoveryURLs = dec.EthDiscoveryURLs
	}
	if dec.Prune != nil {
		c.Prune = *dec.Prune
	}
	if dec.BatchSize != nil {
		c.BatchSize = *dec.BatchSize
	}
	if dec.ImportMode != nil {
		c.ImportMode = *dec.ImportMode
	}
	if dec.BadBlockHash != nil {
		c.BadBlockHash = *dec.BadBlockHash
	}
	if dec.Snapshot != nil {
		c.Snapshot = *dec.Snapshot
	}
	if dec.ExternalSnapshotDownloaderAddr != nil {
		c.ExternalSnapshotDownloaderAddr = *dec.ExternalSnapshotDownloaderAddr
	}
	if dec.Whitelist != nil {
		c.Whitelist = dec.Whitelist
	}
	if dec.Miner != nil {
		c.Miner = *dec.Miner
	}
	if dec.Ethash != nil {
		c.Ethash = *dec.Ethash
	}
	if dec.Clique != nil {
		c.Clique = *dec.Clique
	}
	if dec.Aura != nil {
		c.Aura = *dec.Aura
	}
	if dec.TxPool != nil {
		c.DeprecatedTxPool = *dec.TxPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.RPCGasCap != nil {
		c.RPCGasCap = *dec.RPCGasCap
	}
	if dec.RPCTxFeeCap != nil {
		c.RPCTxFeeCap = *dec.RPCTxFeeCap
	}
	if dec.StateStream != nil {
		c.StateStream = *dec.StateStream
	}
	return nil
}
