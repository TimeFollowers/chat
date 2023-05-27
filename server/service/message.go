package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"wutool.cn/chat/server/module"
	"wutool.cn/chat/server/utils"
)

func MessageList(ctx *gin.Context) {
	roomIdstr := ctx.Query("room_id")
	roomId, err := strconv.Atoi(roomIdstr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "房间号错误",
		})
		return
	}

	//判断用户是否属于该房间
	uc := ctx.MustGet("user_claims").(*utils.UserClaims)
	ur := new(module.UserRoom)
	ur = module.GetUserRoomByUserIdRoomId(uc.Id, roomId)
	if ur == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "非法访问",
		})
		return
	}

	pageIndex, _ := strconv.ParseInt(ctx.Query("page_index"), 10, 32)
	pageSize, _ := strconv.ParseInt(ctx.Query("page_size"), 10, 32)
	skip := (pageIndex - 1) * pageSize
	urs := make([]*module.Message, 0)
	urs, _ = module.GetMessageListByRoomId(roomId, pageSize, skip)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": gin.H{
			"list": urs,
		},
	})

}
