package ws

// session_type 会话类型
const (
	ST_UnKnow = 0 // 未知
	ST_Single = 1 // 单聊
	ST_Group  = 2 // 群聊
)

// content_type 消息内容类型
const (
	CT_UnKnow = 0 // 未知
	CT_Text   = 1 //文本类型
)

// message_type 消息类型
const (
	MT_UnKnow = 0 //未知
	MT_Login  = 1 //登录
)

// 消息
type Message struct {
	Type int // 消息类型
	data []byte
}

type LoginMessage struct {
	Token string
}

type Conn struct {
}
