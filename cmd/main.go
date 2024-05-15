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

	app.Get("/posts", handler.GetPosts)
	app.Get("/posts/:id", handler.GetPostID)
	app.Post("/posts", handler.CreatePosts)
	app.Put("/posts/:id", handler.UpdatePost)
	app.Delete("/posts/:id", handler.DeletePost)
	app.Listen(":8080")
}
