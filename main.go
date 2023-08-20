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
	Table1.AddProperty("userId", "int")
	Table1.AddProperty("username", "string")
	Table1.AddProperty("password", "string")

	// Table1
	Table2 := gojson.CreateTable("config")
	Table2.AddProperty("secretKey", "string")
	Table2.AddProperty("Timeout", "int")

	// Adds the table to the database
	Database1.AddTable(Table1)
	Database1.AddTable(Table2)

	Database1.Start()

	fmt.Println(Database1)
	// fmt.Println(Table1)
}
