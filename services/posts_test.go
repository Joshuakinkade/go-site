package services_test

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/joshuakinkade/go-site/models"
	"github.com/joshuakinkade/go-site/services"
)

type mockRepository struct{}

func (m mockRepository) ListPosts(offset, limit int) ([]models.Post, error) {
	if offset == 0 {
		return []models.Post{{
			Title:       "Hello, World!",
			Slug:        "hello-world",
			Body:        "This is a test post. It's not very interesting, but it's a start.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			PublishedAt: nil,
		}}, nil
	} else if offset == 10 {
		return []models.Post{}, nil
	} else { // simulate a repository error
		return nil, errors.New("could not list posts")
	}
}

func (m mockRepository) GetPostBySlug(slug string) (models.Post, error) {
	return models.Post{}, nil
}

func (m mockRepository) CreatePost(post models.Post) (models.Post, error) {
	return post, nil
}

func (m mockRepository) UpdatePost(post models.Post) (models.Post, error) {
	return post, nil
}

var _ = Describe("Posts", func() {
	var posts services.PostService
	BeforeEach(func() {
		posts = services.NewPostService(mockRepository{})
	})

	Describe("ListPosts", func() {

		It("should return a list of posts", func() {
			list, err := posts.ListPosts(0, 10)
			Expect(err).To(BeNil())
			Expect(list).ToNot(BeEmpty())
			Expect(len(list)).To(BeNumerically("<", 11))
		})

		It("should return an empty slice with no error if no posts are available", func() {
			list, err := posts.ListPosts(10, 10)
			Expect(err).To(BeNil())
			Expect(list).To(BeEmpty())
		})

		It("should return an error if the repository fails", func() {
			list, err := posts.ListPosts(20, 10)
			Expect(err).ToNot(BeNil())
			Expect(list).To(BeEmpty())
		})
	})

	Describe("GetPostBySlug", func() {
		It("Should return the post", func() {
			post, err := posts.GetPostBySlug("hello-world")
			Expect(err).To(BeNil())
			Expect(post.Title).To(Equal("Hello, World!"))
		})

		It("Should return an empty post", func() {
			post, err := posts.GetPostBySlug("not-found")
			Expect(err).To(BeNil())
			Expect(post.Title).To(BeEmpty())
		})
	})

	Describe("CreatePost", func() {
		It("Should create a post", func() {
			post, err := posts.CreatePost(models.Post{
				Title: "Hello, World!",
				Slug:  "hello-world",
				Body:  "This is a test post. It's not very interesting, but it's a start.",
			})
			Expect(err).To(BeNil())
			Expect(post.Title).To(Equal("Hello, World!"))
		})

		It("Should set defaults for missing data", func() {
			post, err := posts.CreatePost(models.Post{
				Title: "Hello, World!",
				Body:  "This is a test post. It's not very interesting, but it's a start.",
			})
			Expect(err).To(BeNil())
			Expect(post.Slug).To(Equal("hello-world"))
			Expect(post.CreatedAt).To(Equal(time.Now()))
		})
	})
})
