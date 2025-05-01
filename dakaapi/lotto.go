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
	return ninecli.Rest[dakarpc.LottoMemberClaimReq, chain.Tx](dakarpc.LottoMemberClaimPath, req)
}
func LottoMerchantClaim(req *dakarpc.LottoMerchantClaimReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.LottoMerchantClaimReq, chain.Tx](dakarpc.LottoMerchantClaimPath, req)
}

func LottoAreaSet(req *dakarpc.LottoAreaSetReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.LottoAreaSetReq, chain.Tx](dakarpc.LottoAreaSetPath, req)
}

func LottoCommunitySet(req *dakarpc.LottoCommunitySetReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.LottoCommunitySetReq, chain.Tx](dakarpc.LottoCommunitySetPath, req)
}

func LottoSupplySet(req *dakarpc.LottoSupplySetReq) (*chain.Tx, *errors.Error) {
	return ninecli.Rest[dakarpc.LottoSupplySetReq, chain.Tx](dakarpc.LottoSupplySetPath, req)
}
