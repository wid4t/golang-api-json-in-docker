package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {

	app := fiber.New()

	api := app.Group("/module/user", logger.New())

	api.Get("/check", func(c *fiber.Ctx) error {

		var result = Response{
			Status:  "OK",
			Message: "Hello, i'm from golang-api-json-in-docker",
		}

		bytes, _ := json.Marshal(result)

		return c.Send(bytes)
	})

	log.Fatal(app.Listen(":3000"))
}
