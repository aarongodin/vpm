build:
	@go build -o bin/vpm main.go

ci:
	@golangci-lint run ./...

build-linux-amd64:
	@mkdir -p bin
	@docker buildx build \
		-t vpm-linux-amd64 \
		-f Dockerfile \
		--platform linux/amd64 \
		--load \
		.
	@docker run --name vpm-build-linux vpm-linux-amd64 > /dev/null
	@docker cp vpm-build-linux:/app/vpm ./bin/vpm-linux-amd64
	@docker rm vpm-build-linux > /dev/null
