SERVICE = phone-validator
PKG_LIST = $(shell go list ./... | grep -v mock)

sqlite :
	$(MAKE) -C ./db

clean:
	rm -rf ./bin
	mkdir bin

build: sqlite clean
	GOOS=linux GOARCH=amd64 go build -o bin/$(SERVICE) main.go

build-osx: clean
	go build -o bin/$(SERVICE) main.go

run:
	go run main.go

test:
	go test ./...

lint:
	golint -set_exit_status $(PKG_LIST)

docker-build:
	docker build . -t $(SERVICE)

docker-run:
	docker run -it -p 3000:3000 phone-validator