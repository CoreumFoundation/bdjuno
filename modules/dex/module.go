package dex

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/callisto/v4/database"
	dexsource "github.com/forbole/callisto/v4/modules/dex/source"
	"github.com/forbole/juno/v6/modules"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
)

// Module represents the x/dex module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source dexsource.Source
}

// NewModule returns a new Module instance
func NewModule(
	source dexsource.Source,
	cdc codec.Codec, db *database.Db,
) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		source: source,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "dex"
}
