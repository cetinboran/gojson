package gojson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/cetinboran/gojson/errorhandler"
)

// Here we compare property names with incoming names if they are not the same we send an error. At the same time, we send an error if the same property name is written more than once.
func CheckNames(names []string, t *Table) {
	check := false

	for _, v := range names {
		// Buradaki count karakterleri saydığı için a yazınca propertiy olarak sıkıntı çıkıyor.
		if countWord(names, v) != 1 {
			fmt.Println(errorhandler.GetErrorTable(2, v))
			os.Exit(2)
		}
		for _, p := range t.Properties {
			if v == p.Name {
				check = true
				break
			} else {
				check = false

			}
		}
		if !check {
			fmt.Println(errorhandler.GetErrorTable(1, v))
			os.Exit(1)
		}
	}

	// Buraya kadar geldiyse exit e gelmemiştir o zaman sorun yok.
}

func CheckValues(values []interface{}, types []string, t *Table) {
	check := false

	var typeMustBe string
	// Bütün propery ler girilmiş gibi düşünülüp kontrol ediliyor.
	for i, v := range values {
		typeStr := fmt.Sprint(reflect.TypeOf(v))
		for j, t := range types {
			// Bütün value kendi type'ına baksın diye i==j kontrolü var
			if i == j {
				if t == typeStr {
					check = true
					break
				} else {
					check = false
					typeMustBe = t
				}
			}
		}
		if !check {
			// gelentype->olmasıgerekentype şeklinde bir hata mesajı veriyoruz.
			fmt.Println(errorhandler.GetErrorTable(3, fmt.Sprintf("The type of %v value ", v)+"must be "+typeMustBe))
			os.Exit(3)
		}
	}
}

// Writes to json file.
func WriteToJson(data map[string]interface{}, t *Table) {
	filePath := t.PathDatabase + t.TableName + ".json"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		os.Exit(9)
	}
	defer file.Close()

	// Mevcut veriyi oku
	byteValue, _ := io.ReadAll(file)
	var existingData []map[string]interface{}
	err = json.Unmarshal(byteValue, &existingData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(9)
	}

	// Jsonu okduğumuz yerden aldığımız dataya eklieyeceğim datayı ekliyoruz
	existingData = append(existingData, data)

	// JSON dosyasını yeniden yaz Buradaki marshallIndet insanların okuyacağı şekilde yazar.
	newJSONData, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		os.Exit(9)
	}

	err = os.WriteFile(filePath, newJSONData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		os.Exit(9)
	}
}

// Saves Updated Json.
func SaveUpdatedData(updatedData []map[string]interface{}, t *Table) {
	filePath := t.PathDatabase + t.TableName + ".json"

	newJSONData, err := json.Marshal(updatedData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		os.Exit(9)
	}

	err = os.WriteFile(filePath, newJSONData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		os.Exit(9)
	}
}

// Değişebilecek bütün dataların indexini buluyorum ve Update Fonksiyonuna yolluyorum.
func GetIndex(jsonData []map[string]interface{}, uniqueStr string, uniqueStrValue interface{}) []int {
	if fmt.Sprint(reflect.TypeOf(uniqueStrValue)) == "int" {
		uniqueStrValue = float64(uniqueStrValue.(int))
	}

	var index []int

	for i, v := range jsonData {
		for k, v2 := range v {
			if k == uniqueStr && v2 == uniqueStrValue {
				index = append(index, i)
			}
		}
	}

	return index
}

func GetMapForJson(names []string, values []interface{}, t *Table) map[string]interface{} {
	data := make(map[string]interface{})

	for i, v := range names {
		for i2, v2 := range values {
			if i == i2 {
				data[v] = v2
			}
		}
	}

	return data
}

func FindPkName(t *Table) string {
	for _, p := range t.Properties {
		if p.Mode == "PK" {
			return p.Name
		}
	}

	return ""
}

func countWord(names []string, name string) int {
	var count int

	// Kelimeyi sayıyoruz
	for _, v := range names {
		if v == name {
			count++
		}
	}

	return count
}
