package models

type User struct {
	ID uint `json:"id"`
	Name string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique not null"`
	Password string `json:"password" gore:"not null"`
}