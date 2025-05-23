package main

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/hexx"
	"github.com/nineora/nine-cli/examples/robot"
	"github.com/nineora/nine-cli/ninecli"
)

/**

{"level":"info","time":"2025-04-27 10:46:46.484","msg":"daka.guard","id":"d06pl5iubk6gcfpvifk0","pub":"2e7c4fbdb886eaa7a8c88d97c2d0ba089b2654d1f6ff1e2e8ce426a448eb46d2","pri":"b2eff61f98a7c482d46d1ca0f80fecd0465b2c0a2bd7470a2cfdcf52b72c4c3e2e7c4fbdb886eaa7a8c88d97c2d0ba089b2654d1f6ff1e2e8ce426a448eb46d2"}
*/

func main() {
	priStr := "c7be627e71307f424879a66ba939c615aa2887fb9a0848a7ea01167371e449e90ea081ff61a6b0e3af38c6d5112866b5104256b4f1c1363dfdea33460172ddbc"
	pri, err := hexx.Decode(priStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	ninecli.UsePrivateKey(pri)

	robot.Running()
}
