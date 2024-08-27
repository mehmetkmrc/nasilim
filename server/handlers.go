package server

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
	User        User `json:"user"`
}

func (server *Server) login(c *fiber.Ctx){
	var req loginRequest
	if err := (*c).BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	for _, user := range users{
		if user.Username == req.Username {
			if user.Password == req.Password{
				accessToken, err := server.tokenMaker.CreateToken(req.Username, time.Minute)
			if err != nil {
				return (*c).Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			rsp := loginResponse{
				AccessToken: accessToken,
				User: user,
			}
			return (*c).Status(http.StatusOK).JSON(rsp)
			}
			return (*c).Status(http.StatusForbidden).JSON(fiber.Map{"error":"Incorrect password"})
		}
	}
	return (*c).Status(http.StatusNotFound).JSON(fiber.Map{"error":"User not found"})
}

type createUserRequest struct{
	Username string `json:"username" binding:"required"`
	Email	 string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) createUser(c *fiber.Ctx){
	var user User

	if err := (*c).BodyParser(&user); err !=nil{
		return (*c).Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	id, _ := rand.Int(rand.Reader, big.NewInt(1000))
	user.ID = strconv.Itoa(int(id.Int64()))
	users = append(users, user)

	return (*c).Status(http.StatusOK).JSON(users)
}

type deleteUserRequest struct{
	ID string `params:"id" validate:"required"`
}

func (server *Server) deleteUser(c *fiber.Ctx) error{
	id := (*c).Params("id")

	
	for idx, user := range users {
		if user.ID == id{
			users = append(users[:idx], users[idx+1:]... )
			return (*c).Status(http.StatusOK).JSON(users)
		}
	}
	return (*c).Status(http.StatusNotFound).JSON(fiber.Map{"error":"User not found"})
}

