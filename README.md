# nine-cli
The Client For Nineora

## import
```shell
cd $WORK_HOME
go get https://github.com/nineora/nine-cli.git
```

## config
### use ENV
```ini
NINEORA_LOC_GUARD: cvqjvokdhg6lal704kn0 (Get From The Nineora Team)
NINEORA_LOC_GW: http://localhost:8080 (Get From The Nineora Team)
# Other config see: ninecli.init()
```
### use config file
```yaml
nineora:
  loc:
    guard: cvqjvokdhg6lal704kn0 (Get From The Nineora Team)
    gw: http://localhost:8080 (Get From The Nineora Team)
```

## Code With Cli

```go
package main

import (
	"github.com/nineora/nine-cli/ninecli"
	"github.com/nineora/nineora/nine/ninekey"
	"github.com/nineora/nineora/nine/nineora"
	"github.com/nineora/nineora/nine/ninerpc"
)
const NetworkID = "*********"

func UserWalletLink(uid string) nineora.Link {
	return nineora.LinkOf(NetworkID, "user.wallet", uid)
}

func main() {
	privateKey := (ninekey.PrivateKey)"******"
	ninecli.UsePrivateKey(privateKey)
	
	uid := "********" //my biz system user id for example
	userMobi := "+8618688888888" //my biz system user mobile phone for example
	userWalletLink := UserWalletLink(uid)
	_, err := nineapi.WalletBind(&ninerpc.WalletBindReq{
		Link: userWalletLink,
		Seed: nineora.SeedOfCell(userMobi),
		Meta: map[string]interface{}{
			"uid": uid,
			"birth": "19900808",
		},
	})
	if err != nil {
		fmt.Println("WalletBind err", err)
		return
	}
	// ...
}
```

## Example
more example see: the package examples