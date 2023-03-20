package routes

import (
	"auth-system/controllers"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func SetUpRoutes(app *fiber.App) {
    app.Post("api/auth/signup", controller.CreateUserHandler)
    app.Post("api/auth/login", controller.LoginHandler)
    app.Get("/api/getuser/:id", controller.GetUserHandler)
    app.Delete("/api/deleteuser/:id", controller.DeleteUserHandler)
    app.Get("/api/getallusers", controller.GetAllUsersHandler)
	fmt.Printf("hey")
}

