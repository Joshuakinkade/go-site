package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
	posts, _, err := h.posts.ListPosts(0, 3)
	if err != nil {
		return err
	}

	return ctx.Render("index", fiber.Map{
		"Message":     "Hello, World!",
		"RecentPosts": posts,
	}, "layouts/base")
}

func (h PagesHandler) ShowPostList(ctx *fiber.Ctx) error {
	pageParam := ctx.Query("page")
	var page int64 = 1
	if len(pageParam) > 0 {
		var err error
		page, err = strconv.ParseInt(pageParam, 10, 64)
		if err != nil {
			ctx.WriteString("Invalid page number")
			ctx.Status(fiber.StatusBadRequest)
			return nil
		}
	}

	pageSize := 2
	start := (int(page) - 1) * pageSize
	end := start + pageSize

	posts, totalPosts, err := h.posts.ListPosts(start, end)
	if err != nil {
		return err
	}

	// previous link
	prevLink := ""
	if page > 1 {
		prevLink = fmt.Sprintf("/posts?page=%v", page-1)
	}
	// next link, if there are more posts
	nextLink := ""
	if end < totalPosts {
		nextLink = fmt.Sprintf("/posts?page=%v", page+1)
	}

	ctx.Render("posts", fiber.Map{
		"Posts": posts,
		"Next":  nextLink,
		"Prev":  prevLink,
	}, "layouts/base")

	return nil
}

// ShowPost looks for the post with the given slug and renders it
func (h PagesHandler) ShowPost(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	post, err := h.posts.GetPostBySlug(slug)
	if err != nil {
		return err
	}

	return ctx.Render("post", fiber.Map{
		"Post": post,
	}, "layouts/base")
}

func (h PagesHandler) ShowError(ctx *fiber.Ctx, err error) error {
	var e *fiber.Error
	if errors.As(err, &e) {
		if e.Code == fiber.StatusNotFound {
			return ctx.Render("not-found", nil, "layouts/base")
		}
		return ctx.Render("internal-error", nil, "layouts/base")
	}
	return nil
}
