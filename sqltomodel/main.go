package main

import (
	"modules/sqltomodel/conf"
	"modules/sqltomodel/dbtools"
)

func main() {
	// 首先初始化struct对象，用于保存配置文件信息
	var confS conf.Config
	// 通过toml.DecodeFile将toml配置文件的内容，解析到struct对象
	err := conf.GetConf("C:/Users/cyz/Desktop/git-project/modules/sqltomodel/conf/mysql_map.toml", &confS)
	if err != nil {
		return
	}
	for _, v := range confS.Database {
		var MasterDbConfig = &dbtools.DbConf{
			Host:   v.Host,
			Port:   v.Port,
			User:   v.User,
			Pwd:    v.Pwd,
			DbName: "test",
		}
		dbtools.Init(MasterDbConfig)             //初始化数据库
		dbtools.Generate("C:/Users/cyz/Desktop/git-project/modules/sqltomodel/models/", "test") //生成指定表信息，可变参数可传入多个表名
	}

}
