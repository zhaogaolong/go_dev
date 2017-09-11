package model

import (
	"github.com/astaxie/beego/logs"
)

type LogInfo struct {
	AppID       int    `db:"app_id"`
	AppName	string   `db:"app_name"`
	LogID    int `db:"log_id"`
	Status int `db:"log_id"`
	LogPath     string `db:"log_path"`
	CreateTime  string `db:"create_time"`
	Topic string `db:"topic"`
}


func GetAllLogInfo() (logInfo []LogInfo, err error) {
	sqlCmd := "select a.app_id, b.app_name, a.topic, a.create_time, a.status,a.log_path from tbl_log_info a, tbl_app_info b where a.app_id=b.app_id"
	err = Db.Select(&logInfo, sqlCmd)
	if err != nil {
		logs.Warn("Get all log info failed, err:%v", err)
	}
	return
}

func CreateLog(info *LogInfo)(err error){
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
	var appId []int
	err = Db.Select(&appId, "select app_id from tbl_app_info where app_name=?", info.AppName)
	if err != nil || len(appId) == 0{
		logs.Warn("Appid failed, Db.exec err:%v", err)
		return
	}
	info.AppID = appId[0]


	sqlCmd := "insert into tbl_log_info(app_id, log_path, topic)value(?,?,?)"
	_, err = conn.Exec(sqlCmd, info.AppID, info.LogPath, info.Topic)
	if err != nil{
		logs.Warn("CreateApp failed, Db.exec err:%v", err)
	}
	conn.Commit()
	return
}