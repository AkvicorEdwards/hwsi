FROM golang:1.13.5-buster
COPY ./ /go/src/github.com/AkvicorEdwards/hwsi
ENV GO111MODULE=on
WORKDIR /go/src/github.com/AkvicorEdwards/hwsi
RUN go mod download
RUN go build hwsi.go
RUN chmod +x hwsi
CMD ["./hwsi"]

