package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jackc/pgx/v5"
	"github.com/joshuakinkade/go-site/db"
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

	dbConn, err := pgx.Connect(context.Background(), "postgres://user:secret@postgres:5432/go_site")
	if err != nil {
		log.Panic(err)
	}

	// Intialize repositories
	postsRepo := db.NewPosts(dbConn)

	// Initialize services
	postsService := services.NewPostService(postsRepo)

	// Initialize handlers
	pageHandler := handlers.NewPagesHandler(postsService)
	apiHandler := handlers.NewAPIHandler(postsService)

	// Configure routes
	app.Get("/", pageHandler.ShowHome)
	app.Get("/posts", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Posts"})
	})
	app.Get("/posts/:slug", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": c.Params("slug")})
	})

	app.Get("/api/v1/posts", apiHandler.ListPosts)
	app.Post("/api/v1/posts", apiHandler.CreatePost)

	log.Panic(app.Listen(":8080"))
}
