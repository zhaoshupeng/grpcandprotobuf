package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	port         = ":5000"
	echoEndpoint = flag.String("echo_endpoint", "localhost"+port, "endpoint of YourService")
)

type myGrpcServer struct{}

func (s *myGrpcServer) Get(ctx context.Context, in *StringMessage) (*StringMessage, error) {
	return &StringMessage{Value: "Get: " + in.Value}, nil
}
func (s *myGrpcServer) Post(ctx context.Context, in *StringMessage) (*StringMessage, error) {
	return &StringMessage{Value: "Post: " + in.Value}, nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 首先通过 runtime.NewServeMux() 函数创建路由处理器，然后通过 RegisterRestServiceHandlerFromEndpoint 函数
	//将 RestService 服务相关的 REST 接口中转到后面的 gRPC 服务。
	mux := runtime.NewServeMux() //grpc-gateway 提供的 runtime.ServeMux 类也实现了 http.Handler 接口，因此可以和标准库中的相关函数配合使用。
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// RegisterRestServiceHandlerFromEndpoint 函数用于将定义了 Rest 接口的请求转发到真正的 gRPC 服务。
	//注册路由处理函数之后就可以启动 Web 服务了：
	err := RegisterRestServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

// $ curl localhost:8080/get/gopher
// {"value":"Get: gopher"}

// $ curl localhost:8080/post -X POST --data '{"value":"grpc"}'
// {"value":"Post: grpc"}

func main() {
	flag.Parse()
	defer glog.Flush()

	go startGrpcServer()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

// 启动 grpc 服务 , 端口 5000
func startGrpcServer() {
	server := grpc.NewServer()
	RegisterRestServiceServer(server, new(myGrpcServer))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicf("could not list on %s: %s", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}
