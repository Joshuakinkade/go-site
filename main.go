package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Configure the app
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/css", "./public/css")

	// Configure routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{"Message": "Hello, World!"}, "layouts/base")
	})
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
