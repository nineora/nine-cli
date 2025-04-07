package main

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/hexx"
	"github.com/hootuu/gelato/io/rest"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/nineora"
	"github.com/nineora/nineora/ninerpc"
)

func main() {
	priStr := "43e57262b2632824c8ffbc016a945b89b5854f1322f6c488ad2aabcc5adb533ec55a59054a7d3ab0133d845aabb01efce8d088adb8fe778521db1323cce05362"
	pri, err := hexx.Decode(priStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	ninecli.SetPriKey(pri)
	resp := ninecli.Rest[rest.Empty, nineora.Nineora](
		ninerpc.NineoraGet,
		rest.NewEmpty(),
	)
	fmt.Println(resp.JSON())
}
