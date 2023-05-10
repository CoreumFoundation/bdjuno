package consensus

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/bdjuno/v3/modules/actions/logging"
	tmtypes "github.com/tendermint/tendermint/types"
)

func (m *Module) RunAsyncOperations() {
	const maxCacheSize = 100

	realProposers := map[int64]tmtypes.Address{}
	expectedProposers := map[int64]tmtypes.Address{}

	for p := range m.proposerQueue {
		// This handles the case when block is received in order.
		if expectedProposer := expectedProposers[p.Height]; expectedProposer != nil {
			delete(expectedProposers, p.Height)
			updateProposerMetric(expectedProposer, p.CurrentProposer)
		} else {
			realProposers[p.Height] = p.CurrentProposer
		}

		// This handles the case when block is received out of order.
		if realProposer := realProposers[p.Height+1]; realProposer != nil {
			delete(realProposers, p.Height+1)
			updateProposerMetric(p.NextProposer, realProposer)
		} else {
			expectedProposers[p.Height+1] = p.NextProposer
		}

		// Protection against memory leaks when blocks are missed by the indexer.
		// This is naive approach which might lead to loosing some measures, but it won't break statistics.
		if len(realProposers) > maxCacheSize {
			realProposers = map[int64]tmtypes.Address{}
		}
		if len(expectedProposers) > maxCacheSize {
			expectedProposers = map[int64]tmtypes.Address{}
		}
	}
}

func updateProposerMetric(expected, real tmtypes.Address) {
	var value float64
	if bytes.Equal(expected, real) {
		value = 1.0
	}
	logging.ProposalCounter.WithLabelValues(sdk.ConsAddress(expected).String()).Observe(value)
}
