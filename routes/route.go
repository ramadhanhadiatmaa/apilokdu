package routes

import (
	"apilokdu/controllers"
	"apilokdu/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Routes(r *fiber.App) {
	ant := r.Group("/api")

	ant.Get("/", middlewares.AuthMiddleware, controllers.Index)
	ant.Post("/", middlewares.AuthMiddleware, controllers.Create)
	ant.Get("/:id", middlewares.AuthMiddleware, controllers.Show)
	ant.Put("/:id", middlewares.AuthMiddleware, controllers.Update)
	ant.Put("/reset/:id", middlewares.AuthMiddleware, controllers.Reset)
}