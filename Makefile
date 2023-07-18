.PHONY: build tidy protoc clean

GO_BIN := $(shell echo ${GOPATH}/bin)
SERVER_SRCS := $(wildcard ./cmd/server/*.go)
CLIENT_SRCS := $(wildcard ./cmd/client/*.go)
SERVER_EXECUTABLE_FILE := bin/server
CLIENT_EXECUTABLE_FILE := bin/client

build: $(SERVER_EXECUTABLE_FILE) $(CLIENT_EXECUTABLE_FILE)
$(SERVER_EXECUTABLE_FILE): $(SERVER_SRCS)
	@go build -o $(SERVER_EXECUTABLE_FILE) $(SERVER_SRCS)
$(CLIENT_EXECUTABLE_FILE): $(CLIENT_SRCS)
	@go build -o $(CLIENT_EXECUTABLE_FILE) $(CLIENT_SRCS)
clean:
	@rm -rf bin
tidy:
	@go mod tidy
install:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
version: build
	@go version -m ${SERVER_EXECUTABLE_FILE}
	@go version -m ${CLIENT_EXECUTABLE_FILE}
protoc:
	@protoc \
		--go_out=. --go_opt=Mproto/hello.proto=pkg/grpc \
		--go-grpc_out=. --go-grpc_opt=Mproto/hello.proto=pkg/grpc \
		proto/hello.proto
