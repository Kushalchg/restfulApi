build:
	go build -o bin/main main.go
migration:
	go run migrate/migrate.go
run:
	go run main.go
