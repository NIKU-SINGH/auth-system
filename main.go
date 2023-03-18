package main

import (
    "github.com/gofiber/fiber/v2"
	// "fmt"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)


// Schema 
type User struct {
    ID       uint   `gorm:"primary_key"`
    Name     string `gorm:"not null"`
    Email    string `gorm:"not null;unique"`
    Password string `gorm:"not null"`
	Role     bool 	`gorm:"not null"`
}


func main() {
	// Loading the env variables
	err := godotenv.Load()
	if err != nil {
    	log.Fatal("Error loading .env file")
	}

	URL := os.Getenv("DB_URL")

	// Create a new fiber app
    app := fiber.New()

	// Connet to databse and create a table
	dsn := URL
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Create the schema
    db.AutoMigrate(&User{})


	// SignUp Routes
	
	app.Post("api/auth/signup", func(c *fiber.Ctx) error {
		user := new(User)
        err := c.BodyParser(&user)

		if err != nil {
            return c.Status(400).SendString("Invalid Input")
        }
		result := db.Create(&user)
		// check for errors
			if result.Error != nil {
			return c.JSON(fiber.Map{"error": result.Error})
		}	
		return c.JSON(fiber.Map{"data": user})
	})

	// Login Routes

	// Get All Users Route
	app.Get("api/getallusers",func(c *fiber.Ctx) error{
		var users []User
		result := db.Find(&users)
		if result.Error != nil {
			return c.JSON(fiber.Map{"error": result.Error})
        }
		return c.JSON(fiber.Map{"error":result})
	})

	// Get Single User
	app.Get("api/getuser/:id",func(c *fiber.Ctx) error{
		var user User
        result := db.First(&user, c.Params("id"))
        if result.Error!= nil {
            return c.JSON(fiber.Map{"error": result.Error})
        }
        return c.JSON(fiber.Map{"data": user})
	})

	// Delete User
	app.Delete("api/deleteuser/:id",func(c *fiber.Ctx) error{
        var user User
        result := db.Delete(&user,c.Params("id"))
        if result.Error!=nil {
            return c.JSON(fiber.Map{"error": result.Error})
        }
        return c.JSON(fiber.Map{"data": user})
    })

    // Update User


    // Define a route for GET /users that returns a list of all users
    // app.Get("/users", func(c *fiber.Ctx) error {
    //     return c.JSON(users)
    // })

    // Define a route for GET /users/:id that returns a single user by ID
    // app.Get("/users/:id", func(c *fiber.Ctx) error {
    //     id := c.Params("id")
    //     for _, user := range users {
    //         if fmt.Sprint(user.ID) == id {
    //             return c.JSON(user)
    //         }
    //     }
    //     return c.Status(404).SendString("User not found")
    // })

    // Define a route for POST /users that creates a new user
    app.Post("/users", func(c *fiber.Ctx) error {
        var user User
        err := c.BodyParser(&user)
        if err != nil {
            return c.Status(400).SendString("Invalid request body")
        }
        // users = append(users, user)
        return c.Status(201).JSON(user)
    })

    // Define a route for PUT /users/:id that updates a user by ID
    // app.Put("/users/:id", func(c *fiber.Ctx) error {
    //     id := c.Params("id")
    //     var updatedUser User
    //     err := c.BodyParser(&updatedUser)
    //     if err != nil {
    //         return c.Status(400).SendString("Invalid request body")
    //     }
    //     for i, user := range users {
    //         if fmt.Sprint(user.ID) == id {
    //             users[i] = updatedUser
    //             return c.Status(200).JSON(updatedUser)
    //         }
    //     }
    //     return c.Status(404).SendString("User not found")
    // })

    // Define a route for DELETE /users/:id that deletes a user by ID
    // app.Delete("/users/:id", func(c *fiber.Ctx) error {
    //     id := c.Params("id")
    //     for i, user := range users {
    //         if fmt.Sprint(user.ID) == id {
    //             users = append(users[:i], users[i+1:]...)
    //             return c.Status(204).SendString("")
    //         }
    //     }
    //     return c.Status(404).SendString("User not found")
    // })

    // Start the server on port 8080
    app.Listen(":8080")
}