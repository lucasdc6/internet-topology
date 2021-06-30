FROM golang:1.16 as builder

WORKDIR /go/src/github.com/lucasdc6/internet-topology

ADD . /go/src/github.com/lucasdc6/internet-topology


RUN make small

FROM alpine:3.10

COPY --from=builder /go/src/github.com/lucasdc6/internet-topology/internet-topology /

ENTRYPOINT [ "/internet-topology" ]
