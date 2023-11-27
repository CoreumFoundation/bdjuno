package gov

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CoreumFoundation/coreum/v2/app"
	"github.com/CoreumFoundation/coreum/v2/pkg/config"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gcutils "github.com/cosmos/cosmos-sdk/x/gov/client/utils"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/forbole/bdjuno/v3/types"
	"github.com/forbole/juno/v3/node/remote"
	juno "github.com/forbole/juno/v3/types"
)

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *govtypes.MsgSubmitProposal:
		return m.handleMsgSubmitProposal(tx, index, cosmosMsg)

	case *govtypes.MsgDeposit:
		return m.handleMsgDeposit(tx, cosmosMsg)

	case *govtypes.MsgVote:
		return m.handleMsgVote(tx, cosmosMsg)
	}

	return nil
}

// handleMsgSubmitProposal allows to properly handle a handleMsgSubmitProposal
func (m *Module) handleMsgSubmitProposal(tx *juno.Tx, index int, msg *govtypes.MsgSubmitProposal) error {
	// Get the proposal id
	event, err := tx.FindEventByType(index, govtypes.EventTypeSubmitProposal)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeSubmitProposal: %s", err)
	}

	id, err := tx.FindAttributeByKey(event, govtypes.AttributeKeyProposalID)
	if err != nil {
		return fmt.Errorf("error while searching for AttributeKeyProposalID: %s", err)
	}

	proposalID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return fmt.Errorf("error while parsing proposal id: %s", err)
	}

	// Get the proposal
	height, err := m.getLatestHeight(tx.Height, proposalID)
	if err != nil {
		return err
	}
	proposal, err := m.source.Proposal(height, proposalID)
	if err != nil {
		return fmt.Errorf("error while getting proposal: %s", err)
	}

	// Unpack the content
	var content govtypes.Content
	err = m.cdc.UnpackAny(proposal.Content, &content)
	if err != nil {
		return fmt.Errorf("error while unpacking proposal content: %s", err)
	}

	// Store the proposal
	proposalObj := types.NewProposal(
		proposal.ProposalId,
		proposal.ProposalRoute(),
		proposal.ProposalType(),
		proposal.GetContent(),
		proposal.Status.String(),
		proposal.SubmitTime,
		proposal.DepositEndTime,
		proposal.VotingStartTime,
		proposal.VotingEndTime,
		msg.Proposer,
	)
	err = m.db.SaveProposals([]types.Proposal{proposalObj})
	if err != nil {
		return err
	}

	// Store the deposit
	deposit := types.NewDeposit(proposal.ProposalId, msg.Proposer, msg.InitialDeposit, tx.Height)
	return m.db.SaveDeposits([]types.Deposit{deposit})
}

// handleMsgDeposit allows to properly handle a handleMsgDeposit
func (m *Module) handleMsgDeposit(tx *juno.Tx, msg *govtypes.MsgDeposit) error {
	deposit, err := m.getDeposit(tx, msg)
	if err != nil {
		return err
	}

	return m.db.SaveDeposits([]types.Deposit{
		types.NewDeposit(msg.ProposalId, msg.Depositor, deposit.Amount, tx.Height),
	})
}

func (m *Module) getDeposit(tx *juno.Tx, msg *govtypes.MsgDeposit) (govtypes.Deposit, error) {
	clientCtx, err := m.getClientCtx()
	if err != nil {
		return govtypes.Deposit{}, err
	}

	height, err := m.getLatestHeight(tx.Height, msg.ProposalId)
	if err != nil {
		return govtypes.Deposit{}, err
	}

	proposal, err := m.source.Proposal(height, msg.ProposalId)
	if err != nil {
		return govtypes.Deposit{}, fmt.Errorf("error while getting proposal: %s", err)
	}

	propStatus := proposal.Status
	if !(propStatus == govtypes.StatusVotingPeriod || propStatus == govtypes.StatusDepositPeriod) {
		depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
		if err != nil {
			return govtypes.Deposit{}, fmt.Errorf("error while converting depositor to sdk.AccAddress: %s", err)
		}

		params := govtypes.NewQueryDepositParams(msg.ProposalId, depositor)
		resByTxQuery, err := gcutils.QueryDepositByTxQuery(clientCtx, params)
		if err != nil {
			return govtypes.Deposit{}, err
		}
		deposit := govtypes.Deposit{}
		clientCtx.Codec.MustUnmarshalJSON(resByTxQuery, &deposit)

		return deposit, nil
	}

	deposit, err := m.source.ProposalDeposit(height, msg.ProposalId, msg.Depositor)
	if err != nil {
		return govtypes.Deposit{}, fmt.Errorf("error while getting proposal deposit: %s", err)
	}

	return deposit, nil
}

// handleMsgVote allows to properly handle a handleMsgVote
func (m *Module) handleMsgVote(tx *juno.Tx, msg *govtypes.MsgVote) error {
	vote := types.NewVote(msg.ProposalId, msg.Voter, msg.Option, tx.Height)
	return m.db.SaveVote(vote)
}

func (m *Module) getLatestHeight(targetHeight int64, proposalID uint64) (int64, error) {
	_, err := m.source.Proposal(targetHeight, proposalID)
	if err == nil {
		return targetHeight, nil
	}
	if strings.Contains(err.Error(), "version mismatch on immutable IAVL tree") {
		// zero is equal to latest in terms of the client
		return 0, nil
	}

	return 0, err
}

func (m *Module) getClientCtx() (client.Context, error) {
	remoteNodeCfg, ok := m.cfg.Node.Details.(*remote.Details)
	if !ok {
		return client.Context{}, fmt.Errorf("failed to cast node config to *remote.Details, %+v", remoteNodeCfg)
	}
	rpcClient, err := client.NewClientFromNode(remoteNodeCfg.RPC.Address)
	if err != nil {
		return client.Context{}, err
	}

	encodingConfig := config.NewEncodingConfig(app.ModuleBasics)
	clientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithClient(rpcClient)

	return clientCtx, nil
}
