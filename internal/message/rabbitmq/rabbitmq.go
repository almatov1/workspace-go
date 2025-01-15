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
	// env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	// connect to rabbitmq
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

func SendMessage(queueName, message string) error {
	_, err := channel.QueueDeclare(
		queueName, // имя очереди
		true,      // устойчивая очередь
		false,     // очередь не удаляется после использования
		false,     // не эксклюзивная
		false,     // очередь не автоудаляется
		nil,       // дополнительные аргументы
	)
	if err != nil {
		return err
	}

	err = channel.Publish(
		"",       // обменник по умолчанию
		queueName, // имя очереди
		false,    // не обязательное
		false,    // не подтвержденное
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Сообщение отправлено в очередь %s: %s", queueName, message)
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