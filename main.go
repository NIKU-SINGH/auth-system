package main 	

import (
	"github.com/gofiber/fiber/v2")

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to DB
	// database.ConnectDB()

	// GET
	app.Get("/api/auth/user/:id",func(c *fiber.Ctx) error{
		err := c.SendString("User info is")
		return err
	})
	app.Post("/api/auth/signup", func(c *fiber.Ctx) error {
		err := c.SendString("User")
		return err
	})

	// Listen on port 8000
	app.Listen(":8000")
}