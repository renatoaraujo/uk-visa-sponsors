.PHONY: build

build:
	go build -o ukvisasponsors ./cmd/cli/main.go

test:
	go test -race -v ./...

help: build
	@./sponsors help

find: build
	@./sponsors find --company $(firstword $(filter-out $@,$(MAKECMDGOALS)))

%:
	@: