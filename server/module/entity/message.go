package entity

type Message struct {
	Id         int64  `json:"id" gorm:"column:id"`
	SendId     int64  `json:"send_id" gorm:"column:send_id"`
	RecvId     int64  `json:"recv_id" gorm:"column:recv_id"`
	Content    string `json:"content" gorm:"column:content"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
	DeleteTime int64  `json:"delete_time" gorm:"column:delete_time"`
}

func (Message) TableName() string {
	return "message"
}
