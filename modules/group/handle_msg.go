package group

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	"github.com/forbole/bdjuno/v4/utils"
	juno "github.com/forbole/juno/v5/types"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Events) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *group.MsgExec,
		*group.MsgCreateGroup,
		*group.MsgUpdateGroupMembers,
		*group.MsgSubmitProposal,
		*group.MsgUpdateGroupPolicyDecisionPolicy,
		*group.MsgUpdateGroupPolicyAdmin,
		*group.MsgLeaveGroup,
		*group.MsgCreateGroupWithPolicy,
		*group.MsgVote,
		*group.MsgUpdateGroupMetadata,
		*group.MsgUpdateGroupAdmin,
		*group.MsgCreateGroupPolicy,
		*group.MsgWithdrawProposal,
		*group.MsgUpdateGroupPolicyMetadata:
		return utils.HandleMessageWithAddresses(index, cosmosMsg, tx, m.messageParser, m.cdc, m.db)
	}

	return nil
}
