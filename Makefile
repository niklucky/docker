.PHONY: build run watch reflex

GO_PROJECT_NAME := docker-spawner-example

build:
	go build -o example examples/main.go

run: build
	./example

watch:
	ulimit -n 1000
	reflex -s -r '\.go$$' make run

test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := watch