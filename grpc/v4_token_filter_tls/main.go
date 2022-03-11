package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"time"
)

var (
	port = ":5000"
)

type myGrpcServer struct{}

func (s *myGrpcServer) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	// 首先通过 metadata.FromIncomingContext 从 ctx 上下文中获取元信息，然后取出相应的认证信息进行认证。如果认证失败，则返回一个 codes.Unauthenticated 类型地错误。
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing credentials")
	}

	var (
		appid  string
		appkey string
	)
	fmt.Println("------------",md)
	if val, ok := md["login"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}

	if appid != "gopher" || appkey != "password" {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid token: appid=%s, appkey=%s", appid, appkey)
	}

	return &HelloReply{Message: "Hello " + in.Name}, nil
}

// 要实现对每个 gRPC 方法进行认证，需要实现 grpc.PerRPCCredentials 接口：
// 创建一个 Authentication 类型，用于实现用户名和密码的认证：
type Authentication struct {
	Login    string
	Password string
}

// 在 GetRequestMetadata 方法中返回认证需要的必要信息。
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"login": a.Login, "password": a.Password}, nil
}
// RequireTransportSecurity 方法表示是否要求底层使用安全链接。在真实的环境中建议必须要求底层启用安全的链接，否则认证信息有泄露和被篡改的风险。
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

// 要实现普通方法的截取器，需要为 grpc.UnaryInterceptor 的参数实现一个函数
// 函数的 ctx 和 req 参数就是每个普通的 RPC 方法的前两个参数。第三个 info 参数表示当前是对应的那个 gRPC 方法，第四个 handler 参数对应当前的 gRPC 方法函数。
func filter(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("fileter:", info)
	log.Println("fileter handler:", handler)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	// valiate req

	return handler(ctx, req)
}

func main() {
	//go startServer()
	go startServerFilter()
	time.Sleep(time.Second)

	doClientWork()
}

func startServer() {
	server := grpc.NewServer()
	RegisterGreeterServer(server, new(myGrpcServer))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicf("could not list on %s: %s", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}

func doClientWork() {
	// 在每次请求 gRPC 服务时就可以将 Token 信息作为参数选项传人：
	auth := Authentication{
		Login:    "gopher",
		Password: "password",
	}

	// 通过 grpc.WithPerRPCCredentials 函数将 Authentication 对象转为 grpc.Dial 参数。因为这里没有启用安全链接，需要传人 grpc.WithInsecure() 表示忽略证书认证。
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
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

// gRPC 中的 grpc.UnaryInterceptor 和 grpc.StreamInterceptor 分别对普通方法和流方法提供了截取器的支持。我们这里简单介绍普通方法的截取器用法。
// 不过 gRPC 框架中只能为每个服务设置一个截取器，因此所有的截取工作只能在一个函数中完成。开源的 grpc-ecosystem 项目中的 go-grpc-middleware 包已经基于 gRPC 对截取器实现了链式截取器的支持。
func startServerFilter() {
	server := grpc.NewServer(grpc.UnaryInterceptor(filter))
	RegisterGreeterServer(server, new(myGrpcServer))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicf("could not list on %s: %s", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}