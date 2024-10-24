package querybuilder_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestQueryBuilder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "QueryBuilder Suite")
}
