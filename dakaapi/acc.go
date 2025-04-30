package dakaapi

import (
	"github.com/hootuu/gelato/errors"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/daka/daka"
	"github.com/nineora/nineora/daka/dakarpc"
	"github.com/nineora/nineora/nine/nineora"
)

func ProvinceGet(linkPtr nineora.Link) (*daka.Province, *errors.Error) {
	return ninecli.Rest[nineora.Link, daka.Province](dakarpc.ProvinceGetPath, &linkPtr)
}

func CityGet(linkPtr nineora.Link) (*daka.City, *errors.Error) {
	return ninecli.Rest[nineora.Link, daka.City](dakarpc.CityGetPath, &linkPtr)
}

func ZoneGet(linkPtr nineora.Link) (*daka.Zone, *errors.Error) {
	return ninecli.Rest[nineora.Link, daka.Zone](dakarpc.ZoneGetPath, &linkPtr)
}

func BranchGet(linkPtr nineora.Link) (*daka.Branch, *errors.Error) {
	return ninecli.Rest[nineora.Link, daka.Branch](dakarpc.BranchGetPath, &linkPtr)
}

func DivisionGet(linkPtr nineora.Link) (*daka.Division, *errors.Error) {
	return ninecli.Rest[nineora.Link, daka.Division](dakarpc.DivisionGetPath, &linkPtr)
}

func AssociateGet(linkPtr nineora.Link) (*daka.Associate, *errors.Error) {
	return ninecli.Rest[nineora.Link, daka.Associate](dakarpc.AssociateGetPath, &linkPtr)
}

func MerchantGet(linkPtr nineora.Link) (*daka.Merchant, *errors.Error) {
	return ninecli.Rest[nineora.Link, daka.Merchant](dakarpc.MerchantGetPath, &linkPtr)
}

func MemberGet(linkPtr nineora.Link) (*daka.Member, *errors.Error) {
	return ninecli.Rest[nineora.Link, daka.Member](dakarpc.MemberGetPath, &linkPtr)
}

func LottoGet() (*daka.LottoAccount, *errors.Error) {
	paras := ""
	return ninecli.Rest[string, daka.LottoAccount](dakarpc.LottoGetPath, &paras)
}
