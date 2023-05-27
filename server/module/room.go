package module

import (
	"time"

	"wutool.cn/chat/server/global"
)

type Room struct {
	Id         int    `json:"id" gorm:"column:id"`
	Name       string `json:"name" gorm:"column:name"`
	Info       string `json:"info" gorm:"column:info"`
	UserId     int    `json:"user_id" gorm:"column:user_id"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
	DeleteTime int64  `json:"delete_time" gorm:"column:delete_time"`
}

func (Room) TableName() string {
	return "room"
}

func InsertOneRoom(r *Room) {
	r.Name = "创建房间"
	r.Info = "创建房间"
	r.CreateTime = time.Now().Unix()
	global.DB.Create(r)
}

func DeleteRoom(roomId int) {
	global.DB.Where("room_id = ?", roomId).Update("delete_time", time.Now().Unix())
}
