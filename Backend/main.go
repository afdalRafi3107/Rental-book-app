package main

import (
	"RentalBook/internal/config"
	"RentalBook/internal/infratructure/http"
	"log"
)

func main() {
	cfg := config.LoadEnv()

	app:= http.NewFiberApp(cfg)
	log.Println("Server is Running on port: ", cfg.Port)
	log.Fatal(app.Listen(":"+cfg.Port))
}