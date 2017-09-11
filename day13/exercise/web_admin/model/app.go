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
	IP []string
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

func CreateApp(info *AppInfo)(err error){
	//开始一个mysql事务
	conn, err := Db.Begin()
	if err != nil{
		logs.Warn("CreateApp failed, err:%v", err)
		return
	}

	// 如果一旦出现问题，就利用defer回滚数据库
	defer func(){
		if err != nil{
			conn.Rollback()		
			return
		}
	}()



	sqlCmd := "insert into tbl_app_info(app_name, app_type, develop_path)value(?,?,?)"
	r, err := conn.Exec(sqlCmd, info.AppName, info.AppType, info.DevelopPath)
	if err != nil{
		logs.Warn("CreateApp failed, Db.exec err:%v", err)
	}
	id, err := r.LastInsertId()
	if err != nil{
		logs.Warn("CreateApp failed, r.id err:%v", err)
	}
	for _, ip := range info.IP{
		_, err = conn.Exec("insert into tbl_app_ip(app_id, ip)value(?,?)", id, ip)
		if err != nil{
			logs.Warn("Create App IP error:%v", err)
			return
		}
	}
	conn.Commit()
	return
}