package model

import "time"

type User struct {
	Email        string `json:"email" validate:"required"`
	Password     string `json:"password" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	BirthDate time.Time `json:"birth_date" validate:"required"`
}