package robot

import (
	"fmt"
	"github.com/avast/retry-go"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/idx"
	"github.com/hootuu/gelato/sys"
	"github.com/hootuu/gelato/types/jsonx"
	"github.com/nineora/nine-cli/dakaapi"
	"github.com/nineora/nine-cli/nineapi"
	"github.com/nineora/nineora/daka/dakarpc"
	"github.com/nineora/nineora/nine/ninekey"
	"github.com/nineora/nineora/nine/nineora"
	"github.com/nineora/nineora/nine/ninerpc"
	"math/rand/v2"
	"sync"
	"time"
)

type AreaAgent struct {
	Type     string
	UID      string
	Email    string
	Wallet   *nineora.Wallet
	ChainAcc *nineora.Account
	Superior nineora.Link
	init     bool
	mu       sync.Mutex
}

func (m *AreaAgent) getWalletLink() nineora.Link {
	return nineora.LinkOf(NetworkID, "member.wallet:", m.Type, m.UID)
}

func (m *AreaAgent) getChainAccLink() nineora.Link {
	return nineora.LinkOf(NetworkID, "area.account:", m.Type, m.UID)
}

func (m *AreaAgent) getNodeLink() nineora.Link {
	return nineora.LinkOf(NetworkID, "area.node:", m.Type, m.UID)
}

var gProvinceDB = make(map[string]*AreaAgent)
var gProvinceDBMu sync.Mutex

func RandProvince() *AreaAgent {
	gProvinceDBMu.Lock()
	defer gProvinceDBMu.Unlock()
	if len(gProvinceDB) <= MaxProvinceCount {
		m := NewArea("province", "")
		if !m.init {
			sys.Exit(errors.System("province.init err"))
		}
		gProvinceDB[m.UID] = m
		return m
	}
	var thisM *AreaAgent
	idx := rand.IntN(len(gProvinceDB))
	var i = 0
	for _, m := range gProvinceDB {
		if i == idx {
			thisM = m
			break
		}
	}
	thisM.doInit()
	return thisM
}

var gCityDB = make(map[string]*AreaAgent)
var gCityDBMu sync.Mutex

func RandCity() *AreaAgent {
	province := RandProvince()
	gCityDBMu.Lock()
	defer gCityDBMu.Unlock()
	if len(gCityDB) <= MaxCityCount {
		m := NewArea("city", province.getNodeLink())
		if !m.init {
			sys.Exit(errors.System("city.init err"))
		}
		gCityDB[m.UID] = m
		return m
	}
	var thisM *AreaAgent
	idx := rand.IntN(len(gCityDB))
	var i = 0
	for _, m := range gCityDB {
		if i == idx {
			thisM = m
			break
		}
	}
	thisM.doInit()
	return thisM
}

var gZoneDB = make(map[string]*AreaAgent)
var gZoneDBMu sync.Mutex

func RandZone() *AreaAgent {
	city := RandCity()
	gZoneDBMu.Lock()
	defer gZoneDBMu.Unlock()
	if len(gZoneDB) <= MaxZoneCount {
		m := NewArea("zone", city.getNodeLink())
		if !m.init {
			sys.Exit(errors.System("zone.init err"))
		}
		gZoneDB[m.UID] = m
		return m
	}
	var thisM *AreaAgent
	idx := rand.IntN(len(gZoneDB))
	var i = 0
	for _, m := range gZoneDB {
		if i == idx {
			thisM = m
			break
		}
	}
	thisM.doInit()
	return thisM
}

func NewArea(t string, superior nineora.Link) *AreaAgent {
	m := &AreaAgent{
		Type:     t,
		UID:      idx.New(),
		Email:    idx.New() + "@nineora.com",
		Wallet:   nil,
		ChainAcc: nil,
		Superior: superior,
		init:     false,
		mu:       sync.Mutex{},
	}
	retry.Do(func() error {
		m.doInit()
		return nil
	}, retry.Attempts(10), retry.MaxDelay(300*time.Millisecond))
	superiorPut(m.getNodeLink(), m.Superior)
	return m
}

func (m *AreaAgent) doInit() {
	if m.init {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.init {
		return
	}
	walletLink := m.getWalletLink()
	wallet, err := nineapi.WalletBind(&ninerpc.WalletBindReq{
		Link: walletLink,
		Seed: nineora.SeedOfEmail(m.Email),
		Meta: nil,
	})
	if err != nil {
		fmt.Println("area.WalletBind:", err)
		return
	}
	m.Wallet = wallet
	accLink := m.getChainAccLink()
	acc, err := nineapi.AccCreate(&ninerpc.AccCreateReq{
		Link:       accLink,
		WalletLink: walletLink,
		KeyType:    ninekey.Ed25519,
		NetworkID:  NetworkID,
		Password:   []byte("12345678"),
		Meta:       nil,
	})
	if err != nil {
		fmt.Println("area.AccCreate: ", err)
		return
	}
	m.ChainAcc = acc

	switch m.Type {
	case "province":
		_, err = dakaapi.ProvinceCreate(&dakarpc.ProvinceCreateReq{
			Link:    m.getNodeLink(),
			Benefit: m.getChainAccLink(),
		})
	case "city":
		_, err = dakaapi.CityCreate(&dakarpc.CityCreateReq{
			Link:     m.getNodeLink(),
			Province: m.Superior,
			Benefit:  m.getChainAccLink(),
		})
	case "zone":
		_, err = dakaapi.ZoneCreate(&dakarpc.ZoneCreateReq{
			Link:    m.getNodeLink(),
			City:    m.Superior,
			Benefit: m.getChainAccLink(),
		})
	}
	if err != nil {
		fmt.Println("area.AccCreate: ", m.Type, err)
		return
	}

	m.init = true

	fmt.Println("\n\n[new "+m.Type+"]\n", jsonx.MustJSON2String(m))
}
