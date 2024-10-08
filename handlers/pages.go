package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshuakinkade/go-site/models"
	"github.com/joshuakinkade/go-site/services"
)

type PagesHandler struct {
	posts services.PostService
}

func NewPagesHandler(posts services.PostService) PagesHandler {
	return PagesHandler{posts: posts}
}

// ShowHome renders the home page
func (h PagesHandler) ShowHome(ctx *fiber.Ctx) error {
	posts, err := h.posts.ListPosts(0, 3)
	if err != nil {
		return err
	}

	return ctx.Render("index", fiber.Map{
		"Message":     "Hello, World!",
		"RecentPosts": posts,
	}, "layouts/base")
}

// ShowPost looks for the post with the given slug and renders it
func (h PagesHandler) ShowPost(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	// post, err := h.posts.GetPostBySlug(slug)
	post := models.Post{
		Title: "Sample Post",
		Slug:  slug,
		Body:  "This is a sample post.",
	}

	return ctx.Render("post", fiber.Map{
		"Post": post,
	})
}
