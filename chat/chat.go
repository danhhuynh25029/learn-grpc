package chat

import (
	"context"
	"fmt"
	"io"
	"time"
)

type Server struct {
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {
	//TODO implement me
	fmt.Println("implement me")
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	fmt.Println("Receive message : %v", in.Body)
	return &Message{Body: "hello grpc"}, nil
}

func (s *Server) Channel(stream ChatService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Println("receive from client : ", args.GetBody())
		for i := 0; i < 5; i++ {
			reply := &Message{Body: "hello: client" + time.Now().String()}
			err = stream.Send(reply)
			if err != nil {
				return err
			}
		}
	}
}
