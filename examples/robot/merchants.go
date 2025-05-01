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

var gMcDB = make(map[string]*Merchant)
var gMu4McDB sync.Mutex

func GetMerchantByLink(mcLink nineora.Link) *Merchant {
	for _, mc := range gMcDB {
		if mc.getNodeLink() == mcLink {
			return mc
		}
	}
	return nil
}

func RandMerchant() *Merchant {
	gMu4McDB.Lock()
	defer gMu4McDB.Unlock()
	if len(gMcDB) <= MaxMerchantCount {
		m := NewMerchant()
		if !m.init {
			sys.Exit(errors.System("mc.init err"))
		}
		gMcDB[m.UID] = m
		superiorPut(m.getNodeLink(), m.Associate)
		return m
	}

	var thisM *Merchant
	idx := rand.IntN(len(gMcDB))
	var i = 0
	for _, m := range gMcDB {
		if i == idx {
			thisM = m
			break
		}
	}
	thisM.doInit()
	return thisM
}

func NewMerchant() *Merchant {
	m := &Merchant{
		UID:         idx.New(),
		Email:       idx.New() + "@nineora.com",
		Wallet:      nil,
		ChainAcc:    nil,
		InvestRatio: rand.Uint64N(10000),
		init:        false,
		mu:          sync.Mutex{},
	}
	fmt.Println("mc.invest_ratio:", m.InvestRatio)
	retry.Do(func() error {
		m.doInit()
		return nil
	}, retry.Attempts(10), retry.MaxDelay(300*time.Millisecond))

	return m
}

type Merchant struct {
	UID         string
	Email       string
	Wallet      *nineora.Wallet
	ChainAcc    *nineora.Account
	Node        *nineora.Node
	Associate   nineora.Link
	InvestRatio uint64
	init        bool
	mu          sync.Mutex
}

func (m *Merchant) getWalletLink() nineora.Link {
	return nineora.LinkOf(NetworkID, "member.wallet:", m.UID)
}

func (m *Merchant) getChainAccLink() nineora.Link {
	return nineora.LinkOf(NetworkID, "merchant.account:", m.UID)
}

func (m *Merchant) getNodeLink() nineora.Link {
	return nineora.LinkOf(NetworkID, "merchant.node:", m.UID)
}

func (m *Merchant) doInit() {
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
		fmt.Println("mc.WalletBind: ", err)
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
		fmt.Println("mc.AccCreate: ", err)
		return
	}
	m.ChainAcc = acc

	associate := RandAssociate()

	_, err = dakaapi.MerchantCreate(&dakarpc.MerchantCreateReq{
		Link:        m.getNodeLink(),
		Associate:   associate.getNodeLink(),
		InvestRatio: m.InvestRatio,
		Benefit:     m.getChainAccLink(),
	})
	if err != nil {
		fmt.Println("mc.MerchantCreate: ", err)
		return
	}

	m.Associate = associate.getNodeLink()

	m.init = true

	fmt.Println("\n\n[new merchant]\n", jsonx.MustJSON2String(m))
}
