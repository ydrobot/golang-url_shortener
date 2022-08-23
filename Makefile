# fixed it
bin-deps:
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
	go get github.com/envoyproxy/protoc-gen-validate@latest && go install github.com/envoyproxy/protoc-gen-validate@latest \
	go get google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \

deps:
	go mod tidy && go mod vendor

# export GOBIN==GOPATH
generate:
	buf lint && buf generate

# прописать локальную установку https://golangci-lint.run/usage/install/
# настроить его, чтобы названия тоже пождчеркивала, что не поканонам
lint:
	golangci-lint run

clean:
	rm -rf pkg

# получать переменные из конфига (?)
init-db:
	docker run --name url_shortener -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432  -d postgres

migration-status:
	cd migrations; goose postgres "user=postgres password=postgres dbname=postgres sslmode=disable" status

migration-up:
	cd migrations; goose postgres "user=postgres password=postgres dbname=postgres sslmode=disable" up

migration-down:
	cd migrations; goose postgres "user=postgres password=postgres dbname=postgres sslmode=disable" down
