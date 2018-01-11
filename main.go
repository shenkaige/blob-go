package main

import (
	"flag"
	"strconv"

	"./admin"
	"./db"
	"./index"
	"./post"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"time"
)

func main() {
	var devMode = flag.Bool("dev", false, "Enable dev mode")
	var port = flag.Int("port", 8080, "the port blob listens")
	flag.Parse()

	if *devMode {
		println("DEV MODE ENABLED. DO NOT USE IT IN PRODUCTION AS IT WILL CAUSE SIGNIFICANT PERFORMANCE LAG.")
	}

	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	sql := db.NewDb("./sqlite.db", "sqlite3")
	iris.RegisterOnInterrupt(func() {
		sql.Close()
	})
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	sql.SetDefaultCacher(cacher)
	sql.Sync2(new(db.CoreDb))
	sql.Sync2(new(db.UserDb))
	sql.Sync2(new(db.PostDb))

	tmpl := iris.HTML("./templates", ".html").Reload(*devMode)
	getCore := db.GetCoreFunc(sql)
	tmpl.AddFunc("getCore", getCore)
	tmpl.AddFunc("add", add)
	tmpl.AddFunc("minus", minus)

	sessManager := sessions.New(sessions.Config{
		Cookie:  "blobsess",
		Expires: 24 * time.Hour,
	})
	session := sessManager.Start

	app.OnErrorCode(iris.StatusNotFound, fzfHandler)
	app.OnErrorCode(iris.StatusInternalServerError, fzzHandler)

	app.RegisterView(tmpl)
	app.StaticWeb("/assets", "./assets")
	mvc.New(app.Party("/post").Layout("shared/main.html")).Register(sql).Handle(new(post.PostController))
	mvc.New(app.Party("/admin").Layout("shared/admin.html")).Register(sql, session).Handle(new(admin.AdminController))
	mvc.New(app.Party("/").Layout("shared/main.html")).Register(sql).Handle(new(index.IndexController))

	app.Run(iris.Addr(":"+strconv.Itoa(*port)), iris.WithOptimizations)
}

func fzfHandler(ctx iris.Context) {
	ctx.ViewLayout("shared/main.html")
	ctx.View("httperr/404.html")
}

func fzzHandler(ctx iris.Context) {
	ctx.ViewLayout("shared/main.html")
	ctx.View("httperr/500.html")
}

func add(a, b int) int { return a + b }

func minus(a, b int) int { return a - b }
