# Build stage
FROM golang:latest AS builder

WORKDIR /go/src/github.com/openbrisk/brisk-runtime-go
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o server ./src/server.go

# Release stage
FROM golang:latest

WORKDIR /app

RUN apt-get update && apt-get install -y wget

ARG DEP_VERSION=0.4.1
RUN wget -q https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -O /usr/local/bin/dep && chmod +x /usr/local/bin/dep

COPY --from=builder /go/src/github.com/openbrisk/brisk-runtime-go/server /app/server
COPY startup.sh .
RUN chmod +x startup.sh

EXPOSE 8080
ENTRYPOINT [ "./startup.sh" ]