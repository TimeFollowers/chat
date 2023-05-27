package module

import (
	"log"
	"time"

	"wutool.cn/chat/server/global"
)

type UserRoom struct {
	Id         int   `json:"id" gorm:"column:id"`
	UserId     int   `json:"user_id" gorm:"column:user_id"`
	RoomId     int   `json:"room_id" gorm:"column:room_id"`
	RoomType   int   `json:"room_type" gorm:"column:room_type"` // 房间类型 [1-单独聊房间 2-群聊房间]
	CreateTime int64 `json:"create_time" gorm:"column:create_time"`
	UpdateTime int64 `json:"update_time" gorm:"column:update_time"`
	DeleteTime int64 `json:"delete_time" gorm:"column:delete_time"`
}

func (UserRoom) TableName() string {

	return "user_room"
}

func GetUserRoomByUserIdRoomId(userId, roomId int) *UserRoom {
	ur := new(UserRoom)
	global.DB.Select("id", "user_id", "room_id", "room_type").Where("user_id = ? and room_id = ?", userId, roomId).First(ur)
	return ur
}

func GetUserRoomByRoomId(roomId int) []*UserRoom {
	urs := make([]*UserRoom, 0)
	global.DB.Select("id", "room_id", "room_type").Where("room_id = ?", roomId).Find(urs)
	return urs
}

// IsFriend 判断两个用户是否是好友
func IsFriend(userId, userId2 int) bool {
	// 获取两个userId的单聊房间列表

	log.Printf("userId:%d", userId)
	roomIds := make([]int, 0)
	urs := make([]UserRoom, 0)
	if err := global.DB.Select("room_id").Where("user_id = ?", userId).Find(&urs).Error; err != nil {
		return false
	}
	log.Printf("查询用户房间列表\n")
	for _, v := range urs {
		roomIds = append(roomIds, v.RoomId)
	}

	// 获取userId2的单聊房间列表
	urs2 := make([]*UserRoom, 0)
	var count int64
	if err := global.DB.Select("id").Where("user_id = ? and romm_id in (?) and room_type = 1", userId2, roomIds).Find(urs2).Count(&count).Error; err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

func InsertOneUserRoom(ur *UserRoom) {
	ur.CreateTime = time.Now().Unix()
	ur.UpdateTime = time.Now().Unix()
	ur.DeleteTime = 0
	global.DB.Create(ur)
}

func GetUserRoomId(userId, userId2 int) int {
	roomIds := make([]int, 0)
	urs := make([]*UserRoom, 0)
	global.DB.Select("room_id").Where("user_id = ?", userId).Find(urs)
	for _, v := range urs {
		roomIds = append(roomIds, v.RoomId)
	}

	// 获取userId2的单聊房间列表
	ur := new(UserRoom)
	global.DB.Select("id").Where("user_id = ? and romm_id in (?) and room_type = 1", userId2, roomIds).First(ur)

	return ur.RoomId
}

func DeleteUserRoom(roomId int) {
	global.DB.Where("room_id = ?", roomId).Update("delete_time", time.Now().Unix())
}

func GetUserRoomByUserId(userId int) ([]UserRoom, error) {
	urs := make([]UserRoom, 0)
	if err := global.DB.Select("id", "room_id", "user_id", "room_type").Where("user_id = ?", userId).Find(&urs).Error; err != nil {
		return nil, err
	}
	return urs, nil
}
