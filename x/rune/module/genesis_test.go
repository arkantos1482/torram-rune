package rune_test

import (
	"testing"

	keepertest "torram/testutil/keeper"
	"torram/testutil/nullify"
	rune "torram/x/rune/module"
	"torram/x/rune/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RuneKeeper(t)
	rune.InitGenesis(ctx, k, genesisState)
	got := rune.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
