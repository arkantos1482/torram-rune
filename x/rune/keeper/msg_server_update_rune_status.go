package keeper

import (
	"context"
	"torram/x/rune/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateRuneStatus(goCtx context.Context, msg *types.MsgUpdateRuneStatus) (*types.MsgUpdateRuneStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Verify the caller is authorized (should be the bridge module account)
	if msg.Authority != k.GetAuthority() {
		return nil, types.ErrUnauthorized
	}

	// Get the staked rune
	store := k.storeService.OpenKVStore(ctx)
	key := types.StakedRuneKey(msg.RuneId)
	bz, err := store.Get(key)
	if err != nil {
		return nil, err
	}
	if bz == nil {
		return nil, types.ErrRuneNotStaked
	}

	var stakedRune types.StakedRune
	if err := k.cdc.Unmarshal(bz, &stakedRune); err != nil {
		return nil, err
	}

	// Verify the current status is pending
	if stakedRune.Status != types.RuneStatusPending {
		return nil, types.ErrInvalidRuneStatus
	}

	// Update rune status
	stakedRune.Status = msg.Status
	bz, err = k.cdc.Marshal(&stakedRune)
	if err != nil {
		return nil, err
	}
	if err := store.Set(key, bz); err != nil {
		return nil, err
	}

	return &types.MsgUpdateRuneStatusResponse{}, nil
}
