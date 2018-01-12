package index

import (
	"github.com/blob-go/blob-go/db"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var fzfResp = mvc.Response{Code: 404}

//IndexController is the controller to / page.
type IndexController struct {
	Sql *xorm.Engine
	Ctx iris.Context
}

//Get is the function when / is called.
func (c *IndexController) Get() mvc.Result {
	return c.GetBy(1)
}

//Get is the function when / is called.
func (c *IndexController) GetBy(pageid int) mvc.Result {
	if index, ok := db.GetIndexBy(pageid, c.Sql); ok {
		return mvc.View{
			Name: "index.html",
			Data: IndexStruct{
				Index:     pageid,
				IndexData: *index,
			},
		}
	}
	return fzfResp
}

//GetCategory is the function when /category/ is called.
func (c *IndexController) GetCategory() mvc.Result {
	return mvc.Response{Path: "/"}
}

//GetCategoryBy is the function when /category/<string> is called.
func (c *IndexController) GetCategoryBy(categ string) mvc.Result {
	return mvc.Response{Path: "/category/" + categ + "/1"}
}

//GetCategoryBy is the function when /category/<string> is called.
func (c *IndexController) GetCategoryByBy(categ string, pageid int) mvc.Result {
	if index, ok := db.GetIndexByCategoryBy(pageid, categ, c.Sql); ok {
		return mvc.View{
			Name: "index.html",
			Data: IndexStruct{
				Index:     pageid,
				IndexData: *index,
			},
		}
	}
	return fzfResp
}

//GetAuthor is the function when /author/ is called.
func (c *IndexController) GetAuthor() mvc.Result {
	return mvc.Response{Path: "/"}
}

func (c *IndexController) GetAuthorBy(autho string) mvc.Result {
	return mvc.Response{Path: "/author/" + autho + "/1"}
}

//GetAuthorBy is the function when /author/<string> is called.
func (c *IndexController) GetAuthorByBy(autho string, pageid int) mvc.Result {
	if index, ok := db.GetIndexByAuthorBy(pageid, autho, c.Sql); ok {
		return mvc.View{
			Name: "index.html",
			Data: IndexStruct{
				Index:     pageid,
				IndexData: *index,
			},
		}
	}
	return fzfResp
}
