package main

import (
	"log"

	"pelacakan-fruit-transport/config"
	"pelacakan-fruit-transport/database"
	"pelacakan-fruit-transport/handler"
	"pelacakan-fruit-transport/repository"
	"pelacakan-fruit-transport/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// 0. Load .env file (ignore error if file doesn't exist)
	_ = godotenv.Load()

	// 1. Load configuration from environment variables
	cfg := config.LoadConfig()

	// 2. Connect to PostgreSQL database and run auto-migration
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 3. Initialize repository (data access layer)
	transportRepo := repository.NewTransportRepository(db)

	// 4. Initialize handler (HTTP layer)
	transportHandler := handler.NewTransportHandler(transportRepo)

	// 5. Create Fiber app, setup middleware and routes
	app := fiber.New()
	router.SetupRoutes(app, transportHandler)

	// 6. Start HTTP server
	log.Printf("Server starting on port %s", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
