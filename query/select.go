package query

import (
	"github.com/raflyafrzl/rasql/schema"
	"strings"
)

type SelectQueryBuilder struct {
	column    string
	where     string
	tableName string
	result    string
	stmt      string
}

func Select(column ...string) *SelectQueryBuilder {
	resultEval := evalValue(column...)
	q := SelectQueryBuilder{column: strings.Join(resultEval, ",")}
	q.stmt = "SELECT "
	return &q
}

func evalValue(value ...string) []string {

	var result []string
	for _, v := range value {
		result = append(result, `"`+v+`"`)
	}
	return result

}

func (q *SelectQueryBuilder) From(schema *schema.Schema) *SelectQueryBuilder {
	q.tableName = schema.GetTableName()
	return q
}

func (q *SelectQueryBuilder) Where(cond ...string) *SelectQueryBuilder {

	for _, v := range cond {
		q.where += strings.Replace(v, "<<table>>", q.tableName+".", -1)
	}
	return q
}

func (q *SelectQueryBuilder) Result() string {

	q.result = q.stmt + q.column + " " + q.tableName
	if q.where != "" {
		q.result += " WHERE " + q.where
	}
	return q.result
}
