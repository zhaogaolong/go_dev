package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (p *MainController) Get() {
	p.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/index", &MainController{})
	beego.Run()
}
