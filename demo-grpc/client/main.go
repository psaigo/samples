package main

import (
	"context"
	"fmt"

	pb "github.com/samples/demo-grpc/pb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("dial error: ", err)
		return
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	// Calling Server
	text := &pb.Input{Text: "hell world"}
	resp, err := client.EchoEcho(context.Background(), text)
	if err != nil {
		fmt.Println("client call error: ", err)
		return
	}
	fmt.Println("response is:", resp.Text)

}
