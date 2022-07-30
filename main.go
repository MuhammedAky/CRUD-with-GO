package main

import (
	"github.com/MuhammedAky/crud-go-react/bootstrap"
	"github.com/MuhammedAky/crud-go-react/repository"
	"github.com/gofiber/fiber/v2"
)

type Repository repository.Repository

func main() {
	app := fiber.New()
	bootstrap.InitializeApp(app)
}