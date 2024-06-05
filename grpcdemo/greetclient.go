package grpcdemo

import (
	"context"
	"fmt"

	pb "github.com/hsit18/golang-practice/grpcdemo/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ExecuteGreetClient() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Failed to ")
	}

	client := pb.NewGreetServiceClient(conn)

	res, err := client.Greet(context.Background(), &pb.GreetRequest{Name: "greet"})
	fmt.Println(res)
}
