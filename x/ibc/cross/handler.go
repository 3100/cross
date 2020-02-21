package cross

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"
)

// NewHandler returns a handler
func NewHandler(keeper Keeper, contractHandler ContractHandler) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgInitiate:
			return handleMsgInitiate(ctx, keeper, msg)
		case MsgConfirm:
			return handleMsgConfirm(ctx, keeper, msg)
		case channeltypes.MsgPacket:
			switch data := msg.Data.(type) {
			case PacketDataPrepare:
				return handlePacketDataPrepare(ctx, keeper, contractHandler, msg, data)
			case PacketDataCommit:
				return handlePacketDataCommit(ctx, keeper, contractHandler, msg, data)
			default:
				return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized packet data type: %T", data)
			}
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized IBC message type: %T", msg)
		}
	}
}

/*
Steps:
- Ensure that all channels in ContractTransactions are correct
- Multicast a Prepare packet to each participants
*/
func handleMsgInitiate(ctx sdk.Context, k Keeper, msg MsgInitiate) (*sdk.Result, error) {
	err := k.MulticastPreparePacket(ctx, msg.Sender, msg, msg.ContractTransactions)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}

/*
Precondition:
- Given proof of packet is valid.
Steps:
- Try to apply given contract transaction to our state.
- If it was success, precommit these changes and get locks for concerned keys. Furthermore, send a Prepacket with status 'OK' to coordinator.
- If it was failed, discard theses changes. Furthermore, send a Prepacket with status 'Failed' to coordinator.
*/
func handlePacketDataPrepare(ctx sdk.Context, k Keeper, contractHandler ContractHandler, msg channeltypes.MsgPacket, data PacketDataPrepare) (*sdk.Result, error) {
	err := k.PrepareTransaction(ctx, contractHandler, msg.SourcePort, msg.SourceChannel, msg.DestinationPort, msg.DestinationChannel, data, msg.Signer)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}

/*
Precondition:
- All Given proof of packet are valid.
Steps:
- Verify all status of each PrepareResultPacket
- If any packet with status 'Failed' exist, send a CommitPacket with status 'Abort' to all participants.
- If all packet with status 'Ok' only exist, send a CommitPacket with status 'Commit' to all participants.
*/
func handleMsgConfirm(ctx sdk.Context, k Keeper, msg MsgConfirm) (*sdk.Result, error) {
	err := k.MulticastCommitPacket(ctx, msg.TxID, msg.PreparePackets, msg.Signer, msg.IsCommittable())
	if err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}

/*
Precondition:
- Given proof of packet is valid.
Steps:
- If PacketDataCommit indicates committable, commit precommitted state and unlock locked keys.
- If PacketDataCommit indicates not committable, rollback precommitted state and unlock locked keys.
*/
func handlePacketDataCommit(ctx sdk.Context, k Keeper, contractHandler ContractHandler, msg channeltypes.MsgPacket, data PacketDataCommit) (*sdk.Result, error) {
	err := k.ReceiveCommitPacket(ctx, contractHandler, msg.SourcePort, msg.SourceChannel, msg.DestinationPort, msg.DestinationChannel, data)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}
