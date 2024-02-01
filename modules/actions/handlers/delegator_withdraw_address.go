package handlers

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/modules/actions/types"

	"github.com/rs/zerolog/log"
)

func DelegatorWithdrawAddressHandler(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	log.Debug().Str("address", payload.GetAddress()).
		Msg("executing delegator withdraw address action")

	// Get delegator's total rewards
	withdrawAddress, err := ctx.Sources.DistrSource.DelegatorWithdrawAddress(payload.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("error while getting delegator withdraw address: %s", err)
	}

	return types.Address{
		Address: withdrawAddress,
	}, nil
}
