# Build stage
FROM golang:latest AS builder

WORKDIR /go/src/github.com/openbrisk/brisk-runtime-go
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o server ./src/server.go

# Release stage
FROM golang:latest

WORKDIR /app

RUN apt-get update

COPY --from=builder /go/src/github.com/openbrisk/brisk-runtime-go/server /app/server
COPY startup.sh .

RUN chmod +x startup.sh

EXPOSE 8080
ENTRYPOINT [ "./startup.sh" ]