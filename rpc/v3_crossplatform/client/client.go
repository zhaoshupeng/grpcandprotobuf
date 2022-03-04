package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var flagAddr = flag.String("addr", "localhost:1234", "server address")


// 实现 json 版本的客户端：
func main() {
	flag.Parse()

	// nc localhost 1234
	// {"method":"HelloService.Hello","params":["hello"],"id":0}
	// echo -e '{"method":"HelloService.Hello","params":["hello2222"],"id":3}' | nc localhost 1234
	// echo -e '{"method":"HelloService.Hello","params":["hello2222"],"id":3}{"method":"HelloService.Hello","params":["hello33"],"id":4}' | nc localhost 1234

	// 先手工调用 net.Dial 函数建立 TCP 链接，然后基于该链接建立针对客户端的 json 编解码器。
	conn, err := net.Dial("tcp", *flagAddr)
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
// 在确保客户端可以正常调用 RPC 服务的方法之后，我们用一个普通的 TCP 服务代替 Go 语言版本的 RPC 服务，
//这样可以查看客户端调用时发送的数据格式。比如通过 nc 命令 nc -l 1234 在同样的端口启动一个 TCP 服务。然后再次执行
//一次 RPC 调用将会发现 nc 输出了以下的信息： {"method":"HelloService.Hello","params":["hello"],"id":0}
// 这是一个 json 编码的数据，其中 method 部分对应要调用的 rpc 服务和方法组合成的名字，params 部分的第一个元素为参数，id 是由调用端维护的一个唯一的调用编号


//在获取到 RPC 调用对应的 json 数据后，我们可以通过直接向架设了 RPC 服务的 TCP 服务器发送 json 数据模拟 RPC 方法调用(即向真实的服务端)：
// $ echo -e '{"method":"HelloService.Hello","params":["hello"],"id":1}' | nc localhost 1234

//返回的结果也是一个 json 格式的数据：
//{"id":1,"result":"hello:hello","error":null}


