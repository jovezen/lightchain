package keeper

import (
	"github.com/jovezen/lightchain/x/lightchain/types"
)

var _ types.QueryServer = Keeper{}
