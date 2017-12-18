package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"./post"
	"./index"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.RegisterView(iris.HTML("./templates", ".html").Layout("shared/main.html").Reload(true))

	app.StaticWeb("/assets", "./assets")

	app.Controller("/post", new(post.PostController))

	app.Controller("/", new(index.IndexController))

	app.Run(iris.Addr(":8080"))
}
