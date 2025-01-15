package main

import (
	"log"
	"workspace-go/internal/message/database"
	"workspace-go/internal/message/rabbitmq"
	"workspace-go/internal/message/server"
	"workspace-go/internal/message/websocket"
)

func main() {
	// rabbitmq
	err := rabbitmq.Initialize()
	if err != nil {
		log.Fatal("Ошибка инициализации RabbitMQ:", err)
	}
	defer rabbitmq.Close()

	// mongo
	database.ConnectMongoDB()

	// websocket
	go websocket.RunWebSocketServer()

	// gRPC
	server.RunGRPCServer()
}