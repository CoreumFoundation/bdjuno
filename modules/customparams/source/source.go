package source

import (
	customparamstypes "github.com/CoreumFoundation/coreum/v5/x/customparams/types"
)

type Source interface {
	GetParams(height int64) (customparamstypes.StakingParams, error)
}
