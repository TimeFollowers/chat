package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"wutool.cn/chat/server/global"
	"wutool.cn/chat/server/module"
	"wutool.cn/chat/server/module/entity"
	"wutool.cn/chat/server/utils"
)

type User struct {
	UserName   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember"`
}

func Login(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": "获取参数失败"})
	}
	fmt.Println(user)
	entityUser := &entity.User{}
	entityUser.UserName = user.UserName
	entityUser.Password = user.Password
	module.CreateUser(entityUser)
	fmt.Println(entityUser)
	store, err := global.SESSION.Get(ctx.Request, "go-session")
	if err != nil {
		utils.Fail(ctx)
		return
	}
	store.Values["user_name"] = entityUser.UserName
	store.Values["user_id"] = entityUser.Id
	fmt.Println(store.Values["user_id"])
	store.Save(ctx.Request, ctx.Writer)
	ctx.JSON(http.StatusOK, gin.H{"username": entityUser.UserName, "password": entityUser.Password, "remeber": user.RememberMe})
}
