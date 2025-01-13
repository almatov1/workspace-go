package main

import (
	"workspace-go/internal/message/database"
	"workspace-go/internal/message/server"
)

func main() {
	database.ConnectMongoDB()
	server.RunGRPCServer()
}