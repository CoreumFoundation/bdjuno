package customparams

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/callisto/v4/database"
	customparamssource "github.com/forbole/callisto/v4/modules/customparams/source"
	"github.com/forbole/juno/v6/modules"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
)

// Module represents the x/customparams module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source customparamssource.Source
}

// NewModule returns a new Module instance
func NewModule(
	source customparamssource.Source,
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
	return "customparams"
}
