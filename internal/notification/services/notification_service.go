package services

import (
	"workspace-go/internal/notification/rabbitmq"
)

func StartListening() {
	queueName := "messages_queue"
	go rabbitmq.GetMessage(queueName)
}
