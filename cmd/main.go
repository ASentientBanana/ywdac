package main

import (
	"fmt"

	"github.com/asentientbanana/ywdac/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api.Init(app)

	app.Listen(":8888")
	fmt.Println("Hello world")
}
