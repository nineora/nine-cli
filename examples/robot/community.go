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

type CommunityAgent struct {
	Type     string
	UID      string
	Email    string
	Wallet   *nineora.Wallet
	ChainAcc *nineora.Account
	Superior nineora.Link
	init     bool
	mu       sync.Mutex
}

func (m *CommunityAgent) getWalletLink() nineora.Link {
	return nineora.LinkOf(NetworkID, "member.wallet:", m.Type, m.UID)
}

func (m *CommunityAgent) getChainAccLink() nineora.Link {
	return nineora.LinkOf(NetworkID, "community.account:", m.Type, m.UID)
}

func (m *CommunityAgent) getNodeLink() nineora.Link {
	return nineora.LinkOf(NetworkID, "community.node:", m.Type, m.UID)
}

var gBranchDB = make(map[string]*CommunityAgent)
var gBranchDBMu sync.Mutex

func RandBranch() *CommunityAgent {
	gBranchDBMu.Lock()
	defer gBranchDBMu.Unlock()
	if len(gBranchDB) <= MaxBranchCount {
		m := NewCommunity("branch", "")
		if !m.init {
			sys.Exit(errors.System("branch.init err"))
		}
		gBranchDB[m.UID] = m
		return m
	}

	var thisM *CommunityAgent
	idx := rand.IntN(len(gBranchDB))
	var i = 0
	for _, m := range gBranchDB {
		if i == idx {
			thisM = m
			break
		}
	}
	thisM.doInit()
	return thisM
}

var gDivisionDB = make(map[string]*CommunityAgent)
var gDivisionDBMu sync.Mutex

func RandDivision() *CommunityAgent {
	branch := RandBranch()
	gDivisionDBMu.Lock()
	defer gDivisionDBMu.Unlock()
	if len(gDivisionDB) <= MaxDivisionCount {
		m := NewCommunity("division", branch.getNodeLink())
		if !m.init {
			sys.Exit(errors.System("division.init err"))
		}
		gDivisionDB[m.UID] = m
		return m
	}
	var thisM *CommunityAgent
	idx := rand.IntN(len(gDivisionDB))
	var i = 0
	for _, m := range gDivisionDB {
		if i == idx {
			thisM = m
			break
		}
	}
	thisM.doInit()
	return thisM
}

var gAssociateDB = make(map[string]*CommunityAgent)
var gAssociateDBMu sync.Mutex

func RandAssociate() *CommunityAgent {
	division := RandDivision()
	gAssociateDBMu.Lock()
	defer gAssociateDBMu.Unlock()
	if len(gAssociateDB) <= MaxAssociateCount {
		agent := NewCommunity("associate", division.getNodeLink())
		if !agent.init {
			sys.Exit(errors.System("associate.init err"))
		}
		gAssociateDB[agent.UID] = agent
		return agent
	}

	var thisM *CommunityAgent
	idx := rand.IntN(len(gAssociateDB))
	var i = 0
	for _, m := range gAssociateDB {
		if i == idx {
			thisM = m
			break
		}
	}
	thisM.doInit()
	return thisM
}

func NewCommunity(t string, superior nineora.Link) *CommunityAgent {
	m := &CommunityAgent{
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

	return m
}

func (m *CommunityAgent) doInit() {
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
		fmt.Println("comm WalletBind:", m.Type, err)
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
		fmt.Println("comm AccCreate:", m.Type, err)
		return
	}
	m.ChainAcc = acc

	switch m.Type {
	case "branch":
		_, err = dakaapi.BranchCreate(&dakarpc.BranchCreateReq{
			Link:    m.getNodeLink(),
			Benefit: m.getChainAccLink(),
		})
	case "division":
		_, err = dakaapi.DivisionCreate(&dakarpc.DivisionCreateReq{
			Link:    m.getNodeLink(),
			Branch:  m.Superior,
			Benefit: m.getChainAccLink(),
		})
	case "associate":
		_, err = dakaapi.AssociateCreate(&dakarpc.AssociateCreateReq{
			Link:     m.getNodeLink(),
			Division: m.Superior,
			Benefit:  m.getChainAccLink(),
		})
	}
	if err != nil {
		fmt.Println("comm Create:", m.Type, err)
		return
	}

	m.init = true

	fmt.Println("\n\n[new "+m.Type+"]\n", jsonx.MustJSON2String(m))
}
