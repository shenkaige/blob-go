package admin

import (
	"../db"
	"../index"
	"../post"
	"encoding/hex"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"golang.org/x/crypto/sha3"
	"html/template"
)

//AdminController is the controller to /admin page.
type AdminController struct {
	Session *sessions.Session
	Ctx     iris.Context
	Sql     *xorm.Engine
}

func (c AdminController) IsLogin() bool {
	userid, err1 := c.Session.GetIntDefault("UserID", 0)
	hash := c.Session.GetStringDefault("UserPass", "")
	if err1 == nil && hash != "" {
		ok, _ := db.AuthUserID(userid, hash, c.Sql)
		return ok
	}
	return false
}

//Get is the function when /admin/ is called.
func (c *AdminController) Get() mvc.Result {
	if c.IsLogin() {
		return c.GetOverview()
	}
	return mvc.Response{
		Path: "/admin/login?notify=you should login first.",
	}
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

func (c *AdminController) GetLogin(ctx iris.Context) mvc.Result {
	if c.IsLogin() {
		return mvc.Response{
			Path: "/admin",
		}
	} else {
		ctx.ViewLayout("shared/logres.html")
		return mvc.View{
			Name: "admin/login.html",
			Data: iris.Map{"Title": "Login"},
		}
	}
}

func (c *AdminController) PostLogin(ctx iris.Context) mvc.Result {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	hasher := sha3.New512()
	hasher.Write([]byte(password))
	hash := hex.EncodeToString(hasher.Sum(nil))

	if ok, id, _ := db.AuthUserName(username, hash, c.Sql); ok {
		c.Session.Set("UserID", id)
		c.Session.Set("UserPass", hash)
		return mvc.Response{
			Path: "/admin",
		}
	} else {
		return mvc.Response{
			Path: "/admin/login?notify=login failed.",
		}
	}
}

func (c *AdminController) GetLogout() mvc.Result {
	c.Session.Delete("UserID")
	c.Session.Delete("UserPass")
	return mvc.Response{
		Path: "/admin/login?notify=you are logged out.",
	}
}
