package nineapi

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/nine/nineora"
	"github.com/nineora/nineora/nine/ninerpc"
)

// NetworkQuery : ninerpc.NetworkQuery
func NetworkQuery(req *ninerpc.NetworkQueryReq) (*pagination.Pagination[nineora.Network], *errors.Error) {
	return ninecli.Rest[ninerpc.NetworkQueryReq, pagination.Pagination[nineora.Network]](ninerpc.NetworkQueryPath, req)
}

// NetworkGet : ninerpc.NetworkGet
func NetworkGet(req *ninerpc.NetworkGetReq) (*nineora.Network, *errors.Error) {
	return ninecli.Rest[ninerpc.NetworkGetReq, nineora.Network](ninerpc.NetworkGetPath, req)
}
