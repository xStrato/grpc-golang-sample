package services

import (
	"context"

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
