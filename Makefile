.PHONY: build

build:
	go build -o sponsors

find: build
	@./sponsors find --company $(filter-out $@,$(MAKECMDGOALS))

%:
	@: