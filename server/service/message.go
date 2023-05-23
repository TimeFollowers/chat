package service

import (
	"github.com/gin-gonic/gin"
	"wutool.cn/chat/server/chat"
	"wutool.cn/chat/server/global"
	"wutool.cn/chat/server/utils"
)

func MessageList(ctx *gin.Context) {
	db := global.DB
	var messageList []chat.RecvMessage
	db.Table("message").Select("content", "id", "send_id", "recv_id").Find(&messageList)
	utils.OkWithData(messageList, ctx)
}
