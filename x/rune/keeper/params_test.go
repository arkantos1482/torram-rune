package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "torram/testutil/keeper"
	"torram/x/rune/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.RuneKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
