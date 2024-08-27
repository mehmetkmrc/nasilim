package server

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mehmetkmrc/nasilim.git/token"
)

type Server struct {
	tokenMaker *token.PasetoMaker
	router		*fiber.Route}


	func NewServer(address string) (*Server, error) {
		tokenMaker, err := token.NewPaseto("abcdefghijkl12345678901234567890")
		if err != nil{
			return nil, fmt.Errorf("Could not create token maker: %w", err)
		}

		server := &Server{
			tokenMaker: tokenMaker,
		}

		server.setRoutes()
		server.router.Run(address)

		return server, nil

	}

	func (server *Server) setRoutes(){
		router := fiber.New()

		auth := router.Group("/").Use(authMiddleware(*server.tokenMaker))
		auth.Delete("/delete/:id", server.deleteUser())
		router.Post("/create", server.createUser)
		router.Post("/login", server.login)

		server.router = router 
	}