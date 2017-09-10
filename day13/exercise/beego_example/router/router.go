package router

import (
	"go_dev/day13/exercise/beego_example/controller"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &controller.IndexController{}, "*:Index")
}
