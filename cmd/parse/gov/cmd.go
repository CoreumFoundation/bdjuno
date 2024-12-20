package gov

import (
	"github.com/forbole/callisto/v4/modules"
	parsecmdtypes "github.com/forbole/juno/v6/cmd/parse/types"
	"github.com/forbole/juno/v6/modules/messages"
	"github.com/spf13/cobra"
)

// NewGovCmd returns the Cobra command allowing to fix various things related to the x/gov module
func NewGovCmd(parseConfig *parsecmdtypes.Config, parser messages.MessageAddressesParser) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gov",
		Short: "Fix things related to the x/gov module",
	}

	cmd.AddCommand(
		proposalCmd(parseConfig, modules.UniqueAddressesParser(parser)),
		paramsCmd(parseConfig, modules.UniqueAddressesParser(parser)),
	)

	return cmd
}
