package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/hyperspatial-labs/superstruct/testutil/keeper"

	"github.com/hyperspatial-labs/superstruct/x/feeburn/keeper"
	"github.com/hyperspatial-labs/superstruct/x/feeburn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FeeburnKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
