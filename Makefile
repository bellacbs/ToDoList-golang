build:
	@go build -o bin/todolist-golang

run: build
	@./bin/todolist-golang

dev:
	@gin run *go

test:
	@go test -v ./...