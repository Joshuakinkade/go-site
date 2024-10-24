package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/joshuakinkade/go-site/lib/validation"
)

type TestStruct struct {
}

var _ = Describe("Validation", func() {
	It("should not return errors for a valid struct", func() {

		err := validation.ValidateStruct(ex)
		Expect(err).To(BeNil())
	})
})
