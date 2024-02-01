package remote

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/forbole/bdjuno/v4/utils"
)

// GetDelegationsWithPagination implements stakingsource.Source
func (s Source) GetDelegationsWithPagination(
	delegator string, pagination *query.PageRequest,
) (*stakingtypes.QueryDelegatorDelegationsResponse, error) {
	ctx := utils.GetHeightRequestContext(s.Ctx)
	res, err := s.stakingClient.DelegatorDelegations(
		ctx,
		&stakingtypes.QueryDelegatorDelegationsRequest{
			DelegatorAddr: delegator,
			Pagination:    pagination,
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetUnbondingDelegations implements stakingsource.Source
func (s Source) GetUnbondingDelegations(delegator string, pagination *query.PageRequest) (*stakingtypes.QueryDelegatorUnbondingDelegationsResponse, error) {
	ctx := utils.GetHeightRequestContext(s.Ctx)

	unbondingDelegations, err := s.stakingClient.DelegatorUnbondingDelegations(
		ctx,
		&stakingtypes.QueryDelegatorUnbondingDelegationsRequest{
			DelegatorAddr: delegator,
			Pagination:    pagination,
		},
	)
	if err != nil {
		return nil, err
	}

	return unbondingDelegations, nil
}

// GetRedelegations implements stakingsource.Source
func (s Source) GetRedelegations(request *stakingtypes.QueryRedelegationsRequest) (*stakingtypes.QueryRedelegationsResponse, error) {
	ctx := utils.GetHeightRequestContext(s.Ctx)

	redelegations, err := s.stakingClient.Redelegations(ctx, request)
	if err != nil {
		return nil, err
	}
	return redelegations, nil
}

// GetValidatorDelegationsWithPagination implements stakingsource.Source
func (s Source) GetValidatorDelegationsWithPagination(
	validator string, pagination *query.PageRequest,
) (*stakingtypes.QueryValidatorDelegationsResponse, error) {
	ctx := utils.GetHeightRequestContext(s.Ctx)

	res, err := s.stakingClient.ValidatorDelegations(
		ctx,
		&stakingtypes.QueryValidatorDelegationsRequest{
			ValidatorAddr: validator,
			Pagination:    pagination,
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetUnbondingDelegationsFromValidator implements stakingsource.Source
func (s Source) GetUnbondingDelegationsFromValidator(
	validator string, pagination *query.PageRequest,
) (*stakingtypes.QueryValidatorUnbondingDelegationsResponse, error) {
	ctx := utils.GetHeightRequestContext(s.Ctx)

	unbondingDelegations, err := s.stakingClient.ValidatorUnbondingDelegations(
		ctx,
		&stakingtypes.QueryValidatorUnbondingDelegationsRequest{
			ValidatorAddr: validator,
			Pagination:    pagination,
		},
	)
	if err != nil {
		return nil, err
	}

	return unbondingDelegations, nil
}
