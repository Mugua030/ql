package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

var (
	ServerName string
	ServerAddr string
	ServerPort string

	RedisAddr  string
	RedisPort  string
	RedisDbNum string

	MysqlAddr   string
	MysqlPort   string
	MysqlDbName string

	FastDfsAddr string
	FastDfsPort string
)

func InitConfig() {
	appconf, err := config.NewConfig("ini", "./conf/app.conf")
	if err != nil {
		beego.Debug(err)
		return
	}
	ServerName := appconf.String("appname")
	ServerAddr := appconf.String("httpaddr")
	ServerPort := appconf.String("httpport")

	RedisAddr := appconf.String("redisaddr")
	RedisPort := appconf.String("redisport")
	RedisDbNum := appconf.String("redisdbnum")

	MysqlAddr := appconf.String("mysqladdr")
	MysqlPort := appconf.String("mysqlport")
	MysqlDbName := appconf.String("mysqldbname")

	FastDfsAddr := appconf.String("fastdfsaddr")
	FastDfsPort := appconf.String("fastdfsport")

	return
}

func init() {
	InitConfig()
}
