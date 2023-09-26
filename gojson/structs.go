package gojson

// Database struct
type Database struct {
	DatabaseName string
	path         string
	tables       map[string]*Table
}

// Table Struct
type Table struct {
	Name    string
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
