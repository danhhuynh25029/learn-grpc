gen:
	protoc --go-grpc_out=. ./chat/service.proto
	protoc --go_out=. ./chat/service.proto
