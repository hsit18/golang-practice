package grpcdemo

import (
	pb "cards/grpcdemo/generated"
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

var DB *gorm.DB
var err error

func InitialMigration() {
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Error conecting DB...")
	}
	DB.AutoMigrate(&User{})
}

func Execute() {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9090))

	if err != nil {
		log.Fatalf("Failed to connect TCP %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &UserService{})
	fmt.Printf("GRPC server started at %v", listen.Addr().String())
	InitialMigration()
	grpcErr := grpcServer.Serve(listen)
	if grpcErr != nil {
		log.Fatalf("Failed to start GRPC server %v", err)
	}

}

func (u *UserService) CreateUser(ctx context.Context, req *pb.CreeateUserRequest) (*pb.CreateUserResponse, error) {

	fmt.Println(req)

	user := req.GetUser()
	data := User{
		Firstname: user.GetFirstname(),
		Lastname:  user.GetLastname(),
		Email:     user.GetEmail(),
	}

	fmt.Printf("%+v\n", data)

	res := DB.Create(&data)

	fmt.Printf("AFTER DB insert: %+v\n", data)

	if res.RowsAffected == 0 {
		return nil, errors.New("failed to create user")
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			ID:        int32(data.ID),
			Firstname: data.Firstname,
			Lastname:  data.Lastname,
			Email:     data.Email,
		},
	}, nil

}
