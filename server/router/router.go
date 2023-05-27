package router

import (
	"github.com/gin-gonic/gin"
	"wutool.cn/chat/server/middleware"
	"wutool.cn/chat/server/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors()) // 跨域
	r.GET("/ping", service.Ping)
	// 用户登录
	r.POST("/login", service.Login)

	// 发送验证码
	r.POST("/send/code", service.SendCode)
	// 用户注册
	r.POST("/register", service.Register)
	// 发送接收消息
	r.GET("/ws", service.Ws)
	auth := r.Group("/u", middleware.AuthCheck())

	// 用户详情
	auth.GET("/user/detail", service.UserDetail)
	// 查询指定用户的个人信息
	auth.GET("/user/query", service.UserQuery)
	// 聊天列表
	auth.GET("/chat/list", service.MessageList)
	// 添加用户
	auth.POST("/user/add", service.UserAdd)
	// 获取用户列表
	auth.GET("/user/list", service.UserList)
	//删除好友
	auth.DELETE("/user/delete", service.UserDelete)
	return r
}
