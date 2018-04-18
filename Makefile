REVISION := $(shell git rev-parse --short HEAD)

.PHONY: build-cli \
	fetch-dep

build-cli:
	go build -v -o ./build/bin/cmd -ldflags "-X main.revision=${REVISION}" main.go

fetch-dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	