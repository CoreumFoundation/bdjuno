package handlers

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/modules/actions/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
)

func ValidatorDelegation(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	log.Debug().Str("address", payload.GetAddress()).
		Msg("executing validator delegation action")

	// Get validator's total delegations
	res, err := ctx.Sources.StakingSource.GetValidatorDelegationsWithPagination(payload.GetAddress(), payload.GetPagination())
	if err != nil {
		return nil, fmt.Errorf("error while getting validator delegations: %s", err)
	}

	delegations := make([]types.Delegation, len(res.DelegationResponses))
	for index, del := range res.DelegationResponses {
		delegations[index] = types.Delegation{
			DelegatorAddress: del.Delegation.DelegatorAddress,
			ValidatorAddress: del.Delegation.ValidatorAddress,
			Coins:            types.ConvertCoins([]sdk.Coin{del.Balance}),
		}
	}

	return types.DelegationResponse{
		Delegations: delegations,
		Pagination:  res.Pagination,
	}, nil
}
