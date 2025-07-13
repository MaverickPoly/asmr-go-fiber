package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type WeatherData struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
		Pressure int     `json:"pressure"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Visibility int    `json:"visibility"`
	Name       string `json:"name"`
	Timezone   int    `json:"timezone"`
	ID         int    `json:"id"`
}

func IndexHandler(c *fiber.Ctx) error {
	return c.SendFile("./templates/index.html")
}

func CityWeather(c *fiber.Ctx) error {
	cityName := c.Params("city")
	API_KEY := os.Getenv("API_KEY")

	client := resty.New()

	res, err := client.R().
		Post(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&appid=%v", cityName, API_KEY))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error fetching weather: %v", err.Error()),
		})
	}

	var weatherData WeatherData
	if err := json.Unmarshal(res.Body(), &weatherData); err != nil {
		log.Error("Error parsing weather response:", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error parsing weather response.",
		})
	}

	// return c.SendString(string(res.Body()))
	return c.JSON(fiber.Map{
		"temp":     weatherData.Main.Temp,
		"humidity": weatherData.Main.Humidity,
		"pressure": weatherData.Main.Pressure,

		"description": weatherData.Weather[0].Description,
		"icon":        weatherData.Weather[0].Icon,

		"wind_speed": weatherData.Wind.Speed,
		"wind_deg":   weatherData.Wind.Deg,

		"visibility": weatherData.Visibility,
		"name":       weatherData.Name,
		"timezone":   weatherData.Timezone,
		"id":         weatherData.ID,
	})
}

/*
POST REQUEST EXAMPLE:


	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"username": "maverick",
			"email":    "maverick@example.com",
		}).
		Post("https://example.com/api/user")

	if err != nil {
		panic(err)
	}

	fmt.Println("Status:", resp.Status())
	fmt.Println("Response:", resp.String())
*/
