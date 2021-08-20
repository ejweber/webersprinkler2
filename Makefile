.PHONY: all
all: build

.PHONY: build
build: generate
	go build cmd/...

.PHONY: generate
generate:
	protoc --go_out=. --go-grpc_out=. --go-grpc_opt=module=github.com/ejweber/webersprinkler2 --go_opt=module=github.com/ejweber/webersprinkler2 proto/*

.PHONY: clean
clean:
	rm -rf bin
