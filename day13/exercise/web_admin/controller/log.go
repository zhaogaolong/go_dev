package controller

import (
	"fmt"
	"go_dev/day13/exercise/web_admin/model"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LogController struct {
	beego.Controller
}

func (p *LogController) LogList() {
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

func (p *LogController) LogApply() {
	logs.Debug("appApply html")
	p.Layout = "layout/layout.html"
	p.TplName = "log/apply.html"
}

func (p *LogController) LogCreate() {
	p.Layout = "layout/layout.html"

	// 获取所有到信息
	appName := p.GetString("project-name")
	logPath := p.GetString("log-path")
	topic := p.GetString("topic")
	if len(appName) == 0 || len(logPath) == 0 || len(topic) == 0{
		p.Data["Error"] = fmt.Sprintf("非法到参数")
		p.TplName = "app/error.html"
		return
	}

	logInfo := &model.LogInfo{}
	logInfo.AppName = appName
	logInfo.LogPath = logPath
	logInfo.Topic = topic

	err := model.CreateLog(logInfo)
	if err != nil{
		logs.Warn("create log filed, err:%v", err)
		p.Data["Error"] = fmt.Sprintf("创建项目失败，数据库繁忙")
		p.TplName = "app/error.html"
		return
	}
	// 创建完成后重定向到项目列表
	p.Redirect("/app/list", 302)

}