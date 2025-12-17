// package main

// func main() {

// 	// // Initialize logger
// 	// logger, err := zap.NewProduction()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// defer logger.Sync() // flushes buffer, if any

// 	// // Log server start
// 	// logger.Info("Starting Fiber server...")

// 	// app := fiber.New()

// 	// app.Get("/health", func(c *fiber.Ctx) error {
// 	// 	logger.Info("Health check requested")
// 	// 	return c.JSON(fiber.Map{
// 	// 		"status": "ok",
// 	// 	})
// 	// })

// 	// addr := ":3000"
// 	// logger.Info("Listening on address", zap.String("address", addr))

//		// if err := app.Listen(addr); err != nil {
//		// 	logger.Fatal("Failed to start server", zap.Error(err))
//		// }
//	}
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	db "github.com/girishsinha/user-manage/db/sqlc"
	"github.com/girishsinha/user-manage/internal/handler"
	"github.com/girishsinha/user-manage/internal/repository"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/girishsinha/user-manage/internal/service"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	if err := godotenv.Load(); err != nil {
		logger.Fatal("No .env file found, using system environment")
	}
	// 1. Database Connection
	// connStr := "postgres://myuser:girish@localhost:5432/mydatabase?sslmode=disable"
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatal("Database connection failed", zap.String("address", err.Error()))
	}

	// 2. Dependency Injection (Wiring)
	queries := db.New(dbConn)
	repo := repository.NewUserRepository(queries)
	svc := service.NewUserService(repo)
	userH := handler.NewUserHandler(svc)

	// 3. Initialize Fiber
	app := fiber.New()

	// 4. Routes
	app.Post("/users", userH.CreateUser)
	app.Get("/users", userH.ListUsers)
	app.Get("/users/:id", userH.GetUser)
	app.Put("/users/:id", userH.UpdateUser)
	app.Delete("/users/:id", userH.DeleteUser)

	// log.Fatal(app.Listen(":8000"))
	port := os.Getenv("APP_PORT")
	log.Fatal(app.Listen(":" + port))
}
