package robot

import (
	"fmt"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/hootuu/gelato/types/jsonx"
	"github.com/nineora/nine-cli/dakaapi"
	"github.com/nineora/nine-cli/nineapi"
	"github.com/nineora/nineora/daka/dakarpc"
	"github.com/nineora/nineora/nine/nineora"
	"github.com/nineora/nineora/nine/ninerpc"
	"math/rand/v2"
	"time"
)

const NetworkID nineora.NetworkID = "9ce9bfd80ea246955b7f7997456dcd00"

const MaxMemberCount = 2000
const MaxMerchantCount = 500

const MaxProvinceCount = 20
const MaxCityCount = MaxProvinceCount * 20
const MaxZoneCount = MaxCityCount * 20

const MaxBranchCount = 50
const MaxDivisionCount = MaxBranchCount * 50
const MaxAssociateCount = MaxDivisionCount * 20

func Running() {
	for i := 0; i < 1; i++ {
		go OneScene()
	}
	time.Sleep(100 * time.Hour)
}

func OneScene() {
	member := RandMember()
	zone := RandZone()
	mc := RandMerchant()

	for i := 0; i < 1000; i++ {
		OneTime(member, zone, mc)
		time.Sleep(100 * time.Millisecond)
	}
}

func OneTime(member *Member, zone *AreaAgent, mc *Merchant) {
	orderAmount := uint64(rand.IntN(1000)) * 1000000
	contributionAmount := orderAmount * mc.InvestRatio / 10000
	tx, err := dakaapi.LottoTrigger(&dakarpc.LottoTriggerReq{
		OrderAmount:        orderAmount,
		ContributionAmount: contributionAmount,
		Member:             member.getNodeLink(),
		Merchant:           mc.getNodeLink(),
		Zone:               zone.getNodeLink(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\n[lotto.trigger] "+tx.Digest+"\n",
		"order.amount: ", orderAmount,
		"mc.InvestRatio: ", mc.InvestRatio,
		"contributionAmount: ", contributionAmount)

	func() {
		//time.Sleep(200 * time.Millisecond)
		//for i := 0; i < 1; i++ {
		//	memberLink := member.getNodeLink()
		//	dakaMember, err := dakaapi.MemberGet(&memberLink)
		//	if err != nil {
		//		fmt.Println(err)
		//		return
		//	}
		//	fmt.Println("\n[daka-member]\n", jsonx.MustJSON2String(dakaMember))
		//
		//	mcLink := member.getNodeLink()
		//	dakaMc, err := dakaapi.MerchantGet(&mcLink)
		//	if err != nil {
		//		fmt.Println(err)
		//		return
		//	}
		//	fmt.Println("\n[daka-mc]\n", jsonx.MustJSON2String(dakaMc))
		//
		//	time.Sleep(200 * time.Millisecond)
		//}
		tx, err = dakaapi.LottoMemberClaim(&dakarpc.LottoMemberClaimReq{Member: member.getNodeLink()})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("LottoMemberClaim.tx: ", tx)
		tx, err = dakaapi.LottoMerchantClaim(&dakarpc.LottoMerchantClaimReq{Merchant: mc.getNodeLink()})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("LottoMerchantClaim.tx: ", tx)
		//time.Sleep(200 * time.Millisecond)
		//for i := 0; i < 1; i++ {
		//	memberLink := member.getNodeLink()
		//	dakaMember, err := dakaapi.MemberGet(&memberLink)
		//	if err != nil {
		//		fmt.Println(err)
		//		return
		//	}
		//	fmt.Println("\n[daka-member]\n", jsonx.MustJSON2String(dakaMember))
		//
		//	mcLink := member.getNodeLink()
		//	dakaMc, err := dakaapi.MerchantGet(&mcLink)
		//	if err != nil {
		//		fmt.Println(err)
		//		return
		//	}
		//	fmt.Println("\n[daka-mc]\n", jsonx.MustJSON2String(dakaMc))
		//
		//	time.Sleep(200 * time.Millisecond)
		//}

		tx, err = dakaapi.AssociateWithdraw(&dakarpc.AssociateWithdrawReq{Associate: mc.Associate})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("AssociateWithdraw.tx: ", tx)
		tx, err = dakaapi.ZoneWithdraw(&dakarpc.ZoneWithdrawReq{Zone: zone.getNodeLink()})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ZoneWithdraw.tx: ", tx)
		//time.Sleep(200 * time.Millisecond)
		//for i := 0; i < 1; i++ {
		//	dakaMc, err := dakaapi.AssociateGet(&(mc.Associate))
		//	if err != nil {
		//		fmt.Println(err)
		//		return
		//	}
		//	fmt.Println("\n[daka-associate]\n", jsonx.MustJSON2String(dakaMc))
		//
		//	zoneLink := zone.getNodeLink()
		//	dakaZone, err := dakaapi.ZoneGet(&zoneLink)
		//	if err != nil {
		//		fmt.Println(err)
		//		return
		//	}
		//	fmt.Println("\n[daka-zone]\n", jsonx.MustJSON2String(dakaZone))
		//
		//	time.Sleep(200 * time.Millisecond)
		//}
		balancePage, err := nineapi.BalQueryByAcc(&ninerpc.BalQueryByAccReq{
			AccountLink: zone.getChainAccLink(),
			Page:        nil,
		})
		if err != nil {
			fmt.Println("nineapi.BalQueryByAcc err", err)
			return
		}
		fmt.Println("\n[daka-zone.bal]\n", jsonx.MustJSON2String(balancePage))
		for _, balItem := range balancePage.Data {
			ballPage, err := nineapi.BillQueryByBal(&ninerpc.BillQueryByBalReq{
				BalanceID: balItem.NID,
				Page:      pagination.PageNormal(),
			})
			if err != nil {
				fmt.Println("nineapi.BalQueryByAcc err", err)
				return
			}
			fmt.Println("[daka-zone.bal]balance:", balItem.GetBalance().Display())
			fmt.Println("====>>>>[daka-zone.bal]", balItem.NID, jsonx.MustJSON2String(ballPage))
		}
	}()

}
