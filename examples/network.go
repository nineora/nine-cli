package main

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/hexx"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/hootuu/gelato/io/serializer"
	"github.com/nineora/nine-cli/nineapi"
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/nine/ninerpc"
)

func main() {
	priStr := "ced49ab201255208746af52a4717e8851cfbc21ba4f686f5473abe0ff9047d01dabcc7e3a1276fbb562130f4036589ddb563717216c937681cc7cb0359934664"
	pri, err := hexx.Decode(priStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	ninecli.SetPriKey(pri)
	m, err := nineapi.NetworkQuery(&ninerpc.NetworkQueryReq{})
	fmt.Println(serializer.JsonMustTo(m), err)

	const networkID = "a2506828304e5dc0bc736014bad6fdaf"
	network, err := nineapi.NetworkGet(&ninerpc.NetworkGetReq{NID: networkID})
	fmt.Println(serializer.JsonMustTo(network), err)

	nodePage, err := nineapi.NodeQueryByNetwork(&ninerpc.NodeQueryByNetworkReq{
		NetworkID: networkID,
		Page: &pagination.Page{
			Size: 10,
			Numb: 1,
		},
		WithCore: false,
	})

	fmt.Println("-------network------")
	fmt.Println(serializer.JsonMustTo(nodePage), err)

	if len(nodePage.Data) == 0 {
		return
	}

	var superiorID string
	for _, n := range nodePage.Data {
		if n.Deep == 1 {
			superiorID = n.NID
			fmt.Println("superiorID:", superiorID)
			fmt.Println("tt", n.Ctrl)
			break
		}
	}

	nodePage, err = nineapi.NodeQueryBySuperior(&ninerpc.NodeQueryBySuperiorReq{
		Superior: superiorID,
		Page:     nil,
		WithCore: false,
		Deep:     []uint32{},
	})
	fmt.Println("-------sub------")
	fmt.Println(serializer.JsonMustTo(nodePage), err)
}
