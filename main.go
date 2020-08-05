package main

import (
	"github.com/astaxie/beego"
	_ "github.com/chenxifun/bsn-sdk-example-go/routers"
	"github.com/chenxifun/bsn-sdk-example-go/service/config"
	"github.com/wonderivan/logger"
)

func main() {

	if err := config.InitConfig(); err != nil {
		logger.Error("Config init failed")
		return
	}

	beego.Run()
}
