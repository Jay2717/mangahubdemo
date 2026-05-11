package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{hub: hub}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // dev mode; production nên check origin
	},
}

// HandleWebSocket handles websocket connection for chat
func (h *Handler) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &Client{
		conn: conn,
		send: make(chan []byte, 256),
	}

	// register client to hub
	h.hub.register <- client

	// READ LOOP (receive message from client)
	go func() {
		defer func() {
			h.hub.unregister <- client
			conn.Close()
		}()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}

			// broadcast to all clients
			h.hub.broadcast <- msg
		}
	}()

	// WRITE LOOP (send message to client)
	for msg := range client.send {
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
