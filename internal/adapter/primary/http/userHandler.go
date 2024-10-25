package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mehmetkmrc/nasilim.git/internal/core/domain/model"
	"github.com/mehmetkmrc/nasilim.git/internal/core/service"
	"github.com/mehmetkmrc/nasilim.git/internal/dto"
)

type UserHandler struct {
	UserService *service.UserService
}


func (s *UserHandler) Signup(c fiber.Ctx) error{
	var req dto.SignupRequest

	if err := c.Bind().Body(req) ; err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	user := &model.User{
		Email: req.Email,
		Password: req.Password,
		FirstName: req.FirstName,
		BirthDate: req.BirthDate,
	}
	if err := s.UserService.Signup(user); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user created successfully"})
}