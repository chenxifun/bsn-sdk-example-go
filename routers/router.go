package routers

import (
	"github.com/astaxie/beego"
	"github.com/chenxifun/bsn-sdk-example-go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
