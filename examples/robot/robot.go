package robot

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/hexx"
	"github.com/hootuu/gelato/types/jsonx"
	"github.com/nineora/nine-cli/dakaapi"
	"github.com/nineora/nine-cli/nineapi"
	"github.com/nineora/nineora/daka/dakarpc"
	"github.com/nineora/nineora/nine/chain"
	"github.com/nineora/nineora/nine/ninekey"
	"github.com/nineora/nineora/nine/nineora"
	"github.com/nineora/nineora/nine/ninerpc"
	"time"
)

const NetworkID nineora.NetworkID = "7304da03e436ea87b0426440fb01b272"

const MaxMemberCount = 2000
const MaxMerchantCount = 500

const MaxProvinceCount = 20
const MaxCityCount = MaxProvinceCount * 20
const MaxZoneCount = MaxCityCount * 20

const MaxBranchCount = 50
const MaxDivisionCount = MaxBranchCount * 50
const MaxAssociateCount = MaxDivisionCount * 20

func Running() {
	//tx, err := dakaapi.LottoAreaSet(&dakarpc.LottoAreaSetReq{
	//	ProvinceRatio: 100,
	//	CityRatio:     200,
	//	ZoneRatio:     200,
	//})
	//fmt.Println(tx, err)
	//tx, err = dakaapi.LottoCommunitySet(&dakarpc.LottoCommunitySetReq{
	//	BranchRatio:    500,
	//	DivisionRatio:  100,
	//	AssociateRatio: 800,
	//})
	//fmt.Println(tx, err)
	//tx, err = dakaapi.LottoSupplySet(&dakarpc.LottoSupplySetReq{
	//	ShiftRatio:     10000,
	//	LockedRatio:    8000,
	//	NxtUnlockRatio: 1000,
	//})
	//fmt.Println(tx, err)
	for i := 0; i < 1; i++ {
		go OneScene()
	}
	//NewMerchant()
	time.Sleep(100 * time.Hour)
}

func OneScene() {
	member := RandMember()
	zone := RandZone()
	mc := RandMerchant()

	for i := 0; i < 1; i++ {
		OneTime(member, zone, mc)
		//time.Sleep(10000 * time.Millisecond)
	}
}

func OneTime(member *Member, zone *AreaAgent, mc *Merchant) {
	mnemonicShow := func() {
		asso := AssociateGet(mc.Associate)
		m, err := nineapi.AccMnemonic(&ninerpc.AccMnemonicReq{
			Link:     asso.getChainAccLink(),
			Password: []byte("12345678"),
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		addr, _ := ninekey.CalcAddress(chain.SUI, m.PublicKey)
		fmt.Println("addr: ", addr)
		fmt.Println("mno: ", m.Mnemonic)
		fmt.Println("pri: ", hexx.Encode(m.PrivateKey))
	}
	mnemonicShow()

	orderAmount := uint64(10000) //uint64(rand.IntN(1000)) * 100
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
	//
	time.Sleep(10000 * time.Millisecond)

	infoShow := func() {
		lottoInfo, err := dakaapi.LottoGet()
		if err != nil {
			fmt.Println("lottoGet err", err)
			return
		}

		fmt.Println("\n\n[lotto.info] \n", jsonx.MustJSON2String(lottoInfo))

		mcData, err := dakaapi.MerchantGet(mc.getNodeLink())
		if err != nil {
			fmt.Println("dakaapi.MerchantGet err", err)
			return
		}

		fmt.Println("\n[daka-mc.data]\n", mc.getNodeLink(),
			"\n",
			jsonx.MustJSON2String(mcData))

		memberData, err := dakaapi.MemberGet(member.getNodeLink())
		if err != nil {
			fmt.Println("dakaapi.MemberGet err", err)
			return
		}
		fmt.Println("\n[daka-member.data]\n", member.getNodeLink(),
			"\n",
			jsonx.MustJSON2String(memberData))

		mcSrcLink := GetSuperior(member.getNodeLink())
		mcSrcData, err := dakaapi.MerchantGet(mcSrcLink)
		if err != nil {
			fmt.Println("dakaapi.MemberGet.SRC err", err)
			return
		}
		fmt.Println("\n[daka-mc-src.data]\n", mcSrcLink,
			"\n",
			jsonx.MustJSON2String(mcSrcData))

		zoneData, err := dakaapi.ZoneGet(zone.getNodeLink())
		if err != nil {
			fmt.Println("dakaapi.ZoneGet err", err)
			return
		}

		fmt.Println("\n[daka-zone.data]\n", zone.getNodeLink(),
			"\n",
			jsonx.MustJSON2String(zoneData))

		cityLink := GetSuperior(zone.getNodeLink())
		cityData, err := dakaapi.CityGet(cityLink)
		if err != nil {
			fmt.Println("dakaapi.CityGet err", err)
			return
		}

		fmt.Println("\n[daka-city.data]\n", cityLink,
			"\n",
			jsonx.MustJSON2String(cityData))

		provinceLink := GetSuperior(cityLink)
		provinceData, err := dakaapi.ProvinceGet(provinceLink)
		if err != nil {
			fmt.Println("dakaapi.ProvinceGet err", err)
			return
		}

		fmt.Println("\n[daka-province.data]\n", provinceLink,
			"\n",
			jsonx.MustJSON2String(provinceData))

		assoData, err := dakaapi.AssociateGet(mc.Associate)
		if err != nil {
			fmt.Println("dakaapi.ZoneGet err", err)
			return
		}
		fmt.Println("\n[daka-asso.data]\n", mc.Associate,
			"\n",
			jsonx.MustJSON2String(assoData))

		divisionLink := GetSuperior(mc.Associate)
		divData, err := dakaapi.DivisionGet(divisionLink)
		if err != nil {
			fmt.Println("dakaapi.DivisionGet err", err)
			return
		}
		fmt.Println("\n[daka-div.data]\n", divisionLink,
			"\n",
			jsonx.MustJSON2String(divData))

		branchLink := GetSuperior(divisionLink)
		branchData, err := dakaapi.BranchGet(branchLink)
		if err != nil {
			fmt.Println("dakaapi.BranchGet err", err)
			return
		}
		fmt.Println("\n[daka-branch.data]\n", branchLink,
			"\n",
			jsonx.MustJSON2String(branchData))
	}

	infoShow()
	//if true {
	//	os.Exit(1)
	//}

	//func() {
	//	//time.Sleep(200 * time.Millisecond)
	//	//for i := 0; i < 1; i++ {
	//	//	memberLink := member.getNodeLink()
	//	//	dakaMember, err := dakaapi.MemberGet(&memberLink)
	//	//	if err != nil {
	//	//		fmt.Println(err)
	//	//		return
	//	//	}
	//	//	fmt.Println("\n[daka-member]\n", jsonx.MustJSON2String(dakaMember))
	//	//
	//	//	mcLink := member.getNodeLink()
	//	//	dakaMc, err := dakaapi.MerchantGet(&mcLink)
	//	//	if err != nil {
	//	//		fmt.Println(err)
	//	//		return
	//	//	}
	//	//	fmt.Println("\n[daka-mc]\n", jsonx.MustJSON2String(dakaMc))
	//	//
	//	//	time.Sleep(200 * time.Millisecond)
	//	//}
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
	tx, err = dakaapi.AssociateWithdraw(&dakarpc.AssociateWithdrawReq{Associate: mc.Associate})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("AssociateWithdraw.tx: ", tx)

	time.Sleep(10 * time.Second)
	infoShow()
	//	//time.Sleep(200 * time.Millisecond)
	//	//for i := 0; i < 1; i++ {
	//	//	memberLink := member.getNodeLink()
	//	//	dakaMember, err := dakaapi.MemberGet(&memberLink)
	//	//	if err != nil {
	//	//		fmt.Println(err)
	//	//		return
	//	//	}
	//	//	fmt.Println("\n[daka-member]\n", jsonx.MustJSON2String(dakaMember))
	//	//
	//	//	mcLink := member.getNodeLink()
	//	//	dakaMc, err := dakaapi.MerchantGet(&mcLink)
	//	//	if err != nil {
	//	//		fmt.Println(err)
	//	//		return
	//	//	}
	//	//	fmt.Println("\n[daka-mc]\n", jsonx.MustJSON2String(dakaMc))
	//	//
	//	//	time.Sleep(200 * time.Millisecond)
	//	//}
	//

	//	//tx, err = dakaapi.ZoneWithdraw(&dakarpc.ZoneWithdrawReq{Zone: zone.getNodeLink()})
	//	//if err != nil {
	//	//	fmt.Println(err)
	//	//	return
	//	//}
	//	//fmt.Println("ZoneWithdraw.tx: ", tx)
	//	//time.Sleep(200 * time.Millisecond)
	//	//for i := 0; i < 1; i++ {
	//	//	dakaMc, err := dakaapi.AssociateGet(&(mc.Associate))
	//	//	if err != nil {
	//	//		fmt.Println(err)
	//	//		return
	//	//	}
	//	//	fmt.Println("\n[daka-associate]\n", jsonx.MustJSON2String(dakaMc))
	//	//
	//	//	zoneLink := zone.getNodeLink()
	//	//	dakaZone, err := dakaapi.ZoneGet(&zoneLink)
	//	//	if err != nil {
	//	//		fmt.Println(err)
	//	//		return
	//	//	}
	//	//	fmt.Println("\n[daka-zone]\n", jsonx.MustJSON2String(dakaZone))
	//	//
	//	time.Sleep(5000 * time.Millisecond)
	//	//}
	//	for i := 0; i < 1; i++ {
	//
	//		mcData, err := dakaapi.MerchantGet(mc.getNodeLink())
	//		if err != nil {
	//			fmt.Println("dakaapi.MerchantGet err", err)
	//			return
	//		}
	//
	//		fmt.Println("\n[daka-mc.data]\n", mc.getNodeLink(),
	//			"\n",
	//			jsonx.MustJSON2String(mcData))
	//		break
	//	}
	//
	//	for i := 0; i < 5; i++ {
	//		asso := AssociateGet(mc.Associate)
	//		balancePage, err := nineapi.BalQueryByAcc(&ninerpc.BalQueryByAccReq{
	//			AccountLink: asso.getChainAccLink(),
	//			Page:        nil,
	//		})
	//		if err != nil {
	//			fmt.Println("nineapi.BalQueryByAcc err", err)
	//			return
	//		}
	//		if len(balancePage.Data) == 0 {
	//			time.Sleep(1000 * time.Millisecond)
	//			continue
	//		}
	//		fmt.Println("\n[daka-asso.bal]\n", jsonx.MustJSON2String(balancePage))
	//		for _, balItem := range balancePage.Data {
	//			ballPage, err := nineapi.BillQueryByBal(&ninerpc.BillQueryByBalReq{
	//				BalanceID: balItem.NID,
	//				Page:      pagination.PageNormal(),
	//			})
	//			if err != nil {
	//				fmt.Println("nineapi.BalQueryByAcc err", err)
	//				return
	//			}
	//			fmt.Println("[daka-asso.bal]balance:", balItem.GetBalance().Display())
	//			fmt.Println("====>>>>[daka-asso.bal]", balItem.NID, jsonx.MustJSON2String(ballPage))
	//		}
	//		break
	//	}
	//
	//	for i := 0; i < 5; i++ {
	//		balancePage, err := nineapi.BalQueryByAcc(&ninerpc.BalQueryByAccReq{
	//			AccountLink: zone.getChainAccLink(),
	//			Page:        nil,
	//		})
	//		if err != nil {
	//			fmt.Println("nineapi.BalQueryByAcc err", err)
	//			return
	//		}
	//		if len(balancePage.Data) == 0 {
	//			time.Sleep(1000 * time.Millisecond)
	//			continue
	//		}
	//		fmt.Println("\n[daka-zone.bal]\n", jsonx.MustJSON2String(balancePage))
	//		for _, balItem := range balancePage.Data {
	//			ballPage, err := nineapi.BillQueryByBal(&ninerpc.BillQueryByBalReq{
	//				BalanceID: balItem.NID,
	//				Page:      pagination.PageNormal(),
	//			})
	//			if err != nil {
	//				fmt.Println("nineapi.BalQueryByAcc err", err)
	//				return
	//			}
	//			fmt.Println("[daka-zone.bal]balance:", balItem.GetBalance().Display())
	//			fmt.Println("====>>>>[daka-zone.bal]", balItem.NID, jsonx.MustJSON2String(ballPage))
	//		}
	//		break
	//	}
	//	fmt.Println("=====>>>on-time=====>>>>ovner")
	//}()

}
