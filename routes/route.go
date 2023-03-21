package routes

import (
	"auth-system/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
    app.Post("api/auth/signup", controller.VerifyAdmin,controller.CreateUserHandler)
    app.Post("api/auth/login", controller.LoginHandler)
    app.Get("/api/getuser/:id", controller.GetUserHandler)
    app.Delete("/api/deleteuser/:id", controller.VerifyAdmin,controller.DeleteUserHandler)
    app.Get("/api/getallusers", controller.GetAllUsersHandler)
}

