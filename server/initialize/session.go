package initialize

import (
	"github.com/gorilla/sessions"
)

func InitSession() *sessions.CookieStore {

	// something-very-secret应该是一个你自己的密匙，只要不被别人知道就行
	return sessions.NewCookieStore([]byte("go-session-secret"))
}
