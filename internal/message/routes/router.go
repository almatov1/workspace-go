package routes

import (
	"log"
	"net"

	"workspace-go/internal/message/handlers"
	pb "workspace-go/internal/message/proto"

	"google.golang.org/grpc"
)

func RunGRPCServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessageServiceServer(grpcServer, &handlers.MessageHandler{})

	log.Println("gRPC-сервер запущен на :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}