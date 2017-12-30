package index

import (
	"../db"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/mvc"
)

//IndexController is the controller to / page.
type IndexController struct {
	mvc.C
	Sql *xorm.Engine
}

//Get is the function when / is called.
func (c *IndexController) Get() mvc.Result {
	if index, ok := db.GetIndex(c.Sql); ok {
		return mvc.View{
			Name: "index.html",
			Data: IndexStruct{
				IndexData: *index,
			},
		}
	}
	return mvc.View{
		Name: "httperr/404.html",
		Code: 404,
	}
}

//GetCategory is the function when /category/ is called.
func (c *IndexController) GetCategory() mvc.Result {
	return mvc.Response{Path: "/"}
}

//GetCategoryBy is the function when /category/<id> is called.
func (c *IndexController) GetCategoryBy(categ string) mvc.Result {
	if index, ok := db.GetIndexByCategory(categ, c.Sql); ok {
		return mvc.View{
			Name: "index.html",
			Data: IndexStruct{
				IndexData: *index,
			},
		}
	}
	return mvc.View{
		Name: "httperr/404.html",
		Code: 404,
	}
}

//GetAuthor is the function when /author/ is called.
func (c *IndexController) GetAuthor() mvc.Result {
	return mvc.Response{Path: "/"}
}

//GetAuthorBy is the function when /author/<id> is called.
func (c *IndexController) GetAuthorBy(autho string) mvc.Result {
	if index, ok := db.GetIndexByAuthor(autho, c.Sql); ok {
		return mvc.View{
			Name: "index.html",
			Data: IndexStruct{
				IndexData: *index,
			},
		}
	}
	return mvc.View{
		Name: "httperr/404.html",
		Code: 404,
	}
}
