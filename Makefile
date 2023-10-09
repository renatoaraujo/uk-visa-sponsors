.PHONY: build test help find

build:
	go build -o uk-visa-sponsors ./cmd/cli/main.go

test:
	go test -race -v ./...

help: build
	@./uk-visa-sponsors help

find: build
	@./uk-visa-sponsors find --company $(firstword $(filter-out $@,$(MAKECMDGOALS)))

%:
	@: