package main

import (
	grpcserver "mangahub/internal/grpc"
	"mangahub/pkg/database"
)

func main() {
	database.Init()
	grpcserver.StartGRPCServer()
}