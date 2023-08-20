package gojson

// initialize Database
func CreateDatabase(dbName string, path string) Database {
	return Database{DatabaseName: dbName, Path: path}
}

// Adds table to the database
func (d *Database) AddTable(table Table) {
	// Eğer table'ı sonra değiştiriceksen * kullanmalısın ama şuanlık sıkıntı yok
	d.Tables = append(d.Tables, table)
}
