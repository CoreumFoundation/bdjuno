package utils

import (
	"strings"

	tmtypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/bdjuno/v4/database"
	junomessages "github.com/forbole/juno/v5/modules/messages"
	juno "github.com/forbole/juno/v5/types"
	"github.com/gogo/protobuf/proto"
	"github.com/samber/lo"
)

// HandleMessageWithAddresses handle messages and takes addresses from events.
func HandleMessageWithAddresses(
	index int,
	msg sdk.Msg,
	tx *juno.Tx,
	messageParser junomessages.MessageAddressesParser,
	cdc codec.Codec, db *database.Db,
) error {
	addresses, err := collectAddresses(cdc, messageParser, msg, tx)
	if err != nil {
		return err
	}

	// marshal the value properly
	bz, err := cdc.MarshalJSON(msg)
	if err != nil {
		return err
	}

	return db.SaveMessage(juno.NewMessage(
		tx.TxHash,
		index,
		proto.MessageName(msg),
		string(bz),
		addresses,
		tx.Height,
	))
}

func collectAddresses(
	cdc codec.Codec,
	messageParser junomessages.MessageAddressesParser,
	msg sdk.Msg, tx *juno.Tx,
) ([]string, error) {
	// get the involved addresses with general parser first
	messageAddresses, err := messageParser(cdc, msg)
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
