package remote

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	"github.com/forbole/bdjuno/v4/utils"
)

// DelegatorTotalRewards implements distrsource.Source
func (s Source) DelegatorTotalRewards(delegator string) ([]distrtypes.DelegationDelegatorReward, error) {
	ctx := utils.GetHeightRequestContext(s.Ctx)
	res, err := s.distrClient.DelegationTotalRewards(
		ctx,
		&distrtypes.QueryDelegationTotalRewardsRequest{DelegatorAddress: delegator},
	)
	if err != nil {
		return nil, fmt.Errorf("error while getting delegation total rewards for for delegator %s: %s", delegator, err)
	}

	return res.Rewards, nil
}

// DelegatorWithdrawAddress implements distrsource.Source
func (s Source) DelegatorWithdrawAddress(delegator string) (string, error) {
	ctx := utils.GetHeightRequestContext(s.Ctx)
	res, err := s.distrClient.DelegatorWithdrawAddress(
		ctx,
		&distrtypes.QueryDelegatorWithdrawAddressRequest{DelegatorAddress: delegator},
	)
	if err != nil {
		return "", err
	}

	return res.WithdrawAddress, nil
}

// ValidatorCommission implements distrsource.Source
func (s Source) ValidatorCommission(valOperAddr string) (sdk.DecCoins, error) {
	ctx := utils.GetHeightRequestContext(s.Ctx)
	res, err := s.distrClient.ValidatorCommission(
		ctx,
		&distrtypes.QueryValidatorCommissionRequest{ValidatorAddress: valOperAddr},
	)
	if err != nil {
		return nil, err
	}

	return res.Commission.Commission, nil
}
