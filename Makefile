.PHONY: all
all: build

.PHONY: build
build: generate
	go build cmd/...

.PHONY: generate
generate:
	protoc --go_out=. --go_opt=module=github.com/ejweber/webersprinkler2 proto/*

.PHONY: clean
clean:
	rm -rf bin
