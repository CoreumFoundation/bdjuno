package source

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type Source interface {
	GetValidator(height int64, valOper string) (stakingtypes.Validator, error)
	GetValidatorsWithStatus(height int64, status string) ([]stakingtypes.Validator, error)
	GetDelegationsWithPagination(delegator string, pagination *query.PageRequest) (*stakingtypes.QueryDelegatorDelegationsResponse, error)
	GetRedelegations(request *stakingtypes.QueryRedelegationsRequest) (*stakingtypes.QueryRedelegationsResponse, error)
	GetPool(height int64) (stakingtypes.Pool, error)
	GetParams() (stakingtypes.Params, error)
	GetUnbondingDelegations(delegator string, pagination *query.PageRequest) (*stakingtypes.QueryDelegatorUnbondingDelegationsResponse, error)
	GetValidatorDelegationsWithPagination(validator string, pagination *query.PageRequest) (*stakingtypes.QueryValidatorDelegationsResponse, error)
	GetUnbondingDelegationsFromValidator(validator string, pagination *query.PageRequest) (*stakingtypes.QueryValidatorUnbondingDelegationsResponse, error)
}
