package consensus

import (
	"github.com/forbole/bdjuno/v3/database"
	"github.com/forbole/juno/v3/modules"
	tmtypes "github.com/tendermint/tendermint/types"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
	_ modules.GenesisModule            = &Module{}
	_ modules.BlockModule              = &Module{}
)

type proposer struct {
	Height          int64
	CurrentProposer tmtypes.Address
	NextProposer    tmtypes.Address
}

// Module implements the consensus utils
type Module struct {
	db            *database.Db
	proposerQueue chan proposer
}

// NewModule builds a new Module instance
func NewModule(db *database.Db) *Module {
	return &Module{
		db:            db,
		proposerQueue: make(chan proposer),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "consensus"
}
