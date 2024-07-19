package query

import (
	"fmt"
	"github.com/raflyafrzl/rasql/schema"
	"strings"
)

type SelectQueryBuilder struct {
	rawColumn []string
	columns   string
	where     string
	tableName string
	result    string
	stmt      string
}

func Select(column ...string) *SelectQueryBuilder {
	q := SelectQueryBuilder{rawColumn: column}
	return &q
}

func (s *SelectQueryBuilder) evalValue(value ...string) []string {

	var result []string
	for _, v := range value {
		c := fmt.Sprintf(`%s."%s"`, s.tableName, v)
		result = append(result, c)
	}
	return result

}

func (q *SelectQueryBuilder) From(schema *schema.Schema) *SelectQueryBuilder {
	q.tableName = schema.GetTableName()

	q.columns = strings.Join(q.evalValue(q.rawColumn...), ", ")

	return q
}

func (q *SelectQueryBuilder) Where(cond ...string) *SelectQueryBuilder {
	for _, v := range cond {
		q.where += strings.Replace(v, "<<table>>", q.tableName+".", -1)
	}
	return q
}

func (q *SelectQueryBuilder) Construct() string {

	q.result = fmt.Sprintf("SELECT %s FROM %s", q.columns, q.tableName)
	if q.where != "" {
		q.result = q.result + fmt.Sprintf(" WHERE %s", q.where)
	}
	return q.result
}
