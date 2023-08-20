package gojson

import (
	"strings"
)

// Initialize Table
func CreateTable(tableName string) Table {
	return Table{TableName: tableName}
}

// Adds Property to the table
func (t *Table) AddProperty(name string, valueType string) {
	t.Properties = append(t.Properties, Property{Name: name, Type: valueType})
}

// Addes data to the table
func (t *Table) Save(name string, values []interface{}) {
	name = strings.ReplaceAll(name, " ", "") // Bütün boşluklardan kurtuluyorum
	nameArr := strings.Split(name, ",")      // , ile split atıyorum.

	// Burası hata çıkarırsa programdan exit atıyor. yoksa devam ediyor.
	CheckArgs(nameArr, values, t)

	WriteToJson(GetMapForJson(nameArr, values), t)
}
