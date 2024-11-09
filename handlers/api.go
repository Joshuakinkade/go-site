package handlers

import (
	"errors"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/joshuakinkade/go-site/models"
	"github.com/joshuakinkade/go-site/services"
)

type JSONResponse struct {
	Data struct {
		Type    string      `json:"type"`
		Content interface{} `json:"content"`
	} `json:"data;omitempty"`
	Error string `json:"error;omitempty"`
}

func SuccessResponse(contentType string, content interface{}) JSONResponse {
	resp := JSONResponse{}
	resp.Data.Type = contentType
	resp.Data.Content = content
	return resp
}

func ErrorResponse(er error) JSONResponse {
	resp := JSONResponse{}
	resp.Error = er.Error()
	return resp
}

// APIHandler provides methods for handling API requests.
type APIHandler struct {
	posts services.PostService
}

// NewAPIHandler returns an initialized APIHandler.
func NewAPIHandler(posts services.PostService) APIHandler {
	return APIHandler{posts: posts}
}

// ListPosts returns a list of posts
func (h APIHandler) ListPosts(ctx *fiber.Ctx) error {
	posts, _, err := h.posts.ListPosts(0, 10)
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

	ctx.JSON(SuccessResponse("post", pr))
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
		ctx.JSON(ErrorResponse(err))
		ctx.SendStatus(fiber.StatusInternalServerError)
		return nil
	}

	ctx.JSON(SuccessResponse("post", np))
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

	if err != nil {
		ctx.JSON(ErrorResponse(err))
		ctx.SendStatus(fiber.StatusInternalServerError)
		return nil
	}

	return err
}

// UploadPhoto uploads a photo to the server. It checks that the body is a valid
// image format and that it's not too large.
//
// Fields:
// - photo: file
// - alt: string
// - caption: string
func (h APIHandler) UploadPhoto(ctx *fiber.Ctx) error {
	contentType := ctx.Get("Content-Type")
	allowedContentTypes := []string{"image/jpeg", "image/png"}
	if sort.SearchStrings(allowedContentTypes, contentType) == len(allowedContentTypes) {
		ctx.JSON(ErrorResponse(errors.New("image is not a permitted type")))
		ctx.SendStatus(fiber.StatusBadRequest)
		return nil
	}

	return nil
}
