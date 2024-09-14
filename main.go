package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joshuakinkade/go-site/handlers"
	"github.com/joshuakinkade/go-site/services"
)

func main() {
	// Configure the app
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/css", "./public/css")

	pageHandler := handlers.NewPagesHandler(services.NewPostService())

	// Configure routes
	app.Get("/", pageHandler.ShowHome)
	app.Get("/posts", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Posts"})
	})
	app.Get("/posts/:slug", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": c.Params("slug")})
	})

	app.Get("/api/v1/posts", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Posts"})
	})

	log.Panic(app.Listen(":8080"))
}
