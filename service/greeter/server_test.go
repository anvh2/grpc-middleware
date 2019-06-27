package greeter

import (
	"context"
	"fmt"
	pb "grpc-middleware/grpc-gen/proto"
	"testing"

	"google.golang.org/grpc"
)

func getConnDev() *grpc.ClientConn {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Can't connect to server", err)
	}

	return conn
}

func Test_Say(t *testing.T) {
	client := pb.NewGreeterServiceClient(getConnDev())

	_, err := client.Say(context.Background(), &pb.Request{
		Name: "anvh2",
	})

	if err != nil {
		fmt.Println(err)
	}
}
