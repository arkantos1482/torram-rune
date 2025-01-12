package keeper

import (
	"context"

	"torram/x/rune/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateRuneStatus(goCtx context.Context, msg *types.MsgUpdateRuneStatus) (*types.MsgUpdateRuneStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateRuneStatusResponse{}, nil
}
