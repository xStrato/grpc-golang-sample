package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/xStrato/grpc-golang-sample/pb"
)

type userService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *userService {
	return &userService{}
}

func (*userService) Add(ctx context.Context, req *pb.User) (*pb.User, error) {
	//Insert DB
	return &pb.User{
		Id:    "12345",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*userService) AddVerbose(req *pb.User, stream pb.UserService_AddVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	user := &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	stream.Send(&pb.UserResultStream{
		Status: "user has been inserted",
		User:   user,
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User:   user,
	})
	return nil
}

func (*userService) AddStream(stream pb.UserService_AddStreamServer) error {
	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{User: users})
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})
		fmt.Println("Adding user: ", req.GetName())
	}
}
