package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	port = ":5000"
)

type myGrpcServer struct{}

func (s *myGrpcServer) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	go startServer()
	time.Sleep(5 *time.Second)

	doClientWork()
}

func startServer() {
	creds, err := credentials.NewServerTLSFromFile("tls_san_config/server.crt", "tls_san_config/server.key")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	RegisterGreeterServer(grpcServer, new(myGrpcServer))

	mux := http.NewServeMux()
	// curl https://localhost:5000/ -k
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	// 因为 gRPC 服务已经实现了 ServeHTTP 方法，可以直接作为 Web 路由处理对象。如果将 gRPC 和 Web 服务放在一起，会导致 gRPC 和 Web 路径的冲突，在处理时我们需要区分两类服务。
	// 我们可以通过以下方式生成同时支持 Web 和 gRPC 协议的路由处理函数：
	//首先 gRPC 是建立在 HTTP/2 版本之上，如果 HTTP 不是 HTTP/2 协议则必然无法提供 gRPC 支持。同时，每个 gRPC 调用请求的 Content-Type 类型会被标注为 "application/grpc" 类型。
	//这样我们就可以在 gRPC 端口上同时提供 Web 服务了。
	http.ListenAndServeTLS(port, "tls_san_config/server.crt", "tls_san_config/server.key", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks
		// https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			mux.ServeHTTP(w, r)
		}
	}))
}

func doClientWork() {
	creds, err := credentials.NewClientTLSFromFile("tls_san_config/server.crt", "localhost")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &HelloRequest{Name: "gopher"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("doClientWork: %s", r.Message)
}



// 启用普通的 https 服务器则非常简单：
func startNormalHttps() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "hello")
	})

	http.ListenAndServeTLS(port, "tls_san_config/server.crt", "tls_san_config/server.key",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mux.ServeHTTP(w, r)
			return
		}),
	)
}