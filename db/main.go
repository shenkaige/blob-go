package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Db struct {
	dialect string
	file string
}

func NewSqlite(file string) *Db {
	return &Db{
		dialect: "sqlite3",
		file: file,
	}
}

func (db *Db) Create(dataStruct interface{}) {
	dbase, err := gorm.Open(db.dialect, db.file)
	if err != nil {
		panic("failed to connect database")
	}
	defer dbase.Close()
	dbase.AutoMigrate(&dataStruct)
	dbase.Create(&dataStruct)
}