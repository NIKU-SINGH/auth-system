package main

import (
    "github.com/gofiber/fiber/v2"
    "auth-system/routes"
    "auth-system/database"
)


func main() {

    // Create a new Fiber app
    app := fiber.New()

    // Connect to the database
    database.ConnectDB()

    // Register your routes
    routes.SetUpRoutes(app)

    // Start the server on port 8080
    app.Listen(":8080")
}