package main

import (
	"context"
	"log"

	hello "golearning/grpc/v1/pbgo/hello"
	pb "golearning/grpc/v2_pubsub/pbgo/pubsubservice"
	"google.golang.org/grpc"
)

func main() {
	//此示例中的 gRPC 的服务都没有提供证书支持，因此客户端在链接服务器中通过 grpc.WithInsecure() 选项跳过了对服务器证书的验证。没有启用证书的 gRPC 服务在和客户端进行的是明文通讯，
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewPubsubServiceClient(conn)

	_, err = client.Publish(context.Background(), &hello.String{Value: "golang: hello Go!!!"})
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(context.Background(), &hello.String{Value: "docker: hello Docker"})
	if err != nil {
		log.Fatal(err)
	}
}
