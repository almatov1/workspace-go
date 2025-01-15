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

func (s *MessageService) InsertMessage(ctx context.Context, req *pb.InsertMessageRequest) (*pb.InsertMessageResponse, error) {	
	// write record
	message := models.Message{
		Message: req.Message,
	}

	collection := database.Client.Database("messages_db").Collection("messages")
    result, err := collection.InsertOne(ctx, message)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Не удалось вставить сообщение в базу данных: %v", err)
    }

    insertedID := result.InsertedID.(primitive.ObjectID)
    return &pb.InsertMessageResponse{
		Id: insertedID.Hex(),
        Message: req.Message,
    }, nil
}