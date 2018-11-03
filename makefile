all:
	go build

runtest:
	@go test ./...

clean:
	- rm -r internet-topology
