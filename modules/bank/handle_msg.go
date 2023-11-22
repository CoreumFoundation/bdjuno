package bank

import (
	"fmt"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"
)

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if err := m.UpdateAccountsBalances(index, tx); err != nil {
		return err
	}

	return nil
}

func (m *Module) UpdateAccountsBalances(index int, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	if err := m.updateBalanceForEventType(index, tx, banktypes.EventTypeCoinReceived); err != nil {
		return err
	}

	return m.updateBalanceForEventType(index, tx, banktypes.EventTypeCoinSpent)
}

func (m *Module) updateBalanceForEventType(index int, tx *juno.Tx, eventType string) error {
	accountAttribute := banktypes.AttributeKeySpender
	if eventType == banktypes.EventTypeCoinReceived {
		accountAttribute = banktypes.AttributeKeyReceiver
	}

	events := FindAllEventsByType(index, tx, eventType)
	for _, event := range events {
		account, err := tx.FindAttributeByKey(event, accountAttribute)
		if err != nil {
			return err
		}

		coinString, err := tx.FindAttributeByKey(event, sdk.AttributeKeyAmount)
		if err != nil {
			return err
		}

		coin, err := sdk.ParseCoinNormalized(coinString)
		if err != nil {
			return err
		}

		quriedBalance, err := m.keeper.GetAccountDenomBalance(account, coin.Denom)
		if err != nil {
			return err
		}

		if quriedBalance == nil {
			return fmt.Errorf("query balance return nil, account: %s, denom:%s", account, coin.Denom)
		}

		if quriedBalance.Amount.IsZero() {
			if err := m.db.DeleteAccountDenomBalance(account, *quriedBalance); err != nil {
				return err
			}
		} else {
			if err := m.db.SaveAccountDenomBalance(account, *quriedBalance); err != nil {
				return err
			}
		}
	}

	return nil
}

func FindAllEventsByType(index int, tx *juno.Tx, eventType string) []sdk.StringEvent {
	var list []sdk.StringEvent
	for _, ev := range tx.Logs[index].Events {
		if ev.Type == eventType {
			list = append(list, ev)
		}
	}
	return list
}
