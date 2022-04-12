package main

import (
	"context"
	pb "gogrpc/example"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedDemoServer
}

func NewServer() pb.DemoServer {
	return &server{}
}

func (s *server) GetDemo(ctx context.Context, in *pb.GetDemoReq) (*pb.GetDemoReply, error) {
	log.Printf("Received: %d", in.UserId)
	return &pb.GetDemoReply{UserId: in.UserId}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//pb.RegisterDemoServer(s, &server{})
	pb.RegisterDemoServer(s, NewServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
