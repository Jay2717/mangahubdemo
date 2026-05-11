# MangaHub

## Features
- JWT Authentication
- REST API
- WebSocket Chat
- TCP Reading Progress Sync
- UDP Notification
- gRPC Manga Service
- Nginx Reverse Proxy
- Load Balancing

## Run API Server ($env:PORT="8081")
go run cmd/api-server/main.go

## Run TCP Server
go run cmd/tcp-server/main.go

## Run UDP Server
go run cmd/udp-server/main.go

## Run gRPC Server
go run cmd/grpc-server/main.go

## Run Chat Server
go run cmd/chat-server/main.go

## Run gRPC Client
go run cmd/grpc-client/main.go

## Health Check
curl http://localhost:8080/health

## Login
curl -X POST http://localhost:8080/login

## Get Manga List
curl http://localhost:8080/manga
