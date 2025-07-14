package middleware

import (
	"errors"
	"fmt"
	"os"

	"13.ExpenseTracker/config"
	"13.ExpenseTracker/models"
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

	// Check if the token is valid or not
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

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid accessToken!",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := uint(claims["userId"].(float64))
	var user models.User

	if err := config.DB.First(&user, userId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		log.Warn("User not found in the db during auth middleware!")
		c.ClearCookie("accessToken")

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not found!",
		})
	}

	c.Locals("userId", userId)

	return c.Next()
}
