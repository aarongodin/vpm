build:
	@go build -o bin/vpm main.go

ci:
	@golangci-lint run ./...

