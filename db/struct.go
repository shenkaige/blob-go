package db

//CoreDb defines the structure of core_db.
type CoreDb struct {
	Id       int `xorm:"pk"`
	Title    string
	SubTitle string
}

//UserDb defines the structure of user_db.
type UserDb struct {
	Id       int `xorm:"pk"`
	Username string
	Passwd   string
}

//PostDb defines the structure of post_db.
type PostDb struct {
	Id       int `xorm:"pk"`
	Title    string
	SubTitle string
	Author   string
	Category string
	Content  string
}

//OverviewDb defines the returned overview's structure.
type OverviewDb struct {
	PostCount    int
	CommentCount int
}
