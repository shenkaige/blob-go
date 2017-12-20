package post

import (
	"../core"
	"../db"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/mvc"
	_ "github.com/mattn/go-sqlite3"
)

type PostController struct {
	mvc.C
	Sql *xorm.Engine
}

func (c *PostController) Get() mvc.Result {
	post := db.PostDb{Id: 1}
	if ok, _ := c.Sql.Get(&post); ok {
		return mvc.View{
			Name: "post.html",
			Data: PostStruct{
				Core:     *(core.GetCore()),
				Title:    post.Title,
				SubTitle: post.SubTitle,
				Author:   post.Author,
				Category: post.Category,
				Content:  post.Content,
			},
		}
	}
	return mvc.Response{
		Code: 404,
	}
}

func (c *PostController) GetBy(id int) mvc.Result {
	post := db.PostDb{Id: id}
	if ok, _ := c.Sql.Get(&post); ok {
		return mvc.View{
			Name: "post.html",
			Data: PostStruct{
				Core:     *(core.GetCore()),
				Title:    post.Title,
				SubTitle: post.SubTitle,
				Author:   post.Author,
				Category: post.Category,
				Content:  post.Content,
			},
		}
	}
	return mvc.Response{
		Code: 404,
	}
}
