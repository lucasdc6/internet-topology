FROM golang:1.13.5 as builder

WORKDIR /go/src/github.com/lucasdc6/internet-topology

ADD . /go/src/github.com/lucasdc6/internet-topology


RUN make small

FROM busybox

COPY --from=builder /go/src/github.com/lucasdc6/internet-topology/internet-topology /

ENTRYPOINT [ "/internet-topology" ]
