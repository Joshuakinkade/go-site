package validation_test

import (
	"regexp"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/joshuakinkade/go-site/lib/validation"
)

var _ = Describe("Validation", func() {
	Describe("StringValidator", func() {
		It("should not reject a valid string", func() {
			v := validation.String().MinLength(1).MaxLength(5).Regexp(regexp.MustCompile(".*"))
			Expect(v.Validate("hello")).To(BeNil())
		})
		It("should reject non-string value", func() {
			v := validation.String()
			Expect(v.Validate(123)).To(MatchError("value is not a string"))
		})
		It("should reject a short string", func() {
			v := validation.String().MinLength(5).MaxLength(10).Regexp(regexp.MustCompile(".*"))
			Expect(v.Validate("hi")).To(MatchError("value is too short"))
		})
		It("should reject a long string", func() {
			v := validation.String().MinLength(1).MaxLength(5).Regexp(regexp.MustCompile(".*"))
			Expect(v.Validate("hello, world!")).To(MatchError("value is too long"))
		})
		It("should reject a string that doesn't match the regexp", func() {
			v := validation.String().MinLength(1).MaxLength(10).Regexp(regexp.MustCompile("^[A-z]+$"))
			Expect(v.Validate("hello123")).To(MatchError("value does not match pattern"))
		})
		It("should ignore empty regexp", func() {
			v := validation.String().MaxLength(1).MaxLength(6)
			Expect(v.Validate("hello")).To(BeNil())
		})
		It("should ignore empty (0) minLength", func() {
			v := validation.String().MaxLength(6).Regexp(regexp.MustCompile(".*"))
			Expect(v.Validate("")).To(BeNil())
		})
		It("should ignore empty (0) maxLength", func() {
			v := validation.String().MinLength(1).Regexp(regexp.MustCompile(".*"))
			Expect(v.Validate("hello")).To(BeNil())
		})
	})
	Describe("MapValidator", func() {
		It("should validate its children", func() {
			v := validation.Map()
			v.Add("name", validation.String().MinLength(1))
			Expect(v.Validate(map[string]interface{}{"name": "hello"})).To(BeNil())
		})
		It("should return it's childrens' errors", func() {
			v := validation.Map()
			v.Add("name", validation.String().MinLength(1).MaxLength(2))
			Expect(v.Validate(map[string]interface{}{"name": "hello"})).To(MatchError("field name: value is too long"))
		})
	})
})
