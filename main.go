package main

import (
	"myGo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func main() {
	// template engine
	engine := handlebars.New("./views", ".hbs")

	// configs
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	// routes
	routes.UseRoute(app)

	// Server
	app.Listen(":3000")
}
