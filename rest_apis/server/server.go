package main

import (
	"context"
	"errors"
	"fmt"
	proto "learning_grpc_go/rest_apis/protoc"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedChatServer
}

func main() {
	fmt.Println("this is the server part for the rest apis section")

	listen, tcpError := net.Listen("tcp", ":9000")

	if tcpError != nil {
		panic(tcpError)
	}

	srv := grpc.NewServer()

	proto.RegisterChatServer(srv, &server{})

	reflection.Register(srv)

	e := srv.Serve(listen)

	if e != nil {
		panic(e)
	}

}

func (s *server) StartChat(c context.Context, req *proto.AnshResposne) (*proto.RonakReply, error) {
	fmt.Println("the response is starting from the here")
	fmt.Println("the name is  :", req.Data)
	fmt.Println("the age is  :", req.Age)
	fmt.Println("the height is  :", req.Height)
	fmt.Println("the marks is  :", req.Marks)
	fmt.Println("the isresult is  :", req.IsResultDeclared)

	return &proto.RonakReply{
		Height:           req.Height,
		Marks:            req.Marks,
		IsResultDeclared: req.IsResultDeclared,
		Data:             req.Data,
		Age:              req.Age,
	}, errors.New("invalid request")

}
