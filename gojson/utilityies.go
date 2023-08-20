package gojson

import (
	"fmt"
	"os"
)

// For Database.go
func HasFile(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		fmt.Println(err)
	}

	return false
}
