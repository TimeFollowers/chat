package entity

type User struct {
	Id         int64  `json:"id" gorm:"column:id"`
	UserName   string `json:"user_name" gorm:"column:user_name"`
	Password   string `json:"password" gorm:"column:password"`
	RememberMe string `json:"remember_me" gorm:"column:remember_me"` //记住我
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
	DeleteTime int64  `json:"delete_time" gorm:"column:delete_time"`
}

func (User) TableName() string {
	return "user"
}
