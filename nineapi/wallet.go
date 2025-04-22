package nineapi

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/nine/nineora"
	"github.com/nineora/nineora/nine/ninerpc"
)

// WalletBind : ninerpc.WalletBind
func WalletBind(
	req *ninerpc.WalletBindReq,
) (*nineora.Wallet, *errors.Error) {
	return ninecli.Rest[ninerpc.WalletBindReq, nineora.Wallet](ninerpc.WalletBindPath, req)

}

// AccCreate : ninerpc.AccCreate
func AccCreate(
	req *ninerpc.AccCreateReq,
) (*nineora.Account, *errors.Error) {
	return ninecli.Rest[ninerpc.AccCreateReq, nineora.Account](ninerpc.AccCreatePath, req)

}

// AccQuery : ninerpc.AccQuery
func AccQuery(
	req *ninerpc.AccQueryReq,
) (*pagination.Pagination[nineora.Account], *errors.Error) {
	return ninecli.Rest[ninerpc.AccQueryReq, pagination.Pagination[nineora.Account]](
		ninerpc.AccQueryPath,
		req,
	)
}

// BalQueryByAcc : ninerpc.BalQueryByAcc
func BalQueryByAcc(
	req *ninerpc.BalQueryByAccReq,
) (*pagination.Pagination[nineora.Balance], *errors.Error) {
	return ninecli.Rest[ninerpc.BalQueryByAccReq, pagination.Pagination[nineora.Balance]](
		ninerpc.BalQueryByAccPath, req,
	)
}

// BillQueryByBal : ninerpc.BillQueryByBal
func BillQueryByBal(
	req *ninerpc.BillQueryByBalReq,
) (*pagination.Pagination[nineora.Bill], *errors.Error) {
	return ninecli.Rest[ninerpc.BillQueryByBalReq, pagination.Pagination[nineora.Bill]](
		ninerpc.BillQueryByBalPath,
		req,
	)
}

func WalletGet(link *nineora.Link) (*nineora.Wallet, *errors.Error) {
	return ninecli.Rest[nineora.Link, nineora.Wallet](
		ninerpc.WalletGetPath,
		link,
	)
}

func AccGet(link *nineora.Link) (*nineora.Account, *errors.Error) {
	return ninecli.Rest[nineora.Link, nineora.Account](
		ninerpc.AccGetPath,
		link,
	)
}
