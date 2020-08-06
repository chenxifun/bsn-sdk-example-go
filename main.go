package main

import (
	"github.com/astaxie/beego"
	_ "github.com/chenxifun/bsn-sdk-example-go/routers"
)

func main() {

	//if err := config.InitFabricConfig(); err != nil {
	//	logger.Error("Config init failed")
	//	return
	//}
	//注释
	beego.Run()
}
