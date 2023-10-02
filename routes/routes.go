package routes

import (
	"go_rest_api/handler"
	"go_rest_api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// public routes
	var publicRoutes fiber.Router = app.Group("/api/v1")

	publicRoutes.Post("/signup", handler.Signup)
	publicRoutes.Post("/login", handler.Login)
	publicRoutes.Get("/items", handler.GetAllItems)
	publicRoutes.Get("/items/:id", handler.GetItemByID)

	// private routes, authentication is required
	var privateRoutes fiber.Router = app.Group("/api/v1", middlewares.CreateMiddleware())

	privateRoutes.Post("/items", handler.CreateItem)
	privateRoutes.Put("/items/:id", handler.UpdateItem)
	privateRoutes.Delete("/items/:id", handler.DeleteItem)
}