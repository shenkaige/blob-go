package main

import (
	"./db"
	"./index"
	"./post"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/go-xorm/xorm"
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

	app.RegisterView(iris.HTML("./templates", ".html").Layout("shared/main.html").Reload(true))

	app.StaticWeb("/assets", "./assets")

	app.Controller("/post", new(post.PostController), sql)
	app.Controller("/", new(index.IndexController), sql)

	app.Run(iris.Addr(":8080"))
}
