package dakaapi

import (
	"github.com/hootuu/gelato/errors"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/daka/dakarpc"
	"github.com/nineora/nineora/nine/chain"
)

func LottoTrigger(req *dakarpc.LottoTriggerReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.LottoTriggerReq, chain.Tx](dakarpc.LottoTriggerPath, req)
}
func LottoMemberClaim(req *dakarpc.LottoMemberClaimReq) (*chain.Tx, *errors.Error) {
	return nil, nil
}
func LottoMerchantClaim(req *dakarpc.LottoMerchantClaimReq) (*chain.Tx, *errors.Error) {
	return nil, nil
}
