package module

import (
	"fmt"
	"time"

	"wutool.cn/chat/server/global"
	"wutool.cn/chat/server/module/entity"
)

func CreateMessage(m *entity.Message) {

	fmt.Println("createMessage")
	m.CreateTime = time.Now().Unix()
	m.UpdateTime = time.Now().Unix()
	m.DeleteTime = 0
	result := global.DB.Create(m)
	err := result.Error
	if err != nil {
		fmt.Println("消息保存失败")
	}
}
