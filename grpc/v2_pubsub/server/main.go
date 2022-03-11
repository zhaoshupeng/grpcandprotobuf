package main

import (
	"context"
	"github.com/docker/docker/pkg/pubsub"
	hello "golearning/grpc/v1/pbgo/hello"
	pb "golearning/grpc/v2_pubsub/pbgo/pubsubservice"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)
// 实现发布方法和订阅方法：
type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(
	ctx context.Context, arg *hello.String,
) (*hello.String, error) {
	p.pub.Publish(arg.GetValue())
	//debug
	//reply := &String{Value: "<Publish>  " + arg.GetValue()}
	//fmt.Println(reply.GetValue())
	return &hello.String{}, nil
}

func (p *PubsubService) Subscribe(
	arg *hello.String, stream pb.PubsubService_SubscribeServer,
) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			//debug
			//fmt.Printf("<debug> %t %s %s %t\n",
			//	ok,arg.GetValue(),key,strings.HasPrefix(key,arg.GetValue()))
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&hello.String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterPubsubServiceServer(grpcServer, NewPubsubService())

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(lis)
}