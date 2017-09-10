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

	// json result 
	data := make(map[string]interface{})
	data["code"] = 200
	data["message"] = "success"
	p.Data["json"] = &data
	p.ServeJSON(true)

}
