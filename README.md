# GoJson

## What is This?
+ In this project, you can create your own json database with just a few lines of code.

## What can you do?
+ You can create the database in one line and you can add tables to the database you created in a few lines
+ You can specify json place holders, value types of inputs, and modes specific to inputs.

## How to Install?
+ Open your project and write the code below in the terminal. The project will be added.
    + `go get github.com/cetinboran/gojson@v1.0.0`

## How to Use?
+ First, let me explain the methods.
    + `Database1 := gojson.CreateDatabase(DatabaseName string, PATH string)`: Initializes the Database
        + First argument is the name of the database
        + The second argument is the path where the database will be created.
    + `Table1 := gojson.CreateTable(tablename string)`: Initializes the Table
        + This function requests table name as argument.
    + `Table1.AddProperty(propertyName string, type string, mode string)`: This function adds table properties
        + The first argument property name
        + The second argument is the data type you will receive from the user -like int string-
        + The third argument is property mode.
            + Mode PK: It provides automatic id assignment.
    + `Database1.AddTable(&Table1)`: This function adds the table to the database. 
    + `Database1.CreateFiles()`: This creates database files.
    + `Table1.Save("userId,username,password", []interface{}{-1, "value2", "value3"})`: This one takes 2 arguments according to table1 in properties. One of the values it takes is a string and the other is an interface array.
        + First argument must contain all property names you create
        + The second argument contains their values. It should still include all
    + `Table1.Get()`: This function returns all registered data as []map[string]interface{}.
+ Now a setup example for the project.

```
package main

import (
	"fmt"

	"github.com/cetinboran/gojson/gojson"
)

func main() {
	// Database1
	Database1 := gojson.CreateDatabase("Database1", "./")

	// Table1
	Table1 := gojson.CreateTable("users")
	Table1.AddProperty("userId", "int", "PK")
	Table1.AddProperty("username", "string", "")
	Table1.AddProperty("password", "string", "")

	// Adds the table to the database
	Database1.AddTable(&Table1)

	Database1.CreateFiles()

	Table1.Save("userId,username,password", []interface{}{-1, "cetinboran", "ThePass"})
	Table1.Save("userId,username,password", []interface{}{-1, "arzu", "ThePass"})
	
	fmt.Println(Table1.Get())
}

```

+ If we write the following code to the terminal `go run .\main.go`.
+ we can see the result below.
    + map[password:ThePass userId:1 username:cetinboran] map[password:ThePass userId:2 username:arzu]]


# Contact

[<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/github.svg' alt='github' height='40'>](https://github.com/cetinboran)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/linkedin.svg' alt='linkedin' height='40'>](https://www.linkedin.com/in/cetinboran-mesum/)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/instagram.svg' alt='instagram' height='40'>](https://www.instagram.com/2023an_m/)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/twitter.svg' alt='twitter' height='40'>](https://twitter.com/2023anM)  