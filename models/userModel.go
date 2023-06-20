package models

import (
	"gorm.io/gorm"
	// "github.com/ethereum/go-ethereum/common"
)

type User struct {
	gorm.Model
	Name string 
	Email string `gorm:"unique"`
	Public string `gorm:"unique"`
	Password string `gorm:"not null"`
}
