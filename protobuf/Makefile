generateV2:
	protoc -I . --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true todo.proto
	protoc -I . --swagger_out=logtostderr=true,grpc_api_configuration=todo.yaml:. todo.proto
	protoc -I . --grpc-gateway_out . --go_out=plugins=grpc:. todo.proto
	protoc -I . --grpc-gateway_out=logtostderr=true,grpc_api_configuration=todo.yaml:. todo.proto