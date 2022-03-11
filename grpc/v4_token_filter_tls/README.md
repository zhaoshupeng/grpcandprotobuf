# grpcandprotobuf

#go version
go version go1.17.6 linux/amd64

gRPC 还为每个 gRPC 方法调用提供了认证支持，这样就基于用户 Token 对不同的方法访问进行权限管理。
要实现对每个 gRPC 方法进行认证，需要实现 grpc.PerRPCCredentials 接口

#1.介绍

gRPC建立在HTTP/2协议之上，对TLS提供了很好的支持。当不需要证书认证时,可通过grpc.WithInsecure()选项跳过了对服务器证书的验证，没有启用证书的gRPC服务和客户端进行的是明文通信，信息面临被任何第三方监听的风险。为了保证gRPC通信不被第三方监听、篡改或伪造，可以对服务器启动TLS加密特性。
#2. 概念简述
   2.1 什么是CA

CA是Certificate Authority的缩写，也叫“证书授权中心”。它是负责管理和签发证书的第三方机构，作用是检查证书持有者身份的合法性，并签发证书，以防证书被伪造或篡改。

    CA实际上是一个机构，负责“证件”印制核发。就像负责颁发身份证的公安局、负责发放行驶证、驾驶证的车管所。

2.2 什么是CA证书

CA 证书就是CA颁发的证书。我们常听到的数字证书就是CA证书,CA证书包含信息有:证书拥有者的身份信息，CA机构的签名，公钥和私钥,

    身份信息: 用于证明证书持有者的身份
    CA机构的签名: 用于保证身份的真实性
    公钥和私钥: 用于通信过程中加解密，从而保证通讯信息的安全性

2.3 什么是SAN

    SAN(Subject Alternative Name) 是 SSL 标准 x509 中定义的一个扩展。使用了 SAN 字段的 SSL 证书，可以扩展此证书支持的域名，使得一个证书可以支持多个不同域名的解析。

#注意
(1)openssl 生成证书上 grpc 报 legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0

    如果出现上述报错，是因为 go 1.15 版本开始废弃 CommonName，因此推荐使用 SAN 证书。
    SAN(Subject Alternative Name) 是 SSL 标准 x509 中定义的一个扩展。使用了 SAN 字段的 SSL 证书，可以扩展此证书支持的域名，使得一个证书可以支持多个不同域名的解析。

(2) 上面生成证书请求时的几个字段的意义：

    C  => Country
    ST => State
    L  => City
    O  => Organization
    OU => Organization Unit
    CN => Common Name (证书所请求的域名)
    emailAddress => main administrative point of contact for the certificate

    字段	    字段含义	                示例
    /C=	    Country 国家	            CN
    /ST=	State or Province 省	Guangdong
    /L=	    Location or City 城市	Guangzhou
    /O=	    Organization 组织或企业	xdevops
    /OU=	Organization Unit 部门	xdevops
    /CN=	Common Name 域名或IP	    gitlab.xdevops.cn


## 参考：
https://www.jianshu.com/p/d3c4c8eb8c30

https://www.cnblogs.com/jackluo/p/13841286.html

http://gaodongfei.com/archives/start-grpc 单向双向验证

https://github.com/cloud-org/grpc-auth-sample