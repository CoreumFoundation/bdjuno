package handlers

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/modules/actions/types"

	"github.com/rs/zerolog/log"
)

func ValidatorCommissionAmountHandler(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	log.Debug().Str("address", payload.GetAddress()).
		Msg("executing validator commission action")

	// Get validator total commission value
	commission, err := ctx.Sources.DistrSource.ValidatorCommission(payload.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("error while getting validator commission: %s", err)
	}

	return types.ValidatorCommissionAmount{
		Coins: types.ConvertDecCoins(commission),
	}, nil
}
