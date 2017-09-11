package main

import (
	"go_dev/day13/exercise/web_admin/model"
	_ "go_dev/day13/exercise/web_admin/router"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"
)

func initDb() (err error) {
	database, err := sqlx.Open("mysql", "root:root@tcp(192.168.1.201:3306)/golang")
	if err != nil {
		return
	}
	model.InitDb(database)
	return
}

func main() {
	err := initDb()
	if err != nil {
		logs.Error("initDb failed, err:%v", err)
		return
	}
	beego.Run()
}
