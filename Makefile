build:
	@go build -o bin/go-bank ./cmd

run: build
	@./bin/go-bank

test: 
	@go test -v ./...

