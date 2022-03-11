
run:
	@go build -o a.out && ./a.out
	-@rm ./a.out

aa:
	cd ../../.. && pwd

gen:
#	protoc -I=. -I=../../.. --go_out=plugins=grpc:. helloworld.proto
#	protoc -I=. -I=../../.. --govalidators_out=. helloworld.proto
	protoc  \
		--proto_path=${GOPATH}/src \
		--proto_path=. \
		--govalidators_out=. --go_out=plugins=grpc:.\
		helloworld.proto

clean:
	-rm *.pb.go