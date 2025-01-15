package main

import (
	"workspace-go/internal/message/database"
	"workspace-go/internal/message/server"
	"workspace-go/internal/message/websocket"
)

func main() {
	database.ConnectMongoDB()
	go websocket.RunWebSocketServer()
	server.RunGRPCServer()
}