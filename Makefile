build:
	go build -v ./cmd/qask

run:
	go run ./cmd/qask

test:
	go test -v -race ./...
clean:

.DEFAULT_GOAL := run
