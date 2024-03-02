package main

import (
	"fmt"

	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world.")
	})
	app.Get("/greetings", func(c *fiber.Ctx) error {
		var name = c.Query("name")
		return c.SendString(fmt.Sprintf("Hello, %s", name))
	})
	log.Fatal(app.Listen(":3000"))
}
