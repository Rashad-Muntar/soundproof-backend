package config

import (
	"os"
	"github.com/Rashad-Muntar/soundproof/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})
	DB = db
}

