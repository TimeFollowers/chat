package module

import (
	"fmt"
	"log"
	"time"

	"wutool.cn/chat/server/global"
)

type User struct {
	Id         int    `json:"id" gorm:"column:id"`
	UserName   string `json:"user_name" gorm:"column:user_name"`
	Password   string `json:"password" gorm:"column:password"`
	Email      string `json:"email" gorm:"column:email"`
	Avatar     string `json:"avatar" gorm:"column:avatar"`
	Sex        int    `json:"sex" gorm:"column:sex"`
	RememberMe string `json:"remember_me" gorm:"column:remember_me"` //记住我
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
	DeleteTime int64  `json:"delete_time" gorm:"column:delete_time"`
}

func (User) TableName() string {
	return "user"
}
func CreateUser(u *User) {
	u.CreateTime = time.Now().Unix()
	u.UpdateTime = time.Now().Unix()
	u.DeleteTime = 0
	result := global.DB.Create(u)
	err := result.Error
	if err != nil {
		fmt.Println("创建用户失败")
	}
}

func GetUserByUsernamePassword(userName, password string) (*User, error) {
	u := new(User)
	if err := global.DB.Select("id", "user_name", "email", "avatar", "sex").Where("user_name = ? and password = ?", userName, password).First(u).Error; err != nil {
		log.Printf("GetUserByUsernamePassword,user:%v \n", u)
		return nil, err
	}
	return u, nil
}

func GetUserById(id int) (*User, error) {
	u := new(User)
	if err := global.DB.Select("id", "user_name", "email", "avatar", "sex").Where("id = ?", id).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByUserName(username string) (*User, error) {
	u := new(User)
	if err := global.DB.Select("id", "user_name", "email", "avatar", "sex").Where("user_name = ?", username).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserCountByUserName(username string) (int64, error) {
	var count int64
	if err := global.DB.Model(&User{}).Select("id", "user_name", "email", "avatar", "sex").Where("user_name = ?", username).Count(&count).Error; err != nil {

		return 0, err
	}
	return count, nil
}

func GetUserByUserEmail(email string) (*User, error) {
	u := new(User)
	global.DB.Select("id", "user_name", "email", "avatar", "sex").Where("user_name = ?", email).First(u)
	return u, nil
}

func GetUserCountByUserEmail(email string) (int64, error) {
	u := new(User)
	var count int64
	global.DB.Select("id", "user_name", "email", "avatar", "sex").Where("user_name = ?", email).First(u).Count(&count)
	return count, nil
}

func GetUserFirend(userId int) ([]User, error) {
	roomIds := make([]int, 0)
	urs := make([]UserRoom, 0)
	urs, err := GetUserRoomByUserId(userId)
	if err != nil {
		return nil, err
	}
	for _, ur := range urs {
		roomIds = append(roomIds, ur.RoomId)
	}
	urs2 := make([]UserRoom, 0)
	if err := global.DB.Select("user_id").Where("room_id in (?) and user_id != ?", roomIds, userId).Find(&urs2).Error; err != nil {
		return nil, err
	}
	uIds := make([]int, 0)
	for uId := range uIds {
		uIds = append(uIds, uId)
	}
	users := make([]User, 0)
	if err := global.DB.Select("id", "user_name", "email", "avatar", "sex").Where("id in (?)", uIds).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserAll() ([]User, error) {
	users := make([]User, 0)
	if err := global.DB.Select("id", "user_name", "email", "avatar", "sex").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
