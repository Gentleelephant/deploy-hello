.PHONY: build clean tool lint help

all: build

build:
	@bash script.sh
check:
	@go langci-lint run
help:
        @echo "make build: build this project"
        @echo "make check: code check"
