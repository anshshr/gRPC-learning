package main

import (
	"context"
	"fmt"
	proto "learning_grpc_go/unary/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	fmt.Println("we are building th client part here")

	conn, err := grpc.NewClient("localhost:9000"  , grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil{
		panic(err)
	}

	client := proto.NewExampleClient(conn)

	req := &proto.HelloRequest{SomeString: "hello from  the client side this is ansh shrivastav learning the backend application"}

	client.ServerReply(context.TODO() , req)
}