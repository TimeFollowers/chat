package main

import (
	"wutool.cn/chat/server/global"
	"wutool.cn/chat/server/initialize"
	"wutool.cn/chat/server/router"
	"wutool.cn/chat/server/service"
)

func main() {
	global.VIPER = initialize.Viper()
	global.DB = initialize.InitMysql()
	global.REDIS = initialize.InitRedis()
	global.SESSION = initialize.InitSession()
	if global.DB != nil {
		initialize.RegisterTables()
	}

	go service.Manager.Start()

	r := router.Router()

	r.Run(":8080")
}
