package gojson

// Database struct
type Database struct {
	DatabaseName string
	Path         string
	Tables       []*Table
}

// Table Struct
type Table struct {
	TableName    string
	Properties   []Property
	PathDatabase string
}

// Json Properties Struct
type Property struct {
	Name string
	Type string
	Mode string
}
type Data struct {
	Names  []string
	Values []interface{}
	Types  []string
	Mods   []string
}
