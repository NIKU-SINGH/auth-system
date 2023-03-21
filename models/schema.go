package models

import (
    "time"
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