package types

import (
	"github.com/smartcontractkit/chainlink-starknet/pkg/relay/starknet/db"

	"github.com/smartcontractkit/chainlink/core/chains"
)

type ORM interface {
	// Chain(string, ...pg.QOpt) (DBChain, error)
	// Chains(offset, limit int, qopts ...pg.QOpt) ([]DBChain, int, error)
	// CreateChain(id string, config *db.ChainCfg, qopts ...pg.QOpt) (DBChain, error)
	// UpdateChain(id string, enabled bool, config *db.ChainCfg, qopts ...pg.QOpt) (DBChain, error)
	// DeleteChain(id string, qopts ...pg.QOpt) error
	// GetChainsByIDs(ids []string) (chains []DBChain, err error)
	// EnabledChains(...pg.QOpt) ([]DBChain, error)

	// CreateNode(db.Node, ...pg.QOpt) (db.Node, error)
	// DeleteNode(int32, ...pg.QOpt) error
	// GetNodesByChainIDs(chainIDs []string, qopts ...pg.QOpt) (nodes []db.Node, err error)
	// Node(int32, ...pg.QOpt) (db.Node, error)
	// NodeNamed(string, ...pg.QOpt) (db.Node, error)
	// Nodes(offset, limit int, qopts ...pg.QOpt) (nodes []db.Node, count int, err error)
	// NodesForChain(chainID string, offset, limit int, qopts ...pg.QOpt) (nodes []db.Node, count int, err error)

	// SetupNodes([]db.Node, []string) error

	// StoreString(chainID string, key, val string) error
	// Clear(chainID string, key string) error
}

type DBChain = chains.DBChain[string, *db.ChainCfg]
