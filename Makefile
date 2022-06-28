.PHONY: build clean tool lint help

all: build

build:
	@go build -o ./deploy-hello ./main.go
	@sh script.sh
image:
	@go build -o ./deploy-hello ./main.go
check:
	@go langci-lint run
push:
	@sh script.sh
help:
        @echo "make build: build this project"
        @echo "make check: code check"
