package index

import (
	"../db"
)

//IndexStruct is the the structure of indexData.
type IndexStruct struct {
	Index     int
	IndexData []db.PostDb
}
