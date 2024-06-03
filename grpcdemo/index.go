package grpcdemo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/hsit18/golang-practice/grpcdemo/generated"

	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

type User struct {
	ID        int32     `gorm:"primarykey"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:false"`
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

func (u *UserService) GetUsers(ctx context.Context, req *pb.NoParams) (*pb.GetUsersResponse, error) {
	user := []*pb.User{}

	DB.Find(&user)

	return &pb.GetUsersResponse{
		Users: user,
	}, nil
}

func (u *UserService) GetUser(ctx context.Context, req *pb.UserRequestParam) (*pb.CreateUserResponse, error) {
	var user User

	res := DB.Find(&user, "id = ?", req.GetID())

	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			ID:        int32(user.ID),
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
		},
	}, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *pb.UserRequestParam) (*pb.DeleteUserResponse, error) {
	var user User

	res := DB.Where("id = ?", req.GetID()).Delete(&user)

	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}

func (u *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
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
