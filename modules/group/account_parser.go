package group

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/group"
)

// MessagesParser returns the list of all the accounts involved in the given
// message if it's related to the x/assetft module
func MessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {
	case *group.MsgExec:
		return []string{msg.Executor}, nil
	case *group.MsgCreateGroup:
		addresses := make([]string, len(msg.Members))
		for _, member := range msg.Members {
			addresses = append(addresses, member.Address)
		}
		return addresses, nil
	case *group.MsgCreateGroupWithPolicy:
		addresses := make([]string, len(msg.Members))
		for _, member := range msg.Members {
			addresses = append(addresses, member.Address)
		}
		return addresses, nil
	case *group.MsgSubmitProposal:
		return []string{msg.GroupPolicyAddress}, nil
	case *group.MsgUpdateGroupMembers:
		addresses := make([]string, len(msg.MemberUpdates))
		for _, member := range msg.MemberUpdates {
			addresses = append(addresses, member.Address)
		}
		return addresses, nil
	case *group.MsgUpdateGroupPolicyAdmin:
		return []string{msg.GroupPolicyAddress, msg.NewAdmin}, nil
	case *group.MsgUpdateGroupPolicyDecisionPolicy:
		return []string{msg.GroupPolicyAddress}, nil
	case *group.MsgUpdateGroupPolicyMetadata:
		return []string{msg.GroupPolicyAddress}, nil
	}

	return nil, nil
}
