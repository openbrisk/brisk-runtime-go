#!/bin/bash

# Build the function shared object.
GOOS=linux GOARCH=amd64 go build -buildmode=plugin -o /openbrisk/$MODULE_NAME.so /openbrisk/$MODULE_NAME.go

# TODO: Load dependencies

# Start the server.
./server