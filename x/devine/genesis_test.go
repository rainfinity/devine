package devine_test

import (
	"testing"

	keepertest "github.com/rainfinity/devine/testutil/keeper"
	"github.com/rainfinity/devine/testutil/nullify"
	"github.com/rainfinity/devine/x/devine"
	"github.com/rainfinity/devine/x/devine/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DevineKeeper(t)
	devine.InitGenesis(ctx, *k, genesisState)
	got := devine.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
