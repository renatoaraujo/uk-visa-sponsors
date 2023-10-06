.PHONY: build

build:
	go build -o sponsors

find:
	@./sponsors find --company $(filter-out $@,$(MAKECMDGOALS))

%:
	@: