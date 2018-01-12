package index

import "github.com/blob-go/blob-go/db"

//Struct is the the structure of indexData.
type Struct struct {
	Index     int
	IndexData []db.PostDb
}
