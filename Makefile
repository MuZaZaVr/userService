migrate_up:
	migrate -path ./pkg/migrations -database 'postgres://admin:admin@localhost:5432/userService?sslmode=disable' up

migrate_down:
	migrate -path ./pkg/migrations -database 'postgres://admin:admin@localhost:5432/userService?sslmode=disable' down
