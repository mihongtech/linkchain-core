package poa

import (
	"errors"
	"github.com/mihongtech/linkchain-core/common/util/event"
	"github.com/mihongtech/linkchain-core/node/config"
	event2 "github.com/mihongtech/linkchain-core/node/event"
	"sync"
	"time"

	"github.com/mihongtech/linkchain-core/common/util/log"
	"github.com/mihongtech/linkchain-core/core/meta"
	"github.com/mihongtech/linkchain-core/node/bcsi"
	"github.com/mihongtech/linkchain-core/node/chain"
	"github.com/mihongtech/linkchain-core/node/pool"
)

type Config struct {
	chain         chain.Chain
	txPool        pool.TxPool
	bcsiAPI       bcsi.BCSI
	newBlockEvent *event.TypeMux
}

func NewConfig(chain chain.Chain, txPool pool.TxPool, bcsiAPI bcsi.BCSI, newBlockEvent *event.TypeMux) *Config {
	return &Config{chain, txPool, bcsiAPI, newBlockEvent}
}

type Miner struct {
	poa           *Poa
	chain         chain.Chain
	txPool        pool.TxPool
	bcsiAPI       bcsi.BCSI
	isMining      bool
	minerMtx      sync.Mutex
	newBlockEvent *event.TypeMux
}

func NewMiner(poa *Poa) *Miner {
	return &Miner{isMining: false, poa: poa}
}

func (m *Miner) Setup(i interface{}) bool {
	cfg := i.(*Config)
	m.chain = cfg.chain
	m.txPool = cfg.txPool
	m.bcsiAPI = cfg.bcsiAPI
	m.newBlockEvent = cfg.newBlockEvent
	return true
}

func (m *Miner) Start() bool {
	log.Info("Miner start...")
	go m.StartMine()
	return true
}

func (m *Miner) Stop() {
	log.Info("Miner stop...")
	go m.StopMine()
}

func (m *Miner) MineBlock() (*meta.Block, error) {
	best := m.chain.GetBestBlock()
	block, err := CreateBlock(best.GetHeight(), *best.GetBlockID())
	if err != nil {
		log.Error("Miner", "New Block error", err)
		return nil, err
	}
	signer := m.poa.getBlockSigner(block)
	//coinbase := CreateCoinBaseTx(signer, meta.NewAmount(config.DefaultBlockReward), block.GetHeight())
	//block.SetTx(*coinbase)

	txs := m.txPool.GetAllTransaction()
	txs = m.bcsiAPI.FilterTx(txs)
	block.SetTx(txs...)

	if !IsBestBlockOffspring(m.chain, block) {
		m.removeBlockTxs(block)
		return nil, errors.New("current block is not block prev")
	}

	block.Header.Status, err = m.bcsiAPI.GetBlockState(*best.GetBlockID()) //The block status is prev block status
	if err != nil {
		log.Error("Miner", "Get Last Block State error", err)
		m.removeBlockTxs(block)
		return nil, err
	}
	block, err = RebuildBlock(block)
	if err != nil {
		log.Error("Miner", "Rebuild Block error", err)
		m.removeBlockTxs(block)
		return nil, err
	}

	err = m.signBlock(signer, block)
	log.Debug("Miner", "signer", signer.String())
	if err != nil {
		log.Error("Miner", "sign Block status error", err)
		m.removeBlockTxs(block)
		return nil, err
	}

	err = m.chain.ProcessBlock(block)
	if err != nil {
		m.removeBlockTxs(block)
		log.Error("Miner", "ProcessBlocks error", err)
		return nil, err
	}
	m.newBlockEvent.Post(event2.NewMinedBlockEvent{Block: block})
	return block, nil
}

func (m *Miner) signBlock(signer meta.Address, block *meta.Block) error {
	//TODO need to add poa sign
	//sign, err := m.walletAPI.SignMessage(signer, block.GetBlockID().CloneBytes())
	//if err != nil {
	//	return err
	//}
	//block.SetSign(sign)
	return nil
}

func (m *Miner) StartMine() error {
	m.minerMtx.Lock()
	if m.isMining {
		m.minerMtx.Unlock()
		return errors.New("the node is mining")
	}
	m.isMining = true
	m.minerMtx.Unlock()
	for true {
		m.minerMtx.Lock()
		tempMing := m.isMining
		m.minerMtx.Unlock()
		if !tempMing {
			break
		}
		m.MineBlock()
		time.Sleep(time.Duration(config.DefaultPeriod) * time.Second)
	}
	return nil
}

func (m *Miner) StopMine() {
	m.minerMtx.Lock()
	defer m.minerMtx.Unlock()
	m.isMining = false
}

func (m *Miner) GetInfo() bool {
	m.minerMtx.Lock()
	defer m.minerMtx.Unlock()
	return m.isMining
}

func (m *Miner) removeBlockTxs(block *meta.Block) {
	for index := range block.TXs.Txs {
		m.txPool.RemoveTransaction(*block.TXs.Txs[index].GetTxID())
	}
}

func IsBestBlockOffspring(chain chain.ChainReader, block *meta.Block) bool {
	return block.GetPrevBlockID().IsEqual(chain.GetBestBlock().GetBlockID())
}
