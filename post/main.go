package post

import (
	"github.com/blob-go/blob-go/db"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/mvc"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
)

var fzfResp = mvc.Response{Code: 404}

//Controller is the controller to /post page.
type Controller struct {
	SQL *xorm.Engine
}

//Get is the function when /post/ is called.
func (c *Controller) Get() mvc.Result {
	return mvc.Response{Path: "/"}
}

//GetBy is the function when /post/<id> is called.
func (c *Controller) GetBy(id int) mvc.Result {
	if post, ok := db.GetPost(id, c.SQL); ok {
		return mvc.View{
			Name: "post.html",
			Data: Struct{
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
