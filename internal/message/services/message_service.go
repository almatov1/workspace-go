package services

import (
	"log"
	"os"
	pb "workspace-go/internal/message/proto"

	"github.com/joho/godotenv"
)

type MessageService struct{}

func (s *MessageService) GetMessage(req *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	// env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
	hello := os.Getenv("HELLO")

	response := &pb.GetMessageResponse{
		Message: "Сообщение с ID: " + req.Id + ", env: " + hello,
	}

	return response, nil
}