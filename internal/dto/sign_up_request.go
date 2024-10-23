package dto

import "time"

type SignupRequest struct {
	Email     string `json:"email" validate:"required, email"`
	Password  string `json:"password" validate:"required, min=8"`
	FirstName string `json:"first_name" validate:"required"`
	BirthDate time.Time `json:"birth_date" validate:"required"`
}