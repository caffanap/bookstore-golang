package routers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleRouter(app *fiber.App) {
	router := app.Group("/api/v1")

	// Register router
	BooksRoutes(router)
}
