package feeburn_test

import (
	"testing"

	keepertest "github.com/hyperspatial-labs/superstruct/testutil/keeper"
	"github.com/hyperspatial-labs/superstruct/testutil/nullify"

	"github.com/hyperspatial-labs/superstruct/x/feeburn"
	"github.com/hyperspatial-labs/superstruct/x/feeburn/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.FeeburnKeeper(t)
	feeburn.InitGenesis(ctx, *k, genesisState)
	got := feeburn.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
