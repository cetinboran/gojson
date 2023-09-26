package gojson

import (
	"log"
	"os"

	"github.com/cetinboran/gojson/errorhandler"
)

// initialize Database
func CreateDatabase(dbName string, path string) Database {
	return Database{DatabaseName: dbName, path: path, Tables: make(map[string]*Table)}
}

// Adds table to the database
func (d *Database) AddTable(table *Table) {

	for _, t := range d.Tables {
		if t.Name == table.Name {
			log.Fatal(errorhandler.GetErrorTable(4, t.Name))
		}
	}

	// Eğer table'ı sonra değiştiriceksen * kullanmalısın ama şuanlık sıkıntı yok
	table.PathDatabase = d.path + d.DatabaseName + "/"
	d.Tables[table.Name] = table
}

func (d *Database) CreateFiles() {
	// 777 => 7 7 7 => 111 111 111 => ilk 1111 root ikinci 1111 group, son 1111 ise diğer kullanıcılar
	// 111 de ilk kısım exeute yetkisi ikinci kısım yazma yetkisi diğeri ise okuma yetkisidir

	DatabasePath := d.path + d.DatabaseName
	if !HasFile(DatabasePath) {
		if err := os.Mkdir(d.path+d.DatabaseName, 777); err != nil {
			log.Fatal(err)
		}

		// Direkt Table'ları oluşturuyoruz.
		for _, t := range d.Tables {
			TablePath := DatabasePath + "/" + t.Name + ".json"
			// Database klasörünün içine table dosyası oluşturulucak
			file, err := os.Create(TablePath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			err = os.WriteFile(TablePath, []byte("[]"), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		// Eğer Db klasörü var ise içindeki tablelar tam olarak var mı bakıyoruz.
		// Tam değil ise eksikleri oluşturuyoruz
		for _, t := range d.Tables {
			TablePath := DatabasePath + "/" + t.Name + ".json"
			if !HasFile(TablePath) {
				// Database klasörünün içine table dosyası oluşturulucak
				file, err := os.Create(TablePath)
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()

				err = os.WriteFile(TablePath, []byte("[]"), 0644)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

}

// Buraya bir tane gettable func yaz table döndürsün. Belki ilerde table için query yazma felan yaparım.
