package keeper

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"

	"torram/x/rune/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type RuneMetadata struct {
	RuneID         string    `json:"rune_id"`
	State          string    `json:"state"`
	Chain          string    `json:"chain"`
	Owner          string    `json:"owner"`
	TxHash         string    `json:"tx_hash"`
	StakeTimestamp time.Time `json:"stake_timestamp"`
}

func (k msgServer) StakeRune(goCtx context.Context, msg *types.MsgStakeRune) (*types.MsgStakeRuneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create staked rune record with pending status
	stakedRune := types.StakedRune{
		RuneId:      msg.RuneId,
		Owner:       msg.Owner,
		Amount:      msg.Amount,
		StakingTime: uint64(ctx.BlockTime().Unix()),
		Status:      types.RuneStatusPending,
	}

	// Store the staked rune
	store := k.storeService.OpenKVStore(ctx)
	key := types.StakedRuneKey(msg.RuneId)
	bz, err := k.cdc.Marshal(&stakedRune)
	if err != nil {
		return nil, err
	}
	if err := store.Set(key, bz); err != nil {
		return nil, err
	}

	// Create Bitcoin Ordinal metadata
	metadata := RuneMetadata{
		RuneID:         msg.RuneId,
		State:          types.RuneStatusPending,
		Chain:          "Torram",
		Owner:          msg.Owner,
		TxHash:         hex.EncodeToString(ctx.TxBytes()),
		StakeTimestamp: ctx.BlockTime(),
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
			"rune_stake_initiated",
			sdk.NewAttribute("rune_id", msg.RuneId),
			sdk.NewAttribute("owner", msg.Owner),
			sdk.NewAttribute("amount", amount),
			sdk.NewAttribute("metadata", string(metadataJSON)),
		),
	)

	return &types.MsgStakeRuneResponse{}, nil
}
