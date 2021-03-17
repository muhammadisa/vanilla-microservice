generate:
	protoc -I/usr/local/include -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway-ff --go_out=plugins=grpc:. todo.proto
	protoc -I/usr/local/include -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway-ff --grpc-gateway_out=logtostderr=true,grpc_api_configuration=todo.yaml:. todo.proto
	protoc -I/usr/local/include -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway-ff --swagger_out=logtostderr=true,grpc_api_configuration=todo.yaml:. todo.proto