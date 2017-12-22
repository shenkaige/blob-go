package db

type PostDb struct {
	Id       int
	Title    string
	SubTitle string
	Author   string
	Category string
	Content  string
}

type OverviewDb struct {
	PostCount    int
	CommentCount int
}
