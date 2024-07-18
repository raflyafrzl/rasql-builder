package schema

import (
	"fmt"
	"reflect"
)

type schemaResult struct {
	TableName string
	Fields    []string
}
type Schema schemaResult

func DefineSchema(schema interface{}) *Schema {

	v := reflect.TypeOf(schema)
	if v.Kind() != reflect.Ptr {
		panic(fmt.Errorf("should be a pointer"))
	}
	v = v.Elem()
	result := Schema{
		TableName: `"` + v.Name() + `"`,
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		result.Fields = append(result.Fields, `"`+f.Tag.Get("db_field")+`"`)
	}
	return &result
}

func (s *Schema) GetTableName() string {
	return s.TableName
}
func (s *Schema) GetFields() []string {
	return s.Fields
}
