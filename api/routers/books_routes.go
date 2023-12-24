package routers

import (
	"caffanap/bookstore/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func BooksRoutes(router fiber.Router) {
	router_books := router.Group("/books")
	controllers := controllers.NewBookController()

	router_books.Get("/", controllers.Index)
	router_books.Post("/", controllers.Create)
	router_books.Get("/:id", controllers.Show)
	router_books.Delete("/:id", controllers.Delete)
	router_books.Put("/:id", controllers.Update)
}
