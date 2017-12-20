package main

import (
	"./db"
	"./index"
	"./post"
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
	sql.Sync2(new(db.PostDb))

	//postdb := db.PostDb{Id: 1}
	//
	//sql.Get(&postdb)
	//log.Println(postdb.Title)

	app.RegisterView(iris.HTML("./templates", ".html").Layout("shared/main.html").Reload(true))

	app.StaticWeb("/assets", "./assets")

	app.Controller("/post", new(post.PostController), sql)
	app.Controller("/", new(index.IndexController), sql)

	app.Run(iris.Addr(":8080"))
}
