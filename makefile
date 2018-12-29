all:
	go build

small:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s"

runtest:
	@go test ./...

clean:
	- rm -r internet-topology

docker:
	docker build -t internet-topology:latest .
