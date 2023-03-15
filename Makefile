BINARY_NAME=gptq

build:
	go build -o ${BINARY_NAME} cmd/gptq/main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows