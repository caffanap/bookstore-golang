package controllers

import (
	"caffanap/bookstore/api/models"
	"caffanap/bookstore/api/repositories"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	bookRepository *repositories.BookRepository
}

func NewBookController() *BookController {
	return &BookController{
		bookRepository: repositories.NewBookRepository(),
	}
}

func (controller BookController) Index(c *fiber.Ctx) error {
	var books = controller.bookRepository.GetAllBooks()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Books Listed Successfully!",
		"data":    books,
	})
}

func (controller BookController) Show(c *fiber.Ctx) error {
	id := c.Params("id")
	var book, err = controller.bookRepository.GetOneBook(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Books found!",
		"data":    book,
	})
}

func (controller BookController) Create(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := controller.bookRepository.CreateBook(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully Created Book!",
		"data":    book,
	})
}

func (controller BookController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	newBook, err := controller.bookRepository.UpdateBook(id, &book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully Updated Book!",
		"data":    newBook,
	})
}

func (controller BookController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	book, err := controller.bookRepository.DeleteBook(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully Deleted Book!",
		"data":    book,
	})
}
