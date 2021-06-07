migrate_up:
	migrate -path ./pkg/migrations -database 'postgres://postgres:admin@localhost:5436/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./pkg/migrations -database 'postgres://postgres:admin@localhost:5436/postgres?sslmode=disable' down

linterCheck:
	golangci-lint run --config .golangci.yml