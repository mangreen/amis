FROM golang:1.12.5

ENV GO111MODULE=on

WORKDIR $GOPATH/src/amis
COPY . .

CMD go run cmd/main.go

EXPOSE 1323