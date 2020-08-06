package service

import (
	"github.com/chenxifun/bsn-sdk-example-go/service/config"
	"github.com/chenxifun/bsn-sdk-example-go/service/fabricChain"
	"time"
)

func NewChainService() (ChainService, error) {
	chainCodeId := ""

	fcs, err := fabricChain.NewFabricChainService(config.GetFabricConfig(), chainCodeId)

	return fcs, err
}

type ChainService interface {

	//向链上存储数据
	ChainSet(key, value string) (txId string, err error)
	//读取链上数据
	ChainGet(key string) (value string, err error)
	//获取一个交易的详情
	ChainTransInfo(txid string) (*TransInfo, error)
}

type TransInfo struct {
	BlockHash   string
	BlockNumber uint64

	CreateName string
	CreateTime time.Time
}
