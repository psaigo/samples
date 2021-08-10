package main

import (
	"fmt"
	"net"

	pb "github.com/psaigo/samples/demo-grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("listening...")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("could not listen:", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type server struct {
	pb.UnimplementedEchoServer
}

func (server) DemoHello(ctx context.Context, inp *pb.Input) (*pb.Output, error) {

	o := pb.Output{Text: inp.Text}
	return &o, nil
}
