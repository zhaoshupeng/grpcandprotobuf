
run:
	@go build -o a.out && ./a.out
	-@rm ./a.out

aa:
	cd ../../.. && pwd

gen:
	protoc  \
		--proto_path=${GOPATH}/src \
		--proto_path=. \
		--grpc-gateway_out=. \
		--go_out=plugins=grpc:.\
		--openapiv2_out . \
		helloworld.proto

## 参考： https://go-kratos.dev/docs/guide/openapi/
## https://juejin.cn/post/7017396592428711972
##		--swagger_out  已经被 --openapiv2_out 替换; not supporting gopath since v.2.0.0 but a non-functioning build is very problematic
## 官方安装文档错误： https://github.com/grpc-ecosystem/grpc-gateway/issues/1992
## 可以接着使用tag 2之下的，不然生成的gw.go文件中导入路径含有v2



#	protoc \
#		-I=. -I=../../../github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#		--go_out=plugins=grpc:. \
#		helloworld.proto
#	protoc \
#		-I=. -I=../../../github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#		--grpc-gateway_out=. \
#		helloworld.proto
#	protoc \
#		-I=. -I=../../../github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#		--swagger_out=. \
#		helloworld.proto


clean:
	-rm *.pb.go *.pb.gw.go