package post

import (
	"../core"
)

type PostStruct struct {
	Core core.CoreStruct
	ID int
	Title string
	SubTitle string
	Author string
	Category string
}
