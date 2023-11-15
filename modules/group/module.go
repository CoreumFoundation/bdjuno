package group

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/bdjuno/v4/database"
	"github.com/forbole/juno/v5/modules"
)

var _ modules.Module = &Module{}

// Module represents the x/asset/ft module
type Module struct {
	cdc codec.Codec
	db  *database.Db
}

// NewModule returns a new Module instance
func NewModule(
	cdc codec.Codec, db *database.Db,
) *Module {
	return &Module{
		cdc: cdc,
		db:  db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "group"
}
