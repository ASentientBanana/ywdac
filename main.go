package main

import (
	"fmt"

	"github.com/asentientbanana/ywdac/server"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	server.Init(app)

	app.Listen(":8888")
	fmt.Println("Hello world")
}
