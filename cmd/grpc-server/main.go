package main

import (
	grpcserver "mangahub/internal/grpc"
	"mangahub/pkg/database"

	_ "modernc.org/sqlite"
)

func main() {
	database.Init()
	grpcserver.StartGRPCServer()
}
