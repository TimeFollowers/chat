package module

import (
	"fmt"
	"time"

	"wutool.cn/chat/server/global"
	"wutool.cn/chat/server/module/entity"
)

func CreateUser(u *entity.User) {

	fmt.Println("createUser")
	u.CreateTime = time.Now().Unix()
	u.UpdateTime = time.Now().Unix()
	u.DeleteTime = 0
	result := global.DB.Create(u)
	err := result.Error
	if err != nil {
		fmt.Println("创建用户失败")
	}
}

func SelectUser() *[]entity.User {
	db := global.DB
	var userlist []entity.User
	db.Select("user_name", "id").Find(&userlist)
	return &userlist
}
