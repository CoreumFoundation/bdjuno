package assetft

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/callisto/v4/database"
	assetftsource "github.com/forbole/callisto/v4/modules/assetft/source"
	"github.com/forbole/juno/v6/modules"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
)

// Module represents the x/asset/ft module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source assetftsource.Source
}

// NewModule returns a new Module instance
func NewModule(
	source assetftsource.Source,
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
	return "assetft"
}
