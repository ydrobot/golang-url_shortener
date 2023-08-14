export GOBIN=$(shell pwd)/bin

deps:
	go mod tidy && go mod vendor

# прописать локальную установку https://golangci-lint.run/usage/install/
# настроить его, чтобы названия тоже пождчеркивала, что не поканонам
lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0

lint: lint-install
	$(GOBIN)/golangci-lint --version
	$(GOBIN)/golangci-lint run --config=.golangci_critical.yml

test:
	go test -mod=vendor ./...