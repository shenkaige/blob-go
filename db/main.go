package db

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type PostDb struct {
	Id       int
	Title    string
	SubTitle string
	Author   string
	Category string
	Content  string
}

func NewDb(file string, dialect string) *xorm.Engine {
	orm, err := xorm.NewEngine(dialect, file)
	if err != nil {
		log.Println(err)
	}
	return orm
}
