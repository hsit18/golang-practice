package grpcdemo

import (
	pb "cards/grpcdemo/generated"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type UserService struct {
	pb.UserServiceServer
}

func Execute() {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9090))

	if err != nil {
		log.Fatalf("Failed to connect TCP %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &UserService{})
	fmt.Printf("GRPC server started at %v", listen.Addr().String())
	grpcErr := grpcServer.Serve(listen)
	if grpcErr != nil {
		log.Fatalf("Failed to start GRPC server %v", err)
	}

}

func (u *UserService) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.CreateUserResponse, error) {

	return &pb.CreateUserResponse{
		ID:        1,
		Firstname: "Hello",
		Lastname:  "Singh",
		Email:     "trest@adsad.com",
		CreatedAt: "22-22-22",
	}, nil

}
