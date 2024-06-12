package router

import (
	"main/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/page", handler.GetPage)
	app.Get("/get/:key", handler.GetPageByKey)
	// app.Get("/page/:id", handler.GetPage)
	app.Post("/page", handler.CreatePage)
	app.Patch("/page", handler.UpdatePage)
	app.Delete("/delete/:key", handler.DeletePageByKey)
}
