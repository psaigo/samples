package main

import (
	"context"
	"fmt"

	pb "github.com/psaigo/samples/demo-grpc/pb"
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
	text := &pb.Input{Text: "hello world"}
	resp, err := client.DemoHello(context.Background(), text)
	if err != nil {
		fmt.Println("client call error: ", err)
		return
	}
	fmt.Println("response is:", resp.Text)

}
