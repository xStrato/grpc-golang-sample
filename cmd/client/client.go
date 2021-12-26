package main

import (
	"context"
	"fmt"
	"io"
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
	addUserVerbose(client)
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

func addUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "123",
		Name:  "Gilvan",
		Email: "gilvan@gmail.com",
	}

	res, err := client.AddVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status: ", stream.GetStatus(), " - ", stream.GetUser())
	}
}
