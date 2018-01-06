package admin

import (
	"../db"
	"../index"
	"../post"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"html/template"
	"github.com/kataras/iris"
)

//AdminController is the controller to /admin page.
type AdminController struct {
	Manager *sessions.Sessions
	Session *sessions.Session
	Ctx     iris.Context
	Sql     *xorm.Engine
}

//Get is the function when /admin/ is called.
func (c *AdminController) Get() mvc.Result {
	return c.GetOverview()
}

//GetOverview is the function when /admin/overview/ is called.
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

//GetPost is the function when /admin/post/ is called.
func (c *AdminController) GetPost() mvc.Result {
	if ind, ok := db.GetIndexBy(1, c.Sql); ok {
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

//GetPostEdit is the function when /admin/post/edit/ is called.
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

//GetPostEditBy is the function when /admin/post/edit/<id> is called.
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

//GetSetting is the function when /admin/setting/ is called.
func (c *AdminController) GetSetting() mvc.Result {
	return mvc.View{
		Name: "admin/setting.html",
		Data: db.GetCore(c.Sql),
	}
}

func (c *AdminController) getCurrentUserID() int {
	userID, _ := c.Session.GetIntDefault("UserID", 0)
	return userID
}

type AuthController struct {
	Manager *sessions.Sessions
	Session *sessions.Session
	Ctx     iris.Context
	Sql     *xorm.Engine
}

func (c *AuthController) GetLogin() mvc.Result {
	return mvc.View{
		Name: "admin/login.html",
		Data: iris.Map{"Title": "Login"},
	}
}

func (c *AuthController) PostLogin(ctx iris.Context) mvc.Result {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if ok, err := db.AuthUser(username, password, c.Sql); ok {
		return mvc.Response{
			Path: "/admin",
		}
	} else {
		println(err.Error())
		return mvc.View{
			Name: "admin/login.html",
			Data: iris.Map{"Title": "Login Failed"},
		}
	}
}
