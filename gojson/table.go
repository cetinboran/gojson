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
	return Table{TableName: tableName}
}

// Adds Property to the table
func (t *Table) AddProperty(name string, valueType string, mode string) {
	t.Properties = append(t.Properties, Property{Name: name, Type: valueType, Mode: mode})
}

func (t *Table) Save(data Data) {
	CheckNames(data.Names, t)

	// Gets all the data form propeties like Mode and needed value type
	data.GetDataFromProperties(t.Properties)
	data.CheckMods(t)

	// Bütün data gelince data.getDataformproperties den bütünü kontrol et.
	CheckValues(data.Values, data.Types, t)

	newData := GetMapForJson(data.Names, data.Values, t)
	WriteToJson(newData, t)

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

func (t *Table) Update(uniqueStr string, uniqueStrValue interface{}, data Data) {
	jsonData := t.Get()

	indexArr := GetIndex(jsonData, uniqueStr, uniqueStrValue)

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
						value := FindPkName(t)
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
					}
				}
			}
		}

		SaveUpdatedData(jsonData, t)
	}
}

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

	// Bu mantıklı ancak içeride bir sürü sorun çıkarabiliyor
	// birbine bağlı bir sürü database varken sıkıntı çıkıyor çünkü diyelim user sildik id si değişti
	// config dosyasındaki userıd aynı kaldı ancak userın ıd si değiştiği için bağlantı kayıyor.

	// // Burada o table'ın PK nameini buluyorum ve idlerini en baştan atıyorum ki id ye göre işlem yapanlarda sıkıntı çıkmasın.
	// // Çünkü Pk name unique dir.
	// count := 1
	// for i, v := range all {
	// 	for k := range v {
	// 		if FindPkName(t) == k {
	// 			all[i][k] = count
	// 			break
	// 		}
	// 	}
	// 	count++
	// }

	SaveUpdatedData(all, t)
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
