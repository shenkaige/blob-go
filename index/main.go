package index

import (
	"../db"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

//IndexController is the controller to / page.
type IndexController struct {
	Sql *xorm.Engine
	Ctx iris.Context
}

//Get is the function when / is called.
func (c *IndexController) Get(ctx iris.Context) mvc.Result {
	page := getURLParamInt(ctx, "page", 1)
	if index, ok := db.GetIndexBy(page, c.Sql); ok {
		return mvc.View{
			Name: "index.html",
			Data: IndexStruct{
				Index:     page,
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

//GetCategoryBy is the function when /category/<string> is called.
func (c *IndexController) GetCategoryBy(ctx iris.Context, categ string) mvc.Result {
	page := getURLParamInt(ctx, "page", 1)
	if index, ok := db.GetIndexByCategoryBy(page, categ, c.Sql); ok {
		return mvc.View{
			Name: "index.html",
			Data: IndexStruct{
				Index:     page,
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

//GetAuthorBy is the function when /author/<string> is called.
func (c *IndexController) GetAuthorBy(ctx iris.Context, autho string) mvc.Result {
	page := getURLParamInt(ctx, "page", 1)
	if index, ok := db.GetIndexByAuthorBy(page, autho, c.Sql); ok {
		return mvc.View{
			Name: "index.html",
			Data: IndexStruct{
				Index:     page,
				IndexData: *index,
			},
		}
	}
	return mvc.View{
		Name: "httperr/404.html",
		Code: 404,
	}
}

func getURLParamInt(ctx iris.Context, param string, defau int) int {
	page, err := ctx.URLParamInt(param)
	if err != nil {
		return defau
	}
	return page
}
