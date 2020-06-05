.PHONY: all test

all: test
	go build

small:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s"

test:
	@echo "Run test..."
	@go test ./...

clean:
	- rm -r internet-topology

docker:
	docker build -t internet-topology:latest .
