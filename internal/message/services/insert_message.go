package services

import (
	"context"
	"log"
	"os"
	configs "workspace-go/configs/message"
	"workspace-go/internal/message/database"
	"workspace-go/internal/message/models"
	pb "workspace-go/internal/message/proto"
	"workspace-go/internal/message/rabbitmq"
	"workspace-go/internal/message/redis"
	"workspace-go/internal/message/websocket"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InsertMessageImpl struct{}

func (s *InsertMessageImpl) InsertMessage(ctx context.Context, req *pb.InsertMessageRequest) (*pb.InsertMessageResponse, error) {	
	// write record
	message := models.Message{
		Message: req.Message,
	}

	collection := database.Client.Database(configs.DBName).Collection(configs.MessageCollectionName)
    result, err := collection.InsertOne(ctx, message)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Не удалось вставить сообщение в базу данных: %v", err)
    }

	// send message to broadcast
	websocket.Broadcast <- req.Message

	// Send message to RabbitMQ
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
	queueName := os.Getenv("RABBITMQ_QUEUE")

	err = rabbitmq.SendMessage(queueName, req.Message)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Не удалось отправить сообщение в RabbitMQ: %v", err)
	}

	// reset redis
	err = redis.RedisClient.Del(ctx, "first_page").Err()
	if err != nil {
		log.Printf("Ошибка при удалении кэша Redis: %v", err)
	}

    insertedID := result.InsertedID.(primitive.ObjectID)
    return &pb.InsertMessageResponse{
		Id: insertedID.Hex(),
        Message: req.Message,
    }, nil
}