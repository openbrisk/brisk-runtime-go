#!/bin/bash

# Load dependencies
#if [ -e /openbrisk/$MODULE_NAME.toml ]
#then
#    cd /openbrisk/
#    cp $MODULE_NAME.toml Gopkg.toml
#    dep ensure
#    cd /app
#fi

# Build the function shared object.
GOOS=linux GOARCH=amd64 go build -buildmode=plugin -o /openbrisk/$MODULE_NAME.so /openbrisk/$MODULE_NAME.go

# Start the server.
./server