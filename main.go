package main

import (
	"./admin"
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

	app.RegisterView(iris.HTML("./templates", ".html").Reload(true))
	app.StaticWeb("/assets", "./assets")
	app.Party("/post").Layout("shared/main.html").Controller("/", new(post.PostController), sql)
	app.Party("/admin").Layout("shared/admin.html").Controller("/", new(admin.AdminController), sql)
	app.Party("/").Layout("shared/main.html").Controller("/", new(index.IndexController), sql)

	app.Run(iris.Addr(":8080"), iris.WithOptimizations)
}
