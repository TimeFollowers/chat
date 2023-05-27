package module

import (
	"fmt"
	"time"

	"wutool.cn/chat/server/global"
)

type Message struct {
	Id         int    `json:"id" gorm:"column:id"`
	SendId     int    `json:"send_id" gorm:"column:send_id"`
	RoomId     int    `json:"room_id" gorm:"column:room_id"`
	Content    string `json:"content" gorm:"column:content"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
	DeleteTime int64  `json:"delete_time" gorm:"column:delete_time"`
}

func (Message) TableName() string {
	return "message"
}
func CreateMessage(m *Message) {
	m.CreateTime = time.Now().Unix()
	m.UpdateTime = time.Now().Unix()
	m.DeleteTime = 0
	result := global.DB.Create(m)
	err := result.Error
	if err != nil {
		fmt.Println("消息保存失败")
		return
	}
}
func GetMessageListByRoomId(roomId int, limt, skip int64) ([]*Message, error) {
	data := make([]*Message, 0)
	global.DB.Select("room_id = ? and offset = ? and limit = ?", roomId, skip, limt).Find(data)
	return data, nil
}
