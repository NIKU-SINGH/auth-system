package database

import(
	"fmt"
	"log"
	"strconv"
	"github.com/percoguru/notes-api-fiber/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variable for the Database
var DB *gorm.DB

// Function to connect with DB
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port,err := strconv.ParseUint(p,10,32)

	if err != nil {
		log.Println("Error: port number is not a number")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
        panic("failed to connect database")
    }

	fmt.Println("Connection Opened to Database")

}