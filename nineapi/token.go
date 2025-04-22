package nineapi

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/nine/nineora"
	"github.com/nineora/nineora/nine/ninerpc"
)

// TokenQuery : ninerpc.TokenQuery
func TokenQuery(req *ninerpc.TokenQueryReq) (*pagination.Pagination[nineora.Token], *errors.Error) {
	return ninecli.Rest[ninerpc.TokenQueryReq, pagination.Pagination[nineora.Token]](
		ninerpc.TokenQueryPath,
		req,
	)
}

// TokenQueryByNetwork : ninerpc.TokenQueryByNetwork
func TokenQueryByNetwork(req *ninerpc.TokenQueryByNetworkReq) (*pagination.Pagination[nineora.Token], *errors.Error) {
	return ninecli.Rest[ninerpc.TokenQueryByNetworkReq, pagination.Pagination[nineora.Token]](
		ninerpc.TokenQueryByNetworkPath,
		req,
	)
}

// TokenGet : ninerpc.TokenGet
func TokenGet(req *ninerpc.TokenGetReq) (*nineora.Token, *errors.Error) {
	return ninecli.Rest[ninerpc.TokenGetReq, nineora.Token](
		ninerpc.TokenGetPath,
		req,
	)
}
