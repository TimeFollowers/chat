package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"wutool.cn/chat/server/utils"
)

type LoginMessage struct {
	token string
}

// ClientManager is a websocket manager
type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// Client is a websocket client
type Client struct {
	ID     int
	Socket *websocket.Conn
	Send   chan []byte
}

// Message is an object for websocket message which is mapped to json type
type Message struct {
	Sender    int    `json:"sender,omitempty"`
	Recipient int    `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
	Token     string `json:"token,omitempty"`
}

// Manager define a ws server manager
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

// Start is to start a ws server
func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.Send(jsonMessage, conn)
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
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

// Send is to send ws message to ws client
func (manager *ClientManager) Send(message []byte, ignore *Client) {
	for conn := range manager.Clients {
		if conn != ignore {
			conn.Send <- message
		}
	}
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, messageStr, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			c.Socket.Close()
			break
		}
		message := Message{}
		err = json.Unmarshal(messageStr, &message)
		if err != nil {
			Manager.Broadcast <- []byte("链接失败")
			continue
		}

		token := message.Token
		userClaims, err := utils.AnalyseToken(token)
		if err != nil {
			log.Printf("token:%s", token)
			log.Printf("err:%s", err.Error())
			log.Printf("userClaims:%v", userClaims)
			Manager.Broadcast <- []byte("链接失败")
			continue
		}
		Manager.Broadcast <- messageStr
	}
}

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

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

// Ws websocket
func Ws(ctx *gin.Context) {
	// change the reqest to websocket model
	// conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(ctx.Writer, ctx.Request, nil)
	conn, error := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if error != nil {
		log.Printf("error:%s", error.Error())
		http.NotFound(ctx.Writer, ctx.Request)
		return
	}

	// websocket connect
	// client := &Client{ID: 0, Socket: conn, Send: make(chan []byte)}

	// Manager.Register <- client

	go client,Read()
	go client.Write()
}

func (c *ClientManager) handleMessage(message string) {

}
