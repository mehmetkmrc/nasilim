package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/mehmetkmrc/nasilim.git/internal/adapter/secondary/storage/postgres"
	
)

func main() {
	app := fiber.New()

	connString := "postgres://username:password@localhost:5432/nasilim_1?sslmode=disable"
	postgres.NewPostgresUserRepository(connString)
	log.Fatal(app.Listen(":3000"))

}