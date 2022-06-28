.PHONY: build clean tool lint help

all: build

build:
        @go build -o ./app main.go
check:
        @go langci-lint run
help:
        @echo "make build: build this project"
        @echo "make check: code check"
