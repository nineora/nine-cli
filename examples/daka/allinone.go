package main

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/hexx"
	"github.com/hootuu/gelato/types/jsonx"
	"github.com/nineora/nine-cli/dakaapi"
	"github.com/nineora/nine-cli/nineapi"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/daka/dakarpc"
	"github.com/nineora/nineora/nine/chain"
	"github.com/nineora/nineora/nine/ninekey"
	"github.com/nineora/nineora/nine/nineora"
	"github.com/nineora/nineora/nine/ninerpc"
	"github.com/spf13/cast"
	"time"
)

const DakaNetworkAddr chain.Address = "0xabb2d16f0f52c6feb1c5f4f182886133b3f27ebd60b52b02443a8bd48a6d475f"
const DakaUnimktAddr chain.Address = "0xc8266b056c88b09135b8ca5047647f27ef542987bdc791fc9e49d46e2158f275"
const DakaUniPwd = "88888888"

var DakaNetwork = nineora.NewNetworkID(chain.SUI, DakaNetworkAddr)
var DataUnimkt = nineora.NewBusinessID(chain.SUI, DakaUnimktAddr)

func newWallet(code string, id string) nineora.Link {
	walletLink := nineora.LinkOf(DakaNetwork, code, id)
	_, err := nineapi.WalletBind(&ninerpc.WalletBindReq{
		Link: walletLink,
		Seed: nineora.SeedOfCell("+8618688888888"),
		Meta: map[string]interface{}{
			"a": "A",
			"b": "B",
		},
	})
	if err != nil {
		fmt.Println("newWallet::", err)
		return ""
	}
	return walletLink
}

func newAccount(walletLink nineora.Link, code string, id string) nineora.Link {
	accLink := nineora.LinkOf(DakaNetwork, code, id)
	_, err := nineapi.AccCreate(&ninerpc.AccCreateReq{
		Link:       accLink,
		WalletLink: walletLink,
		KeyType:    ninekey.Ed25519,
		NetworkID:  DakaNetwork,
		Password:   []byte(DakaUniPwd),
		Meta:       nil,
	})
	if err != nil {
		fmt.Println("newAccount::", err)
		return ""
	}
	return accLink
}

func main() {
	priStr := "ced49ab201255208746af52a4717e8851cfbc21ba4f686f5473abe0ff9047d01dabcc7e3a1276fbb562130f4036589ddb563717216c937681cc7cb0359934664"
	pri, err := hexx.Decode(priStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	ninecli.SetPriKey(pri)

	provinceWallet := newWallet("province.wallet", cast.ToString(time.Now().UnixMilli()))
	provinceAcc := newAccount(provinceWallet, "province.acc", cast.ToString(time.Now().UnixMilli()))
	provinceNodeLink := nineora.LinkOf(DakaNetwork, "province.node", cast.ToString(time.Now().UnixMilli()))
	tx, err := dakaapi.ProvinceCreate(&dakarpc.ProvinceCreateReq{
		Link:    provinceNodeLink,
		Benefit: provinceAcc,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("province: ", jsonx.MustJSON2String(tx))

	cityWallet := newWallet("city.wallet", cast.ToString(time.Now().UnixMilli()))
	cityAcc := newAccount(cityWallet, "city.acc", cast.ToString(time.Now().UnixMilli()))
	cityLink := nineora.LinkOf(DakaNetwork, "city.node", cast.ToString(time.Now().UnixMilli()))
	tx, err = dakaapi.CityCreate(&dakarpc.CityCreateReq{
		Link:     cityLink,
		Province: provinceNodeLink,
		Benefit:  cityAcc,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("city: ", jsonx.MustJSON2String(tx))

	zoneWallet := newWallet("zone.wallet", cast.ToString(time.Now().UnixMilli()))
	zoneAcc := newAccount(zoneWallet, "zone.acc", cast.ToString(time.Now().UnixMilli()))
	zoneLink := nineora.LinkOf(DakaNetwork, "zone.node", cast.ToString(time.Now().UnixMilli()))
	tx, err = dakaapi.ZoneCreate(&dakarpc.ZoneCreateReq{
		Link:    zoneLink,
		City:    cityLink,
		Benefit: zoneAcc,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("zone: ", jsonx.MustJSON2String(tx))

	branchWallet := newWallet("branch.wallet", cast.ToString(time.Now().UnixMilli()))
	branchAcc := newAccount(branchWallet, "branch.acc", cast.ToString(time.Now().UnixMilli()))
	branchLink := nineora.LinkOf(DakaNetwork, "branch.node", cast.ToString(time.Now().UnixMilli()))
	tx, err = dakaapi.BranchCreate(&dakarpc.BranchCreateReq{
		Link:    branchLink,
		Benefit: branchAcc,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("branch: ", jsonx.MustJSON2String(tx))

	divWallet := newWallet("div.wallet", cast.ToString(time.Now().UnixMilli()))
	divAcc := newAccount(divWallet, "div.acc", cast.ToString(time.Now().UnixMilli()))
	divLink := nineora.LinkOf(DakaNetwork, "div.node", cast.ToString(time.Now().UnixMilli()))
	tx, err = dakaapi.DivisionCreate(&dakarpc.DivisionCreateReq{
		Link:    divLink,
		Branch:  branchLink,
		Benefit: divAcc,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("div: ", jsonx.MustJSON2String(tx))

	assoWallet := newWallet("asso.wallet", cast.ToString(time.Now().UnixMilli()))
	assoDiv := newAccount(assoWallet, "asso.acc", cast.ToString(time.Now().UnixMilli()))
	assoLink := nineora.LinkOf(DakaNetwork, "asso.node", cast.ToString(time.Now().UnixMilli()))
	tx, err = dakaapi.AssociateCreate(&dakarpc.AssociateCreateReq{
		Link:     assoLink,
		Division: divLink,
		Benefit:  assoDiv,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("asso: ", jsonx.MustJSON2String(tx))

	mcWallet := newWallet("mc.wallet", cast.ToString(time.Now().UnixMilli()))
	mcAcc := newAccount(mcWallet, "mc.acc", cast.ToString(time.Now().UnixMilli()))
	mcLink := nineora.LinkOf(DakaNetwork, "mc.node", cast.ToString(time.Now().UnixMilli()))
	tx, err = dakaapi.MerchantCreate(&dakarpc.MerchantCreateReq{
		Link:        mcLink,
		Associate:   assoLink,
		Benefit:     mcAcc,
		InvestRatio: 4000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("mc: ", jsonx.MustJSON2String(tx))

	mcSrcWallet := newWallet("mc.wallet", cast.ToString(time.Now().UnixMilli()))
	mcSrcAcc := newAccount(mcSrcWallet, "mc.acc", cast.ToString(time.Now().UnixMilli()))
	mcSrcLink := nineora.LinkOf(DakaNetwork, "mc.node", cast.ToString(time.Now().UnixMilli()))
	tx, err = dakaapi.MerchantCreate(&dakarpc.MerchantCreateReq{
		Link:        mcSrcLink,
		Associate:   assoLink,
		Benefit:     mcSrcAcc,
		InvestRatio: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("mc.src: ", jsonx.MustJSON2String(tx))

	memberWallet := newWallet("member.wallet", cast.ToString(time.Now().UnixMilli()))
	memberAcc := newAccount(memberWallet, "member.acc", cast.ToString(time.Now().UnixMilli()))
	memberLink := nineora.LinkOf(DakaNetwork, "member.node", cast.ToString(time.Now().UnixMilli()))
	tx, err = dakaapi.MemberCreate(&dakarpc.MemberCreateReq{
		Link:     memberLink,
		Merchant: mcSrcLink,
		Benefit:  memberAcc,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("member: ", jsonx.MustJSON2String(tx))

	for i := 0; i < 3; i++ {
		tx, err = dakaapi.LottoTrigger(&dakarpc.LottoTriggerReq{
			OrderAmount:        1000 * 1000000,
			ContributionAmount: 400 * 1000000,
			Member:             memberLink,
			Merchant:           mcLink,
			Zone:               zoneLink,
		})

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("lotto trigger: ", jsonx.MustJSON2String(tx))
		time.Sleep(100 * time.Millisecond)
	}

	for i := 0; i < 100; i++ {
		provinceData, err := dakaapi.ProvinceGet(&provinceNodeLink)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("province data: ", jsonx.MustJSON2String(provinceData))

		cityData, err := dakaapi.CityGet(&cityLink)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("city data: ", jsonx.MustJSON2String(cityData))

		zoneData, err := dakaapi.ZoneGet(&zoneLink)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("zone data: ", jsonx.MustJSON2String(zoneData))

		branchData, err := dakaapi.BranchGet(&branchLink)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("branch data: ", jsonx.MustJSON2String(branchData))

		divData, err := dakaapi.DivisionGet(&divLink)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("div data: ", jsonx.MustJSON2String(divData))

		asso, err := dakaapi.AssociateGet(&assoLink)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("asso data: ", jsonx.MustJSON2String(asso))

		mcSrcData, err := dakaapi.MerchantGet(&mcSrcLink)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("mc.src data: ", jsonx.MustJSON2String(mcSrcData))

		mcData, err := dakaapi.MerchantGet(&mcLink)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("merchant data: ", jsonx.MustJSON2String(mcData))

		memberData, err := dakaapi.MemberGet(&memberLink)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("member data: ", jsonx.MustJSON2String(memberData))

		memberAccInfo, err := nineapi.AccGet(&memberAcc)
		if err != nil {
			fmt.Println(err)
			return
		}
		memberAddr, _ := ninekey.CalcAddress(chain.SUI, memberAccInfo.PublicKey)
		fmt.Println("[member.sui] ", memberAddr)
		pageData, err := nineapi.BalQueryByAcc(&ninerpc.BalQueryByAccReq{
			AccountLink: memberAcc,
			Page:        nil,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("member.BalQueryByAcc---------->>>>>>>>>")
		fmt.Println(jsonx.MustJSON2String(pageData))
		fmt.Println("member.BalQueryByAcc<<<<<<<<<<---------")
		time.Sleep(10 * time.Second)
		fmt.Println("---->:::::---->")
	}

}
