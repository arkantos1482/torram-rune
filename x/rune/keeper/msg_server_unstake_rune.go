package keeper

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"torram/x/rune/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UnstakeRune(goCtx context.Context, msg *types.MsgUnstakeRune) (*types.MsgUnstakeRuneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

	// Verify ownership
	if stakedRune.Owner != msg.Owner {
		return nil, types.ErrUnauthorized
	}

	// Verify current status is staked
	if stakedRune.Status != types.RuneStatusStaked {
		return nil, types.ErrInvalidRuneStatus
	}

	// Update rune status to pending unstake
	stakedRune.Status = types.RuneStatusPending
	bz, err = k.cdc.Marshal(&stakedRune)
	if err != nil {
		return nil, err
	}
	if err := store.Set(key, bz); err != nil {
		return nil, err
	}

	// Create Bitcoin Ordinal metadata for unstaking
	metadata := RuneMetadata{
		RuneID:         msg.RuneId,
		State:          types.RuneStatusPending,
		Chain:          "Torram",
		Owner:          msg.Owner,
		TxHash:         "",          // Will be set by the Bitcoin network
		StakeTimestamp: time.Time{}, // Clear stake timestamp
	}

	// Convert metadata to JSON
	metadataJSON, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	// Emit event with metadata for Bitcoin Ordinal update
	amount := strconv.FormatUint(msg.Amount, 10)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"rune_unstake_initiated",
			sdk.NewAttribute("rune_id", msg.RuneId),
			sdk.NewAttribute("owner", msg.Owner),
			sdk.NewAttribute("amount", amount),
			sdk.NewAttribute("metadata", string(metadataJSON)),
		),
	)

	return &types.MsgUnstakeRuneResponse{}, nil
}
