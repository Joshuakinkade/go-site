package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joshuakinkade/go-site/models"
	"github.com/joshuakinkade/go-site/services"
)

type APIHandler struct {
	posts services.PostService
}

func NewAPIHandler(posts services.PostService) APIHandler {
	return APIHandler{posts: posts}
}

// ListPosts returns a list of posts
func (h APIHandler) ListPosts(ctx *fiber.Ctx) error {
	posts, err := h.posts.ListPosts(0, 10)
	if err != nil {
		return err
	}

	ctx.JSON(posts)
	return nil
}

// GetPost returns a single post
func (h APIHandler) GetPost(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	post, err := h.posts.GetPostBySlug(slug)
	if err != nil {
		return err
	}
	type postReponse struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	body, err := post.Render()
	if err != nil {
		return err
	}
	pr := postReponse{
		Title: post.Title,
		Body:  body,
	}
	ctx.JSON(pr)
	return nil
}

// CreatePost adds a new post to the blog.
func (h APIHandler) CreatePost(ctx *fiber.Ctx) error {
	var np models.Post
	err := ctx.BodyParser(&np)
	if errors.Is(fiber.ErrUnprocessableEntity, err) {
		ctx.SendStatus(fiber.StatusUnprocessableEntity)
		return nil
	}

	np, err = h.posts.CreatePost(np)
	if err != nil {
		ctx.Write([]byte(fmt.Sprintf("could not create post: %s", err.Error())))
		ctx.SendStatus(fiber.StatusInternalServerError)
		return nil
	}

	ctx.JSON(np)
	ctx.SendStatus(fiber.StatusCreated)
	return nil
}

// UpdatePost updates an existing post.
//
// Fields:
// - title: string
// - body: string
// - published: bool
func (h APIHandler) UpdatePost(ctx *fiber.Ctx) error {
	var body map[string]interface{}
	err := ctx.BodyParser(&body)
	if errors.Is(fiber.ErrUnprocessableEntity, err) {
		ctx.SendStatus(fiber.StatusUnprocessableEntity)
		return nil
	}

	// Todo: validate the body here before using it.

	slug := ctx.Params("slug")

	err = h.posts.UpdatePost(slug, body)
	return err
}
