package source

import (
	assetfttypes "github.com/CoreumFoundation/coreum/v5/x/asset/ft/types"
)

type Source interface {
	GetParams(height int64) (assetfttypes.Params, error)
}
