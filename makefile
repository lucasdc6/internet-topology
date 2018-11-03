DRAW_SCRIPT=draw/draw.py
GO_OUTPUT_DATA=datsets/
DRAW_OUTPUT_DATA=media

all: folders
	go build

runtest:
	@go test ./...

folders:
	- mkdir $(GO_OUTPUT_DATA) $(DRAW_OUTPUT_DATA)

clean:
	- rm -r $(DRAW_OUTPUT_DATA) internet-topology

graph: all
	./internet-topology --asn=5692 -pi
	dot -Tsvg /tmp/internet-topology.dot > /tmp/internet-topology.svg
