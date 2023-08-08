package wasm

import (
	"strings"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	"github.com/gogo/protobuf/proto"
	"github.com/samber/lo"
	tmtypes "github.com/tendermint/tendermint/abci/types"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Events) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) { //nolint:gocritic //bdjuno style
	case *wasmtypes.MsgInstantiateContract,
		*wasmtypes.MsgInstantiateContract2,
		*wasmtypes.MsgMigrateContract,
		*wasmtypes.MsgExecuteContract:
		return m.handleWasmRelatedAddress(index, cosmosMsg, tx)
	}

	return nil
}

func (m *Module) handleWasmRelatedAddress(index int, msg sdk.Msg, tx *juno.Tx) error {
	eventsAddresses := m.findBech32EventValues(tx.Events)
	if len(eventsAddresses) == 0 {
		return nil
	}

	// get the involved addresses with general parser first
	addresses, err := m.messageParser(m.cdc, msg)
	if err != nil {
		return err
	}

	// we join and then do the Uniq since the receivers might be duplicated
	addresses = lo.Uniq(append(addresses, eventsAddresses...))

	// Marshal the value properly
	bz, err := m.cdc.MarshalJSON(msg)
	if err != nil {
		return err
	}

	return m.db.SaveMessage(juno.NewMessage(
		tx.TxHash,
		index,
		proto.MessageName(msg),
		string(bz),
		addresses,
		tx.Height,
	))
}

func (m *Module) findBech32EventValues(events []tmtypes.Event) []string {
	values := make([]string, 0)
	for _, ev := range sdk.StringifyEvents(events) {
		for _, attrItem := range ev.Attributes {
			address := strings.Trim(strings.TrimSpace(attrItem.Value), `"`)
			if !m.isBech32Address(address) {
				continue
			}
			values = append(values, address)
		}
	}

	return values
}

func (m *Module) isBech32Address(address string) bool {
	_, err := sdk.AccAddressFromBech32(address)
	return err == nil
}
