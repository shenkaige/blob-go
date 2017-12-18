package post

import (
	"github.com/kataras/iris/mvc"
	"../core"
)

type PostController struct {
	mvc.C
}

func (c *PostController) Get() mvc.Result {
	return mvc.View{
		Name: "post.html",
		Data: PostStruct{
			Core:     *(core.GetCore()),
			Title:    "Hello Page",
			SubTitle: "Welcome to my awesome website",
			Author:   "Black Hat",
			Category: "Megan",
		},
	}
}

func (c *PostController) GetBy(id int) mvc.Result {
	return mvc.View{
		Name: "post.html",
		Data: PostStruct{
			Core:     *(core.GetCore()),
			Title:    "Hello Page",
			SubTitle: "Welcome to my awesome website",
			Author:   "Black Hat",
			Category: "Megan",
		},
	}
}
