package router

import (
	"github.com/astaxie/beego"
	"go_dev/day13/exercise/web_admin/controller"
)


func init(){
	beego.Router("/index", &controller.AppController{}, "*:Index")
}