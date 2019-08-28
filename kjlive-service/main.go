package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"kjlive-service/routers"
	"kjlive-service/conf"
	"kjlive-service/logs"
)

func init() {
	// 初始化数据库
	// read
	orm.RegisterDataBase("default", "mysql", conf.Settings.DatabaseSource)
	if conf.Settings.DatabaseSourceReplica != "" {
		// write
		orm.RegisterDataBase("replica", "mysql", conf.Settings.DatabaseSourceReplica)
	} else {
		// write
		orm.RegisterDataBase("replica", "mysql", conf.Settings.DatabaseSource)
	}
	//orm.Debug = true
}

func main() {
	//  初始化log文件
	infoLog, errorLog, actionLog := logs.LogInit()
	defer infoLog.Close()
	defer errorLog.Close()
	defer actionLog.Close()
	routers.Run()
}
