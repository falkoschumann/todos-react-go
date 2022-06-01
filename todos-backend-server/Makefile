all: test linux mac-intel mac-m1 win

run:
	go run main.go

build:
	go build

clean:
	go clean
	rm -rf ./build

test:
	go test

reformat:
	go fmt

linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/linux/todos_backend_server

mac-intel:
	GOOS=darwin GOARCH=amd64 go build -o ./build/mac/intel/todos_backend_server

mac-m1:
	GOOS=darwin GOARCH=arm64 go build -o ./build/mac/m1/todos_backend_server

win:
	GOOS=windows GOARCH=amd64 go build -o ./build/win/todos_backend_server.exe
