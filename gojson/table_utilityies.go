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

func CheckValues(values []interface{}, t *Table) {
	check := false

	for i, v := range values {
		typeStr := fmt.Sprint(reflect.TypeOf(v))
		for j, p := range t.Properties {
			// Burada 0. index value ile 0. index table type ı kayaslanıyor. Yani i = j olmalı diğerlerine bakmanın anlamı yok.
			if i == j {
				if p.Type == typeStr {
					check = true
					break
				} else {
					check = false
				}
			}
		}
		if !check {
			// gelentype->olmasıgerekentype şeklinde bir hata mesajı veriyoruz.
			fmt.Println(errorhandler.GetErrorTable(3, typeStr+"->"+t.Properties[i].Type+" in the "+fmt.Sprint(i+1)+" th row."))
			os.Exit(3)
		}
	}
}

func CheckArgs(names []string, values []interface{}, t *Table) {
	// Artık kaç tane properties var ise o kadar input girmesi lazım values ve names olarak
	// Eksik ise hata atıyoruz.

	if len(names) != len(t.Properties) {
		v := fmt.Sprintf("%v != %v", len(names), len(t.Properties))
		fmt.Println(errorhandler.GetErrorTable(4, v))
		os.Exit(4)
	}

	if len(values) != len(t.Properties) {
		v := fmt.Sprintf("%v != %v", len(values), len(t.Properties))
		fmt.Println(errorhandler.GetErrorTable(4, v))
		os.Exit(4)
	}

	CheckNames(names, t)
	CheckValues(values, t)
	// Eğer buraya geçerse yukarıdaki fonksiyonlar os.exit çalıştırmamıştır.

}

// Writes to json file.
func WriteToJson(data map[string]interface{}, t *Table) {
	filePath := t.PathDatabase + t.TableName + ".json"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// Mevcut veriyi oku
	byteValue, _ := io.ReadAll(file)
	var existingData []map[string]interface{}
	err = json.Unmarshal(byteValue, &existingData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Jsonu okduğumuz yerden aldığımız dataya eklieyeceğim datayı ekliyoruz
	existingData = append(existingData, data)

	// JSON dosyasını yeniden yaz
	newJSONData, err := json.Marshal(existingData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	err = os.WriteFile(filePath, newJSONData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}
}

// Finds Primary Key
func FindPrimaryKey(t *Table) string {
	var name string
	for _, p := range t.Properties {
		if p.Mode == "PK" {
			name = p.Name
			break
		}
	}

	return name
}

func GetMapForJson(names []string, values []interface{}, t *Table) map[string]interface{} {
	data := make(map[string]interface{})

	for i, v := range names {
		for i2, v2 := range values {
			if i == i2 {
				data[v] = v2

				// Eğer v primary'key ise value'yu jsona göre ayarla.
				if v == FindPrimaryKey(t) {
					data[v] = len(t.Get()) + 1
				}
			}
		}
	}

	return data
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
