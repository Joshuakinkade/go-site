package handlers

import (
	"errors"
	"fmt"
	"regexp"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/joshuakinkade/go-site/models"
	"github.com/joshuakinkade/go-site/services"
)

type JSONResponse struct {
	Data struct {
		Type    string      `json:"type,omitempty"`
		Content interface{} `json:"content,omitempty"`
	} `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
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

type StringChecker struct {
	fieldName string
	required  bool
	maxLength int
	minLength int
	pattern   *regexp.Regexp
}

// Check validates that a map has a field and that its value is valid.
// Can this check multiple data types? map, string, array
// If string, assume its the value to validate
// If array, check each value
// If map, look for value with fieldName
func (c StringChecker) Check(data map[string]interface{}) (string, error) {
	if value, ok := data[c.fieldName]; ok {
		switch value := value.(type) {
		case string:
			return value, nil
		default:
			return "", fmt.Errorf("%v must be a string", c.fieldName)
		}
	} else {
		if c.required {
			return "", fmt.Errorf("%v is required", c.fieldName)
		}
		return "", nil
	}
}

// CreatePost adds a new post to the blog.
func (h APIHandler) CreatePost(ctx *fiber.Ctx) error {
	var np map[string]interface{}
	err := ctx.BodyParser(&np)
	if errors.Is(fiber.ErrUnprocessableEntity, err) {
		ctx.SendStatus(fiber.StatusUnprocessableEntity)
		return nil
	}

	var newPost models.Post

	// I prefer the builder syntax for defining rules
	titleChecker := StringChecker{
		fieldName: "title",
		required:  true,
	}
	title, err := titleChecker.Check(np) // I like that this returns a typed value
	// This is a lot of error handling code for one field, and it's going to be
	// repeated for each field.
	if err != nil {
		ctx.JSON(ErrorResponse(err))
		ctx.SendStatus(fiber.StatusBadRequest)
		return nil
	}
	newPost.Title = title

	bodyChecker := StringChecker{
		fieldName: "body",
		required:  true,
	}
	body, err := bodyChecker.Check(np)
	if err != nil {
		ctx.JSON(ErrorResponse(err))
		ctx.SendStatus(fiber.StatusBadRequest)
	}
	newPost.Body = body

	newPost, err = h.posts.CreatePost(newPost)
	if err != nil {
		ctx.JSON(ErrorResponse(err))
		ctx.SendStatus(fiber.StatusInternalServerError)
		return nil
	}

	ctx.JSON(SuccessResponse("post", newPost))
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
