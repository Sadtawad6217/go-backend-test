package main

import (
	"go-backend-test/pkg/common/database"
	"go-backend-test/pkg/core"
	"go-backend-test/pkg/handlers"
	"go-backend-test/pkg/repository"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to the database
	err := database.ConnectDB("postgres://posts:38S2GPNZut4Tmvan@dev.opensource-technology.com:5523/posts?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.DB.Close()

	// Initialize repositories
	examRepo := repository.NewExamRepository()

	// Initialize services
	examService := core.NewExamService(examRepo)

	// Initialize handlers
	examHandler := handlers.NewExamHandler(examService)

	// Setup Fiber app
	app := fiber.New()

	// Routes
	app.Get("/exams/:id", examHandler.GetExamByID)
	app.Get("/exams", examHandler.GetAllExams)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
