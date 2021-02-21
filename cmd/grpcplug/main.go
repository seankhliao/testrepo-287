package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"go.seankhliao.com/testrepo-287/internal/grpcplug"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println(exec.Command("./a", "127.0.0.1:8888").Start())
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	client := grpcplug.NewPlugClient(conn)

	res, err := client.SayHello(context.TODO(), &grpcplug.HelloReq{
		N: "a",
	})
	fmt.Println(res.Msg, err)
}
