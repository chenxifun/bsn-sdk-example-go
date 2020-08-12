package fabricChain

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client/fabric"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/node"
	res "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/fabric/node"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/crypto"

	"github.com/chenxifun/bsn-sdk-example-go/models/errors"
	"github.com/chenxifun/bsn-sdk-example-go/service"
)

func NewFabricChainService(conf *config.Config, chainCode string) (*fabricChainService, error) {

	if conf == nil {
		return nil, errors.New("Config has error")
	}

	client, err := fabric.InitFabricClient(conf)
	if err != nil {
		return nil, err
	}

	fcs := &fabricChainService{
		client:      client,
		chainCodeId: chainCode,
	}

	return fcs, nil
}

type fabricChainService struct {
	client      *fabric.FabricClient
	chainCodeId string
}

func (f *fabricChainService) reqChainCode(funcName string, args []string) (*res.TranResDataBody, error) {
	name := ""
	nonce, _ := crypto.GetRandomNonce()

	body := req.TransReqDataBody{
		UserName:     name,
		Nonce:        base64.StdEncoding.EncodeToString(nonce),
		ChainCode:    f.chainCodeId,
		FuncName:     funcName,
		Args:         args,
		TransientMap: make(map[string]string),
	}

	res, err := f.client.ReqChainCode(body)
	if err != nil {
		return nil, err
	}
	if res.Header.Code == 0 && res.Body.CCRes.CCCode == 200 {
		return res.Body, nil
	} else {
		if res.Header.Code == -1 {
			//信息不存在 是链码返回的异常信息，在这里处理，
			//
			if strings.Contains(res.Header.Msg, "信息不存在") {
				return nil, errors.New("key is not found")
			} else if strings.Contains(res.Header.Msg, "该键值信息已存在") {
				return nil, errors.New("key is exist")
			} else {
				return nil, errors.New(res.Header.Msg)
			}
		} else {
			return nil, errors.New("save has error")
		}

	}

}

func (f *fabricChainService) ChainSet(key, value string) (txId string, err error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}
	if value == "" {
		return "", errors.New("value cannot be empty")
	}
	var args []string
	args = append(args, fmt.Sprintf("{\"baseKey\":\"%s\",\"baseValue\":\"%s\"}", key, value))
	funcName := "set"

	res, err := f.reqChainCode(funcName, args)
	if err != nil {
		return "", err
	}
	return res.BlockInfo.TxId, err
}

func (f *fabricChainService) ChainGet(key string) (value string, err error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}

	var args []string
	args = append(args, key)
	funcName := "get"
	res, err := f.reqChainCode(funcName, args)
	if err != nil {
		return "", err
	}

	return res.CCRes.CCData, nil
}

func (f *fabricChainService) ChainTransInfo(txid string) (*service.TransInfo, error) {

	body := req.TxTransReqDataBody{
		TxId: txid,
	}

	res, err := f.client.GetTransInfo(body)
	if err != nil {
		return nil, err
	}
	if res.Header.Code == 0 {

		info := &service.TransInfo{
			BlockHash:   res.Body.BlockHash,
			BlockNumber: res.Body.BlockNumber,
			CreateName:  res.Body.CreateName,
			CreateTime:  time.Unix(res.Body.TimeSpanSec, res.Body.TimeSpanNsec),
		}

		return info, nil
	} else {
		return nil, errors.New(res.Header.Msg)
	}

}
