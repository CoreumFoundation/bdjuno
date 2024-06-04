package customparams

import (
	"encoding/json"
	"fmt"

	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/forbole/bdjuno/v4/types"
	"github.com/rs/zerolog/log"

	customparamstypes "github.com/CoreumFoundation/coreum/v4/x/customparams/types"
)

// HandleGenesis implements GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "customparams").Msg("parsing genesis")

	// Read the genesis state
	var genState customparamstypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[customparamstypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while unmarshaling customparams state: %s", err)
	}

	// Save the params
	err = m.db.SaveCustomParamsParams(types.NewCustomParamsParams(types.CustomParamsStakingParams(genState.StakingParams), doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error while storing genesis customparams staking params: %s", err)
	}

	return nil
}
