package dakaapi

import (
	"github.com/hootuu/gelato/errors"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/daka/dakarpc"
	"github.com/nineora/nineora/nine/chain"
)

func MemberWithdraw(req *dakarpc.MemberWithdrawReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.MemberWithdrawReq, chain.Tx](dakarpc.MemberWithdrawPath, req)
}

func MerchantWithdraw(req *dakarpc.MerchantWithdrawReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.MerchantWithdrawReq, chain.Tx](dakarpc.MerchantWithdrawPath, req)
}

func ProvinceWithdraw(req *dakarpc.ProvinceWithdrawReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.ProvinceWithdrawReq, chain.Tx](dakarpc.ProvinceWithdrawPath, req)
}

func CityWithdraw(req *dakarpc.CityWithdrawReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.CityWithdrawReq, chain.Tx](dakarpc.CityWithdrawPath, req)
}

func ZoneWithdraw(req *dakarpc.ZoneWithdrawReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.ZoneWithdrawReq, chain.Tx](dakarpc.ZoneWithdrawPath, req)
}

func BranchWithdraw(req *dakarpc.BranchWithdrawReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.BranchWithdrawReq, chain.Tx](dakarpc.BranchWithdrawPath, req)
}

func DivisionWithdraw(req *dakarpc.DivisionWithdrawReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.DivisionWithdrawReq, chain.Tx](dakarpc.DivisionWithdrawPath, req)
}

func AssociateWithdraw(req *dakarpc.AssociateWithdrawReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.AssociateWithdrawReq, chain.Tx](dakarpc.AssociateWithdrawPath, req)
}
