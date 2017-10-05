FROM golang:1-alpine

WORKDIR /go/src/b3-propagation-test
COPY . .

RUN go build

FROM alpine:3.6
COPY --from=0 /go/src/b3-propagation-test/b3-propagation-test /
CMD ["/b3-propagation-test"]
EXPOSE 8080
