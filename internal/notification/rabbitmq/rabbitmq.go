package rabbitmq

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel

func Initialize() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	uri := os.Getenv("RABBITMQ_URI")

	var err error
	conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s", user, password, uri))
	if err != nil {
		return err
	}

	channel, err = conn.Channel()
	if err != nil {
		return err
	}

	log.Println("RabbitMQ успешно инициализировано")
	return nil
}

func Close() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

func GetMessage(queueName string) error {
	msgs, err := channel.Consume(
		queueName, // Имя очереди
		"",        // Consumer имя (пустое)
		true,      // Автоматическое подтверждение
		false,     // Не эксклюзивная очередь
		false,     // Не автоудаляемая
		false,     // Не обязательная
		nil,       // Дополнительные аргументы
	)
	if err != nil {
        return fmt.Errorf("Ошибка подключения к очереди: %v", err)
    }

    for msg := range msgs {
        log.Printf("Получено сообщение: %s", string(msg.Body))
    }
	
	return nil
}