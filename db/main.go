package db

import (
	"errors"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

//NewDb creates a new database with chosen dialect.
func NewDb(file string, dialect string) *xorm.Engine {
	orm, err := xorm.NewEngine(dialect, file)
	if err != nil {
		println(err)
	}
	return orm
}

//GetInfo gets core information from db.
func GetInfo(sql *xorm.Engine) CoreDb {
	var coreData CoreDb
	if ok, _ := sql.ID(1).Get(&coreData); ok {
		return coreData
	}
	return CoreDb{}
}

//GetInfoFunc returns a function that get core information using defined sql.
func GetInfoFunc(sql *xorm.Engine) func() CoreDb {
	return func() CoreDb {
		var coreData CoreDb
		if ok, _ := sql.ID(1).Get(&coreData); ok {
			return coreData
		}
		return CoreDb{}
	}
}

func SetCore(title string, subTitle string, sql *xorm.Engine) bool {
	coreData := CoreDb{
		Id:       1,
		Title:    title,
		SubTitle: subTitle,
	}
	_, err := sql.Update(&coreData)
	return err == nil
}

//AuthUserName authenticates the user's identity.
func AuthUserName(usernm string, hash string, sql *xorm.Engine) (bool, int, error) {
	userData := UserDb{Username: usernm}
	if ok, _ := sql.Get(&userData); ok {
		if userData.Id != 0 {
			if hash == userData.Passwd {
				return true, userData.Id, nil
			}
			return false, 0, errors.New("passwd not match")
		}
		return false, 0, errors.New("user not found")
	}
	return false, 0, errors.New("database err")
}

func AuthUserID(userid int, hash string, sql *xorm.Engine) (bool, error) {
	userData := UserDb{Id: userid}
	if ok, _ := sql.Get(&userData); ok {
		if userData.Username != "" {
			if hash == userData.Passwd {
				return true, nil
			}
			return false, errors.New("passwd not match")
		}
		return false, errors.New("user not found")
	}
	return false, errors.New("database err")
}

//GetIndexBy gets index of the first 10 posts.
func GetIndexBy(index int, sql *xorm.Engine) (*[]PostDb, bool) {
	var datas []PostDb
	if index > 0 {
		if err := sql.Desc("id").Limit(10, 10*(index-1)).Find(&datas);
			err == nil && len(datas) != 0 {
			return &datas, true
		}
	}
	return nil, false
}

//GetIndexByCategoryBy gets index of the first 10 posts of the destinated category.
func GetIndexByCategoryBy(index int, categ string, sql *xorm.Engine) (*[]PostDb, bool) {
	var datas []PostDb
	if index > 0 {
		if err := sql.Desc("id").Limit(10, 10*(index-1)).Find(&datas, &PostDb{Category: categ});
			err == nil && len(datas) != 0 {
			return &datas, true
		}
	}
	return nil, false
}

//GetIndexByAuthorBy gets index of the first 10 posts of the destinated author.
func GetIndexByAuthorBy(index int, autho string, sql *xorm.Engine) (*[]PostDb, bool) {
	var datas []PostDb
	if index > 0 {
		if err := sql.Desc("id").Limit(10, 10*(index-1)).Find(&datas, &PostDb{Author: autho});
			err == nil && len(datas) != 0 {
			return &datas, true
		}
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

func SetPost(id int, post *PostDb, sql *xorm.Engine) bool {
	_, err := sql.ID(id).Update(post)
	return err == nil
}

//GetOverview gets post number.
func GetOverview(sql *xorm.Engine) (*OverviewDb, bool) {
	post := new(PostDb)
	if postCount, err := sql.Count(post); err == nil {
		return &OverviewDb{int(postCount), 0}, true
	}
	return &OverviewDb{}, false
}
