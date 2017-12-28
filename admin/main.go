package admin

import (
	"../db"
	"../index"
	"../post"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"html/template"
)

type AdminController struct {
	mvc.C
	Manager *sessions.Sessions
	Session *sessions.Session
	Sql     *xorm.Engine
}

func (c *AdminController) Get() mvc.Result {
	return c.GetOverview()
}

func (c *AdminController) GetOverview() mvc.Result {
	if overv, ok := db.GetOverview(c.Sql); ok {
		return mvc.View{
			Name: "admin/overview.html",
			Data: OverviewStruct{
				OverviewData: *overv,
			},
		}
	}
	return mvc.View{
		Name: "admin/overview.html",
		Data: OverviewStruct{
			OverviewData: db.OverviewDb{},
		},
	}
}

func (c *AdminController) GetPost() mvc.Result {
	if ind, ok := db.GetIndex(c.Sql); ok {
		return mvc.View{
			Name: "admin/post_index.html",
			Data: index.IndexStruct{
				IndexData: *ind,
			},
		}
	}
	return mvc.View{
		Name: "admin/post_index.html",
		Data: index.IndexStruct{
			IndexData: []db.PostDb{},
		},
	}
}

func (c *AdminController) GetPostEdit() mvc.Result {
	return mvc.View{
		Name: "admin/post_edit.html",
		Data: post.PostStruct{
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
				Title:    pos.Title,
				SubTitle: pos.SubTitle,
				Author:   pos.Author,
				Category: pos.Category,
				Content:  template.HTML(pos.Content),
			},
		}
	}
	return mvc.Response{
		Code: 404,
	}
}

func (c *AdminController) GetSetting() mvc.Result {
	return mvc.View{
		Name: "admin/setting.html",
		Data: db.GetCore(c.Sql),
	}
}
