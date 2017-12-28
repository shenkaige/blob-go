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

func GetCore(sql *xorm.Engine) CoreDb {
	var coreData CoreDb
	if ok, _ := sql.Get(&coreData); ok {
		return coreData
	}
	return CoreDb{}
}

func GetCoreFunc(sql *xorm.Engine) func() CoreDb {
	return func() CoreDb {
		var coreData CoreDb
		if ok, _ := sql.Get(&coreData); ok {
			return coreData
		}
		return CoreDb{}
	}
}

func GetIndex(sql *xorm.Engine) (*[]PostDb, bool) {
	var datas []PostDb
	if err := sql.Desc("id").Limit(10).Find(&datas); err == nil {
		return &datas, true
	}
	return nil, false
}

func GetIndexByCategory(categ string, sql *xorm.Engine) (*[]PostDb, bool) {
	var datas []PostDb
	if err := sql.Desc("id").Limit(10).Find(&datas, &PostDb{Category: categ}); err == nil && len(datas) != 0 {
		return &datas, true
	}
	return nil, false
}

func GetIndexByAuthor(autho string, sql *xorm.Engine) (*[]PostDb, bool) {
	var datas []PostDb
	if err := sql.Desc("id").Limit(10).Find(&datas, &PostDb{Author: autho}); err == nil && len(datas) != 0 {
		return &datas, true
	}
	return nil, false
}

func GetPost(id int, sql *xorm.Engine) (*PostDb, bool) {
	post := PostDb{Id: id}
	if ok, _ := sql.Get(&post); ok && post.Title != "" {
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
