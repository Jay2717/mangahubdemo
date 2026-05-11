package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var hub = NewHub()

func init() {
	go hub.Run()
}

func ChatHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &Client{
		conn: conn,
		send: make(chan []byte),
	}

	hub.register <- client

	// read
	go func() {
		defer func() {
			hub.unregister <- client
			conn.Close()
		}()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			hub.broadcast <- msg
		}
	}()

	// write
	for msg := range client.send {
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}