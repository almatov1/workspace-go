package handlers

import (
	"context"
	pb "workspace-go/internal/message/proto"
	"workspace-go/internal/message/services"
)

type MessageHandler struct {
	pb.UnimplementedMessageServiceServer 
	InsertMessageService *services.InsertMessageImpl
	GetMessageService *services.GetMessageImpl
}

func (h *MessageHandler) InsertMessage(ctx context.Context, req *pb.InsertMessageRequest) (*pb.InsertMessageResponse, error) {
	return h.InsertMessageService.InsertMessage(ctx, req)
}

func (h *MessageHandler) GetMessage(ctx context.Context, req *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	return h.GetMessageService.GetMessage(ctx, req)
}