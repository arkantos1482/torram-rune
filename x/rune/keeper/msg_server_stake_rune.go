package keeper

import (
	"context"

	"torram/x/rune/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) StakeRune(goCtx context.Context, msg *types.MsgStakeRune) (*types.MsgStakeRuneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgStakeRuneResponse{}, nil
}
