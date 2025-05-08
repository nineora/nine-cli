package dakaapi

import (
	"github.com/hootuu/gelato/errors"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/daka/dakarpc"
	"github.com/nineora/nineora/nine/chain"
)

func ProvinceCreate(req *dakarpc.ProvinceCreateReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.ProvinceCreateReq, chain.Tx](dakarpc.ProvinceCreatePath, req)
}

func CityCreate(req *dakarpc.CityCreateReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.CityCreateReq, chain.Tx](dakarpc.CityCreatePath, req)
}

func ZoneCreate(req *dakarpc.ZoneCreateReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.ZoneCreateReq, chain.Tx](dakarpc.ZoneCreatePath, req)
}

func BranchCreate(req *dakarpc.BranchCreateReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.BranchCreateReq, chain.Tx](dakarpc.BranchCreatePath, req)
}

func DivisionCreate(req *dakarpc.DivisionCreateReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.DivisionCreateReq, chain.Tx](dakarpc.DivisionCreatePath, req)
}

func AssociateCreate(req *dakarpc.AssociateCreateReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.AssociateCreateReq, chain.Tx](dakarpc.AssociateCreatePath, req)
}

func MerchantCreate(req *dakarpc.MerchantCreateReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.MerchantCreateReq, chain.Tx](dakarpc.MerchantCreatePath, req)
}

func MemberCreate(req *dakarpc.MemberCreateReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.MemberCreateReq, chain.Tx](dakarpc.MemberCreatePath, req)
}

func MerchantSetInvest(req *dakarpc.MerchantSetInvestReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.MerchantSetInvestReq, chain.Tx](dakarpc.MerchantSetInvestPath, req)
}
