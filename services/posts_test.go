package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/joshuakinkade/go-site/models"
	"github.com/joshuakinkade/go-site/services"
)

var _ = Describe("Posts", func() {
	Describe("ListPosts", func() {
		It("should return a list of posts", func() {
			posts := services.NewPostService()
			list, err := posts.ListPosts(0, 10)
			Expect(err).To(BeNil())
			Expect(list).ToNot(BeEmpty())
			Expect(len(list)).To(BeNumerically("<", 11))
		})
	})

	Describe("GetPostBySlug", func() {
		It("Should return the post", func() {
			posts := services.NewPostService()
			post, err := posts.GetPostBySlug("hello-world")
			Expect(err).To(BeNil())
			Expect(post.Title).To(Equal("Hello, World!"))
		})

		It("Should return an empty post", func() {
			posts := services.NewPostService()
			post, err := posts.GetPostBySlug("not-found")
			Expect(err).To(BeNil())
			Expect(post.Title).To(BeEmpty())
		})
	})

	Describe("CreatePost", func() {
		It("Should create a post", func() {
			posts := services.NewPostService()
			post, err := posts.CreatePost(models.Post{
				Title:   "Hello, World!",
				Slug:    "hello-world",
				Content: "This is a test post. It's not very interesting, but it's a start.",
			})
			Expect(err).To(BeNil())
			Expect(post.Title).To(Equal("Hello, World!"))
		})

		It("Should return generate a slug if ones not given", func() {
			posts := services.NewPostService()
			post, err := posts.CreatePost(models.Post{
				Title:   "Hello, World!",
				Content: "This is a test post. It's not very interesting, but it's a start.",
			})
			Expect(err).To(BeNil())
			Expect(post.Slug).To(Equal("hello-world"))
		})
	})
})
