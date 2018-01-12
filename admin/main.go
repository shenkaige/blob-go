package admin

import (
	"encoding/hex"
	"github.com/blob-go/blob-go/db"
	"github.com/blob-go/blob-go/index"
	"github.com/blob-go/blob-go/post"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"golang.org/x/crypto/sha3"
	"html/template"
	"strings"
)

var fzfResp = mvc.Response{Code: 404}
var fzzResp = mvc.Response{Code: 500}

//Controller is the controller to /admin page.
type Controller struct {
	Session *sessions.Session
	Ctx     iris.Context
	SQL     *xorm.Engine
}

func (c *Controller) isLogin() bool {
	userID, err1 := c.Session.GetIntDefault("UserID", 0)
	hash := c.Session.GetStringDefault("UserPass", "")
	if err1 == nil && hash != "" {
		ok, _ := db.AuthUserID(userID, hash, c.SQL)
		return ok
	}
	return false
}

func (c *Controller) checkLogin(callback func() mvc.Result) mvc.Result {
	if c.isLogin() {
		return callback()
	}
	return mvc.Response{Path: "/admin/login?notify=you should login first."}
}

//Get is the function when /admin/ is called.
func (c *Controller) Get() mvc.Result {
	return c.checkLogin(func() mvc.Result {
		return c.GetOverview()
	})
}

//GetOverview is the function when /admin/overview/ is called.
func (c *Controller) GetOverview() mvc.Result {
	return c.checkLogin(func() mvc.Result {
		if overviewData, ok := db.GetOverview(c.SQL); ok {
			return mvc.View{
				Name: "admin/overview.html",
				Data: OverviewStruct{
					OverviewData: *overviewData,
				},
			}
		}
		return mvc.View{
			Name: "admin/overview.html",
			Data: OverviewStruct{
				OverviewData: db.OverviewDb{},
			},
		}
	})
}

//GetPost is the function when /admin/post/ is called.
func (c *Controller) GetPost() mvc.Result {
	return c.checkLogin(func() mvc.Result {
		if indexData, ok := db.GetIndexBy(1, c.SQL); ok {
			return mvc.View{
				Name: "admin/post_index.html",
				Data: index.Struct{
					IndexData: *indexData,
				},
			}
		}
		return mvc.View{
			Name: "admin/post_index.html",
			Data: index.Struct{
				IndexData: []db.PostDb{},
			},
		}
	})
}

//GetPostEdit is the function when /admin/post/edit/ is called.
func (c *Controller) GetPostEdit() mvc.Result {
	return c.checkLogin(func() mvc.Result {
		return mvc.View{
			Name: "admin/post_edit.html",
			Data: post.Struct{
				Title:    "New Post",
				SubTitle: "This is a Subtitle",
				Author:   "Author",
				Category: "",
				Content:  "",
			},
		}
	})
}

//GetPostEditBy is the function when /admin/post/edit/<id> is called.
func (c *Controller) GetPostEditBy(id int) mvc.Result {
	return c.checkLogin(func() mvc.Result {
		if postData, ok := db.GetPost(id, c.SQL); ok {
			return mvc.View{
				Name: "admin/post_edit.html",
				Data: post.Struct{
					ID:       postData.Id,
					Title:    postData.Title,
					SubTitle: postData.SubTitle,
					Author:   postData.Author,
					Category: postData.Category,
					Content:  template.HTML(postData.Content),
				},
			}
		}
		return fzfResp
	})
}

//PostPostEditBy is the function when /admin/post/edit/<id> is posted.
func (c *Controller) PostPostEditBy(id int, ctx iris.Context) mvc.Result {
	return c.checkLogin(func() mvc.Result {
		r := strings.NewReplacer("\r\n", "\n")
		postData := db.PostDb{
			Title:    ctx.FormValue("title"),
			SubTitle: ctx.FormValue("sub-title"),
			Author:   ctx.FormValue("author"),
			Category: ctx.FormValue("category"),
			Content:  r.Replace(ctx.FormValue("content")),
		}
		if db.SetPost(id, &postData, c.SQL) {
			return mvc.Response{Path: "/admin/post/"}
		}
		return fzzResp
	})
}

//GetSetting is the function when /admin/setting/ is called.
func (c *Controller) GetSetting() mvc.Result {
	return c.checkLogin(func() mvc.Result {
		return mvc.View{
			Name: "admin/setting.html",
			Data: db.GetInfo(c.SQL),
		}
	})
}

//PostSetting is the function when /admin/setting/ is posted.
func (c *Controller) PostSetting(ctx iris.Context) mvc.Result {
	return c.checkLogin(func() mvc.Result {
		title := ctx.FormValue("title")
		subTitle := ctx.FormValue("sub-title")
		if db.SetCore(title, subTitle, c.SQL) {
			return mvc.Response{Path: "/admin/setting"}
		}
		return fzzResp
	})
}

//GetLogin is the function when /admin/login/ is called.
func (c *Controller) GetLogin(ctx iris.Context) mvc.Result {
	if c.isLogin() {
		return mvc.Response{Path: "/admin"}
	}
	ctx.ViewLayout("shared/logres.html")
	return mvc.View{
		Name: "admin/login.html",
		Data: iris.Map{"Title": "Login"},
	}
}

//PostLogin is the function when /admin/login/ is posted.
func (c *Controller) PostLogin(ctx iris.Context) mvc.Result {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	hasher := sha3.New512()
	hasher.Write([]byte(password))
	hash := hex.EncodeToString(hasher.Sum(nil))

	if ok, id, _ := db.AuthUserName(username, hash, c.SQL); ok {
		c.Session.Set("UserID", id)
		c.Session.Set("UserPass", hash)
		return mvc.Response{Path: "/admin"}
	}
	return mvc.Response{Path: "/admin/login?notify=login failed."}
}

//GetLogout is the function when /admin/logout/ is called.
func (c *Controller) GetLogout() mvc.Result {
	c.Session.Delete("UserID")
	c.Session.Delete("UserPass")
	return mvc.Response{Path: "/admin/login?notify=you are logged out."}
}
