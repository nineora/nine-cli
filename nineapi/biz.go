package nineapi

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/nine/nineora"
	"github.com/nineora/nineora/nine/ninerpc"
)

// BusinessQuery : ninerpc.BusinessQuery
func BusinessQuery(
	req *ninerpc.BusinessQueryReq,
) (*pagination.Pagination[nineora.Business], *errors.Error) {
	return ninecli.Rest[ninerpc.BusinessQueryReq, pagination.Pagination[nineora.Business]](
		ninerpc.BusinessQueryPath,
		req,
	)
}

// BusinessQueryByNetwork : ninerpc.BusinessQueryByNetwork
func BusinessQueryByNetwork(
	req *ninerpc.BusinessQueryByNetworkReq,
) (*pagination.Pagination[nineora.Business], *errors.Error) {
	return ninecli.Rest[ninerpc.BusinessQueryByNetworkReq, pagination.Pagination[nineora.Business]](
		ninerpc.BusinessQueryByNetworkPath,
		req,
	)
}

// BusinessGet : ninerpc.BusinessGet
func BusinessGet(
	req *ninerpc.BusinessGetReq,
) (*nineora.Business, *errors.Error) {
	return ninecli.Rest[ninerpc.BusinessGetReq, nineora.Business](
		ninerpc.BusinessGetPath,
		req,
	)
}
