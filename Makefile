# fixed it
bin-deps:
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
	go get google.golang.org/protobuf/cmd/protoc-gen-go && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate:
	buf mod update && buf generate && go mod vendor

