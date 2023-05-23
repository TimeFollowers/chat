package initialize

import (
	"github.com/gin-gonic/gin"
	"wutool.cn/chat/server/service"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.POST("/login", service.Login)
	r.GET("/user/list", service.List)
	r.GET("/message/list", service.MessageList)
	r.GET("/ws", service.Ws)
	return r
}
