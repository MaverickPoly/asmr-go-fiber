package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func HandleRegister(c *fiber.Ctx) error {
	user := new(User)

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

	result := DB.Create(user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error creating user: %v", result.Error.Error()),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully!",
	})
}

func HandleLogin(c *fiber.Ctx) error {
	user := new(User)

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

	dbUser := new(User)
	DB.Where("username = ?", user.Username).First(dbUser)

	if dbUser.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found!",
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

	fmt.Println("Claims:", claims)

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(JWT_SECRET)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error creating token: %v", err.Error()),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		HTTPOnly: !c.IsFromLocal(),
		Secure:   !c.IsFromLocal(),
		MaxAge:   60 * 60 * 24 * 7, // 7 days
	})

	return c.JSON(fiber.Map{
		"message":     "Logged in successfully!",
		"accessToken": accessToken,
	})
}

func HandleLogout(c *fiber.Ctx) error {
	c.ClearCookie("accessToken")
	return c.JSON(fiber.Map{"message": "Logged in successfully!"})
}

func HandleGetUser(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	user := new(User)

	result := DB.First(&user, userId)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to fetch user: %v", result.Error.Error()),
		})
	}

	return c.JSON(user)
}
