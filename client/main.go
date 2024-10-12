package main

import (
	"log"

	"github.com/ryuryu5121/go-environment/grpc_sample"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	c := grpc_sample.NewSampleServiceClient(conn)

	response, err := c.GetData(context.Background(), &grpc_sample.Message{Body: "送信データ"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Print(response.Body)

	defer conn.Close()
}
