package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/chenxifun/bsn-sdk-example-go/models"
)

type baseController struct {
	beego.Controller
}

func (c *baseController) GetReqData(data interface{}) {
	s := c.Ctx.Input.RequestBody
	err := json.Unmarshal(s, data)
	if err != nil {
		c.ResultError(err)
	}
}

func (c *baseController) ResultSuccess(data interface{}) {
	result := models.Result{
		Code:   1,
		Msg:    "success",
		Result: data,
	}
	c.Result(result)
}

func (c *baseController) ResultError(err error) {
	result := models.Result{
		Code: 0,
		Msg:  err.Error(),
	}
	c.Result(result)
}

func (c *baseController) Result(data interface{}) {
	c.Data["json"] = data
	c.ServeJSON()
	c.StopRun()
}
