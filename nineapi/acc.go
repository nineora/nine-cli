package nineapi

// todo
//import (
//	"github.com/hootuu/gelato/errors"
//	"github.com/hootuu/gelato/io/pagination"
//	"github.com/nineora/nine-cli/ninecli"
//	"github.com/nineora/nineora/nine/nineora"
//	"github.com/nineora/nineora/nine/ninerpc"
//)
//
////type AccountQuery func(req *AccountQueryReq) (*pagination.Pagination[nineora.Account], *errors.Error)
////type AccountGet func(req *AccountGetReq) (*nineora.Account, *errors.Error)
////type BillQuery func(req *BillQueryReq) (*pagination.Pagination[nineora.Bill], *errors.Error)
//
//// AccQuery : ninerpc.AccQuery
//func AccQuery(
//	req *ninerpc.AccQueryReq,
//) (*pagination.Pagination[nineora.Account], *errors.Error) {
//	return ninecli.Rest[ninerpc.AccQueryReq, pagination.Pagination[nineora.Account]](
//		ninerpc.AccQueryPath,
//		req,
//	)
//}
//
//// BillQueryByBal : ninerpc.BillQueryByBal
//func BillQueryByBal(
//	req *ninerpc.BillQueryByBalReq,
//) (*pagination.Pagination[nineora.Bill], *errors.Error) {
//	return ninecli.Rest[ninerpc.BillQueryByBalReq, pagination.Pagination[nineora.Bill]](
//		ninerpc.BillQueryByBalPath,
//		req,
//	)
//}
