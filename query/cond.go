package query

import (
	"fmt"
	"strings"
)

type Cond []string

type EqsClause struct {
	Op    Operator
	Field Cond
}

type Operator string

func Eqs(value EqsClause) string {

	var where []string
	for _, v := range value.Field {
		where = append(where, fmt.Sprintf(`<<table>>"%s" = ?`, v))
	}

	return strings.Join(where, string(value.Op))
}
