package keeper

import (
	"context"

	"torram/x/rune/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UnstakeRune(goCtx context.Context, msg *types.MsgUnstakeRune) (*types.MsgUnstakeRuneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUnstakeRuneResponse{}, nil
}
