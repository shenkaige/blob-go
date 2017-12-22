package admin

import (
	"../core"
	"../db"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/mvc"
)

type AdminController struct {
	mvc.C
	Sql *xorm.Engine
}

func (c *AdminController) Get() mvc.Result {
	if overv, ok := db.GetOverview(c.Sql); ok {
		return mvc.View{
			Name: "overview.html",
			Data: OverviewStruct{
				Core:    *(core.GetCore()),
				SubData: *overv,
			},
		}
	}
	return mvc.View{
		Name: "overview.html",
		Data: OverviewStruct{
			Core:    *(core.GetCore()),
			SubData: db.OverviewDb{},
		},
	}
}

func (c *AdminController) GetOverview() mvc.Result {
	if overv, ok := db.GetOverview(c.Sql); ok {
		return mvc.View{
			Name: "overview.html",
			Data: OverviewStruct{
				Core:    *(core.GetCore()),
				SubData: *overv,
			},
		}
	}
	return mvc.View{
		Name: "overview.html",
		Data: OverviewStruct{
			Core:    *(core.GetCore()),
			SubData: db.OverviewDb{},
		},
	}
}
