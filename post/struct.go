package post

import "html/template"

//PostStruct is the the structure of post in /post/.
type PostStruct struct {
	ID       int
	Title    string
	SubTitle string
	Author   string
	Category string
	Content  template.HTML
}
