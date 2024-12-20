package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/callisto/v4/database"
	authsource "github.com/forbole/callisto/v4/modules/auth/source"
	"github.com/forbole/juno/v6/modules"
	"github.com/forbole/juno/v6/modules/messages"
)

var (
	_ modules.Module             = &Module{}
	_ modules.GenesisModule      = &Module{}
	_ modules.MessageModule      = &Module{}
	_ modules.AuthzMessageModule = &Module{}
)

// Module represents the x/auth module
type Module struct {
	cdc            codec.Codec
	db             *database.Db
	source         authsource.Source
	messagesParser messages.MessageAddressesParser
}

// NewModule builds a new Module instance
func NewModule(source authsource.Source, messagesParser messages.MessageAddressesParser, cdc codec.Codec,
	db *database.Db,
) *Module {
	return &Module{
		messagesParser: messagesParser,
		cdc:            cdc,
		db:             db,
		source:         source,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "auth"
}
