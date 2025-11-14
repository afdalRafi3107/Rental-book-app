package handler

import (
	"rental-book/internal/entity"
	"rental-book/internal/service"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(s *service.BookService) *BookHandler {
	return &BookHandler{s}
}

func (h *BookHandler) Create(c *fiber.Ctx) error {
	var input entity.Book
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	err := h.service.Create(&input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to save"})
	}

	return c.JSON(input)
}

func (h *BookHandler) List(c *fiber.Ctx) error {
	books, _ := h.service.GetAll()
	return c.JSON(books)
}
