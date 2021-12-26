package main

import (
	"log"
	"net"

	"github.com/xStrato/grpc-golang-sample/pb"
	"github.com/xStrato/grpc-golang-sample/services"
	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
	log.Println("server is listening on port: 3030")
}
