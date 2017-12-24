package main

import (
	"./admin"
	"./core"
	"./db"
	"./index"
	"./post"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	sql := db.NewDb("./sqlite.db", "sqlite3")
	iris.RegisterOnInterrupt(func() {
		sql.Close()
	})
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	sql.SetDefaultCacher(cacher)
	sql.Sync2(new(db.PostDb))

	tmpl := iris.HTML("./templates", ".html").Reload(true)
	tmpl.AddFunc("getCore", core.GetCore)
	app.RegisterView(tmpl)
	app.StaticWeb("/assets", "./assets")
	app.Layout("shared/main.html").Controller("/post", new(post.PostController), sql)
	app.Layout("shared/admin.html").Controller("/admin", new(admin.AdminController), sql)
	app.Layout("shared/main.html").Controller("/", new(index.IndexController), sql)

	app.Run(iris.Addr(":8080"), iris.WithOptimizations)
}
