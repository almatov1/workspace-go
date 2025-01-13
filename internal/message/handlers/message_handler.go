package handlers

import (
	"context"
	pb "workspace-go/internal/message/proto"
)

type MessageHandler struct {
	pb.UnimplementedMessageServiceServer
}

func (h *MessageHandler) GetMessage(ctx context.Context, req *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	response := &pb.GetMessageResponse{
		Message: "Сообщение с ID: " + req.Id,
	}
	return response, nil
}