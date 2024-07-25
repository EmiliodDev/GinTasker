build:
	@go build -o bin/ginTasker cmd/ginTasker/main.go

test:
	@go test -v ./...

run: build
	@./bin/ginTasker
