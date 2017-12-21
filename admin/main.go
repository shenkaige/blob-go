package admin

import (
	"../core"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/mvc"
)

type AdminController struct {
	mvc.C
	Sql *xorm.Engine
}

func (c *AdminController) Get() mvc.Result {
	return mvc.View{
		Name: "overview.html",
		Data: AdminStruct{
			Core: *(core.GetCore()),
		},
	}
}

func (c *AdminController) GetOverview() mvc.Result {
	return mvc.View{
		Name: "overview.html",
		Data: AdminStruct{
			Core: *(core.GetCore()),
		},
	}
}
