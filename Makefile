# fixed it
bin-deps:
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
	go get github.com/envoyproxy/protoc-gen-validate@latest && go install github.com/envoyproxy/protoc-gen-validate@latest \
	go get google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

deps:
	go mod tidy && go mod vendor

# export GOBIN==GOPATH
generate:
	buf generate

# прописать локальную установку https://golangci-lint.run/usage/install/
# настроить его, чтобы названия тоже пождчеркивала, что не поканонам
lint:
	golangci-lint run

clean:
	rm -rf pkg
