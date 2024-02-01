package source

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

type Source interface {
	ValidatorCommission(valOperAddr string) (sdk.DecCoins, error)
	DelegatorTotalRewards(delegator string) ([]distrtypes.DelegationDelegatorReward, error)
	DelegatorWithdrawAddress(delegator string) (string, error)
	CommunityPool(height int64) (sdk.DecCoins, error)
	Params(height int64) (distrtypes.Params, error)
}
