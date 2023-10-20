build:
	@go build -o bin/todolist-golang

run: build
	@./bin/todolist-golang

test:
	@go test -v ./...