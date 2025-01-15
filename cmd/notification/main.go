package main

import (
	"log"
	"workspace-go/internal/notification/rabbitmq"
	"workspace-go/internal/notification/services"
)

func main() {
	err := rabbitmq.Initialize()
	if err != nil {
		log.Fatal("Ошибка инициализации RabbitMQ:", err)
	}
	defer rabbitmq.Close()

	services.StartListening();
	select {}
}