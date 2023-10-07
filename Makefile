.PHONY: build

build:
	go build -o sponsors

help: build
	@./sponsors help

find: build
	@./sponsors find --company $(firstword $(filter-out $@,$(MAKECMDGOALS))) --details $(word 2, $(filter-out $@,$(MAKECMDGOALS)))

%:
	@: