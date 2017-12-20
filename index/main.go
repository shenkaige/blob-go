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
	return mvc.View{
		Name: "index.html",
		Data: IndexStruct{
			Core:    *(core.GetCore()),
			SubData: *GetIndex(c.Sql),
		},
	}
}

func GetIndex(sql *xorm.Engine) *[]db.PostDb {
	var datas []db.PostDb
	sql.Limit(10).Desc("id").Find(&datas)
	return &datas
}
