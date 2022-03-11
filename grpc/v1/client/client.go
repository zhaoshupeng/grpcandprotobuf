package main

import (
	"context"
	"fmt"
	"golearning/grpc/v1/pbgo/hello"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

// 客户端链接 gRPC 服务了
func main() {
	// 其中 grpc.Dial 负责和 gRPC 服务建立链接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 返回的 client 其实是一个 HelloServiceClient 接口对象，通过接口定义的方法就可以调用服务端对应的 gRPC 服务提供的方法。
	client := hello.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &hello.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

	testChannel(client)
}

func testChannel(client hello.HelloServiceClient) {
	// 客户端需要先调用 Channel 方法获取返回的流对象：
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 在客户端我们将发送和接收操作放到两个独立的 Goroutine。首先是向服务端发送数据：
	go func() {
		for {
			if err := stream.Send(&hello.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	// 然后在循环中接收服务端返回的数据：
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
