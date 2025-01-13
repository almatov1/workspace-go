package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongoDB() {
    // env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	// connect to db
	user := os.Getenv("MONGO_USER")
    password := os.Getenv("MONGO_PASSWORD")
    uri := os.Getenv("MONGO_URI")
    
    var err error
    mongoURI := fmt.Sprintf("mongodb://%s:%s@%s", user, password, uri)
    clientOptions := options.Client().ApplyURI(mongoURI)
    Client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatalf("Ошибка подключения к MongoDB: %v", err)
    }

    // Проверяем подключение
    err = Client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatalf("Ошибка пинга MongoDB: %v", err)
    }
    fmt.Println("Успешно подключено к MongoDB")
}
