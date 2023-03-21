package database

import(
	"fmt"
	"log"
	// "strconv"
	"os"
	"time"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "github.com/gofiber/fiber/v2"
	// "auth-system/models"
)

// Schema 
type User struct {
    ID           uint   `gorm:"primary_key"`
    Username     string `gorm:"not null"`
    Email        string `gorm:"not null;unique"`
    Password     string `gorm:"not null"`
	Role         string `gorm:"not null;default:user"`
	CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// Variable for the Database
var DB *gorm.DB

// Function to connect with DB
func ConnectDB() *gorm.DB {

	// userSchema := models.User{}

	// Loading the env variables
	err := godotenv.Load()
	if err != nil {
    	log.Fatal("Error loading .env file")
	}

	URL := os.Getenv("DB_URL")

	// Connet to databse and create a table
	dsn := URL
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    } else {fmt.Printf("database connection established")}

    // Create the schema
    db.AutoMigrate(&User{})

	return db

}