package controller

import (
	"fmt"
	"go_dev/day13/exercise/web_admin/model"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AppController struct {
	beego.Controller
}

func (p *AppController) AppList() {
	appList, err := model.GetAllAppInfo()

	p.Layout = "layout/layout.html"
	// 如果出现错误就吧错误的信息渲染到错误页面即可
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("服务器繁忙")
		p.TplName = "app/error.html"
		logs.Warn("get app list failed , err:%v", err)
		return
	}

	logs.Debug("get app list success, data:%v", appList)
	// 如果成功获取信息，就把信息渲染到页面中
	p.Data["applist"] = appList
	p.TplName = "app/index.html"

}

func (p *AppController) AppApply() {
	logs.Debug("appApply html")

}
