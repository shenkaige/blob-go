package db

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/core/errors"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/sha3"
	"log"
)

//NewDb creates a new database with chosen dialect.
func NewDb(file string, dialect string) *xorm.Engine {
	orm, err := xorm.NewEngine(dialect, file)
	if err != nil {
		log.Println(err)
	}
	return orm
}

//GetCore gets core information from db.
func GetCore(sql *xorm.Engine) CoreDb {
	var coreData CoreDb
	if ok, _ := sql.Get(&coreData); ok {
		return coreData
	}
	return CoreDb{}
}

//GetCoreFunc returns a function that get core information using defined sql.
func GetCoreFunc(sql *xorm.Engine) func() CoreDb {
	return func() CoreDb {
		var coreData CoreDb
		if ok, _ := sql.Get(&coreData); ok {
			return coreData
		}
		return CoreDb{}
	}
}

//AuthUser authenticates the user's identity.
func AuthUser(usernm string, passwd string, sql *xorm.Engine) (bool, error) {
	userData := UserDb{Username: usernm}
	if ok, _ := sql.Get(&userData); ok {
		if userData.Id != 0 {
			hasher := sha3.New512()
			hasher.Write([]byte(passwd))
			if string(hasher.Sum(nil)) == userData.Passwd {
				return true, nil
			}
			return false, errors.New("passwd not match")
		}
		return false, errors.New("user not found")
	}
	return false, errors.New("database err")
}

//GetIndex gets index of the first 10 posts.
func GetIndex(sql *xorm.Engine) (*[]PostDb, bool) {
	var datas []PostDb
	if err := sql.Desc("id").Limit(10).Find(&datas); err == nil {
		return &datas, true
	}
	return nil, false
}

//GetIndexByCategory gets index of the first 10 posts of the destinated category.
func GetIndexByCategory(categ string, sql *xorm.Engine) (*[]PostDb, bool) {
	var datas []PostDb
	if err := sql.Desc("id").Limit(10).Find(&datas, &PostDb{Category: categ}); err == nil && len(datas) != 0 {
		return &datas, true
	}
	return nil, false
}

//GetIndexByAuthor gets index of the first 10 posts of the destinated author.
func GetIndexByAuthor(autho string, sql *xorm.Engine) (*[]PostDb, bool) {
	var datas []PostDb
	if err := sql.Desc("id").Limit(10).Find(&datas, &PostDb{Author: autho}); err == nil && len(datas) != 0 {
		return &datas, true
	}
	return nil, false
}

//GetPost gets post from ID.
func GetPost(id int, sql *xorm.Engine) (*PostDb, bool) {
	post := PostDb{Id: id}
	if ok, _ := sql.Get(&post); ok && post.Title != "" {
		return &post, true
	}
	return &PostDb{}, false
}

//GetOverview gets post number.
func GetOverview(sql *xorm.Engine) (*OverviewDb, bool) {
	post := new(PostDb)
	if postCount, err := sql.Count(post); err == nil {
		return &OverviewDb{int(postCount), 0}, true
	}
	return &OverviewDb{}, false
}
