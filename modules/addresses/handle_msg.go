package addresses

import (
	"strings"

	tmtypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	juno "github.com/forbole/juno/v5/types"
	"github.com/samber/lo"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Events) == 0 {
		return nil
	}

	addresses, err := m.collectAddresses(msg, tx)
	if err != nil {
		return err
	}

	// marshal the value properly
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

func (m *Module) collectAddresses(msg sdk.Msg, tx *juno.Tx) ([]string, error) {
	// get the involved addresses with general parser first
	messageAddresses, err := m.messageParser(m.cdc, msg)
	if err != nil {
		return nil, err
	}

	addresses := make(map[string]struct{})
	for _, address := range messageAddresses {
		addresses[address] = struct{}{}
	}
	// add address from event values
	addBech32EventValues(addresses, tx.Events)

	return lo.Keys(addresses), nil
}

func addBech32EventValues(addressSet map[string]struct{}, events []tmtypes.Event) {
	for _, ev := range sdk.StringifyEvents(events) {
		for _, attrItem := range ev.Attributes {
			address := strings.Trim(strings.TrimSpace(attrItem.Value), `"`)
			if !isBech32Address(address) {
				continue
			}
			addressSet[address] = struct{}{}
		}
	}
}

func isBech32Address(address string) bool {
	_, err := sdk.AccAddressFromBech32(address)
	return err == nil
}
