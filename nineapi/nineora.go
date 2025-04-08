package nineapi

import (
	"github.com/hootuu/gelato/errors"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/nineora"
	"github.com/nineora/nineora/ninerpc"
)

// NineoraGet : ninerpc.NineoraGet
func NineoraGet(req *ninerpc.NineoraGetReq) (*nineora.Nineora, *errors.Error) {
	return ninecli.Rest[ninerpc.NineoraGetReq, nineora.Nineora](ninerpc.NineoraGetPath, req)
}
