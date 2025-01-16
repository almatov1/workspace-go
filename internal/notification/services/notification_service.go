package services

import (
	"log"
	"os"
	"workspace-go/internal/notification/rabbitmq"

	"github.com/joho/godotenv"
)

func StartListening() {
	// env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
	queueName := os.Getenv("RABBITMQ_QUEUE")

	go rabbitmq.GetMessage(queueName)
}