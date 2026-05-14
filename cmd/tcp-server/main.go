package main

import (
	"mangahub/internal/tcp"
	//"mangahub/pkg/database"
)

func main() {
	//database.Init()
	tcp.StartTCPServer()
}
