package router

import (
	"github.com/astaxie/beego"
)


func init(){
	beego.Router("/index", &)
}