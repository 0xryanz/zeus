package schema

import (
	"go/ast"
	"reflect"
	"zorm/dialect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destVlues := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destVlues.FieldByName(field.Name).Interface())
	}

	return fieldValues
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelTyep := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelTyep.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelTyep.NumField(); i++ {
		p := modelTyep.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}

			if v, ok := p.Tag.Lookup("zorm"); ok {
				field.Tag = v
			}

			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
