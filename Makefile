gen-all: generate-grpc generate-http generate-swagger gen-statik

generate-grpc:
	protoc -I/usr/local/include -I. -I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc,paths=source_relative:. pkg/pb/service.proto

generate-http:
	protoc -I/usr/local/include -I. -I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true,paths=source_relative:. pkg/pb/service.proto

generate-swagger:
	protoc -I/usr/local/include -I. -I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--swagger_out=logtostderr=true:. pkg/pb/service.proto

gen-statik:
	rm -rf statik/
	rm -f swaggerui/service.swagger.json
	cp pkg/pb/service.swagger.json swaggerui/
	statik -src swaggerui
