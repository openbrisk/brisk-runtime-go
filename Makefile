.PHONY: build run

build:
	docker build -t openbrisk/brisk-runtime-go .

run:
	docker run -it \
	-p 8080:8080 \
	openbrisk/brisk-runtime-go