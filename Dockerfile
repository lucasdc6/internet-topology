FROM golang:1.10.4 as builder

WORKDIR /go/src/github.com/lucasdc6/internet-topology

ADD . /go/src/github.com/lucasdc6/internet-topology


RUN go get -u github.com/golang/dep/... \
    && dep ensure -v \
    && make

FROM alpine

COPY --from=builder /go/src/github.com/lucasdc6/internet-topology/internet-topology /app/internet-topology

CMD [ "/app/internet-topology" ]
