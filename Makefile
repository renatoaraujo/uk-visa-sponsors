.PHONY: build

build:
	go build -o sponsors ./cmd/cli/main.go

test:
	go test -race -v ./...

help: build
	@./sponsors help

find: build
	@./sponsors find --company $(firstword $(filter-out $@,$(MAKECMDGOALS)))

%:
	@: