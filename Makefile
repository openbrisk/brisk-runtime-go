.PHONY: build run compile start

build:
	docker build -t openbrisk/brisk-runtime-go .

run:
	docker run -it \
	-p 8080:8080 \
	-e MODULE_NAME=echo \
	-e FUNCTION_HANDLER=Execute \
	-e FUNCTION_TIMEOUT=10 \
	-v `pwd`/examples:/openbrisk \
	openbrisk/brisk-runtime-go

compile:
	go build ./src/server.go	

start:
	./server