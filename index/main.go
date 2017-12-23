package index

import (
	"../core"
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
				Core:    *(core.GetCore()),
				SubData: *index,
			},
		}
	}
	return mvc.View{
		Name: "index.html",
		Data: IndexStruct{
			Core:    *(core.GetCore()),
			SubData: []db.PostDb{},
		},
	}
}
