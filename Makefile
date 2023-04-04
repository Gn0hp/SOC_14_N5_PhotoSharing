install:
	echo Download go.mod dependencies
	go mod download

test:
	echo Run all test files in directories
	go run test.go

migrate:
	echo Migrate database schema
	go run cli/db_seed.go