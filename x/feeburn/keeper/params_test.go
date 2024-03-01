package keeper_test

import (
	"testing"

	testkeeper "github.com/hyperspatial-labs/superstruct/testutil/keeper"

	"github.com/hyperspatial-labs/superstruct/x/feeburn/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FeeburnKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
