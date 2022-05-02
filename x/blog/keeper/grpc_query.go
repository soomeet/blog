package keeper

import (
	"github.com/soomeet/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
