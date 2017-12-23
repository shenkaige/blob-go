package db

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func NewDb(file string, dialect string) *xorm.Engine {
	orm, err := xorm.NewEngine(dialect, file)
	if err != nil {
		log.Println(err)
	}
	return orm
}

func GetIndex(sql *xorm.Engine) (*[]PostDb, bool) {
	var datas []PostDb
	if err := sql.Limit(10).Desc("id").Find(&datas); err == nil {
		return &datas, true
	}
	return nil, false
}

func GetPost(id int, sql *xorm.Engine) (*PostDb, bool) {
	post := PostDb{Id: id}
	if ok, _ := sql.Get(&post); ok {
		return &post, true
	}
	return &PostDb{}, false
}

func GetOverview(sql *xorm.Engine) (*OverviewDb, bool) {
	post := new(PostDb)
	if postCount, err := sql.Count(post); err == nil {
		return &OverviewDb{int(postCount), 0}, true
	}
	return &OverviewDb{}, false
}
