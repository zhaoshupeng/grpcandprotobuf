package rpcRegular

import "net/rpc"

/**
	在涉及RPC的应用中，作为开发人员一般至少有三种角色：首先是服务端实现RPC方法的开发人员，其次是客户端调用RPC方法的人员，最后也
是最重要的是制定服务端和客户端RPC接口规范的设计人员。
	在v1例子中我们为了简化将以上几种角色的工作全部放到了一起，虽然看似实现简单，但是不利于后期的维护和工作的切割。

*/

// 1. 如果要重构HelloService服务，第一步需要明确服务的名字和接口：
const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
// 1.1 我们将RPC服务的接口规范分为三个部分：首先是服务的名字，然后是服务要实现的详细方法列表，最后是注册该类型服务的函数。
// 为了避免名字冲突，我们在RPC服务的名字中增加了包路径前缀（这个是RPC服务抽象的包路径，并非完全等价Go语言的包路径）。
//	RegisterHelloService注册服务时，编译器会要求传入的对象满足HelloServiceInterface接口。


// 2. 在定义了RPC服务接口规范之后，客户端就可以根据规范编写RPC调用的代码了：





