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

## Chain
### SUI
#### DEV
```ini
NineoraPackageAddr: 0x06fbe9d1e6542bcd1f5192c74bbd751bf4535c31681f04f487c1753f64f4714f
Nineora: 0x0c47f39c954cc316a14453019fe576a38faa4c113dbbde23b976112d55bddf49

DakaNetworkAddr: 0xb855c8aa2421b7e454a12a9115d41671e8af6e9c9e26e4ab9026156f2e9b6db6
DakaPackageAddr: 0x6045389f3757c0c9b003e47dcedebcc7c8b9eeb2e4d98d1ede03fc631b2408bc
UnimktBizAddr: 0xf7ba9db527c700e2c77bc7957fba36531c4cc93248a674a67810f55b0aa838ec
KAKA: 0xe741eafb7ccadafea18fec20ed03f1547e3965c407cf0cc27f8894beb08be38f
DAKA: 0xb0df5fa49048167065e27962475305b748ca9d902b03153a74fba2a93cc1bd77
DIGI: 0x82d22d6a60305d57d92ff3de77eb42d7a7fa091c79274903f24066a1852cbe67
DADA: 0xad8bd4a03353981fc4fef207921abd0af9b09ea47a6773b6b701cced3aaa48c4
```