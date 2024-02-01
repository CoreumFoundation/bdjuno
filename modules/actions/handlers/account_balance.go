package handlers

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/modules/actions/types"

	"github.com/rs/zerolog/log"
)

func AccountBalanceHandler(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	log.Debug().Str("address", payload.GetAddress()).
		Msg("executing account balance action")

	balance, err := ctx.Sources.BankSource.GetAccountBalance(payload.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("error while getting account balance: %s", err)
	}

	return types.Balance{
		Coins: types.ConvertCoins(balance),
	}, nil
}
