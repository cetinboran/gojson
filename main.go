package main

import (
	"fmt"

	"github.com/cetinboran/gojson/gojson"
)

func main() {
	// Database1
	Database1 := gojson.CreateDatabase("Database1", "../")

	// Table1
	Table1 := gojson.CreateTable("users")
	Table1.AddProperty("userId", "1")
	Table1.AddProperty("username", "Boran")
	Table1.AddProperty("password", "123")

	// Table1
	Table2 := gojson.CreateTable("config")
	Table2.AddProperty("secretKey", "1645854515")
	Table2.AddProperty("Timeout", "4")

	// Adds the table to the database
	Database1.AddTable(Table1)
	Database1.AddTable(Table2)

	fmt.Println(Database1)
	fmt.Println(Table1)
}
