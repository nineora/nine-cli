package ninecli

import (
	"github.com/hootuu/gelato/configure"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/rest"
	"github.com/nineora/nineora/nine/ninekey"
	"resty.dev/v3"
	"time"
)

var gRest = rest.NewRest()
var gCliPriKey []byte
var gCliGuardID string

func Rest[REQ any, RESP any](path string, data *REQ) (*RESP, *errors.Error) {
	if len(gCliPriKey) == 0 {
		return nil, errors.System("must set pri key first")
	}
	req := rest.NewRequest[REQ](gCliGuardID, data)
	resp := rest.Call[REQ, RESP](gRest, path, req, gCliPriKey)
	if !resp.Success {
		return nil, resp.Error
	}
	return resp.Data, nil
}

func UsePrivateKey(privateKey ninekey.PrivateKey) {
	gCliPriKey = privateKey
}

func init() {
	gCliGuardID = configure.GetString("nineora.loc.guard", "d08pq7kdhg6ivokiba5g")
	cfgBaseUrl := configure.GetString("nineora.loc.gw", "http://localhost:8080")
	cfgRetryWaitTime := configure.GetDuration("nineora.loc.wait.retry", 2)
	cfgRetryMaxWaitTime := configure.GetDuration("nineora.loc.wait.retry.max", 10)
	cfgTimeout := configure.GetDuration("nineora.loc.timeout", 60)
	gRest.SetBaseURL(cfgBaseUrl)
	gRest.WithTimeSetter(func(cli *resty.Client) {
		cli.SetRetryWaitTime(cfgRetryWaitTime * time.Second).
			SetRetryMaxWaitTime(cfgRetryMaxWaitTime * time.Second).
			SetTimeout(cfgTimeout * time.Second)
	})
}
