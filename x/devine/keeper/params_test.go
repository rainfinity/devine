package keeper_test

import (
	"testing"

	testkeeper "github.com/rainfinity/devine/testutil/keeper"
	"github.com/rainfinity/devine/x/devine/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DevineKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
