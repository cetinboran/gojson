package gojson

import (
	"fmt"
	"os"
	"strings"

	"github.com/cetinboran/gojson/errorhandler"
)

// Here we compare property names with incoming names if they are not the same we send an error. At the same time, we send an error if the same property name is written more than once.
func CheckNames(names []string, t *Table) bool {
	check := false

	for _, v := range names {
		if strings.Count(strings.Join(names, " "), v) != 1 {
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
	return true
}

func CheckValues(values []interface{}, t *Table) bool {

	return true
}
