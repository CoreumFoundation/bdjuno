package customparams

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/types"
	"github.com/rs/zerolog/log"
)

// UpdateParams gets the updated params and stores them inside the database
func (m *Module) UpdateParams(height int64) error {
	log.Debug().Str("module", "customparams").Int64("height", height).
		Msg("updating params")

	stakingParams, err := m.source.GetParams(height)
	if err != nil {
		return fmt.Errorf("error while getting params: %s", err)
	}

	return m.db.SaveCustomParamsParams(types.NewCustomParamsParams(types.CustomParamsStakingParams(stakingParams), height))
}
