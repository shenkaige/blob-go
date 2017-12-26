package db

type CoreDb struct {
	Title    string
	SubTitle string
}

type PostDb struct {
	Id       int `xorm:"pk"`
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
