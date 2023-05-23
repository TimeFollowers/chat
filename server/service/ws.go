package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"wutool.cn/chat/server/chat"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 建立websocket链接
func Ws(ctx *gin.Context) {

	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Printf("升级发生错误")
		fmt.Println(err)
		// http.NotFound(ctx.Writer, ctx.Request)
		ctx.JSON(http.StatusExpectationFailed, gin.H{"msg": "链接失败"})
		return
	}
	id, err := uuid.NewV4()
	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{"msg": "id生成失败"})
		return
	}

	// websocket connect
	client := chat.Client{ID: id.String(), Socket: conn, Send: make(chan []byte)}
	fmt.Printf(client.ID)
	chat.Manager.Register <- &client
	fmt.Println("调试测试")
	go client.Read()
	go client.Write()
}
