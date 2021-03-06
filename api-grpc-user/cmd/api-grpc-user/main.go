package main

import (
	"api-grpc-user/internal/client"
	"api-grpc-user/internal/service"
	"api-grpc-user/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	githubClient := client.CreateGithubClient()
	userService := service.NewUserService(githubClient)

	baseServer := grpc.NewServer()
	proto.RegisterUserServiceServer(baseServer, userService)
	reflection.Register(baseServer)

	port := os.Getenv("port")
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := baseServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}
}
