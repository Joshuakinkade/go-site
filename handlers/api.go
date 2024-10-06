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

// CreatePost adds a new post to the blog.
func (h APIHandler) CreatePost(ctx *fiber.Ctx) error {
	var np models.Post
	err := ctx.BodyParser(&np)
	if err != nil {
		return errors.New("could not read body")
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
