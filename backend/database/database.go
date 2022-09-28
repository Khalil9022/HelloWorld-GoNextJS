package database

type Database struct {
	Data []string
}

var dbase *Database

func StartDB() {
	dbase = &Database{Data: []string{}}
}

func GetDB() *Database {
	return dbase
}

func (db *Database) Append(text string) {
	db.Data = append(db.Data, text)
}
