package post

import (
	"../db"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/mvc"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
)

var fzfResp = mvc.Response{Code: 404}

//PostController is the controller to /post page.
type PostController struct {
	Sql *xorm.Engine
}

//Get is the function when /post/ is called.
func (c *PostController) Get() mvc.Result {
	return mvc.Response{Path: "/"}
}

//GetBy is the function when /post/<id> is called.
func (c *PostController) GetBy(id int) mvc.Result {
	if post, ok := db.GetPost(id, c.Sql); ok {
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
	return fzfResp
}
