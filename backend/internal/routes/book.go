package routes

import (
	"rental-book/internal/config"
	"rental-book/internal/handler"
	"rental-book/internal/repository"
	"rental-book/internal/service"

	"github.com/gofiber/fiber/v2"
)

func BookRoute(r *fiber.App){
	db:=config.ConnectDB()

	bookRepo:= repository.NewBookRepository(db)
	bookService:=service.NewBookService(bookRepo)
	bookHandler:=handler.NewBookHandler(bookService)
	
	r.Post("/books", bookHandler.Create)
	r.Get("/books", bookHandler.List)
}