build:
	@go build -o bin/redCards

run: build
	@./bin/redCards