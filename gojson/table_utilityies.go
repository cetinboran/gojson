package gojson

import (
	"fmt"
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
	if len(names) != len(values) {
		v := fmt.Sprintf("%v != %v", len(names), len(values))
		fmt.Println(errorhandler.GetErrorTable(4, v))
		os.Exit(4)
	}

	CheckNames(names, t)
	CheckValues(values, t)
	// Eğer buraya geçerse yukarıdaki fonksiyonlar os.exit çalıştırmamıştır.

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
