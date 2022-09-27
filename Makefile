BINARY_NAME=MessageSecure

build:
	GOARCH=amd64 GOOS=darwin go build -o bin/darwin/MessageSecure main.go
	cp -R views/ bin/darwin/views
	cp -R public/ bin/darwin/public

	GOARCH=amd64 GOOS=linux go build -o bin/linux/MessageSecure main.go
	cp -R views/ bin/linux/
	cp -R public/ bin/linux/

# GOARCH=amd64 GOOS=window go build -o bin/MessageSecure-windows main.go

# run:
#  ./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm -rf bin/darwin
	rm -rf bin/linux
# rm ${BINARY_NAME}-windows