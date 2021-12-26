package main

import (
	"context"
	"log"

	"github.com/xStrato/grpc-golang-sample/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":3030", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	addUser(client)
}

func addUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "123",
		Name:  "Gilvan",
		Email: "gilvan@gmail.com",
	}

	res, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}
	log.Println(res)
}
