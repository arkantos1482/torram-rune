package keeper

import (
	"context"

	"torram/x/rune/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetStakedRune(goCtx context.Context, req *types.QueryGetStakedRuneRequest) (*types.QueryGetStakedRuneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryGetStakedRuneResponse{}, nil
}
