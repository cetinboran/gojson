package gojson

import (
	"fmt"
	"os"

	"github.com/cetinboran/gojson/errorhandler"
)

func DataInit(names []string, values []interface{}, t *Table) Data {
	if len(names) != len(values) {
		fmt.Println(errorhandler.GetErrorMods(2, "Property Names & Values"))
		os.Exit(2)
	}

	// Şuanlık çalışıyor İlere Sıkıntı çıkarsa bakarsın.
	// Eksik girilen name'lerin yerini defaut value'lar alıyor.
	if len(names) != len(t.Properties) {
		myMap := make(map[string]interface{})

		for i, v := range names {
			for j, v2 := range values {
				// Herkes kendi value'sını alsın diye i==j mi diye bakıyoruz.
				if i == j {
					myMap[v] = v2
				}
			}
		}

		check := false
		for _, p := range t.Properties {
			for _, v := range names {
				if p.Name == v {
					check = false
					break
				} else {
					check = true
				}
			}
			// Eğer buraya geldiği halde check true ise properties içindeki dönen name datanın içindeki name de yok. Bunu data ya ekleyelim.
			if check {
				switch p.Type {
				case "int":
					myMap[p.Name] = 0
					break
				case "string":
					myMap[p.Name] = ""
					break
				case "bool":
					myMap[p.Name] = false
					break
				}
			}
		}

		// make in ordaki 0 hiç bir eleman ile başlatılmasın anlamına geliyor
		// Normalde make ile oluşturursak boş eleman açıyor birkaç tane.
		newNames := make([]string, 0, len(myMap))
		newValues := make([]interface{}, 0, len(myMap))
		for k, v := range myMap {
			newNames = append(newNames, k)
			newValues = append(newValues, v)
		}

		return Data{Names: newNames, Values: newValues, Types: make([]string, len(t.Properties)), Mods: make([]string, len(t.Properties))}

	}

	// Girilen proporties sayısı kadar boyutta types ve mods oluşturdum.
	return Data{Names: names, Values: values, Types: make([]string, len(t.Properties)), Mods: make([]string, len(t.Properties))}
}

func (d *Data) GetDataFromProperties(properties []Property) {

	// Property deki mode ve type bilgilerini buraya atıyorum.
	for _, p := range properties {
		for i, v := range d.Names {
			if p.Name == v {
				d.Mods[i] = p.Mode
				d.Types[i] = p.Type
				break
			}
		}
	}
}

func (d *Data) CheckMods(t *Table) {
	// Find which index of data has mod
	for i, v := range d.Mods {
		if v != "" {

			// PK
			switch d.Mods[i] {
			case "PK":
				if d.Types[i] == "int" {
					// Eğer PK doğru kullanılmış ise bunun değeri otomatik atanacaktır.
					d.Values[i] = len(t.Get()) + 1
				} else {
					fmt.Println(errorhandler.GetErrorMods(1, d.Names[i]))
					os.Exit(1)
				}
				break
			}

		}
	}

}
