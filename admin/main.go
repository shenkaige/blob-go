package admin

import (
	"../core"
	"../db"
	"../index"
	"../post"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/mvc"
)

type AdminController struct {
	mvc.C
	Sql *xorm.Engine
}

func (c *AdminController) Get() mvc.Result {
	return c.GetOverview()
}

func (c *AdminController) GetOverview() mvc.Result {
	if overv, ok := db.GetOverview(c.Sql); ok {
		return mvc.View{
			Name: "admin/overview.html",
			Data: OverviewStruct{
				Core:    *(core.GetCore()),
				SubData: *overv,
			},
		}
	}
	return mvc.View{
		Name: "admin/overview.html",
		Data: OverviewStruct{
			Core:    *(core.GetCore()),
			SubData: db.OverviewDb{},
		},
	}
}

func (c *AdminController) GetPost() mvc.Result {
	if ind, ok := db.GetIndex(c.Sql); ok {
		return mvc.View{
			Name: "admin/post_index.html",
			Data: index.IndexStruct{
				Core:    *(core.GetCore()),
				SubData: *ind,
			},
		}
	}
	return mvc.View{
		Name: "admin/post_index.html",
		Data: index.IndexStruct{
			Core:    *(core.GetCore()),
			SubData: []db.PostDb{},
		},
	}
}

func (c *AdminController) GetPostEdit() mvc.Result {
	return mvc.View{
		Name: "admin/post_edit.html",
		Data: post.PostStruct{
			Core:     *(core.GetCore()),
			Title:    "New Post",
			SubTitle: "This is a Subtitle",
			Author:   "Author",
			Category: "",
			Content:  "",
		},
	}
}

func (c *AdminController) GetPostEditBy(id int) mvc.Result {
	if pos, ok := db.GetPost(id, c.Sql); ok {
		return mvc.View{
			Name: "admin/post_edit.html",
			Data: post.PostStruct{
				Core:     *(core.GetCore()),
				Title:    pos.Title,
				SubTitle: pos.SubTitle,
				Author:   pos.Author,
				Category: pos.Category,
				Content:  pos.Content,
			},
		}
	}
	return mvc.Response{
		Code: 404,
	}
}
