package main

import (
	"golearning/rpc/v2/rpcRegular"
	"log"
	"net"
	"net/rpc"
)

// 3.0 最后是基于 RPC 接口规范编写真实的服务端代码：

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	// 3.1 在新的 RPC 服务端实现中，我们用 RegisterHelloService 函数来注册函数，这样不仅可以避免命名服务名称的工作，同时也保证了传入的服务对象满足了 RPC 接口的定义
	rpcRegular.RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 3.2 最后我们新的服务改为支持多个 TCP 链接，然后为每个 TCP 链接提供 RPC 服务。
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}


