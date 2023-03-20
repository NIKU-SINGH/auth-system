package main

import (
    "github.com/gofiber/fiber/v2"
	// "fmt"
	// "gorm.io/driver/postgres"
    // "gorm.io/gorm"
	// "github.com/joho/godotenv"
	// "log"
	// "os"
    // "fmt"
    // "time"
    // "github.com/golang-jwt/jwt/v5"
    // "golang.org/x/crypto/bcrypt"
    "auth-system/routes"
    "auth-system/database"
)


// Schema 
// type User struct {
//     ID           uint   `gorm:"primary_key"`
//     Username     string `gorm:"not null"`
//     Email        string `gorm:"not null;unique"`
//     Password     string `gorm:"not null"`
// 	Role         bool 	`gorm:"not null"`
// }

// Claims for the token

// type MyCustomClaims struct {
// 	Foo string `json:"foo"`
// 	jwt.RegisteredClaims
// }


func main() {
    // mySigningKey := []byte("AllYourBase")

	// Loading the env variables
	// err := godotenv.Load()
	// if err != nil {
    // 	log.Fatal("Error loading .env file")
	// }

	// URL := os.Getenv("DB_URL")

	// // Create a new fiber app
    app := fiber.New()

	// // Connet to databse and create a table
	// dsn := URL
    // db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    // if err != nil {
    //     panic("failed to connect database")
    // } else {fmt.Printf("database connection established")}

    // Create the schema
    // db.AutoMigrate(&User{})


	// SignUp Routes
	
	// app.Post("api/auth/signup", func(c *fiber.Ctx) error {
	// 	user := new(User)
    //     err := c.BodyParser(&user)

	// 	if err != nil {
    //         return c.Status(400).SendString("Invalid Input")
    //     }

    //     // Hashing the passowrd
    //     hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    //     if err != nil {
    //         return err
    //     }

    //     user.Password = string(hashedPassword)

	// 	result := db.Create(&user)
	// 	// check for errors
	// 		if result.Error != nil {
	// 		return c.JSON(fiber.Map{"error": result.Error})
	// 	}	
	// 	return c.JSON(fiber.Map{"User created": user})
	// })

	// Login Routes
    // app.Post("api/auth/login",func(c *fiber.Ctx) error {

    //     var user User
    //     // Parsing the body request
    //     if err := c.BodyParser(&user); err != nil {
    //         return c.Status(400).SendString("Invalid Input")
    //     }

    //      // Find user in database
    //     var dbUser User

    //     if err:=db.Where("username = ?",user.Username).Find(&dbUser).Error; err != nil{
    //         // if err != gorm.ErrRecordNotFound {
    //         // }
    //         return fiber.NewError(fiber.StatusUnauthorized, "User not found, please sign up")
    //     }

    //     // Comparing the passwords
    //     if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password),[]byte(user.Password)); err != nil {
    //         return fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password")
    //     }

    //     // Create claims while leaving out some of the optional fields
    //     claims := MyCustomClaims{
	//             "bar",
	//             jwt.RegisteredClaims{
	// 	// Also fixed dates can be used for the NumericDate
	// 	        ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
	// 	        Issuer:    "test",
	//             },
    //     }

    //     token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    //     tokenString, err := token.SignedString(mySigningKey)

    //     // Send JWT token in a cookie
    //     c.Cookie(&fiber.Cookie{
    //         Name:  "jwt",
    //         Value: tokenString,
    //         Expires:  time.Now().Add(time.Hour * 24),
    //         HTTPOnly: true,
    //     })

    //     return c.JSON(fiber.Map{"User token is": tokenString})
    //     return err

    // })

	// Get All Users Route
	// app.Get("api/getallusers",func(c *fiber.Ctx) error{
	// 	var users []User
	// 	result := db.Find(&users)
	// 	if result.Error != nil {
	// 		return c.JSON(fiber.Map{"error": result.Error})
    //     }
	// 	return c.JSON(fiber.Map{"data":result})
	// })

	// Get Single User
	// app.Get("api/getuser/:id",func(c *fiber.Ctx) error{
	// 	var user User
    //     result := db.First(&user, c.Params("id"))
    //     if result.Error!= nil {
    //         return c.JSON(fiber.Map{"User not found": result.Error})
    //     }
    //     return c.JSON(fiber.Map{"data": user})
	// })

	// Delete User
	// app.Delete("api/deleteuser/:id",func(c *fiber.Ctx) error{
    //     var user User
    //     // Check if user is there or not
    //     if err := db.Where("id = ?", c.Params("id")).First(&user).Error; err != nil {
    //         return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    //     }

    //     // Delete user
    //     result := db.Delete(&user,c.Params("id"))
    //     if result.Error!=nil {
    //         return c.JSON(fiber.Map{"error": result.Error})
    //     }
    //     return c.JSON(fiber.Map{"data": user})
    // })

    // Update User


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
    // app.Post("/users", func(c *fiber.Ctx) error {
    //     var user User
    //     err := c.BodyParser(&user)
    //     if err != nil {
    //         return c.Status(400).SendString("Invalid request body")
    //     }
    //     // users = append(users, user)
    //     return c.Status(201).JSON(user)
    // })

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
    
    // Connect to the database
    database.ConnectDB()

    // Register your routes
    routes.SetUpRoutes(app)

    // Start the server on port 8080
    app.Listen(":8080")
}