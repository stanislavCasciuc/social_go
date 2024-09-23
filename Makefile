build:
	@go build -o bin/social cmd/api/*.go
run: build
	@./bin/social