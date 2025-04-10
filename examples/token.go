package main

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/hexx"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/hootuu/gelato/io/serializer"
	"github.com/nineora/nine-cli/nineapi"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/ninerpc"
)

func main() {
	priStr := "ced49ab201255208746af52a4717e8851cfbc21ba4f686f5473abe0ff9047d01dabcc7e3a1276fbb562130f4036589ddb563717216c937681cc7cb0359934664"
	pri, err := hexx.Decode(priStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	ninecli.SetPriKey(pri)
	page, err := nineapi.TokenQuery(&ninerpc.TokenQueryReq{
		Page: pagination.PageNormal(),
	})
	fmt.Println(serializer.JsonMustTo(page), err)
}
