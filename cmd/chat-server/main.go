package main

import (
	"mangahub/internal/chat"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/chat", chat.ChatHandler)

	r.Run(":9093")
}
