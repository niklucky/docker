.PHONY: build

build:
	go build -o example -v ./examples/main.go

.DEFAULT_GOAL := build