package main

import (
	"context"
	"errors"
	"fmt"
	proto "learning_grpc_go/unary/protoc"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type exampleServer struct {
	proto.UnimplementedExampleServer
}

func main() {
	fmt.Println("we are creating the server of the grpc using the golang ")

	listen , tcpError := net.Listen("tcp" , ":9000") 

	if tcpError != nil{
		panic(tcpError)
	}

	srv := grpc.NewServer()
	proto.RegisterExampleServer(srv , &exampleServer{})

	reflection.Register(srv)

	e := srv.Serve(listen)

	if e != nil{
		panic(e)
	}

}


func (s *exampleServer) ServerReply(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error){
	fmt.Println("receive reqiest from the client" , req.SomeString)
	fmt.Println("hello from thr server")
	return &proto.HelloResponse{} , errors.New("")
}