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
    + `gojson.DataInit([]string{"username"}, []interface{}{"Boran"}, &Table1)`: This one takes 3 arguments according to table1 in properties. One of the values it takes is a string and the other is an interface array and the last one is the table that you are going to save.
        + The first argument will cover the property names to which you will add the value.
        + The second argument contains their values.
        + The last argument is the table that you are going to save. Otherwise it wont save.
    + `Table1.Save(data Data)`: This one takes one argument and it's the gojson data struct.
    + `Table1.Get()`: This function returns all registered data as []map[string]interface{}.
    + `Table1.Update(uniqueStr string, uniqueStrValue interface{}, data Data)`: This takes 3 argument.
        + The first argument tells the function which data to change. Preferably unique data should be used like userId
        + The second argument represents the value of the property name we get in the first argument, like 1, 2, hi
            + if you write userId in the first and 2 in the second. You change the userId 2 from inside the json.
        + The last one is the Data struct. While saving, you were entering the property name and values to add to the data struct, here you will enter the property name and values to update.
+ Now a setup example for the project.

```
package main

import (
	"fmt"

	"github.com/cetinboran/gojson/gojson"
)

func main() {
	// Database1
	Database := gojson.CreateDatabase("Database", "./")

	// Table1
	Table1 := gojson.CreateTable("users")
	Table1.AddProperty("userId", "int", "PK")
	Table1.AddProperty("username", "string", "")
	Table1.AddProperty("password", "string", "")

	// Adds the table to the database
	Database.AddTable(&Table1)

    // Creates the json files.
	Database.CreateFiles()

    // You don't need to include the whole property name, gojson will do it for you and give the initial values.
    Table1.Save(gojson.DataInit([]string{"username", "password"}, []interface{}{"Boran","cetin"}, &Table1))
    Table1.Save(gojson.DataInit([]string{"username", "password"}, []interface{}{"Arzu","1597"}, &Table1))
	
    // You can get the values of json files using Table1.Get() it returns map.
	fmt.Println(Table1.Get())

    // Update Example
    // I changed the username of the data with a userId to "arzu"
    newData := gojson.DataInit([]string{"username"}, []interface{}{"arzu"}, &Table1)
	Table1.Update("userId", 1, newData)
}

```

+ If we write the following code to the terminal `go run .\main.go`.
+ we can see the result below.
    + map[password:ThePass userId:1 username:cetinboran] map[password:ThePass userId:2 username:arzu]]

## What's new
+ 23.08.2023
    + `Table1.Find(PropertyName string, PropertyValue interface{})`: Added find function to table. Now giving the propert name and its value []map[string]interface{}{} 
    value has a return value. If it returns blank, nothing was found.
    + Note that if the property type is int and you enter value as a string, nothing will be returned.

# Contact

[<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/github.svg' alt='github' height='40'>](https://github.com/cetinboran)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/linkedin.svg' alt='linkedin' height='40'>](https://www.linkedin.com/in/cetinboran-mesum/)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/instagram.svg' alt='instagram' height='40'>](https://www.instagram.com/2023an_m/)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/twitter.svg' alt='twitter' height='40'>](https://twitter.com/2023anM)  