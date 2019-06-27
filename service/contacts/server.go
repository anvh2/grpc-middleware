package contacts

import (
	"context"
	"fmt"
	contacts "grpc-middleware/grpc-gen/contacts/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

//Server ...
type Server struct{}

//NewServer ...
func NewServer() *Server {
	return &Server{}
}

//Run ...
func (s *Server) Run() error {
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()

	contacts.RegisterContactsServiceServer(grpcServer, s)

	fmt.Println("Server is runing on port 8000")

	return grpcServer.Serve(lis)
}

//Validate ...
func (s *Server) Validate(ctx context.Context, req *contacts.Request) (*contacts.ValidateResponse, error) {
	if req.Contact.GetName() == "" {

	}

	return &contacts.ValidateResponse{
		ErrorCode:   0,
		MessageCode: "OK",
		Status:      contacts.Status_VALID,
	}, nil
}

//Create ...
func (s *Server) Create(ctx context.Context, req *contacts.Request) (*contacts.Response, error) {
	return &contacts.Response{
		ErrorCode:   0,
		MessageCode: "OK",
	}, nil
}

//GetContact ...
func (s *Server) GetContact(ctx context.Context, req *contacts.GetContactRequest) (*contacts.GetContactResponse, error) {
	contact := new(contacts.Contact)

	contact = &contacts.Contact{
		UserID: "",
		Name:   "",
		Phone:  "",
		Status: contacts.Status_VALID,
	}

	return &contacts.GetContactResponse{
		ErrorCode:   0,
		MessageCode: "OK",
		Contact:     contact,
	}, nil
}
