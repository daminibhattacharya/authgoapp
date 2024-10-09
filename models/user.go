package models

import (
	"time"
)

type User struct {
	ID        string 
	FirstName string     `json:"firstname" validate:"required"`
	LastName  string     `json:"lastname" validate:"required"`
	Email     string     `json:"email" validate:"required"`
	Password  string     `json:"password" validate:"required"`
	CreatedAt time.Time
}