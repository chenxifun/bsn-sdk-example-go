package routers

import (
	"github.com/astaxie/beego"
	"github.com/chenxifun/bsn-sdk-example-go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	fabric := &controllers.FabricController{}

	fabricNS := beego.NSNamespace("/fabric", beego.NSInclude(fabric))
	ns := beego.NewNamespace("/api", fabricNS)

	beego.AddNamespace(ns)
}
