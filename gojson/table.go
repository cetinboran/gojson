package gojson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// Initialize Table
func CreateTable(tableName string) Table {
	return Table{TableName: tableName}
}

// Adds Property to the table
func (t *Table) AddProperty(name string, valueType string, mode string) {
	t.Properties = append(t.Properties, Property{Name: name, Type: valueType, Mode: mode})
}

// Addes data to the table
func (t *Table) Save(name string, values []interface{}) {
	name = strings.ReplaceAll(name, " ", "") // Bütün boşluklardan kurtuluyorum
	nameArr := strings.Split(name, ",")      // , ile split atıyorum.

	// Burası hata çıkarırsa programdan exit atıyor. yoksa devam ediyor.
	CheckArgs(nameArr, values, t)

	newData := GetMapForJson(nameArr, values, t)
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
