package wasm

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/bdjuno/v4/utils"
	juno "github.com/forbole/juno/v5/types"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Events) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *wasmtypes.MsgInstantiateContract,
		*wasmtypes.MsgInstantiateContract2,
		*wasmtypes.MsgMigrateContract,
		*wasmtypes.MsgExecuteContract:
		return utils.HandleMessageWithAddresses(index, cosmosMsg, tx, m.messageParser, m.cdc, m.db)
	}

	return nil
}
