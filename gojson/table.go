package gojson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/cetinboran/gojson/errorhandler"
)

// Initialize Table
func CreateTable(tableName string) Table {
	return Table{Name: tableName}
}

// Adds Property to the table
func (t *Table) AddProperty(name string, valueType string, mode string) {
	t.Properties = append(t.Properties, Property{Name: name, Type: valueType, Mode: mode})
}

func (t *Table) Save(data Data) {
	checkNames(data.Names, t)

	// Gets all the data form propeties like Mode and needed value type
	data.getDataFromProperties(t.Properties)
	data.checkMods(t)

	// Bütün data gelince data.getDataformproperties den bütünü kontrol et.
	checkValues(data.Values, data.Types, t)

	newData := getMapForJson(data.Names, data.Values, t)
	writeToJson(newData, t)

}

func (t *Table) Find(uniqueStr string, uniqueStrValue interface{}) []map[string]interface{} {
	jsonData := t.Get()

	if fmt.Sprint(reflect.TypeOf(uniqueStrValue)) == "int" {
		uniqueStrValue = float64(uniqueStrValue.(int))
	}

	var all []map[string]interface{}

	for i, v := range jsonData {
		for k, v2 := range v {
			if k == uniqueStr && v2 == uniqueStrValue {
				all = append(all, jsonData[i])
			}
		}
	}

	return all
}

// Updates Spesific Data
func (t *Table) Update(uniqueStr string, uniqueStrValue interface{}, data Data) {
	jsonData := t.Get()

	indexArr := getIndex(jsonData, uniqueStr, uniqueStrValue)

	// All bütün property lere sahip olucak çünkü data tipinde. O yüzden default value ise değiştirmeyeceğiz. Değilse değişilecek
	all := make(map[string]interface{})

	for i, v := range data.Names {
		for i2, v2 := range data.Values {
			if i == i2 {
				all[v] = v2
			}
		}
	}

	// Değişebilecek bütün değerlerin indexlerini buluyorum sonra hepsine değişimi uyguluyorum.
	for _, index := range indexArr {
		for k := range jsonData[index] {
			for k2, v2 := range all {
				// Eğer propert isimleri aynı ise bazı kontrollerden geçip değiştireceğiz.
				if k == k2 {
					switch fmt.Sprint(reflect.TypeOf(v2)) {
					case "int":
						value := findPkName(t)
						// Primary Key ise değiştirlmesin.
						// Eğer biri PK olan bir değeri değiştirmeye çalışıyorsa diye alttaki if var. PK ise değiştirme yapmaz.
						if v2 != 0 && value != "PK" {
							jsonData[index][k] = v2
						}
						break
					case "string":
						if v2 != "" {
							jsonData[index][k] = v2
						}
					case "bool":
						jsonData[index][k] = v2
						break
					case "[]string":
						jsonData[index][k] = v2
						break
					case "[]int":
						jsonData[index][k] = v2
						break
					}
				}
			}
		}

		saveUpdatedData(jsonData, t)
	}
}

// Deletes spesific data
func (t *Table) Delete(uniqueStr string, uniqueStrValue interface{}) {
	if fmt.Sprint(reflect.TypeOf(uniqueStrValue)) == "int" {
		uniqueStrValue = float64(uniqueStrValue.(int))
	}

	// Bunda sıkıntı yok
	if len(t.Find(uniqueStr, uniqueStrValue)) == 0 {
		fmt.Printf(errorhandler.GetErrorTable(5, fmt.Sprint(uniqueStrValue)))
		os.Exit(5)
	}

	jsonData := t.Get()

	var all []map[string]interface{}

	for i, v := range jsonData {
		for k, v2 := range v {
			if k == uniqueStr && v2 != uniqueStrValue {
				all = append(all, jsonData[i])
				break
			}
		}
	}

	saveUpdatedData(all, t)
}

// Returns existing Data
func (t *Table) Get() []map[string]interface{} {
	filePath := t.PathDatabase + t.Name + ".json"

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
