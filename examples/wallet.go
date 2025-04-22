package main

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/hexx"
	"github.com/hootuu/gelato/types/jsonx"
	"github.com/nineora/nine-cli/nineapi"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/nine/chain"
	"github.com/nineora/nineora/nine/ninekey"
	"github.com/nineora/nineora/nine/nineora"
	"github.com/nineora/nineora/nine/ninerpc"
	"github.com/spf13/cast"
	"time"
)

func main() {
	priStr := "ced49ab201255208746af52a4717e8851cfbc21ba4f686f5473abe0ff9047d01dabcc7e3a1276fbb562130f4036589ddb563717216c937681cc7cb0359934664"
	pri, err := hexx.Decode(priStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	ninecli.SetPriKey(pri)

	walletLink := nineora.LinkOf("WALLET", "a", "b", cast.ToString(time.Now().UnixMilli()))
	accLink := nineora.LinkOf("ACC", "a", "b", cast.ToString(time.Now().UnixMilli()))
	wallet, err := nineapi.WalletBind(&ninerpc.WalletBindReq{
		Link: walletLink,
		Seed: nineora.SeedOfCell("+8618688888888"),
		Meta: map[string]interface{}{
			"a": "A",
			"b": "B",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jsonx.JSON2String(wallet))
	fmt.Println(ninekey.CalcAddress(chain.Nineora, wallet.PublicKey))

	walletByGet, err := nineapi.WalletGet(&walletLink)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("walletGet: ", jsonx.MustJSON2String(walletByGet))

	account, err := nineapi.AccCreate(&ninerpc.AccCreateReq{
		Link:       accLink,
		WalletLink: walletLink,
		KeyType:    ninekey.Ed25519,
		NetworkID:  "xxxxxx",
		Password:   []byte("1234566"),
		Meta:       nil,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jsonx.JSON2String(account))
	addr, err := ninekey.CalcAddress(chain.SUI, account.PublicKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("addr: ", addr)

	balPage, err := nineapi.BalQueryByAcc(&ninerpc.BalQueryByAccReq{
		AccountLink: accLink,
		Page:        nil,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jsonx.JSON2String(balPage))

}
