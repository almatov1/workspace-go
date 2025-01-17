package main

import (
	"log"
	"workspace-go/internal/message/database"
	"workspace-go/internal/message/rabbitmq"
	"workspace-go/internal/message/redis"
	"workspace-go/internal/message/server"
	"workspace-go/internal/message/websocket"
)

func main() {
	// RabbitMQ
	err := rabbitmq.Initialize()
	if err != nil {
		log.Fatal("Ошибка инициализации RabbitMQ:", err)
	}
	defer rabbitmq.Close()

	// MongoDB
	database.ConnectMongoDB()
	
	// Redis
	redis.InitRedis()
	defer redis.Close()

	// WebSocket
	go websocket.RunWebSocketServer()

	// gRPC
	server.RunGRPCServer()
}