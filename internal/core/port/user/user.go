package user

import "github.com/mehmetkmrc/nasilim.git/internal/core/domain/model"

type UserRepository interface {
	Create(user *model.User) error
}