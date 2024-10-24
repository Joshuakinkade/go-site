package querybuilder

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

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
