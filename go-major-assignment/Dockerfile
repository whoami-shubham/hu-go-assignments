FROM golang:1.13

WORKDIR /go/src/go-major-assignment
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

CMD ["./go-major-assignment"]