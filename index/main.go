package index

import (
	"github.com/kataras/iris/mvc"
	"../core"
	"../post"
)

type IndexController struct {
	mvc.C
}

func (c *IndexController) Get() mvc.Result {
	return mvc.View{
		Name: "index.html",
		Data: IndexStruct{
			Core:    *(core.GetCore()),
			SubData: *GetDatas(),
		},
	}
}

func GetDatas() *[]interface{} {
	datas := make([]interface{}, 0)
	datas = append(datas, post.PostStruct{
		ID:       1,
		Title:    "THERE'S MEGAN IN TY",
		SubTitle: "THAT'S TRUE",
		Author:   "Black Hat",
		Category: "Megan",
	})
	return &datas
}
