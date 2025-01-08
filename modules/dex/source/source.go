package source

import (
	dextypes "github.com/CoreumFoundation/coreum/v5/x/dex/types"
)

type Source interface {
	GetParams(height int64) (dextypes.Params, error)
}
