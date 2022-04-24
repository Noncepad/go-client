BINARY_NAME=rpc-server

build:
	mkdir ./bin
	go build -o ./bin/${BINARY_NAME} ./cmd/rpc
#	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin ./cmd/rpc
#	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux ./cmd/rpc
#	GOARCH=amd64 GOOS=window go build -o ./bin/${BINARY_NAME}-windows ./cmd/rpc

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm -r ./bin
