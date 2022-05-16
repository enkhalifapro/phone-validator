SERVICE = phone-validator
PKG_LIST = $(shell go list ./... | grep -v mock)


clean:
	rm -rf ./bin

build: clean
	mkdir bin
	GOOS=linux GOARCH=amd64 go build -o bin/$(SERVICE) main.go

build-osx: clean
	mkdir bin
	go build -o bin/$(SERVICE) main.go

run:
	go run main.go

test:
	go test ./...

lint:
	golint -set_exit_status $(PKG_LIST)