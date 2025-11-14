package routes

import (
	"rental-book/internal/config"
	"rental-book/internal/handler"
	"rental-book/internal/repository"
	"rental-book/internal/service"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r *fiber.App){
	db:=config.ConnectDB()

	authRepo := repository.NewUserRpository(db)
	authService := service.NewUserService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	r.Post("/auth/register", authHandler.Register)
	r.Post("/auth/login", authHandler.Login)

}