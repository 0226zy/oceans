package main

import (
	context "context"
	pb "gin-blog/protocol/services/hello_world"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

const (
	addr = ":50051"
)

type server struct {
	pb.UnimplementedHelloWorldServer
}

func (s *server) Hi(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	log.Printf("Revceived:%v", req.GetName())
	return &pb.HelloResp{Msg: "Hello" + req.GetName() + " from grpc-go"}, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server:%v", err)
	}
}
