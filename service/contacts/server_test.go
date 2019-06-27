package contacts

import (
	"context"
	"fmt"
	"testing"

	pb "grpc-middleware/grpc-gen/contacts/proto"

	"google.golang.org/grpc"
)

func getConnDev() *grpc.ClientConn {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Can't connect to server", err)
	}

	return conn
}

func Test_Validate(t *testing.T) {
	client := pb.NewContactsServiceClient(getConnDev())

	_, err := client.Validate(context.Background(), &pb.Request{
		Contact: &pb.Contact{
			UserID: "anvh2",
			Name:   "Hoang An",
			Phone:  "0921412532",
			Status: pb.Status_VALID,
		},
	})

	if err != nil {
		fmt.Println(err)
	}
}

func Test_Create(t *testing.T) {

}

func Test_GetContact(t *testing.T) {

}
