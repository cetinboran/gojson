package gojson

import (
	"fmt"
	"os"
	"reflect"
	"strings"

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

func SaveToTheTable(names []string, values []interface{}) {
	theJson := createTheJson(names, values)

	fmt.Print(theJson)
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

func createTheJson(names []string, values []interface{}) string {
	// İnternetten dız

	var jsonData strings.Builder
	jsonData.WriteString("{")
	for i, key := range names {
		if i != 0 {
			jsonData.WriteString(",")
		}

		jsonData.WriteString(`"`)
		jsonData.WriteString(key)
		jsonData.WriteString(`":`)

		// Type'a göre yazılış değişiyor.
		value := values[i]
		switch v := value.(type) {
		case int:
			jsonData.WriteString(fmt.Sprintf("%d", v))
		case string:
			jsonData.WriteString(`"`)
			jsonData.WriteString(v)
			jsonData.WriteString(`"`)
		case bool:
			if v {
				jsonData.WriteString(fmt.Sprintf("%d", 1))
			} else {
				jsonData.WriteString(fmt.Sprintf("%d", 0))
			}
		}
	}
	jsonData.WriteString("}")

	return jsonData.String()
}
