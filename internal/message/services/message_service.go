package services

import (
	"context"
	"workspace-go/internal/message/database"
	"workspace-go/internal/message/models"
	pb "workspace-go/internal/message/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MessageService struct{}

func (s *MessageService) GetMessage(ctx context.Context, req *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {	
	// write record
	message := models.Message{
		Content: req.Id,
	}

	collection := database.Client.Database("messages_db").Collection("messages")
    result, err := collection.InsertOne(ctx, message)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Не удалось вставить сообщение в базу данных: %v", err)
    }

    insertedID := result.InsertedID.(primitive.ObjectID)
    return &pb.GetMessageResponse{
        Message: "Сообщение с ID: " + insertedID.Hex(),
    }, nil
}