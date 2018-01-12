package db

import (
	"errors"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"//Imports package necessary for xorm.
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

//SetCore is the function to set core information of blob.
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
func AuthUserName(userName string, hash string, sql *xorm.Engine) (bool, int, error) {
	userData := UserDb{Username: userName}
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

//AuthUserID authenticates user with ID and hash.
func AuthUserID(userID int, hash string, sql *xorm.Engine) (bool, error) {
	userData := UserDb{Id: userID}
	if ok, _ := sql.Get(&userData); ok {
		if userData.Username != "" {
			if hash == userData.Passwd {
				return true, nil
			}
			return false, errors.New("password not match")
		}
		return false, errors.New("user not found")
	}
	return false, errors.New("database err")
}

//GetIndexBy gets index of the first 10 posts.
func GetIndexBy(index int, sql *xorm.Engine) (*[]PostDb, bool) {
	var arrPostDb []PostDb
	if index > 0 {
		if err := sql.Desc("id").Limit(10, 10*(index-1)).Find(&arrPostDb); err == nil && len(arrPostDb) != 0 {
			return &arrPostDb, true
		}
	}
	return nil, false
}

//GetIndexByCategoryBy gets index of the first 10 posts of the destine category.
func GetIndexByCategoryBy(index int, category string, sql *xorm.Engine) (*[]PostDb, bool) {
	var arrPostDb []PostDb
	if index > 0 {
		if err := sql.Desc("id").Limit(10, 10*(index-1)).Find(&arrPostDb, &PostDb{Category: category}); err == nil && len(arrPostDb) != 0 {
			return &arrPostDb, true
		}
	}
	return nil, false
}

//GetIndexByAuthorBy gets index of the first 10 posts of the destine author.
func GetIndexByAuthorBy(index int, author string, sql *xorm.Engine) (*[]PostDb, bool) {
	var arrPostDb []PostDb
	if index > 0 {
		if err := sql.Desc("id").Limit(10, 10*(index-1)).Find(&arrPostDb, &PostDb{Author: author}); err == nil && len(arrPostDb) != 0 {
			return &arrPostDb, true
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

//SetPost updates post from PostDb.
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
