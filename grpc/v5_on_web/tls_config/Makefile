
default:
	#为了避免证书的传递过程中被篡改，可以通过一个安全可靠的根证书分别对服务器和客户端的证书进行签名。这样客户端或服务器在收到对方的证书后可以通过根证书进行验证证书的有效性。

	# ca
	#生成根证书
	openssl genrsa -out ca.key 2048
	openssl req -new -x509 -days 3650 \
		-subj "/C=GB/L=China/O=gobook/CN=github.com" \
		-key ca.key -out ca.crt

	# server
	# 为服务器生成私钥
	openssl genrsa \
		-out server.key 2048
	# 签名的过程中引入了一个新的以. csr 为后缀名的文件，它表示证书签名请求文件。在证书签名完成之后可以删除. csr 文件。
	openssl req -new \
		-subj "/C=GB/L=China/O=server/CN=server.io" \
		-key server.key \
		-out server.csr
	# 对服务器端证书进行签名：
	openssl x509 -req -sha256 \
		-CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 \
		-in server.csr \
		-out server.crt

	# client
	# 如果客户端的证书也采用 CA 根证书签名的话，服务器端也可以对客户端进行证书认证。我们用 CA 根证书对客户端证书签名：
	# 为客户端生成私钥
	openssl genrsa \
		-out client.key 2048
	openssl req -new \
		-subj "/C=GB/L=China/O=client/CN=client.io" \
		-key client.key \
		-out client.csr
	openssl x509 -req -sha256 \
		-CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 \
		-in client.csr \
		-out client.crt

firstversion:
	# server
	# 为服务器生成私钥
	openssl genrsa -out server.key 2048
	# 为服务器生成证书
	openssl req -new -x509 -days 3650 \
    	-subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io" \
    	-key server.key -out server.crt
    #其中以. key 为后缀名的是私钥文件，需要妥善保管。以. crt 为后缀名是证书文件，也可以简单理解为公钥文件，并不需要秘密保存。
    # 在 subj 参数中的 /CN=server.grpc.io 表示服务器的名字为 server.grpc.io，在验证服务器的证书时需要用到该信息。

    # client
    # 为客户端生成私钥
	openssl genrsa -out client.key 2048
	# 为客户端生成证书
	openssl req -new -x509 -days 3650 \
    	-subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io" \
    	-key client.key -out client.crt

clean:
	-rm *.key *.crt *.csr *.srl

