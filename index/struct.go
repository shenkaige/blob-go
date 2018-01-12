package index

import (
	"github.com/blob-go/blob-go/db"
)

//IndexStruct is the the structure of indexData.
type IndexStruct struct {
	Index     int
	IndexData []db.PostDb
}
