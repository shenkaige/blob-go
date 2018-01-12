package index

import (
	"github.com/blob-go/blob-go/db"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var fzfResp = mvc.Response{Code: 404}

//Controller is the controller to / page.
type Controller struct {
	SQL *xorm.Engine
	Ctx iris.Context
}

//Get is the function when / is called.
func (c *Controller) Get() mvc.Result {
	return c.GetBy(1)
}

//GetBy is the function when /<int> is called.
func (c *Controller) GetBy(pageID int) mvc.Result {
	if index, ok := db.GetIndexBy(pageID, c.SQL); ok {
		return mvc.View{
			Name: "index.html",
			Data: Struct{
				Index:     pageID,
				IndexData: *index,
			},
		}
	}
	return fzfResp
}

//GetCategory is the function when /category/ is called.
func (c *Controller) GetCategory() mvc.Result {
	return mvc.Response{Path: "/"}
}

//GetCategoryBy is the function when /category/<string> is called.
func (c *Controller) GetCategoryBy(category string) mvc.Result {
	return mvc.Response{Path: "/category/" + category + "/1"}
}

//GetCategoryByBy is the function when /category/<string>/<int> is called.
func (c *Controller) GetCategoryByBy(category string, pageID int) mvc.Result {
	if index, ok := db.GetIndexByCategoryBy(pageID, category, c.SQL); ok {
		return mvc.View{
			Name: "index.html",
			Data: Struct{
				Index:     pageID,
				IndexData: *index,
			},
		}
	}
	return fzfResp
}

//GetAuthor is the function when /author/ is called.
func (c *Controller) GetAuthor() mvc.Result {
	return mvc.Response{Path: "/"}
}

//GetAuthorBy is the function when /author/<string> is called.
func (c *Controller) GetAuthorBy(author string) mvc.Result {
	return mvc.Response{Path: "/author/" + author + "/1"}
}

//GetAuthorByBy is the function when /author/<string>/<int> is called.
func (c *Controller) GetAuthorByBy(author string, pageID int) mvc.Result {
	if index, ok := db.GetIndexByAuthorBy(pageID, author, c.SQL); ok {
		return mvc.View{
			Name: "index.html",
			Data: Struct{
				Index:     pageID,
				IndexData: *index,
			},
		}
	}
	return fzfResp
}
