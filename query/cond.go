package query

import (
	"fmt"
	"strings"
)

type Cond map[string]any

type RaWhereCondition string

func Eq(op string, value Cond) string {

	var where []string
	for k, v := range value {

		where = append(where, fmt.Sprintf(`<<table>>"%s" = %v`, k, v))
	}
	return strings.Join(where, " AND ")

}
