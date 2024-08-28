package server

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mehmetkmrc/nasilim.git/token"
)

type Server struct {
	tokenMaker *token.PasetoMaker
	router		*fiber.App
}


	func NewServer(address string) (*Server, error) {
		tokenMaker, err := token.NewPaseto("abcdefghijkl12345678901234567890")
		if err != nil{
			return nil, fmt.Errorf("Could not create token maker: %w", err)
		}

		server := &Server{
			tokenMaker: tokenMaker,
		}

		server.setRoutes()
		server.router.Listen(address)

		return server, nil

	}

	