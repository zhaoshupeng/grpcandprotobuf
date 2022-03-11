## 参考：https://www.jianshu.com/p/d3c4c8eb8c30
#https://github.com/cloud-org/grpc-auth-sample
Country :="CN"
State :="GZ"
Location :="GZ"
Organization :="ronething"
Organizational :="ronething"
CommonName :="localhost"

default:
	# ca
	#生成根证书私钥
	openssl genrsa -out ca.key 2048
#生成CA证书请求，不能加 -x509 参数；不然生成ca.crt时报错:error:0906D06C:PEM routines:PEM_read_bio:no start line:pem_lib.c:701:Expecting: CERTIFICATE REQUEST
#https://stackoverflow.com/questions/67721280/openssl-csr-generation-gives-expecting-certificate-request
# openssl req -text -in ca.csr -noout -verify
	openssl req -new  -days 3650 \
		-subj "/C=${Country}/ST=${State}/L=${Location}/O=${Organization}/OU=${Organizational}/CN=${CommonName}" \
		-key ca.key -out ca.csr
	#自签名得到根证书
	openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.crt

#	openssl req -new -x509 -key ca.key -out ca.crt -days 3650 -subj "/C=${Country}/ST=${State}/L=${Location}/O=${Organization}/OU=${Organizational}/CN=${CommonName}"



#	# server
#	# 为服务器生成私钥
	openssl genpkey -algorithm RSA \
		-out server.key
#	# 签名的过程中引入了一个新的以. csr 为后缀名的文件，它表示证书签名请求文件。在证书签名完成之后可以删除. csr 文件。
#	#根据私钥server.key生成证书请求文件server.csr
	openssl req -new -nodes -key server.key \
		-out server.csr  \
		-subj "/C=${Country}/O=${Organization}/OU=${Organizational}/CN=${CommonName}" -config ./bak.openssl.cnf \
		-extensions v3_req
#	#请求CA对证书请求文件签名，生成最终证书文件
	openssl x509 -req -days 365 -in server.csr -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./bak.openssl.cnf -extensions v3_req



	# client san
	openssl genrsa  -out client.key
#	openssl genpkey -algorithm RSA -out client.key#也可以采用这个命令
	openssl req -new -nodes -key client.key -out client.csr -days 3650 -subj "/C=${Country}/O=${Organization}/OU=${Organizational}/CN=${CommonName}" -config ./bak.openssl.cnf -extensions v3_req

	openssl x509 -req -days 3650 -in client.csr -out client.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./bak.openssl.cnf -extensions v3_req


#	# client without san
#    # 为客户端生成私钥
#	openssl genrsa -out client.key 2048
#	# 为客户端生成证书
#	openssl req -new -x509 -days 3650 \
#    	-subj "/C=GB/L=China/O=client/CN=client.io" \
#    	-key client.key -out client.crt

clean:
	-rm *.key *.crt *.csr *.srl

