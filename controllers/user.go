package controller

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"fmt"
	// "log"
	"auth-system/database"
	// "gorm.io/gorm"
	// "gorm.io/driver/postgres"

)

type User struct {
    ID           uint   `gorm:"primary_key"`
    Username     string `gorm:"not null;unique"`
    Email        string `gorm:"not null;unique"`
    Password     string `gorm:"not null"`
	Role         string `gorm:"not null;default:user"`
	CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// Claims for the token

// type MyCustomClaims struct {
// 	Foo string `json:"foo"`
// 	jwt.RegisteredClaims
// }

type Claims struct {
    Username string `json:"username"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

// Import the DB object and initialize it
var db  = database.ConnectDB()

// Verify Admin
func VerifyAdmin(c *fiber.Ctx) error {
	// Gettig the JWT token
	cookie :=c.Cookies("jwt")
	mySigningKey := []byte("AllYourBase")


	// Parsing the JWT token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
        // check signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, 
			fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        // return secret key for signature validation
        return mySigningKey, nil
    })
	claims := token.Claims.(jwt.MapClaims)
	if claims["role"] == "admin" {
		fmt.Printf("Verified admin")
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized user, can only read the data"})
	}
	
	return c.JSON(fiber.Map{"Error while parsing the token": err})
}

// CreateUserHandler
func CreateUserHandler(c *fiber.Ctx) error {
    user := new(User)
        err := c.BodyParser(&user)

		if err != nil {
            return c.Status(400).SendString("Invalid Input")
        }

        // Hashing the passowrd
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
            return err
        }

        user.Password = string(hashedPassword)

		 // Find user in database
		var dbUser User

		// If user is already registered
		if err:=db.Where("username = ?",user.Username).Find(&dbUser).Error; err != nil{
			// if err != gorm.ErrRecordNotFound {
			// }
			return fiber.NewError(fiber.StatusUnauthorized, "User is already registered")
		}

		result := db.Create(&user)
		// check for errors
			if result.Error != nil {
			return c.JSON(fiber.Map{"error": result.Error})
		}	
		return c.JSON(fiber.Map{"User created": user})
}

// LoginHandler
func LoginHandler(c *fiber.Ctx) error {
	mySigningKey := []byte("AllYourBase")

	var user User
	// Parsing the body request
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Invalid Input")
	}

	 // Find user in database
	var dbUser User

	if err:=db.Where("username = ?",user.Username).Find(&dbUser).Error; err != nil{
		// if err != gorm.ErrRecordNotFound {
		// }
		return fiber.NewError(fiber.StatusUnauthorized, "User not found, please sign up")
	}

	// Comparing the passwords
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password),[]byte(user.Password)); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password")
	}

	// Create claims while leaving out some of the optional fields
	claims := Claims{
		Username: dbUser.Username,
		Role: dbUser.Role,
		RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
            Issuer:    "my-app",
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)

	// Send JWT token in a cookie
	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"User token is": tokenString})
	return err
}


func GetUserHandler(c *fiber.Ctx) error {
	var user User
        result := db.First(&user, c.Params("id"))
        if result.Error!= nil {
            return c.JSON(fiber.Map{"User not found": result.Error})
        }
        return c.JSON(fiber.Map{"data": user})
}

// func UpdateUserHandler(c *fiber.Ctx) error {}

func GetAllUsersHandler(c *fiber.Ctx) error {
	var users []User
		result := db.Find(&users)
		if result.Error != nil {
			return c.JSON(fiber.Map{"error in finding users": result.Error})
        }
		fmt.Printf("list of users is: %v",result)
		return c.JSON(fiber.Map{"Data of all users":result})
}

func DeleteUserHandler(c *fiber.Ctx) error {
	var user User
        // Check if user is there or not
        if err := db.Where("id = ?", c.Params("id")).First(&user).Error; err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
        }

        // Delete user
        result := db.Delete(&user,c.Params("id"))
        if result.Error!=nil {
            return c.JSON(fiber.Map{"error": result.Error})
        }
        return c.JSON(fiber.Map{"deleted user": user})
}
