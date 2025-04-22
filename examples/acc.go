package main

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/aesx"
	"github.com/hootuu/gelato/crtpto/hexx"
	"github.com/nineora/nine-cli/ninecli"
)

func main() {
	r, err := aesx.Encrypt([]byte("abc"), []byte("love"))
	if err != nil {
		fmt.Println(err)
		return
	}
	dr, err := aesx.Decrypt(r, []byte("love"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(dr))
	if true {
		return
	}
	priStr := "ced49ab201255208746af52a4717e8851cfbc21ba4f686f5473abe0ff9047d01dabcc7e3a1276fbb562130f4036589ddb563717216c937681cc7cb0359934664"
	pri, err := hexx.Decode(priStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	ninecli.SetPriKey(pri)
	//m, err := nineapi.NetworkQuery(&ninerpc.NetworkQueryReq{})
	//fmt.Println(serializer.JsonMustTo(m), err)

	//const networkID = "a2506828304e5dc0bc736014bad6fdaf"
	//network, err := nineapi.NetworkGet(&ninerpc.NetworkGetReq{NID: networkID})
	//fmt.Println(serializer.JsonMustTo(network), err)

	//nodePage, err := nineapi.AccountQuery(&ninerpc.AccountQueryReq{
	//	WithBalance: true,
	//	Page: &pagination.Page{
	//		Size: 10,
	//		Numb: 1,
	//	},
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(jsonx.JSON2String(nodePage))

}
