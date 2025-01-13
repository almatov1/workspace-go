package handlers

import (
	"context"
	pb "workspace-go/internal/message/proto"
	"workspace-go/internal/message/services"
)

type MessageHandler struct {
	pb.UnimplementedMessageServiceServer 
	MessageService *services.MessageService
}

func (h *MessageHandler) GetMessage(ctx context.Context, req *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	return h.MessageService.GetMessage(req)
}