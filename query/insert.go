package query

import (
	"fmt"
	"github.com/raflyafrzl/rasql/schema"
	"strings"
)

type InsertQueryBuilder struct {
	table  string
	column string
	length int
	result string
}

func Insert(schema *schema.Schema) *InsertQueryBuilder {

	insertStatement := strings.Join(schema.GetFields(), ", ")

	iQb := &InsertQueryBuilder{column: insertStatement, table: schema.TableName, length: len(schema.Fields)}

	return iQb

}

func (q *InsertQueryBuilder) Construct() string {
	values := strings.Repeat("?,", q.length)
	q.result = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", q.table, q.column, values[:len(values)-1])

	return q.result
}
