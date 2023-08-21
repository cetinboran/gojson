package gojson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Initialize Table
func CreateTable(tableName string) Table {
	return Table{TableName: tableName}
}

// Adds Property to the table
func (t *Table) AddProperty(name string, valueType string, mode string) {
	t.Properties = append(t.Properties, Property{Name: name, Type: valueType, Mode: mode})
}

func (t *Table) Save(data Data) {
	// Gets all the data form propeties like Mode and needed value type
	data.GetDataFromProperties(t.Properties)
	data.CheckMods(t)

	CheckNames(data.Names, t)
	CheckValues(data.Values, data.Types, t)

	newData := GetMapForJson(data.Names, data.Values, t)
	WriteToJson(newData, t)

}

func (t *Table) Get() []map[string]interface{} {
	filePath := t.PathDatabase + t.TableName + ".json"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	var existingData []map[string]interface{}
	err = json.Unmarshal(byteValue, &existingData)
	if err != nil {
		fmt.Println(err)
	}

	return existingData
}
