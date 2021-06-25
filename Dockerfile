FROM golang:1.13

RUN ["go", "get", "-tags", "postgres,mysql", "-v", "-u", "github.com/golang-migrate/migrate/cmd/migrate"]