package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func AuthRequired(c *fiber.Ctx) error {
	tokenCookie := c.Cookies("accessToken")

	if tokenCookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Not authorized",
		})
	}

	// check if token is valid or not
	JWT_SECRET := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenCookie, func(token *jwt.Token) (any, error) {
		if token.Method.Alg() != jwt.GetSigningMethod("HS256").Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_SECRET, nil
	})

	if err != nil || !token.Valid {
		log.Warn("Invalid access Token!")
		c.ClearCookie("accessToken")

		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "Invalid accessToken",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := uint(claims["userId"].(float64))
	var user User

	// check if the user exists in the database
	if err := DB.First(&user, userId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		log.Warn("User not found in the db during auth middleware!")
		c.ClearCookie("accessToken")

		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "User not found!",
		})
	}

	c.Locals("userId", userId)

	return c.Next()
}
