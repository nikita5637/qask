build:
	go build -v ./cmd/qask

run:
	go run ./cmd/qask --config ./configs/qask_develop.conf

test:
	go test -v -race ./...

migrate-mysql: 
	migrate -source file://internal/app/migrations/mysql -database 'mysql://$(QASK_MYSQL_DATABASE_USER):$(QASK_MYSQL_DATABASE_PASSWORD)@tcp($(QASK_MYSQL_DATABASE_ADDRESS):$(QASK_MYSQL_DATABASE_PORT))/$(QASK_MYSQL_DATABASE_DB)' up 
	migrate -source file://internal/app/migrations/mysql -database 'mysql://$(QASK_MYSQL_DATABASE_USER):$(QASK_MYSQL_DATABASE_PASSWORD)@tcp($(QASK_MYSQL_DATABASE_ADDRESS):$(QASK_MYSQL_DATABASE_PORT))/$(QASK_MYSQL_DATABASE_DB_TEST)' up 

clean:
	rm qask

.DEFAULT_GOAL := run
