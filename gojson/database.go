package gojson

import (
	"fmt"
	"os"
)

// initialize Database
func CreateDatabase(dbName string, path string) Database {
	return Database{DatabaseName: dbName, Path: path}
}

// Adds table to the database
func (d *Database) AddTable(table *Table) {
	// Eğer table'ı sonra değiştiriceksen * kullanmalısın ama şuanlık sıkıntı yok
	table.PathDatabase = d.Path + d.DatabaseName + "/"
	d.Tables = append(d.Tables, table)
}

func (d *Database) CreateFiles() {
	// 777 => 7 7 7 => 111 111 111 => ilk 1111 root ikinci 1111 group, son 1111 ise diğer kullanıcılar
	// 111 de ilk kısım exeute yetkisi ikinci kısım yazma yetkisi diğeri ise okuma yetkisidir

	DatabasePath := d.Path + d.DatabaseName
	if !HasFile(DatabasePath) {
		if err := os.Mkdir(d.Path+d.DatabaseName, 777); err != nil {
			fmt.Println(err)
		}

		// Direkt Table'ları oluşturuyoruz.
		for _, t := range d.Tables {
			TablePath := DatabasePath + "/" + t.TableName + ".json"
			// Database klasörünün içine table dosyası oluşturulucak
			file, err := os.Create(TablePath)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
		}
	} else {
		// Eğer Db klasörü var ise içindeki tablelar tam olarak var mı bakıyoruz.
		// Tam değil ise eksikleri oluşturuyoruz
		for _, t := range d.Tables {
			TablePath := DatabasePath + "/" + t.TableName + ".json"
			if !HasFile(TablePath) {
				// Database klasörünün içine table dosyası oluşturulucak
				file, err := os.Create(TablePath)
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()
			}
		}
	}
}

// Buraya bir tane gettable func yaz table döndürsün. Belki ilerde table için query yazma felan yaparım.
