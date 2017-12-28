package index

import (
	"../db"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/mvc"
)

type IndexController struct {
	mvc.C
	Sql *xorm.Engine
}

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

func (c *IndexController) GetCategory() mvc.Result {
	return mvc.Response{Path: "/"}
}

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

func (c *IndexController) GetAuthor() mvc.Result {
	return mvc.Response{Path: "/"}
}

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
