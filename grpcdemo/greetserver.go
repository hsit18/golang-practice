package grpcdemo

import (
	"context"

	pb "github.com/hsit18/golang-practice/grpcdemo/generated"
)

type GreetServer struct {
	pb.GreetServiceServer
}

func (s *GreetServer) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{Message: "Hello " + req.Name}, nil
}
