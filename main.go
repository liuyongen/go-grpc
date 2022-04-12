package main

import (
	"context"
	"flag"
	"fmt"
	pb "gogrpc/example"
	"google.golang.org/grpc"
	"log"
	"net"
)

/*const (
	port = ":50051"
)*/

var (
	port = flag.Int("port", 50051, "The server port")
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
	flag.Parse()
	// lis, err := net.Listen("tcp", port)
	// socket = fmt.Sprintf("172.20.131.130:%d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
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
