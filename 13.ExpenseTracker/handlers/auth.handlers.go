package handlers

import (
	"fmt"
	"os"

	"13.ExpenseTracker/config"
	"13.ExpenseTracker/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegister(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body!",
		})
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Some fields are missing!",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user.Password = string(hashedPassword)

	if err := config.DB.Create(user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error creating user: %v", err.Error()),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully!",
		"data":    user,
	})

}

func HandleLogin(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Body!",
		})
	}

	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Some fields are missing!",
		})
	}

	dbUser := new(models.User)
	config.DB.Where("username = ?", user.Username).First(&dbUser)

	if dbUser.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("User with username %v not found!", user.Username),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password!",
		})
	}

	JWT_SECRET := []byte(os.Getenv("JWT_SECRET"))
	fmt.Println("JWT_SECRET:", JWT_SECRET)
	claims := jwt.MapClaims{
		"userId": dbUser.ID,
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(JWT_SECRET)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error creating access token: %v", err.Error()),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		HTTPOnly: !c.IsFromLocal(),
		Secure:   !c.IsFromLocal(),
		MaxAge:   60 * 60 * 24 * 7,
	})

	return c.JSON(fiber.Map{
		"message":     "Logged in successfully!",
		"accessToken": accessToken,
	})
}

func HandleLogout(c *fiber.Ctx) error {
	c.ClearCookie("accessToken")
	return c.JSON(fiber.Map{"message": "Logged out successfully!"})
}

// Fetch my profile
func GetMe(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	user := new(models.User)

	if err := config.DB.First(&user, userId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to fetch user: %v", err.Error()),
		})
	}

	return c.JSON(user)
}
