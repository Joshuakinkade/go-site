package main

import (
	"bytes"
	"context"
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jackc/pgx/v5"
	"github.com/joshuakinkade/go-site/db"
	"github.com/joshuakinkade/go-site/handlers"
	"github.com/joshuakinkade/go-site/services"
	"github.com/yuin/goldmark"
)

func main() {
	// Configure templates
	engine := html.New("./templates", ".html")
	engine.AddFunc("unescape", func(s string) template.HTML {
		return template.HTML(s)
	})
	engine.AddFunc("RenderMarkdown", func(s string) string {
		var buf bytes.Buffer
		goldmark.Convert([]byte(s), &buf)
		return buf.String()
	})
	engine.AddFunc("StringSlice", func(s string, start, end int) string {
		if end > len(s) {
			end = len(s)
		}
		return s[start:end]
	})

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

	// Configure app
	app := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: pageHandler.ShowError,
	})

	// configure routes
	app.Static("/css", "./public/css")
	app.Get("/", pageHandler.ShowHome)
	app.Get("/posts", pageHandler.ShowPostList)
	app.Get("/posts/:slug", pageHandler.ShowPost)
	app.Get("/api/v1/posts", apiHandler.ListPosts)
	app.Get("/api/v1/posts/:slug", apiHandler.GetPost)
	app.Post("/api/v1/posts", apiHandler.CreatePost)
	app.Patch("/api/v1/posts/:slug", apiHandler.UpdatePost)

	log.Panic(app.Listen(":8080"))
}
