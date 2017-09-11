package model

import (
	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"
)

type AppInfo struct {
	AppID       int    `db:"app_id"`
	AppName     string `db:"app_name"`
	AppType     string `db:"app_type"`
	CreateTime  string `db:"create_time"`
	DevelopPath string `db:"develop_path"`
}

var (
	Db *sqlx.DB
)

func InitDb(db *sqlx.DB) {
	Db = db
}

func GetAllAppInfo() (appList []AppInfo, err error) {
	sqlCmd := "select app_id, app_name, app_type, create_time, develop_path from tbl_app_info"
	err = Db.Select(&appList, sqlCmd)
	if err != nil {
		logs.Warn("Get all app info failed, err:%v", err)
	}
	return
}
