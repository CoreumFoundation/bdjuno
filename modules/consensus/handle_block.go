package consensus

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/bdjuno/v3/modules/actions/logging"
	"github.com/forbole/juno/v3/types"

	"github.com/rs/zerolog/log"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// HandleBlock implements modules.Module
func (m *Module) HandleBlock(
	b *tmctypes.ResultBlock, _ *tmctypes.ResultBlockResults, _ []*types.Tx, vals *tmctypes.ResultValidators,
) error {
	if err := m.updateBlockTimeFromGenesis(b); err != nil {
		log.Error().Str("module", "consensus").Int64("height", b.Block.Height).
			Err(err).Msg("error while updating block time from genesis")
	}

	if err := m.updateBlockTimeByValidator(b, vals); err != nil {
		log.Error().Str("module", "consensus").Int64("height", b.Block.Height).
			Err(err).Msg("error while updating block time by validator")
	}

	return nil
}

// updateBlockTimeFromGenesis insert average block time from genesis
func (m *Module) updateBlockTimeFromGenesis(block *tmctypes.ResultBlock) error {
	log.Trace().Str("module", "consensus").Int64("height", block.Block.Height).
		Msg("updating block time from genesis")

	genesis, err := m.db.GetGenesis()
	if err != nil {
		return fmt.Errorf("error while getting genesis: %s", err)
	}
	if genesis == nil {
		return fmt.Errorf("genesis table is empty")
	}

	// Skip if the genesis does not exist
	if genesis == nil {
		return nil
	}

	newBlockTime := block.Block.Time.Sub(genesis.Time).Seconds() / float64(block.Block.Height-genesis.InitialHeight)
	return m.db.SaveAverageBlockTimeGenesis(newBlockTime, block.Block.Height)
}

func (m *Module) updateBlockTimeByValidator(block *tmctypes.ResultBlock, vals *tmctypes.ResultValidators) error {
	if m.expectedProposer != nil && !bytes.Equal(block.Block.ProposerAddress, m.expectedProposer) {
		logging.MissedProposerCounter.WithLabelValues(sdk.ConsAddress(m.expectedProposer).String()).Inc()
	}

	expectedNextProposer := vals.Validators[0]
	if len(vals.Validators) > 1 {
		for _, v := range vals.Validators[1:] {
			if v.ProposerPriority > expectedNextProposer.ProposerPriority {
				expectedNextProposer = v
			}
		}
	}
	m.expectedProposer = expectedNextProposer.Address

	return nil
}
