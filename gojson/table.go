package gojson

// Initialize Table
func CreateTable(tableName string) Table {
	return Table{TableName: tableName}
}

// Adds Property to the table
func (t *Table) AddProperty(name string, valueType string) {
	t.Properties = append(t.Properties, Property{Name: name, Type: valueType})
}
