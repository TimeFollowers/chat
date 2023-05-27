package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"wutool.cn/chat/server/utils"
)

func AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("authorization")
		userClaims, err := utils.AnalyseToken(token)
		if err != nil {
			log.Printf("token:%s", token)
			log.Printf("err:%s", err.Error())
			log.Printf("userClaims:%v", userClaims)
			ctx.Abort()
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户验证不通过",
			})
			return
		}

		ctx.Set("user_claims", userClaims)
		ctx.Next()
	}
}
