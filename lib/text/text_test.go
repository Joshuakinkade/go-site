package text_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/joshuakinkade/go-site/lib/text"
)

var _ = Describe("Text", func() {
	Describe("Slugify", func() {
		It("should convert a string into a URL-friendly slug", func() {
			title := "Hello, World!--@1"
			slug := text.Slugify(title)
			Expect(slug).To(Equal("hello-world-1"))
		})
	})
})
