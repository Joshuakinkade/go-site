package querybuilder

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

// BuildUpdateClause creates the update clause of a sql statement from a map of fields and values.
// It takes a slice of strings with the names of fields that are allowed in the update and returns
// an error if there are fields in the input map that are not in the allowed fields slice. On
// success, it returns the update clause as a string with parameter placeholders and a slice of with
// the values. The allowed fields **must** be given in alphabetical order.
func BuildUpdateClause(updates map[string]interface{}, allowedFields []string) (string, []interface{}, error) {
	var wheres []string
	var args []interface{}
	var i int64 = 1
	for field, value := range updates {
		if sort.SearchStrings(allowedFields, field) < len(allowedFields) {
			wheres = append(wheres, field+" = $"+strconv.FormatInt(i, 10))
			args = append(args, value)
			i++
		} else {
			return "", nil, errors.New("Invalid field: " + field)
		}
	}
	return strings.Join(wheres, ", "), args, nil
}
