# Go - Backend Development Task

Build a RESTful API using Go to manage users with their name and dob (date of birth). The API should calculate and return a user’s age dynamically when fetching user detail

## Tech Stack

- [GoFiber](https://gofiber.io/)
- SQL PostgreSQL + [SQLC](https://sqlc.dev/) (docker)
- [Uber Zap](https://github.com/uber-go/zap) for logging
- [go-playground/validator](https://github.com/go-playground/validator) for input validation

## Run Locally

Clone the project

```bash
  git clone https://github.com/girishsinha/usermanagement.git
```

### create .env file using [example.env](example.env)

make sure the .env file is in root directory

Run docker compose

```bash
  docker-compose up -d
```

Run Migrations

```bash
  go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path db/migrations -database "postgres://{DB_USER}:{DB_PASSWORD}@localhost:5432/{DB_NAME}?sslmode=disable" up
```

### Replace DB_USER, DB_PASSWORD, DB_NAME with your credentials

Start the server

```bash
  go run cmd/server/main.go
```
