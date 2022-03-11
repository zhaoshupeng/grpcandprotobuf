package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

var (
	port = ":5000"

	tlsDir        = "./tls_san_config"
	//tlsServerName = "server.io"
	tlsServerName = "localhost"

	ca         = tlsDir + "/ca.crt"
	server_crt = tlsDir + "/server.crt"
	server_key = tlsDir + "/server.key"
	client_crt = tlsDir + "/client.crt"
	client_key = tlsDir + "/client.key"
)

type myGrpcServerSan struct{}

func (s *myGrpcServerSan) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello " + in.Name}, nil
}



func startServerSan() {
	//从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	certificate, err := tls.LoadX509KeyPair(server_crt, server_key)
	if err != nil {
		log.Panicf("could not load server key pair: %s", err)
	}

	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(ca)
	if err != nil {
		log.Panicf("could not read ca certificate: %s", err)
	}

	//尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Panic("failed to append client certs")
	}

	//构建基于 TLS 的 TransportCredentials 选项
	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		//要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		//设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{certificate},
		//设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs:    certPool,
	})

	server := grpc.NewServer(grpc.Creds(creds))
	RegisterGreeterServer(server, new(myGrpcServerSan))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicf("could not list on %s: %s", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}

func doClientWorkSan() {
	certificate, err := tls.LoadX509KeyPair(client_crt, client_key)
	if err != nil {
		log.Panicf("could not load client key pair: %s", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(ca)
	if err != nil {
		log.Panicf("could not read ca certificate: %s", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Panic("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: false,         // NOTE: this is required! // 实验后不是必须
		ServerName:         tlsServerName, // NOTE: this is required!  // 实验后不是必须
		Certificates:       []tls.Certificate{certificate},
		RootCAs:            certPool,
	})

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


//func startServerSanGen() {
//
//	lis, err := net.Listen("tcp", port)
//	if err != nil {
//		grpclog.Fatalf("Failed to listen: %v", err)
//	}
//
//	// TLS 认证
//	creds, err := credentials.NewServerTLSFromFile("../../tls/server.pem", "../../tls/server.key")
//	if err != nil {
//		grpclog.Fatalf("Failed to generate credentials %v", err)
//	}
//
//	// 实例化 grpc Server, 并开启 TLS 认证
//	server := grpc.NewServer(grpc.Creds(creds))
//
//	// 注册 HelloService
//	RegisterGreeterServer(server, new(myGrpcServerSan))
//
//	if err := server.Serve(lis); err != nil {
//		log.Panicf("grpc serve error: %s", err)
//	}
//
//}

//func doClientWorkSanGen() {
//	certificate, err := tls.LoadX509KeyPair(client_crt, client_key)
//	if err != nil {
//		log.Panicf("could not load client key pair: %s", err)
//	}
//
//	certPool := x509.NewCertPool()
//	ca, err := ioutil.ReadFile(ca)
//	if err != nil {
//		log.Panicf("could not read ca certificate: %s", err)
//	}
//	if ok := certPool.AppendCertsFromPEM(ca); !ok {
//		log.Panic("failed to append ca certs")
//	}
//
//	creds := credentials.NewTLS(&tls.Config{
//		InsecureSkipVerify: false,         // NOTE: this is required!
//		ServerName:         tlsServerName, // NOTE: this is required!
//		Certificates:       []tls.Certificate{certificate},
//		RootCAs:            certPool,
//	})
//
//	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(creds))
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer conn.Close()
//
//	c := NewGreeterClient(conn)
//
//	r, err := c.SayHello(context.Background(), &HelloRequest{Name: "gopher"})
//	if err != nil {
//		log.Fatalf("could not greet: %v", err)
//	}
//	log.Printf("doClientWork: %s", r.Message)
//}


//------
func startServerSanPem() {
	//从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert,_ := tls.LoadX509KeyPair("tls_san_cert_config/server.pem","tls_san_cert_config/server.key")
	// 创建一个新的、空的 CertPool
	certPool  := x509.NewCertPool()
	ca,_ := ioutil.ReadFile("tls_san_cert_config/ca.pem")
	//尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	//构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		//设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		//要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ClientAuth: tls.RequireAndVerifyClientCert,
		//设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})
	rpcServer := grpc.NewServer(grpc.Creds(creds))

	RegisterGreeterServer(rpcServer, new(myGrpcServerSan))

	lis,_ := net.Listen("tcp",":8081")

	rpcServer.Serve(lis)
}

func doClientWorkSanPem() {
	//从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert,_ := tls.LoadX509KeyPair("tls_san_cert_config/client.pem","tls_san_cert_config/client.key")
	// 创建一个新的、空的 CertPool
	certPool  := x509.NewCertPool()
	ca,_ := ioutil.ReadFile("tls_san_cert_config/ca.pem")
	//尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	//构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		//设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		//要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ServerName: "localhost",
		RootCAs:certPool,
	})
	conn,err := grpc.Dial(":8081",grpc.WithTransportCredentials(creds))
	if err != nil{
		log.Fatal(err)
		return
	}
	defer conn.Close()

	c := NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &HelloRequest{Name: "gopher"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("doClientWork: %s", r.Message)

	//prodClient := service.NewProdServiceClient(conn)
	//prdRes,err :=prodClient.GetProdStock(context.Background(),
	//	&service.ProdRequest{ProdId:1})
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(prdRes.ProdStock)
}