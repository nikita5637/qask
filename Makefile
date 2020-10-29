build:
	go build -v ./cmd/qask

run:
	go run ./cmd/qask --config ./configs/qask_develop.conf

test:
	go test -v -race ./...

clean:
	rm qask

.DEFAULT_GOAL := run
