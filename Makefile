.PHONY: all build run gotool clean

BINARY_NAME="eagle"

all: gotool build

build:
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}

run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY_NAME} ];then rm -f ${BINARY_NAME}; fi

version:
	@go version
