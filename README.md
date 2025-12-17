go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path db/migrations -database "postgres://myuser:girish@localhost:5432/mydatabase?sslmode=disable" up
