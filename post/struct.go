package post

import "html/template"

//Struct is the the structure of post in /post/.
type Struct struct {
	ID       int
	Title    string
	SubTitle string
	Author   string
	Category string
	Content  template.HTML
}
