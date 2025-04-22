package nineapi

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/nine/nineora"
	"github.com/nineora/nineora/nine/ninerpc"
)

// NodeQueryByNetwork : ninerpc.NodeQueryByNetwork
func NodeQueryByNetwork(req *ninerpc.NodeQueryByNetworkReq) (*pagination.Pagination[nineora.Node], *errors.Error) {
	return ninecli.Rest[ninerpc.NodeQueryByNetworkReq, pagination.Pagination[nineora.Node]](
		ninerpc.NodeQueryByNetworkPath,
		req,
	)
}

// NodeQueryBySuperior : ninerpc.NodeQueryBySuperior
func NodeQueryBySuperior(req *ninerpc.NodeQueryBySuperiorReq) (*pagination.Pagination[nineora.Node], *errors.Error) {
	return ninecli.Rest[ninerpc.NodeQueryBySuperiorReq, pagination.Pagination[nineora.Node]](
		ninerpc.NodeQueryBySuperiorPath,
		req,
	)
}

// NodeGet : ninerpc.NodeGet
func NodeGet(req *ninerpc.NodeGetReq) (*nineora.Node, *errors.Error) {
	return ninecli.Rest[ninerpc.NodeGetReq, nineora.Node](
		ninerpc.NodeGetPath,
		req,
	)
}
