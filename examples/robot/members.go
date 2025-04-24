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

var gMemberDB = make(map[string]*Member)
var gMu4MemberDB sync.Mutex

func RandMember() *Member {
	gMu4MemberDB.Lock()
	defer gMu4MemberDB.Unlock()
	if len(gMemberDB) <= MaxMemberCount {
		m := NewMember()
		if !m.init {
			sys.Exit(errors.System("member.init err"))
		}
		gMemberDB[m.UID] = m
		return m
	}

	var thisM *Member
	idx := rand.IntN(len(gMemberDB))
	var i = 0
	for _, m := range gMemberDB {
		if i == idx {
			thisM = m
			break
		}
	}
	thisM.doInit()
	return thisM
}

func NewMember() *Member {
	m := &Member{
		UID:      idx.New(),
		Email:    idx.New() + "@nineora.com",
		Wallet:   nil,
		ChainAcc: nil,
		init:     false,
		mu:       sync.Mutex{},
	}
	retry.Do(func() error {
		m.doInit()
		return nil
	}, retry.Attempts(10), retry.MaxDelay(300*time.Millisecond))

	return m
}

type Member struct {
	UID      string
	Email    string
	Wallet   *nineora.Wallet
	ChainAcc *nineora.Account
	Merchant nineora.Link
	init     bool
	mu       sync.Mutex
}

func (m *Member) getWalletLink() nineora.Link {
	return nineora.LinkOf(NetworkID, "member.wallet:", m.UID)
}

func (m *Member) getChainAccLink() nineora.Link {
	return nineora.LinkOf(NetworkID, "member.account:", m.UID)
}

func (m *Member) getNodeLink() nineora.Link {
	return nineora.LinkOf(NetworkID, "member.node:", m.UID)
}

func (m *Member) doInit() {
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
		fmt.Println("members.nineapi.WalletBind", err)
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
		fmt.Println("members.nineapi.AccCreate", err)
		return
	}
	m.ChainAcc = acc

	merchant := RandMerchant()
	_, err = dakaapi.MemberCreate(&dakarpc.MemberCreateReq{
		Link:     m.getNodeLink(),
		Merchant: merchant.getNodeLink(),
		Benefit:  m.getChainAccLink(),
	})
	if err != nil {
		fmt.Println("members.nineapi.MemberCreate", err)
		return
	}

	m.Merchant = merchant.getNodeLink()

	m.init = true

	fmt.Println("\n\n[new member]\n", jsonx.MustJSON2String(m))
}
