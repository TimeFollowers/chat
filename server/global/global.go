package global

import (
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"wutool.cn/chat/server/config"
)

var (
	CONFIG  config.Server
	VIPER   *viper.Viper
	DB      *gorm.DB
	REDIS   *redis.Client
	SESSION *sessions.CookieStore
)
