package index

import (
	"../core"
	"../db"
)

type IndexStruct struct {
	Core    core.CoreStruct
	SubData []db.PostDb
}
