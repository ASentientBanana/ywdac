package api

import (
	"github.com/asentientbanana/ywdac/controllers"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {

	//torent
	app.Get("/torrents", controllers.GetTorrent)

}
