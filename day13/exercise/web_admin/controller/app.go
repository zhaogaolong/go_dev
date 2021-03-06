package controller

import (
	"strings"
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
	p.Layout = "layout/layout.html"
	p.TplName = "app/apply.html"
}

func (p *AppController) AppCreate() {
	p.Layout = "layout/layout.html"

	// 获取所有到信息
	appName := p.GetString("project-name")
	appType := p.GetString("project-type")
	appIps := p.GetString("ip-list")
	appDevelopPath := p.GetString("develop-path")
	if len(appName) == 0 || len(appType) == 0 || len(appIps) == 0 || len(appDevelopPath) == 0{
		p.Data["Error"] = fmt.Sprintf("非法到参数")
		p.TplName = "app/error.html"
		return
	}
	appInfo := &model.AppInfo{}
	appInfo.AppName = appName
	appInfo.AppType = appType
	appInfo.DevelopPath = appDevelopPath
	appInfo.IP = strings.Split(appIps, ",")

	err := model.CreateApp(appInfo)
	if err != nil{
		p.Data["Error"] = fmt.Sprintf("创建项目失败，数据库繁忙")
		p.TplName = "app/error.html"
		return
	}

	// 创建完成后重定向到项目列表
	p.Redirect("/app/list", 302)

}