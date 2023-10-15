package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"service/chat"
	"time"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("not start service listen 9000 : %v", err)
	}

	defer conn.Close()
	c := chat.NewChatServiceClient(conn)

	//response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
	//if err != nil {
	//	log.Fatal("err when calling sayhello: %v", err)
	//}

	stream, err := c.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// send value to serve
	//go func() {
	//	for {
	if err := stream.Send(&chat.Message{Body: "hi"}); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
	//	}
	//}()

	// get value from serve
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetBody())
	}
}
