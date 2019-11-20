package main

import (
	"context"
	pb "gin-blog/protocol/services/hello_world"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	//address = "localhost:50051"
	address = "106.52.111.96:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}

	defer conn.Close()
	c := pb.NewHelloWorldClient(conn)
	name := "zhangyong"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Hi(ctx, &pb.HelloReq{Name: name})
	if err != nil {
		log.Fatalf("could not hi:%v", err)
	}

	log.Printf("hi:%s", r.GetMsg())

}
