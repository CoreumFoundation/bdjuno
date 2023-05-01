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

// Module implements the consensus utils
type Module struct {
	db               *database.Db
	expectedProposer tmtypes.Address
}

// NewModule builds a new Module instance
func NewModule(db *database.Db) *Module {
	return &Module{
		db: db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "consensus"
}
