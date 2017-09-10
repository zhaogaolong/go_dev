package router

import (
	"github.com/astaxie/beego"
	"go_dev/day13/example/beego_example/controller"
)


func init(){
	beego.Router("/index", &controller.IndexController{}, "*:Index")
}