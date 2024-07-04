run:
	@go run ./cmd/main.go $(ARGS)

test:
	go test ./...
