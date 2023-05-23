package service

import (
	"github.com/gin-gonic/gin"
	"wutool.cn/chat/server/global"
	"wutool.cn/chat/server/utils"
)

type UserReponse struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
}

func List(ctx *gin.Context) {
	db := global.DB
	var userlist []UserReponse
	db.Table("user").Select("user_name", "id").Find(&userlist)
	utils.OkWithData(userlist, ctx)
}
