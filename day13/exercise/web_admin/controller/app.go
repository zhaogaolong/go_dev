package controller

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AppController struct {
	beego.Controller
}

func (p *AppController) Index() {
	logs.Debug("enter index controller...")
	p.Layout = "layout/layout.html"
	p.TplName = "app/index.html"

}
