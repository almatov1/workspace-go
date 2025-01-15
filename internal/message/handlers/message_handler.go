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

func (h *MessageHandler) InsertMessage(ctx context.Context, req *pb.InsertMessageRequest) (*pb.InsertMessageResponse, error) {
	return h.MessageService.InsertMessage(ctx, req)
}