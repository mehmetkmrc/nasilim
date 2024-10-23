package service

import (
	"errors"

	"github.com/mehmetkmrc/nasilim.git/internal/core/domain/model"
	"github.com/mehmetkmrc/nasilim.git/internal/core/port/db"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo db.UserRepository
}


func (s *UserService) Signup(user *model.User) error{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("password hashing failed")
	}
	user.Password = string(hashedPassword)

	//Save into user database
	return s.UserRepo.Create(user)
}
