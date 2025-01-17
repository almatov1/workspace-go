package services

import (
	"context"
	"encoding/json"
	"log"
	"workspace-go/configs"
	"workspace-go/internal/message/database"
	"workspace-go/internal/message/models"
	pb "workspace-go/internal/message/proto"
	"workspace-go/internal/message/redis"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetMessageImpl struct{}

func (s *GetMessageImpl) GetMessage(ctx context.Context, req *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	// check redis
	if req.Page == 1 {
		cachedData, err := redis.RedisClient.Get(ctx, "first_page").Result()
		if err == nil && cachedData != "" {
			var cachedMessages []*pb.Message
			if err := json.Unmarshal([]byte(cachedData), &cachedMessages); err != nil {
				log.Printf("Ошибка при распаковке данных из Redis: %v", err)
			}
			return &pb.GetMessageResponse{Messages: cachedMessages}, nil
		}
	}

	// pagination
	pageSize := req.PageSize
	pageNumber := req.Page

	skip := int64((pageNumber - 1)) * int64(pageSize)
	limit := int64(pageSize)

	collection := database.Client.Database(configs.DBName).Collection(configs.MessageCollectionName)

	cursor, err := collection.Find(ctx, bson.D{}, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []models.Message
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	var responseMessages []*pb.Message
	for _, msg := range messages {
		responseMessages = append(responseMessages, &pb.Message{
			Id:      msg.ID.Hex(),
			Message: msg.Message,
		})
	}

	// caching first page
	if req.Page == 1 {
		cachedMessages, err := json.Marshal(responseMessages)
		if err != nil {
			log.Printf("Ошибка при сериализации данных для Redis: %v", err)
		} else {
			err = redis.RedisClient.Set(ctx, "first_page", cachedMessages, 0).Err()
			if err != nil {
				log.Printf("Ошибка при записи данных в Redis: %v", err)
			}
		}
	}

	return &pb.GetMessageResponse{
		Messages: responseMessages,
	}, nil
}