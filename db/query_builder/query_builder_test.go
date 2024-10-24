package querybuilder_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	querybuilder "github.com/joshuakinkade/go-site/db/query_builder"
)

var _ = Describe("QueryBuilder", func() {
	It("should build a query with the given parameters", func() {
		params := map[string]interface{}{
			"name": "James",
			"age":  25,
		}

		updates, args, err := querybuilder.BuildUpdateClause(params, []string{"age", "name"})
		Expect(err).To(BeNil())
		Expect(updates).To(Equal("name = $1, age = $2"))
		Expect(len(updates)).To(Equal(len(args)))
	})

	It("should reject fields that aren't allowed", func() {
		params := map[string]interface{}{
			"name": "James",
			"age":  25,
		}

		_, _, err := querybuilder.BuildUpdateClause(params, []string{"name"})
		Expect(err).To(Not(BeNil()))
		Expect(err.Error()).To(Equal("Invalid field: age"))
	})
})
