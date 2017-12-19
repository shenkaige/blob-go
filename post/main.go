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
			Title:    "THERE'S MEGAN IN TY",
			SubTitle: "THAT'S TRUE!",
			Author:   "Black Hat",
			Category: "Megan",
			Content:  "# MEGAN \n\n Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		},
	}
}

func (c *PostController) GetBy(id int) mvc.Result {
	return mvc.View{
		Name: "post.html",
		Data: PostStruct{
			Core:     *(core.GetCore()),
			Title:    "THERE'S MEGAN IN TY",
			SubTitle: "THAT'S TRUE!",
			Author:   "Black Hat",
			Category: "Megan",
			Content:  "# MEGAN \n\n Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		},
	}
}
