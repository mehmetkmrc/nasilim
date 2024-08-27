package server

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/mehmetkmrc/nasilim.git/token"
)

const (
	authorizationHeaderKey       = "Authorization"
	authorizationHeaderBearerType = "bearer"
)

func authMiddleware(maker token.PasetoMaker) fiber.Handler{
	return func(c fiber.Ctx) error {
		authHeader := c.Get(authorizationHeaderKey)
		if authHeader == "" {
			c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "No Header was passed"})
		}
		fields := strings.Fields(authHeader)
		if len(fields) != 2 {
			c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid or Missing Bearer Token"})
		}
		authType := fields [0]
		if strings.ToLower(authType) != authorizationHeaderBearerType{
			c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization Type Not Supported"})
		}
		token := fields[1]
		_, err := maker.VerifyToken(token)
		if err != nil{
			c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Access Token Not Valid"})
			
		}
		return c.Next()


	}
}