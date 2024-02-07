build:
	@go build -o bin/gb-backend

run: build
	@./bin/gb-backend

test:
	@go test -v ./...