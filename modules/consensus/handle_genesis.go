package consensus

import (
	"encoding/json"
	"fmt"

	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/forbole/callisto/v4/types"
	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.Module
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, _ map[string]json.RawMessage) error {
	log.Debug().Str("module", "consensus").Msg("parsing genesis")

	// Save the genesis time
	err := m.db.SaveGenesis(types.NewGenesis(doc.ChainID, doc.GenesisTime, doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error while storing genesis time: %s", err)
	}

	return nil
}
