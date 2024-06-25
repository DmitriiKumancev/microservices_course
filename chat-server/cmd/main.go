package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/DmitriiKumacnev/microservices_course/chat-server/pkg/chat_server_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	desc.UnimplementedChatServerV1Server
}

func (s *server) Create(_ context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	fmt.Println("Create chat with users:", req.Usernames)
	return &desc.CreateChatResponse{Id: 1}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	fmt.Println("Delete chat with id:", req.Id)
	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	fmt.Println("Send message from:", req.From, "test:", req.Text, "timestamp:", req.Timestamp)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterChatServerV1Server(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
