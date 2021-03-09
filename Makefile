.PHONY: build
build: 
	go build .

.PHONY: migrate
migrate:
	go run db/migrate.go