.PHONY: all
all: build

.PHONY: build
build:
	protoc --go_out=. --go_opt=module=github.com/ejweber/webersprinkler2 proto/*