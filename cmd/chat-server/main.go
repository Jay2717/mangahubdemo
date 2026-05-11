package main

import (
	"github.com/gin-gonic/gin"
	"mangahub/internal/chat"
)

func main() {
	r := gin.Default()

	r.GET("/chat", chat.ChatHandler)

	r.Run(":9093")
}