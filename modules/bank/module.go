package bank

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/forbole/bdjuno/v4/database"
	"github.com/forbole/bdjuno/v4/modules/bank/source"

	junomessages "github.com/forbole/juno/v5/modules/messages"

	"github.com/forbole/juno/v5/modules"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

// Module represents the x/bank module
type Module struct {
	cdc codec.Codec
	db  *database.Db

	messageParser junomessages.MessageAddressesParser
	keeper        source.Source
	govDenom      string
}

// NewModule returns a new Module instance
func NewModule(
	messageParser junomessages.MessageAddressesParser,
	keeper source.Source,
	cdc codec.Codec,
	db *database.Db,
	addressPrefix string,
) *Module {
	return &Module{
		cdc:           cdc,
		db:            db,
		messageParser: messageParser,
		keeper:        keeper,
		govDenom:      getGovTokenFromAddressPrefix(addressPrefix),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "bank"
}

func getGovTokenFromAddressPrefix(addressPrefix string) string {
	switch addressPrefix {
	case "core":
		return "ucore"
	case "testcore":
		return "utestcore"
	case "devcore":
		return "udevcore"
	default:
		return ""
	}
}
