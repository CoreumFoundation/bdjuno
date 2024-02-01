package handlers

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/modules/actions/types"

	"github.com/rs/zerolog/log"
)

func DelegationRewardHandler(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	log.Debug().Str("address", payload.GetAddress()).
		Msg("executing delegation rewards action")

	// Get delegator's total rewards
	rewards, err := ctx.Sources.DistrSource.DelegatorTotalRewards(payload.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("error while getting delegator total rewards: %s", err)
	}

	delegationRewards := make([]types.DelegationReward, len(rewards))
	for index, rew := range rewards {
		delegationRewards[index] = types.DelegationReward{
			Coins:            types.ConvertDecCoins(rew.Reward),
			ValidatorAddress: rew.ValidatorAddress,
		}
	}

	return delegationRewards, nil
}
