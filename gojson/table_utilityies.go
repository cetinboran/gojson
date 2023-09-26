package gojson

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"

	"github.com/cetinboran/gojson/errorhandler"
)

// Here we compare property names with incoming names if they are not the same we send an error. At the same time, we send an error if the same property name is written more than once.
func checkNames(names []string, t *Table) {
	check := false

	for _, v := range names {
		// Buradaki count karakterleri saydığı için a yazınca propertiy olarak sıkıntı çıkıyor.
		if countWord(names, v) != 1 {
			log.Fatal(errorhandler.GetErrorTable(2, v))
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
			log.Fatal(errorhandler.GetErrorTable(1, v))
		}
	}

	// Buraya kadar geldiyse exit e gelmemiştir o zaman sorun yok.
}

func checkValues(values []interface{}, types []string, t *Table) {
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
			log.Fatal(errorhandler.GetErrorTable(3, fmt.Sprintf("The type of %v value ", v)+"must be "+typeMustBe))
		}
	}
}

// Writes to json file.
func writeToJson(data map[string]interface{}, t *Table) {
	filePath := t.PathDatabase + t.Name + ".json"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error opening JSON file:", err)
	}
	defer file.Close()

	// Mevcut veriyi oku
	byteValue, _ := io.ReadAll(file)
	var existingData []map[string]interface{}
	err = json.Unmarshal(byteValue, &existingData)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	// Jsonu okduğumuz yerden aldığımız dataya eklieyeceğim datayı ekliyoruz
	existingData = append(existingData, data)

	// JSON dosyasını yeniden yaz Buradaki marshallIndet insanların okuyacağı şekilde yazar.
	newJSONData, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		log.Fatal("Error encoding JSON:", err)
	}

	err = os.WriteFile(filePath, newJSONData, 0644)
	if err != nil {
		log.Fatal("Error writing JSON file:", err)
	}
}

// Saves Updated Json.
func saveUpdatedData(updatedData []map[string]interface{}, t *Table) {
	filePath := t.PathDatabase + t.Name + ".json"

	newJSONData, err := json.MarshalIndent(updatedData, "", "  ")
	if err != nil {
		log.Fatal("Error encoding JSON:", err)
	}

	err = os.WriteFile(filePath, newJSONData, 0644)
	if err != nil {
		log.Fatal("Error writing JSON file:", err)
	}
}

// Değişebilecek bütün dataların indexini buluyorum ve Update Fonksiyonuna yolluyorum.
func getIndex(jsonData []map[string]interface{}, uniqueStr string, uniqueStrValue interface{}) []int {
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

func getMapForJson(names []string, values []interface{}, t *Table) map[string]interface{} {
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

func findPkName(t *Table) string {
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
