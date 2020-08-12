package controllers

import (
	"github.com/chenxifun/bsn-sdk-example-go/models"
	"github.com/chenxifun/bsn-sdk-example-go/service"
)

type FabricController struct {
	baseController
}

func (f *FabricController) Set() {
	data := &models.ChainData{}
	f.GetReqData(data)

	fcs, err := service.NewChainService()
	if err != nil {
		f.ResultError(err)
	}

	txid, err := fcs.ChainSet(data.Key, data.Value)

	if err != nil {
		f.ResultError(err)
	}

	f.ResultSuccess(&models.TxData{TxId: txid})

}
