package controllers

import (
	"github.com/astaxie/beego"
	"github.com/chenxifun/bsn-sdk-example-go/models"
)

type baseController struct {
	beego.Controller
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
