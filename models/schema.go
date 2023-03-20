package models

// Schema 
type User struct {
    ID           uint   `gorm:"primary_key"`
    Username     string `gorm:"not null"`
    Email        string `gorm:"not null;unique"`
    Password     string `gorm:"not null"`
	Role         bool 	`gorm:"not null"`
}