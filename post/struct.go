package post

import "html/template"

type PostStruct struct {
	ID       int
	Title    string
	SubTitle string
	Author   string
	Category string
	Content  template.HTML
}
