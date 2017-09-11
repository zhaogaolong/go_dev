package router

import (
	"go_dev/day13/exercise/web_admin/controller"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &controller.AppController{}, "*:AppList")
	beego.Router("/app/list", &controller.AppController{}, "*:AppList")
	beego.Router("/app/apply", &controller.AppController{}, "*:AppApply")
	beego.Router("/app/create", &controller.AppController{}, "*:AppCreate")
	
	beego.Router("/log/list", &controller.LogController{}, "*:LogList")
	beego.Router("/log/apply", &controller.LogController{}, "*:LogApply")
	beego.Router("/log/create", &controller.LogController{}, "*:LogCreate")
}
