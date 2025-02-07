package main

import (
	"context"
	"fmt"
	proto "learning_grpc_go/rest_apis/protoc"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("this is the client part of the rest apis craeation")

	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)

	}
	client := proto.NewChatClient(conn)

	route := gin.Default()

	route.GET("/getdata/:message", func(ctx *gin.Context) {
		data := ctx.Param("message")

		req := proto.AnshResposne{Data: data, Age: 20, Height: 30, Marks: 100, IsResultDeclared: false}
		resp, err := client.StartChat(context.TODO(), &req)
		
		if err != nil {
			// panic(err)
			ctx.JSON(http.StatusBadRequest , gin.H{
				"error occured" : err,
			})
		}
		if resp != nil{
			ctx.JSON(http.StatusBadRequest , gin.H{
				"error occured in getting response" : resp,
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"values":  "your message is data" + data,
			"data is": resp,
		})
	})

	route.Run(":3000")
}
