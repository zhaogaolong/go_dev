package controller

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type IndexController struct {
	beego.Controller
}

func (p *IndexController) Index() {
	logs.Debug("index get ....")
	p.TplName = "index/index.html"
}
func init() {
	beego.Router("/index", &IndexController{}, "*:Index")
}
