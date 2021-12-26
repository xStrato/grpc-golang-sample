package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	addUsers(client)
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
func addUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{Id: "1", Name: "Gilvan1", Email: "gilvan1@email.com"},
		{Id: "2", Name: "Gilvan2", Email: "gilvan2@email.com"},
		{Id: "3", Name: "Gilvan3", Email: "gilvan3@email.com"},
	}

	stream, err := client.AddStream(context.Background())
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}
	fmt.Println(res)
}
