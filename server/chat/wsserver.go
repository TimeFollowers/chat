package chat

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"wutool.cn/chat/server/module"
	"wutool.cn/chat/server/module/entity"
)

// ClientManager is a websocket manager
type ClientManager struct {
	Clients    map[*Client]bool //客户端
	Broadcast  chan []byte      // 广播
	Register   chan *Client     // 注册
	Unregister chan *Client     // 注销
}

// Client is a websocket client
type Client struct {
	ID     string          // 客户端id
	Socket *websocket.Conn // 链接
	Send   chan []byte     //发送
}

type RecvMessage struct {
	SendId  int64  `json:"sendId,omitempty"`  //发送者
	RecvId  int64  `json:"recvId,omitempty"`  // 收件人
	Content string `json:"content,omitempty"` // 发送内容
}

var Manager = ClientManager{
	Broadcast:  make(chan []byte),      // 广播通道
	Register:   make(chan *Client),     // 注册
	Unregister: make(chan *Client),     // 注销
	Clients:    make(map[*Client]bool), // 链接客户端
}

// 启动websocket服务
func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			jsonMessage, _ := json.Marshal(&RecvMessage{Content: "/A new socket is has connected"})
			manager.Send(jsonMessage, conn)
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&RecvMessage{Content: "/A socket has disconnected"})
				manager.Send(jsonMessage, conn)
			}
		case message := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(manager.Clients, conn)
				}
			}
		}

	}
}

// 向websocket的管道chan写入数据
func (manager *ClientManager) Send(message []byte, ignore *Client) {
	for conn := range manager.Clients {
		if conn != ignore {
			conn.Send <- message
		}
	}
}

// 读取在websocket管道中的数据
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			c.Socket.Close()
			break
		}
		var recvMessage RecvMessage
		err = json.Unmarshal(message, &recvMessage)
		if err != nil {
			fmt.Println("消息接收失败:" + err.Error())
			return
		}
		c.SaveMessage(&recvMessage)
		fmt.Println(recvMessage)
		fmt.Println("收到消息:" + recvMessage.Content)
		Manager.Broadcast <- []byte(recvMessage.Content)
	}
}

// 通过websocket协议向链接到ws的客户端发送数据
func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c *Client) SaveMessage(r *RecvMessage) {
	var message entity.Message

	message.SendId = r.SendId
	message.RecvId = r.RecvId
	message.Content = r.Content
	module.CreateMessage(&message)
	fmt.Println("message保存成功，Id: %d", message.Id)
}
