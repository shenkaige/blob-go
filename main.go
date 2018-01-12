package main

import (
	// first, import the standar libraries.
	"flag"
	"fmt"
	"time"

	// second, the subpackages of your own application.
	"github.com/blob-go/blob-go/admin"
	"github.com/blob-go/blob-go/db"
	"github.com/blob-go/blob-go/index"
	"github.com/blob-go/blob-go/post"

	// third, all expect the the web framework itself.
	"github.com/go-xorm/xorm"
	"github.com/gorilla/securecookie"

	// last, the web framework.
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

var themeDic = "./usr/themes/rbreaker/"

func main() {
	var devMode = flag.Bool("dev", false, "Enable dev mode")
	var isCache = flag.Bool("cache", false, "Enable iris cache")
	var port = flag.String("port", "8080", "the port blob listens")
	flag.Parse()

	app := iris.New()
	app.Use(recover.New())

	// if devMode enabled then print a message
	// and add the logger middleware.
	if *devMode {
		app.Logger().SetLevel("debug")
		app.Logger().Infof("DEV MODE ENABLED. DO NOT USE IT IN PRODUCTION AS IT WILL CAUSE SIGNIFICANT PERFORMANCE LAG DUE TO MIDDLEWARES AND TEMPLATE RELOADS.")
		app.Use(logger.New())
	}

	// Register a view engine with its template functions.
	tmpl := iris.HTML(themeDic+"templates", ".html").Reload(*devMode)
	tmpl.AddFunc("getInfo", db.GetInfoFunc(sql))
	tmpl.AddFunc("add", add)
	tmpl.AddFunc("minus", minus)
	// Add a default layout.
	// This will be used to all routes otherwise it's overriden
	// by the party's `.Layout`.
	// Only the "/admin" sub router(party) has a different
	// one, therefore we will change it only there (look below).
	tmpl.Layout("shared/main.html")
	app.RegisterView(tmpl)

	// register custom errors handlers.
	app.OnErrorCode(iris.StatusNotFound, errHandler(iris.StatusNotFound))
	app.OnErrorCode(iris.StatusInternalServerError, errHandler(iris.StatusInternalServerError))

	if *isCache {
		// if cache is true enable the cache middleware,
		// including the statis assets (so register it before the app.StaticWeb).
		// It's root, so all sub routers
		// will make use of this middleware, so it's no need
		// to register it on each party
		// (it was wrong that you did it only for two of them,
		// all of them were using it).
		app.Use(cache.Handler(10 * time.Second))
	}

	// register static assets routes.
	app.StaticWeb("/assets/css", themeDic+"assets/css")
	app.StaticWeb("/assets/js", themeDic+"assets/js")

	// Open the sql database connection.
	sql := db.NewDb("./sqlite.db", "sqlite3")
	defer sql.Close()
	iris.RegisterOnInterrupt(func() {
		sql.Close()
	})
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	sql.SetDefaultCacher(cacher)
	sql.Sync2(new(db.CoreDb))
	sql.Sync2(new(db.UserDb))
	sql.Sync2(new(db.PostDb))

	// Prepare the web modules that may used
	// on many sub routers(parties), like the sessions manager.

	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You need to provide exactly that amount.
	hashKey := []byte("the-big-and-secret-fash-key-here")
	blockKey := []byte("lot-secret-of-characters-big-too")

	sessManager := sessions.New(sessions.Config{
		Cookie:   "blobsess",
		Expires:  24 * time.Hour,
		Encoding: securecookie.New(hashKey, blockKey),
	})

	// Start MVC.

	// All dependencies that are registered with
	// `.Register` on root mvc application will be cloned
	// to its sub apps as well, so no need to register 'sql' more than once.
	mvcApp := mvc.New(app).Register(sql).Handle(new(index.Controller))

	// handle "/post" controller.
	mvcApp.Clone(r.Party("/post")).Handle(new(post.Controller))

	// handle "/admin" controller with sessions
	// and admin.html layout.
	mvcApp.Clone(r.Party("/admin").Layout("shared/admin.html")).
		Register(sessManager.Start).Handle(new(admin.Controller))

	// End MVC.

	// Start the web server.
	app.Run(iris.Addr(":"+*port), iris.WithOptimizations)
}

func errHandler(statusCode int) iris.Handler {
	return func(ctx iris.Context) {
		// this is not necessary, the same layout file is registered as the default layout.
		// ctx.ViewLayout("shared/main.html")
		ctx.View(fmt.Sprintf("httperr/%s.html", statusCode)) // i.e httperr/404.html
	}
}

func add(a, b int) int { return a + b }

func minus(a, b int) int { return a - b }
