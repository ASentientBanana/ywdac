package controllers

import "github.com/gofiber/fiber/v2"

type response struct {
	Data string
}

func GetTorrent(c *fiber.Ctx) error {
	response := response{}
	response.Data = "Some data"
	c.JSON(response)
	return nil
}
