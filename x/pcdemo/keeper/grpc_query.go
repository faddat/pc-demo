package keeper

import (
	"github.com/fivesixnine/pc-demo/x/pcdemo/types"
)

var _ types.QueryServer = Keeper{}
