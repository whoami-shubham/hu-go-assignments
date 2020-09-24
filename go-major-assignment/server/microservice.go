package main

import (
	context "context"
	email "go-major-assignment/email"
	"go-major-assignment/model"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type server struct{}

func (s *server) Report(ctx context.Context, req *email.ReportRequest) (*email.ReportResponse, error) {
	go SendMail(req.Email, req.Report)
	return &email.ReportResponse{Res: "Ok"}, nil
}
func (s *server) Update(ctx context.Context, req *email.UpdateRequest) (*email.UpdateResponse, error) {
	go SendMail(req.Email, req.Title)
	return &email.UpdateResponse{Res: "ok"}, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("server started ....")

	s := grpc.NewServer()
	email.RegisterEmailServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to serve: %v\n", err)
	}

}

// SendMail ...
func SendMail(email string, message interface{}) {
	log.Println("Sending Mail to : ", email)
	log.Println("msg : ", string(model.GetSampleJSONData(message)))
}
