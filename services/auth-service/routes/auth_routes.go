package routes

import (
	"auth-service/internal/handler"
	"auth-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, authHandler *handler.AuthHandler) {
	api := app.Group("/api/auth")

	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)
	api.Get("/me", middleware.AuthMiddleware(), authHandler.Me)
}
