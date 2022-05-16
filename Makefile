SERVICE = phone-validator
PKG_LIST = $(shell go list ./... | grep -v mock)

sqlite :
	$(MAKE) -C ./db

clean:
	rm -rf ./bin

build: sqlite clean
	mkdir bin
	GOOS=linux GOARCH=amd64 go build -o bin/$(SERVICE) main.go
	ls

build-osx: clean
	mkdir bin
	go build -o bin/$(SERVICE) main.go

run:
	go run main.go

test:
	go test ./...

lint:
	golint -set_exit_status $(PKG_LIST)