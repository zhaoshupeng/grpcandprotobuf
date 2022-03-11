package main

import (
	"context"
	"golearning/grpc/v1/pbgo/hello"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type HelloServiceImpl struct {}

// 基于服务端的 HelloServiceServer 接口可以重新实现 HelloService 服务：
func (p *HelloServiceImpl) Hello(ctx context.Context, args *hello.String,) (*hello.String,error) {
	reply := &hello.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream hello.HelloService_ChannelServer) error {
	for {
		// 服务端在循环中接收客户端发来的数据，如果遇到 io.EOF 表示客户端流被关闭，如果函数退出表示服务端流关闭。
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &hello.String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}




// gRPC 服务的启动流程和标准库的 RPC 服务启动流程类似：
func main() {

	// 首先是通过 grpc.NewServer() 构造一个 gRPC 服务对象
	grpcServer := grpc.NewServer()
	// 然后通过 gRPC 插件生成的 RegisterHelloServiceServer 函数注册我们实现的 HelloServiceImpl 服务。
	hello.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	// 然后通过 grpcServer.Serve(lis) 在一个监听端口上提供 gRPC 服务。
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)

}