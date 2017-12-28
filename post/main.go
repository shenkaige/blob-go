package post

import (
	"../db"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/mvc"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
)

type PostController struct {
	mvc.C
	Sql *xorm.Engine
}

func (c *PostController) Get() mvc.Result {
	if post, ok := db.GetPost(1, c.Sql); ok {
		return mvc.View{
			Name: "post.html",
			Data: PostStruct{
				Title:    post.Title,
				SubTitle: post.SubTitle,
				Author:   post.Author,
				Category: post.Category,
				Content:  template.HTML(string(blackfriday.Run([]byte(post.Content)))),
			},
		}
	}
	return mvc.Response{
		Code: 404,
	}
}

func (c *PostController) GetBy(id int) mvc.Result {
	if post, ok := db.GetPost(id, c.Sql); ok {
		return mvc.View{
			Name: "post.html",
			Data: PostStruct{
				Title:    post.Title,
				SubTitle: post.SubTitle,
				Author:   post.Author,
				Category: post.Category,
				Content: template.HTML(string(blackfriday.Run([]byte(post.Content)))),
			},
		}
	}
	return mvc.Response{
		Code: 404,
	}
}
