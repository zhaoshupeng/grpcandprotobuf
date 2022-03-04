package main

import (
	"fmt"
	"golearning/rpc/v2/rpcRegular"
	"log"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

// 2.1 为了简化客户端用户调用RPC函数，我们在可以在接口规范部分增加对客户端的简单包装：

type HelloServiceClient struct {
	*rpc.Client
}

// 得到编译器提供的安全保障。
var _ rpcRegular.HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
// 2.1.1 我们在接口规范中针对客户端新增加了 HelloServiceClient 类型，
//该类型也必须满足 HelloServiceInterface 接口，这样客户端用户就可以直接通过接口对应的方法调用 RPC 函数。
//同时提供了一个 DialHelloService 方法，直接拨号 HelloService 服务。

// 现在客户端用户不用再担心 RPC 方法名字或参数类型不匹配等低级错误的发生。

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
