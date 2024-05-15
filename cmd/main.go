package main

import (
	"gobackend/pkg/common/database"
	"gobackend/pkg/core/service"
	"gobackend/pkg/handlers"
	"gobackend/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := database.Connect()

	repo := repository.NewRepo(db)

	srv := service.NewService(repo)

	handler := handlers.NewHandler(srv)
	api := app.Group("/api/v1")
	api.Get("/posts", handler.GetPosts)
	api.Get("/posts/:id", handler.GetPostID)
	api.Post("/posts", handler.CreatePosts)
	api.Put("/posts/:id", handler.UpdatePost)
	api.Delete("/posts/:id", handler.DeletePost)
	app.Listen(":8080")
}
