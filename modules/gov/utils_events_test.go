package gov_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/forbole/callisto/v4/modules/gov"
	"github.com/stretchr/testify/require"
)

func TestWeightVoteOptionFromEvents(t *testing.T) {
	tests := []struct {
		name      string
		events    sdk.StringEvents
		expected  []govtypesv1.WeightedVoteOption
		shouldErr bool
	}{
		{
			"json option from vote event returns properly",
			sdk.StringEvents{
				sdk.StringEvent{
					Type: "vote",
					Attributes: []sdk.Attribute{
						sdk.NewAttribute(govtypes.AttributeKeyOption, "{\"option\":1,\"weight\":\"1.000000000000000000\"}"),
					},
				},
			},
			[]govtypesv1.WeightedVoteOption{{Option: govtypesv1.OptionYes, Weight: "1.000000000000000000"}},
			false,
		},
		{
			"json options from vote event returns properly",
			sdk.StringEvents{
				sdk.StringEvent{
					Type: "vote",
					Attributes: []sdk.Attribute{
						sdk.NewAttribute(govtypes.AttributeKeyOption, "[{\"option\":3,\"weight\":\"0.300000000000000000\"},{\"option\":4,\"weight\":\"0.700000000000000000\"}]"),
					},
				},
			},
			[]govtypesv1.WeightedVoteOption{
				{Option: govtypesv1.OptionNo, Weight: "0.300000000000000000"},
				{Option: govtypesv1.OptionNoWithVeto, Weight: "0.700000000000000000"},
			},
			false,
		},
		{
			"string option from vote event returns properly",
			sdk.StringEvents{
				sdk.StringEvent{
					Type: "vote",
					Attributes: []sdk.Attribute{
						sdk.NewAttribute(govtypes.AttributeKeyOption, "option:VOTE_OPTION_NO weight:\"1.000000000000000000\""),
					},
				},
			},
			[]govtypesv1.WeightedVoteOption{{Option: govtypesv1.OptionNo, Weight: "1.000000000000000000"}},
			false,
		},
		{
			"invalid option from vote event returns error",
			sdk.StringEvents{
				sdk.StringEvent{
					Type: "vote",
					Attributes: []sdk.Attribute{
						sdk.NewAttribute("other", "value"),
					},
				},
			},
			nil,
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := gov.WeightVoteOptionsFromEvents(test.events)
			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.Equal(t, test.expected, result)
			}
		})
	}
}
