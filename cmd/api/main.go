package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/priyansh-garg1/ticket-verse/config"
	"github.com/priyansh-garg1/ticket-verse/db"
	"github.com/priyansh-garg1/ticket-verse/handelers"
	"github.com/priyansh-garg1/ticket-verse/repositories"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig)

	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Initialize repositories
	eventRepository := repositories.NewEventRepository(nil)

	server := app.Group("/api")

	// Initialize handlers
	handelers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
